package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	goredis "github.com/go-redis/redis/v9"
	"github.com/jakob-moeller-cloud/octi-sync-server/api"
	"github.com/jakob-moeller-cloud/octi-sync-server/config"
	"github.com/jakob-moeller-cloud/octi-sync-server/service/redis"
)

// Run will run the HTTP Server.
func Run(ctx context.Context, cfg *config.Config) error {
	startUpContext, cancelStartUpContext := context.WithCancel(ctx)
	defer cancelStartUpContext()

	clients, err := redis.NewClientsWithRegularPing(
		startUpContext, cfg,
		DefaultUniversalClient(),
		DefaultClientMutators("default"),
	)
	if err != nil {
		return fmt.Errorf("error while starting up redis client")
	}

	accounts := &redis.Accounts{Client: clients["default"]}
	cfg.Services.Accounts = accounts
	cfg.Services.Sharing = accounts
	cfg.Services.Modules = &redis.Modules{Client: clients["default"]}
	cfg.Services.Devices = &redis.Devices{Client: clients["default"]}

	// Define server options
	srv := &http.Server{
		Addr:              cfg.Server.Host + ":" + cfg.Server.Port,
		Handler:           api.New(startUpContext, cfg),
		ReadTimeout:       cfg.Server.Timeout.Read,
		ReadHeaderTimeout: cfg.Server.Timeout.Read,
		WriteTimeout:      cfg.Server.Timeout.Write,
		IdleTimeout:       cfg.Server.Timeout.Idle,
	}

	idleConsClosed := make(chan struct{})
	closeServer := func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint // We received an interrupt signal, shut down.
		// Set up a context to allow for graceful server shutdowns in the event
		// of an OS interrupt (defers the cancel just in case)
		ctx, cancel := context.WithTimeout(
			startUpContext,
			cfg.Server.Timeout.Server,
		)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			// Error from closing listeners, or context timeout:
			cfg.Logger.Warn().Msg("server shutdown error: " + err.Error())
		}

		close(idleConsClosed)
	}

	go closeServer()

	// service connections
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		cfg.Logger.Fatal().Msg("listen: %s" + err.Error())
	}

	<-idleConsClosed
	cfg.Logger.Info().Msg("server shut down finished")

	return nil
}

func DefaultClientMutators(identifier string) redis.ClientMutators {
	return redis.ClientMutators{
		identifier: nil,
	}
}

func DefaultUniversalClient() func(cfg *config.Config) goredis.UniversalClient {
	return func(cfg *config.Config) goredis.UniversalClient {
		return goredis.NewUniversalClient(&cfg.Redis.UniversalOptions)
	}
}

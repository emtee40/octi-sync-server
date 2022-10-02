// Code generated by MockGen. DO NOT EDIT.
// Source: metadata.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	service "github.com/jakob-moeller-cloud/octi-sync-server/service"
)

// MockMetadataProvider is a mock of MetadataProvider interface.
type MockMetadataProvider struct {
	ctrl     *gomock.Controller
	recorder *MockMetadataProviderMockRecorder
}

// MockMetadataProviderMockRecorder is the mock recorder for MockMetadataProvider.
type MockMetadataProviderMockRecorder struct {
	mock *MockMetadataProvider
}

// NewMockMetadataProvider creates a new mock instance.
func NewMockMetadataProvider(ctrl *gomock.Controller) *MockMetadataProvider {
	mock := &MockMetadataProvider{ctrl: ctrl}
	mock.recorder = &MockMetadataProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetadataProvider) EXPECT() *MockMetadataProviderMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockMetadataProvider) Get(ctx context.Context, id service.MetadataID) (service.Metadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(service.Metadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockMetadataProviderMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMetadataProvider)(nil).Get), ctx, id)
}

// Set mocks base method.
func (m *MockMetadataProvider) Set(ctx context.Context, meta service.Metadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, meta)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockMetadataProviderMockRecorder) Set(ctx, meta interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockMetadataProvider)(nil).Set), ctx, meta)
}

// MockMetadata is a mock of Metadata interface.
type MockMetadata struct {
	ctrl     *gomock.Controller
	recorder *MockMetadataMockRecorder
}

// MockMetadataMockRecorder is the mock recorder for MockMetadata.
type MockMetadataMockRecorder struct {
	mock *MockMetadata
}

// NewMockMetadata creates a new mock instance.
func NewMockMetadata(ctrl *gomock.Controller) *MockMetadata {
	mock := &MockMetadata{ctrl: ctrl}
	mock.recorder = &MockMetadataMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetadata) EXPECT() *MockMetadataMockRecorder {
	return m.recorder
}

// GetID mocks base method.
func (m *MockMetadata) GetID() service.MetadataID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(service.MetadataID)
	return ret0
}

// GetID indicates an expected call of GetID.
func (mr *MockMetadataMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockMetadata)(nil).GetID))
}

// GetModifiedAt mocks base method.
func (m *MockMetadata) GetModifiedAt() service.ModifiedAt {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModifiedAt")
	ret0, _ := ret[0].(service.ModifiedAt)
	return ret0
}

// GetModifiedAt indicates an expected call of GetModifiedAt.
func (mr *MockMetadataMockRecorder) GetModifiedAt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModifiedAt", reflect.TypeOf((*MockMetadata)(nil).GetModifiedAt))
}
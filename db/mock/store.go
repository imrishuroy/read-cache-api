// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/imrishuroy/read-cache/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/imrishuroy/read-cache-api/db/sqlc"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateCache mocks base method.
func (m *MockStore) CreateCache(arg0 context.Context, arg1 db.CreateCacheParams) (db.Cache, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCache", arg0, arg1)
	ret0, _ := ret[0].(db.Cache)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCache indicates an expected call of CreateCache.
func (mr *MockStoreMockRecorder) CreateCache(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCache", reflect.TypeOf((*MockStore)(nil).CreateCache), arg0, arg1)
}

// DeleteCache mocks base method.
func (m *MockStore) DeleteCache(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCache", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCache indicates an expected call of DeleteCache.
func (mr *MockStoreMockRecorder) DeleteCache(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCache", reflect.TypeOf((*MockStore)(nil).DeleteCache), arg0, arg1)
}

// GetCache mocks base method.
func (m *MockStore) GetCache(arg0 context.Context, arg1 int64) (db.Cache, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCache", arg0, arg1)
	ret0, _ := ret[0].(db.Cache)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCache indicates an expected call of GetCache.
func (mr *MockStoreMockRecorder) GetCache(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCache", reflect.TypeOf((*MockStore)(nil).GetCache), arg0, arg1)
}

// ListCaches mocks base method.
func (m *MockStore) ListCaches(arg0 context.Context, arg1 db.ListCachesParams) ([]db.Cache, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCaches", arg0, arg1)
	ret0, _ := ret[0].([]db.Cache)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCaches indicates an expected call of ListCaches.
func (mr *MockStoreMockRecorder) ListCaches(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCaches", reflect.TypeOf((*MockStore)(nil).ListCaches), arg0, arg1)
}

// UpdateCache mocks base method.
func (m *MockStore) UpdateCache(arg0 context.Context, arg1 db.UpdateCacheParams) (db.Cache, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCache", arg0, arg1)
	ret0, _ := ret[0].(db.Cache)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCache indicates an expected call of UpdateCache.
func (mr *MockStoreMockRecorder) UpdateCache(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCache", reflect.TypeOf((*MockStore)(nil).UpdateCache), arg0, arg1)
}

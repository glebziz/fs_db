// Code generated by MockGen. DO NOT EDIT.
// Source: ../../db/badger/manager.go
//
// Generated by this command:
//
//	mockgen -source ../../db/badger/manager.go -destination mocks/manager_mocks.go -typed true
//
// Package mock_badger is a generated GoMock package.
package mock_badger

import (
	context "context"
	reflect "reflect"

	badger "github.com/glebziz/fs_db/internal/db/badger"
	transactor "github.com/glebziz/fs_db/internal/model/transactor"
	gomock "go.uber.org/mock/gomock"
)

// MockQueryManager is a mock of QueryManager interface.
type MockQueryManager struct {
	ctrl     *gomock.Controller
	recorder *MockQueryManagerMockRecorder
}

// MockQueryManagerMockRecorder is the mock recorder for MockQueryManager.
type MockQueryManagerMockRecorder struct {
	mock *MockQueryManager
}

// NewMockQueryManager creates a new mock instance.
func NewMockQueryManager(ctrl *gomock.Controller) *MockQueryManager {
	mock := &MockQueryManager{ctrl: ctrl}
	mock.recorder = &MockQueryManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueryManager) EXPECT() *MockQueryManagerMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockQueryManager) Delete(key []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockQueryManagerMockRecorder) Delete(key any) *QueryManagerDeleteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockQueryManager)(nil).Delete), key)
	return &QueryManagerDeleteCall{Call: call}
}

// QueryManagerDeleteCall wrap *gomock.Call
type QueryManagerDeleteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *QueryManagerDeleteCall) Return(arg0 error) *QueryManagerDeleteCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *QueryManagerDeleteCall) Do(f func([]byte) error) *QueryManagerDeleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *QueryManagerDeleteCall) DoAndReturn(f func([]byte) error) *QueryManagerDeleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get mocks base method.
func (m *MockQueryManager) Get(key []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockQueryManagerMockRecorder) Get(key any) *QueryManagerGetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockQueryManager)(nil).Get), key)
	return &QueryManagerGetCall{Call: call}
}

// QueryManagerGetCall wrap *gomock.Call
type QueryManagerGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *QueryManagerGetCall) Return(arg0 []byte, arg1 error) *QueryManagerGetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *QueryManagerGetCall) Do(f func([]byte) ([]byte, error)) *QueryManagerGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *QueryManagerGetCall) DoAndReturn(f func([]byte) ([]byte, error)) *QueryManagerGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetAll mocks base method.
func (m *MockQueryManager) GetAll(prefix []byte) ([]badger.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", prefix)
	ret0, _ := ret[0].([]badger.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockQueryManagerMockRecorder) GetAll(prefix any) *QueryManagerGetAllCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockQueryManager)(nil).GetAll), prefix)
	return &QueryManagerGetAllCall{Call: call}
}

// QueryManagerGetAllCall wrap *gomock.Call
type QueryManagerGetAllCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *QueryManagerGetAllCall) Return(arg0 []badger.Item, arg1 error) *QueryManagerGetAllCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *QueryManagerGetAllCall) Do(f func([]byte) ([]badger.Item, error)) *QueryManagerGetAllCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *QueryManagerGetAllCall) DoAndReturn(f func([]byte) ([]badger.Item, error)) *QueryManagerGetAllCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Set mocks base method.
func (m *MockQueryManager) Set(key, val []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", key, val)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockQueryManagerMockRecorder) Set(key, val any) *QueryManagerSetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockQueryManager)(nil).Set), key, val)
	return &QueryManagerSetCall{Call: call}
}

// QueryManagerSetCall wrap *gomock.Call
type QueryManagerSetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *QueryManagerSetCall) Return(arg0 error) *QueryManagerSetCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *QueryManagerSetCall) Do(f func([]byte, []byte) error) *QueryManagerSetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *QueryManagerSetCall) DoAndReturn(f func([]byte, []byte) error) *QueryManagerSetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockProvider is a mock of Provider interface.
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider.
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance.
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// DB mocks base method.
func (m *MockProvider) DB(ctx context.Context) badger.QueryManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DB", ctx)
	ret0, _ := ret[0].(badger.QueryManager)
	return ret0
}

// DB indicates an expected call of DB.
func (mr *MockProviderMockRecorder) DB(ctx any) *ProviderDBCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*MockProvider)(nil).DB), ctx)
	return &ProviderDBCall{Call: call}
}

// ProviderDBCall wrap *gomock.Call
type ProviderDBCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ProviderDBCall) Return(arg0 badger.QueryManager) *ProviderDBCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ProviderDBCall) Do(f func(context.Context) badger.QueryManager) *ProviderDBCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ProviderDBCall) DoAndReturn(f func(context.Context) badger.QueryManager) *ProviderDBCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// RunTransaction mocks base method.
func (m *MockProvider) RunTransaction(ctx context.Context, fn transactor.TransactionFn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunTransaction", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunTransaction indicates an expected call of RunTransaction.
func (mr *MockProviderMockRecorder) RunTransaction(ctx, fn any) *ProviderRunTransactionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunTransaction", reflect.TypeOf((*MockProvider)(nil).RunTransaction), ctx, fn)
	return &ProviderRunTransactionCall{Call: call}
}

// ProviderRunTransactionCall wrap *gomock.Call
type ProviderRunTransactionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ProviderRunTransactionCall) Return(arg0 error) *ProviderRunTransactionCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ProviderRunTransactionCall) Do(f func(context.Context, transactor.TransactionFn) error) *ProviderRunTransactionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ProviderRunTransactionCall) DoAndReturn(f func(context.Context, transactor.TransactionFn) error) *ProviderRunTransactionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

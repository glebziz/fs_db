// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go
//
// Generated by this command:
//
//	mockgen -source usecase.go -destination mocks/mocks.go -typed true
//
// Package mock_transaction is a generated GoMock package.
package mock_transaction

import (
	context "context"
	reflect "reflect"

	model "github.com/glebziz/fs_db/internal/model"
	gomock "go.uber.org/mock/gomock"
)

// Mockcleaner is a mock of cleaner interface.
type Mockcleaner struct {
	ctrl     *gomock.Controller
	recorder *MockcleanerMockRecorder
}

// MockcleanerMockRecorder is the mock recorder for Mockcleaner.
type MockcleanerMockRecorder struct {
	mock *Mockcleaner
}

// NewMockcleaner creates a new mock instance.
func NewMockcleaner(ctrl *gomock.Controller) *Mockcleaner {
	mock := &Mockcleaner{ctrl: ctrl}
	mock.recorder = &MockcleanerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockcleaner) EXPECT() *MockcleanerMockRecorder {
	return m.recorder
}

// Send mocks base method.
func (m *Mockcleaner) Send(contentIds []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Send", contentIds)
}

// Send indicates an expected call of Send.
func (mr *MockcleanerMockRecorder) Send(contentIds any) *cleanerSendCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*Mockcleaner)(nil).Send), contentIds)
	return &cleanerSendCall{Call: call}
}

// cleanerSendCall wrap *gomock.Call
type cleanerSendCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *cleanerSendCall) Return() *cleanerSendCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *cleanerSendCall) Do(f func([]string)) *cleanerSendCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *cleanerSendCall) DoAndReturn(f func([]string)) *cleanerSendCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockfileRepository is a mock of fileRepository interface.
type MockfileRepository struct {
	ctrl     *gomock.Controller
	recorder *MockfileRepositoryMockRecorder
}

// MockfileRepositoryMockRecorder is the mock recorder for MockfileRepository.
type MockfileRepositoryMockRecorder struct {
	mock *MockfileRepository
}

// NewMockfileRepository creates a new mock instance.
func NewMockfileRepository(ctrl *gomock.Controller) *MockfileRepository {
	mock := &MockfileRepository{ctrl: ctrl}
	mock.recorder = &MockfileRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockfileRepository) EXPECT() *MockfileRepositoryMockRecorder {
	return m.recorder
}

// HardDelete mocks base method.
func (m *MockfileRepository) HardDelete(ctx context.Context, txId string, filter *model.FileFilter) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HardDelete", ctx, txId, filter)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HardDelete indicates an expected call of HardDelete.
func (mr *MockfileRepositoryMockRecorder) HardDelete(ctx, txId, filter any) *fileRepositoryHardDeleteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HardDelete", reflect.TypeOf((*MockfileRepository)(nil).HardDelete), ctx, txId, filter)
	return &fileRepositoryHardDeleteCall{Call: call}
}

// fileRepositoryHardDeleteCall wrap *gomock.Call
type fileRepositoryHardDeleteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *fileRepositoryHardDeleteCall) Return(contentIds []string, err error) *fileRepositoryHardDeleteCall {
	c.Call = c.Call.Return(contentIds, err)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *fileRepositoryHardDeleteCall) Do(f func(context.Context, string, *model.FileFilter) ([]string, error)) *fileRepositoryHardDeleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *fileRepositoryHardDeleteCall) DoAndReturn(f func(context.Context, string, *model.FileFilter) ([]string, error)) *fileRepositoryHardDeleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateTx mocks base method.
func (m *MockfileRepository) UpdateTx(ctx context.Context, oldTxId, newTxId string, filter *model.FileFilter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTx", ctx, oldTxId, newTxId, filter)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTx indicates an expected call of UpdateTx.
func (mr *MockfileRepositoryMockRecorder) UpdateTx(ctx, oldTxId, newTxId, filter any) *fileRepositoryUpdateTxCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTx", reflect.TypeOf((*MockfileRepository)(nil).UpdateTx), ctx, oldTxId, newTxId, filter)
	return &fileRepositoryUpdateTxCall{Call: call}
}

// fileRepositoryUpdateTxCall wrap *gomock.Call
type fileRepositoryUpdateTxCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *fileRepositoryUpdateTxCall) Return(arg0 error) *fileRepositoryUpdateTxCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *fileRepositoryUpdateTxCall) Do(f func(context.Context, string, string, *model.FileFilter) error) *fileRepositoryUpdateTxCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *fileRepositoryUpdateTxCall) DoAndReturn(f func(context.Context, string, string, *model.FileFilter) error) *fileRepositoryUpdateTxCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MocktxRepository is a mock of txRepository interface.
type MocktxRepository struct {
	ctrl     *gomock.Controller
	recorder *MocktxRepositoryMockRecorder
}

// MocktxRepositoryMockRecorder is the mock recorder for MocktxRepository.
type MocktxRepositoryMockRecorder struct {
	mock *MocktxRepository
}

// NewMocktxRepository creates a new mock instance.
func NewMocktxRepository(ctrl *gomock.Controller) *MocktxRepository {
	mock := &MocktxRepository{ctrl: ctrl}
	mock.recorder = &MocktxRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocktxRepository) EXPECT() *MocktxRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MocktxRepository) Delete(ctx context.Context, id string) (*model.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(*model.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MocktxRepositoryMockRecorder) Delete(ctx, id any) *txRepositoryDeleteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MocktxRepository)(nil).Delete), ctx, id)
	return &txRepositoryDeleteCall{Call: call}
}

// txRepositoryDeleteCall wrap *gomock.Call
type txRepositoryDeleteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *txRepositoryDeleteCall) Return(arg0 *model.Transaction, arg1 error) *txRepositoryDeleteCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *txRepositoryDeleteCall) Do(f func(context.Context, string) (*model.Transaction, error)) *txRepositoryDeleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *txRepositoryDeleteCall) DoAndReturn(f func(context.Context, string) (*model.Transaction, error)) *txRepositoryDeleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Store mocks base method.
func (m *MocktxRepository) Store(ctx context.Context, tx model.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", ctx, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MocktxRepositoryMockRecorder) Store(ctx, tx any) *txRepositoryStoreCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MocktxRepository)(nil).Store), ctx, tx)
	return &txRepositoryStoreCall{Call: call}
}

// txRepositoryStoreCall wrap *gomock.Call
type txRepositoryStoreCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *txRepositoryStoreCall) Return(arg0 error) *txRepositoryStoreCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *txRepositoryStoreCall) Do(f func(context.Context, model.Transaction) error) *txRepositoryStoreCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *txRepositoryStoreCall) DoAndReturn(f func(context.Context, model.Transaction) error) *txRepositoryStoreCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Mockgenerator is a mock of generator interface.
type Mockgenerator struct {
	ctrl     *gomock.Controller
	recorder *MockgeneratorMockRecorder
}

// MockgeneratorMockRecorder is the mock recorder for Mockgenerator.
type MockgeneratorMockRecorder struct {
	mock *Mockgenerator
}

// NewMockgenerator creates a new mock instance.
func NewMockgenerator(ctrl *gomock.Controller) *Mockgenerator {
	mock := &Mockgenerator{ctrl: ctrl}
	mock.recorder = &MockgeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockgenerator) EXPECT() *MockgeneratorMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *Mockgenerator) Generate() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate")
	ret0, _ := ret[0].(string)
	return ret0
}

// Generate indicates an expected call of Generate.
func (mr *MockgeneratorMockRecorder) Generate() *generatorGenerateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*Mockgenerator)(nil).Generate))
	return &generatorGenerateCall{Call: call}
}

// generatorGenerateCall wrap *gomock.Call
type generatorGenerateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *generatorGenerateCall) Return(arg0 string) *generatorGenerateCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *generatorGenerateCall) Do(f func() string) *generatorGenerateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *generatorGenerateCall) DoAndReturn(f func() string) *generatorGenerateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

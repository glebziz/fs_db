// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go
//
// Generated by this command:
//
//	mockgen -source usecase.go -destination mocks/mocks.go -typed true
//
// Package mock_cleaner is a generated GoMock package.
package mock_cleaner

import (
	context "context"
	reflect "reflect"

	model "github.com/glebziz/fs_db/internal/model"
	sequence "github.com/glebziz/fs_db/internal/model/sequence"
	wpool "github.com/glebziz/fs_db/internal/utils/wpool"
	gomock "go.uber.org/mock/gomock"
)

// Mockcore is a mock of core interface.
type Mockcore struct {
	ctrl     *gomock.Controller
	recorder *MockcoreMockRecorder
}

// MockcoreMockRecorder is the mock recorder for Mockcore.
type MockcoreMockRecorder struct {
	mock *Mockcore
}

// NewMockcore creates a new mock instance.
func NewMockcore(ctrl *gomock.Controller) *Mockcore {
	mock := &Mockcore{ctrl: ctrl}
	mock.recorder = &MockcoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockcore) EXPECT() *MockcoreMockRecorder {
	return m.recorder
}

// DeleteOld mocks base method.
func (m *Mockcore) DeleteOld(ctx context.Context, txId string, beforeSeq sequence.Seq) []model.File {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOld", ctx, txId, beforeSeq)
	ret0, _ := ret[0].([]model.File)
	return ret0
}

// DeleteOld indicates an expected call of DeleteOld.
func (mr *MockcoreMockRecorder) DeleteOld(ctx, txId, beforeSeq any) *coreDeleteOldCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOld", reflect.TypeOf((*Mockcore)(nil).DeleteOld), ctx, txId, beforeSeq)
	return &coreDeleteOldCall{Call: call}
}

// coreDeleteOldCall wrap *gomock.Call
type coreDeleteOldCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *coreDeleteOldCall) Return(arg0 []model.File) *coreDeleteOldCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *coreDeleteOldCall) Do(f func(context.Context, string, sequence.Seq) []model.File) *coreDeleteOldCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *coreDeleteOldCall) DoAndReturn(f func(context.Context, string, sequence.Seq) []model.File) *coreDeleteOldCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockcontentRepository is a mock of contentRepository interface.
type MockcontentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockcontentRepositoryMockRecorder
}

// MockcontentRepositoryMockRecorder is the mock recorder for MockcontentRepository.
type MockcontentRepositoryMockRecorder struct {
	mock *MockcontentRepository
}

// NewMockcontentRepository creates a new mock instance.
func NewMockcontentRepository(ctrl *gomock.Controller) *MockcontentRepository {
	mock := &MockcontentRepository{ctrl: ctrl}
	mock.recorder = &MockcontentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockcontentRepository) EXPECT() *MockcontentRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockcontentRepository) Delete(ctx context.Context, path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, path)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockcontentRepositoryMockRecorder) Delete(ctx, path any) *contentRepositoryDeleteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockcontentRepository)(nil).Delete), ctx, path)
	return &contentRepositoryDeleteCall{Call: call}
}

// contentRepositoryDeleteCall wrap *gomock.Call
type contentRepositoryDeleteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *contentRepositoryDeleteCall) Return(arg0 error) *contentRepositoryDeleteCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *contentRepositoryDeleteCall) Do(f func(context.Context, string) error) *contentRepositoryDeleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *contentRepositoryDeleteCall) DoAndReturn(f func(context.Context, string) error) *contentRepositoryDeleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MockcontentFileRepository is a mock of contentFileRepository interface.
type MockcontentFileRepository struct {
	ctrl     *gomock.Controller
	recorder *MockcontentFileRepositoryMockRecorder
}

// MockcontentFileRepositoryMockRecorder is the mock recorder for MockcontentFileRepository.
type MockcontentFileRepositoryMockRecorder struct {
	mock *MockcontentFileRepository
}

// NewMockcontentFileRepository creates a new mock instance.
func NewMockcontentFileRepository(ctrl *gomock.Controller) *MockcontentFileRepository {
	mock := &MockcontentFileRepository{ctrl: ctrl}
	mock.recorder = &MockcontentFileRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockcontentFileRepository) EXPECT() *MockcontentFileRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockcontentFileRepository) Delete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockcontentFileRepositoryMockRecorder) Delete(ctx, id any) *contentFileRepositoryDeleteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockcontentFileRepository)(nil).Delete), ctx, id)
	return &contentFileRepositoryDeleteCall{Call: call}
}

// contentFileRepositoryDeleteCall wrap *gomock.Call
type contentFileRepositoryDeleteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *contentFileRepositoryDeleteCall) Return(arg0 error) *contentFileRepositoryDeleteCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *contentFileRepositoryDeleteCall) Do(f func(context.Context, string) error) *contentFileRepositoryDeleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *contentFileRepositoryDeleteCall) DoAndReturn(f func(context.Context, string) error) *contentFileRepositoryDeleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get mocks base method.
func (m *MockcontentFileRepository) Get(ctx context.Context, id string) (model.ContentFile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(model.ContentFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockcontentFileRepositoryMockRecorder) Get(ctx, id any) *contentFileRepositoryGetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockcontentFileRepository)(nil).Get), ctx, id)
	return &contentFileRepositoryGetCall{Call: call}
}

// contentFileRepositoryGetCall wrap *gomock.Call
type contentFileRepositoryGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *contentFileRepositoryGetCall) Return(arg0 model.ContentFile, arg1 error) *contentFileRepositoryGetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *contentFileRepositoryGetCall) Do(f func(context.Context, string) (model.ContentFile, error)) *contentFileRepositoryGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *contentFileRepositoryGetCall) DoAndReturn(f func(context.Context, string) (model.ContentFile, error)) *contentFileRepositoryGetCall {
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

// Delete mocks base method.
func (m *MockfileRepository) Delete(ctx context.Context, file model.File) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, file)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockfileRepositoryMockRecorder) Delete(ctx, file any) *fileRepositoryDeleteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockfileRepository)(nil).Delete), ctx, file)
	return &fileRepositoryDeleteCall{Call: call}
}

// fileRepositoryDeleteCall wrap *gomock.Call
type fileRepositoryDeleteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *fileRepositoryDeleteCall) Return(arg0 error) *fileRepositoryDeleteCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *fileRepositoryDeleteCall) Do(f func(context.Context, model.File) error) *fileRepositoryDeleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *fileRepositoryDeleteCall) DoAndReturn(f func(context.Context, model.File) error) *fileRepositoryDeleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Mocksender is a mock of sender interface.
type Mocksender struct {
	ctrl     *gomock.Controller
	recorder *MocksenderMockRecorder
}

// MocksenderMockRecorder is the mock recorder for Mocksender.
type MocksenderMockRecorder struct {
	mock *Mocksender
}

// NewMocksender creates a new mock instance.
func NewMocksender(ctrl *gomock.Controller) *Mocksender {
	mock := &Mocksender{ctrl: ctrl}
	mock.recorder = &MocksenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mocksender) EXPECT() *MocksenderMockRecorder {
	return m.recorder
}

// Send mocks base method.
func (m *Mocksender) Send(ctx context.Context, event wpool.Event) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Send", ctx, event)
}

// Send indicates an expected call of Send.
func (mr *MocksenderMockRecorder) Send(ctx, event any) *senderSendCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*Mocksender)(nil).Send), ctx, event)
	return &senderSendCall{Call: call}
}

// senderSendCall wrap *gomock.Call
type senderSendCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *senderSendCall) Return() *senderSendCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *senderSendCall) Do(f func(context.Context, wpool.Event)) *senderSendCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *senderSendCall) DoAndReturn(f func(context.Context, wpool.Event)) *senderSendCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MocktransactionRepository is a mock of transactionRepository interface.
type MocktransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MocktransactionRepositoryMockRecorder
}

// MocktransactionRepositoryMockRecorder is the mock recorder for MocktransactionRepository.
type MocktransactionRepositoryMockRecorder struct {
	mock *MocktransactionRepository
}

// NewMocktransactionRepository creates a new mock instance.
func NewMocktransactionRepository(ctrl *gomock.Controller) *MocktransactionRepository {
	mock := &MocktransactionRepository{ctrl: ctrl}
	mock.recorder = &MocktransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocktransactionRepository) EXPECT() *MocktransactionRepositoryMockRecorder {
	return m.recorder
}

// Oldest mocks base method.
func (m *MocktransactionRepository) Oldest(ctx context.Context) (model.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Oldest", ctx)
	ret0, _ := ret[0].(model.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Oldest indicates an expected call of Oldest.
func (mr *MocktransactionRepositoryMockRecorder) Oldest(ctx any) *transactionRepositoryOldestCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Oldest", reflect.TypeOf((*MocktransactionRepository)(nil).Oldest), ctx)
	return &transactionRepositoryOldestCall{Call: call}
}

// transactionRepositoryOldestCall wrap *gomock.Call
type transactionRepositoryOldestCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *transactionRepositoryOldestCall) Return(arg0 model.Transaction, arg1 error) *transactionRepositoryOldestCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *transactionRepositoryOldestCall) Do(f func(context.Context) (model.Transaction, error)) *transactionRepositoryOldestCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *transactionRepositoryOldestCall) DoAndReturn(f func(context.Context) (model.Transaction, error)) *transactionRepositoryOldestCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

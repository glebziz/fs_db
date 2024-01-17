// Code generated by MockGen. DO NOT EDIT.
// Source: cleaner.go
//
// Generated by this command:
//
//	mockgen -source cleaner.go -destination mocks/mocks.go -typed true
//
// Package mock_cleaner is a generated GoMock package.
package mock_cleaner

import (
	context "context"
	reflect "reflect"

	model "github.com/glebziz/fs_db/internal/model"
	quartz "github.com/reugn/go-quartz/quartz"
	gomock "go.uber.org/mock/gomock"
)

// Mockscheduler is a mock of scheduler interface.
type Mockscheduler struct {
	ctrl     *gomock.Controller
	recorder *MockschedulerMockRecorder
}

// MockschedulerMockRecorder is the mock recorder for Mockscheduler.
type MockschedulerMockRecorder struct {
	mock *Mockscheduler
}

// NewMockscheduler creates a new mock instance.
func NewMockscheduler(ctrl *gomock.Controller) *Mockscheduler {
	mock := &Mockscheduler{ctrl: ctrl}
	mock.recorder = &MockschedulerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockscheduler) EXPECT() *MockschedulerMockRecorder {
	return m.recorder
}

// Clear mocks base method.
func (m *Mockscheduler) Clear() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clear")
	ret0, _ := ret[0].(error)
	return ret0
}

// Clear indicates an expected call of Clear.
func (mr *MockschedulerMockRecorder) Clear() *schedulerClearCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*Mockscheduler)(nil).Clear))
	return &schedulerClearCall{Call: call}
}

// schedulerClearCall wrap *gomock.Call
type schedulerClearCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *schedulerClearCall) Return(arg0 error) *schedulerClearCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *schedulerClearCall) Do(f func() error) *schedulerClearCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *schedulerClearCall) DoAndReturn(f func() error) *schedulerClearCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ScheduleJob mocks base method.
func (m *Mockscheduler) ScheduleJob(jobDetail *quartz.JobDetail, trigger quartz.Trigger) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScheduleJob", jobDetail, trigger)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScheduleJob indicates an expected call of ScheduleJob.
func (mr *MockschedulerMockRecorder) ScheduleJob(jobDetail, trigger any) *schedulerScheduleJobCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleJob", reflect.TypeOf((*Mockscheduler)(nil).ScheduleJob), jobDetail, trigger)
	return &schedulerScheduleJobCall{Call: call}
}

// schedulerScheduleJobCall wrap *gomock.Call
type schedulerScheduleJobCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *schedulerScheduleJobCall) Return(arg0 error) *schedulerScheduleJobCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *schedulerScheduleJobCall) Do(f func(*quartz.JobDetail, quartz.Trigger) error) *schedulerScheduleJobCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *schedulerScheduleJobCall) DoAndReturn(f func(*quartz.JobDetail, quartz.Trigger) error) *schedulerScheduleJobCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Start mocks base method.
func (m *Mockscheduler) Start(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start", arg0)
}

// Start indicates an expected call of Start.
func (mr *MockschedulerMockRecorder) Start(arg0 any) *schedulerStartCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*Mockscheduler)(nil).Start), arg0)
	return &schedulerStartCall{Call: call}
}

// schedulerStartCall wrap *gomock.Call
type schedulerStartCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *schedulerStartCall) Return() *schedulerStartCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *schedulerStartCall) Do(f func(context.Context)) *schedulerStartCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *schedulerStartCall) DoAndReturn(f func(context.Context)) *schedulerStartCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Stop mocks base method.
func (m *Mockscheduler) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockschedulerMockRecorder) Stop() *schedulerStopCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*Mockscheduler)(nil).Stop))
	return &schedulerStopCall{Call: call}
}

// schedulerStopCall wrap *gomock.Call
type schedulerStopCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *schedulerStopCall) Return() *schedulerStopCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *schedulerStopCall) Do(f func()) *schedulerStopCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *schedulerStopCall) DoAndReturn(f func()) *schedulerStopCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Wait mocks base method.
func (m *Mockscheduler) Wait(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Wait", arg0)
}

// Wait indicates an expected call of Wait.
func (mr *MockschedulerMockRecorder) Wait(arg0 any) *schedulerWaitCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*Mockscheduler)(nil).Wait), arg0)
	return &schedulerWaitCall{Call: call}
}

// schedulerWaitCall wrap *gomock.Call
type schedulerWaitCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *schedulerWaitCall) Return() *schedulerWaitCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *schedulerWaitCall) Do(f func(context.Context)) *schedulerWaitCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *schedulerWaitCall) DoAndReturn(f func(context.Context)) *schedulerWaitCall {
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
func (m *MockcontentFileRepository) Delete(ctx context.Context, ids []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, ids)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockcontentFileRepositoryMockRecorder) Delete(ctx, ids any) *contentFileRepositoryDeleteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockcontentFileRepository)(nil).Delete), ctx, ids)
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
func (c *contentFileRepositoryDeleteCall) Do(f func(context.Context, []string) error) *contentFileRepositoryDeleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *contentFileRepositoryDeleteCall) DoAndReturn(f func(context.Context, []string) error) *contentFileRepositoryDeleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetIn mocks base method.
func (m *MockcontentFileRepository) GetIn(ctx context.Context, ids []string) ([]model.ContentFile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIn", ctx, ids)
	ret0, _ := ret[0].([]model.ContentFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIn indicates an expected call of GetIn.
func (mr *MockcontentFileRepositoryMockRecorder) GetIn(ctx, ids any) *contentFileRepositoryGetInCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIn", reflect.TypeOf((*MockcontentFileRepository)(nil).GetIn), ctx, ids)
	return &contentFileRepositoryGetInCall{Call: call}
}

// contentFileRepositoryGetInCall wrap *gomock.Call
type contentFileRepositoryGetInCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *contentFileRepositoryGetInCall) Return(arg0 []model.ContentFile, arg1 error) *contentFileRepositoryGetInCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *contentFileRepositoryGetInCall) Do(f func(context.Context, []string) ([]model.ContentFile, error)) *contentFileRepositoryGetInCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *contentFileRepositoryGetInCall) DoAndReturn(f func(context.Context, []string) ([]model.ContentFile, error)) *contentFileRepositoryGetInCall {
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
func (m *MocktransactionRepository) Oldest(ctx context.Context) (*model.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Oldest", ctx)
	ret0, _ := ret[0].(*model.Transaction)
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
func (c *transactionRepositoryOldestCall) Return(arg0 *model.Transaction, arg1 error) *transactionRepositoryOldestCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *transactionRepositoryOldestCall) Do(f func(context.Context) (*model.Transaction, error)) *transactionRepositoryOldestCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *transactionRepositoryOldestCall) DoAndReturn(f func(context.Context) (*model.Transaction, error)) *transactionRepositoryOldestCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

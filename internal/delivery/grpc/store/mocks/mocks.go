// Code generated by MockGen. DO NOT EDIT.
// Source: service.go
//
// Generated by this command:
//
//	mockgen -source service.go -destination mocks/mocks.go -typed true
//
// Package mock_store is a generated GoMock package.
package mock_store

import (
	context "context"
	model "github.com/glebziz/fs_db/internal/model"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUsecase is a mock of Usecase interface.
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase.
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance.
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockUsecase) Delete(ctx context.Context, key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUsecaseMockRecorder) Delete(ctx, key any) *UsecaseDeleteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUsecase)(nil).Delete), ctx, key)
	return &UsecaseDeleteCall{Call: call}
}

// UsecaseDeleteCall wrap *gomock.Call
type UsecaseDeleteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *UsecaseDeleteCall) Return(arg0 error) *UsecaseDeleteCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *UsecaseDeleteCall) Do(f func(context.Context, string) error) *UsecaseDeleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *UsecaseDeleteCall) DoAndReturn(f func(context.Context, string) error) *UsecaseDeleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get mocks base method.
func (m *MockUsecase) Get(ctx context.Context, key string) (*model.Content, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(*model.Content)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUsecaseMockRecorder) Get(ctx, key any) *UsecaseGetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUsecase)(nil).Get), ctx, key)
	return &UsecaseGetCall{Call: call}
}

// UsecaseGetCall wrap *gomock.Call
type UsecaseGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *UsecaseGetCall) Return(arg0 *model.Content, arg1 error) *UsecaseGetCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *UsecaseGetCall) Do(f func(context.Context, string) (*model.Content, error)) *UsecaseGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *UsecaseGetCall) DoAndReturn(f func(context.Context, string) (*model.Content, error)) *UsecaseGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Set mocks base method.
func (m *MockUsecase) Set(ctx context.Context, key string, content *model.Content) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, key, content)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockUsecaseMockRecorder) Set(ctx, key, content any) *UsecaseSetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockUsecase)(nil).Set), ctx, key, content)
	return &UsecaseSetCall{Call: call}
}

// UsecaseSetCall wrap *gomock.Call
type UsecaseSetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *UsecaseSetCall) Return(arg0 error) *UsecaseSetCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *UsecaseSetCall) Do(f func(context.Context, string, *model.Content) error) *UsecaseSetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *UsecaseSetCall) DoAndReturn(f func(context.Context, string, *model.Content) error) *UsecaseSetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

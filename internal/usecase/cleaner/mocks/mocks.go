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
	gomock "go.uber.org/mock/gomock"
)

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
func (m *MockcontentFileRepository) Delete(ctx context.Context, ids []string) ([]model.ContentFile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, ids)
	ret0, _ := ret[0].([]model.ContentFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
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
func (c *contentFileRepositoryDeleteCall) Return(arg0 []model.ContentFile, arg1 error) *contentFileRepositoryDeleteCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *contentFileRepositoryDeleteCall) Do(f func(context.Context, []string) ([]model.ContentFile, error)) *contentFileRepositoryDeleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *contentFileRepositoryDeleteCall) DoAndReturn(f func(context.Context, []string) ([]model.ContentFile, error)) *contentFileRepositoryDeleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

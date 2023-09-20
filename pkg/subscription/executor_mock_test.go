// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/wundergraph/graphql-go-tools/pkg/subscription (interfaces: Executor,ExecutorPool)

// Package subscription is a generated GoMock package.
package subscription

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	ast "github.com/wundergraph/graphql-go-tools/pkg/ast"
	resolve "github.com/wundergraph/graphql-go-tools/pkg/engine/resolve"
)

// MockExecutor is a mock of Executor interface.
type MockExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockExecutorMockRecorder
}

// MockExecutorMockRecorder is the mock recorder for MockExecutor.
type MockExecutorMockRecorder struct {
	mock *MockExecutor
}

// NewMockExecutor creates a new mock instance.
func NewMockExecutor(ctrl *gomock.Controller) *MockExecutor {
	mock := &MockExecutor{ctrl: ctrl}
	mock.recorder = &MockExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExecutor) EXPECT() *MockExecutorMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockExecutor) Execute(arg0 resolve.FlushWriter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockExecutorMockRecorder) Execute(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockExecutor)(nil).Execute), arg0)
}

// OperationType mocks base method.
func (m *MockExecutor) OperationType() ast.OperationType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OperationType")
	ret0, _ := ret[0].(ast.OperationType)
	return ret0
}

// OperationType indicates an expected call of OperationType.
func (mr *MockExecutorMockRecorder) OperationType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OperationType", reflect.TypeOf((*MockExecutor)(nil).OperationType))
}

// Reset mocks base method.
func (m *MockExecutor) Reset() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset.
func (mr *MockExecutorMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockExecutor)(nil).Reset))
}

// SetContext mocks base method.
func (m *MockExecutor) SetContext(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetContext", arg0)
}

// SetContext indicates an expected call of SetContext.
func (mr *MockExecutorMockRecorder) SetContext(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetContext", reflect.TypeOf((*MockExecutor)(nil).SetContext), arg0)
}

// MockExecutorPool is a mock of ExecutorPool interface.
type MockExecutorPool struct {
	ctrl     *gomock.Controller
	recorder *MockExecutorPoolMockRecorder
}

// MockExecutorPoolMockRecorder is the mock recorder for MockExecutorPool.
type MockExecutorPoolMockRecorder struct {
	mock *MockExecutorPool
}

// NewMockExecutorPool creates a new mock instance.
func NewMockExecutorPool(ctrl *gomock.Controller) *MockExecutorPool {
	mock := &MockExecutorPool{ctrl: ctrl}
	mock.recorder = &MockExecutorPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExecutorPool) EXPECT() *MockExecutorPoolMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockExecutorPool) Get(arg0 []byte) (Executor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(Executor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockExecutorPoolMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockExecutorPool)(nil).Get), arg0)
}

// Put mocks base method.
func (m *MockExecutorPool) Put(arg0 Executor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockExecutorPoolMockRecorder) Put(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockExecutorPool)(nil).Put), arg0)
}
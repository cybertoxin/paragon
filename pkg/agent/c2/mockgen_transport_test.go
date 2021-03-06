// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kcarretto/paragon/pkg/agent/transport (interfaces: ServerMessageWriter)

// Package c2_test is a generated GoMock package.
package c2_test

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	transport "github.com/kcarretto/paragon/pkg/agent/transport"
	reflect "reflect"
)

// MockServerMessageWriter is a mock of ServerMessageWriter interface
type MockServerMessageWriter struct {
	ctrl     *gomock.Controller
	recorder *MockServerMessageWriterMockRecorder
}

// MockServerMessageWriterMockRecorder is the mock recorder for MockServerMessageWriter
type MockServerMessageWriterMockRecorder struct {
	mock *MockServerMessageWriter
}

// NewMockServerMessageWriter creates a new mock instance
func NewMockServerMessageWriter(ctrl *gomock.Controller) *MockServerMessageWriter {
	mock := &MockServerMessageWriter{ctrl: ctrl}
	mock.recorder = &MockServerMessageWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockServerMessageWriter) EXPECT() *MockServerMessageWriterMockRecorder {
	return m.recorder
}

// WriteServerMessage mocks base method
func (m *MockServerMessageWriter) WriteServerMessage(arg0 context.Context, arg1 transport.ServerMessage) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WriteServerMessage", arg0, arg1)
}

// WriteServerMessage indicates an expected call of WriteServerMessage
func (mr *MockServerMessageWriterMockRecorder) WriteServerMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteServerMessage", reflect.TypeOf((*MockServerMessageWriter)(nil).WriteServerMessage), arg0, arg1)
}

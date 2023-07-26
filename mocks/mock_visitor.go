// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rohit-tambe/mockgen-example/party (interfaces: VisitorLister,Greeter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	party "github.com/rohit-tambe/mockgen-example/party"
)

// MockVisitorLister is a mock of VisitorLister interface.
type MockVisitorLister struct {
	ctrl     *gomock.Controller
	recorder *MockVisitorListerMockRecorder
}

// MockVisitorListerMockRecorder is the mock recorder for MockVisitorLister.
type MockVisitorListerMockRecorder struct {
	mock *MockVisitorLister
}

// NewMockVisitorLister creates a new mock instance.
func NewMockVisitorLister(ctrl *gomock.Controller) *MockVisitorLister {
	mock := &MockVisitorLister{ctrl: ctrl}
	mock.recorder = &MockVisitorListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVisitorLister) EXPECT() *MockVisitorListerMockRecorder {
	return m.recorder
}

// ListVisitors mocks base method.
func (m *MockVisitorLister) ListVisitors(arg0 party.VisitorGroup) ([]party.Visitor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVisitors", arg0)
	ret0, _ := ret[0].([]party.Visitor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVisitors indicates an expected call of ListVisitors.
func (mr *MockVisitorListerMockRecorder) ListVisitors(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVisitors", reflect.TypeOf((*MockVisitorLister)(nil).ListVisitors), arg0)
}

// MockGreeter is a mock of Greeter interface.
type MockGreeter struct {
	ctrl     *gomock.Controller
	recorder *MockGreeterMockRecorder
}

// MockGreeterMockRecorder is the mock recorder for MockGreeter.
type MockGreeterMockRecorder struct {
	mock *MockGreeter
}

// NewMockGreeter creates a new mock instance.
func NewMockGreeter(ctrl *gomock.Controller) *MockGreeter {
	mock := &MockGreeter{ctrl: ctrl}
	mock.recorder = &MockGreeterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGreeter) EXPECT() *MockGreeterMockRecorder {
	return m.recorder
}

// Hello mocks base method.
func (m *MockGreeter) Hello(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hello", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Hello indicates an expected call of Hello.
func (mr *MockGreeterMockRecorder) Hello(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hello", reflect.TypeOf((*MockGreeter)(nil).Hello), arg0)
}
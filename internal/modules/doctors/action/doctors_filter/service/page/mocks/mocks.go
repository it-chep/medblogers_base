// Code generated by MockGen. DO NOT EDIT.
// Source: medblogers_base/internal/modules/doctors/action/doctors_filter/service/page (interfaces: Storage)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	dto "medblogers_base/internal/modules/doctors/action/doctors_filter/dto"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// GetDoctorsCountByFilter mocks base method.
func (m *MockStorage) GetDoctorsCountByFilter(arg0 context.Context, arg1 dto.Filter) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDoctorsCountByFilter", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDoctorsCountByFilter indicates an expected call of GetDoctorsCountByFilter.
func (mr *MockStorageMockRecorder) GetDoctorsCountByFilter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDoctorsCountByFilter", reflect.TypeOf((*MockStorage)(nil).GetDoctorsCountByFilter), arg0, arg1)
}

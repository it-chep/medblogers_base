// Code generated by MockGen. DO NOT EDIT.
// Source: medblogers_base/internal/modules/doctors/action/create_doctor/service/validate (interfaces: CityStorage,SpecialityStorage)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	city "medblogers_base/internal/modules/doctors/domain/city"
	speciality "medblogers_base/internal/modules/doctors/domain/speciality"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCityStorage is a mock of CityStorage interface.
type MockCityStorage struct {
	ctrl     *gomock.Controller
	recorder *MockCityStorageMockRecorder
}

// MockCityStorageMockRecorder is the mock recorder for MockCityStorage.
type MockCityStorageMockRecorder struct {
	mock *MockCityStorage
}

// NewMockCityStorage creates a new mock instance.
func NewMockCityStorage(ctrl *gomock.Controller) *MockCityStorage {
	mock := &MockCityStorage{ctrl: ctrl}
	mock.recorder = &MockCityStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCityStorage) EXPECT() *MockCityStorageMockRecorder {
	return m.recorder
}

// GetAllCities mocks base method.
func (m *MockCityStorage) GetAllCities(arg0 context.Context) ([]*city.City, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCities", arg0)
	ret0, _ := ret[0].([]*city.City)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCities indicates an expected call of GetAllCities.
func (mr *MockCityStorageMockRecorder) GetAllCities(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCities", reflect.TypeOf((*MockCityStorage)(nil).GetAllCities), arg0)
}

// MockSpecialityStorage is a mock of SpecialityStorage interface.
type MockSpecialityStorage struct {
	ctrl     *gomock.Controller
	recorder *MockSpecialityStorageMockRecorder
}

// MockSpecialityStorageMockRecorder is the mock recorder for MockSpecialityStorage.
type MockSpecialityStorageMockRecorder struct {
	mock *MockSpecialityStorage
}

// NewMockSpecialityStorage creates a new mock instance.
func NewMockSpecialityStorage(ctrl *gomock.Controller) *MockSpecialityStorage {
	mock := &MockSpecialityStorage{ctrl: ctrl}
	mock.recorder = &MockSpecialityStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpecialityStorage) EXPECT() *MockSpecialityStorageMockRecorder {
	return m.recorder
}

// GetAllSpecialities mocks base method.
func (m *MockSpecialityStorage) GetAllSpecialities(arg0 context.Context) ([]*speciality.Speciality, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSpecialities", arg0)
	ret0, _ := ret[0].([]*speciality.Speciality)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSpecialities indicates an expected call of GetAllSpecialities.
func (mr *MockSpecialityStorageMockRecorder) GetAllSpecialities(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSpecialities", reflect.TypeOf((*MockSpecialityStorage)(nil).GetAllSpecialities), arg0)
}

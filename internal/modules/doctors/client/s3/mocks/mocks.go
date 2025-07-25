// Code generated by MockGen. DO NOT EDIT.
// Source: medblogers_base/internal/modules/doctors/client/s3 (interfaces: S3Client,S3PresignClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	gomock "github.com/golang/mock/gomock"
)

// MockS3Client is a mock of S3Client interface.
type MockS3Client struct {
	ctrl     *gomock.Controller
	recorder *MockS3ClientMockRecorder
}

// MockS3ClientMockRecorder is the mock recorder for MockS3Client.
type MockS3ClientMockRecorder struct {
	mock *MockS3Client
}

// NewMockS3Client creates a new mock instance.
func NewMockS3Client(ctrl *gomock.Controller) *MockS3Client {
	mock := &MockS3Client{ctrl: ctrl}
	mock.recorder = &MockS3ClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3Client) EXPECT() *MockS3ClientMockRecorder {
	return m.recorder
}

// ListObjectsV2 mocks base method.
func (m *MockS3Client) ListObjectsV2(arg0 context.Context, arg1 *s3.ListObjectsV2Input, arg2 ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListObjectsV2", varargs...)
	ret0, _ := ret[0].(*s3.ListObjectsV2Output)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListObjectsV2 indicates an expected call of ListObjectsV2.
func (mr *MockS3ClientMockRecorder) ListObjectsV2(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjectsV2", reflect.TypeOf((*MockS3Client)(nil).ListObjectsV2), varargs...)
}

// PutObject mocks base method.
func (m *MockS3Client) PutObject(arg0 context.Context, arg1 *s3.PutObjectInput, arg2 ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutObject", varargs...)
	ret0, _ := ret[0].(*s3.PutObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutObject indicates an expected call of PutObject.
func (mr *MockS3ClientMockRecorder) PutObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObject", reflect.TypeOf((*MockS3Client)(nil).PutObject), varargs...)
}

// MockS3PresignClient is a mock of S3PresignClient interface.
type MockS3PresignClient struct {
	ctrl     *gomock.Controller
	recorder *MockS3PresignClientMockRecorder
}

// MockS3PresignClientMockRecorder is the mock recorder for MockS3PresignClient.
type MockS3PresignClientMockRecorder struct {
	mock *MockS3PresignClient
}

// NewMockS3PresignClient creates a new mock instance.
func NewMockS3PresignClient(ctrl *gomock.Controller) *MockS3PresignClient {
	mock := &MockS3PresignClient{ctrl: ctrl}
	mock.recorder = &MockS3PresignClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3PresignClient) EXPECT() *MockS3PresignClientMockRecorder {
	return m.recorder
}

// PresignGetObject mocks base method.
func (m *MockS3PresignClient) PresignGetObject(arg0 context.Context, arg1 *s3.GetObjectInput, arg2 ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PresignGetObject", varargs...)
	ret0, _ := ret[0].(*v4.PresignedHTTPRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PresignGetObject indicates an expected call of PresignGetObject.
func (mr *MockS3PresignClientMockRecorder) PresignGetObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PresignGetObject", reflect.TypeOf((*MockS3PresignClient)(nil).PresignGetObject), varargs...)
}

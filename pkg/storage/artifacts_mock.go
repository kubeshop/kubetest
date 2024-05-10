// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubeshop/testkube/pkg/storage (interfaces: ArtifactsStorage)

// Package storage is a generated GoMock package.
package storage

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	testkube "github.com/kubeshop/testkube/pkg/api/v1/testkube"
)

// MockArtifactsStorage is a mock of ArtifactsStorage interface.
type MockArtifactsStorage struct {
	ctrl     *gomock.Controller
	recorder *MockArtifactsStorageMockRecorder
}

// MockArtifactsStorageMockRecorder is the mock recorder for MockArtifactsStorage.
type MockArtifactsStorageMockRecorder struct {
	mock *MockArtifactsStorage
}

// NewMockArtifactsStorage creates a new mock instance.
func NewMockArtifactsStorage(ctrl *gomock.Controller) *MockArtifactsStorage {
	mock := &MockArtifactsStorage{ctrl: ctrl}
	mock.recorder = &MockArtifactsStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArtifactsStorage) EXPECT() *MockArtifactsStorageMockRecorder {
	return m.recorder
}

// DownloadArchive mocks base method.
func (m *MockArtifactsStorage) DownloadArchive(arg0 context.Context, arg1 string, arg2 []string) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadArchive", arg0, arg1, arg2)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadArchive indicates an expected call of DownloadArchive.
func (mr *MockArtifactsStorageMockRecorder) DownloadArchive(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadArchive", reflect.TypeOf((*MockArtifactsStorage)(nil).DownloadArchive), arg0, arg1, arg2)
}

// DownloadFile mocks base method.
func (m *MockArtifactsStorage) DownloadFile(arg0 context.Context, arg1, arg2, arg3, arg4, arg5 string) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadFile", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadFile indicates an expected call of DownloadFile.
func (mr *MockArtifactsStorageMockRecorder) DownloadFile(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadFile", reflect.TypeOf((*MockArtifactsStorage)(nil).DownloadFile), arg0, arg1, arg2, arg3, arg4, arg5)
}

// GetValidBucketName mocks base method.
func (m *MockArtifactsStorage) GetValidBucketName(arg0, arg1 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidBucketName", arg0, arg1)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetValidBucketName indicates an expected call of GetValidBucketName.
func (mr *MockArtifactsStorageMockRecorder) GetValidBucketName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidBucketName", reflect.TypeOf((*MockArtifactsStorage)(nil).GetValidBucketName), arg0, arg1)
}

// ListFiles mocks base method.
func (m *MockArtifactsStorage) ListFiles(arg0 context.Context, arg1, arg2, arg3, arg4 string) ([]testkube.Artifact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFiles", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]testkube.Artifact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFiles indicates an expected call of ListFiles.
func (mr *MockArtifactsStorageMockRecorder) ListFiles(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFiles", reflect.TypeOf((*MockArtifactsStorage)(nil).ListFiles), arg0, arg1, arg2, arg3, arg4)
}

// PlaceFiles mocks base method.
func (m *MockArtifactsStorage) PlaceFiles(arg0 context.Context, arg1 []string, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PlaceFiles", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PlaceFiles indicates an expected call of PlaceFiles.
func (mr *MockArtifactsStorageMockRecorder) PlaceFiles(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlaceFiles", reflect.TypeOf((*MockArtifactsStorage)(nil).PlaceFiles), arg0, arg1, arg2)
}

// UploadFile mocks base method.
func (m *MockArtifactsStorage) UploadFile(arg0 context.Context, arg1, arg2 string, arg3 io.Reader, arg4 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadFile indicates an expected call of UploadFile.
func (mr *MockArtifactsStorageMockRecorder) UploadFile(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockArtifactsStorage)(nil).UploadFile), arg0, arg1, arg2, arg3, arg4)
}
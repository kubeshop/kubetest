// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubeshop/testkube/pkg/storage (interfaces: Client)

// Package storage is a generated GoMock package.
package storage

import (
	context "context"
	io "io"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	testkube "github.com/kubeshop/testkube/pkg/api/v1/testkube"
	minio "github.com/minio/minio-go/v7"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// BucketExists mocks base method.
func (m *MockClient) BucketExists(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BucketExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BucketExists indicates an expected call of BucketExists.
func (mr *MockClientMockRecorder) BucketExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BucketExists", reflect.TypeOf((*MockClient)(nil).BucketExists), arg0, arg1)
}

// CreateBucket mocks base method.
func (m *MockClient) CreateBucket(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBucket", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBucket indicates an expected call of CreateBucket.
func (mr *MockClientMockRecorder) CreateBucket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucket", reflect.TypeOf((*MockClient)(nil).CreateBucket), arg0, arg1)
}

// DeleteBucket mocks base method.
func (m *MockClient) DeleteBucket(arg0 context.Context, arg1 string, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBucket", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBucket indicates an expected call of DeleteBucket.
func (mr *MockClientMockRecorder) DeleteBucket(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBucket", reflect.TypeOf((*MockClient)(nil).DeleteBucket), arg0, arg1, arg2)
}

// DeleteFile mocks base method.
func (m *MockClient) DeleteFile(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFile", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFile indicates an expected call of DeleteFile.
func (mr *MockClientMockRecorder) DeleteFile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFile", reflect.TypeOf((*MockClient)(nil).DeleteFile), arg0, arg1, arg2)
}

// DeleteFileFromBucket mocks base method.
func (m *MockClient) DeleteFileFromBucket(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFileFromBucket", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFileFromBucket indicates an expected call of DeleteFileFromBucket.
func (mr *MockClientMockRecorder) DeleteFileFromBucket(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFileFromBucket", reflect.TypeOf((*MockClient)(nil).DeleteFileFromBucket), arg0, arg1, arg2, arg3)
}

// DownloadArchive mocks base method.
func (m *MockClient) DownloadArchive(arg0 context.Context, arg1 string, arg2 []string) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadArchive", arg0, arg1, arg2)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadArchive indicates an expected call of DownloadArchive.
func (mr *MockClientMockRecorder) DownloadArchive(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadArchive", reflect.TypeOf((*MockClient)(nil).DownloadArchive), arg0, arg1, arg2)
}

// DownloadArchiveFromBucket mocks base method.
func (m *MockClient) DownloadArchiveFromBucket(arg0 context.Context, arg1, arg2 string, arg3 []string) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadArchiveFromBucket", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadArchiveFromBucket indicates an expected call of DownloadArchiveFromBucket.
func (mr *MockClientMockRecorder) DownloadArchiveFromBucket(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadArchiveFromBucket", reflect.TypeOf((*MockClient)(nil).DownloadArchiveFromBucket), arg0, arg1, arg2, arg3)
}

// DownloadFile mocks base method.
func (m *MockClient) DownloadFile(arg0 context.Context, arg1, arg2 string) (*minio.Object, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadFile", arg0, arg1, arg2)
	ret0, _ := ret[0].(*minio.Object)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadFile indicates an expected call of DownloadFile.
func (mr *MockClientMockRecorder) DownloadFile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadFile", reflect.TypeOf((*MockClient)(nil).DownloadFile), arg0, arg1, arg2)
}

// DownloadFileFromBucket mocks base method.
func (m *MockClient) DownloadFileFromBucket(arg0 context.Context, arg1, arg2, arg3 string) (io.Reader, minio.ObjectInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadFileFromBucket", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(minio.ObjectInfo)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// DownloadFileFromBucket indicates an expected call of DownloadFileFromBucket.
func (mr *MockClientMockRecorder) DownloadFileFromBucket(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadFileFromBucket", reflect.TypeOf((*MockClient)(nil).DownloadFileFromBucket), arg0, arg1, arg2, arg3)
}

// GetValidBucketName mocks base method.
func (m *MockClient) GetValidBucketName(arg0, arg1 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidBucketName", arg0, arg1)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetValidBucketName indicates an expected call of GetValidBucketName.
func (mr *MockClientMockRecorder) GetValidBucketName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidBucketName", reflect.TypeOf((*MockClient)(nil).GetValidBucketName), arg0, arg1)
}

// IsConnectionPossible mocks base method.
func (m *MockClient) IsConnectionPossible(arg0 context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsConnectionPossible", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsConnectionPossible indicates an expected call of IsConnectionPossible.
func (mr *MockClientMockRecorder) IsConnectionPossible(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsConnectionPossible", reflect.TypeOf((*MockClient)(nil).IsConnectionPossible), arg0)
}

// ListBuckets mocks base method.
func (m *MockClient) ListBuckets(arg0 context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBuckets", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBuckets indicates an expected call of ListBuckets.
func (mr *MockClientMockRecorder) ListBuckets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBuckets", reflect.TypeOf((*MockClient)(nil).ListBuckets), arg0)
}

// ListFiles mocks base method.
func (m *MockClient) ListFiles(arg0 context.Context, arg1 string) ([]testkube.Artifact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFiles", arg0, arg1)
	ret0, _ := ret[0].([]testkube.Artifact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFiles indicates an expected call of ListFiles.
func (mr *MockClientMockRecorder) ListFiles(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFiles", reflect.TypeOf((*MockClient)(nil).ListFiles), arg0, arg1)
}

// PlaceFiles mocks base method.
func (m *MockClient) PlaceFiles(arg0 context.Context, arg1 []string, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PlaceFiles", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PlaceFiles indicates an expected call of PlaceFiles.
func (mr *MockClientMockRecorder) PlaceFiles(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PlaceFiles", reflect.TypeOf((*MockClient)(nil).PlaceFiles), arg0, arg1, arg2)
}

// PresignDownloadFileFromBucket mocks base method.
func (m *MockClient) PresignDownloadFileFromBucket(arg0 context.Context, arg1, arg2, arg3 string, arg4 time.Duration) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PresignDownloadFileFromBucket", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PresignDownloadFileFromBucket indicates an expected call of PresignDownloadFileFromBucket.
func (mr *MockClientMockRecorder) PresignDownloadFileFromBucket(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PresignDownloadFileFromBucket", reflect.TypeOf((*MockClient)(nil).PresignDownloadFileFromBucket), arg0, arg1, arg2, arg3, arg4)
}

// PresignUploadFileToBucket mocks base method.
func (m *MockClient) PresignUploadFileToBucket(arg0 context.Context, arg1, arg2, arg3 string, arg4 time.Duration) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PresignUploadFileToBucket", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PresignUploadFileToBucket indicates an expected call of PresignUploadFileToBucket.
func (mr *MockClientMockRecorder) PresignUploadFileToBucket(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PresignUploadFileToBucket", reflect.TypeOf((*MockClient)(nil).PresignUploadFileToBucket), arg0, arg1, arg2, arg3, arg4)
}

// SaveFile mocks base method.
func (m *MockClient) SaveFile(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveFile", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveFile indicates an expected call of SaveFile.
func (mr *MockClientMockRecorder) SaveFile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveFile", reflect.TypeOf((*MockClient)(nil).SaveFile), arg0, arg1, arg2)
}

// UploadFile mocks base method.
func (m *MockClient) UploadFile(arg0 context.Context, arg1, arg2 string, arg3 io.Reader, arg4 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadFile indicates an expected call of UploadFile.
func (mr *MockClientMockRecorder) UploadFile(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockClient)(nil).UploadFile), arg0, arg1, arg2, arg3, arg4)
}

// UploadFileToBucket mocks base method.
func (m *MockClient) UploadFileToBucket(arg0 context.Context, arg1, arg2, arg3 string, arg4 io.Reader, arg5 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFileToBucket", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadFileToBucket indicates an expected call of UploadFileToBucket.
func (mr *MockClientMockRecorder) UploadFileToBucket(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFileToBucket", reflect.TypeOf((*MockClient)(nil).UploadFileToBucket), arg0, arg1, arg2, arg3, arg4, arg5)
}

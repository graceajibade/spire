// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spiffe/spire/proto/api/node (interfaces: NodeClient,NodeServer)

// Package mock_node is a generated GoMock package.
package mock_node

import (
	gomock "github.com/golang/mock/gomock"
	node "github.com/spiffe/spire/proto/api/node"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockNodeClient is a mock of NodeClient interface
type MockNodeClient struct {
	ctrl     *gomock.Controller
	recorder *MockNodeClientMockRecorder
}

// MockNodeClientMockRecorder is the mock recorder for MockNodeClient
type MockNodeClientMockRecorder struct {
	mock *MockNodeClient
}

// NewMockNodeClient creates a new mock instance
func NewMockNodeClient(ctrl *gomock.Controller) *MockNodeClient {
	mock := &MockNodeClient{ctrl: ctrl}
	mock.recorder = &MockNodeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeClient) EXPECT() *MockNodeClientMockRecorder {
	return m.recorder
}

// FetchBaseSVID mocks base method
func (m *MockNodeClient) FetchBaseSVID(arg0 context.Context, arg1 *node.FetchBaseSVIDRequest, arg2 ...grpc.CallOption) (*node.FetchBaseSVIDResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FetchBaseSVID", varargs...)
	ret0, _ := ret[0].(*node.FetchBaseSVIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBaseSVID indicates an expected call of FetchBaseSVID
func (mr *MockNodeClientMockRecorder) FetchBaseSVID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBaseSVID", reflect.TypeOf((*MockNodeClient)(nil).FetchBaseSVID), varargs...)
}

// FetchCPBundle mocks base method
func (m *MockNodeClient) FetchCPBundle(arg0 context.Context, arg1 *node.FetchCPBundleRequest, arg2 ...grpc.CallOption) (*node.FetchCPBundleResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FetchCPBundle", varargs...)
	ret0, _ := ret[0].(*node.FetchCPBundleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCPBundle indicates an expected call of FetchCPBundle
func (mr *MockNodeClientMockRecorder) FetchCPBundle(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCPBundle", reflect.TypeOf((*MockNodeClient)(nil).FetchCPBundle), varargs...)
}

// FetchFederatedBundle mocks base method
func (m *MockNodeClient) FetchFederatedBundle(arg0 context.Context, arg1 *node.FetchFederatedBundleRequest, arg2 ...grpc.CallOption) (*node.FetchFederatedBundleResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FetchFederatedBundle", varargs...)
	ret0, _ := ret[0].(*node.FetchFederatedBundleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchFederatedBundle indicates an expected call of FetchFederatedBundle
func (mr *MockNodeClientMockRecorder) FetchFederatedBundle(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchFederatedBundle", reflect.TypeOf((*MockNodeClient)(nil).FetchFederatedBundle), varargs...)
}

// FetchSVID mocks base method
func (m *MockNodeClient) FetchSVID(arg0 context.Context, arg1 *node.FetchSVIDRequest, arg2 ...grpc.CallOption) (*node.FetchSVIDResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FetchSVID", varargs...)
	ret0, _ := ret[0].(*node.FetchSVIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchSVID indicates an expected call of FetchSVID
func (mr *MockNodeClientMockRecorder) FetchSVID(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchSVID", reflect.TypeOf((*MockNodeClient)(nil).FetchSVID), varargs...)
}

// MockNodeServer is a mock of NodeServer interface
type MockNodeServer struct {
	ctrl     *gomock.Controller
	recorder *MockNodeServerMockRecorder
}

// MockNodeServerMockRecorder is the mock recorder for MockNodeServer
type MockNodeServerMockRecorder struct {
	mock *MockNodeServer
}

// NewMockNodeServer creates a new mock instance
func NewMockNodeServer(ctrl *gomock.Controller) *MockNodeServer {
	mock := &MockNodeServer{ctrl: ctrl}
	mock.recorder = &MockNodeServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeServer) EXPECT() *MockNodeServerMockRecorder {
	return m.recorder
}

// FetchBaseSVID mocks base method
func (m *MockNodeServer) FetchBaseSVID(arg0 context.Context, arg1 *node.FetchBaseSVIDRequest) (*node.FetchBaseSVIDResponse, error) {
	ret := m.ctrl.Call(m, "FetchBaseSVID", arg0, arg1)
	ret0, _ := ret[0].(*node.FetchBaseSVIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBaseSVID indicates an expected call of FetchBaseSVID
func (mr *MockNodeServerMockRecorder) FetchBaseSVID(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBaseSVID", reflect.TypeOf((*MockNodeServer)(nil).FetchBaseSVID), arg0, arg1)
}

// FetchCPBundle mocks base method
func (m *MockNodeServer) FetchCPBundle(arg0 context.Context, arg1 *node.FetchCPBundleRequest) (*node.FetchCPBundleResponse, error) {
	ret := m.ctrl.Call(m, "FetchCPBundle", arg0, arg1)
	ret0, _ := ret[0].(*node.FetchCPBundleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCPBundle indicates an expected call of FetchCPBundle
func (mr *MockNodeServerMockRecorder) FetchCPBundle(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCPBundle", reflect.TypeOf((*MockNodeServer)(nil).FetchCPBundle), arg0, arg1)
}

// FetchFederatedBundle mocks base method
func (m *MockNodeServer) FetchFederatedBundle(arg0 context.Context, arg1 *node.FetchFederatedBundleRequest) (*node.FetchFederatedBundleResponse, error) {
	ret := m.ctrl.Call(m, "FetchFederatedBundle", arg0, arg1)
	ret0, _ := ret[0].(*node.FetchFederatedBundleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchFederatedBundle indicates an expected call of FetchFederatedBundle
func (mr *MockNodeServerMockRecorder) FetchFederatedBundle(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchFederatedBundle", reflect.TypeOf((*MockNodeServer)(nil).FetchFederatedBundle), arg0, arg1)
}

// FetchSVID mocks base method
func (m *MockNodeServer) FetchSVID(arg0 context.Context, arg1 *node.FetchSVIDRequest) (*node.FetchSVIDResponse, error) {
	ret := m.ctrl.Call(m, "FetchSVID", arg0, arg1)
	ret0, _ := ret[0].(*node.FetchSVIDResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchSVID indicates an expected call of FetchSVID
func (mr *MockNodeServerMockRecorder) FetchSVID(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchSVID", reflect.TypeOf((*MockNodeServer)(nil).FetchSVID), arg0, arg1)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: kitex_gen/douyin/base/baseservice/client.go

// Package basecli is a generated GoMock package.
package basecli

import (
	base "ByteTech-7355608/douyin-server/kitex_gen/douyin/base"
	context "context"
	reflect "reflect"

	callopt "github.com/cloudwego/kitex/client/callopt"
	gomock "github.com/golang/mock/gomock"
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

// Feed mocks base method.
func (m *MockClient) Feed(ctx context.Context, req *base.DouyinFeedRequest, callOptions ...callopt.Option) (*base.DouyinFeedResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range callOptions {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Feed", varargs...)
	ret0, _ := ret[0].(*base.DouyinFeedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Feed indicates an expected call of Feed.
func (mr *MockClientMockRecorder) Feed(ctx, req interface{}, callOptions ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, callOptions...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Feed", reflect.TypeOf((*MockClient)(nil).Feed), varargs...)
}

// PublishAction mocks base method.
func (m *MockClient) PublishAction(ctx context.Context, req *base.DouyinPublishActionRequest, callOptions ...callopt.Option) (*base.DouyinPublishActionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range callOptions {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PublishAction", varargs...)
	ret0, _ := ret[0].(*base.DouyinPublishActionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublishAction indicates an expected call of PublishAction.
func (mr *MockClientMockRecorder) PublishAction(ctx, req interface{}, callOptions ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, callOptions...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishAction", reflect.TypeOf((*MockClient)(nil).PublishAction), varargs...)
}

// PublishList mocks base method.
func (m *MockClient) PublishList(ctx context.Context, req *base.DouyinPublishListRequest, callOptions ...callopt.Option) (*base.DouyinPublishListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range callOptions {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PublishList", varargs...)
	ret0, _ := ret[0].(*base.DouyinPublishListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PublishList indicates an expected call of PublishList.
func (mr *MockClientMockRecorder) PublishList(ctx, req interface{}, callOptions ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, callOptions...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishList", reflect.TypeOf((*MockClient)(nil).PublishList), varargs...)
}

// UserLogin mocks base method.
func (m *MockClient) UserLogin(ctx context.Context, req *base.DouyinUserLoginRequest, callOptions ...callopt.Option) (*base.DouyinUserLoginResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range callOptions {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UserLogin", varargs...)
	ret0, _ := ret[0].(*base.DouyinUserLoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserLogin indicates an expected call of UserLogin.
func (mr *MockClientMockRecorder) UserLogin(ctx, req interface{}, callOptions ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, callOptions...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserLogin", reflect.TypeOf((*MockClient)(nil).UserLogin), varargs...)
}

// UserMsg mocks base method.
func (m *MockClient) UserMsg(ctx context.Context, req *base.DouyinUserRequest, callOptions ...callopt.Option) (*base.DouyinUserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range callOptions {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UserMsg", varargs...)
	ret0, _ := ret[0].(*base.DouyinUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserMsg indicates an expected call of UserMsg.
func (mr *MockClientMockRecorder) UserMsg(ctx, req interface{}, callOptions ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, callOptions...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserMsg", reflect.TypeOf((*MockClient)(nil).UserMsg), varargs...)
}

// UserRegister mocks base method.
func (m *MockClient) UserRegister(ctx context.Context, req *base.DouyinUserRegisterRequest, callOptions ...callopt.Option) (*base.DouyinUserRegisterResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, req}
	for _, a := range callOptions {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UserRegister", varargs...)
	ret0, _ := ret[0].(*base.DouyinUserRegisterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserRegister indicates an expected call of UserRegister.
func (mr *MockClientMockRecorder) UserRegister(ctx, req interface{}, callOptions ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, req}, callOptions...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserRegister", reflect.TypeOf((*MockClient)(nil).UserRegister), varargs...)
}

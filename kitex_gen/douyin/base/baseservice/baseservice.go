// Code generated by Kitex v0.4.4. DO NOT EDIT.

package baseservice

import (
	base "ByteTech-7355608/douyin-server/kitex_gen/douyin/base"
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return baseServiceServiceInfo
}

var baseServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "BaseService"
	handlerType := (*base.BaseService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Feed":          kitex.NewMethodInfo(feedHandler, newBaseServiceFeedArgs, newBaseServiceFeedResult, false),
		"UserRegister":  kitex.NewMethodInfo(userRegisterHandler, newBaseServiceUserRegisterArgs, newBaseServiceUserRegisterResult, false),
		"UserLogin":     kitex.NewMethodInfo(userLoginHandler, newBaseServiceUserLoginArgs, newBaseServiceUserLoginResult, false),
		"UserMsg":       kitex.NewMethodInfo(userMsgHandler, newBaseServiceUserMsgArgs, newBaseServiceUserMsgResult, false),
		"PublishAction": kitex.NewMethodInfo(publishActionHandler, newBaseServicePublishActionArgs, newBaseServicePublishActionResult, false),
		"PublishList":   kitex.NewMethodInfo(publishListHandler, newBaseServicePublishListArgs, newBaseServicePublishListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "base",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func feedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*base.BaseServiceFeedArgs)
	realResult := result.(*base.BaseServiceFeedResult)
	success, err := handler.(base.BaseService).Feed(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBaseServiceFeedArgs() interface{} {
	return base.NewBaseServiceFeedArgs()
}

func newBaseServiceFeedResult() interface{} {
	return base.NewBaseServiceFeedResult()
}

func userRegisterHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*base.BaseServiceUserRegisterArgs)
	realResult := result.(*base.BaseServiceUserRegisterResult)
	success, err := handler.(base.BaseService).UserRegister(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBaseServiceUserRegisterArgs() interface{} {
	return base.NewBaseServiceUserRegisterArgs()
}

func newBaseServiceUserRegisterResult() interface{} {
	return base.NewBaseServiceUserRegisterResult()
}

func userLoginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*base.BaseServiceUserLoginArgs)
	realResult := result.(*base.BaseServiceUserLoginResult)
	success, err := handler.(base.BaseService).UserLogin(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBaseServiceUserLoginArgs() interface{} {
	return base.NewBaseServiceUserLoginArgs()
}

func newBaseServiceUserLoginResult() interface{} {
	return base.NewBaseServiceUserLoginResult()
}

func userMsgHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*base.BaseServiceUserMsgArgs)
	realResult := result.(*base.BaseServiceUserMsgResult)
	success, err := handler.(base.BaseService).UserMsg(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBaseServiceUserMsgArgs() interface{} {
	return base.NewBaseServiceUserMsgArgs()
}

func newBaseServiceUserMsgResult() interface{} {
	return base.NewBaseServiceUserMsgResult()
}

func publishActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*base.BaseServicePublishActionArgs)
	realResult := result.(*base.BaseServicePublishActionResult)
	success, err := handler.(base.BaseService).PublishAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBaseServicePublishActionArgs() interface{} {
	return base.NewBaseServicePublishActionArgs()
}

func newBaseServicePublishActionResult() interface{} {
	return base.NewBaseServicePublishActionResult()
}

func publishListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*base.BaseServicePublishListArgs)
	realResult := result.(*base.BaseServicePublishListResult)
	success, err := handler.(base.BaseService).PublishList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBaseServicePublishListArgs() interface{} {
	return base.NewBaseServicePublishListArgs()
}

func newBaseServicePublishListResult() interface{} {
	return base.NewBaseServicePublishListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Feed(ctx context.Context, req *base.DouyinFeedRequest) (r *base.DouyinFeedResponse, err error) {
	var _args base.BaseServiceFeedArgs
	_args.Req = req
	var _result base.BaseServiceFeedResult
	if err = p.c.Call(ctx, "Feed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserRegister(ctx context.Context, req *base.DouyinUserRegisterRequest) (r *base.DouyinUserRegisterResponse, err error) {
	var _args base.BaseServiceUserRegisterArgs
	_args.Req = req
	var _result base.BaseServiceUserRegisterResult
	if err = p.c.Call(ctx, "UserRegister", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserLogin(ctx context.Context, req *base.DouyinUserLoginRequest) (r *base.DouyinUserLoginResponse, err error) {
	var _args base.BaseServiceUserLoginArgs
	_args.Req = req
	var _result base.BaseServiceUserLoginResult
	if err = p.c.Call(ctx, "UserLogin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserMsg(ctx context.Context, req *base.DouyinUserRequest) (r *base.DouyinUserResponse, err error) {
	var _args base.BaseServiceUserMsgArgs
	_args.Req = req
	var _result base.BaseServiceUserMsgResult
	if err = p.c.Call(ctx, "UserMsg", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishAction(ctx context.Context, req *base.DouyinPublishActionRequest) (r *base.DouyinPublishActionResponse, err error) {
	var _args base.BaseServicePublishActionArgs
	_args.Req = req
	var _result base.BaseServicePublishActionResult
	if err = p.c.Call(ctx, "PublishAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishList(ctx context.Context, req *base.DouyinPublishListRequest) (r *base.DouyinPublishListResponse, err error) {
	var _args base.BaseServicePublishListArgs
	_args.Req = req
	var _result base.BaseServicePublishListResult
	if err = p.c.Call(ctx, "PublishList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

// Code generated by Kitex v0.4.4. DO NOT EDIT.

package socialservice

import (
	social "ByteTech-7355608/douyin-server/kitex_gen/douyin/social"
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return socialServiceServiceInfo
}

var socialServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "SocialService"
	handlerType := (*social.SocialService)(nil)
	methods := map[string]kitex.MethodInfo{
		"FollowAction": kitex.NewMethodInfo(followActionHandler, newSocialServiceFollowActionArgs, newSocialServiceFollowActionResult, false),
		"FollowList":   kitex.NewMethodInfo(followListHandler, newSocialServiceFollowListArgs, newSocialServiceFollowListResult, false),
		"FollowerList": kitex.NewMethodInfo(followerListHandler, newSocialServiceFollowerListArgs, newSocialServiceFollowerListResult, false),
		"FriendList":   kitex.NewMethodInfo(friendListHandler, newSocialServiceFriendListArgs, newSocialServiceFriendListResult, false),
		"MessageList":  kitex.NewMethodInfo(messageListHandler, newSocialServiceMessageListArgs, newSocialServiceMessageListResult, false),
		"SendMessage":  kitex.NewMethodInfo(sendMessageHandler, newSocialServiceSendMessageArgs, newSocialServiceSendMessageResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "social",
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

func followActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*social.SocialServiceFollowActionArgs)
	realResult := result.(*social.SocialServiceFollowActionResult)
	success, err := handler.(social.SocialService).FollowAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSocialServiceFollowActionArgs() interface{} {
	return social.NewSocialServiceFollowActionArgs()
}

func newSocialServiceFollowActionResult() interface{} {
	return social.NewSocialServiceFollowActionResult()
}

func followListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*social.SocialServiceFollowListArgs)
	realResult := result.(*social.SocialServiceFollowListResult)
	success, err := handler.(social.SocialService).FollowList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSocialServiceFollowListArgs() interface{} {
	return social.NewSocialServiceFollowListArgs()
}

func newSocialServiceFollowListResult() interface{} {
	return social.NewSocialServiceFollowListResult()
}

func followerListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*social.SocialServiceFollowerListArgs)
	realResult := result.(*social.SocialServiceFollowerListResult)
	success, err := handler.(social.SocialService).FollowerList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSocialServiceFollowerListArgs() interface{} {
	return social.NewSocialServiceFollowerListArgs()
}

func newSocialServiceFollowerListResult() interface{} {
	return social.NewSocialServiceFollowerListResult()
}

func friendListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*social.SocialServiceFriendListArgs)
	realResult := result.(*social.SocialServiceFriendListResult)
	success, err := handler.(social.SocialService).FriendList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSocialServiceFriendListArgs() interface{} {
	return social.NewSocialServiceFriendListArgs()
}

func newSocialServiceFriendListResult() interface{} {
	return social.NewSocialServiceFriendListResult()
}

func messageListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*social.SocialServiceMessageListArgs)
	realResult := result.(*social.SocialServiceMessageListResult)
	success, err := handler.(social.SocialService).MessageList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSocialServiceMessageListArgs() interface{} {
	return social.NewSocialServiceMessageListArgs()
}

func newSocialServiceMessageListResult() interface{} {
	return social.NewSocialServiceMessageListResult()
}

func sendMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*social.SocialServiceSendMessageArgs)
	realResult := result.(*social.SocialServiceSendMessageResult)
	success, err := handler.(social.SocialService).SendMessage(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSocialServiceSendMessageArgs() interface{} {
	return social.NewSocialServiceSendMessageArgs()
}

func newSocialServiceSendMessageResult() interface{} {
	return social.NewSocialServiceSendMessageResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) FollowAction(ctx context.Context, req *social.DouyinFollowActionRequest) (r *social.DouyinFollowActionResponse, err error) {
	var _args social.SocialServiceFollowActionArgs
	_args.Req = req
	var _result social.SocialServiceFollowActionResult
	if err = p.c.Call(ctx, "FollowAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FollowList(ctx context.Context, req *social.DouyinFollowingListRequest) (r *social.DouyinFollowingListResponse, err error) {
	var _args social.SocialServiceFollowListArgs
	_args.Req = req
	var _result social.SocialServiceFollowListResult
	if err = p.c.Call(ctx, "FollowList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FollowerList(ctx context.Context, req *social.DouyinFollowerListRequest) (r *social.DouyinFollowerListResponse, err error) {
	var _args social.SocialServiceFollowerListArgs
	_args.Req = req
	var _result social.SocialServiceFollowerListResult
	if err = p.c.Call(ctx, "FollowerList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FriendList(ctx context.Context, req *social.DouyinRelationFriendListRequest) (r *social.DouyinRelationFriendListResponse, err error) {
	var _args social.SocialServiceFriendListArgs
	_args.Req = req
	var _result social.SocialServiceFriendListResult
	if err = p.c.Call(ctx, "FriendList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessageList(ctx context.Context, req *social.DouyinMessageChatRequest) (r *social.DouyinMessageChatResponse, err error) {
	var _args social.SocialServiceMessageListArgs
	_args.Req = req
	var _result social.SocialServiceMessageListResult
	if err = p.c.Call(ctx, "MessageList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SendMessage(ctx context.Context, req *social.DouyinMessageActionRequest) (r *social.DouyinMessageActionResponse, err error) {
	var _args social.SocialServiceSendMessageArgs
	_args.Req = req
	var _result social.SocialServiceSendMessageResult
	if err = p.c.Call(ctx, "SendMessage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

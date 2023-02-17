package handler

import (
	api "ByteTech-7355608/douyin-server/cmd/api/biz/model/douyin/social"
	rpc "ByteTech-7355608/douyin-server/kitex_gen/douyin/social"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

// FollowerList
// @router /douyin/relation/follower/list/ [GET]
func (h *Handler) FollowerList(ctx context.Context, c *app.RequestContext) {
	req := &api.DouyinFollowerListRequest{}
	rpcReq := &rpc.DouyinFollowerListRequest{}
	if h.Pre(ctx, c, req, rpcReq) {
		rpcReq.BaseReq = h.GetReqBase(c)
		rpcResp, err := h.RPC().Social().Client().FollowerList(ctx, rpcReq)
		if err != nil {
			return
		}
		resp := &api.DouyinFollowerListResponse{}
		h.After(ctx, c, resp, rpcResp, err)
	}
}

// FriendList
// @router /douyin/relation/friend/list/ [GET]
func (h *Handler) FriendList(ctx context.Context, c *app.RequestContext) {
	req := &api.DouyinRelationFriendListRequest{}
	rpcReq := &rpc.DouyinRelationFriendListRequest{}
	if h.Pre(ctx, c, req, rpcReq) {
		rpcReq.BaseReq = h.GetReqBase(c)
		rpcResp, err := h.RPC().Social().Client().FriendList(ctx, rpcReq)
		if err != nil {
			return
		}
		resp := &api.DouyinRelationFriendListResponse{}
		h.After(ctx, c, resp, rpcResp, err)
	}
}

// RelationAction .
// @router /douyin/relation/action [POST]
func (h *Handler) FollowAction(ctx context.Context, c *app.RequestContext) {
	req := &api.DouyinFollowActionRequest{}
	rpcReq := &rpc.DouyinFollowActionRequest{}
	if h.Pre(ctx, c, req, rpcReq) {
		rpcReq.BaseReq = h.GetReqBase(c)
		rpcResp, err := h.RPC().Social().Client().FollowAction(ctx, rpcReq)
		if err != nil {
			return
		}
		resp := &api.DouyinFollowActionResponse{}
		h.After(ctx, c, resp, rpcResp, err)
	}
}

// SendMessage .
// @router /douyin/message/action/ [POST]
func (h *Handler) SendMessage(ctx context.Context, c *app.RequestContext) {
	req := &api.DouyinMessageActionRequest{}
	rpcReq := &rpc.DouyinMessageActionRequest{}
	if h.Pre(ctx, c, req, rpcReq) {
		rpcReq.BaseReq = h.GetReqBase(c)
		rpcResp, err := h.RPC().Social().Client().SendMessage(ctx, rpcReq)
		if err != nil {
			return
		}
		resp := &api.DouyinMessageActionResponse{}
		h.After(ctx, c, &resp, rpcResp, err)
	}
}

// MessageList .
// @router /douyin/message/chat/ [GET]
func (h *Handler) MessageList(ctx context.Context, c *app.RequestContext) {
	req := &api.DouyinMessageChatRequest{}
	rpcReq := &rpc.DouyinMessageChatRequest{}
	if h.Pre(ctx, c, req, rpcReq) {
		rpcReq.BaseReq = h.GetReqBase(c)
		rpcResp, err := h.RPC().Social().Client().MessageList(ctx, rpcReq)
		if err != nil {
			return
		}
		resp := &api.DouyinMessageChatResponse{}
		h.After(ctx, c, &resp, rpcResp, err)
	}
}

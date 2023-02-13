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

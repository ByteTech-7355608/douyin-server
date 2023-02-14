package handler

import (
	api "ByteTech-7355608/douyin-server/cmd/api/biz/model/douyin/social"
	rpc "ByteTech-7355608/douyin-server/kitex_gen/douyin/social"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

// FollowAction .
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
		resp := api.DouyinFollowActionResponse{}
		h.After(ctx, c, &resp, rpcResp, err)
	}
}

// FollowList .
// @router /douyin/relation/list [GET]
func (h *Handler) FollowList(ctx context.Context, c *app.RequestContext) {
	req := &api.DouyinFollowingListRequest{}
	rpcReq := &rpc.DouyinFollowingListRequest{}
	if h.Pre(ctx, c, req, rpcReq) {
		//rpcReq.BaseReq = h.GetReqBase(c)
		rpcResp, err := h.RPC().Social().Client().FollowList(ctx, rpcReq)
		if err != nil {
			return
		}
		resp := api.DouyinFollowingListResponse{}
		h.After(ctx, c, &resp, rpcResp, err)
	}
}

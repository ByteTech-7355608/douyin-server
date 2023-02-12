package handler

import (
	api "ByteTech-7355608/douyin-server/cmd/api/biz/model/douyin/social"
	rpc "ByteTech-7355608/douyin-server/kitex_gen/douyin/social"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

// RelationAction .
// @router /douyin/relation/action [POST]
func (h *Handler) RelationAction(ctx context.Context, c *app.RequestContext) {
	req := api.DouyinFollowActionRequest{}
	rpcReq := rpc.DouyinFollowActionRequest{}
	if h.Pre(ctx, c, &req, &rpcReq) {
		userID, ok := c.Get("userid") // 当前用户id
		if !ok {
			return
		}
		rpcReq.FollowingId = userID.(int64) //关注者的id
		rpcReq.FollowerId = req.FollowerID  //被关注者的id
		rpcResp, err := h.RPC().Social().Client().FollowAction(ctx, &rpcReq)
		if err != nil {
			return
		}
		resp := rpc.DouyinFollowActionResponse{}
		h.After(ctx, c, &resp, &rpcResp, err)
	}
}

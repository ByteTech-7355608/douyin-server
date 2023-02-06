package handler

import (
	api "ByteTech-7355608/douyin-server/cmd/api/biz/model/douyin/base"
	rpc "ByteTech-7355608/douyin-server/kitex_gen/douyin/base"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
)

// UserRegister .
// @router /douyin/user/register [POST]
func (h *Handler) UserRegister(ctx context.Context, c *app.RequestContext) {
	req := &api.DouyinUserRegisterRequest{}
	rpcReq := &rpc.DouyinUserRegisterRequest{}
	if h.Pre(ctx, c, req, rpcReq) {
		rpcResp, err := h.RPC().Base().Client().UserRegister(ctx, rpcReq)
		if err != nil {
			Log.Error(err)
			return
		}
		resp := rpc.DouyinUserRegisterResponse{}
		h.After(ctx, c, &resp, rpcResp, err)
	}

}

// PublishAction .
// @router /douyin/publish/action [POST]
func (h *Handler) PublishAction(ctx context.Context, c *app.RequestContext) {
	req := &api.DouyinPublishActionRequest{}
	rpcReq := &rpc.DouyinPublishActionRequest{}
	if h.Pre(ctx, c, req, rpcReq) {
		rpcResp, err := h.RPC().Base().Client().PublishAction(ctx, rpcReq)
		if err != nil {
			fmt.Println(err)
			Log.Error(err)
			return
		}
		resp := rpc.DouyinPublishActionResponse{}
		h.After(ctx, c, &resp, rpcResp, err)
	}
}

package handler

import (
	api "ByteTech-7355608/douyin-server/cmd/api/biz/model/douyin/base"
	rpc "ByteTech-7355608/douyin-server/kitex_gen/douyin/base"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/model"
	"context"
	"strings"

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
			return
		}
		resp := &api.DouyinUserRegisterResponse{}
		h.After(ctx, c, resp, rpcResp, err)
	}
}

// UserLogin
// @router /douyin/user/Login [POST]
func (h *Handler) UserLogin(ctx context.Context, c *app.RequestContext) {
	req := &api.DouyinUserLoginRequest{}
	rpcReq := &rpc.DouyinUserLoginRequest{}
	if h.Pre(ctx, c, req, rpcReq) {
		rpcResp, err := h.RPC().Base().Client().UserLogin(ctx, rpcReq)
		if err != nil {
			return
		}
		resp := rpc.DouyinUserLoginResponse{}
		h.After(ctx, c, &resp, rpcResp, err)
	}

}

// PublishAction
// @router /douyin/publish/action [POST]
func (h *Handler) PublishAction(ctx context.Context, c *app.RequestContext) {
	req := &api.DouyinPublishActionRequest{}
	rpcReq := &rpc.DouyinPublishActionRequest{}
	if h.Pre(ctx, c, req, rpcReq) {
		token := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
		rpcReq.Token = token[1]
		req.Token = token[1]
		rpcResp, err := h.RPC().Base().Client().PublishAction(ctx, rpcReq)
		if err != nil {
			return
		}
		resp := rpc.DouyinPublishActionResponse{}
		h.After(ctx, c, &resp, rpcResp, err)
	}
}

// UserMsg
// @router /douyin/user/ [GET]
func (h *Handler) UserMsg(ctx context.Context, c *app.RequestContext) {
	req := &api.DouyinUserRequest{}
	rpcReq := &rpc.DouyinUserRequest{}
	if h.Pre(ctx, c, req, rpcReq) {
		userID, ok := c.Get("userid")
		if !ok {
			return
		}
		username, ok := c.Get("username")
		if !ok {
			return
		}
		rpcReq.BaseReq = new(model.BaseReq)
		rpcReq.BaseReq.UserId = new(int64)
		rpcReq.BaseReq.Username = new(string)
		*rpcReq.BaseReq.UserId = userID.(int64)
		*rpcReq.BaseReq.Username = username.(string)
		rpcResp, err := h.RPC().Base().Client().UserMsg(ctx, rpcReq)
		if err != nil {
			return
		}
		resp := rpc.DouyinUserResponse{}
		h.After(ctx, c, &resp, rpcResp, err)
	}
}

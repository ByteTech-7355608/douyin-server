package handler

import (
	"ByteTech-7355608/douyin-server/rpc"
	"ByteTech-7355608/douyin-server/util"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"reflect"
)

type Handler struct {
	rpc *rpc.RPC
}

func (h *Handler) RPC() *rpc.RPC {
	return h.rpc
}

func NewHandler(rpc *rpc.RPC) *Handler {
	return &Handler{
		rpc: rpc,
	}
}

// Pre 绑定参数，并且copy到rpcReq
// rpcReq: nil 或 指针
func (h *Handler) Pre(ctx context.Context, c *app.RequestContext, req interface{}, rpcReq interface{}) (ok bool) {
	// TODO 鉴权？

	if err := c.BindAndValidate(req); err != nil {
		logrus.Errorf("bind to %v error: %v", req, err)
		c.String(consts.StatusBadRequest, err.Error())
		return false
	}
	logrus.Infof("req %T: %v", req, util.LogStr(req))
	if rpcReq != nil {
		if err := copier.Copy(rpcReq, req); err != nil {
			logrus.Errorf("copy from %T to %T error: %v", req, rpcReq, err)
			c.String(consts.StatusBadRequest, err.Error())
			return false
		}
		logrus.Infof("rpc req %T: %v", rpcReq, util.LogStr(rpcReq))
	}
	return true
}

// After resp has to be pointer, optional thriftResp and err is from rpcImpl response
func (h *Handler) After(ctx context.Context, c *app.RequestContext, resp interface{}, rpcResp interface{}, err error) (ok bool) {
	if h.afterData(ctx, c, resp, rpcResp, err); resp != nil {
		logrus.Infof("api response: %v", util.LogStr(resp))
		c.JSON(consts.StatusOK, resp)
	}
	return true
}

// AfterData 只返回resp的data字段
func (h *Handler) afterData(ctx context.Context, c *app.RequestContext, resp interface{}, rpcResp interface{}, err error) {
	logrus.Infof("resp type: %T, rpcResp type: %T err: %v", resp, rpcResp, err)
	if rpcResp != nil && !reflect.ValueOf(rpcResp).IsZero() {
		if err := copier.Copy(resp, rpcResp); err != nil {
			logrus.Errorf("copy from %T to %T error: %v", rpcResp, resp, err)
			c.String(consts.StatusInternalServerError, err.Error())
			return
		}
	}
	return
}

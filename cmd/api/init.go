package api

import (
	"ByteTech-7355608/douyin-server/cmd/api/biz/handler"
	"ByteTech-7355608/douyin-server/cmd/api/biz/router"
	"ByteTech-7355608/douyin-server/rpc"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func NewDouyinApiHertz() *server.Hertz {
	svc := server.Default()
	h := handler.NewHandler(rpc.NewRPC())
	router.Register(svc, h)
	return svc
}

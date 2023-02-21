package api

import (
	"ByteTech-7355608/douyin-server/cmd/api/biz/handler"
	"ByteTech-7355608/douyin-server/cmd/api/biz/router"
	"ByteTech-7355608/douyin-server/dal/cache"
	"ByteTech-7355608/douyin-server/pkg/tracer"
	"ByteTech-7355608/douyin-server/rpc"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	hertztracer "github.com/hertz-contrib/tracer/hertz"
)

func NewDouyinApiHertz() *server.Hertz {
	hTracer, _ := tracer.InitTracer("douyin.api")
	cache.NewRedisCache()
	svc := server.Default(
		server.WithMaxRequestBodySize(30*1024*1024),
		server.WithTracer(hertztracer.NewTracer(hTracer, func(c *app.RequestContext) string {
			return "hertz.server" + "::" + c.FullPath()
		})))
	svc.Use(hertztracer.ServerCtx())
	h := handler.NewHandler(rpc.NewRPC())
	router.Register(svc, h)
	return svc
}

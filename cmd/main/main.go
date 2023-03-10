package main

import (
	"ByteTech-7355608/douyin-server/cmd/api"
	"ByteTech-7355608/douyin-server/cmd/handlers"
	"ByteTech-7355608/douyin-server/dal/cache"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/base/baseservice"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/interaction/interactionservice"
	"ByteTech-7355608/douyin-server/kitex_gen/douyin/social/socialservice"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"ByteTech-7355608/douyin-server/pkg/tracer"
	"ByteTech-7355608/douyin-server/rpc"
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/go-co-op/gocron"
	etcd "github.com/kitex-contrib/registry-etcd"
	opentracing "github.com/kitex-contrib/tracer-opentracing"
)

func Init() {
	InitLogger()
}

// main 服务入口，一个main启动多个服务
func main() {
	Init()
	psm := os.Getenv("ServiceName")
	var svr server.Server
	switch psm {
	case constants.APIServiceName:
		startDouyinApi()
		return
	case constants.BaseServiceName:
		svr = startDouyinBase()
	case constants.InteractionServiceName:
		svr = startDouyinInteraction()
	case constants.SocialServiceName:
		svr = startDouyinSocial()
	case constants.CronServiceName:
		// 启动定时任务
		s := gocron.NewScheduler(time.Local)
		_, _ = s.Every(5).Minutes().Do(cache.SyncDataToDB)
		s.StartBlocking()
		return
	}
	Log.Infof("start service: %s", psm)
	if svr == nil {
		panic(fmt.Sprintf("no server for (%s) to run, support PSM: %v", psm,
			[]string{constants.APIServiceName, constants.BaseServiceName, constants.InteractionServiceName}))
	}
	err := svr.Run()
	if err != nil {
		panic(err)
	}
}

func startDouyinApi() {
	svc := api.NewDouyinApiHertz()
	Log.Infof("start service: %s", constants.APIServiceName)
	svc.Spin()
}

func startDouyinBase() (svr server.Server) {
	svc := handlers.NewBaseServiceImpl()
	svc.Init(rpc.NewRPC())
	svr = baseservice.NewServer(svc, serverOptions(constants.BaseServiceName, constants.BaseTCPAddr)...)
	return
}

func startDouyinInteraction() (svr server.Server) {
	svc := handlers.NewInteractionServiceImpl()
	svc.Init(rpc.NewRPC())
	svr = interactionservice.NewServer(svc, serverOptions(constants.InteractionServiceName, constants.InteractionTCPAddr)...)
	return
}

func startDouyinSocial() (svr server.Server) {
	svc := handlers.NewSocialServiceImpl()
	svc.Init(rpc.NewRPC())
	svr = socialservice.NewServer(svc, serverOptions(constants.SocialServiceName, constants.SocialTCPAddr)...)
	return
}

// serverOptions server启动配置
func serverOptions(serverName, tcpAddr string) (opts []server.Option) {
	kTracer, _ := tracer.InitTracer("kitex.server::" + serverName)
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}
	//Init()
	opts = []server.Option{
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serverName}), // server name
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithSuite(opentracing.NewServerSuite(kTracer, func(ctx context.Context) string {
			endpoint := rpcinfo.GetRPCInfo(ctx).To()
			return endpoint.ServiceName() + "::" + endpoint.Method()
		})), // tracer
		//server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
		server.WithRegistry(r),
	}
	return
}

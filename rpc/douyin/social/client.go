package socialcli

import (
	svc "ByteTech-7355608/douyin-server/kitex_gen/douyin/social/socialservice"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)

//go:generate mockgen -destination rpc/douyin/socialcli/mock_client.go -package socialcli -source kitex_gen/douyin/social/socialservice/client.go  Client

func GetKitexClient(opts ...client.Option) svc.Client {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	// 配置服务发现、超时时间等
	return svc.MustNewClient(
		constants.SocialServiceName,
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		//client.WithSuite(trace.NewDefaultClientSuite()), // tracer
		client.WithResolver(r), // resolver
	)
}

type Client struct {
	cli svc.Client
}

func NewClient(cli svc.Client) *Client {
	if cli == nil {
		cli = GetKitexClient()
	}
	return &Client{
		cli: cli,
	}
}

func (t *Client) Client() svc.Client {
	if t != nil {
		return t.cli
	}
	return nil
}

func (t *Client) MockClient() *MockClient {
	if t == nil || t.cli == nil {
		return nil
	}
	if v, ok := t.cli.(*MockClient); ok {
		return v
	}
	return nil
}

var _ svc.Client = new(MockClient)

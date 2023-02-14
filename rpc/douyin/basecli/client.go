package basecli

import (
	svc "ByteTech-7355608/douyin-server/kitex_gen/douyin/base/baseservice"
	"ByteTech-7355608/douyin-server/pkg/constants"

	"github.com/cloudwego/kitex/client"
)

//go:generate mockgen -destination rpc/douyin/basecli/mock_client.go -package basecli -source kitex_gen/douyin/base/baseservice/client.go  Client

func GetKitexClient(opts ...client.Option) svc.Client {
	return svc.MustNewClient(constants.BaseServiceName, opts...)
}

type Client struct {
	cli svc.Client
}

func NewClient(cli svc.Client, opts ...client.Option) *Client {
	if cli == nil {
		cli = GetKitexClient(opts...)
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
	if t == nil {
		return nil
	}
	if v, ok := t.cli.(*MockClient); ok {
		return v
	}
	return nil
}

var _ svc.Client = new(MockClient)

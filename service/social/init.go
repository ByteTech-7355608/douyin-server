package social

import (
	"ByteTech-7355608/douyin-server/dal/cache"
	"ByteTech-7355608/douyin-server/dal/dao"
	"ByteTech-7355608/douyin-server/rpc"
)

type Service struct {
	dao   *dao.Dao
	rpc   *rpc.RPC
	cache *cache.RedisCache
}

func NewService(rpc *rpc.RPC) *Service {
	return &Service{
		dao:   dao.NewDao(),
		rpc:   rpc,
		cache: cache.NewRedisCache(),
	}
}

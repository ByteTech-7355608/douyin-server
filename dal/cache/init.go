package cache

import (
	"ByteTech-7355608/douyin-server/pkg/constants"

	"github.com/redis/go-redis/v9"
)

var cli *redis.Client

type RedisCache struct {
	User     User
	Video    Video
	Like     Like
	Relation Relation
}

func NewRedisCache() *RedisCache {
	if cli == nil {
		cli = redis.NewClient(&redis.Options{
			Addr:     constants.RedisIPPort,
			Password: "",
			DB:       0, // 默认DB 0
		})
	}
	return &RedisCache{
		User:     User{},
		Video:    Video{},
		Like:     Like{},
		Relation: Relation{},
	}
}

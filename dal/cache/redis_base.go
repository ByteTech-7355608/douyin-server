package cache

import (
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/redis/go-redis/v9"
)

var redisCli *redis.Client

type RedisCache struct {
	cli *redis.Client
}

func NewRedisCache() *RedisCache {
	if redisCli == nil {
		redisCli = redis.NewClient(&redis.Options{
			Addr:     constants.RedisIPPort,
			Password: "",
			DB:       0, // 默认DB 0
		})
	}
	return &RedisCache{
		cli: redisCli,
	}
}

// 将缓存的值设置到dest中
func (r *RedisCache) setDest(ctx context.Context, dest, val interface{}) error {
	vd := reflect.ValueOf(dest)
	vvd := reflect.ValueOf(val)
	if val == nil || vvd.IsNil() {
		return nil
	}
	vd, vvd = reflect.Indirect(vd), reflect.Indirect(vvd)
	if vd.CanSet() && vd.Type() == vvd.Type() {
		vd.Set(vvd)
	} else {
		return fmt.Errorf("dest %T in not settable or type mismatch for %T", dest, val)
	}
	return nil
}

func (r *RedisCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (ok bool) {
	if value == nil {
		return false
	}
	bs, err := json.Marshal(value)
	if err != nil {
		Log.Warnf("marshal value err: %v", err)
		return false
	}
	if expiration == 0 {
		expiration = time.Hour
	}
	if status := r.cli.Set(ctx, key, bs, expiration); status.Err() != nil {
		Log.Errorf("set %v to redis err: %v", key, status.Err())
		return false
	}
	return true
}

// Get 注意dest需要为指针
func (r *RedisCache) Get(ctx context.Context, key string, dest interface{}) (err error) {
	bs, err := r.cli.Get(ctx, key).Bytes()
	if err != nil {
		Log.Warnf("get %v from redis err: %v", key, err)
		return err
	}
	err = json.Unmarshal(bs, dest)
	if err != nil {
		Log.Errorf("unmarshal err: %v", err)
		return err
	}
	return nil
}

func (r *RedisCache) Delete(ctx context.Context, key string) (ok bool) {
	if status := r.cli.Del(ctx, key); status.Err() != nil {
		Log.Errorf("delete %v from redis err: %v", key, status.Err())
		return false
	}
	return true
}

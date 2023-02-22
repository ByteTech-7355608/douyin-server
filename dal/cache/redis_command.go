package cache

import (
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"context"
	"encoding/json"
	"time"
)

// ==========通用操作============

func Exists(ctx context.Context, key ...string) int64 {
	return cli.Exists(ctx, key...).Val()
}

// ==========String操作============

func Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (ok bool) {
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
	if status := cli.Set(ctx, key, bs, expiration); status.Err() != nil {
		Log.Errorf("set %v to redis err: %v", key, status.Err())
		return false
	}
	return true
}

// Get 注意dest需要为指针
func Get(ctx context.Context, key string, dest interface{}) (err error) {
	bs, err := cli.Get(ctx, key).Bytes()
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

func Delete(ctx context.Context, key string) (ok bool) {
	if status := cli.Del(ctx, key); status.Err() != nil {
		Log.Errorf("delete %v from redis err: %v", key, status.Err())
		return false
	}
	return true
}

func Incr(ctx context.Context, key string) (cnt int64, err error) {
	cnt, err = cli.Incr(ctx, key).Result()
	if err != nil {
		Log.Warnf("Incr %v redis err: %v", key, err)
		return
	}
	return
}

// ==========Hash操作============

func HSet(ctx context.Context, key string, value interface{}) (ok bool) {
	if status := cli.HSet(ctx, key, value); status.Err() != nil {
		Log.Errorf("set %v to redis err: %v", key, status.Err())
		return false
	}
	return true
}

// HGetAll 返回值为map，需要调用方自行转换
func HGetAll(ctx context.Context, key string) map[string]string {
	return cli.HGetAll(ctx, key).Val()
}

func HMGet(ctx context.Context, key string, field ...string) []interface{} {
	return cli.HMGet(ctx, key, field...).Val()
}

func HIncr(ctx context.Context, key, field string, incr int64) (ok bool) {
	if status := cli.HIncrBy(ctx, key, field, incr); status.Err() != nil {
		Log.Errorf("incr %v %v err: %v", key, field, status.Err())
		return false
	}
	return true
}

func HKeys(ctx context.Context, key string) []string {
	return cli.HKeys(ctx, key).Val()
}

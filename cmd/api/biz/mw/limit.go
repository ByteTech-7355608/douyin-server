package mw

import (
	"ByteTech-7355608/douyin-server/cmd/api/biz/handler"
	"ByteTech-7355608/douyin-server/dal/cache"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

func IPLimitMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		clientip := c.ClientIP()
		key := constants.GetFavoriteLmtKey(clientip)
		cache.NewRedisCache()
		if cache.Exists(ctx, key) == 0 {
			fmt.Println("2")
			// redis中不存在当前key
			cache.Set(ctx, key, 1, time.Minute)
		} else {
			fmt.Println("3")
			cnt, err := cache.Incr(ctx, key)
			if err != nil {
				fmt.Println("limite,err:", err)
				handler.Response(ctx, c, constants.ErrWriteCache)
				c.Abort()
				return
			}
			if cnt >= 20 {
				// 操作过于频繁
				handler.Response(ctx, c, constants.ErrIPLimited)
				c.Abort()
				return
			}
		}

		c.Next(ctx)
	}
}

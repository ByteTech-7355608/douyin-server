package mw

import (
	"ByteTech-7355608/douyin-server/cmd/api/biz/handler"
	"ByteTech-7355608/douyin-server/dal/cache"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

func IPLimitMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		clientip := c.ClientIP()
		key := constants.GetFavoriteLmtKey(clientip)
		if cache.Exists(ctx, key) == 0 {
			// redis中不存在当前key
			cache.Set(ctx, key, 1, time.Minute)
		} else {
			cnt, _ := cache.Incr(ctx, key)
			if cnt >= constants.Limits_per_sec {
				// 操作过于频繁
				handler.Response(ctx, c, constants.ErrIPLimited)
				c.Abort()
				return
			}
		}

		c.Next(ctx)
	}
}

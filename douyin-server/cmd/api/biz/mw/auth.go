package mw

import (
	"ByteTech-7355608/douyin-server/cmd/api/biz/handler"
	. "ByteTech-7355608/douyin-server/pkg/configs"
	"ByteTech-7355608/douyin-server/pkg/constants"
	"ByteTech-7355608/douyin-server/pkg/jwt"
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
)

func JWTAuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		auth := c.Query("token")
		// URL中为检测到token
		if auth == "" {
			auth = c.PostForm("token")
		}

		mc, err := jwt.ParseToken(auth)
		if err != nil {
			if errors.Is(err, constants.ErrTokenExpires) {
				handler.Response(ctx, c, constants.ErrTokenExpires)
			} else {
				handler.Response(ctx, c, constants.ErrInvalidToken)
			}
			c.Abort()
			return
		}

		// 将当前请求的username信息保存到请求的上下文c上
		// 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
		c.Set("userid", mc.UserID)
		c.Set("username", mc.Username)
		c.Next(ctx)
	}
}

func FeedAuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		auth := c.Query("token")
		if auth == "" {
			return
		}

		mc, err := jwt.ParseToken(auth)
		if err != nil {
			Log.Infof("鉴权失败: %v", err)
			return
		}

		// 将当前请求的username信息保存到请求的上下文c上
		// 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
		c.Set("userid", mc.UserID)
		c.Set("username", mc.Username)
	}
}

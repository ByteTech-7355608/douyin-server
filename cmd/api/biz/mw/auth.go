package mw

import (
	"ByteTech-7355608/douyin-server/pkg/constants"
	"ByteTech-7355608/douyin-server/pkg/jwt"
	"context"
	"errors"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
)

func JWTAuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// Authorization: Bearer XXXX.XXXX.XXXXX
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			constants.Response(ctx, c, constants.CodeNotLogin)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			constants.Response(ctx, c, constants.CodeInvalidAuth)
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			if errors.Is(err, constants.ErrTokenExpires) {
				constants.Response(ctx, c, constants.CodeTokenExpires)
			} else {
				constants.Response(ctx, c, constants.CodeInvalidToken)
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

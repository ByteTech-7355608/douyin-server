package constants

import "errors"

var (
	ErrUserExist       = errors.New("用户名已存在")
	ErrUserNotExist    = errors.New("用户名不存在")
	ErrInvalidPassword = errors.New("用户密码错误")
	ErrTokenExpires    = errors.New("Token已过期")
)

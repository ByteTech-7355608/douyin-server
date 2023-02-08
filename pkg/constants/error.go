package constants

import "errors"

type RespStatus struct {
	StatusCode int32
	Err        error
}

func (s *RespStatus) Error() string {
	return s.Err.Error()
}

func (s *RespStatus) Errormsg() *string {
	er := s.Err.Error()
	return &er
}

var (
	ErrUserExist       = &RespStatus{201, errors.New("用户名已存在")}
	ErrUserNotExist    = &RespStatus{202, errors.New("用户名不存在")}
	ErrInvalidPassword = &RespStatus{203, errors.New("用户密码错误")}
	ErrTokenExpires    = &RespStatus{204, errors.New("token已过期")}
	ErrNotLogin        = &RespStatus{205, errors.New("未登录")}
	ErrInvalidAuth     = &RespStatus{206, errors.New("认证格式有误")}
	ErrInvalidToken    = &RespStatus{207, errors.New("无效的Token")}
)

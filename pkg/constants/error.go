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
	ErrUserExist            = &RespStatus{4001, errors.New("用户名已存在")}
	ErrUserNotExist         = &RespStatus{4002, errors.New("用户名不存在")}
	ErrInvalidPassword      = &RespStatus{4003, errors.New("用户密码错误")}
	ErrTokenExpires         = &RespStatus{4004, errors.New("token已过期")}
	ErrNotLogin             = &RespStatus{4005, errors.New("未登录")}
	ErrInvalidToken         = &RespStatus{4007, errors.New("无效的Token")}
	ErrCreateRecord         = &RespStatus{4008, errors.New("创建记录错误")}
	ErrQueryRecord          = &RespStatus{4009, errors.New("查询记录错误")}
	ErrUpdateRecord         = &RespStatus{4010, errors.New("更新记录错误")}
	ErrDeleteRecord         = &RespStatus{4011, errors.New("删除记录错误")}
	ErrUnsupportedOperation = &RespStatus{4012, errors.New("不支持的操作")}
	ErrUserNameOverSize     = &RespStatus{4013, errors.New("用户名长度应小于32")}
	ErrPassWordOverSize     = &RespStatus{4014, errors.New("密码长度应大于5")}
	ErrPassWordBelowSize    = &RespStatus{4014, errors.New("密码长度应小于32")}
	ErrPassWordSymbols      = &RespStatus{4014, errors.New("密码需包含字母、数字及特殊字符")}
)

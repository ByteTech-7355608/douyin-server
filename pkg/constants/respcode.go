package constants

type ResCode int

const (
	CodeSignupSuccess ResCode = 200 + iota
	CodeSigninSuccess
	CodeOpeartSuccess
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeTokenExpires
	CodeInvalidToken
	CodeNotLogin
	CodeInvalidAuth
	CodeServerBusy
)

var codeMsgmap = map[ResCode]string{
	CodeSignupSuccess:   "注册成功",
	CodeSigninSuccess:   "登录成功",
	CodeOpeartSuccess:   "操作成功",
	CodeUserExist:       "用户名已存在",
	CodeUserNotExist:    "用户名不存在",
	CodeInvalidPassword: "用户密码错误",
	CodeTokenExpires:    "Token已过期",
	CodeInvalidToken:    "无效的Token",
	CodeNotLogin:        "未登录",
	CodeInvalidAuth:     "认证格式有误",
	CodeServerBusy:      "服务器繁忙",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgmap[c]
	if !ok {
		msg = codeMsgmap[CodeServerBusy]
	}
	return msg
}

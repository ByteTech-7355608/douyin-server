package util

import (
	"ByteTech-7355608/douyin-server/pkg/constants"
	"regexp"
)

// 检查密码是否同时包含大小写、数字、特殊字符，且长度5-32
// true->检验成功，false->检验失败
func CheckPassword(password string) (bool, *constants.RespStatus) {
	pats := []string{"[a-zA-Z]", "[0-9]", "[^\\d\\w]"}
	switch {
	case len(password) < constants.PasswordMinLen:
		return false, constants.ErrPassWordBelowSize
	case len(password) > constants.PassWordMaxLen:
		return false, constants.ErrPassWordOverSize
	}
	for _, pat := range pats {
		if ok, _ := regexp.MatchString(pat, password); !ok {
			return false, constants.ErrPassWordSymbols
		}
	}
	return true, nil
}

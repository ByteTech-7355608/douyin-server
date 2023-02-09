package util

import (
	"crypto/md5"
	"encoding/hex"
)

const secret = "7355608"

// EncryptPassword 使用md5加盐加密
func EncryptPassword(pwd string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(pwd)))
}

package jwt

import (
	"ByteTech-7355608/douyin-server/pkg/constants"
	"errors"

	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySecret = []byte("douyin service")

// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段s
// 假设我们这里需要额外记录一个userid字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	// 可根据需要自行添加字段
	UserID               int64  `json:"user_id"`
	Username             string `json:"username"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

// GenToken 生成JWT
func GenToken(userid int64, username string) (aToken string, err error) {
	// 实例化一个我们带创建的加密声明
	aclaims := MyClaims{
		// 自定义字段
		userid,
		username,
		jwt.RegisteredClaims{
			// 过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(1))),
			// 签发人
			Issuer: "bluebell",
		},
	}

	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, aclaims).SignedString(mySecret)
	return
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	var mc = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			err = constants.ErrTokenExpires
		}
		return nil, err
	}
	// 校验token
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}

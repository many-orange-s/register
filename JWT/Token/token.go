package Token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"time"
)

// GetToken token的获取
func GetToken(department string) (string, error) {
	msg := Token{
		department,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDurnTime).Unix(), // 过期时间
			Issuer:    "go_project",                               // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, msg)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析token
func ParseToken(tokenstring string) (*Token, error) {
	//验证密码
	tokenmsg := new(Token)
	token, err := jwt.ParseWithClaims(tokenstring, tokenmsg, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}

	//二次判断
	if token.Valid {
		return tokenmsg, nil
	}
	return nil, errors.New("invalid token")
}

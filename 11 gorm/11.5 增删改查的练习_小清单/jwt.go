package main

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const TokenExpireDuration = time.Second * 10
// MySecret 用于加盐的字符串
var MySecret = []byte("夏天夏天悄悄过去")
type MyClaims struct {
	Uid int64 `json:"uid"`
	Name string `json:"name"`
	jwt.StandardClaims
}


// GenToken 生成JWT
func GenToken(uid int64,name string) (string, error) {
	// 创建一个我们自己的声明
	claims := MyClaims{
		uid, // 自定义字段
		name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "todo-app", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil //加盐字符串通过这个函数返回
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
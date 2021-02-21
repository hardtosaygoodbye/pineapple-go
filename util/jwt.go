package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/cast"
)

// CreateToken 创建Token
func CreateToken(uid uint, secret string) (string, error) {

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": cast.ToString(uid),
		"exp": time.Now().Add(time.Hour * 24 * 9999).Unix(),
	})
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

// ParseToken 解析Token
func ParseToken(token string, secret string) (string, error) {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}
	return claim.Claims.(jwt.MapClaims)["uid"].(string), err
}

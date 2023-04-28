package jwtx

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"metaserver/pkg/changetool"
)

// GetToken 申请token
func GetToken(secretKey string, iat int64, seconds int64, uid string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["uid"] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

// AuthToken 验证token
func AuthToken(secretKey string, accessToken string) (string, error) {
	// 解析 token
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Verification failed 1")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return "", errors.New("Verification failed 2")
	}

	// 验证token是否有效
	if !token.Valid {
		return "", errors.New("token lose efficacy")
	}
	cliams, _ := token.Claims.(jwt.MapClaims)

	return changetool.InterfaceToString(cliams["uid"]), nil
}

package jwtx

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type Claims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

// CreateToken 签发用户Token
func CreateToken(userId string, AccessSecret string, AccessExpire int64) (string, uint64, error) {
	expiresAt := time.Now().Unix() + AccessExpire
	claims := Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "IkunHouse-app",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(AccessSecret))
	return token, uint64(expiresAt), err
}

// ValidToken 验证用户token
// bool: 是否过期 default: true 过期
// error: 解析是否成功 default: nil
func ValidToken(token string, AccessSecret string) (bool, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(AccessSecret), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			expiresTime := claims.ExpiresAt
			now := time.Now().Unix()
			if now > expiresTime {
				//token过期了
				return true, nil
			} else {
				return false, nil
			}
		}
	}
	return true, err
}

// GetUserIDFormToken 从token中获取用户id
func GetUserIDFormToken(token string, AccessSecret string) (string, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(AccessSecret), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims.UserId, nil
		}
	}
	return "", err
}

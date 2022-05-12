package auth

import (
	"errors"
	"time"

	"api/config"

	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

const (
	TokenExpireDuration = time.Minute * 5
)

var ErrInvalidToken = errors.New("invalid jwt token")

func GenToken(username string) (string, error) {
	c := MyClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "qianxi",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString([]byte(config.MySecretKey))
}

func ParseToken(tokenStr string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.MySecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

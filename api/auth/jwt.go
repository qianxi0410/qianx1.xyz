package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

const (
	TokenExpireDuration = time.Minute * 5
	MySecretKey         = ""
)

var ErrInvalidToken = errors.New("invalid jwt token")

func GenToken(username string) (string, error) {
	c := MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "qianxi",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString([]byte(MySecretKey))
}

func ParseToken(tokenStr string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(MySecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/felixlambertv/go-cleanplate/internal/model"
	"github.com/golang-jwt/jwt"
)

type TokenClaims struct {
	jwt.StandardClaims
	Authorized bool        `json:"authorized"`
	User       *model.User `json:"user"`
	Expire     int64       `json:"expire"`
}

type Token struct {
	Token   string
	Expires time.Time
}

func GenerateToken(user *model.User, lifespan int, secret string) (*Token, error) {
	token := &Token{}

	expTime := time.Now().Add(time.Hour * time.Duration(lifespan))

	claims := TokenClaims{}
	claims.Authorized = true
	claims.User = user
	claims.Expire = expTime.Unix()
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, errSign := unsignedToken.SignedString([]byte(secret))

	token.Token = signedToken
	token.Expires = expTime

	return token, errSign
}

func ParseToken(tokenString string, secret string) (*TokenClaims, error) {
	claims := &TokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token invalid")
}

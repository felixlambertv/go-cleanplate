package utils

import (
	"time"

	"github.com/felixlambertv/go-cleanplate/config"
	"github.com/felixlambertv/go-cleanplate/internal/model"
	"github.com/golang-jwt/jwt"
)

type TokenStruct struct {
	Token   string
	Expires time.Time
}

func GenerateToken(user model.User) (TokenStruct, error) {
	conf := config.GetInstance()

	token := TokenStruct{}

	expTime := time.Now().Add(time.Hour * time.Duration(conf.App.TokenLifespan))

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user.ID
	claims["expire"] = expTime.Unix()
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, errSign := unsignedToken.SignedString([]byte(conf.App.Secret))

	token.Token = signedToken
	token.Expires = expTime

	return token, errSign
}

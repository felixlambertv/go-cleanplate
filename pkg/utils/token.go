package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/felixlambertv/go-cleanplate/config"
	"github.com/felixlambertv/go-cleanplate/internal/model"
	"github.com/gin-gonic/gin"
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
	claims["user"] = user
	claims["expire"] = expTime.Unix()
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, errSign := unsignedToken.SignedString([]byte(conf.App.Secret))

	token.Token = signedToken
	token.Expires = expTime

	return token, errSign
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	conf := config.GetInstance()
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(conf.App.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return parsedToken, nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")

	// Apple already reserved header for Authorization
	// https://developer.apple.com/documentation/foundation/nsurlrequest
	if bearerToken == "" {
		bearerToken = c.Request.Header.Get("X-Authorization")
	}

	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractUserLevel(token *jwt.Token) (int, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		ulevel, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_level"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return int(ulevel), nil
	}
	return 0, nil
}

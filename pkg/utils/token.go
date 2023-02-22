package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/felixlambertv/go-cleanplate/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Token struct {
	Token   string
	Expires time.Time
}

func GenerateToken(user *model.User, lifespan int, secret string) (*Token, error) {
	token := &Token{}

	expTime := time.Now().Add(time.Hour * time.Duration(lifespan))

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user"] = user
	claims["expire"] = expTime.Unix()
	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, errSign := unsignedToken.SignedString([]byte(secret))

	token.Token = signedToken
	token.Expires = expTime

	return token, errSign
}

func ParseToken(tokenString string, secret string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
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

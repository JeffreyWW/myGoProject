package jwtTool

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var key = []byte("crfChina")

func CreateToken() string {
	claim := jwt.StandardClaims{
		NotBefore: time.Now().Unix(),
		ExpiresAt: time.Now().Unix() + 1000,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, _ := token.SignedString(key)
	return tokenString

}

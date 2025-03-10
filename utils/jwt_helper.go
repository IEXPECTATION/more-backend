package utils

import (
	"crypto/rand"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var key []byte

func init() {
	key = make([]byte, 32)
	rand.Read(key)
}

func GenerateJWT(method jwt.SigningMethod, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(method, claims)
	return token.SignedString(key)
}

func ParseJWT(tokenString string, options ...jwt.ParserOption) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	}, options...)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token.Claims, nil

}

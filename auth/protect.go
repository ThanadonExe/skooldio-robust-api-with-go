package auth

import (
	"fmt"
	"strings"

	jwt "github.com/golang-jwt/jwt/v4"
)

func Protect(token string) error {
	tokenString := strings.TrimPrefix(token, "Bearer ")

	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("==Signature=="), nil
	})

	return err
}

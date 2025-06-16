package utils

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(tokenString string, secretOrPublicKey []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// Kiểm tra thuật toán ký có đúng không (ví dụ HMAC)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretOrPublicKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Parse thành công, kiểm tra claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

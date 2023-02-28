package jwtParser

import (
	"github.com/golang-jwt/jwt/v4"
)

type Content struct {
	XId string `json:"_id"`
	jwt.RegisteredClaims
}

// Redeem returns _id, error message (if applicable)
func (secret *JwtSecret) Redeem(userToken string) (string, error) {
	token, err := jwt.ParseWithClaims(userToken, &Content{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret.phrase), nil
	})
	if err == nil {
		if claims, ok := token.Claims.(*Content); ok && token.Valid {
			return claims.XId, nil
		}
	}

	return "", err
}

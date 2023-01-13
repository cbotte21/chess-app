package jwt

import (
	"github.com/golang-jwt/jwt/v4"
)

type Content struct {
	Id string `json:"email"`
	jwt.RegisteredClaims
}

// Redeem returns the _id belonging to a jwt. nil on success
func (secret *JwtSecret) Redeem(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Content{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret.phrase), nil
	})
	if err == nil {
		if claims, ok := token.Claims.(*Content); ok && token.Valid {
			return claims.Id, nil
		}
	}

	return "", err
}

package jwtParser

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

type Content struct {
	XId string `json:"_id"`
	jwt.RegisteredClaims
}

// Redeem returns _id, error message (if applicable)
func (secret JwtSecret) Redeem(userToken string) (string, error) {
	token, err := jwt.ParseWithClaims(userToken, &Content{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err == nil {
		if claims, ok := token.Claims.(*Content); ok && token.Valid {
			return claims.XId, nil
		}
	}

	return "", err
}

// ValidateJWT returns nil if the token is valid
func (secret JwtSecret) ValidateJWT(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, errors.New("could not parse token")
	})

	if err != nil {
		return errors.New("could not parse token")
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("could not claim token")
}

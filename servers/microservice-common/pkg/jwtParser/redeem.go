package jwtParser

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

// Redeem returns the contents of a jwt
func (secret JwtSecret) Redeem(userToken string) (*JwtContent, error) {
	token, err := jwt.ParseWithClaims(userToken, &JwtContent{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err == nil {
		if claims, ok := token.Claims.(*JwtContent); ok && token.Valid {
			return claims, nil
		}
	}

	return nil, err
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

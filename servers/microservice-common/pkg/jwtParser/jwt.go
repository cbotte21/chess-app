package jwtParser

import (
	"github.com/cbotte21/microservice-common/pkg/schema"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtContent struct {
	Id   string `json:"_id"`
	Role int    `json:"role"`
	jwt.RegisteredClaims
}

const EXPIRY_HOURS time.Duration = 14

type JwtSecret string

func (secret JwtSecret) GenerateJWT(user schema.User) (string, error) { //time.Now().Unix() + int64(60*60*EXPIRY_HOURS)
	claims := JwtContent{
		user.Id,
		user.Role,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(EXPIRY_HOURS * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "cbotte21",
			Subject:   "jwt",
			ID:        "1",
			Audience:  []string{"client"},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	println()
	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}

package utilities

import (
	"errors"
	"github.com/cbotte21/auth-go/internal/schema"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const PHRASE string = "mysupersecretjwtphrase" //TODO: Export to main
const EXPIRY_HOURS time.Duration = 14

type JwtContent struct {
	Id   string `json:"_id"`
	Role int    `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(user schema.User) (string, error) { //time.Now().Unix() + int64(60*60*EXPIRY_HOURS)
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
	tokenString, err := token.SignedString([]byte(PHRASE))
	return tokenString, err
}

func ValidateJWT(tokenString string) error { //Returns role, and if valid token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(PHRASE), nil
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

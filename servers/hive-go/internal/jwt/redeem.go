package jwt

import (
	"github.com/cbotte21/hive-go/internal/playerbase"
	"github.com/golang-jwt/jwt/v4"
)

type Content struct {
	Id   string `json:"_id"`
	Role int    `json:"role"`
	jwt.RegisteredClaims
}

// Redeem returns the _id belonging to a jwt. nil on success
func (secret *JwtSecret) Redeem(tokenStr string) (playerbase.Player, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Content{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret.phrase), nil
	})
	if err == nil {
		if claims, ok := token.Claims.(*Content); ok && token.Valid {
			return playerbase.Player{Id: claims.Id, Role: claims.Role}, nil
		}
	}

	return playerbase.Player{}, err
}

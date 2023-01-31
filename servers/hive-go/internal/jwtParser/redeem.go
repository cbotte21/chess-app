package jwtParser

import (
	"github.com/cbotte21/hive-go/internal/playerbase"
	"github.com/cbotte21/hive-go/pb"
	"github.com/golang-jwt/jwt/v4"
)

type Content struct {
	XId string `json:"_id"`
	jwt.RegisteredClaims
}

// Redeem returns the _id belonging to a jwtParser. nil on success
func (secret *JwtSecret) Redeem(request *pb.Jwt) (playerbase.Player, error) {
	token, err := jwt.ParseWithClaims(request.GetJwt(), &Content{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret.phrase), nil
	})
	if err == nil {
		if claims, ok := token.Claims.(*Content); ok && token.Valid {
			return playerbase.Player{XID: claims.XId, Jwt: request.Jwt}, nil
		}
	}

	return playerbase.Player{}, err
}

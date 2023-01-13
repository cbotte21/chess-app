package pb

import (
	"golang.org/x/net/context"
	"log"
)

type Server struct {
}

func (server *Server) Join(ctx context.Context, joinRequest *JoinRequest) *Player {
	log.Printf("Join requested: %s", joinRequest.Jwt)
	return &Player{XId: "players id"}
}

func (server *Server) Disconnect(ctx context.Context, player *Player) *Void {
	log.Printf("Player has disconnected: %s", player.XId)
	return &Void{}
}

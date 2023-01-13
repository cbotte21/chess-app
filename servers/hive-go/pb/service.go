package pb

import (
	"golang.org/x/net/context"
	"log"
)

type HiveServer struct {
}

func (hive *HiveServer) Join(ctx context.Context, joinRequest *JoinRequest) *Player {
	log.Printf("Join requested: %s", joinRequest.Jwt)
	return &Player{XId: "players id"}
}

func (hive *HiveServer) Disconnect(ctx context.Context, player *Player) *Void {
	log.Printf("Player has disconnected: %s", player.XId)
	return &Void{}
}

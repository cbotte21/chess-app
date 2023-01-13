package pb

import (
	"golang.org/x/net/context"
	"log"
)

type HiveServer struct {
	UnimplementedHiveServiceServer
}

func (hive *HiveServer) Join(ctx context.Context, joinRequest *JoinRequest) (*Player, error) {
	log.Printf("Join requested: %s", joinRequest.Jwt)
	return &Player{XId: "players id"}, nil
}

func (hive *HiveServer) Disconnect(ctx context.Context, player *Player) (*Void, error) {
	log.Printf("Player has disconnected: %s", player.XId)
	return &Void{}, nil
}

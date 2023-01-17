package internal

import (
	"github.com/cbotte21/hive-go/internal/jwt"
	"github.com/cbotte21/hive-go/internal/playerbase"
	"github.com/cbotte21/hive-go/pb"
	"golang.org/x/net/context"
)

type Hive struct {
	PlayerBase  *playerbase.PlayerBase
	JwtRedeemer *jwt.JwtSecret
	pb.UnimplementedHiveServiceServer
}

func NewHive(playerBase *playerbase.PlayerBase, jwtRedeemer *jwt.JwtSecret) Hive {
	return Hive{PlayerBase: playerBase, JwtRedeemer: jwtRedeemer}
}

// Join appends {_id, jwt} to the active players, if joinRequest.jwt is valid
func (hive *Hive) Join(ctx context.Context, joinRequest *pb.JoinRequest) (*pb.Player, error) {
	player, err := hive.JwtRedeemer.Redeem(joinRequest.Jwt)
	if err == nil {
		player.Jwt = joinRequest.Jwt
		hive.PlayerBase.AppendUnique(player)
		return &pb.Player{XId: player.Id}, err //TODO: Add role
	}
	return &pb.Player{}, err
}

// Disconnect removes the player from the PlayerBase
func (hive *Hive) Disconnect(ctx context.Context, player *pb.Player) (*pb.Void, error) {
	hive.PlayerBase.Disconnect(player.XId)
	return &pb.Void{}, nil
}

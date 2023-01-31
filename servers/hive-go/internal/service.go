package internal

import (
	"github.com/cbotte21/hive-go/internal/jwtParser"
	"github.com/cbotte21/hive-go/internal/playerbase"
	"github.com/cbotte21/hive-go/pb"
	"golang.org/x/net/context"
)

type Hive struct {
	PlayerBase  *playerbase.PlayerBase
	JwtRedeemer *jwtParser.JwtSecret
	pb.UnimplementedHiveServiceServer
}

func NewHive(playerBase *playerbase.PlayerBase, jwtRedeemer *jwtParser.JwtSecret) Hive {
	return Hive{PlayerBase: playerBase, JwtRedeemer: jwtRedeemer}
}

// Join appends {_id, jwtParser} to the active players, if joinRequest.jwtParser is valid
func (hive *Hive) Join(ctx context.Context, jwt *pb.Jwt) (*pb.Player, error) {
	player, err := hive.JwtRedeemer.Redeem(jwt)
	if err == nil {
		hive.PlayerBase.AppendUnique(player)
		return &pb.Player{XId: player.XID}, err
	}
	return &pb.Player{}, err
}

// Disconnect removes the player from the PlayerBase
func (hive *Hive) Disconnect(ctx context.Context, player *pb.Player) (*pb.Void, error) {
	hive.PlayerBase.Disconnect(playerbase.Player{XID: player.XId})
	return &pb.Void{}, nil
}

// Online returns true if a player is online
func (hive *Hive) Online(ctx context.Context, player *pb.Player) (*pb.Bool, error) {
	_, err := hive.PlayerBase.GetId(playerbase.Player{Jwt: player.XId})
	if err != nil {
		return &pb.Bool{Status: 0}, err
	}
	return &pb.Bool{Status: 1}, err
}

func (hive *Hive) Redeem(ctx context.Context, jwt *pb.Jwt) (*pb.Player, error) {
	id, err := hive.PlayerBase.GetId(playerbase.Player{Jwt: jwt.GetJwt()})
	return &pb.Player{XId: id}, err
}

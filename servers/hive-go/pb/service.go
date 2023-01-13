package pb

import (
	"github.com/cbotte21/hive-go/internal/jwt"
	"github.com/cbotte21/hive-go/internal/playerbase"
	"golang.org/x/net/context"
)

type Hive struct {
	PlayerBase  *playerbase.PlayerBase
	JwtRedeemer *jwt.JwtSecret
	UnimplementedHiveServiceServer
}

func NewHive(playerBase *playerbase.PlayerBase, jwtRedeemer *jwt.JwtSecret) Hive {
	return Hive{PlayerBase: playerBase, JwtRedeemer: jwtRedeemer}
}

// Join appends {_id, jwt} to the active players, if joinRequest.jwt is valid
func (hive *Hive) Join(ctx context.Context, joinRequest *JoinRequest) (*Player, error) {
	id, err := hive.JwtRedeemer.Redeem(joinRequest.Jwt)
	if err == nil {
		hive.PlayerBase.AppendUnique(playerbase.Player{Id: id, Jwt: joinRequest.Jwt})
		return &Player{XId: id}, err
	}
	return &Player{}, err
}

func (hive *Hive) Disconnect(ctx context.Context, player *Player) (*Void, error) {
	hive.PlayerBase.Disconnect(player.XId)
	return &Void{}, nil
}

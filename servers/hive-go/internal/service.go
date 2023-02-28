package internal

import (
	"errors"
	"github.com/cbotte21/hive-go/internal/jwtParser"
	"github.com/cbotte21/hive-go/internal/playerbase"
	"github.com/cbotte21/hive-go/pb"
	judicial "github.com/cbotte21/judicial-go/pb"
	"golang.org/x/net/context"
)

type Hive struct {
	PlayerBase     *playerbase.PlayerBase
	JwtRedeemer    *jwtParser.JwtSecret
	JudicialClient *judicial.JudicialServiceClient
	pb.UnimplementedHiveServiceServer
}

func NewHive(playerBase *playerbase.PlayerBase, jwtRedeemer *jwtParser.JwtSecret, judicialClient *judicial.JudicialServiceClient) Hive {
	return Hive{PlayerBase: playerBase, JwtRedeemer: jwtRedeemer, JudicialClient: judicialClient}
}

// Join appends {_id, jwtParser} to the active players, if joinRequest.jwtParser is valid
func (hive *Hive) Join(ctx context.Context, joinRequest *pb.JoinRequest) (*pb.JoinResponse, error) {
	_id, err := hive.JwtRedeemer.Redeem(joinRequest.GetJwt())
	if err == nil {
		integrity, err := (*hive.JudicialClient).Integrity(ctx, &judicial.IntegrityRequest{XId: _id})
		if err == nil {
			if integrity.Status {
				hive.PlayerBase.AppendUnique(joinRequest.GetJwt(), _id)
				return &pb.JoinResponse{Status: 1}, nil
			}
		}
	}
	return &pb.JoinResponse{Status: 0}, err
}

// Disconnect removes the player from the PlayerBase
func (hive *Hive) Disconnect(ctx context.Context, disconnectRequest *pb.DisconnectRequest) (*pb.DisconnectResponse, error) {
	hive.PlayerBase.Disconnect(disconnectRequest.GetJwt())
	return &pb.DisconnectResponse{}, nil
}

// Online returns true if a player is online
func (hive *Hive) Online(ctx context.Context, onlineRequest *pb.OnlineRequest) (*pb.OnlineResponse, error) {
	if hive.PlayerBase.Online(onlineRequest.XId) {
		return &pb.OnlineResponse{Status: 1}, nil
	}
	return &pb.OnlineResponse{Status: 0}, errors.New("onlineRequest is not online")
}

// Redeem returns the _id pertaining to a player
func (hive *Hive) Redeem(ctx context.Context, redeemRequest *pb.RedeemRequest) (*pb.RedeemResponse, error) {
	id, err := hive.PlayerBase.GetId(redeemRequest.GetJwt())
	return &pb.RedeemResponse{XId: id}, err
}

// Role returns a player's Role
func (hive *Hive) Role(ctx context.Context, roleRequest *pb.RoleRequest) (*pb.RoleResponse, error) {
	role, err := hive.PlayerBase.Role(roleRequest.GetJwt())
	return &pb.RoleResponse{Value: role}, err
}

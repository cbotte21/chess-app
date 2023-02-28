package internal

import (
	"context"
	"errors"
	hive "github.com/cbotte21/hive-go/pb"
	"github.com/cbotte21/judicial-go/internal/datastore"
	"github.com/cbotte21/judicial-go/internal/schema"
	pb "github.com/cbotte21/judicial-go/pb"
	"time"
)

type Judicial struct {
	HiveClient *hive.HiveServiceClient
	pb.UnimplementedJudicialServiceServer
}

func NewJudicial(hiveClient *hive.HiveServiceClient) Judicial {
	return Judicial{HiveClient: hiveClient}
}

func canBan(role int32) bool {
	return role > 0
}

func (judicial *Judicial) Ban(ctx context.Context, banRequest *pb.BanRequest) (*pb.BanResponse, error) {
	role, err := (*judicial.HiveClient).Role(context.Background(), &hive.RoleRequest{Jwt: banRequest.GetGod()})
	admin, err2 := (*judicial.HiveClient).Redeem(context.Background(), &hive.RedeemRequest{Jwt: banRequest.GetGod()})

	if err == nil && err2 == nil {
		if canBan(role.GetValue()) {
			err := datastore.Create(schema.Ban{
				Player:    banRequest.XId,
				God:       admin.GetXId(),
				Reason:    banRequest.GetReason(),
				Expiry:    banRequest.GetExpiry().String(),
				Timestamp: time.Now().String(),
			})
			if err == nil { //Success
				//TODO: Disconnect player from hive-go
				return &pb.BanResponse{Status: true}, nil
			}
			println(err.Error())
		}
	}
	return &pb.BanResponse{Status: false}, err
}

func (judicial *Judicial) Unban(ctx context.Context, unbanRequest *pb.UnbanRequest) (*pb.UnbanResponse, error) {
	role, err := (*judicial.HiveClient).Role(context.Background(), &hive.RoleRequest{Jwt: unbanRequest.GetGod()})
	admin, err2 := (*judicial.HiveClient).Redeem(context.Background(), &hive.RedeemRequest{Jwt: unbanRequest.GetGod()})

	if err == nil && err2 == nil {
		if canBan(role.GetValue()) {
			err := datastore.Delete(schema.Ban{Player: unbanRequest.GetXId()})
			if err == nil {
				err := datastore.Create(schema.Unban{
					Player:    unbanRequest.GetXId(),
					God:       admin.GetXId(),
					Timestamp: time.Now().String(),
				})
				if err == nil { //Success
					return &pb.UnbanResponse{Status: true}, nil
				}
			}
		}
	}
	return &pb.UnbanResponse{Status: false}, err
}

func (judicial *Judicial) Integrity(ctx context.Context, integrityRequest *pb.IntegrityRequest) (*pb.IntegrityResponse, error) {
	ban, err := datastore.Find(schema.Ban{Player: integrityRequest.GetXId()})
	if err != nil { // Player is not banned
		return &pb.IntegrityResponse{Status: true}, nil
	}
	return &pb.IntegrityResponse{Status: false}, errors.New("integrityRequest is banned until " + ban.Expiry)
}

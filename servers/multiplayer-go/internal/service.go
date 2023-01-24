package internal

import (
	"github.com/cbotte21/multiplayer-go/internal/players"
	"github.com/cbotte21/multiplayer-go/pb"
	"golang.org/x/net/context"
	"io"
)

type Multiplayer struct {
	players *players.Players
	pb.UnimplementedMultiplayerServiceServer
}

func NewMultiplayer() Multiplayer {
	return Multiplayer{}
}

func (multiplayer *Multiplayer) Join(ctx context.Context, stream *pb.MultiplayerService_PlayServer) error {
	var player *players.Player
	defer func() {
		multiplayer.players.Leave(*player)
	}()
	for {
		playerInfo, err := stream.Recv()
		if player == nil { //Must register player
			multiplayer.players.Join(playerInfo)
			player = multiplayer.players.Get(playerInfo)
		}

		player.Update(playerInfo.GetX(), playerInfo.GetY(), playerInfo.GetTexture(), playerInfo.GetJwt())

		if player.GetTick()%10 == 1 { //Only update data every 10 ticks
			for _, player := range *multiplayer.players.GetPlayers() {
				playerLocation := pb.PlayerLocation{
					Texture: player.GetTexture(),
					X:       player.GetX(),
					Y:       player.GetY(),
				}
				if err := stream.Send(playerLocation); err != nil {
					return err
				}
			}
		}

		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
	}
}

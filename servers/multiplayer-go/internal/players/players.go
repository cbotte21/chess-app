package players

import "github.com/cbotte21/multiplayer-go/pb"

type Players []Player

func (players *Players) Get(playerInfo pb.PlayerInfo) *Player {
	for i := range *players {
		if (*players)[i].GetToken() == playerInfo.GetJwt() {
			return &(*players)[i]
		}
	}
	return nil
}

func (players *Players) Join(playerInfo *pb.PlayerInfo) {
	*players = append(*players, Player{
		x:       playerInfo.GetX(),
		y:       playerInfo.GetY(),
		token:   playerInfo.GetJwt(),
		texture: playerInfo.GetTexture(),
	})
}

func (players *Players) Leave(player Player) {
	for i := range *players {
		if player.token == (*players)[i].token {
			*players = append((*players)[:i], (*players)[i+1:]...)
			break
		}
	}
}

func (players *Players) GetPlayers() *[]Player {
	return (*[]Player)(players)
}

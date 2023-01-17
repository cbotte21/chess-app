package playerbase

import "log"

type Player struct {
	Id, Jwt string
	Role    int
}

type PlayerBase []Player

// AppendUnique follows a delete existing policy
func (playerBase *PlayerBase) AppendUnique(player Player) {
	playerBase.Disconnect(player.Id)
	*playerBase = append(*playerBase, player)
}

func (playerBase *PlayerBase) Disconnect(id string) {
	for i := range *playerBase {
		if (*playerBase)[i].Id == id {
			*playerBase = append((*playerBase)[:i], (*playerBase)[i+1:]...)
			return
		}
	}
}

func (playerBase *PlayerBase) Log() {
	var i int
	for i = range *playerBase {
		(*playerBase)[i].Log()
	}
	log.Printf("%d active players.", i)
}

func (player Player) Log() {
	log.Printf(player.Jwt + " " + player.Id)
}

package playerbase

type Player struct {
	Id, Jwt string
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

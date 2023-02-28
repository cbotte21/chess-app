package playerbase

import (
	"errors"
)

type PlayerBase map[string]Player

type Player struct {
	Id   string
	Role int32 //Access level
}

// AppendUnique follows a delete existing policy
func (playerBase *PlayerBase) AppendUnique(jwt, _id string) {
	playerBase.Disconnect(_id)
	(*playerBase)[jwt] = Player{Id: _id, Role: 1} //TODO: set Role on append
}

// Disconnect removes a player from the active players map
func (playerBase *PlayerBase) Disconnect(_id string) {
	delete(*playerBase, _id)
}

func (playerBase *PlayerBase) Online(_id string) bool {
	for k := range *playerBase {
		if (*playerBase)[k].Id == _id {
			return true
		}
	}
	return false
}

// GetId returns the XID belonging to a player
func (playerBase *PlayerBase) GetId(jwt string) (string, error) {
	player, found := (*playerBase)[jwt]
	if found {
		return player.Id, nil
	}
	return "", errors.New("player is not online")
}

// Role returns a player's Role
func (playerBase *PlayerBase) Role(jwt string) (int32, error) {
	player := (*playerBase)[jwt]
	if player.Id != "" {
		return player.Role, nil
	}
	return 0, errors.New("player not online")
}

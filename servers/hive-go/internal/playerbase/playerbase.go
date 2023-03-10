package playerbase

import (
	"errors"
)

type PlayerBase map[string]string

// AppendUnique follows a delete existing policy
func (playerBase *PlayerBase) AppendUnique(jwt, _id string) {
	playerBase.Disconnect(_id)
	(*playerBase)[jwt] = _id
}

// Disconnect removes a player from the active players map
func (playerBase *PlayerBase) Disconnect(_id string) {
	delete(*playerBase, _id)
}

// Online returns true if a player is online
func (playerBase *PlayerBase) Online(_id string) bool {
	for k := range *playerBase {
		if (*playerBase)[k] == _id {
			return true
		}
	}
	return false
}

// GetId returns the XID belonging to a player
func (playerBase *PlayerBase) GetId(jwt string) (string, error) {
	player, found := (*playerBase)[jwt]
	if found {
		return player, nil
	}
	return "", errors.New("player is not online")
}

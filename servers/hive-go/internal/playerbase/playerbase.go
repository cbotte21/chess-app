package playerbase

import "errors"

type PlayerBase map[string]string

type Player struct {
	XID, Jwt string
}

// AppendUnique follows a delete existing policy
func (playerBase *PlayerBase) AppendUnique(player Player) {
	playerBase.Disconnect(player)
	(*playerBase)[player.Jwt] = player.XID
}

// Disconnect removes a player from the active players map
func (playerBase *PlayerBase) Disconnect(player Player) {
	delete(*playerBase, player.XID)
}

// GetId returns the XID belonging to a player
func (playerBase *PlayerBase) GetId(player Player) (string, error) {
	id, found := (*playerBase)[player.Jwt]
	if found {
		return id, nil
	}
	return "", errors.New("player is not online")
}

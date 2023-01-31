package internal

type Archive struct {
	//UnimplementedArchiveServiceServer
}

func NewArchive() Archive {
	return Archive{}
}

/*
// Join appends {_id, jwtParser} to the active players, if joinRequest.jwtParser is valid
func (hive *Hive) Join(ctx context.Context, joinRequest *JoinRequest) (*Player, error) {
	player, err := hive.JwtRedeemer.Redeem(joinRequest.Jwt)
	if err == nil {
		player.Jwt = joinRequest.Jwt
		hive.PlayerBase.AppendUnique(player)
		return &Player{XId: player.Id}, err //TODO: Add role
	}
	return &Player{}, err
}

func (hive *Hive) Disconnect(ctx context.Context, player *Player) (*Void, error) {
	hive.PlayerBase.Disconnect(player.XId)
	return &Void{}, nil
}
*/

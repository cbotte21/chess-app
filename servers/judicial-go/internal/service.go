package internal

type Judicial struct {
	//UnimplementedJudicialServiceServer
}

func NewJudicial() Judicial {
	return Judicial{}
}

// Join appends {_id, jwt} to the active players, if joinRequest.jwt is valid
/*
func (hive *Hive) Join(ctx context.Context, joinRequest *JoinRequest) (*Player, error) {
	return &Player{}, err
}

func (hive *Hive) Disconnect(ctx context.Context, player *Player) (*Void, error) {
	return &Void{}, nil
}
*/

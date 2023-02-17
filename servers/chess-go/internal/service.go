package internal

import (
	"context"
	"errors"
	"github.com/cbotte21/chess-go/pb"
)

type Chess struct {
	pb.UnimplementedChessServiceServer
}

func NewChess() Chess {
	return Chess{}
}

func (chess *Chess) Move(ctx context.Context, moveRequest *pb.MoveRequest) (*pb.Bool, error) {
	return &pb.Bool{Status: true}, errors.New("")
}

func (chess *Chess) Update(jwt *pb.Jwt, stream pb.ChessService_UpdateServer) error {

	return errors.New("")
}

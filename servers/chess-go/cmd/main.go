package main

import (
	"github.com/cbotte21/chess-go/internal"
	"github.com/cbotte21/chess-go/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

const (
	PORT int = 9002
)

func main() {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if err != nil {
		log.Fatalf("Failed to listen on port: %d", PORT)
	}
	grpcServer := grpc.NewServer()

	//Initialize hive
	chessServer := internal.NewChess()

	pb.RegisterChessServiceServer(grpcServer, &chessServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to initialize grpc server.")
	}
}

package main

import (
	"github.com/cbotte21/hive-go/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

const (
	PORT int = 9000
)

func main() {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if err != nil {
		log.Fatalf("Failed to listen on port: %d", PORT)
	}
	hiveService := pb.Server{}
	server := grpc.NewServer()

	pb.RegisterChatServiceServer(server, &hiveService)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to initialize grpc server.")
	}
}

package main

import (
	"github.com/cbotte21/hive-go/internal"
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

	//Register handlers to attach

	//Initialize hive
	_ = internal.NewQueue()

	//pb.RegisterQueueServiceServer(grpcServer, &hive)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to initialize grpc server.")
	}
}

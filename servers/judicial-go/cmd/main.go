package main

import (
	"github.com/cbotte21/hive-go/internal"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

const (
	PORT int = 9001
)

func main() {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if err != nil {
		log.Fatalf("Failed to listen on port: %d", PORT)
	}
	grpcServer := grpc.NewServer()

	//Register handler(s) to attach

	//Initialize judicial
	_ = internal.NewJudicial()

	//pb.RegisterJudicialServiceServer(grpcServer, &hive)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to initialize grpc server.")
	}
}

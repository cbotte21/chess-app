package main

import (
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
	server := grpc.NewServer()

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to initialize grpc server.")
	}
}

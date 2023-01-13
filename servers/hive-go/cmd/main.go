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
	hive := pb.HiveServer{}
	grpcServer := grpc.NewServer()

	pb.RegisterHiveServiceServer(grpcServer, &hive)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to initialize grpc server.")
	}
}

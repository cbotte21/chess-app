package main

import (
	"github.com/cbotte21/hive-go/internal"
	"github.com/cbotte21/hive-go/internal/jwtParser"
	"github.com/cbotte21/hive-go/internal/playerbase"
	"github.com/cbotte21/hive-go/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

const (
	PORT   int    = 9000
	SECRET string = "mysupersecretjwtphrase"
)

func main() {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if err != nil {
		log.Fatalf("Failed to listen on port: %d", PORT)
	}
	grpcServer := grpc.NewServer()

	//Register handlers to attach
	playerBase := playerbase.PlayerBase{}
	jwtRedeemer := jwtParser.NewJwtSecret(SECRET)
	//Initialize hive
	hive := internal.NewHive(&playerBase, &jwtRedeemer)

	pb.RegisterHiveServiceServer(grpcServer, &hive)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to initialize grpc server.")
	}
}

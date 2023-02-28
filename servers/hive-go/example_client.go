package main

import (
	"github.com/cbotte21/hive-go/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//TODO: Test should use HTTP client to dynamically obtain JWT

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		println(err)
		return
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	hive := pb.NewHiveServiceClient(conn)

	player, err := hive.Join(context.Background(), &pb.Jwt{Jwt: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2Mzc1NzhjOTZlMTQ0YjQxYzJhOWMzYjkiLCJyb2xlIjowLCJpc3MiOiJjYm90dGUyMSIsInN1YiI6Imp3dCIsImF1ZCI6WyJjbGllbnQiXSwiZXhwIjoxNjc3NjYwMjc4LCJuYmYiOjE2Nzc2MDk4NzgsImlhdCI6MTY3NzYwOTg3OCwianRpIjoiMSJ9.plDGQKyxHT5XMxT3bN7etVLoa-mJTlKN2EKlF9_39qg"})
	if err != nil {
		println(err)
		return
	}
	println("success... ")
	println(player.Status)
}

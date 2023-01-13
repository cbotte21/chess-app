package tests

import (
	"github.com/cbotte21/hive-go/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"testing"
)

func TestJoin(t *testing.T) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer func(conn *grpc.ClientConn) {
		_ = conn.Close()
	}(conn)

	hive := pb.NewHiveServiceClient(conn)

	player, err := hive.Join(context.Background(), &pb.JoinRequest{Jwt: "myjwt"})
	if err != nil {
		t.Error(err.Error())
	}
	println(player.XId)
}

package internal

import (
	"context"
	"github.com/cbotte21/queue-go/internal/adt_queue"
	pb "github.com/cbotte21/queue-go/pb"
)

type Queue struct {
	AdtQueue *adt_queue.AdtQueue
	pb.UnimplementedQueueServiceServer
}

func NewQueue(adtQueue *adt_queue.AdtQueue) Queue {
	return Queue{AdtQueue: adtQueue}
}

func (queue *Queue) Join(ctx context.Context, joinRequest *pb.JoinRequest) (*pb.JoinResponse, error) {
	return &pb.JoinResponse{Status: 1}, nil
}

func (queue *Queue) Leave(ctx context.Context, leaveRequest *pb.LeaveRequest) (*pb.LeaveResponse, error) {
	return &pb.LeaveResponse{}, nil
}

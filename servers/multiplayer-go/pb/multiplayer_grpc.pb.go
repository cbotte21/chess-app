// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: multiplayer.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MultiplayerServiceClient is the client API for MultiplayerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MultiplayerServiceClient interface {
	Play(ctx context.Context, opts ...grpc.CallOption) (MultiplayerService_PlayClient, error)
}

type multiplayerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMultiplayerServiceClient(cc grpc.ClientConnInterface) MultiplayerServiceClient {
	return &multiplayerServiceClient{cc}
}

func (c *multiplayerServiceClient) Play(ctx context.Context, opts ...grpc.CallOption) (MultiplayerService_PlayClient, error) {
	stream, err := c.cc.NewStream(ctx, &MultiplayerService_ServiceDesc.Streams[0], "/MultiplayerService/Play", opts...)
	if err != nil {
		return nil, err
	}
	x := &multiplayerServicePlayClient{stream}
	return x, nil
}

type MultiplayerService_PlayClient interface {
	Send(*PlayerInfo) error
	Recv() (*PlayerLocation, error)
	grpc.ClientStream
}

type multiplayerServicePlayClient struct {
	grpc.ClientStream
}

func (x *multiplayerServicePlayClient) Send(m *PlayerInfo) error {
	return x.ClientStream.SendMsg(m)
}

func (x *multiplayerServicePlayClient) Recv() (*PlayerLocation, error) {
	m := new(PlayerLocation)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MultiplayerServiceServer is the server API for MultiplayerService service.
// All implementations must embed UnimplementedMultiplayerServiceServer
// for forward compatibility
type MultiplayerServiceServer interface {
	Play(MultiplayerService_PlayServer) error
	mustEmbedUnimplementedMultiplayerServiceServer()
}

// UnimplementedMultiplayerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMultiplayerServiceServer struct {
}

func (UnimplementedMultiplayerServiceServer) Play(MultiplayerService_PlayServer) error {
	return status.Errorf(codes.Unimplemented, "method Play not implemented")
}
func (UnimplementedMultiplayerServiceServer) mustEmbedUnimplementedMultiplayerServiceServer() {}

// UnsafeMultiplayerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MultiplayerServiceServer will
// result in compilation errors.
type UnsafeMultiplayerServiceServer interface {
	mustEmbedUnimplementedMultiplayerServiceServer()
}

func RegisterMultiplayerServiceServer(s grpc.ServiceRegistrar, srv MultiplayerServiceServer) {
	s.RegisterService(&MultiplayerService_ServiceDesc, srv)
}

func _MultiplayerService_Play_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MultiplayerServiceServer).Play(&multiplayerServicePlayServer{stream})
}

type MultiplayerService_PlayServer interface {
	Send(*PlayerLocation) error
	Recv() (*PlayerInfo, error)
	grpc.ServerStream
}

type multiplayerServicePlayServer struct {
	grpc.ServerStream
}

func (x *multiplayerServicePlayServer) Send(m *PlayerLocation) error {
	return x.ServerStream.SendMsg(m)
}

func (x *multiplayerServicePlayServer) Recv() (*PlayerInfo, error) {
	m := new(PlayerInfo)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MultiplayerService_ServiceDesc is the grpc.ServiceDesc for MultiplayerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MultiplayerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MultiplayerService",
	HandlerType: (*MultiplayerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Play",
			Handler:       _MultiplayerService_Play_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "multiplayer.proto",
}
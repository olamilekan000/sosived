// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: receiver.proto

package receiver

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

// MessagingClient is the client API for Messaging service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessagingClient interface {
	ProcessMessage(ctx context.Context, opts ...grpc.CallOption) (Messaging_ProcessMessageClient, error)
}

type messagingClient struct {
	cc grpc.ClientConnInterface
}

func NewMessagingClient(cc grpc.ClientConnInterface) MessagingClient {
	return &messagingClient{cc}
}

func (c *messagingClient) ProcessMessage(ctx context.Context, opts ...grpc.CallOption) (Messaging_ProcessMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &Messaging_ServiceDesc.Streams[0], "/receiver.Messaging/ProcessMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &messagingProcessMessageClient{stream}
	return x, nil
}

type Messaging_ProcessMessageClient interface {
	Send(*MessageItem) error
	CloseAndRecv() (*MessageItem, error)
	grpc.ClientStream
}

type messagingProcessMessageClient struct {
	grpc.ClientStream
}

func (x *messagingProcessMessageClient) Send(m *MessageItem) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messagingProcessMessageClient) CloseAndRecv() (*MessageItem, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MessageItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessagingServer is the server API for Messaging service.
// All implementations must embed UnimplementedMessagingServer
// for forward compatibility
type MessagingServer interface {
	ProcessMessage(Messaging_ProcessMessageServer) error
	mustEmbedUnimplementedMessagingServer()
}

// UnimplementedMessagingServer must be embedded to have forward compatible implementations.
type UnimplementedMessagingServer struct {
}

func (UnimplementedMessagingServer) ProcessMessage(Messaging_ProcessMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method ProcessMessage not implemented")
}
func (UnimplementedMessagingServer) mustEmbedUnimplementedMessagingServer() {}

// UnsafeMessagingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessagingServer will
// result in compilation errors.
type UnsafeMessagingServer interface {
	mustEmbedUnimplementedMessagingServer()
}

func RegisterMessagingServer(s grpc.ServiceRegistrar, srv MessagingServer) {
	s.RegisterService(&Messaging_ServiceDesc, srv)
}

func _Messaging_ProcessMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessagingServer).ProcessMessage(&messagingProcessMessageServer{stream})
}

type Messaging_ProcessMessageServer interface {
	SendAndClose(*MessageItem) error
	Recv() (*MessageItem, error)
	grpc.ServerStream
}

type messagingProcessMessageServer struct {
	grpc.ServerStream
}

func (x *messagingProcessMessageServer) SendAndClose(m *MessageItem) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messagingProcessMessageServer) Recv() (*MessageItem, error) {
	m := new(MessageItem)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Messaging_ServiceDesc is the grpc.ServiceDesc for Messaging service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Messaging_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "receiver.Messaging",
	HandlerType: (*MessagingServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ProcessMessage",
			Handler:       _Messaging_ProcessMessage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "receiver.proto",
}
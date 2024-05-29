// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: talk.proto

package gotalk

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

const (
	ChatService_Send_FullMethodName     = "/gotalk.ChatService/Send"
	ChatService_Listener_FullMethodName = "/gotalk.ChatService/Listener"
)

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	Send(ctx context.Context, opts ...grpc.CallOption) (ChatService_SendClient, error)
	Listener(ctx context.Context, in *SessionId, opts ...grpc.CallOption) (ChatService_ListenerClient, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) Send(ctx context.Context, opts ...grpc.CallOption) (ChatService_SendClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], ChatService_Send_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceSendClient{stream}
	return x, nil
}

type ChatService_SendClient interface {
	Send(*ChatMessage) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type chatServiceSendClient struct {
	grpc.ClientStream
}

func (x *chatServiceSendClient) Send(m *ChatMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceSendClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) Listener(ctx context.Context, in *SessionId, opts ...grpc.CallOption) (ChatService_ListenerClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[1], ChatService_Listener_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceListenerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_ListenerClient interface {
	Recv() (*ChatMessage, error)
	grpc.ClientStream
}

type chatServiceListenerClient struct {
	grpc.ClientStream
}

func (x *chatServiceListenerClient) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	Send(ChatService_SendServer) error
	Listener(*SessionId, ChatService_ListenerServer) error
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) Send(ChatService_SendServer) error {
	return status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedChatServiceServer) Listener(*SessionId, ChatService_ListenerServer) error {
	return status.Errorf(codes.Unimplemented, "method Listener not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_Send_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).Send(&chatServiceSendServer{stream})
}

type ChatService_SendServer interface {
	SendAndClose(*Empty) error
	Recv() (*ChatMessage, error)
	grpc.ServerStream
}

type chatServiceSendServer struct {
	grpc.ServerStream
}

func (x *chatServiceSendServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceSendServer) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatService_Listener_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SessionId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).Listener(m, &chatServiceListenerServer{stream})
}

type ChatService_ListenerServer interface {
	Send(*ChatMessage) error
	grpc.ServerStream
}

type chatServiceListenerServer struct {
	grpc.ServerStream
}

func (x *chatServiceListenerServer) Send(m *ChatMessage) error {
	return x.ServerStream.SendMsg(m)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gotalk.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Send",
			Handler:       _ChatService_Send_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Listener",
			Handler:       _ChatService_Listener_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "talk.proto",
}
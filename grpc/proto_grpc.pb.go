// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ITUDatabaseClient is the client API for ITUDatabase service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ITUDatabaseClient interface {
	//This sends the data from client to server
	ClientSend(ctx context.Context, opts ...grpc.CallOption) (ITUDatabase_ClientSendClient, error)
	//This sends the data from Server to client
	ServerSend(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ITUDatabase_ServerSendClient, error)
	//This sends the data from client to server that the client joins
	Join(ctx context.Context, opts ...grpc.CallOption) (ITUDatabase_JoinClient, error)
	//This sends the data from Server to client that clients has left
	Leave(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ITUDatabase_LeaveClient, error)
}

type iTUDatabaseClient struct {
	cc grpc.ClientConnInterface
}

func NewITUDatabaseClient(cc grpc.ClientConnInterface) ITUDatabaseClient {
	return &iTUDatabaseClient{cc}
}

func (c *iTUDatabaseClient) ClientSend(ctx context.Context, opts ...grpc.CallOption) (ITUDatabase_ClientSendClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ITUDatabase_serviceDesc.Streams[0], "/ITUDatabase/ClientSend", opts...)
	if err != nil {
		return nil, err
	}
	x := &iTUDatabaseClientSendClient{stream}
	return x, nil
}

type ITUDatabase_ClientSendClient interface {
	Send(*ClientMessage) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type iTUDatabaseClientSendClient struct {
	grpc.ClientStream
}

func (x *iTUDatabaseClientSendClient) Send(m *ClientMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *iTUDatabaseClientSendClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *iTUDatabaseClient) ServerSend(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ITUDatabase_ServerSendClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ITUDatabase_serviceDesc.Streams[1], "/ITUDatabase/ServerSend", opts...)
	if err != nil {
		return nil, err
	}
	x := &iTUDatabaseServerSendClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ITUDatabase_ServerSendClient interface {
	Recv() (*ServerMessage, error)
	grpc.ClientStream
}

type iTUDatabaseServerSendClient struct {
	grpc.ClientStream
}

func (x *iTUDatabaseServerSendClient) Recv() (*ServerMessage, error) {
	m := new(ServerMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *iTUDatabaseClient) Join(ctx context.Context, opts ...grpc.CallOption) (ITUDatabase_JoinClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ITUDatabase_serviceDesc.Streams[2], "/ITUDatabase/Join", opts...)
	if err != nil {
		return nil, err
	}
	x := &iTUDatabaseJoinClient{stream}
	return x, nil
}

type ITUDatabase_JoinClient interface {
	Send(*Joining) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type iTUDatabaseJoinClient struct {
	grpc.ClientStream
}

func (x *iTUDatabaseJoinClient) Send(m *Joining) error {
	return x.ClientStream.SendMsg(m)
}

func (x *iTUDatabaseJoinClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *iTUDatabaseClient) Leave(ctx context.Context, in *Empty, opts ...grpc.CallOption) (ITUDatabase_LeaveClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ITUDatabase_serviceDesc.Streams[3], "/ITUDatabase/Leave", opts...)
	if err != nil {
		return nil, err
	}
	x := &iTUDatabaseLeaveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ITUDatabase_LeaveClient interface {
	Recv() (*Leaving, error)
	grpc.ClientStream
}

type iTUDatabaseLeaveClient struct {
	grpc.ClientStream
}

func (x *iTUDatabaseLeaveClient) Recv() (*Leaving, error) {
	m := new(Leaving)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ITUDatabaseServer is the server API for ITUDatabase service.
// All implementations must embed UnimplementedITUDatabaseServer
// for forward compatibility
type ITUDatabaseServer interface {
	//This sends the data from client to server
	ClientSend(ITUDatabase_ClientSendServer) error
	//This sends the data from Server to client
	ServerSend(*Empty, ITUDatabase_ServerSendServer) error
	//This sends the data from client to server that the client joins
	Join(ITUDatabase_JoinServer) error
	//This sends the data from Server to client that clients has left
	Leave(*Empty, ITUDatabase_LeaveServer) error
	mustEmbedUnimplementedITUDatabaseServer()
}

// UnimplementedITUDatabaseServer must be embedded to have forward compatible implementations.
type UnimplementedITUDatabaseServer struct {
}

func (UnimplementedITUDatabaseServer) ClientSend(ITUDatabase_ClientSendServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientSend not implemented")
}
func (UnimplementedITUDatabaseServer) ServerSend(*Empty, ITUDatabase_ServerSendServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerSend not implemented")
}
func (UnimplementedITUDatabaseServer) Join(ITUDatabase_JoinServer) error {
	return status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedITUDatabaseServer) Leave(*Empty, ITUDatabase_LeaveServer) error {
	return status.Errorf(codes.Unimplemented, "method Leave not implemented")
}
func (UnimplementedITUDatabaseServer) mustEmbedUnimplementedITUDatabaseServer() {}

// UnsafeITUDatabaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ITUDatabaseServer will
// result in compilation errors.
type UnsafeITUDatabaseServer interface {
	mustEmbedUnimplementedITUDatabaseServer()
}

func RegisterITUDatabaseServer(s *grpc.Server, srv ITUDatabaseServer) {
	s.RegisterService(&_ITUDatabase_serviceDesc, srv)
}

func _ITUDatabase_ClientSend_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ITUDatabaseServer).ClientSend(&iTUDatabaseClientSendServer{stream})
}

type ITUDatabase_ClientSendServer interface {
	SendAndClose(*Empty) error
	Recv() (*ClientMessage, error)
	grpc.ServerStream
}

type iTUDatabaseClientSendServer struct {
	grpc.ServerStream
}

func (x *iTUDatabaseClientSendServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *iTUDatabaseClientSendServer) Recv() (*ClientMessage, error) {
	m := new(ClientMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ITUDatabase_ServerSend_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ITUDatabaseServer).ServerSend(m, &iTUDatabaseServerSendServer{stream})
}

type ITUDatabase_ServerSendServer interface {
	Send(*ServerMessage) error
	grpc.ServerStream
}

type iTUDatabaseServerSendServer struct {
	grpc.ServerStream
}

func (x *iTUDatabaseServerSendServer) Send(m *ServerMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _ITUDatabase_Join_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ITUDatabaseServer).Join(&iTUDatabaseJoinServer{stream})
}

type ITUDatabase_JoinServer interface {
	SendAndClose(*Empty) error
	Recv() (*Joining, error)
	grpc.ServerStream
}

type iTUDatabaseJoinServer struct {
	grpc.ServerStream
}

func (x *iTUDatabaseJoinServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *iTUDatabaseJoinServer) Recv() (*Joining, error) {
	m := new(Joining)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ITUDatabase_Leave_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ITUDatabaseServer).Leave(m, &iTUDatabaseLeaveServer{stream})
}

type ITUDatabase_LeaveServer interface {
	Send(*Leaving) error
	grpc.ServerStream
}

type iTUDatabaseLeaveServer struct {
	grpc.ServerStream
}

func (x *iTUDatabaseLeaveServer) Send(m *Leaving) error {
	return x.ServerStream.SendMsg(m)
}

var _ITUDatabase_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ITUDatabase",
	HandlerType: (*ITUDatabaseServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ClientSend",
			Handler:       _ITUDatabase_ClientSend_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ServerSend",
			Handler:       _ITUDatabase_ServerSend_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Join",
			Handler:       _ITUDatabase_Join_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Leave",
			Handler:       _ITUDatabase_Leave_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc/proto.proto",
}

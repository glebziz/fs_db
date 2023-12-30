// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: store_service.proto

package store

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
	StoreV1_SetFile_FullMethodName    = "/store.StoreV1/SetFile"
	StoreV1_GetFile_FullMethodName    = "/store.StoreV1/GetFile"
	StoreV1_DeleteFile_FullMethodName = "/store.StoreV1/DeleteFile"
)

// StoreV1Client is the client API for StoreV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoreV1Client interface {
	SetFile(ctx context.Context, opts ...grpc.CallOption) (StoreV1_SetFileClient, error)
	GetFile(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (StoreV1_GetFileClient, error)
	DeleteFile(ctx context.Context, in *DeleteFileRequest, opts ...grpc.CallOption) (*DeleteFileResponse, error)
}

type storeV1Client struct {
	cc grpc.ClientConnInterface
}

func NewStoreV1Client(cc grpc.ClientConnInterface) StoreV1Client {
	return &storeV1Client{cc}
}

func (c *storeV1Client) SetFile(ctx context.Context, opts ...grpc.CallOption) (StoreV1_SetFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &StoreV1_ServiceDesc.Streams[0], StoreV1_SetFile_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &storeV1SetFileClient{stream}
	return x, nil
}

type StoreV1_SetFileClient interface {
	Send(*SetFileRequest) error
	CloseAndRecv() (*SetFileResponse, error)
	grpc.ClientStream
}

type storeV1SetFileClient struct {
	grpc.ClientStream
}

func (x *storeV1SetFileClient) Send(m *SetFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *storeV1SetFileClient) CloseAndRecv() (*SetFileResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SetFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storeV1Client) GetFile(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (StoreV1_GetFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &StoreV1_ServiceDesc.Streams[1], StoreV1_GetFile_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &storeV1GetFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StoreV1_GetFileClient interface {
	Recv() (*GetFileResponse, error)
	grpc.ClientStream
}

type storeV1GetFileClient struct {
	grpc.ClientStream
}

func (x *storeV1GetFileClient) Recv() (*GetFileResponse, error) {
	m := new(GetFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *storeV1Client) DeleteFile(ctx context.Context, in *DeleteFileRequest, opts ...grpc.CallOption) (*DeleteFileResponse, error) {
	out := new(DeleteFileResponse)
	err := c.cc.Invoke(ctx, StoreV1_DeleteFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreV1Server is the server API for StoreV1 service.
// All implementations must embed UnimplementedStoreV1Server
// for forward compatibility
type StoreV1Server interface {
	SetFile(StoreV1_SetFileServer) error
	GetFile(*GetFileRequest, StoreV1_GetFileServer) error
	DeleteFile(context.Context, *DeleteFileRequest) (*DeleteFileResponse, error)
	mustEmbedUnimplementedStoreV1Server()
}

// UnimplementedStoreV1Server must be embedded to have forward compatible implementations.
type UnimplementedStoreV1Server struct {
}

func (UnimplementedStoreV1Server) SetFile(StoreV1_SetFileServer) error {
	return status.Errorf(codes.Unimplemented, "method SetFile not implemented")
}
func (UnimplementedStoreV1Server) GetFile(*GetFileRequest, StoreV1_GetFileServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFile not implemented")
}
func (UnimplementedStoreV1Server) DeleteFile(context.Context, *DeleteFileRequest) (*DeleteFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
}
func (UnimplementedStoreV1Server) mustEmbedUnimplementedStoreV1Server() {}

// UnsafeStoreV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoreV1Server will
// result in compilation errors.
type UnsafeStoreV1Server interface {
	mustEmbedUnimplementedStoreV1Server()
}

func RegisterStoreV1Server(s grpc.ServiceRegistrar, srv StoreV1Server) {
	s.RegisterService(&StoreV1_ServiceDesc, srv)
}

func _StoreV1_SetFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StoreV1Server).SetFile(&storeV1SetFileServer{stream})
}

type StoreV1_SetFileServer interface {
	SendAndClose(*SetFileResponse) error
	Recv() (*SetFileRequest, error)
	grpc.ServerStream
}

type storeV1SetFileServer struct {
	grpc.ServerStream
}

func (x *storeV1SetFileServer) SendAndClose(m *SetFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *storeV1SetFileServer) Recv() (*SetFileRequest, error) {
	m := new(SetFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StoreV1_GetFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetFileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StoreV1Server).GetFile(m, &storeV1GetFileServer{stream})
}

type StoreV1_GetFileServer interface {
	Send(*GetFileResponse) error
	grpc.ServerStream
}

type storeV1GetFileServer struct {
	grpc.ServerStream
}

func (x *storeV1GetFileServer) Send(m *GetFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _StoreV1_DeleteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreV1Server).DeleteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreV1_DeleteFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreV1Server).DeleteFile(ctx, req.(*DeleteFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StoreV1_ServiceDesc is the grpc.ServiceDesc for StoreV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StoreV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "store.StoreV1",
	HandlerType: (*StoreV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteFile",
			Handler:    _StoreV1_DeleteFile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SetFile",
			Handler:       _StoreV1_SetFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetFile",
			Handler:       _StoreV1_GetFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "store_service.proto",
}

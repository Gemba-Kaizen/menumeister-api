// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: ping/pb/ping.proto

// This package would be the name of your microservice and it is important
// for generating pb.go files, this helps proto compiler know where to output your files

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

const (
	PingService_CreatePing_FullMethodName = "/ping.PingService/CreatePing"
	PingService_GetPing_FullMethodName    = "/ping.PingService/GetPing"
)

// PingServiceClient is the client API for PingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PingServiceClient interface {
	CreatePing(ctx context.Context, in *CreatePingRequest, opts ...grpc.CallOption) (*CreatePingResponse, error)
	GetPing(ctx context.Context, in *GetPingRequest, opts ...grpc.CallOption) (*GetPingResponse, error)
}

type pingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPingServiceClient(cc grpc.ClientConnInterface) PingServiceClient {
	return &pingServiceClient{cc}
}

func (c *pingServiceClient) CreatePing(ctx context.Context, in *CreatePingRequest, opts ...grpc.CallOption) (*CreatePingResponse, error) {
	out := new(CreatePingResponse)
	err := c.cc.Invoke(ctx, PingService_CreatePing_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pingServiceClient) GetPing(ctx context.Context, in *GetPingRequest, opts ...grpc.CallOption) (*GetPingResponse, error) {
	out := new(GetPingResponse)
	err := c.cc.Invoke(ctx, PingService_GetPing_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingServiceServer is the server API for PingService service.
// All implementations must embed UnimplementedPingServiceServer
// for forward compatibility
type PingServiceServer interface {
	CreatePing(context.Context, *CreatePingRequest) (*CreatePingResponse, error)
	GetPing(context.Context, *GetPingRequest) (*GetPingResponse, error)
	mustEmbedUnimplementedPingServiceServer()
}

// UnimplementedPingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPingServiceServer struct {
}

func (UnimplementedPingServiceServer) CreatePing(context.Context, *CreatePingRequest) (*CreatePingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePing not implemented")
}
func (UnimplementedPingServiceServer) GetPing(context.Context, *GetPingRequest) (*GetPingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPing not implemented")
}
func (UnimplementedPingServiceServer) mustEmbedUnimplementedPingServiceServer() {}

// UnsafePingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PingServiceServer will
// result in compilation errors.
type UnsafePingServiceServer interface {
	mustEmbedUnimplementedPingServiceServer()
}

func RegisterPingServiceServer(s grpc.ServiceRegistrar, srv PingServiceServer) {
	s.RegisterService(&PingService_ServiceDesc, srv)
}

func _PingService_CreatePing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServiceServer).CreatePing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PingService_CreatePing_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServiceServer).CreatePing(ctx, req.(*CreatePingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PingService_GetPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServiceServer).GetPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PingService_GetPing_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServiceServer).GetPing(ctx, req.(*GetPingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PingService_ServiceDesc is the grpc.ServiceDesc for PingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ping.PingService",
	HandlerType: (*PingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePing",
			Handler:    _PingService_CreatePing_Handler,
		},
		{
			MethodName: "GetPing",
			Handler:    _PingService_GetPing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ping/pb/ping.proto",
}
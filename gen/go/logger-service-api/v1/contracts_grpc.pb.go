// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: logger-service-api/v1/contracts.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Logs_GetLogsLevels_FullMethodName = "/contracts.v1.Logs/GetLogsLevels"
	Logs_GetLogs_FullMethodName       = "/contracts.v1.Logs/GetLogs"
)

// LogsClient is the client API for Logs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogsClient interface {
	GetLogsLevels(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetLogsLevelsResponse, error)
	GetLogs(ctx context.Context, in *GetLogsRequest, opts ...grpc.CallOption) (*GetLogsResponse, error)
}

type logsClient struct {
	cc grpc.ClientConnInterface
}

func NewLogsClient(cc grpc.ClientConnInterface) LogsClient {
	return &logsClient{cc}
}

func (c *logsClient) GetLogsLevels(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetLogsLevelsResponse, error) {
	out := new(GetLogsLevelsResponse)
	err := c.cc.Invoke(ctx, Logs_GetLogsLevels_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logsClient) GetLogs(ctx context.Context, in *GetLogsRequest, opts ...grpc.CallOption) (*GetLogsResponse, error) {
	out := new(GetLogsResponse)
	err := c.cc.Invoke(ctx, Logs_GetLogs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogsServer is the server API for Logs service.
// All implementations must embed UnimplementedLogsServer
// for forward compatibility
type LogsServer interface {
	GetLogsLevels(context.Context, *emptypb.Empty) (*GetLogsLevelsResponse, error)
	GetLogs(context.Context, *GetLogsRequest) (*GetLogsResponse, error)
	mustEmbedUnimplementedLogsServer()
}

// UnimplementedLogsServer must be embedded to have forward compatible implementations.
type UnimplementedLogsServer struct {
}

func (UnimplementedLogsServer) GetLogsLevels(context.Context, *emptypb.Empty) (*GetLogsLevelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogsLevels not implemented")
}
func (UnimplementedLogsServer) GetLogs(context.Context, *GetLogsRequest) (*GetLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLogs not implemented")
}
func (UnimplementedLogsServer) mustEmbedUnimplementedLogsServer() {}

// UnsafeLogsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogsServer will
// result in compilation errors.
type UnsafeLogsServer interface {
	mustEmbedUnimplementedLogsServer()
}

func RegisterLogsServer(s grpc.ServiceRegistrar, srv LogsServer) {
	s.RegisterService(&Logs_ServiceDesc, srv)
}

func _Logs_GetLogsLevels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogsServer).GetLogsLevels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Logs_GetLogsLevels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogsServer).GetLogsLevels(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Logs_GetLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogsServer).GetLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Logs_GetLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogsServer).GetLogs(ctx, req.(*GetLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Logs_ServiceDesc is the grpc.ServiceDesc for Logs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Logs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "contracts.v1.Logs",
	HandlerType: (*LogsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLogsLevels",
			Handler:    _Logs_GetLogsLevels_Handler,
		},
		{
			MethodName: "GetLogs",
			Handler:    _Logs_GetLogs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logger-service-api/v1/contracts.proto",
}

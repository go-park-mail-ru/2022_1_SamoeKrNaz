// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package handler

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

// AuthCheckerClient is the client API for AuthChecker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthCheckerClient interface {
	Create(ctx context.Context, in *SessionModel, opts ...grpc.CallOption) (*Nothing, error)
	Get(ctx context.Context, in *SessionValue, opts ...grpc.CallOption) (*SessionID, error)
	Delete(ctx context.Context, in *SessionValue, opts ...grpc.CallOption) (*Nothing, error)
}

type authCheckerClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthCheckerClient(cc grpc.ClientConnInterface) AuthCheckerClient {
	return &authCheckerClient{cc}
}

func (c *authCheckerClient) Create(ctx context.Context, in *SessionModel, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/handler.AuthChecker/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) Get(ctx context.Context, in *SessionValue, opts ...grpc.CallOption) (*SessionID, error) {
	out := new(SessionID)
	err := c.cc.Invoke(ctx, "/handler.AuthChecker/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authCheckerClient) Delete(ctx context.Context, in *SessionValue, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/handler.AuthChecker/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthCheckerServer is the server API for AuthChecker service.
// All implementations must embed UnimplementedAuthCheckerServer
// for forward compatibility
type AuthCheckerServer interface {
	Create(context.Context, *SessionModel) (*Nothing, error)
	Get(context.Context, *SessionValue) (*SessionID, error)
	Delete(context.Context, *SessionValue) (*Nothing, error)
	mustEmbedUnimplementedAuthCheckerServer()
}

// UnimplementedAuthCheckerServer must be embedded to have forward compatible implementations.
type UnimplementedAuthCheckerServer struct {
}

func (UnimplementedAuthCheckerServer) Create(context.Context, *SessionModel) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAuthCheckerServer) Get(context.Context, *SessionValue) (*SessionID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAuthCheckerServer) Delete(context.Context, *SessionValue) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAuthCheckerServer) mustEmbedUnimplementedAuthCheckerServer() {}

// UnsafeAuthCheckerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthCheckerServer will
// result in compilation errors.
type UnsafeAuthCheckerServer interface {
	mustEmbedUnimplementedAuthCheckerServer()
}

func RegisterAuthCheckerServer(s grpc.ServiceRegistrar, srv AuthCheckerServer) {
	s.RegisterService(&AuthChecker_ServiceDesc, srv)
}

func _AuthChecker_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/handler.AuthChecker/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).Create(ctx, req.(*SessionModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/handler.AuthChecker/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).Get(ctx, req.(*SessionValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthChecker_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthCheckerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/handler.AuthChecker/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthCheckerServer).Delete(ctx, req.(*SessionValue))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthChecker_ServiceDesc is the grpc.ServiceDesc for AuthChecker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthChecker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "handler.AuthChecker",
	HandlerType: (*AuthCheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AuthChecker_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AuthChecker_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AuthChecker_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "handler/session.proto",
}

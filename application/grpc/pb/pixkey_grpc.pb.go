// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// PixServiceClient is the client API for PixService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PixServiceClient interface {
	RegisterPixKey(ctx context.Context, in *PixKeyRegistration, opts ...grpc.CallOption) (*PixKeyCreatedResult, error)
	Find(ctx context.Context, in *PixKey, opts ...grpc.CallOption) (*PixKeyInfo, error)
}

type pixServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPixServiceClient(cc grpc.ClientConnInterface) PixServiceClient {
	return &pixServiceClient{cc}
}

func (c *pixServiceClient) RegisterPixKey(ctx context.Context, in *PixKeyRegistration, opts ...grpc.CallOption) (*PixKeyCreatedResult, error) {
	out := new(PixKeyCreatedResult)
	err := c.cc.Invoke(ctx, "/github.com.vinigracindo.codepix.PixService/RegisterPixKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pixServiceClient) Find(ctx context.Context, in *PixKey, opts ...grpc.CallOption) (*PixKeyInfo, error) {
	out := new(PixKeyInfo)
	err := c.cc.Invoke(ctx, "/github.com.vinigracindo.codepix.PixService/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PixServiceServer is the server API for PixService service.
// All implementations must embed UnimplementedPixServiceServer
// for forward compatibility
type PixServiceServer interface {
	RegisterPixKey(context.Context, *PixKeyRegistration) (*PixKeyCreatedResult, error)
	Find(context.Context, *PixKey) (*PixKeyInfo, error)
	mustEmbedUnimplementedPixServiceServer()
}

// UnimplementedPixServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPixServiceServer struct {
}

func (UnimplementedPixServiceServer) RegisterPixKey(context.Context, *PixKeyRegistration) (*PixKeyCreatedResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterPixKey not implemented")
}
func (UnimplementedPixServiceServer) Find(context.Context, *PixKey) (*PixKeyInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedPixServiceServer) mustEmbedUnimplementedPixServiceServer() {}

// UnsafePixServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PixServiceServer will
// result in compilation errors.
type UnsafePixServiceServer interface {
	mustEmbedUnimplementedPixServiceServer()
}

func RegisterPixServiceServer(s grpc.ServiceRegistrar, srv PixServiceServer) {
	s.RegisterService(&PixService_ServiceDesc, srv)
}

func _PixService_RegisterPixKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PixKeyRegistration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PixServiceServer).RegisterPixKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.vinigracindo.codepix.PixService/RegisterPixKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PixServiceServer).RegisterPixKey(ctx, req.(*PixKeyRegistration))
	}
	return interceptor(ctx, in, info, handler)
}

func _PixService_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PixKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PixServiceServer).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/github.com.vinigracindo.codepix.PixService/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PixServiceServer).Find(ctx, req.(*PixKey))
	}
	return interceptor(ctx, in, info, handler)
}

// PixService_ServiceDesc is the grpc.ServiceDesc for PixService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PixService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.vinigracindo.codepix.PixService",
	HandlerType: (*PixServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterPixKey",
			Handler:    _PixService_RegisterPixKey_Handler,
		},
		{
			MethodName: "Find",
			Handler:    _PixService_Find_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pixkey.proto",
}

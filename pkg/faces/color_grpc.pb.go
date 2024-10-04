// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.3
// source: pkg/faces/color.proto

package faces

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ColorService_Center_FullMethodName = "/ColorService/Center"
	ColorService_Edge_FullMethodName   = "/ColorService/Edge"
)

// ColorServiceClient is the client API for ColorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ColorServiceClient interface {
	Center(ctx context.Context, in *ColorRequest, opts ...grpc.CallOption) (*ColorResponse, error)
	Edge(ctx context.Context, in *ColorRequest, opts ...grpc.CallOption) (*ColorResponse, error)
}

type colorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewColorServiceClient(cc grpc.ClientConnInterface) ColorServiceClient {
	return &colorServiceClient{cc}
}

func (c *colorServiceClient) Center(ctx context.Context, in *ColorRequest, opts ...grpc.CallOption) (*ColorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ColorResponse)
	err := c.cc.Invoke(ctx, ColorService_Center_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *colorServiceClient) Edge(ctx context.Context, in *ColorRequest, opts ...grpc.CallOption) (*ColorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ColorResponse)
	err := c.cc.Invoke(ctx, ColorService_Edge_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ColorServiceServer is the server API for ColorService service.
// All implementations must embed UnimplementedColorServiceServer
// for forward compatibility.
type ColorServiceServer interface {
	Center(context.Context, *ColorRequest) (*ColorResponse, error)
	Edge(context.Context, *ColorRequest) (*ColorResponse, error)
	mustEmbedUnimplementedColorServiceServer()
}

// UnimplementedColorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedColorServiceServer struct{}

func (UnimplementedColorServiceServer) Center(context.Context, *ColorRequest) (*ColorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Center not implemented")
}
func (UnimplementedColorServiceServer) Edge(context.Context, *ColorRequest) (*ColorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Edge not implemented")
}
func (UnimplementedColorServiceServer) mustEmbedUnimplementedColorServiceServer() {}
func (UnimplementedColorServiceServer) testEmbeddedByValue()                      {}

// UnsafeColorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ColorServiceServer will
// result in compilation errors.
type UnsafeColorServiceServer interface {
	mustEmbedUnimplementedColorServiceServer()
}

func RegisterColorServiceServer(s grpc.ServiceRegistrar, srv ColorServiceServer) {
	// If the following call pancis, it indicates UnimplementedColorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ColorService_ServiceDesc, srv)
}

func _ColorService_Center_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ColorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ColorServiceServer).Center(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ColorService_Center_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ColorServiceServer).Center(ctx, req.(*ColorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ColorService_Edge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ColorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ColorServiceServer).Edge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ColorService_Edge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ColorServiceServer).Edge(ctx, req.(*ColorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ColorService_ServiceDesc is the grpc.ServiceDesc for ColorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ColorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ColorService",
	HandlerType: (*ColorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Center",
			Handler:    _ColorService_Center_Handler,
		},
		{
			MethodName: "Edge",
			Handler:    _ColorService_Edge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/faces/color.proto",
}

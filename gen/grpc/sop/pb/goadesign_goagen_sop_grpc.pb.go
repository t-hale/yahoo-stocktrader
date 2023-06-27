// Code generated with goa v3.11.3, DO NOT EDIT.
//
// sop protocol buffer definition
//
// Command:
// $ goa gen stocktrader/design

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: goadesign_goagen_sop.proto

package soppb

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
	Sop_Plan_FullMethodName = "/sop.Sop/Plan"
)

// SopClient is the client API for Sop service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SopClient interface {
	// Plan implements plan.
	Plan(ctx context.Context, in *PlanRequest, opts ...grpc.CallOption) (*PlanResponse, error)
}

type sopClient struct {
	cc grpc.ClientConnInterface
}

func NewSopClient(cc grpc.ClientConnInterface) SopClient {
	return &sopClient{cc}
}

func (c *sopClient) Plan(ctx context.Context, in *PlanRequest, opts ...grpc.CallOption) (*PlanResponse, error) {
	out := new(PlanResponse)
	err := c.cc.Invoke(ctx, Sop_Plan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SopServer is the server API for Sop service.
// All implementations must embed UnimplementedSopServer
// for forward compatibility
type SopServer interface {
	// Plan implements plan.
	Plan(context.Context, *PlanRequest) (*PlanResponse, error)
	mustEmbedUnimplementedSopServer()
}

// UnimplementedSopServer must be embedded to have forward compatible implementations.
type UnimplementedSopServer struct {
}

func (UnimplementedSopServer) Plan(context.Context, *PlanRequest) (*PlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Plan not implemented")
}
func (UnimplementedSopServer) mustEmbedUnimplementedSopServer() {}

// UnsafeSopServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SopServer will
// result in compilation errors.
type UnsafeSopServer interface {
	mustEmbedUnimplementedSopServer()
}

func RegisterSopServer(s grpc.ServiceRegistrar, srv SopServer) {
	s.RegisterService(&Sop_ServiceDesc, srv)
}

func _Sop_Plan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SopServer).Plan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sop_Plan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SopServer).Plan(ctx, req.(*PlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sop_ServiceDesc is the grpc.ServiceDesc for Sop service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sop_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sop.Sop",
	HandlerType: (*SopServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Plan",
			Handler:    _Sop_Plan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goadesign_goagen_sop.proto",
}

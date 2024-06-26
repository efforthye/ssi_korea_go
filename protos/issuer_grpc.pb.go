// protos/issuer.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: protos/issuer.proto

package protos

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
	SimpleIssuer_IssueSimpleVC_FullMethodName = "/issuer.SimpleIssuer/IssueSimpleVC"
)

// SimpleIssuerClient is the client API for SimpleIssuer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimpleIssuerClient interface {
	IssueSimpleVC(ctx context.Context, in *MsgIssueVC, opts ...grpc.CallOption) (*MsgIssueVCResponse, error)
}

type simpleIssuerClient struct {
	cc grpc.ClientConnInterface
}

func NewSimpleIssuerClient(cc grpc.ClientConnInterface) SimpleIssuerClient {
	return &simpleIssuerClient{cc}
}

func (c *simpleIssuerClient) IssueSimpleVC(ctx context.Context, in *MsgIssueVC, opts ...grpc.CallOption) (*MsgIssueVCResponse, error) {
	out := new(MsgIssueVCResponse)
	err := c.cc.Invoke(ctx, SimpleIssuer_IssueSimpleVC_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleIssuerServer is the server API for SimpleIssuer service.
// All implementations must embed UnimplementedSimpleIssuerServer
// for forward compatibility
type SimpleIssuerServer interface {
	IssueSimpleVC(context.Context, *MsgIssueVC) (*MsgIssueVCResponse, error)
	mustEmbedUnimplementedSimpleIssuerServer()
}

// UnimplementedSimpleIssuerServer must be embedded to have forward compatible implementations.
type UnimplementedSimpleIssuerServer struct {
}

func (UnimplementedSimpleIssuerServer) IssueSimpleVC(context.Context, *MsgIssueVC) (*MsgIssueVCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueSimpleVC not implemented")
}
func (UnimplementedSimpleIssuerServer) mustEmbedUnimplementedSimpleIssuerServer() {}

// UnsafeSimpleIssuerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimpleIssuerServer will
// result in compilation errors.
type UnsafeSimpleIssuerServer interface {
	mustEmbedUnimplementedSimpleIssuerServer()
}

func RegisterSimpleIssuerServer(s grpc.ServiceRegistrar, srv SimpleIssuerServer) {
	s.RegisterService(&SimpleIssuer_ServiceDesc, srv)
}

func _SimpleIssuer_IssueSimpleVC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgIssueVC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleIssuerServer).IssueSimpleVC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SimpleIssuer_IssueSimpleVC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleIssuerServer).IssueSimpleVC(ctx, req.(*MsgIssueVC))
	}
	return interceptor(ctx, in, info, handler)
}

// SimpleIssuer_ServiceDesc is the grpc.ServiceDesc for SimpleIssuer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SimpleIssuer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "issuer.SimpleIssuer",
	HandlerType: (*SimpleIssuerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IssueSimpleVC",
			Handler:    _SimpleIssuer_IssueSimpleVC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/issuer.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: pkg/proto/lookup.proto

package proto

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

// LookupClient is the client API for Lookup service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LookupClient interface {
	Query(ctx context.Context, in *LookupQueryRequest, opts ...grpc.CallOption) (*LookupQueryResponse, error)
}

type lookupClient struct {
	cc grpc.ClientConnInterface
}

func NewLookupClient(cc grpc.ClientConnInterface) LookupClient {
	return &lookupClient{cc}
}

func (c *lookupClient) Query(ctx context.Context, in *LookupQueryRequest, opts ...grpc.CallOption) (*LookupQueryResponse, error) {
	out := new(LookupQueryResponse)
	err := c.cc.Invoke(ctx, "/holocron.Lookup/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LookupServer is the server API for Lookup service.
// All implementations must embed UnimplementedLookupServer
// for forward compatibility
type LookupServer interface {
	Query(context.Context, *LookupQueryRequest) (*LookupQueryResponse, error)
	mustEmbedUnimplementedLookupServer()
}

// UnimplementedLookupServer must be embedded to have forward compatible implementations.
type UnimplementedLookupServer struct {
}

func (UnimplementedLookupServer) Query(context.Context, *LookupQueryRequest) (*LookupQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedLookupServer) mustEmbedUnimplementedLookupServer() {}

// UnsafeLookupServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LookupServer will
// result in compilation errors.
type UnsafeLookupServer interface {
	mustEmbedUnimplementedLookupServer()
}

func RegisterLookupServer(s grpc.ServiceRegistrar, srv LookupServer) {
	s.RegisterService(&Lookup_ServiceDesc, srv)
}

func _Lookup_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LookupQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LookupServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/holocron.Lookup/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LookupServer).Query(ctx, req.(*LookupQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Lookup_ServiceDesc is the grpc.ServiceDesc for Lookup service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Lookup_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "holocron.Lookup",
	HandlerType: (*LookupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _Lookup_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/lookup.proto",
}

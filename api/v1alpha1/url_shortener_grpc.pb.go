// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package genproto

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

// URLShorteningClient is the client API for URLShortening service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type URLShorteningClient interface {
	// Creates short URLs from long URLs
	ShortURLs(ctx context.Context, opts ...grpc.CallOption) (URLShortening_ShortURLsClient, error)
	BalanceURLs(ctx context.Context, in *BalanceURLsRequest, opts ...grpc.CallOption) (*BalanceURLsResponse, error)
}

type uRLShorteningClient struct {
	cc grpc.ClientConnInterface
}

func NewURLShorteningClient(cc grpc.ClientConnInterface) URLShorteningClient {
	return &uRLShorteningClient{cc}
}

func (c *uRLShorteningClient) ShortURLs(ctx context.Context, opts ...grpc.CallOption) (URLShortening_ShortURLsClient, error) {
	stream, err := c.cc.NewStream(ctx, &URLShortening_ServiceDesc.Streams[0], "/webengineering.api.v1alpha1.URLShortening/ShortURLs", opts...)
	if err != nil {
		return nil, err
	}
	x := &uRLShorteningShortURLsClient{stream}
	return x, nil
}

type URLShortening_ShortURLsClient interface {
	Send(*ShortURLsRequest) error
	Recv() (*ShortURLsResponse, error)
	grpc.ClientStream
}

type uRLShorteningShortURLsClient struct {
	grpc.ClientStream
}

func (x *uRLShorteningShortURLsClient) Send(m *ShortURLsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uRLShorteningShortURLsClient) Recv() (*ShortURLsResponse, error) {
	m := new(ShortURLsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *uRLShorteningClient) BalanceURLs(ctx context.Context, in *BalanceURLsRequest, opts ...grpc.CallOption) (*BalanceURLsResponse, error) {
	out := new(BalanceURLsResponse)
	err := c.cc.Invoke(ctx, "/webengineering.api.v1alpha1.URLShortening/BalanceURLs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// URLShorteningServer is the server API for URLShortening service.
// All implementations must embed UnimplementedURLShorteningServer
// for forward compatibility
type URLShorteningServer interface {
	// Creates short URLs from long URLs
	ShortURLs(URLShortening_ShortURLsServer) error
	BalanceURLs(context.Context, *BalanceURLsRequest) (*BalanceURLsResponse, error)
	mustEmbedUnimplementedURLShorteningServer()
}

// UnimplementedURLShorteningServer must be embedded to have forward compatible implementations.
type UnimplementedURLShorteningServer struct {
}

func (UnimplementedURLShorteningServer) ShortURLs(URLShortening_ShortURLsServer) error {
	return status.Errorf(codes.Unimplemented, "method ShortURLs not implemented")
}
func (UnimplementedURLShorteningServer) BalanceURLs(context.Context, *BalanceURLsRequest) (*BalanceURLsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BalanceURLs not implemented")
}
func (UnimplementedURLShorteningServer) mustEmbedUnimplementedURLShorteningServer() {}

// UnsafeURLShorteningServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to URLShorteningServer will
// result in compilation errors.
type UnsafeURLShorteningServer interface {
	mustEmbedUnimplementedURLShorteningServer()
}

func RegisterURLShorteningServer(s grpc.ServiceRegistrar, srv URLShorteningServer) {
	s.RegisterService(&URLShortening_ServiceDesc, srv)
}

func _URLShortening_ShortURLs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(URLShorteningServer).ShortURLs(&uRLShorteningShortURLsServer{stream})
}

type URLShortening_ShortURLsServer interface {
	Send(*ShortURLsResponse) error
	Recv() (*ShortURLsRequest, error)
	grpc.ServerStream
}

type uRLShorteningShortURLsServer struct {
	grpc.ServerStream
}

func (x *uRLShorteningShortURLsServer) Send(m *ShortURLsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uRLShorteningShortURLsServer) Recv() (*ShortURLsRequest, error) {
	m := new(ShortURLsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _URLShortening_BalanceURLs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceURLsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(URLShorteningServer).BalanceURLs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/webengineering.api.v1alpha1.URLShortening/BalanceURLs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(URLShorteningServer).BalanceURLs(ctx, req.(*BalanceURLsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// URLShortening_ServiceDesc is the grpc.ServiceDesc for URLShortening service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var URLShortening_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "webengineering.api.v1alpha1.URLShortening",
	HandlerType: (*URLShorteningServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BalanceURLs",
			Handler:    _URLShortening_BalanceURLs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ShortURLs",
			Handler:       _URLShortening_ShortURLs_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/api/v1alpha1/url_shortener.proto",
}

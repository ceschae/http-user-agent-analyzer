// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: httpanalyzer/http_analyzer.proto

package httpanalyzer

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

// HttpAnalyzerClient is the client API for HttpAnalyzer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HttpAnalyzerClient interface {
	// Returns true if the user agent string request is allowed, false otherwise
	Allow(ctx context.Context, in *AllowRequest, opts ...grpc.CallOption) (*IsAllowed, error)
}

type httpAnalyzerClient struct {
	cc grpc.ClientConnInterface
}

func NewHttpAnalyzerClient(cc grpc.ClientConnInterface) HttpAnalyzerClient {
	return &httpAnalyzerClient{cc}
}

func (c *httpAnalyzerClient) Allow(ctx context.Context, in *AllowRequest, opts ...grpc.CallOption) (*IsAllowed, error) {
	out := new(IsAllowed)
	err := c.cc.Invoke(ctx, "/http_analyzer.HttpAnalyzer/Allow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HttpAnalyzerServer is the server API for HttpAnalyzer service.
// All implementations must embed UnimplementedHttpAnalyzerServer
// for forward compatibility
type HttpAnalyzerServer interface {
	// Returns true if the user agent string request is allowed, false otherwise
	Allow(context.Context, *AllowRequest) (*IsAllowed, error)
	mustEmbedUnimplementedHttpAnalyzerServer()
}

// UnimplementedHttpAnalyzerServer must be embedded to have forward compatible implementations.
type UnimplementedHttpAnalyzerServer struct {
}

func (UnimplementedHttpAnalyzerServer) Allow(context.Context, *AllowRequest) (*IsAllowed, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Allow not implemented")
}
func (UnimplementedHttpAnalyzerServer) mustEmbedUnimplementedHttpAnalyzerServer() {}

// UnsafeHttpAnalyzerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HttpAnalyzerServer will
// result in compilation errors.
type UnsafeHttpAnalyzerServer interface {
	mustEmbedUnimplementedHttpAnalyzerServer()
}

func RegisterHttpAnalyzerServer(s grpc.ServiceRegistrar, srv HttpAnalyzerServer) {
	s.RegisterService(&HttpAnalyzer_ServiceDesc, srv)
}

func _HttpAnalyzer_Allow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpAnalyzerServer).Allow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/http_analyzer.HttpAnalyzer/Allow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpAnalyzerServer).Allow(ctx, req.(*AllowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HttpAnalyzer_ServiceDesc is the grpc.ServiceDesc for HttpAnalyzer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HttpAnalyzer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "http_analyzer.HttpAnalyzer",
	HandlerType: (*HttpAnalyzerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Allow",
			Handler:    _HttpAnalyzer_Allow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "httpanalyzer/http_analyzer.proto",
}

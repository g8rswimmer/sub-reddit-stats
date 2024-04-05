// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: protos/reddit.proto

package redditv1

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

// RedditServiceClient is the client API for RedditService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RedditServiceClient interface {
	GetSubredditMostUps(ctx context.Context, in *GetSubredditMostUpsRequest, opts ...grpc.CallOption) (*GetSubredditMostUpsResponse, error)
}

type redditServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRedditServiceClient(cc grpc.ClientConnInterface) RedditServiceClient {
	return &redditServiceClient{cc}
}

func (c *redditServiceClient) GetSubredditMostUps(ctx context.Context, in *GetSubredditMostUpsRequest, opts ...grpc.CallOption) (*GetSubredditMostUpsResponse, error) {
	out := new(GetSubredditMostUpsResponse)
	err := c.cc.Invoke(ctx, "/reddit.v1.RedditService/GetSubredditMostUps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RedditServiceServer is the server API for RedditService service.
// All implementations should embed UnimplementedRedditServiceServer
// for forward compatibility
type RedditServiceServer interface {
	GetSubredditMostUps(context.Context, *GetSubredditMostUpsRequest) (*GetSubredditMostUpsResponse, error)
}

// UnimplementedRedditServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRedditServiceServer struct {
}

func (UnimplementedRedditServiceServer) GetSubredditMostUps(context.Context, *GetSubredditMostUpsRequest) (*GetSubredditMostUpsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubredditMostUps not implemented")
}

// UnsafeRedditServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RedditServiceServer will
// result in compilation errors.
type UnsafeRedditServiceServer interface {
	mustEmbedUnimplementedRedditServiceServer()
}

func RegisterRedditServiceServer(s grpc.ServiceRegistrar, srv RedditServiceServer) {
	s.RegisterService(&RedditService_ServiceDesc, srv)
}

func _RedditService_GetSubredditMostUps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubredditMostUpsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedditServiceServer).GetSubredditMostUps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reddit.v1.RedditService/GetSubredditMostUps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedditServiceServer).GetSubredditMostUps(ctx, req.(*GetSubredditMostUpsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RedditService_ServiceDesc is the grpc.ServiceDesc for RedditService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RedditService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reddit.v1.RedditService",
	HandlerType: (*RedditServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSubredditMostUps",
			Handler:    _RedditService_GetSubredditMostUps_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/reddit.proto",
}
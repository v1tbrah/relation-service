// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.14.0
// source: relation-service.proto

package rpbapi

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
	RelationService_AddFriend_FullMethodName    = "/rpbapi.RelationService/AddFriend"
	RelationService_RemoveFriend_FullMethodName = "/rpbapi.RelationService/RemoveFriend"
	RelationService_GetFriends_FullMethodName   = "/rpbapi.RelationService/GetFriends"
)

// RelationServiceClient is the client API for RelationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RelationServiceClient interface {
	AddFriend(ctx context.Context, in *AddFriendRequest, opts ...grpc.CallOption) (*Empty, error)
	RemoveFriend(ctx context.Context, in *RemoveFriendRequest, opts ...grpc.CallOption) (*Empty, error)
	GetFriends(ctx context.Context, in *GetFriendsRequest, opts ...grpc.CallOption) (*GetFriendsResponse, error)
}

type relationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRelationServiceClient(cc grpc.ClientConnInterface) RelationServiceClient {
	return &relationServiceClient{cc}
}

func (c *relationServiceClient) AddFriend(ctx context.Context, in *AddFriendRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, RelationService_AddFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) RemoveFriend(ctx context.Context, in *RemoveFriendRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, RelationService_RemoveFriend_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationServiceClient) GetFriends(ctx context.Context, in *GetFriendsRequest, opts ...grpc.CallOption) (*GetFriendsResponse, error) {
	out := new(GetFriendsResponse)
	err := c.cc.Invoke(ctx, RelationService_GetFriends_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelationServiceServer is the server API for RelationService service.
// All implementations must embed UnimplementedRelationServiceServer
// for forward compatibility
type RelationServiceServer interface {
	AddFriend(context.Context, *AddFriendRequest) (*Empty, error)
	RemoveFriend(context.Context, *RemoveFriendRequest) (*Empty, error)
	GetFriends(context.Context, *GetFriendsRequest) (*GetFriendsResponse, error)
	mustEmbedUnimplementedRelationServiceServer()
}

// UnimplementedRelationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRelationServiceServer struct {
}

func (UnimplementedRelationServiceServer) AddFriend(context.Context, *AddFriendRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}
func (UnimplementedRelationServiceServer) RemoveFriend(context.Context, *RemoveFriendRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFriend not implemented")
}
func (UnimplementedRelationServiceServer) GetFriends(context.Context, *GetFriendsRequest) (*GetFriendsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriends not implemented")
}
func (UnimplementedRelationServiceServer) mustEmbedUnimplementedRelationServiceServer() {}

// UnsafeRelationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelationServiceServer will
// result in compilation errors.
type UnsafeRelationServiceServer interface {
	mustEmbedUnimplementedRelationServiceServer()
}

func RegisterRelationServiceServer(s grpc.ServiceRegistrar, srv RelationServiceServer) {
	s.RegisterService(&RelationService_ServiceDesc, srv)
}

func _RelationService_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_AddFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).AddFriend(ctx, req.(*AddFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_RemoveFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).RemoveFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_RemoveFriend_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).RemoveFriend(ctx, req.(*RemoveFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationService_GetFriends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFriendsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationServiceServer).GetFriends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RelationService_GetFriends_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationServiceServer).GetFriends(ctx, req.(*GetFriendsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RelationService_ServiceDesc is the grpc.ServiceDesc for RelationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RelationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpbapi.RelationService",
	HandlerType: (*RelationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFriend",
			Handler:    _RelationService_AddFriend_Handler,
		},
		{
			MethodName: "RemoveFriend",
			Handler:    _RelationService_RemoveFriend_Handler,
		},
		{
			MethodName: "GetFriends",
			Handler:    _RelationService_GetFriends_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relation-service.proto",
}

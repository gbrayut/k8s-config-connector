// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/storage/v1/service.proto

package storage

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageClient interface {
	// Returns metadata for the specified bucket.
	GetBucket(ctx context.Context, in *GetBucketRequest, opts ...grpc.CallOption) (*Bucket, error)
	// Creates a new bucket.
	InsertBucket(ctx context.Context, in *InsertBucketRequest, opts ...grpc.CallOption) (*Bucket, error)
	// Permanently deletes an empty bucket.
	DeleteBucket(ctx context.Context, in *DeleteBucketRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Updates a bucket.
	PatchBucket(ctx context.Context, in *PatchBucketRequest, opts ...grpc.CallOption) (*Bucket, error)
}

type storageClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageClient(cc grpc.ClientConnInterface) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) GetBucket(ctx context.Context, in *GetBucketRequest, opts ...grpc.CallOption) (*Bucket, error) {
	out := new(Bucket)
	err := c.cc.Invoke(ctx, "/mockgcp.storage.v1.Storage/GetBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) InsertBucket(ctx context.Context, in *InsertBucketRequest, opts ...grpc.CallOption) (*Bucket, error) {
	out := new(Bucket)
	err := c.cc.Invoke(ctx, "/mockgcp.storage.v1.Storage/InsertBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) DeleteBucket(ctx context.Context, in *DeleteBucketRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/mockgcp.storage.v1.Storage/DeleteBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) PatchBucket(ctx context.Context, in *PatchBucketRequest, opts ...grpc.CallOption) (*Bucket, error) {
	out := new(Bucket)
	err := c.cc.Invoke(ctx, "/mockgcp.storage.v1.Storage/PatchBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServer is the server API for Storage service.
// All implementations must embed UnimplementedStorageServer
// for forward compatibility
type StorageServer interface {
	// Returns metadata for the specified bucket.
	GetBucket(context.Context, *GetBucketRequest) (*Bucket, error)
	// Creates a new bucket.
	InsertBucket(context.Context, *InsertBucketRequest) (*Bucket, error)
	// Permanently deletes an empty bucket.
	DeleteBucket(context.Context, *DeleteBucketRequest) (*empty.Empty, error)
	// Updates a bucket.
	PatchBucket(context.Context, *PatchBucketRequest) (*Bucket, error)
	mustEmbedUnimplementedStorageServer()
}

// UnimplementedStorageServer must be embedded to have forward compatible implementations.
type UnimplementedStorageServer struct {
}

func (UnimplementedStorageServer) GetBucket(context.Context, *GetBucketRequest) (*Bucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBucket not implemented")
}
func (UnimplementedStorageServer) InsertBucket(context.Context, *InsertBucketRequest) (*Bucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertBucket not implemented")
}
func (UnimplementedStorageServer) DeleteBucket(context.Context, *DeleteBucketRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBucket not implemented")
}
func (UnimplementedStorageServer) PatchBucket(context.Context, *PatchBucketRequest) (*Bucket, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PatchBucket not implemented")
}
func (UnimplementedStorageServer) mustEmbedUnimplementedStorageServer() {}

// UnsafeStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageServer will
// result in compilation errors.
type UnsafeStorageServer interface {
	mustEmbedUnimplementedStorageServer()
}

func RegisterStorageServer(s grpc.ServiceRegistrar, srv StorageServer) {
	s.RegisterService(&Storage_ServiceDesc, srv)
}

func _Storage_GetBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).GetBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.storage.v1.Storage/GetBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).GetBucket(ctx, req.(*GetBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_InsertBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).InsertBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.storage.v1.Storage/InsertBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).InsertBucket(ctx, req.(*InsertBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_DeleteBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).DeleteBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.storage.v1.Storage/DeleteBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).DeleteBucket(ctx, req.(*DeleteBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_PatchBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PatchBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).PatchBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.storage.v1.Storage/PatchBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).PatchBucket(ctx, req.(*PatchBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Storage_ServiceDesc is the grpc.ServiceDesc for Storage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Storage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.storage.v1.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBucket",
			Handler:    _Storage_GetBucket_Handler,
		},
		{
			MethodName: "InsertBucket",
			Handler:    _Storage_InsertBucket_Handler,
		},
		{
			MethodName: "DeleteBucket",
			Handler:    _Storage_DeleteBucket_Handler,
		},
		{
			MethodName: "PatchBucket",
			Handler:    _Storage_PatchBucket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockgcp/storage/v1/service.proto",
}

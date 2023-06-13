// proto/serv/favourite/v2/favourite_srv.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: proto/serv/favourite/v2/favourite_srv.proto

package favourite_srvpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	v1 "x-gwi/proto/core/favourite/v1"
	v11 "x-gwi/proto/core/idx/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	FavouriteService_Create_FullMethodName = "/proto.serv.favourite.v2.FavouriteService/Create"
	FavouriteService_Get_FullMethodName    = "/proto.serv.favourite.v2.FavouriteService/Get"
	FavouriteService_Gett_FullMethodName   = "/proto.serv.favourite.v2.FavouriteService/Gett"
	FavouriteService_Update_FullMethodName = "/proto.serv.favourite.v2.FavouriteService/Update"
	FavouriteService_Delete_FullMethodName = "/proto.serv.favourite.v2.FavouriteService/Delete"
	FavouriteService_List_FullMethodName   = "/proto.serv.favourite.v2.FavouriteService/List"
)

// FavouriteServiceClient is the client API for FavouriteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FavouriteServiceClient interface {
	// Create
	Create(ctx context.Context, in *v1.FavouriteAsset, opts ...grpc.CallOption) (*v1.FavouriteAsset, error)
	// Get
	Get(ctx context.Context, in *v11.IDX, opts ...grpc.CallOption) (*v1.FavouriteAsset, error)
	// Gett
	Gett(ctx context.Context, in *v11.IDX, opts ...grpc.CallOption) (*v1.FavouriteAsset, error)
	// Update
	Update(ctx context.Context, in *v1.FavouriteAsset, opts ...grpc.CallOption) (*v1.FavouriteAsset, error)
	// Delete
	Delete(ctx context.Context, in *v11.IDX, opts ...grpc.CallOption) (*v1.FavouriteAsset, error)
	// List - stream favourites of a user
	List(ctx context.Context, in *v11.IDX, opts ...grpc.CallOption) (FavouriteService_ListClient, error)
}

type favouriteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFavouriteServiceClient(cc grpc.ClientConnInterface) FavouriteServiceClient {
	return &favouriteServiceClient{cc}
}

func (c *favouriteServiceClient) Create(ctx context.Context, in *v1.FavouriteAsset, opts ...grpc.CallOption) (*v1.FavouriteAsset, error) {
	out := new(v1.FavouriteAsset)
	err := c.cc.Invoke(ctx, FavouriteService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favouriteServiceClient) Get(ctx context.Context, in *v11.IDX, opts ...grpc.CallOption) (*v1.FavouriteAsset, error) {
	out := new(v1.FavouriteAsset)
	err := c.cc.Invoke(ctx, FavouriteService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favouriteServiceClient) Gett(ctx context.Context, in *v11.IDX, opts ...grpc.CallOption) (*v1.FavouriteAsset, error) {
	out := new(v1.FavouriteAsset)
	err := c.cc.Invoke(ctx, FavouriteService_Gett_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favouriteServiceClient) Update(ctx context.Context, in *v1.FavouriteAsset, opts ...grpc.CallOption) (*v1.FavouriteAsset, error) {
	out := new(v1.FavouriteAsset)
	err := c.cc.Invoke(ctx, FavouriteService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favouriteServiceClient) Delete(ctx context.Context, in *v11.IDX, opts ...grpc.CallOption) (*v1.FavouriteAsset, error) {
	out := new(v1.FavouriteAsset)
	err := c.cc.Invoke(ctx, FavouriteService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *favouriteServiceClient) List(ctx context.Context, in *v11.IDX, opts ...grpc.CallOption) (FavouriteService_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &FavouriteService_ServiceDesc.Streams[0], FavouriteService_List_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &favouriteServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FavouriteService_ListClient interface {
	Recv() (*v1.FavouriteAsset, error)
	grpc.ClientStream
}

type favouriteServiceListClient struct {
	grpc.ClientStream
}

func (x *favouriteServiceListClient) Recv() (*v1.FavouriteAsset, error) {
	m := new(v1.FavouriteAsset)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FavouriteServiceServer is the server API for FavouriteService service.
// All implementations must embed UnimplementedFavouriteServiceServer
// for forward compatibility
type FavouriteServiceServer interface {
	// Create
	Create(context.Context, *v1.FavouriteAsset) (*v1.FavouriteAsset, error)
	// Get
	Get(context.Context, *v11.IDX) (*v1.FavouriteAsset, error)
	// Gett
	Gett(context.Context, *v11.IDX) (*v1.FavouriteAsset, error)
	// Update
	Update(context.Context, *v1.FavouriteAsset) (*v1.FavouriteAsset, error)
	// Delete
	Delete(context.Context, *v11.IDX) (*v1.FavouriteAsset, error)
	// List - stream favourites of a user
	List(*v11.IDX, FavouriteService_ListServer) error
	mustEmbedUnimplementedFavouriteServiceServer()
}

// UnimplementedFavouriteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFavouriteServiceServer struct {
}

func (UnimplementedFavouriteServiceServer) Create(context.Context, *v1.FavouriteAsset) (*v1.FavouriteAsset, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedFavouriteServiceServer) Get(context.Context, *v11.IDX) (*v1.FavouriteAsset, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedFavouriteServiceServer) Gett(context.Context, *v11.IDX) (*v1.FavouriteAsset, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Gett not implemented")
}
func (UnimplementedFavouriteServiceServer) Update(context.Context, *v1.FavouriteAsset) (*v1.FavouriteAsset, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedFavouriteServiceServer) Delete(context.Context, *v11.IDX) (*v1.FavouriteAsset, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedFavouriteServiceServer) List(*v11.IDX, FavouriteService_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedFavouriteServiceServer) mustEmbedUnimplementedFavouriteServiceServer() {}

// UnsafeFavouriteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FavouriteServiceServer will
// result in compilation errors.
type UnsafeFavouriteServiceServer interface {
	mustEmbedUnimplementedFavouriteServiceServer()
}

func RegisterFavouriteServiceServer(s grpc.ServiceRegistrar, srv FavouriteServiceServer) {
	s.RegisterService(&FavouriteService_ServiceDesc, srv)
}

func _FavouriteService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.FavouriteAsset)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavouriteServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FavouriteService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavouriteServiceServer).Create(ctx, req.(*v1.FavouriteAsset))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavouriteService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.IDX)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavouriteServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FavouriteService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavouriteServiceServer).Get(ctx, req.(*v11.IDX))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavouriteService_Gett_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.IDX)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavouriteServiceServer).Gett(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FavouriteService_Gett_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavouriteServiceServer).Gett(ctx, req.(*v11.IDX))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavouriteService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.FavouriteAsset)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavouriteServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FavouriteService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavouriteServiceServer).Update(ctx, req.(*v1.FavouriteAsset))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavouriteService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v11.IDX)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FavouriteServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FavouriteService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FavouriteServiceServer).Delete(ctx, req.(*v11.IDX))
	}
	return interceptor(ctx, in, info, handler)
}

func _FavouriteService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(v11.IDX)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FavouriteServiceServer).List(m, &favouriteServiceListServer{stream})
}

type FavouriteService_ListServer interface {
	Send(*v1.FavouriteAsset) error
	grpc.ServerStream
}

type favouriteServiceListServer struct {
	grpc.ServerStream
}

func (x *favouriteServiceListServer) Send(m *v1.FavouriteAsset) error {
	return x.ServerStream.SendMsg(m)
}

// FavouriteService_ServiceDesc is the grpc.ServiceDesc for FavouriteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FavouriteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.serv.favourite.v2.FavouriteService",
	HandlerType: (*FavouriteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _FavouriteService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _FavouriteService_Get_Handler,
		},
		{
			MethodName: "Gett",
			Handler:    _FavouriteService_Gett_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _FavouriteService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FavouriteService_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _FavouriteService_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/serv/favourite/v2/favourite_srv.proto",
}
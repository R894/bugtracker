// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0--rc3
// source: ports/bug.proto

package ports

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
	BugRepositoryService_SaveBug_FullMethodName            = "/bug.BugRepositoryService/SaveBug"
	BugRepositoryService_GetBugByID_FullMethodName         = "/bug.BugRepositoryService/GetBugByID"
	BugRepositoryService_GetBugsByProjectID_FullMethodName = "/bug.BugRepositoryService/GetBugsByProjectID"
	BugRepositoryService_GetBugs_FullMethodName            = "/bug.BugRepositoryService/GetBugs"
	BugRepositoryService_AssignBugTo_FullMethodName        = "/bug.BugRepositoryService/AssignBugTo"
	BugRepositoryService_UpdateBug_FullMethodName          = "/bug.BugRepositoryService/UpdateBug"
	BugRepositoryService_DeleteBug_FullMethodName          = "/bug.BugRepositoryService/DeleteBug"
)

// BugRepositoryServiceClient is the client API for BugRepositoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BugRepositoryServiceClient interface {
	SaveBug(ctx context.Context, in *SaveBugRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	GetBugByID(ctx context.Context, in *GetBugByIDRequest, opts ...grpc.CallOption) (*BugResponse, error)
	GetBugsByProjectID(ctx context.Context, in *GetBugsByProjectIDRequest, opts ...grpc.CallOption) (*GetBugsByProjectIDResponse, error)
	GetBugs(ctx context.Context, in *GetBugsRequest, opts ...grpc.CallOption) (*GetBugsResponse, error)
	AssignBugTo(ctx context.Context, in *AssignBugToRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	UpdateBug(ctx context.Context, in *UpdateBugRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
	DeleteBug(ctx context.Context, in *DeleteBugRequest, opts ...grpc.CallOption) (*EmptyResponse, error)
}

type bugRepositoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBugRepositoryServiceClient(cc grpc.ClientConnInterface) BugRepositoryServiceClient {
	return &bugRepositoryServiceClient{cc}
}

func (c *bugRepositoryServiceClient) SaveBug(ctx context.Context, in *SaveBugRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, BugRepositoryService_SaveBug_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bugRepositoryServiceClient) GetBugByID(ctx context.Context, in *GetBugByIDRequest, opts ...grpc.CallOption) (*BugResponse, error) {
	out := new(BugResponse)
	err := c.cc.Invoke(ctx, BugRepositoryService_GetBugByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bugRepositoryServiceClient) GetBugsByProjectID(ctx context.Context, in *GetBugsByProjectIDRequest, opts ...grpc.CallOption) (*GetBugsByProjectIDResponse, error) {
	out := new(GetBugsByProjectIDResponse)
	err := c.cc.Invoke(ctx, BugRepositoryService_GetBugsByProjectID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bugRepositoryServiceClient) GetBugs(ctx context.Context, in *GetBugsRequest, opts ...grpc.CallOption) (*GetBugsResponse, error) {
	out := new(GetBugsResponse)
	err := c.cc.Invoke(ctx, BugRepositoryService_GetBugs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bugRepositoryServiceClient) AssignBugTo(ctx context.Context, in *AssignBugToRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, BugRepositoryService_AssignBugTo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bugRepositoryServiceClient) UpdateBug(ctx context.Context, in *UpdateBugRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, BugRepositoryService_UpdateBug_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bugRepositoryServiceClient) DeleteBug(ctx context.Context, in *DeleteBugRequest, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, BugRepositoryService_DeleteBug_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BugRepositoryServiceServer is the server API for BugRepositoryService service.
// All implementations must embed UnimplementedBugRepositoryServiceServer
// for forward compatibility
type BugRepositoryServiceServer interface {
	SaveBug(context.Context, *SaveBugRequest) (*EmptyResponse, error)
	GetBugByID(context.Context, *GetBugByIDRequest) (*BugResponse, error)
	GetBugsByProjectID(context.Context, *GetBugsByProjectIDRequest) (*GetBugsByProjectIDResponse, error)
	GetBugs(context.Context, *GetBugsRequest) (*GetBugsResponse, error)
	AssignBugTo(context.Context, *AssignBugToRequest) (*EmptyResponse, error)
	UpdateBug(context.Context, *UpdateBugRequest) (*EmptyResponse, error)
	DeleteBug(context.Context, *DeleteBugRequest) (*EmptyResponse, error)
	mustEmbedUnimplementedBugRepositoryServiceServer()
}

// UnimplementedBugRepositoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBugRepositoryServiceServer struct {
}

func (UnimplementedBugRepositoryServiceServer) SaveBug(context.Context, *SaveBugRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveBug not implemented")
}
func (UnimplementedBugRepositoryServiceServer) GetBugByID(context.Context, *GetBugByIDRequest) (*BugResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBugByID not implemented")
}
func (UnimplementedBugRepositoryServiceServer) GetBugsByProjectID(context.Context, *GetBugsByProjectIDRequest) (*GetBugsByProjectIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBugsByProjectID not implemented")
}
func (UnimplementedBugRepositoryServiceServer) GetBugs(context.Context, *GetBugsRequest) (*GetBugsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBugs not implemented")
}
func (UnimplementedBugRepositoryServiceServer) AssignBugTo(context.Context, *AssignBugToRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignBugTo not implemented")
}
func (UnimplementedBugRepositoryServiceServer) UpdateBug(context.Context, *UpdateBugRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBug not implemented")
}
func (UnimplementedBugRepositoryServiceServer) DeleteBug(context.Context, *DeleteBugRequest) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBug not implemented")
}
func (UnimplementedBugRepositoryServiceServer) mustEmbedUnimplementedBugRepositoryServiceServer() {}

// UnsafeBugRepositoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BugRepositoryServiceServer will
// result in compilation errors.
type UnsafeBugRepositoryServiceServer interface {
	mustEmbedUnimplementedBugRepositoryServiceServer()
}

func RegisterBugRepositoryServiceServer(s grpc.ServiceRegistrar, srv BugRepositoryServiceServer) {
	s.RegisterService(&BugRepositoryService_ServiceDesc, srv)
}

func _BugRepositoryService_SaveBug_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveBugRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BugRepositoryServiceServer).SaveBug(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BugRepositoryService_SaveBug_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BugRepositoryServiceServer).SaveBug(ctx, req.(*SaveBugRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BugRepositoryService_GetBugByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBugByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BugRepositoryServiceServer).GetBugByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BugRepositoryService_GetBugByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BugRepositoryServiceServer).GetBugByID(ctx, req.(*GetBugByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BugRepositoryService_GetBugsByProjectID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBugsByProjectIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BugRepositoryServiceServer).GetBugsByProjectID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BugRepositoryService_GetBugsByProjectID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BugRepositoryServiceServer).GetBugsByProjectID(ctx, req.(*GetBugsByProjectIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BugRepositoryService_GetBugs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBugsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BugRepositoryServiceServer).GetBugs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BugRepositoryService_GetBugs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BugRepositoryServiceServer).GetBugs(ctx, req.(*GetBugsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BugRepositoryService_AssignBugTo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignBugToRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BugRepositoryServiceServer).AssignBugTo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BugRepositoryService_AssignBugTo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BugRepositoryServiceServer).AssignBugTo(ctx, req.(*AssignBugToRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BugRepositoryService_UpdateBug_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBugRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BugRepositoryServiceServer).UpdateBug(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BugRepositoryService_UpdateBug_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BugRepositoryServiceServer).UpdateBug(ctx, req.(*UpdateBugRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BugRepositoryService_DeleteBug_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBugRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BugRepositoryServiceServer).DeleteBug(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BugRepositoryService_DeleteBug_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BugRepositoryServiceServer).DeleteBug(ctx, req.(*DeleteBugRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BugRepositoryService_ServiceDesc is the grpc.ServiceDesc for BugRepositoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BugRepositoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bug.BugRepositoryService",
	HandlerType: (*BugRepositoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveBug",
			Handler:    _BugRepositoryService_SaveBug_Handler,
		},
		{
			MethodName: "GetBugByID",
			Handler:    _BugRepositoryService_GetBugByID_Handler,
		},
		{
			MethodName: "GetBugsByProjectID",
			Handler:    _BugRepositoryService_GetBugsByProjectID_Handler,
		},
		{
			MethodName: "GetBugs",
			Handler:    _BugRepositoryService_GetBugs_Handler,
		},
		{
			MethodName: "AssignBugTo",
			Handler:    _BugRepositoryService_AssignBugTo_Handler,
		},
		{
			MethodName: "UpdateBug",
			Handler:    _BugRepositoryService_UpdateBug_Handler,
		},
		{
			MethodName: "DeleteBug",
			Handler:    _BugRepositoryService_DeleteBug_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ports/bug.proto",
}

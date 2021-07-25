// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// BaseappInterfaceClient is the client API for BaseappInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BaseappInterfaceClient interface {
	GetCommentSubject(ctx context.Context, in *GetCommentSubjectRequest, opts ...grpc.CallOption) (*GetCommentSubjectReply, error)
	SaveComment(ctx context.Context, in *SaveCommentRequest, opts ...grpc.CallOption) (*SaveCommentReply, error)
	GetCommentList(ctx context.Context, in *GetCommentListRequest, opts ...grpc.CallOption) (*GetCommentListReply, error)
	GetReplyList(ctx context.Context, in *GetReplyListRequest, opts ...grpc.CallOption) (*GetCommentListReply, error)
	GetComment(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*GetCommentReply, error)
	LikeComment(ctx context.Context, in *LikeCommentRequest, opts ...grpc.CallOption) (*LikeCommentReply, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
}

type baseappInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewBaseappInterfaceClient(cc grpc.ClientConnInterface) BaseappInterfaceClient {
	return &baseappInterfaceClient{cc}
}

func (c *baseappInterfaceClient) GetCommentSubject(ctx context.Context, in *GetCommentSubjectRequest, opts ...grpc.CallOption) (*GetCommentSubjectReply, error) {
	out := new(GetCommentSubjectReply)
	err := c.cc.Invoke(ctx, "/api.baseapp.interface.v1.BaseappInterface/GetCommentSubject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseappInterfaceClient) SaveComment(ctx context.Context, in *SaveCommentRequest, opts ...grpc.CallOption) (*SaveCommentReply, error) {
	out := new(SaveCommentReply)
	err := c.cc.Invoke(ctx, "/api.baseapp.interface.v1.BaseappInterface/SaveComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseappInterfaceClient) GetCommentList(ctx context.Context, in *GetCommentListRequest, opts ...grpc.CallOption) (*GetCommentListReply, error) {
	out := new(GetCommentListReply)
	err := c.cc.Invoke(ctx, "/api.baseapp.interface.v1.BaseappInterface/GetCommentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseappInterfaceClient) GetReplyList(ctx context.Context, in *GetReplyListRequest, opts ...grpc.CallOption) (*GetCommentListReply, error) {
	out := new(GetCommentListReply)
	err := c.cc.Invoke(ctx, "/api.baseapp.interface.v1.BaseappInterface/GetReplyList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseappInterfaceClient) GetComment(ctx context.Context, in *GetCommentRequest, opts ...grpc.CallOption) (*GetCommentReply, error) {
	out := new(GetCommentReply)
	err := c.cc.Invoke(ctx, "/api.baseapp.interface.v1.BaseappInterface/GetComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseappInterfaceClient) LikeComment(ctx context.Context, in *LikeCommentRequest, opts ...grpc.CallOption) (*LikeCommentReply, error) {
	out := new(LikeCommentReply)
	err := c.cc.Invoke(ctx, "/api.baseapp.interface.v1.BaseappInterface/LikeComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseappInterfaceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/api.baseapp.interface.v1.BaseappInterface/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BaseappInterfaceServer is the server API for BaseappInterface service.
// All implementations must embed UnimplementedBaseappInterfaceServer
// for forward compatibility
type BaseappInterfaceServer interface {
	GetCommentSubject(context.Context, *GetCommentSubjectRequest) (*GetCommentSubjectReply, error)
	SaveComment(context.Context, *SaveCommentRequest) (*SaveCommentReply, error)
	GetCommentList(context.Context, *GetCommentListRequest) (*GetCommentListReply, error)
	GetReplyList(context.Context, *GetReplyListRequest) (*GetCommentListReply, error)
	GetComment(context.Context, *GetCommentRequest) (*GetCommentReply, error)
	LikeComment(context.Context, *LikeCommentRequest) (*LikeCommentReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	mustEmbedUnimplementedBaseappInterfaceServer()
}

// UnimplementedBaseappInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedBaseappInterfaceServer struct {
}

func (UnimplementedBaseappInterfaceServer) GetCommentSubject(context.Context, *GetCommentSubjectRequest) (*GetCommentSubjectReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentSubject not implemented")
}
func (UnimplementedBaseappInterfaceServer) SaveComment(context.Context, *SaveCommentRequest) (*SaveCommentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveComment not implemented")
}
func (UnimplementedBaseappInterfaceServer) GetCommentList(context.Context, *GetCommentListRequest) (*GetCommentListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentList not implemented")
}
func (UnimplementedBaseappInterfaceServer) GetReplyList(context.Context, *GetReplyListRequest) (*GetCommentListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReplyList not implemented")
}
func (UnimplementedBaseappInterfaceServer) GetComment(context.Context, *GetCommentRequest) (*GetCommentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComment not implemented")
}
func (UnimplementedBaseappInterfaceServer) LikeComment(context.Context, *LikeCommentRequest) (*LikeCommentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikeComment not implemented")
}
func (UnimplementedBaseappInterfaceServer) Login(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedBaseappInterfaceServer) mustEmbedUnimplementedBaseappInterfaceServer() {}

// UnsafeBaseappInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BaseappInterfaceServer will
// result in compilation errors.
type UnsafeBaseappInterfaceServer interface {
	mustEmbedUnimplementedBaseappInterfaceServer()
}

func RegisterBaseappInterfaceServer(s grpc.ServiceRegistrar, srv BaseappInterfaceServer) {
	s.RegisterService(&BaseappInterface_ServiceDesc, srv)
}

func _BaseappInterface_GetCommentSubject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentSubjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseappInterfaceServer).GetCommentSubject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.baseapp.interface.v1.BaseappInterface/GetCommentSubject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseappInterfaceServer).GetCommentSubject(ctx, req.(*GetCommentSubjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaseappInterface_SaveComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseappInterfaceServer).SaveComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.baseapp.interface.v1.BaseappInterface/SaveComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseappInterfaceServer).SaveComment(ctx, req.(*SaveCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaseappInterface_GetCommentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseappInterfaceServer).GetCommentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.baseapp.interface.v1.BaseappInterface/GetCommentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseappInterfaceServer).GetCommentList(ctx, req.(*GetCommentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaseappInterface_GetReplyList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReplyListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseappInterfaceServer).GetReplyList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.baseapp.interface.v1.BaseappInterface/GetReplyList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseappInterfaceServer).GetReplyList(ctx, req.(*GetReplyListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaseappInterface_GetComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseappInterfaceServer).GetComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.baseapp.interface.v1.BaseappInterface/GetComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseappInterfaceServer).GetComment(ctx, req.(*GetCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaseappInterface_LikeComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseappInterfaceServer).LikeComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.baseapp.interface.v1.BaseappInterface/LikeComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseappInterfaceServer).LikeComment(ctx, req.(*LikeCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaseappInterface_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseappInterfaceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.baseapp.interface.v1.BaseappInterface/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseappInterfaceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BaseappInterface_ServiceDesc is the grpc.ServiceDesc for BaseappInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BaseappInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.baseapp.interface.v1.BaseappInterface",
	HandlerType: (*BaseappInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCommentSubject",
			Handler:    _BaseappInterface_GetCommentSubject_Handler,
		},
		{
			MethodName: "SaveComment",
			Handler:    _BaseappInterface_SaveComment_Handler,
		},
		{
			MethodName: "GetCommentList",
			Handler:    _BaseappInterface_GetCommentList_Handler,
		},
		{
			MethodName: "GetReplyList",
			Handler:    _BaseappInterface_GetReplyList_Handler,
		},
		{
			MethodName: "GetComment",
			Handler:    _BaseappInterface_GetComment_Handler,
		},
		{
			MethodName: "LikeComment",
			Handler:    _BaseappInterface_LikeComment_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _BaseappInterface_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/baseapp/interface/v1/baseapp_interface.proto",
}
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package todogrpc

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

// TodoClient is the client API for Todo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	RetrieveView(ctx context.Context, in *RetrieveViewRequest, opts ...grpc.CallOption) (*RetrieveViewResponse, error)
}

type todoClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoClient(cc grpc.ClientConnInterface) TodoClient {
	return &todoClient{cc}
}

func (c *todoClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/todogrpc.Todo/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) RetrieveView(ctx context.Context, in *RetrieveViewRequest, opts ...grpc.CallOption) (*RetrieveViewResponse, error) {
	out := new(RetrieveViewResponse)
	err := c.cc.Invoke(ctx, "/todogrpc.Todo/RetrieveView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServer is the server API for Todo service.
// All implementations must embed UnimplementedTodoServer
// for forward compatibility
type TodoServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
	RetrieveView(context.Context, *RetrieveViewRequest) (*RetrieveViewResponse, error)
	mustEmbedUnimplementedTodoServer()
}

// UnimplementedTodoServer must be embedded to have forward compatible implementations.
type UnimplementedTodoServer struct {
}

func (UnimplementedTodoServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedTodoServer) RetrieveView(context.Context, *RetrieveViewRequest) (*RetrieveViewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveView not implemented")
}
func (UnimplementedTodoServer) mustEmbedUnimplementedTodoServer() {}

// UnsafeTodoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoServer will
// result in compilation errors.
type UnsafeTodoServer interface {
	mustEmbedUnimplementedTodoServer()
}

func RegisterTodoServer(s grpc.ServiceRegistrar, srv TodoServer) {
	s.RegisterService(&Todo_ServiceDesc, srv)
}

func _Todo_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todogrpc.Todo/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_RetrieveView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).RetrieveView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todogrpc.Todo/RetrieveView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).RetrieveView(ctx, req.(*RetrieveViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Todo_ServiceDesc is the grpc.ServiceDesc for Todo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Todo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todogrpc.Todo",
	HandlerType: (*TodoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Todo_Add_Handler,
		},
		{
			MethodName: "RetrieveView",
			Handler:    _Todo_RetrieveView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/todo.proto",
}

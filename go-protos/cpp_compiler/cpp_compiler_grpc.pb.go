// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: protos/cpp_compiler/cpp_compiler.proto

package cpp_compiler

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

// CPPCompilerClient is the client API for CPPCompiler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CPPCompilerClient interface {
	PingPong(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	CompileAndRun(ctx context.Context, in *CompileAndRunRequest, opts ...grpc.CallOption) (*CompileAndRunResponse, error)
}

type cPPCompilerClient struct {
	cc grpc.ClientConnInterface
}

func NewCPPCompilerClient(cc grpc.ClientConnInterface) CPPCompilerClient {
	return &cPPCompilerClient{cc}
}

func (c *cPPCompilerClient) PingPong(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/protobufs.proto.cppcompiler.CPPCompiler/PingPong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cPPCompilerClient) CompileAndRun(ctx context.Context, in *CompileAndRunRequest, opts ...grpc.CallOption) (*CompileAndRunResponse, error) {
	out := new(CompileAndRunResponse)
	err := c.cc.Invoke(ctx, "/protobufs.proto.cppcompiler.CPPCompiler/CompileAndRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CPPCompilerServer is the server API for CPPCompiler service.
// All implementations must embed UnimplementedCPPCompilerServer
// for forward compatibility
type CPPCompilerServer interface {
	PingPong(context.Context, *PingRequest) (*PingResponse, error)
	CompileAndRun(context.Context, *CompileAndRunRequest) (*CompileAndRunResponse, error)
	mustEmbedUnimplementedCPPCompilerServer()
}

// UnimplementedCPPCompilerServer must be embedded to have forward compatible implementations.
type UnimplementedCPPCompilerServer struct {
}

func (UnimplementedCPPCompilerServer) PingPong(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingPong not implemented")
}
func (UnimplementedCPPCompilerServer) CompileAndRun(context.Context, *CompileAndRunRequest) (*CompileAndRunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompileAndRun not implemented")
}
func (UnimplementedCPPCompilerServer) mustEmbedUnimplementedCPPCompilerServer() {}

// UnsafeCPPCompilerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CPPCompilerServer will
// result in compilation errors.
type UnsafeCPPCompilerServer interface {
	mustEmbedUnimplementedCPPCompilerServer()
}

func RegisterCPPCompilerServer(s grpc.ServiceRegistrar, srv CPPCompilerServer) {
	s.RegisterService(&CPPCompiler_ServiceDesc, srv)
}

func _CPPCompiler_PingPong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPPCompilerServer).PingPong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobufs.proto.cppcompiler.CPPCompiler/PingPong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPPCompilerServer).PingPong(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CPPCompiler_CompileAndRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompileAndRunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CPPCompilerServer).CompileAndRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobufs.proto.cppcompiler.CPPCompiler/CompileAndRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CPPCompilerServer).CompileAndRun(ctx, req.(*CompileAndRunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CPPCompiler_ServiceDesc is the grpc.ServiceDesc for CPPCompiler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CPPCompiler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobufs.proto.cppcompiler.CPPCompiler",
	HandlerType: (*CPPCompilerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PingPong",
			Handler:    _CPPCompiler_PingPong_Handler,
		},
		{
			MethodName: "CompileAndRun",
			Handler:    _CPPCompiler_CompileAndRun_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/cpp_compiler/cpp_compiler.proto",
}

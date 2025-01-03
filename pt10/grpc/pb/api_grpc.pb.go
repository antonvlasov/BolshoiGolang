// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: api.proto

package pb

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
	Orchestrator_Run_FullMethodName    = "/streaming.Orchestrator/Run"
	Orchestrator_Get_FullMethodName    = "/streaming.Orchestrator/Get"
	Orchestrator_Cancel_FullMethodName = "/streaming.Orchestrator/Cancel"
)

// OrchestratorClient is the client API for Orchestrator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrchestratorClient interface {
	Run(ctx context.Context, in *RunQuery, opts ...grpc.CallOption) (*RunResp, error)
	Get(ctx context.Context, in *GetQuery, opts ...grpc.CallOption) (*GetResp, error)
	Cancel(ctx context.Context, in *CancelQuery, opts ...grpc.CallOption) (*CancelResp, error)
}

type orchestratorClient struct {
	cc grpc.ClientConnInterface
}

func NewOrchestratorClient(cc grpc.ClientConnInterface) OrchestratorClient {
	return &orchestratorClient{cc}
}

func (c *orchestratorClient) Run(ctx context.Context, in *RunQuery, opts ...grpc.CallOption) (*RunResp, error) {
	out := new(RunResp)
	err := c.cc.Invoke(ctx, Orchestrator_Run_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orchestratorClient) Get(ctx context.Context, in *GetQuery, opts ...grpc.CallOption) (*GetResp, error) {
	out := new(GetResp)
	err := c.cc.Invoke(ctx, Orchestrator_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orchestratorClient) Cancel(ctx context.Context, in *CancelQuery, opts ...grpc.CallOption) (*CancelResp, error) {
	out := new(CancelResp)
	err := c.cc.Invoke(ctx, Orchestrator_Cancel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrchestratorServer is the server API for Orchestrator service.
// All implementations must embed UnimplementedOrchestratorServer
// for forward compatibility
type OrchestratorServer interface {
	Run(context.Context, *RunQuery) (*RunResp, error)
	Get(context.Context, *GetQuery) (*GetResp, error)
	Cancel(context.Context, *CancelQuery) (*CancelResp, error)
	mustEmbedUnimplementedOrchestratorServer()
}

// UnimplementedOrchestratorServer must be embedded to have forward compatible implementations.
type UnimplementedOrchestratorServer struct {
}

func (UnimplementedOrchestratorServer) Run(context.Context, *RunQuery) (*RunResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (UnimplementedOrchestratorServer) Get(context.Context, *GetQuery) (*GetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedOrchestratorServer) Cancel(context.Context, *CancelQuery) (*CancelResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (UnimplementedOrchestratorServer) mustEmbedUnimplementedOrchestratorServer() {}

// UnsafeOrchestratorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrchestratorServer will
// result in compilation errors.
type UnsafeOrchestratorServer interface {
	mustEmbedUnimplementedOrchestratorServer()
}

func RegisterOrchestratorServer(s grpc.ServiceRegistrar, srv OrchestratorServer) {
	s.RegisterService(&Orchestrator_ServiceDesc, srv)
}

func _Orchestrator_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Orchestrator_Run_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).Run(ctx, req.(*RunQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orchestrator_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Orchestrator_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).Get(ctx, req.(*GetQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Orchestrator_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrchestratorServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Orchestrator_Cancel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrchestratorServer).Cancel(ctx, req.(*CancelQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// Orchestrator_ServiceDesc is the grpc.ServiceDesc for Orchestrator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Orchestrator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "streaming.Orchestrator",
	HandlerType: (*OrchestratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Run",
			Handler:    _Orchestrator_Run_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Orchestrator_Get_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _Orchestrator_Cancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

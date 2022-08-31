// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: deploys/v1/deploys.proto

package deploysv1

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

// DeploysAPIClient is the client API for DeploysAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeploysAPIClient interface {
	TriggerManualDeploy(ctx context.Context, in *TriggerManualDeployRequest, opts ...grpc.CallOption) (*TriggerManualDeployResponse, error)
}

type deploysAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewDeploysAPIClient(cc grpc.ClientConnInterface) DeploysAPIClient {
	return &deploysAPIClient{cc}
}

func (c *deploysAPIClient) TriggerManualDeploy(ctx context.Context, in *TriggerManualDeployRequest, opts ...grpc.CallOption) (*TriggerManualDeployResponse, error) {
	out := new(TriggerManualDeployResponse)
	err := c.cc.Invoke(ctx, "/clutch.deploys.v1.DeploysAPI/TriggerManualDeploy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeploysAPIServer is the server API for DeploysAPI service.
// All implementations should embed UnimplementedDeploysAPIServer
// for forward compatibility
type DeploysAPIServer interface {
	TriggerManualDeploy(context.Context, *TriggerManualDeployRequest) (*TriggerManualDeployResponse, error)
}

// UnimplementedDeploysAPIServer should be embedded to have forward compatible implementations.
type UnimplementedDeploysAPIServer struct {
}

func (UnimplementedDeploysAPIServer) TriggerManualDeploy(context.Context, *TriggerManualDeployRequest) (*TriggerManualDeployResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggerManualDeploy not implemented")
}

// UnsafeDeploysAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeploysAPIServer will
// result in compilation errors.
type UnsafeDeploysAPIServer interface {
	mustEmbedUnimplementedDeploysAPIServer()
}

func RegisterDeploysAPIServer(s grpc.ServiceRegistrar, srv DeploysAPIServer) {
	s.RegisterService(&DeploysAPI_ServiceDesc, srv)
}

func _DeploysAPI_TriggerManualDeploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerManualDeployRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeploysAPIServer).TriggerManualDeploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.deploys.v1.DeploysAPI/TriggerManualDeploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeploysAPIServer).TriggerManualDeploy(ctx, req.(*TriggerManualDeployRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeploysAPI_ServiceDesc is the grpc.ServiceDesc for DeploysAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeploysAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "clutch.deploys.v1.DeploysAPI",
	HandlerType: (*DeploysAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TriggerManualDeploy",
			Handler:    _DeploysAPI_TriggerManualDeploy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "deploys/v1/deploys.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: workflows.proto

package contracts

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

// WorkflowServiceClient is the client API for WorkflowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkflowServiceClient interface {
	ListWorkflows(ctx context.Context, in *ListWorkflowsRequest, opts ...grpc.CallOption) (*ListWorkflowsResponse, error)
	PutWorkflow(ctx context.Context, in *PutWorkflowRequest, opts ...grpc.CallOption) (*WorkflowVersion, error)
	ScheduleWorkflow(ctx context.Context, in *ScheduleWorkflowRequest, opts ...grpc.CallOption) (*WorkflowVersion, error)
	TriggerWorkflow(ctx context.Context, in *TriggerWorkflowRequest, opts ...grpc.CallOption) (*TriggerWorkflowResponse, error)
	GetWorkflowByName(ctx context.Context, in *GetWorkflowByNameRequest, opts ...grpc.CallOption) (*Workflow, error)
	ListWorkflowsForEvent(ctx context.Context, in *ListWorkflowsForEventRequest, opts ...grpc.CallOption) (*ListWorkflowsResponse, error)
	DeleteWorkflow(ctx context.Context, in *DeleteWorkflowRequest, opts ...grpc.CallOption) (*Workflow, error)
}

type workflowServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkflowServiceClient(cc grpc.ClientConnInterface) WorkflowServiceClient {
	return &workflowServiceClient{cc}
}

func (c *workflowServiceClient) ListWorkflows(ctx context.Context, in *ListWorkflowsRequest, opts ...grpc.CallOption) (*ListWorkflowsResponse, error) {
	out := new(ListWorkflowsResponse)
	err := c.cc.Invoke(ctx, "/WorkflowService/ListWorkflows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowServiceClient) PutWorkflow(ctx context.Context, in *PutWorkflowRequest, opts ...grpc.CallOption) (*WorkflowVersion, error) {
	out := new(WorkflowVersion)
	err := c.cc.Invoke(ctx, "/WorkflowService/PutWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowServiceClient) ScheduleWorkflow(ctx context.Context, in *ScheduleWorkflowRequest, opts ...grpc.CallOption) (*WorkflowVersion, error) {
	out := new(WorkflowVersion)
	err := c.cc.Invoke(ctx, "/WorkflowService/ScheduleWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowServiceClient) TriggerWorkflow(ctx context.Context, in *TriggerWorkflowRequest, opts ...grpc.CallOption) (*TriggerWorkflowResponse, error) {
	out := new(TriggerWorkflowResponse)
	err := c.cc.Invoke(ctx, "/WorkflowService/TriggerWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowServiceClient) GetWorkflowByName(ctx context.Context, in *GetWorkflowByNameRequest, opts ...grpc.CallOption) (*Workflow, error) {
	out := new(Workflow)
	err := c.cc.Invoke(ctx, "/WorkflowService/GetWorkflowByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowServiceClient) ListWorkflowsForEvent(ctx context.Context, in *ListWorkflowsForEventRequest, opts ...grpc.CallOption) (*ListWorkflowsResponse, error) {
	out := new(ListWorkflowsResponse)
	err := c.cc.Invoke(ctx, "/WorkflowService/ListWorkflowsForEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workflowServiceClient) DeleteWorkflow(ctx context.Context, in *DeleteWorkflowRequest, opts ...grpc.CallOption) (*Workflow, error) {
	out := new(Workflow)
	err := c.cc.Invoke(ctx, "/WorkflowService/DeleteWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkflowServiceServer is the server API for WorkflowService service.
// All implementations must embed UnimplementedWorkflowServiceServer
// for forward compatibility
type WorkflowServiceServer interface {
	ListWorkflows(context.Context, *ListWorkflowsRequest) (*ListWorkflowsResponse, error)
	PutWorkflow(context.Context, *PutWorkflowRequest) (*WorkflowVersion, error)
	ScheduleWorkflow(context.Context, *ScheduleWorkflowRequest) (*WorkflowVersion, error)
	TriggerWorkflow(context.Context, *TriggerWorkflowRequest) (*TriggerWorkflowResponse, error)
	GetWorkflowByName(context.Context, *GetWorkflowByNameRequest) (*Workflow, error)
	ListWorkflowsForEvent(context.Context, *ListWorkflowsForEventRequest) (*ListWorkflowsResponse, error)
	DeleteWorkflow(context.Context, *DeleteWorkflowRequest) (*Workflow, error)
	mustEmbedUnimplementedWorkflowServiceServer()
}

// UnimplementedWorkflowServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWorkflowServiceServer struct {
}

func (UnimplementedWorkflowServiceServer) ListWorkflows(context.Context, *ListWorkflowsRequest) (*ListWorkflowsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkflows not implemented")
}
func (UnimplementedWorkflowServiceServer) PutWorkflow(context.Context, *PutWorkflowRequest) (*WorkflowVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutWorkflow not implemented")
}
func (UnimplementedWorkflowServiceServer) ScheduleWorkflow(context.Context, *ScheduleWorkflowRequest) (*WorkflowVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScheduleWorkflow not implemented")
}
func (UnimplementedWorkflowServiceServer) TriggerWorkflow(context.Context, *TriggerWorkflowRequest) (*TriggerWorkflowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggerWorkflow not implemented")
}
func (UnimplementedWorkflowServiceServer) GetWorkflowByName(context.Context, *GetWorkflowByNameRequest) (*Workflow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkflowByName not implemented")
}
func (UnimplementedWorkflowServiceServer) ListWorkflowsForEvent(context.Context, *ListWorkflowsForEventRequest) (*ListWorkflowsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkflowsForEvent not implemented")
}
func (UnimplementedWorkflowServiceServer) DeleteWorkflow(context.Context, *DeleteWorkflowRequest) (*Workflow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWorkflow not implemented")
}
func (UnimplementedWorkflowServiceServer) mustEmbedUnimplementedWorkflowServiceServer() {}

// UnsafeWorkflowServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkflowServiceServer will
// result in compilation errors.
type UnsafeWorkflowServiceServer interface {
	mustEmbedUnimplementedWorkflowServiceServer()
}

func RegisterWorkflowServiceServer(s grpc.ServiceRegistrar, srv WorkflowServiceServer) {
	s.RegisterService(&WorkflowService_ServiceDesc, srv)
}

func _WorkflowService_ListWorkflows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkflowsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServiceServer).ListWorkflows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WorkflowService/ListWorkflows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServiceServer).ListWorkflows(ctx, req.(*ListWorkflowsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowService_PutWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServiceServer).PutWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WorkflowService/PutWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServiceServer).PutWorkflow(ctx, req.(*PutWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowService_ScheduleWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServiceServer).ScheduleWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WorkflowService/ScheduleWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServiceServer).ScheduleWorkflow(ctx, req.(*ScheduleWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowService_TriggerWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServiceServer).TriggerWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WorkflowService/TriggerWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServiceServer).TriggerWorkflow(ctx, req.(*TriggerWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowService_GetWorkflowByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkflowByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServiceServer).GetWorkflowByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WorkflowService/GetWorkflowByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServiceServer).GetWorkflowByName(ctx, req.(*GetWorkflowByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowService_ListWorkflowsForEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWorkflowsForEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServiceServer).ListWorkflowsForEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WorkflowService/ListWorkflowsForEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServiceServer).ListWorkflowsForEvent(ctx, req.(*ListWorkflowsForEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkflowService_DeleteWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowServiceServer).DeleteWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WorkflowService/DeleteWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowServiceServer).DeleteWorkflow(ctx, req.(*DeleteWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkflowService_ServiceDesc is the grpc.ServiceDesc for WorkflowService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkflowService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "WorkflowService",
	HandlerType: (*WorkflowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListWorkflows",
			Handler:    _WorkflowService_ListWorkflows_Handler,
		},
		{
			MethodName: "PutWorkflow",
			Handler:    _WorkflowService_PutWorkflow_Handler,
		},
		{
			MethodName: "ScheduleWorkflow",
			Handler:    _WorkflowService_ScheduleWorkflow_Handler,
		},
		{
			MethodName: "TriggerWorkflow",
			Handler:    _WorkflowService_TriggerWorkflow_Handler,
		},
		{
			MethodName: "GetWorkflowByName",
			Handler:    _WorkflowService_GetWorkflowByName_Handler,
		},
		{
			MethodName: "ListWorkflowsForEvent",
			Handler:    _WorkflowService_ListWorkflowsForEvent_Handler,
		},
		{
			MethodName: "DeleteWorkflow",
			Handler:    _WorkflowService_DeleteWorkflow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workflows.proto",
}

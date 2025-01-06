// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.29.2
// source: dispatcher.proto

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

// DispatcherClient is the client API for Dispatcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DispatcherClient interface {
	Register(ctx context.Context, in *WorkerRegisterRequest, opts ...grpc.CallOption) (*WorkerRegisterResponse, error)
	Listen(ctx context.Context, in *WorkerListenRequest, opts ...grpc.CallOption) (Dispatcher_ListenClient, error)
	// ListenV2 is like listen, but implementation does not include heartbeats. This should only used by SDKs
	// against engine version v0.18.1+
	ListenV2(ctx context.Context, in *WorkerListenRequest, opts ...grpc.CallOption) (Dispatcher_ListenV2Client, error)
	// Heartbeat is a method for workers to send heartbeats to the dispatcher
	Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error)
	SubscribeToWorkflowEvents(ctx context.Context, in *SubscribeToWorkflowEventsRequest, opts ...grpc.CallOption) (Dispatcher_SubscribeToWorkflowEventsClient, error)
	SubscribeToWorkflowRuns(ctx context.Context, opts ...grpc.CallOption) (Dispatcher_SubscribeToWorkflowRunsClient, error)
	SendStepActionEvent(ctx context.Context, in *StepActionEvent, opts ...grpc.CallOption) (*ActionEventResponse, error)
	SendGroupKeyActionEvent(ctx context.Context, in *GroupKeyActionEvent, opts ...grpc.CallOption) (*ActionEventResponse, error)
	PutOverridesData(ctx context.Context, in *OverridesData, opts ...grpc.CallOption) (*OverridesDataResponse, error)
	Unsubscribe(ctx context.Context, in *WorkerUnsubscribeRequest, opts ...grpc.CallOption) (*WorkerUnsubscribeResponse, error)
	RefreshTimeout(ctx context.Context, in *RefreshTimeoutRequest, opts ...grpc.CallOption) (*RefreshTimeoutResponse, error)
	ReleaseSlot(ctx context.Context, in *ReleaseSlotRequest, opts ...grpc.CallOption) (*ReleaseSlotResponse, error)
	UpsertWorkerLabels(ctx context.Context, in *UpsertWorkerLabelsRequest, opts ...grpc.CallOption) (*UpsertWorkerLabelsResponse, error)
}

type dispatcherClient struct {
	cc grpc.ClientConnInterface
}

func NewDispatcherClient(cc grpc.ClientConnInterface) DispatcherClient {
	return &dispatcherClient{cc}
}

func (c *dispatcherClient) Register(ctx context.Context, in *WorkerRegisterRequest, opts ...grpc.CallOption) (*WorkerRegisterResponse, error) {
	out := new(WorkerRegisterResponse)
	err := c.cc.Invoke(ctx, "/Dispatcher/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherClient) Listen(ctx context.Context, in *WorkerListenRequest, opts ...grpc.CallOption) (Dispatcher_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &Dispatcher_ServiceDesc.Streams[0], "/Dispatcher/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &dispatcherListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Dispatcher_ListenClient interface {
	Recv() (*AssignedAction, error)
	grpc.ClientStream
}

type dispatcherListenClient struct {
	grpc.ClientStream
}

func (x *dispatcherListenClient) Recv() (*AssignedAction, error) {
	m := new(AssignedAction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dispatcherClient) ListenV2(ctx context.Context, in *WorkerListenRequest, opts ...grpc.CallOption) (Dispatcher_ListenV2Client, error) {
	stream, err := c.cc.NewStream(ctx, &Dispatcher_ServiceDesc.Streams[1], "/Dispatcher/ListenV2", opts...)
	if err != nil {
		return nil, err
	}
	x := &dispatcherListenV2Client{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Dispatcher_ListenV2Client interface {
	Recv() (*AssignedAction, error)
	grpc.ClientStream
}

type dispatcherListenV2Client struct {
	grpc.ClientStream
}

func (x *dispatcherListenV2Client) Recv() (*AssignedAction, error) {
	m := new(AssignedAction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dispatcherClient) Heartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error) {
	out := new(HeartbeatResponse)
	err := c.cc.Invoke(ctx, "/Dispatcher/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherClient) SubscribeToWorkflowEvents(ctx context.Context, in *SubscribeToWorkflowEventsRequest, opts ...grpc.CallOption) (Dispatcher_SubscribeToWorkflowEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Dispatcher_ServiceDesc.Streams[2], "/Dispatcher/SubscribeToWorkflowEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &dispatcherSubscribeToWorkflowEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Dispatcher_SubscribeToWorkflowEventsClient interface {
	Recv() (*WorkflowEvent, error)
	grpc.ClientStream
}

type dispatcherSubscribeToWorkflowEventsClient struct {
	grpc.ClientStream
}

func (x *dispatcherSubscribeToWorkflowEventsClient) Recv() (*WorkflowEvent, error) {
	m := new(WorkflowEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dispatcherClient) SubscribeToWorkflowRuns(ctx context.Context, opts ...grpc.CallOption) (Dispatcher_SubscribeToWorkflowRunsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Dispatcher_ServiceDesc.Streams[3], "/Dispatcher/SubscribeToWorkflowRuns", opts...)
	if err != nil {
		return nil, err
	}
	x := &dispatcherSubscribeToWorkflowRunsClient{stream}
	return x, nil
}

type Dispatcher_SubscribeToWorkflowRunsClient interface {
	Send(*SubscribeToWorkflowRunsRequest) error
	Recv() (*WorkflowRunEvent, error)
	grpc.ClientStream
}

type dispatcherSubscribeToWorkflowRunsClient struct {
	grpc.ClientStream
}

func (x *dispatcherSubscribeToWorkflowRunsClient) Send(m *SubscribeToWorkflowRunsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dispatcherSubscribeToWorkflowRunsClient) Recv() (*WorkflowRunEvent, error) {
	m := new(WorkflowRunEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dispatcherClient) SendStepActionEvent(ctx context.Context, in *StepActionEvent, opts ...grpc.CallOption) (*ActionEventResponse, error) {
	out := new(ActionEventResponse)
	err := c.cc.Invoke(ctx, "/Dispatcher/SendStepActionEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherClient) SendGroupKeyActionEvent(ctx context.Context, in *GroupKeyActionEvent, opts ...grpc.CallOption) (*ActionEventResponse, error) {
	out := new(ActionEventResponse)
	err := c.cc.Invoke(ctx, "/Dispatcher/SendGroupKeyActionEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherClient) PutOverridesData(ctx context.Context, in *OverridesData, opts ...grpc.CallOption) (*OverridesDataResponse, error) {
	out := new(OverridesDataResponse)
	err := c.cc.Invoke(ctx, "/Dispatcher/PutOverridesData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherClient) Unsubscribe(ctx context.Context, in *WorkerUnsubscribeRequest, opts ...grpc.CallOption) (*WorkerUnsubscribeResponse, error) {
	out := new(WorkerUnsubscribeResponse)
	err := c.cc.Invoke(ctx, "/Dispatcher/Unsubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherClient) RefreshTimeout(ctx context.Context, in *RefreshTimeoutRequest, opts ...grpc.CallOption) (*RefreshTimeoutResponse, error) {
	out := new(RefreshTimeoutResponse)
	err := c.cc.Invoke(ctx, "/Dispatcher/RefreshTimeout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherClient) ReleaseSlot(ctx context.Context, in *ReleaseSlotRequest, opts ...grpc.CallOption) (*ReleaseSlotResponse, error) {
	out := new(ReleaseSlotResponse)
	err := c.cc.Invoke(ctx, "/Dispatcher/ReleaseSlot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherClient) UpsertWorkerLabels(ctx context.Context, in *UpsertWorkerLabelsRequest, opts ...grpc.CallOption) (*UpsertWorkerLabelsResponse, error) {
	out := new(UpsertWorkerLabelsResponse)
	err := c.cc.Invoke(ctx, "/Dispatcher/UpsertWorkerLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DispatcherServer is the server API for Dispatcher service.
// All implementations must embed UnimplementedDispatcherServer
// for forward compatibility
type DispatcherServer interface {
	Register(context.Context, *WorkerRegisterRequest) (*WorkerRegisterResponse, error)
	Listen(*WorkerListenRequest, Dispatcher_ListenServer) error
	// ListenV2 is like listen, but implementation does not include heartbeats. This should only used by SDKs
	// against engine version v0.18.1+
	ListenV2(*WorkerListenRequest, Dispatcher_ListenV2Server) error
	// Heartbeat is a method for workers to send heartbeats to the dispatcher
	Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error)
	SubscribeToWorkflowEvents(*SubscribeToWorkflowEventsRequest, Dispatcher_SubscribeToWorkflowEventsServer) error
	SubscribeToWorkflowRuns(Dispatcher_SubscribeToWorkflowRunsServer) error
	SendStepActionEvent(context.Context, *StepActionEvent) (*ActionEventResponse, error)
	SendGroupKeyActionEvent(context.Context, *GroupKeyActionEvent) (*ActionEventResponse, error)
	PutOverridesData(context.Context, *OverridesData) (*OverridesDataResponse, error)
	Unsubscribe(context.Context, *WorkerUnsubscribeRequest) (*WorkerUnsubscribeResponse, error)
	RefreshTimeout(context.Context, *RefreshTimeoutRequest) (*RefreshTimeoutResponse, error)
	ReleaseSlot(context.Context, *ReleaseSlotRequest) (*ReleaseSlotResponse, error)
	UpsertWorkerLabels(context.Context, *UpsertWorkerLabelsRequest) (*UpsertWorkerLabelsResponse, error)
	mustEmbedUnimplementedDispatcherServer()
}

// UnimplementedDispatcherServer must be embedded to have forward compatible implementations.
type UnimplementedDispatcherServer struct {
}

func (UnimplementedDispatcherServer) Register(context.Context, *WorkerRegisterRequest) (*WorkerRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedDispatcherServer) Listen(*WorkerListenRequest, Dispatcher_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}
func (UnimplementedDispatcherServer) ListenV2(*WorkerListenRequest, Dispatcher_ListenV2Server) error {
	return status.Errorf(codes.Unimplemented, "method ListenV2 not implemented")
}
func (UnimplementedDispatcherServer) Heartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}
func (UnimplementedDispatcherServer) SubscribeToWorkflowEvents(*SubscribeToWorkflowEventsRequest, Dispatcher_SubscribeToWorkflowEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToWorkflowEvents not implemented")
}
func (UnimplementedDispatcherServer) SubscribeToWorkflowRuns(Dispatcher_SubscribeToWorkflowRunsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToWorkflowRuns not implemented")
}
func (UnimplementedDispatcherServer) SendStepActionEvent(context.Context, *StepActionEvent) (*ActionEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendStepActionEvent not implemented")
}
func (UnimplementedDispatcherServer) SendGroupKeyActionEvent(context.Context, *GroupKeyActionEvent) (*ActionEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGroupKeyActionEvent not implemented")
}
func (UnimplementedDispatcherServer) PutOverridesData(context.Context, *OverridesData) (*OverridesDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutOverridesData not implemented")
}
func (UnimplementedDispatcherServer) Unsubscribe(context.Context, *WorkerUnsubscribeRequest) (*WorkerUnsubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsubscribe not implemented")
}
func (UnimplementedDispatcherServer) RefreshTimeout(context.Context, *RefreshTimeoutRequest) (*RefreshTimeoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshTimeout not implemented")
}
func (UnimplementedDispatcherServer) ReleaseSlot(context.Context, *ReleaseSlotRequest) (*ReleaseSlotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseSlot not implemented")
}
func (UnimplementedDispatcherServer) UpsertWorkerLabels(context.Context, *UpsertWorkerLabelsRequest) (*UpsertWorkerLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertWorkerLabels not implemented")
}
func (UnimplementedDispatcherServer) mustEmbedUnimplementedDispatcherServer() {}

// UnsafeDispatcherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DispatcherServer will
// result in compilation errors.
type UnsafeDispatcherServer interface {
	mustEmbedUnimplementedDispatcherServer()
}

func RegisterDispatcherServer(s grpc.ServiceRegistrar, srv DispatcherServer) {
	s.RegisterService(&Dispatcher_ServiceDesc, srv)
}

func _Dispatcher_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dispatcher/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServer).Register(ctx, req.(*WorkerRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dispatcher_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WorkerListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DispatcherServer).Listen(m, &dispatcherListenServer{stream})
}

type Dispatcher_ListenServer interface {
	Send(*AssignedAction) error
	grpc.ServerStream
}

type dispatcherListenServer struct {
	grpc.ServerStream
}

func (x *dispatcherListenServer) Send(m *AssignedAction) error {
	return x.ServerStream.SendMsg(m)
}

func _Dispatcher_ListenV2_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WorkerListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DispatcherServer).ListenV2(m, &dispatcherListenV2Server{stream})
}

type Dispatcher_ListenV2Server interface {
	Send(*AssignedAction) error
	grpc.ServerStream
}

type dispatcherListenV2Server struct {
	grpc.ServerStream
}

func (x *dispatcherListenV2Server) Send(m *AssignedAction) error {
	return x.ServerStream.SendMsg(m)
}

func _Dispatcher_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dispatcher/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServer).Heartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dispatcher_SubscribeToWorkflowEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeToWorkflowEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DispatcherServer).SubscribeToWorkflowEvents(m, &dispatcherSubscribeToWorkflowEventsServer{stream})
}

type Dispatcher_SubscribeToWorkflowEventsServer interface {
	Send(*WorkflowEvent) error
	grpc.ServerStream
}

type dispatcherSubscribeToWorkflowEventsServer struct {
	grpc.ServerStream
}

func (x *dispatcherSubscribeToWorkflowEventsServer) Send(m *WorkflowEvent) error {
	return x.ServerStream.SendMsg(m)
}

func _Dispatcher_SubscribeToWorkflowRuns_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DispatcherServer).SubscribeToWorkflowRuns(&dispatcherSubscribeToWorkflowRunsServer{stream})
}

type Dispatcher_SubscribeToWorkflowRunsServer interface {
	Send(*WorkflowRunEvent) error
	Recv() (*SubscribeToWorkflowRunsRequest, error)
	grpc.ServerStream
}

type dispatcherSubscribeToWorkflowRunsServer struct {
	grpc.ServerStream
}

func (x *dispatcherSubscribeToWorkflowRunsServer) Send(m *WorkflowRunEvent) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dispatcherSubscribeToWorkflowRunsServer) Recv() (*SubscribeToWorkflowRunsRequest, error) {
	m := new(SubscribeToWorkflowRunsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Dispatcher_SendStepActionEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StepActionEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServer).SendStepActionEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dispatcher/SendStepActionEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServer).SendStepActionEvent(ctx, req.(*StepActionEvent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dispatcher_SendGroupKeyActionEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupKeyActionEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServer).SendGroupKeyActionEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dispatcher/SendGroupKeyActionEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServer).SendGroupKeyActionEvent(ctx, req.(*GroupKeyActionEvent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dispatcher_PutOverridesData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OverridesData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServer).PutOverridesData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dispatcher/PutOverridesData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServer).PutOverridesData(ctx, req.(*OverridesData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dispatcher_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerUnsubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dispatcher/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServer).Unsubscribe(ctx, req.(*WorkerUnsubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dispatcher_RefreshTimeout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTimeoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServer).RefreshTimeout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dispatcher/RefreshTimeout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServer).RefreshTimeout(ctx, req.(*RefreshTimeoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dispatcher_ReleaseSlot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseSlotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServer).ReleaseSlot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dispatcher/ReleaseSlot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServer).ReleaseSlot(ctx, req.(*ReleaseSlotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dispatcher_UpsertWorkerLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertWorkerLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServer).UpsertWorkerLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Dispatcher/UpsertWorkerLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServer).UpsertWorkerLabels(ctx, req.(*UpsertWorkerLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Dispatcher_ServiceDesc is the grpc.ServiceDesc for Dispatcher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dispatcher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Dispatcher",
	HandlerType: (*DispatcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Dispatcher_Register_Handler,
		},
		{
			MethodName: "Heartbeat",
			Handler:    _Dispatcher_Heartbeat_Handler,
		},
		{
			MethodName: "SendStepActionEvent",
			Handler:    _Dispatcher_SendStepActionEvent_Handler,
		},
		{
			MethodName: "SendGroupKeyActionEvent",
			Handler:    _Dispatcher_SendGroupKeyActionEvent_Handler,
		},
		{
			MethodName: "PutOverridesData",
			Handler:    _Dispatcher_PutOverridesData_Handler,
		},
		{
			MethodName: "Unsubscribe",
			Handler:    _Dispatcher_Unsubscribe_Handler,
		},
		{
			MethodName: "RefreshTimeout",
			Handler:    _Dispatcher_RefreshTimeout_Handler,
		},
		{
			MethodName: "ReleaseSlot",
			Handler:    _Dispatcher_ReleaseSlot_Handler,
		},
		{
			MethodName: "UpsertWorkerLabels",
			Handler:    _Dispatcher_UpsertWorkerLabels_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Listen",
			Handler:       _Dispatcher_Listen_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListenV2",
			Handler:       _Dispatcher_ListenV2_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeToWorkflowEvents",
			Handler:       _Dispatcher_SubscribeToWorkflowEvents_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeToWorkflowRuns",
			Handler:       _Dispatcher_SubscribeToWorkflowRuns_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "dispatcher.proto",
}

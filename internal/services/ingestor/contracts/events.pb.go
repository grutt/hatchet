// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.3
// source: events.proto

package contracts

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the tenant id
	TenantId string `protobuf:"bytes,1,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	// the id of the event
	EventId string `protobuf:"bytes,2,opt,name=eventId,proto3" json:"eventId,omitempty"`
	// the key for the event
	Key string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	// the payload for the event
	Payload string `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	// when the event was generated
	EventTimestamp *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=eventTimestamp,proto3" json:"eventTimestamp,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *Event) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *Event) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Event) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *Event) GetEventTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.EventTimestamp
	}
	return nil
}

type PushEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the tenant id
	TenantId string `protobuf:"bytes,1,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	// the key for the event
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// the payload for the event
	Payload string `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	// when the event was generated
	EventTimestamp *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=eventTimestamp,proto3" json:"eventTimestamp,omitempty"`
}

func (x *PushEventRequest) Reset() {
	*x = PushEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushEventRequest) ProtoMessage() {}

func (x *PushEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushEventRequest.ProtoReflect.Descriptor instead.
func (*PushEventRequest) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{1}
}

func (x *PushEventRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *PushEventRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PushEventRequest) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *PushEventRequest) GetEventTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.EventTimestamp
	}
	return nil
}

type ListEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (required) the tenant id
	TenantId string `protobuf:"bytes,1,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	// (optional) the number of events to skip
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// (optional) the key for the event
	Key string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *ListEventRequest) Reset() {
	*x = ListEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEventRequest) ProtoMessage() {}

func (x *ListEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEventRequest.ProtoReflect.Descriptor instead.
func (*ListEventRequest) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{2}
}

func (x *ListEventRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *ListEventRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListEventRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ListEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the events
	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *ListEventResponse) Reset() {
	*x = ListEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEventResponse) ProtoMessage() {}

func (x *ListEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEventResponse.ProtoReflect.Descriptor instead.
func (*ListEventResponse) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{3}
}

func (x *ListEventResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type ReplayEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the tenant id
	TenantId string `protobuf:"bytes,1,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	// the event id to replay
	EventId string `protobuf:"bytes,2,opt,name=eventId,proto3" json:"eventId,omitempty"`
}

func (x *ReplayEventRequest) Reset() {
	*x = ReplayEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplayEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplayEventRequest) ProtoMessage() {}

func (x *ReplayEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplayEventRequest.ProtoReflect.Descriptor instead.
func (*ReplayEventRequest) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{4}
}

func (x *ReplayEventRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *ReplayEventRequest) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

var File_events_proto protoreflect.FileDescriptor

var file_events_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xad, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x42, 0x0a, 0x0e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22,
	0x9e, 0x01, 0x0a, 0x10, 0x50, 0x75, 0x73, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x42, 0x0a, 0x0e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x22, 0x58, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x33, 0x0a, 0x11, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x06, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0x4a, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x32, 0x99, 0x01, 0x0a, 0x0d,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a,
	0x04, 0x50, 0x75, 0x73, 0x68, 0x12, 0x11, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x22, 0x00, 0x12, 0x2f, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x11, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61,
	0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x74, 0x63, 0x68, 0x65, 0x74, 0x2d, 0x64, 0x65,
	0x76, 0x2f, 0x68, 0x61, 0x74, 0x63, 0x68, 0x65, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_events_proto_rawDescOnce sync.Once
	file_events_proto_rawDescData = file_events_proto_rawDesc
)

func file_events_proto_rawDescGZIP() []byte {
	file_events_proto_rawDescOnce.Do(func() {
		file_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_events_proto_rawDescData)
	})
	return file_events_proto_rawDescData
}

var file_events_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_events_proto_goTypes = []interface{}{
	(*Event)(nil),                 // 0: Event
	(*PushEventRequest)(nil),      // 1: PushEventRequest
	(*ListEventRequest)(nil),      // 2: ListEventRequest
	(*ListEventResponse)(nil),     // 3: ListEventResponse
	(*ReplayEventRequest)(nil),    // 4: ReplayEventRequest
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_events_proto_depIdxs = []int32{
	5, // 0: Event.eventTimestamp:type_name -> google.protobuf.Timestamp
	5, // 1: PushEventRequest.eventTimestamp:type_name -> google.protobuf.Timestamp
	0, // 2: ListEventResponse.events:type_name -> Event
	1, // 3: EventsService.Push:input_type -> PushEventRequest
	2, // 4: EventsService.List:input_type -> ListEventRequest
	4, // 5: EventsService.ReplaySingleEvent:input_type -> ReplayEventRequest
	0, // 6: EventsService.Push:output_type -> Event
	3, // 7: EventsService.List:output_type -> ListEventResponse
	0, // 8: EventsService.ReplaySingleEvent:output_type -> Event
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_events_proto_init() }
func file_events_proto_init() {
	if File_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushEventRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEventRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEventResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_events_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplayEventRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_events_proto_goTypes,
		DependencyIndexes: file_events_proto_depIdxs,
		MessageInfos:      file_events_proto_msgTypes,
	}.Build()
	File_events_proto = out.File
	file_events_proto_rawDesc = nil
	file_events_proto_goTypes = nil
	file_events_proto_depIdxs = nil
}

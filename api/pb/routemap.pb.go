// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.21.12
// source: routemap.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TerminalCreateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TerminalType  uint32                 `protobuf:"varint,2,opt,name=terminalType,proto3" json:"terminalType,omitempty"`
	Country       string                 `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	State         string                 `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	City          string                 `protobuf:"bytes,5,opt,name=city,proto3" json:"city,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TerminalCreateRequest) Reset() {
	*x = TerminalCreateRequest{}
	mi := &file_routemap_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TerminalCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminalCreateRequest) ProtoMessage() {}

func (x *TerminalCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routemap_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminalCreateRequest.ProtoReflect.Descriptor instead.
func (*TerminalCreateRequest) Descriptor() ([]byte, []int) {
	return file_routemap_proto_rawDescGZIP(), []int{0}
}

func (x *TerminalCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TerminalCreateRequest) GetTerminalType() uint32 {
	if x != nil {
		return x.TerminalType
	}
	return 0
}

func (x *TerminalCreateRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *TerminalCreateRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *TerminalCreateRequest) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

type TerminalCreateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TerminalID    string                 `protobuf:"bytes,1,opt,name=terminalID,proto3" json:"terminalID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TerminalCreateResponse) Reset() {
	*x = TerminalCreateResponse{}
	mi := &file_routemap_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TerminalCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminalCreateResponse) ProtoMessage() {}

func (x *TerminalCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_routemap_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminalCreateResponse.ProtoReflect.Descriptor instead.
func (*TerminalCreateResponse) Descriptor() ([]byte, []int) {
	return file_routemap_proto_rawDescGZIP(), []int{1}
}

func (x *TerminalCreateResponse) GetTerminalID() string {
	if x != nil {
		return x.TerminalID
	}
	return ""
}

type TerminalGetByIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TerminalID    string                 `protobuf:"bytes,1,opt,name=terminalID,proto3" json:"terminalID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TerminalGetByIDRequest) Reset() {
	*x = TerminalGetByIDRequest{}
	mi := &file_routemap_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TerminalGetByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminalGetByIDRequest) ProtoMessage() {}

func (x *TerminalGetByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routemap_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminalGetByIDRequest.ProtoReflect.Descriptor instead.
func (*TerminalGetByIDRequest) Descriptor() ([]byte, []int) {
	return file_routemap_proto_rawDescGZIP(), []int{2}
}

func (x *TerminalGetByIDRequest) GetTerminalID() string {
	if x != nil {
		return x.TerminalID
	}
	return ""
}

type Terminal struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TerminalType  uint32                 `protobuf:"varint,3,opt,name=terminalType,proto3" json:"terminalType,omitempty"`
	Country       string                 `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	State         string                 `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	City          string                 `protobuf:"bytes,6,opt,name=city,proto3" json:"city,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Terminal) Reset() {
	*x = Terminal{}
	mi := &file_routemap_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Terminal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Terminal) ProtoMessage() {}

func (x *Terminal) ProtoReflect() protoreflect.Message {
	mi := &file_routemap_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Terminal.ProtoReflect.Descriptor instead.
func (*Terminal) Descriptor() ([]byte, []int) {
	return file_routemap_proto_rawDescGZIP(), []int{3}
}

func (x *Terminal) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Terminal) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Terminal) GetTerminalType() uint32 {
	if x != nil {
		return x.TerminalType
	}
	return 0
}

func (x *Terminal) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Terminal) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Terminal) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

type RouteItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Source        *Terminal              `protobuf:"bytes,1,opt,name=Source,proto3" json:"Source,omitempty"`
	Destination   *Terminal              `protobuf:"bytes,2,opt,name=Destination,proto3" json:"Destination,omitempty"`
	RouteNumber   uint32                 `protobuf:"varint,3,opt,name=RouteNumber,proto3" json:"RouteNumber,omitempty"`
	TransportType uint32                 `protobuf:"varint,4,opt,name=TransportType,proto3" json:"TransportType,omitempty"`
	Distance      float32                `protobuf:"fixed32,5,opt,name=Distance,proto3" json:"Distance,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RouteItem) Reset() {
	*x = RouteItem{}
	mi := &file_routemap_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RouteItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteItem) ProtoMessage() {}

func (x *RouteItem) ProtoReflect() protoreflect.Message {
	mi := &file_routemap_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteItem.ProtoReflect.Descriptor instead.
func (*RouteItem) Descriptor() ([]byte, []int) {
	return file_routemap_proto_rawDescGZIP(), []int{4}
}

func (x *RouteItem) GetSource() *Terminal {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *RouteItem) GetDestination() *Terminal {
	if x != nil {
		return x.Destination
	}
	return nil
}

func (x *RouteItem) GetRouteNumber() uint32 {
	if x != nil {
		return x.RouteNumber
	}
	return 0
}

func (x *RouteItem) GetTransportType() uint32 {
	if x != nil {
		return x.TransportType
	}
	return 0
}

func (x *RouteItem) GetDistance() float32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

type CreateRouteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RouteItem     *RouteItem             `protobuf:"bytes,1,opt,name=routeItem,proto3" json:"routeItem,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRouteRequest) Reset() {
	*x = CreateRouteRequest{}
	mi := &file_routemap_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRouteRequest) ProtoMessage() {}

func (x *CreateRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routemap_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRouteRequest.ProtoReflect.Descriptor instead.
func (*CreateRouteRequest) Descriptor() ([]byte, []int) {
	return file_routemap_proto_rawDescGZIP(), []int{5}
}

func (x *CreateRouteRequest) GetRouteItem() *RouteItem {
	if x != nil {
		return x.RouteItem
	}
	return nil
}

type CreateRouteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRouteResponse) Reset() {
	*x = CreateRouteResponse{}
	mi := &file_routemap_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRouteResponse) ProtoMessage() {}

func (x *CreateRouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_routemap_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRouteResponse.ProtoReflect.Descriptor instead.
func (*CreateRouteResponse) Descriptor() ([]byte, []int) {
	return file_routemap_proto_rawDescGZIP(), []int{6}
}

func (x *CreateRouteResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Route struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RouteItem     *RouteItem             `protobuf:"bytes,2,opt,name=routeItem,proto3" json:"routeItem,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Route) Reset() {
	*x = Route{}
	mi := &file_routemap_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route) ProtoMessage() {}

func (x *Route) ProtoReflect() protoreflect.Message {
	mi := &file_routemap_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route.ProtoReflect.Descriptor instead.
func (*Route) Descriptor() ([]byte, []int) {
	return file_routemap_proto_rawDescGZIP(), []int{7}
}

func (x *Route) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Route) GetRouteItem() *RouteItem {
	if x != nil {
		return x.RouteItem
	}
	return nil
}

type GetRouteByIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRouteByIDRequest) Reset() {
	*x = GetRouteByIDRequest{}
	mi := &file_routemap_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRouteByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRouteByIDRequest) ProtoMessage() {}

func (x *GetRouteByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_routemap_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRouteByIDRequest.ProtoReflect.Descriptor instead.
func (*GetRouteByIDRequest) Descriptor() ([]byte, []int) {
	return file_routemap_proto_rawDescGZIP(), []int{8}
}

func (x *GetRouteByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_routemap_proto protoreflect.FileDescriptor

var file_routemap_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x93, 0x01, 0x0a, 0x15, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x22, 0x38, 0x0a, 0x16, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x61, 0x6c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x49, 0x44,
	0x22, 0x38, 0x0a, 0x16, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x49, 0x44, 0x22, 0x96, 0x01, 0x0a, 0x08, 0x54,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x22, 0xbf, 0x01, 0x0a, 0x09, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x21, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x52, 0x06, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x54, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x6c, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x44, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x3e, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x09, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x25, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x05,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x22,
	0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xeb, 0x01, 0x0a, 0x0f, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x6d, 0x61, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x16, 0x2e, 0x54,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x17, 0x2e, 0x54,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c,
	0x12, 0x38, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12,
	0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x47, 0x65,
	0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x71, 0x6f, 0x6c, 0x69, 0x62, 0x61, 0x62, 0x61,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_routemap_proto_rawDescOnce sync.Once
	file_routemap_proto_rawDescData = file_routemap_proto_rawDesc
)

func file_routemap_proto_rawDescGZIP() []byte {
	file_routemap_proto_rawDescOnce.Do(func() {
		file_routemap_proto_rawDescData = protoimpl.X.CompressGZIP(file_routemap_proto_rawDescData)
	})
	return file_routemap_proto_rawDescData
}

var file_routemap_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_routemap_proto_goTypes = []any{
	(*TerminalCreateRequest)(nil),  // 0: TerminalCreateRequest
	(*TerminalCreateResponse)(nil), // 1: TerminalCreateResponse
	(*TerminalGetByIDRequest)(nil), // 2: TerminalGetByIDRequest
	(*Terminal)(nil),               // 3: Terminal
	(*RouteItem)(nil),              // 4: RouteItem
	(*CreateRouteRequest)(nil),     // 5: CreateRouteRequest
	(*CreateRouteResponse)(nil),    // 6: CreateRouteResponse
	(*Route)(nil),                  // 7: Route
	(*GetRouteByIDRequest)(nil),    // 8: GetRouteByIDRequest
}
var file_routemap_proto_depIdxs = []int32{
	3, // 0: RouteItem.Source:type_name -> Terminal
	3, // 1: RouteItem.Destination:type_name -> Terminal
	4, // 2: CreateRouteRequest.routeItem:type_name -> RouteItem
	4, // 3: Route.routeItem:type_name -> RouteItem
	0, // 4: RoutemapService.CreateTerminal:input_type -> TerminalCreateRequest
	2, // 5: RoutemapService.GetTerminal:input_type -> TerminalGetByIDRequest
	5, // 6: RoutemapService.CreateRoute:input_type -> CreateRouteRequest
	8, // 7: RoutemapService.GetRoute:input_type -> GetRouteByIDRequest
	1, // 8: RoutemapService.CreateTerminal:output_type -> TerminalCreateResponse
	3, // 9: RoutemapService.GetTerminal:output_type -> Terminal
	6, // 10: RoutemapService.CreateRoute:output_type -> CreateRouteResponse
	7, // 11: RoutemapService.GetRoute:output_type -> Route
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_routemap_proto_init() }
func file_routemap_proto_init() {
	if File_routemap_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_routemap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_routemap_proto_goTypes,
		DependencyIndexes: file_routemap_proto_depIdxs,
		MessageInfos:      file_routemap_proto_msgTypes,
	}.Build()
	File_routemap_proto = out.File
	file_routemap_proto_rawDesc = nil
	file_routemap_proto_goTypes = nil
	file_routemap_proto_depIdxs = nil
}
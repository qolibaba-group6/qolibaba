
package user
import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)
const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)
type GetUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // شناسه کاربر
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	mi := &file_internal_user_service_user_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*GetUserRequest) ProtoMessage() {}
func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_user_service_user_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_internal_user_service_user_service_proto_rawDescGZIP(), []int{0}
}
func (x *GetUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}
type GetUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // نام کاربر
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}
func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	mi := &file_internal_user_service_user_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}
func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*GetUserResponse) ProtoMessage() {}
func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_user_service_user_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_internal_user_service_user_service_proto_rawDescGZIP(), []int{1}
}
func (x *GetUserResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}
var File_internal_user_service_user_service_proto protoreflect.FileDescriptor
var file_internal_user_service_user_service_proto_rawDesc = []byte{
	0x0a, 0x28, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x25, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x45, 0x0a, 0x0b, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x2f, 0x5a, 0x2d, 0x74, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x2d, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x2d, 0x61, 0x70, 0x70, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x75, 0x73, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}
var (
	file_internal_user_service_user_service_proto_rawDescOnce sync.Once
	file_internal_user_service_user_service_proto_rawDescData = file_internal_user_service_user_service_proto_rawDesc
)
func file_internal_user_service_user_service_proto_rawDescGZIP() []byte {
	file_internal_user_service_user_service_proto_rawDescOnce.Do(func() {
		file_internal_user_service_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_user_service_user_service_proto_rawDescData)
	})
	return file_internal_user_service_user_service_proto_rawDescData
}
var file_internal_user_service_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_user_service_user_service_proto_goTypes = []any{
	(*GetUserRequest)(nil),  // 0: user.GetUserRequest
	(*GetUserResponse)(nil), // 1: user.GetUserResponse
}
var file_internal_user_service_user_service_proto_depIdxs = []int32{
	0, // 0: user.UserService.GetUser:input_type -> user.GetUserRequest
	1, // 1: user.UserService.GetUser:output_type -> user.GetUserResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}
func init() { file_internal_user_service_user_service_proto_init() }
func file_internal_user_service_user_service_proto_init() {
	if File_internal_user_service_user_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_user_service_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_user_service_user_service_proto_goTypes,
		DependencyIndexes: file_internal_user_service_user_service_proto_depIdxs,
		MessageInfos:      file_internal_user_service_user_service_proto_msgTypes,
	}.Build()
	File_internal_user_service_user_service_proto = out.File
	file_internal_user_service_user_service_proto_rawDesc = nil
	file_internal_user_service_user_service_proto_goTypes = nil
	file_internal_user_service_user_service_proto_depIdxs = nil
}

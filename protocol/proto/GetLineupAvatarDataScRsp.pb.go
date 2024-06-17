// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetLineupAvatarDataScRsp.proto

package proto

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

type GetLineupAvatarDataScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarDataList []*LineupAvatarData `protobuf:"bytes,12,rep,name=avatar_data_list,json=avatarDataList,proto3" json:"avatar_data_list,omitempty"`
	Retcode        uint32              `protobuf:"varint,13,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *GetLineupAvatarDataScRsp) Reset() {
	*x = GetLineupAvatarDataScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetLineupAvatarDataScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLineupAvatarDataScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLineupAvatarDataScRsp) ProtoMessage() {}

func (x *GetLineupAvatarDataScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetLineupAvatarDataScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLineupAvatarDataScRsp.ProtoReflect.Descriptor instead.
func (*GetLineupAvatarDataScRsp) Descriptor() ([]byte, []int) {
	return file_GetLineupAvatarDataScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetLineupAvatarDataScRsp) GetAvatarDataList() []*LineupAvatarData {
	if x != nil {
		return x.AvatarDataList
	}
	return nil
}

func (x *GetLineupAvatarDataScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_GetLineupAvatarDataScRsp_proto protoreflect.FileDescriptor

var file_GetLineupAvatarDataScRsp_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x16, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4c,
	0x69, 0x6e, 0x65, 0x75, 0x70, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x63, 0x52, 0x73, 0x70, 0x12, 0x3b, 0x0a, 0x10, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x28, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e,
	0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetLineupAvatarDataScRsp_proto_rawDescOnce sync.Once
	file_GetLineupAvatarDataScRsp_proto_rawDescData = file_GetLineupAvatarDataScRsp_proto_rawDesc
)

func file_GetLineupAvatarDataScRsp_proto_rawDescGZIP() []byte {
	file_GetLineupAvatarDataScRsp_proto_rawDescOnce.Do(func() {
		file_GetLineupAvatarDataScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetLineupAvatarDataScRsp_proto_rawDescData)
	})
	return file_GetLineupAvatarDataScRsp_proto_rawDescData
}

var file_GetLineupAvatarDataScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetLineupAvatarDataScRsp_proto_goTypes = []interface{}{
	(*GetLineupAvatarDataScRsp)(nil), // 0: GetLineupAvatarDataScRsp
	(*LineupAvatarData)(nil),         // 1: LineupAvatarData
}
var file_GetLineupAvatarDataScRsp_proto_depIdxs = []int32{
	1, // 0: GetLineupAvatarDataScRsp.avatar_data_list:type_name -> LineupAvatarData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetLineupAvatarDataScRsp_proto_init() }
func file_GetLineupAvatarDataScRsp_proto_init() {
	if File_GetLineupAvatarDataScRsp_proto != nil {
		return
	}
	file_LineupAvatarData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetLineupAvatarDataScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLineupAvatarDataScRsp); i {
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
			RawDescriptor: file_GetLineupAvatarDataScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetLineupAvatarDataScRsp_proto_goTypes,
		DependencyIndexes: file_GetLineupAvatarDataScRsp_proto_depIdxs,
		MessageInfos:      file_GetLineupAvatarDataScRsp_proto_msgTypes,
	}.Build()
	File_GetLineupAvatarDataScRsp_proto = out.File
	file_GetLineupAvatarDataScRsp_proto_rawDesc = nil
	file_GetLineupAvatarDataScRsp_proto_goTypes = nil
	file_GetLineupAvatarDataScRsp_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PlayerAssistInfo.proto

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

type PlayerAssistInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssistAvatar *DisplayAvatarDetailInfo `protobuf:"bytes,10,opt,name=assist_avatar,json=assistAvatar,proto3" json:"assist_avatar,omitempty"`
	PlayerInfo   *PlayerSimpleInfo        `protobuf:"bytes,15,opt,name=player_info,json=playerInfo,proto3" json:"player_info,omitempty"`
}

func (x *PlayerAssistInfo) Reset() {
	*x = PlayerAssistInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerAssistInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerAssistInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerAssistInfo) ProtoMessage() {}

func (x *PlayerAssistInfo) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerAssistInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerAssistInfo.ProtoReflect.Descriptor instead.
func (*PlayerAssistInfo) Descriptor() ([]byte, []int) {
	return file_PlayerAssistInfo_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerAssistInfo) GetAssistAvatar() *DisplayAvatarDetailInfo {
	if x != nil {
		return x.AssistAvatar
	}
	return nil
}

func (x *PlayerAssistInfo) GetPlayerInfo() *PlayerSimpleInfo {
	if x != nil {
		return x.PlayerInfo
	}
	return nil
}

var File_PlayerAssistInfo_proto protoreflect.FileDescriptor

var file_PlayerAssistInfo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x85, 0x01, 0x0a, 0x10, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3d, 0x0a, 0x0d, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x5f, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x44, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x12, 0x32, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerAssistInfo_proto_rawDescOnce sync.Once
	file_PlayerAssistInfo_proto_rawDescData = file_PlayerAssistInfo_proto_rawDesc
)

func file_PlayerAssistInfo_proto_rawDescGZIP() []byte {
	file_PlayerAssistInfo_proto_rawDescOnce.Do(func() {
		file_PlayerAssistInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerAssistInfo_proto_rawDescData)
	})
	return file_PlayerAssistInfo_proto_rawDescData
}

var file_PlayerAssistInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerAssistInfo_proto_goTypes = []interface{}{
	(*PlayerAssistInfo)(nil),        // 0: PlayerAssistInfo
	(*DisplayAvatarDetailInfo)(nil), // 1: DisplayAvatarDetailInfo
	(*PlayerSimpleInfo)(nil),        // 2: PlayerSimpleInfo
}
var file_PlayerAssistInfo_proto_depIdxs = []int32{
	1, // 0: PlayerAssistInfo.assist_avatar:type_name -> DisplayAvatarDetailInfo
	2, // 1: PlayerAssistInfo.player_info:type_name -> PlayerSimpleInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_PlayerAssistInfo_proto_init() }
func file_PlayerAssistInfo_proto_init() {
	if File_PlayerAssistInfo_proto != nil {
		return
	}
	file_DisplayAvatarDetailInfo_proto_init()
	file_PlayerSimpleInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerAssistInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerAssistInfo); i {
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
			RawDescriptor: file_PlayerAssistInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerAssistInfo_proto_goTypes,
		DependencyIndexes: file_PlayerAssistInfo_proto_depIdxs,
		MessageInfos:      file_PlayerAssistInfo_proto_msgTypes,
	}.Build()
	File_PlayerAssistInfo_proto = out.File
	file_PlayerAssistInfo_proto_rawDesc = nil
	file_PlayerAssistInfo_proto_goTypes = nil
	file_PlayerAssistInfo_proto_depIdxs = nil
}

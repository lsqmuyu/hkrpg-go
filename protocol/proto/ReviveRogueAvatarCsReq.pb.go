// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ReviveRogueAvatarCsReq.proto

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

type ReviveRogueAvatarCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseAvatarIdList []uint32 `protobuf:"varint,2,rep,packed,name=base_avatar_id_list,json=baseAvatarIdList,proto3" json:"base_avatar_id_list,omitempty"`
	PPMPGOGOLIF      []uint32 `protobuf:"varint,13,rep,packed,name=PPMPGOGOLIF,proto3" json:"PPMPGOGOLIF,omitempty"`
	BaseAvatarId     uint32   `protobuf:"varint,4,opt,name=base_avatar_id,json=baseAvatarId,proto3" json:"base_avatar_id,omitempty"`
	MonsterId        uint32   `protobuf:"varint,7,opt,name=monster_id,json=monsterId,proto3" json:"monster_id,omitempty"`
}

func (x *ReviveRogueAvatarCsReq) Reset() {
	*x = ReviveRogueAvatarCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ReviveRogueAvatarCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReviveRogueAvatarCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviveRogueAvatarCsReq) ProtoMessage() {}

func (x *ReviveRogueAvatarCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_ReviveRogueAvatarCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviveRogueAvatarCsReq.ProtoReflect.Descriptor instead.
func (*ReviveRogueAvatarCsReq) Descriptor() ([]byte, []int) {
	return file_ReviveRogueAvatarCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *ReviveRogueAvatarCsReq) GetBaseAvatarIdList() []uint32 {
	if x != nil {
		return x.BaseAvatarIdList
	}
	return nil
}

func (x *ReviveRogueAvatarCsReq) GetPPMPGOGOLIF() []uint32 {
	if x != nil {
		return x.PPMPGOGOLIF
	}
	return nil
}

func (x *ReviveRogueAvatarCsReq) GetBaseAvatarId() uint32 {
	if x != nil {
		return x.BaseAvatarId
	}
	return 0
}

func (x *ReviveRogueAvatarCsReq) GetMonsterId() uint32 {
	if x != nil {
		return x.MonsterId
	}
	return 0
}

var File_ReviveRogueAvatarCsReq_proto protoreflect.FileDescriptor

var file_ReviveRogueAvatarCsReq_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x52, 0x65, 0x76, 0x69, 0x76, 0x65, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae,
	0x01, 0x0a, 0x16, 0x52, 0x65, 0x76, 0x69, 0x76, 0x65, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x2d, 0x0a, 0x13, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x10, 0x62, 0x61, 0x73, 0x65, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x50, 0x4d, 0x50,
	0x47, 0x4f, 0x47, 0x4f, 0x4c, 0x49, 0x46, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x50,
	0x50, 0x4d, 0x50, 0x47, 0x4f, 0x47, 0x4f, 0x4c, 0x49, 0x46, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x42,
	0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67,
	0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_ReviveRogueAvatarCsReq_proto_rawDescOnce sync.Once
	file_ReviveRogueAvatarCsReq_proto_rawDescData = file_ReviveRogueAvatarCsReq_proto_rawDesc
)

func file_ReviveRogueAvatarCsReq_proto_rawDescGZIP() []byte {
	file_ReviveRogueAvatarCsReq_proto_rawDescOnce.Do(func() {
		file_ReviveRogueAvatarCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_ReviveRogueAvatarCsReq_proto_rawDescData)
	})
	return file_ReviveRogueAvatarCsReq_proto_rawDescData
}

var file_ReviveRogueAvatarCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ReviveRogueAvatarCsReq_proto_goTypes = []interface{}{
	(*ReviveRogueAvatarCsReq)(nil), // 0: ReviveRogueAvatarCsReq
}
var file_ReviveRogueAvatarCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ReviveRogueAvatarCsReq_proto_init() }
func file_ReviveRogueAvatarCsReq_proto_init() {
	if File_ReviveRogueAvatarCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ReviveRogueAvatarCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReviveRogueAvatarCsReq); i {
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
			RawDescriptor: file_ReviveRogueAvatarCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ReviveRogueAvatarCsReq_proto_goTypes,
		DependencyIndexes: file_ReviveRogueAvatarCsReq_proto_depIdxs,
		MessageInfos:      file_ReviveRogueAvatarCsReq_proto_msgTypes,
	}.Build()
	File_ReviveRogueAvatarCsReq_proto = out.File
	file_ReviveRogueAvatarCsReq_proto_rawDesc = nil
	file_ReviveRogueAvatarCsReq_proto_goTypes = nil
	file_ReviveRogueAvatarCsReq_proto_depIdxs = nil
}

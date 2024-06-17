// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueUpdateUnlockLevelScNotify.proto

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

type ChessRogueUpdateUnlockLevelScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AreaIdList []uint32 `protobuf:"varint,14,rep,packed,name=area_id_list,json=areaIdList,proto3" json:"area_id_list,omitempty"`
}

func (x *ChessRogueUpdateUnlockLevelScNotify) Reset() {
	*x = ChessRogueUpdateUnlockLevelScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessRogueUpdateUnlockLevelScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessRogueUpdateUnlockLevelScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessRogueUpdateUnlockLevelScNotify) ProtoMessage() {}

func (x *ChessRogueUpdateUnlockLevelScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ChessRogueUpdateUnlockLevelScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessRogueUpdateUnlockLevelScNotify.ProtoReflect.Descriptor instead.
func (*ChessRogueUpdateUnlockLevelScNotify) Descriptor() ([]byte, []int) {
	return file_ChessRogueUpdateUnlockLevelScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ChessRogueUpdateUnlockLevelScNotify) GetAreaIdList() []uint32 {
	if x != nil {
		return x.AreaIdList
	}
	return nil
}

var File_ChessRogueUpdateUnlockLevelScNotify_proto protoreflect.FileDescriptor

var file_ChessRogueUpdateUnlockLevelScNotify_proto_rawDesc = []byte{
	0x0a, 0x29, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x53, 0x63, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x23, 0x43,
	0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x20, 0x0a, 0x0c, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueUpdateUnlockLevelScNotify_proto_rawDescOnce sync.Once
	file_ChessRogueUpdateUnlockLevelScNotify_proto_rawDescData = file_ChessRogueUpdateUnlockLevelScNotify_proto_rawDesc
)

func file_ChessRogueUpdateUnlockLevelScNotify_proto_rawDescGZIP() []byte {
	file_ChessRogueUpdateUnlockLevelScNotify_proto_rawDescOnce.Do(func() {
		file_ChessRogueUpdateUnlockLevelScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueUpdateUnlockLevelScNotify_proto_rawDescData)
	})
	return file_ChessRogueUpdateUnlockLevelScNotify_proto_rawDescData
}

var file_ChessRogueUpdateUnlockLevelScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChessRogueUpdateUnlockLevelScNotify_proto_goTypes = []interface{}{
	(*ChessRogueUpdateUnlockLevelScNotify)(nil), // 0: ChessRogueUpdateUnlockLevelScNotify
}
var file_ChessRogueUpdateUnlockLevelScNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ChessRogueUpdateUnlockLevelScNotify_proto_init() }
func file_ChessRogueUpdateUnlockLevelScNotify_proto_init() {
	if File_ChessRogueUpdateUnlockLevelScNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ChessRogueUpdateUnlockLevelScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessRogueUpdateUnlockLevelScNotify); i {
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
			RawDescriptor: file_ChessRogueUpdateUnlockLevelScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueUpdateUnlockLevelScNotify_proto_goTypes,
		DependencyIndexes: file_ChessRogueUpdateUnlockLevelScNotify_proto_depIdxs,
		MessageInfos:      file_ChessRogueUpdateUnlockLevelScNotify_proto_msgTypes,
	}.Build()
	File_ChessRogueUpdateUnlockLevelScNotify_proto = out.File
	file_ChessRogueUpdateUnlockLevelScNotify_proto_rawDesc = nil
	file_ChessRogueUpdateUnlockLevelScNotify_proto_goTypes = nil
	file_ChessRogueUpdateUnlockLevelScNotify_proto_depIdxs = nil
}

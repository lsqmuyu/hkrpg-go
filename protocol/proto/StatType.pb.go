// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: StatType.proto

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

type StatType int32

const (
	StatType_STAT_TYPE_NONE    StatType = 0
	StatType_STAT_TYPE_ART     StatType = 1
	StatType_STAT_TYPE_CULTURE StatType = 2
	StatType_STAT_TYPE_POPULAR StatType = 3
)

// Enum value maps for StatType.
var (
	StatType_name = map[int32]string{
		0: "STAT_TYPE_NONE",
		1: "STAT_TYPE_ART",
		2: "STAT_TYPE_CULTURE",
		3: "STAT_TYPE_POPULAR",
	}
	StatType_value = map[string]int32{
		"STAT_TYPE_NONE":    0,
		"STAT_TYPE_ART":     1,
		"STAT_TYPE_CULTURE": 2,
		"STAT_TYPE_POPULAR": 3,
	}
)

func (x StatType) Enum() *StatType {
	p := new(StatType)
	*p = x
	return p
}

func (x StatType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatType) Descriptor() protoreflect.EnumDescriptor {
	return file_StatType_proto_enumTypes[0].Descriptor()
}

func (StatType) Type() protoreflect.EnumType {
	return &file_StatType_proto_enumTypes[0]
}

func (x StatType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatType.Descriptor instead.
func (StatType) EnumDescriptor() ([]byte, []int) {
	return file_StatType_proto_rawDescGZIP(), []int{0}
}

var File_StatType_proto protoreflect.FileDescriptor

var file_StatType_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2a, 0x5f, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e,
	0x53, 0x54, 0x41, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x52,
	0x54, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x43, 0x55, 0x4c, 0x54, 0x55, 0x52, 0x45, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54,
	0x41, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4f, 0x50, 0x55, 0x4c, 0x41, 0x52, 0x10,
	0x03, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_StatType_proto_rawDescOnce sync.Once
	file_StatType_proto_rawDescData = file_StatType_proto_rawDesc
)

func file_StatType_proto_rawDescGZIP() []byte {
	file_StatType_proto_rawDescOnce.Do(func() {
		file_StatType_proto_rawDescData = protoimpl.X.CompressGZIP(file_StatType_proto_rawDescData)
	})
	return file_StatType_proto_rawDescData
}

var file_StatType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_StatType_proto_goTypes = []interface{}{
	(StatType)(0), // 0: StatType
}
var file_StatType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_StatType_proto_init() }
func file_StatType_proto_init() {
	if File_StatType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_StatType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_StatType_proto_goTypes,
		DependencyIndexes: file_StatType_proto_depIdxs,
		EnumInfos:         file_StatType_proto_enumTypes,
	}.Build()
	File_StatType_proto = out.File
	file_StatType_proto_rawDesc = nil
	file_StatType_proto_goTypes = nil
	file_StatType_proto_depIdxs = nil
}

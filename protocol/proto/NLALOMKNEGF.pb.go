// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: NLALOMKNEGF.proto

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

type NLALOMKNEGF struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MapId       uint32            `protobuf:"varint,4,opt,name=map_id,json=mapId,proto3" json:"map_id,omitempty"`
	LDEFMDLNCPD MonopolyCellState `protobuf:"varint,7,opt,name=LDEFMDLNCPD,proto3,enum=MonopolyCellState" json:"LDEFMDLNCPD,omitempty"`
	JINKIOJFAAA bool              `protobuf:"varint,5,opt,name=JINKIOJFAAA,proto3" json:"JINKIOJFAAA,omitempty"`
	CellId      uint32            `protobuf:"varint,13,opt,name=cell_id,json=cellId,proto3" json:"cell_id,omitempty"`
}

func (x *NLALOMKNEGF) Reset() {
	*x = NLALOMKNEGF{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NLALOMKNEGF_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NLALOMKNEGF) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NLALOMKNEGF) ProtoMessage() {}

func (x *NLALOMKNEGF) ProtoReflect() protoreflect.Message {
	mi := &file_NLALOMKNEGF_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NLALOMKNEGF.ProtoReflect.Descriptor instead.
func (*NLALOMKNEGF) Descriptor() ([]byte, []int) {
	return file_NLALOMKNEGF_proto_rawDescGZIP(), []int{0}
}

func (x *NLALOMKNEGF) GetMapId() uint32 {
	if x != nil {
		return x.MapId
	}
	return 0
}

func (x *NLALOMKNEGF) GetLDEFMDLNCPD() MonopolyCellState {
	if x != nil {
		return x.LDEFMDLNCPD
	}
	return MonopolyCellState_MONOPOLY_CELL_STATE_IDLE
}

func (x *NLALOMKNEGF) GetJINKIOJFAAA() bool {
	if x != nil {
		return x.JINKIOJFAAA
	}
	return false
}

func (x *NLALOMKNEGF) GetCellId() uint32 {
	if x != nil {
		return x.CellId
	}
	return 0
}

var File_NLALOMKNEGF_proto protoreflect.FileDescriptor

var file_NLALOMKNEGF_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4e, 0x4c, 0x41, 0x4c, 0x4f, 0x4d, 0x4b, 0x4e, 0x45, 0x47, 0x46, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f, 0x6c, 0x79, 0x43, 0x65, 0x6c,
	0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a,
	0x0b, 0x4e, 0x4c, 0x41, 0x4c, 0x4f, 0x4d, 0x4b, 0x4e, 0x45, 0x47, 0x46, 0x12, 0x15, 0x0a, 0x06,
	0x6d, 0x61, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6d, 0x61,
	0x70, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x0b, 0x4c, 0x44, 0x45, 0x46, 0x4d, 0x44, 0x4c, 0x4e, 0x43,
	0x50, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x4d, 0x6f, 0x6e, 0x6f, 0x70,
	0x6f, 0x6c, 0x79, 0x43, 0x65, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x4c, 0x44,
	0x45, 0x46, 0x4d, 0x44, 0x4c, 0x4e, 0x43, 0x50, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x49, 0x4e,
	0x4b, 0x49, 0x4f, 0x4a, 0x46, 0x41, 0x41, 0x41, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x4a, 0x49, 0x4e, 0x4b, 0x49, 0x4f, 0x4a, 0x46, 0x41, 0x41, 0x41, 0x12, 0x17, 0x0a, 0x07, 0x63,
	0x65, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x65,
	0x6c, 0x6c, 0x49, 0x64, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_NLALOMKNEGF_proto_rawDescOnce sync.Once
	file_NLALOMKNEGF_proto_rawDescData = file_NLALOMKNEGF_proto_rawDesc
)

func file_NLALOMKNEGF_proto_rawDescGZIP() []byte {
	file_NLALOMKNEGF_proto_rawDescOnce.Do(func() {
		file_NLALOMKNEGF_proto_rawDescData = protoimpl.X.CompressGZIP(file_NLALOMKNEGF_proto_rawDescData)
	})
	return file_NLALOMKNEGF_proto_rawDescData
}

var file_NLALOMKNEGF_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_NLALOMKNEGF_proto_goTypes = []interface{}{
	(*NLALOMKNEGF)(nil),    // 0: NLALOMKNEGF
	(MonopolyCellState)(0), // 1: MonopolyCellState
}
var file_NLALOMKNEGF_proto_depIdxs = []int32{
	1, // 0: NLALOMKNEGF.LDEFMDLNCPD:type_name -> MonopolyCellState
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_NLALOMKNEGF_proto_init() }
func file_NLALOMKNEGF_proto_init() {
	if File_NLALOMKNEGF_proto != nil {
		return
	}
	file_MonopolyCellState_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_NLALOMKNEGF_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NLALOMKNEGF); i {
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
			RawDescriptor: file_NLALOMKNEGF_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_NLALOMKNEGF_proto_goTypes,
		DependencyIndexes: file_NLALOMKNEGF_proto_depIdxs,
		MessageInfos:      file_NLALOMKNEGF_proto_msgTypes,
	}.Build()
	File_NLALOMKNEGF_proto = out.File
	file_NLALOMKNEGF_proto_rawDesc = nil
	file_NLALOMKNEGF_proto_goTypes = nil
	file_NLALOMKNEGF_proto_depIdxs = nil
}

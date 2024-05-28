// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: EKOPDGAFLDA.proto

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

type EKOPDGAFLDA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  AetherdivideSpiritLineupStatus `protobuf:"varint,1,opt,name=type,proto3,enum=AetherdivideSpiritLineupStatus" json:"type,omitempty"`
	Id    uint32                         `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	SpBar *SpBarInfo                     `protobuf:"bytes,3,opt,name=sp_bar,json=spBar,proto3" json:"sp_bar,omitempty"`
}

func (x *EKOPDGAFLDA) Reset() {
	*x = EKOPDGAFLDA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EKOPDGAFLDA_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EKOPDGAFLDA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EKOPDGAFLDA) ProtoMessage() {}

func (x *EKOPDGAFLDA) ProtoReflect() protoreflect.Message {
	mi := &file_EKOPDGAFLDA_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EKOPDGAFLDA.ProtoReflect.Descriptor instead.
func (*EKOPDGAFLDA) Descriptor() ([]byte, []int) {
	return file_EKOPDGAFLDA_proto_rawDescGZIP(), []int{0}
}

func (x *EKOPDGAFLDA) GetType() AetherdivideSpiritLineupStatus {
	if x != nil {
		return x.Type
	}
	return AetherdivideSpiritLineupStatus_AETHERDIVIDE_SPIRIT_LINEUP_NONE
}

func (x *EKOPDGAFLDA) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EKOPDGAFLDA) GetSpBar() *SpBarInfo {
	if x != nil {
		return x.SpBar
	}
	return nil
}

var File_EKOPDGAFLDA_proto protoreflect.FileDescriptor

var file_EKOPDGAFLDA_proto_rawDesc = []byte{
	0x0a, 0x11, 0x45, 0x4b, 0x4f, 0x50, 0x44, 0x47, 0x41, 0x46, 0x4c, 0x44, 0x41, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x53, 0x70, 0x42, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x41, 0x65, 0x74, 0x68, 0x65, 0x72, 0x64, 0x69, 0x76, 0x69,
	0x64, 0x65, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x0b, 0x45, 0x4b,
	0x4f, 0x50, 0x44, 0x47, 0x41, 0x46, 0x4c, 0x44, 0x41, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x41, 0x65, 0x74, 0x68, 0x65, 0x72,
	0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x4c, 0x69, 0x6e, 0x65,
	0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21,
	0x0a, 0x06, 0x73, 0x70, 0x5f, 0x62, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x53, 0x70, 0x42, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x73, 0x70, 0x42, 0x61,
	0x72, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_EKOPDGAFLDA_proto_rawDescOnce sync.Once
	file_EKOPDGAFLDA_proto_rawDescData = file_EKOPDGAFLDA_proto_rawDesc
)

func file_EKOPDGAFLDA_proto_rawDescGZIP() []byte {
	file_EKOPDGAFLDA_proto_rawDescOnce.Do(func() {
		file_EKOPDGAFLDA_proto_rawDescData = protoimpl.X.CompressGZIP(file_EKOPDGAFLDA_proto_rawDescData)
	})
	return file_EKOPDGAFLDA_proto_rawDescData
}

var file_EKOPDGAFLDA_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EKOPDGAFLDA_proto_goTypes = []interface{}{
	(*EKOPDGAFLDA)(nil),                 // 0: EKOPDGAFLDA
	(AetherdivideSpiritLineupStatus)(0), // 1: AetherdivideSpiritLineupStatus
	(*SpBarInfo)(nil),                   // 2: SpBarInfo
}
var file_EKOPDGAFLDA_proto_depIdxs = []int32{
	1, // 0: EKOPDGAFLDA.type:type_name -> AetherdivideSpiritLineupStatus
	2, // 1: EKOPDGAFLDA.sp_bar:type_name -> SpBarInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_EKOPDGAFLDA_proto_init() }
func file_EKOPDGAFLDA_proto_init() {
	if File_EKOPDGAFLDA_proto != nil {
		return
	}
	file_SpBarInfo_proto_init()
	file_AetherdivideSpiritLineupStatus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EKOPDGAFLDA_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EKOPDGAFLDA); i {
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
			RawDescriptor: file_EKOPDGAFLDA_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EKOPDGAFLDA_proto_goTypes,
		DependencyIndexes: file_EKOPDGAFLDA_proto_depIdxs,
		MessageInfos:      file_EKOPDGAFLDA_proto_msgTypes,
	}.Build()
	File_EKOPDGAFLDA_proto = out.File
	file_EKOPDGAFLDA_proto_rawDesc = nil
	file_EKOPDGAFLDA_proto_goTypes = nil
	file_EKOPDGAFLDA_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: CFGILHCKECI.proto

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

type CFGILHCKECI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JFBOMNGKBFJ        *NLONGNDDAJK            `protobuf:"bytes,3,opt,name=JFBOMNGKBFJ,proto3" json:"JFBOMNGKBFJ,omitempty"`
	ModifierSourceType RogueModifierSourceType `protobuf:"varint,4,opt,name=modifier_source_type,json=modifierSourceType,proto3,enum=RogueModifierSourceType" json:"modifier_source_type,omitempty"`
	NPNBOEJCNCD        uint64                  `protobuf:"varint,1,opt,name=NPNBOEJCNCD,proto3" json:"NPNBOEJCNCD,omitempty"`
	// Types that are assignable to FPDBFGNEDBB:
	//
	//	*CFGILHCKECI_BFHLEPABBIE
	FPDBFGNEDBB isCFGILHCKECI_FPDBFGNEDBB `protobuf_oneof:"FPDBFGNEDBB"`
}

func (x *CFGILHCKECI) Reset() {
	*x = CFGILHCKECI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CFGILHCKECI_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CFGILHCKECI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CFGILHCKECI) ProtoMessage() {}

func (x *CFGILHCKECI) ProtoReflect() protoreflect.Message {
	mi := &file_CFGILHCKECI_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CFGILHCKECI.ProtoReflect.Descriptor instead.
func (*CFGILHCKECI) Descriptor() ([]byte, []int) {
	return file_CFGILHCKECI_proto_rawDescGZIP(), []int{0}
}

func (x *CFGILHCKECI) GetJFBOMNGKBFJ() *NLONGNDDAJK {
	if x != nil {
		return x.JFBOMNGKBFJ
	}
	return nil
}

func (x *CFGILHCKECI) GetModifierSourceType() RogueModifierSourceType {
	if x != nil {
		return x.ModifierSourceType
	}
	return RogueModifierSourceType_ROGUE_MODIFIER_SOURCE_NONE
}

func (x *CFGILHCKECI) GetNPNBOEJCNCD() uint64 {
	if x != nil {
		return x.NPNBOEJCNCD
	}
	return 0
}

func (m *CFGILHCKECI) GetFPDBFGNEDBB() isCFGILHCKECI_FPDBFGNEDBB {
	if m != nil {
		return m.FPDBFGNEDBB
	}
	return nil
}

func (x *CFGILHCKECI) GetBFHLEPABBIE() *HNHFKCHEPNF {
	if x, ok := x.GetFPDBFGNEDBB().(*CFGILHCKECI_BFHLEPABBIE); ok {
		return x.BFHLEPABBIE
	}
	return nil
}

type isCFGILHCKECI_FPDBFGNEDBB interface {
	isCFGILHCKECI_FPDBFGNEDBB()
}

type CFGILHCKECI_BFHLEPABBIE struct {
	BFHLEPABBIE *HNHFKCHEPNF `protobuf:"bytes,1033,opt,name=BFHLEPABBIE,proto3,oneof"`
}

func (*CFGILHCKECI_BFHLEPABBIE) isCFGILHCKECI_FPDBFGNEDBB() {}

var File_CFGILHCKECI_proto protoreflect.FileDescriptor

var file_CFGILHCKECI_proto_rawDesc = []byte{
	0x0a, 0x11, 0x43, 0x46, 0x47, 0x49, 0x4c, 0x48, 0x43, 0x4b, 0x45, 0x43, 0x49, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x48, 0x4e, 0x48, 0x46, 0x4b, 0x43, 0x48, 0x45, 0x50, 0x4e, 0x46,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4e, 0x4c, 0x4f, 0x4e, 0x47, 0x4e, 0x44, 0x44,
	0x41, 0x4a, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x01, 0x0a, 0x0b, 0x43, 0x46, 0x47,
	0x49, 0x4c, 0x48, 0x43, 0x4b, 0x45, 0x43, 0x49, 0x12, 0x2e, 0x0a, 0x0b, 0x4a, 0x46, 0x42, 0x4f,
	0x4d, 0x4e, 0x47, 0x4b, 0x42, 0x46, 0x4a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x4e, 0x4c, 0x4f, 0x4e, 0x47, 0x4e, 0x44, 0x44, 0x41, 0x4a, 0x4b, 0x52, 0x0b, 0x4a, 0x46, 0x42,
	0x4f, 0x4d, 0x4e, 0x47, 0x4b, 0x42, 0x46, 0x4a, 0x12, 0x4a, 0x0a, 0x14, 0x6d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x12, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x50, 0x4e, 0x42, 0x4f, 0x45, 0x4a, 0x43,
	0x4e, 0x43, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x4e, 0x50, 0x4e, 0x42, 0x4f,
	0x45, 0x4a, 0x43, 0x4e, 0x43, 0x44, 0x12, 0x31, 0x0a, 0x0b, 0x42, 0x46, 0x48, 0x4c, 0x45, 0x50,
	0x41, 0x42, 0x42, 0x49, 0x45, 0x18, 0x89, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x48,
	0x4e, 0x48, 0x46, 0x4b, 0x43, 0x48, 0x45, 0x50, 0x4e, 0x46, 0x48, 0x00, 0x52, 0x0b, 0x42, 0x46,
	0x48, 0x4c, 0x45, 0x50, 0x41, 0x42, 0x42, 0x49, 0x45, 0x42, 0x0d, 0x0a, 0x0b, 0x46, 0x50, 0x44,
	0x42, 0x46, 0x47, 0x4e, 0x45, 0x44, 0x42, 0x42, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44,
	0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CFGILHCKECI_proto_rawDescOnce sync.Once
	file_CFGILHCKECI_proto_rawDescData = file_CFGILHCKECI_proto_rawDesc
)

func file_CFGILHCKECI_proto_rawDescGZIP() []byte {
	file_CFGILHCKECI_proto_rawDescOnce.Do(func() {
		file_CFGILHCKECI_proto_rawDescData = protoimpl.X.CompressGZIP(file_CFGILHCKECI_proto_rawDescData)
	})
	return file_CFGILHCKECI_proto_rawDescData
}

var file_CFGILHCKECI_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CFGILHCKECI_proto_goTypes = []interface{}{
	(*CFGILHCKECI)(nil),          // 0: CFGILHCKECI
	(*NLONGNDDAJK)(nil),          // 1: NLONGNDDAJK
	(RogueModifierSourceType)(0), // 2: RogueModifierSourceType
	(*HNHFKCHEPNF)(nil),          // 3: HNHFKCHEPNF
}
var file_CFGILHCKECI_proto_depIdxs = []int32{
	1, // 0: CFGILHCKECI.JFBOMNGKBFJ:type_name -> NLONGNDDAJK
	2, // 1: CFGILHCKECI.modifier_source_type:type_name -> RogueModifierSourceType
	3, // 2: CFGILHCKECI.BFHLEPABBIE:type_name -> HNHFKCHEPNF
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_CFGILHCKECI_proto_init() }
func file_CFGILHCKECI_proto_init() {
	if File_CFGILHCKECI_proto != nil {
		return
	}
	file_HNHFKCHEPNF_proto_init()
	file_NLONGNDDAJK_proto_init()
	file_RogueModifierSourceType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CFGILHCKECI_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CFGILHCKECI); i {
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
	file_CFGILHCKECI_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CFGILHCKECI_BFHLEPABBIE)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_CFGILHCKECI_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CFGILHCKECI_proto_goTypes,
		DependencyIndexes: file_CFGILHCKECI_proto_depIdxs,
		MessageInfos:      file_CFGILHCKECI_proto_msgTypes,
	}.Build()
	File_CFGILHCKECI_proto = out.File
	file_CFGILHCKECI_proto_rawDesc = nil
	file_CFGILHCKECI_proto_goTypes = nil
	file_CFGILHCKECI_proto_depIdxs = nil
}

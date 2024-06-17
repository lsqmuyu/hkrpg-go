// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: WolfBroGameInfo.proto

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

type WolfBroGameInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CHGGJGHBNBM uint32      `protobuf:"varint,12,opt,name=CHGGJGHBNBM,proto3" json:"CHGGJGHBNBM,omitempty"`
	Motion      *MotionInfo `protobuf:"bytes,6,opt,name=motion,proto3" json:"motion,omitempty"`
	LFEHNFOPJCL bool        `protobuf:"varint,2,opt,name=LFEHNFOPJCL,proto3" json:"LFEHNFOPJCL,omitempty"`
	MCJJKBHLAJN []*Vector   `protobuf:"bytes,4,rep,name=MCJJKBHLAJN,proto3" json:"MCJJKBHLAJN,omitempty"`
}

func (x *WolfBroGameInfo) Reset() {
	*x = WolfBroGameInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_WolfBroGameInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WolfBroGameInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WolfBroGameInfo) ProtoMessage() {}

func (x *WolfBroGameInfo) ProtoReflect() protoreflect.Message {
	mi := &file_WolfBroGameInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WolfBroGameInfo.ProtoReflect.Descriptor instead.
func (*WolfBroGameInfo) Descriptor() ([]byte, []int) {
	return file_WolfBroGameInfo_proto_rawDescGZIP(), []int{0}
}

func (x *WolfBroGameInfo) GetCHGGJGHBNBM() uint32 {
	if x != nil {
		return x.CHGGJGHBNBM
	}
	return 0
}

func (x *WolfBroGameInfo) GetMotion() *MotionInfo {
	if x != nil {
		return x.Motion
	}
	return nil
}

func (x *WolfBroGameInfo) GetLFEHNFOPJCL() bool {
	if x != nil {
		return x.LFEHNFOPJCL
	}
	return false
}

func (x *WolfBroGameInfo) GetMCJJKBHLAJN() []*Vector {
	if x != nil {
		return x.MCJJKBHLAJN
	}
	return nil
}

var File_WolfBroGameInfo_proto protoreflect.FileDescriptor

var file_WolfBroGameInfo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x57, 0x6f, 0x6c, 0x66, 0x42, 0x72, 0x6f, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x01, 0x0a, 0x0f, 0x57, 0x6f, 0x6c, 0x66,
	0x42, 0x72, 0x6f, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x43,
	0x48, 0x47, 0x47, 0x4a, 0x47, 0x48, 0x42, 0x4e, 0x42, 0x4d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x43, 0x48, 0x47, 0x47, 0x4a, 0x47, 0x48, 0x42, 0x4e, 0x42, 0x4d, 0x12, 0x23, 0x0a,
	0x06, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x46, 0x45, 0x48, 0x4e, 0x46, 0x4f, 0x50, 0x4a, 0x43,
	0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4c, 0x46, 0x45, 0x48, 0x4e, 0x46, 0x4f,
	0x50, 0x4a, 0x43, 0x4c, 0x12, 0x29, 0x0a, 0x0b, 0x4d, 0x43, 0x4a, 0x4a, 0x4b, 0x42, 0x48, 0x4c,
	0x41, 0x4a, 0x4e, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x0b, 0x4d, 0x43, 0x4a, 0x4a, 0x4b, 0x42, 0x48, 0x4c, 0x41, 0x4a, 0x4e, 0x42,
	0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67,
	0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_WolfBroGameInfo_proto_rawDescOnce sync.Once
	file_WolfBroGameInfo_proto_rawDescData = file_WolfBroGameInfo_proto_rawDesc
)

func file_WolfBroGameInfo_proto_rawDescGZIP() []byte {
	file_WolfBroGameInfo_proto_rawDescOnce.Do(func() {
		file_WolfBroGameInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_WolfBroGameInfo_proto_rawDescData)
	})
	return file_WolfBroGameInfo_proto_rawDescData
}

var file_WolfBroGameInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_WolfBroGameInfo_proto_goTypes = []interface{}{
	(*WolfBroGameInfo)(nil), // 0: WolfBroGameInfo
	(*MotionInfo)(nil),      // 1: MotionInfo
	(*Vector)(nil),          // 2: Vector
}
var file_WolfBroGameInfo_proto_depIdxs = []int32{
	1, // 0: WolfBroGameInfo.motion:type_name -> MotionInfo
	2, // 1: WolfBroGameInfo.MCJJKBHLAJN:type_name -> Vector
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_WolfBroGameInfo_proto_init() }
func file_WolfBroGameInfo_proto_init() {
	if File_WolfBroGameInfo_proto != nil {
		return
	}
	file_Vector_proto_init()
	file_MotionInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_WolfBroGameInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WolfBroGameInfo); i {
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
			RawDescriptor: file_WolfBroGameInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_WolfBroGameInfo_proto_goTypes,
		DependencyIndexes: file_WolfBroGameInfo_proto_depIdxs,
		MessageInfos:      file_WolfBroGameInfo_proto_msgTypes,
	}.Build()
	File_WolfBroGameInfo_proto = out.File
	file_WolfBroGameInfo_proto_rawDesc = nil
	file_WolfBroGameInfo_proto_goTypes = nil
	file_WolfBroGameInfo_proto_depIdxs = nil
}

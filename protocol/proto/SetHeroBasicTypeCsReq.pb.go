// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SetHeroBasicTypeCsReq.proto

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

type SetHeroBasicTypeCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BasicType HeroBasicType `protobuf:"varint,11,opt,name=basic_type,json=basicType,proto3,enum=HeroBasicType" json:"basic_type,omitempty"`
}

func (x *SetHeroBasicTypeCsReq) Reset() {
	*x = SetHeroBasicTypeCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SetHeroBasicTypeCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetHeroBasicTypeCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetHeroBasicTypeCsReq) ProtoMessage() {}

func (x *SetHeroBasicTypeCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_SetHeroBasicTypeCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetHeroBasicTypeCsReq.ProtoReflect.Descriptor instead.
func (*SetHeroBasicTypeCsReq) Descriptor() ([]byte, []int) {
	return file_SetHeroBasicTypeCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *SetHeroBasicTypeCsReq) GetBasicType() HeroBasicType {
	if x != nil {
		return x.BasicType
	}
	return HeroBasicType_None
}

var File_SetHeroBasicTypeCsReq_proto protoreflect.FileDescriptor

var file_SetHeroBasicTypeCsReq_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x53, 0x65, 0x74, 0x48, 0x65, 0x72, 0x6f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79,
	0x70, 0x65, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x48,
	0x65, 0x72, 0x6f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x46, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x48, 0x65, 0x72, 0x6f, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x2d, 0x0a, 0x0a, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0e, 0x2e, 0x48, 0x65, 0x72, 0x6f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x09, 0x62, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SetHeroBasicTypeCsReq_proto_rawDescOnce sync.Once
	file_SetHeroBasicTypeCsReq_proto_rawDescData = file_SetHeroBasicTypeCsReq_proto_rawDesc
)

func file_SetHeroBasicTypeCsReq_proto_rawDescGZIP() []byte {
	file_SetHeroBasicTypeCsReq_proto_rawDescOnce.Do(func() {
		file_SetHeroBasicTypeCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_SetHeroBasicTypeCsReq_proto_rawDescData)
	})
	return file_SetHeroBasicTypeCsReq_proto_rawDescData
}

var file_SetHeroBasicTypeCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SetHeroBasicTypeCsReq_proto_goTypes = []interface{}{
	(*SetHeroBasicTypeCsReq)(nil), // 0: SetHeroBasicTypeCsReq
	(HeroBasicType)(0),            // 1: HeroBasicType
}
var file_SetHeroBasicTypeCsReq_proto_depIdxs = []int32{
	1, // 0: SetHeroBasicTypeCsReq.basic_type:type_name -> HeroBasicType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SetHeroBasicTypeCsReq_proto_init() }
func file_SetHeroBasicTypeCsReq_proto_init() {
	if File_SetHeroBasicTypeCsReq_proto != nil {
		return
	}
	file_HeroBasicType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SetHeroBasicTypeCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetHeroBasicTypeCsReq); i {
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
			RawDescriptor: file_SetHeroBasicTypeCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SetHeroBasicTypeCsReq_proto_goTypes,
		DependencyIndexes: file_SetHeroBasicTypeCsReq_proto_depIdxs,
		MessageInfos:      file_SetHeroBasicTypeCsReq_proto_msgTypes,
	}.Build()
	File_SetHeroBasicTypeCsReq_proto = out.File
	file_SetHeroBasicTypeCsReq_proto_rawDesc = nil
	file_SetHeroBasicTypeCsReq_proto_goTypes = nil
	file_SetHeroBasicTypeCsReq_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ACAJPHPNKAE.proto

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

type ACAJPHPNKAE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IKOHEMFHCJO *CNMOPPANLGC   `protobuf:"bytes,8,opt,name=IKOHEMFHCJO,proto3" json:"IKOHEMFHCJO,omitempty"`
	NGGJLDOLBFO *DJCBKLFAGDI   `protobuf:"bytes,11,opt,name=NGGJLDOLBFO,proto3" json:"NGGJLDOLBFO,omitempty"`
	CMDJFBFCEDI []*CNMOPPANLGC `protobuf:"bytes,14,rep,name=CMDJFBFCEDI,proto3" json:"CMDJFBFCEDI,omitempty"`
}

func (x *ACAJPHPNKAE) Reset() {
	*x = ACAJPHPNKAE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ACAJPHPNKAE_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACAJPHPNKAE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACAJPHPNKAE) ProtoMessage() {}

func (x *ACAJPHPNKAE) ProtoReflect() protoreflect.Message {
	mi := &file_ACAJPHPNKAE_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACAJPHPNKAE.ProtoReflect.Descriptor instead.
func (*ACAJPHPNKAE) Descriptor() ([]byte, []int) {
	return file_ACAJPHPNKAE_proto_rawDescGZIP(), []int{0}
}

func (x *ACAJPHPNKAE) GetIKOHEMFHCJO() *CNMOPPANLGC {
	if x != nil {
		return x.IKOHEMFHCJO
	}
	return nil
}

func (x *ACAJPHPNKAE) GetNGGJLDOLBFO() *DJCBKLFAGDI {
	if x != nil {
		return x.NGGJLDOLBFO
	}
	return nil
}

func (x *ACAJPHPNKAE) GetCMDJFBFCEDI() []*CNMOPPANLGC {
	if x != nil {
		return x.CMDJFBFCEDI
	}
	return nil
}

var File_ACAJPHPNKAE_proto protoreflect.FileDescriptor

var file_ACAJPHPNKAE_proto_rawDesc = []byte{
	0x0a, 0x11, 0x41, 0x43, 0x41, 0x4a, 0x50, 0x48, 0x50, 0x4e, 0x4b, 0x41, 0x45, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x44, 0x4a, 0x43, 0x42, 0x4b, 0x4c, 0x46, 0x41, 0x47, 0x44, 0x49,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x43, 0x4e, 0x4d, 0x4f, 0x50, 0x50, 0x41, 0x4e,
	0x4c, 0x47, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a, 0x0b, 0x41, 0x43,
	0x41, 0x4a, 0x50, 0x48, 0x50, 0x4e, 0x4b, 0x41, 0x45, 0x12, 0x2e, 0x0a, 0x0b, 0x49, 0x4b, 0x4f,
	0x48, 0x45, 0x4d, 0x46, 0x48, 0x43, 0x4a, 0x4f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x43, 0x4e, 0x4d, 0x4f, 0x50, 0x50, 0x41, 0x4e, 0x4c, 0x47, 0x43, 0x52, 0x0b, 0x49, 0x4b,
	0x4f, 0x48, 0x45, 0x4d, 0x46, 0x48, 0x43, 0x4a, 0x4f, 0x12, 0x2e, 0x0a, 0x0b, 0x4e, 0x47, 0x47,
	0x4a, 0x4c, 0x44, 0x4f, 0x4c, 0x42, 0x46, 0x4f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x44, 0x4a, 0x43, 0x42, 0x4b, 0x4c, 0x46, 0x41, 0x47, 0x44, 0x49, 0x52, 0x0b, 0x4e, 0x47,
	0x47, 0x4a, 0x4c, 0x44, 0x4f, 0x4c, 0x42, 0x46, 0x4f, 0x12, 0x2e, 0x0a, 0x0b, 0x43, 0x4d, 0x44,
	0x4a, 0x46, 0x42, 0x46, 0x43, 0x45, 0x44, 0x49, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x43, 0x4e, 0x4d, 0x4f, 0x50, 0x50, 0x41, 0x4e, 0x4c, 0x47, 0x43, 0x52, 0x0b, 0x43, 0x4d,
	0x44, 0x4a, 0x46, 0x42, 0x46, 0x43, 0x45, 0x44, 0x49, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e,
	0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ACAJPHPNKAE_proto_rawDescOnce sync.Once
	file_ACAJPHPNKAE_proto_rawDescData = file_ACAJPHPNKAE_proto_rawDesc
)

func file_ACAJPHPNKAE_proto_rawDescGZIP() []byte {
	file_ACAJPHPNKAE_proto_rawDescOnce.Do(func() {
		file_ACAJPHPNKAE_proto_rawDescData = protoimpl.X.CompressGZIP(file_ACAJPHPNKAE_proto_rawDescData)
	})
	return file_ACAJPHPNKAE_proto_rawDescData
}

var file_ACAJPHPNKAE_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ACAJPHPNKAE_proto_goTypes = []interface{}{
	(*ACAJPHPNKAE)(nil), // 0: ACAJPHPNKAE
	(*CNMOPPANLGC)(nil), // 1: CNMOPPANLGC
	(*DJCBKLFAGDI)(nil), // 2: DJCBKLFAGDI
}
var file_ACAJPHPNKAE_proto_depIdxs = []int32{
	1, // 0: ACAJPHPNKAE.IKOHEMFHCJO:type_name -> CNMOPPANLGC
	2, // 1: ACAJPHPNKAE.NGGJLDOLBFO:type_name -> DJCBKLFAGDI
	1, // 2: ACAJPHPNKAE.CMDJFBFCEDI:type_name -> CNMOPPANLGC
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ACAJPHPNKAE_proto_init() }
func file_ACAJPHPNKAE_proto_init() {
	if File_ACAJPHPNKAE_proto != nil {
		return
	}
	file_DJCBKLFAGDI_proto_init()
	file_CNMOPPANLGC_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ACAJPHPNKAE_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ACAJPHPNKAE); i {
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
			RawDescriptor: file_ACAJPHPNKAE_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ACAJPHPNKAE_proto_goTypes,
		DependencyIndexes: file_ACAJPHPNKAE_proto_depIdxs,
		MessageInfos:      file_ACAJPHPNKAE_proto_msgTypes,
	}.Build()
	File_ACAJPHPNKAE_proto = out.File
	file_ACAJPHPNKAE_proto_rawDesc = nil
	file_ACAJPHPNKAE_proto_goTypes = nil
	file_ACAJPHPNKAE_proto_depIdxs = nil
}

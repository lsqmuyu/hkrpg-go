// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: NNLHJAIPCBO.proto

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

type NNLHJAIPCBO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ANNJPLFCBMB uint32 `protobuf:"varint,15,opt,name=ANNJPLFCBMB,proto3" json:"ANNJPLFCBMB,omitempty"`
	PDBBAEIEJJB uint32 `protobuf:"varint,5,opt,name=PDBBAEIEJJB,proto3" json:"PDBBAEIEJJB,omitempty"`
	OIHDEPMDJBI uint32 `protobuf:"varint,8,opt,name=OIHDEPMDJBI,proto3" json:"OIHDEPMDJBI,omitempty"`
	PanelId     uint32 `protobuf:"varint,4,opt,name=panel_id,json=panelId,proto3" json:"panel_id,omitempty"`
}

func (x *NNLHJAIPCBO) Reset() {
	*x = NNLHJAIPCBO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NNLHJAIPCBO_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NNLHJAIPCBO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NNLHJAIPCBO) ProtoMessage() {}

func (x *NNLHJAIPCBO) ProtoReflect() protoreflect.Message {
	mi := &file_NNLHJAIPCBO_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NNLHJAIPCBO.ProtoReflect.Descriptor instead.
func (*NNLHJAIPCBO) Descriptor() ([]byte, []int) {
	return file_NNLHJAIPCBO_proto_rawDescGZIP(), []int{0}
}

func (x *NNLHJAIPCBO) GetANNJPLFCBMB() uint32 {
	if x != nil {
		return x.ANNJPLFCBMB
	}
	return 0
}

func (x *NNLHJAIPCBO) GetPDBBAEIEJJB() uint32 {
	if x != nil {
		return x.PDBBAEIEJJB
	}
	return 0
}

func (x *NNLHJAIPCBO) GetOIHDEPMDJBI() uint32 {
	if x != nil {
		return x.OIHDEPMDJBI
	}
	return 0
}

func (x *NNLHJAIPCBO) GetPanelId() uint32 {
	if x != nil {
		return x.PanelId
	}
	return 0
}

var File_NNLHJAIPCBO_proto protoreflect.FileDescriptor

var file_NNLHJAIPCBO_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4e, 0x4e, 0x4c, 0x48, 0x4a, 0x41, 0x49, 0x50, 0x43, 0x42, 0x4f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x01, 0x0a, 0x0b, 0x4e, 0x4e, 0x4c, 0x48, 0x4a, 0x41, 0x49, 0x50,
	0x43, 0x42, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x4e, 0x4e, 0x4a, 0x50, 0x4c, 0x46, 0x43, 0x42,
	0x4d, 0x42, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x41, 0x4e, 0x4e, 0x4a, 0x50, 0x4c,
	0x46, 0x43, 0x42, 0x4d, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x44, 0x42, 0x42, 0x41, 0x45, 0x49,
	0x45, 0x4a, 0x4a, 0x42, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x50, 0x44, 0x42, 0x42,
	0x41, 0x45, 0x49, 0x45, 0x4a, 0x4a, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x49, 0x48, 0x44, 0x45,
	0x50, 0x4d, 0x44, 0x4a, 0x42, 0x49, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x49,
	0x48, 0x44, 0x45, 0x50, 0x4d, 0x44, 0x4a, 0x42, 0x49, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x6e,
	0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x61, 0x6e,
	0x65, 0x6c, 0x49, 0x64, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_NNLHJAIPCBO_proto_rawDescOnce sync.Once
	file_NNLHJAIPCBO_proto_rawDescData = file_NNLHJAIPCBO_proto_rawDesc
)

func file_NNLHJAIPCBO_proto_rawDescGZIP() []byte {
	file_NNLHJAIPCBO_proto_rawDescOnce.Do(func() {
		file_NNLHJAIPCBO_proto_rawDescData = protoimpl.X.CompressGZIP(file_NNLHJAIPCBO_proto_rawDescData)
	})
	return file_NNLHJAIPCBO_proto_rawDescData
}

var file_NNLHJAIPCBO_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_NNLHJAIPCBO_proto_goTypes = []interface{}{
	(*NNLHJAIPCBO)(nil), // 0: NNLHJAIPCBO
}
var file_NNLHJAIPCBO_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_NNLHJAIPCBO_proto_init() }
func file_NNLHJAIPCBO_proto_init() {
	if File_NNLHJAIPCBO_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_NNLHJAIPCBO_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NNLHJAIPCBO); i {
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
			RawDescriptor: file_NNLHJAIPCBO_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_NNLHJAIPCBO_proto_goTypes,
		DependencyIndexes: file_NNLHJAIPCBO_proto_depIdxs,
		MessageInfos:      file_NNLHJAIPCBO_proto_msgTypes,
	}.Build()
	File_NNLHJAIPCBO_proto = out.File
	file_NNLHJAIPCBO_proto_rawDesc = nil
	file_NNLHJAIPCBO_proto_goTypes = nil
	file_NNLHJAIPCBO_proto_depIdxs = nil
}

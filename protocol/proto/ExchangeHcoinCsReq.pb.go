// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ExchangeHcoinCsReq.proto

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

type ExchangeHcoinCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num uint32 `protobuf:"varint,11,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *ExchangeHcoinCsReq) Reset() {
	*x = ExchangeHcoinCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ExchangeHcoinCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeHcoinCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeHcoinCsReq) ProtoMessage() {}

func (x *ExchangeHcoinCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_ExchangeHcoinCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeHcoinCsReq.ProtoReflect.Descriptor instead.
func (*ExchangeHcoinCsReq) Descriptor() ([]byte, []int) {
	return file_ExchangeHcoinCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *ExchangeHcoinCsReq) GetNum() uint32 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File_ExchangeHcoinCsReq_proto protoreflect.FileDescriptor

var file_ExchangeHcoinCsReq_proto_rawDesc = []byte{
	0x0a, 0x18, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x63, 0x6f, 0x69, 0x6e, 0x43,
	0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x12, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x63, 0x6f, 0x69, 0x6e, 0x43, 0x73, 0x52, 0x65, 0x71,
	0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6e,
	0x75, 0x6d, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ExchangeHcoinCsReq_proto_rawDescOnce sync.Once
	file_ExchangeHcoinCsReq_proto_rawDescData = file_ExchangeHcoinCsReq_proto_rawDesc
)

func file_ExchangeHcoinCsReq_proto_rawDescGZIP() []byte {
	file_ExchangeHcoinCsReq_proto_rawDescOnce.Do(func() {
		file_ExchangeHcoinCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_ExchangeHcoinCsReq_proto_rawDescData)
	})
	return file_ExchangeHcoinCsReq_proto_rawDescData
}

var file_ExchangeHcoinCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ExchangeHcoinCsReq_proto_goTypes = []interface{}{
	(*ExchangeHcoinCsReq)(nil), // 0: ExchangeHcoinCsReq
}
var file_ExchangeHcoinCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ExchangeHcoinCsReq_proto_init() }
func file_ExchangeHcoinCsReq_proto_init() {
	if File_ExchangeHcoinCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ExchangeHcoinCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeHcoinCsReq); i {
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
			RawDescriptor: file_ExchangeHcoinCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ExchangeHcoinCsReq_proto_goTypes,
		DependencyIndexes: file_ExchangeHcoinCsReq_proto_depIdxs,
		MessageInfos:      file_ExchangeHcoinCsReq_proto_msgTypes,
	}.Build()
	File_ExchangeHcoinCsReq_proto = out.File
	file_ExchangeHcoinCsReq_proto_rawDesc = nil
	file_ExchangeHcoinCsReq_proto_goTypes = nil
	file_ExchangeHcoinCsReq_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: FinishCurTurnScRsp.proto

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

type FinishCurTurnScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EJFIBEKFPMO uint32 `protobuf:"varint,12,opt,name=EJFIBEKFPMO,proto3" json:"EJFIBEKFPMO,omitempty"`
	Retcode     uint32 `protobuf:"varint,14,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *FinishCurTurnScRsp) Reset() {
	*x = FinishCurTurnScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FinishCurTurnScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishCurTurnScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishCurTurnScRsp) ProtoMessage() {}

func (x *FinishCurTurnScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_FinishCurTurnScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishCurTurnScRsp.ProtoReflect.Descriptor instead.
func (*FinishCurTurnScRsp) Descriptor() ([]byte, []int) {
	return file_FinishCurTurnScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *FinishCurTurnScRsp) GetEJFIBEKFPMO() uint32 {
	if x != nil {
		return x.EJFIBEKFPMO
	}
	return 0
}

func (x *FinishCurTurnScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_FinishCurTurnScRsp_proto protoreflect.FileDescriptor

var file_FinishCurTurnScRsp_proto_rawDesc = []byte{
	0x0a, 0x18, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x43, 0x75, 0x72, 0x54, 0x75, 0x72, 0x6e, 0x53,
	0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x12, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x43, 0x75, 0x72, 0x54, 0x75, 0x72, 0x6e, 0x53, 0x63, 0x52, 0x73, 0x70,
	0x12, 0x20, 0x0a, 0x0b, 0x45, 0x4a, 0x46, 0x49, 0x42, 0x45, 0x4b, 0x46, 0x50, 0x4d, 0x4f, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x4a, 0x46, 0x49, 0x42, 0x45, 0x4b, 0x46, 0x50,
	0x4d, 0x4f, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x28, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FinishCurTurnScRsp_proto_rawDescOnce sync.Once
	file_FinishCurTurnScRsp_proto_rawDescData = file_FinishCurTurnScRsp_proto_rawDesc
)

func file_FinishCurTurnScRsp_proto_rawDescGZIP() []byte {
	file_FinishCurTurnScRsp_proto_rawDescOnce.Do(func() {
		file_FinishCurTurnScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_FinishCurTurnScRsp_proto_rawDescData)
	})
	return file_FinishCurTurnScRsp_proto_rawDescData
}

var file_FinishCurTurnScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FinishCurTurnScRsp_proto_goTypes = []interface{}{
	(*FinishCurTurnScRsp)(nil), // 0: FinishCurTurnScRsp
}
var file_FinishCurTurnScRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_FinishCurTurnScRsp_proto_init() }
func file_FinishCurTurnScRsp_proto_init() {
	if File_FinishCurTurnScRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_FinishCurTurnScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishCurTurnScRsp); i {
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
			RawDescriptor: file_FinishCurTurnScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FinishCurTurnScRsp_proto_goTypes,
		DependencyIndexes: file_FinishCurTurnScRsp_proto_depIdxs,
		MessageInfos:      file_FinishCurTurnScRsp_proto_msgTypes,
	}.Build()
	File_FinishCurTurnScRsp_proto = out.File
	file_FinishCurTurnScRsp_proto_rawDesc = nil
	file_FinishCurTurnScRsp_proto_goTypes = nil
	file_FinishCurTurnScRsp_proto_depIdxs = nil
}

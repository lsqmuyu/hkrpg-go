// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: FinishEmotionDialoguePerformanceCsReq.proto

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

type FinishEmotionDialoguePerformanceCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GBLGDPPIFPE uint32 `protobuf:"varint,7,opt,name=GBLGDPPIFPE,proto3" json:"GBLGDPPIFPE,omitempty"`
	FDODEHMNPJP uint32 `protobuf:"varint,13,opt,name=FDODEHMNPJP,proto3" json:"FDODEHMNPJP,omitempty"`
	BPDKOPMKJMM uint32 `protobuf:"varint,10,opt,name=BPDKOPMKJMM,proto3" json:"BPDKOPMKJMM,omitempty"`
}

func (x *FinishEmotionDialoguePerformanceCsReq) Reset() {
	*x = FinishEmotionDialoguePerformanceCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FinishEmotionDialoguePerformanceCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishEmotionDialoguePerformanceCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishEmotionDialoguePerformanceCsReq) ProtoMessage() {}

func (x *FinishEmotionDialoguePerformanceCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_FinishEmotionDialoguePerformanceCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishEmotionDialoguePerformanceCsReq.ProtoReflect.Descriptor instead.
func (*FinishEmotionDialoguePerformanceCsReq) Descriptor() ([]byte, []int) {
	return file_FinishEmotionDialoguePerformanceCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *FinishEmotionDialoguePerformanceCsReq) GetGBLGDPPIFPE() uint32 {
	if x != nil {
		return x.GBLGDPPIFPE
	}
	return 0
}

func (x *FinishEmotionDialoguePerformanceCsReq) GetFDODEHMNPJP() uint32 {
	if x != nil {
		return x.FDODEHMNPJP
	}
	return 0
}

func (x *FinishEmotionDialoguePerformanceCsReq) GetBPDKOPMKJMM() uint32 {
	if x != nil {
		return x.BPDKOPMKJMM
	}
	return 0
}

var File_FinishEmotionDialoguePerformanceCsReq_proto protoreflect.FileDescriptor

var file_FinishEmotionDialoguePerformanceCsReq_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x45, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e,
	0x63, 0x65, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01,
	0x0a, 0x25, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x45, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e,
	0x63, 0x65, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x42, 0x4c, 0x47, 0x44,
	0x50, 0x50, 0x49, 0x46, 0x50, 0x45, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x47, 0x42,
	0x4c, 0x47, 0x44, 0x50, 0x50, 0x49, 0x46, 0x50, 0x45, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x44, 0x4f,
	0x44, 0x45, 0x48, 0x4d, 0x4e, 0x50, 0x4a, 0x50, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x46, 0x44, 0x4f, 0x44, 0x45, 0x48, 0x4d, 0x4e, 0x50, 0x4a, 0x50, 0x12, 0x20, 0x0a, 0x0b, 0x42,
	0x50, 0x44, 0x4b, 0x4f, 0x50, 0x4d, 0x4b, 0x4a, 0x4d, 0x4d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x42, 0x50, 0x44, 0x4b, 0x4f, 0x50, 0x4d, 0x4b, 0x4a, 0x4d, 0x4d, 0x42, 0x28, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FinishEmotionDialoguePerformanceCsReq_proto_rawDescOnce sync.Once
	file_FinishEmotionDialoguePerformanceCsReq_proto_rawDescData = file_FinishEmotionDialoguePerformanceCsReq_proto_rawDesc
)

func file_FinishEmotionDialoguePerformanceCsReq_proto_rawDescGZIP() []byte {
	file_FinishEmotionDialoguePerformanceCsReq_proto_rawDescOnce.Do(func() {
		file_FinishEmotionDialoguePerformanceCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_FinishEmotionDialoguePerformanceCsReq_proto_rawDescData)
	})
	return file_FinishEmotionDialoguePerformanceCsReq_proto_rawDescData
}

var file_FinishEmotionDialoguePerformanceCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FinishEmotionDialoguePerformanceCsReq_proto_goTypes = []interface{}{
	(*FinishEmotionDialoguePerformanceCsReq)(nil), // 0: FinishEmotionDialoguePerformanceCsReq
}
var file_FinishEmotionDialoguePerformanceCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_FinishEmotionDialoguePerformanceCsReq_proto_init() }
func file_FinishEmotionDialoguePerformanceCsReq_proto_init() {
	if File_FinishEmotionDialoguePerformanceCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_FinishEmotionDialoguePerformanceCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishEmotionDialoguePerformanceCsReq); i {
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
			RawDescriptor: file_FinishEmotionDialoguePerformanceCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FinishEmotionDialoguePerformanceCsReq_proto_goTypes,
		DependencyIndexes: file_FinishEmotionDialoguePerformanceCsReq_proto_depIdxs,
		MessageInfos:      file_FinishEmotionDialoguePerformanceCsReq_proto_msgTypes,
	}.Build()
	File_FinishEmotionDialoguePerformanceCsReq_proto = out.File
	file_FinishEmotionDialoguePerformanceCsReq_proto_rawDesc = nil
	file_FinishEmotionDialoguePerformanceCsReq_proto_goTypes = nil
	file_FinishEmotionDialoguePerformanceCsReq_proto_depIdxs = nil
}

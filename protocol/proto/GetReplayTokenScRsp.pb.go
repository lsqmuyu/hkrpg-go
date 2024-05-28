// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetReplayTokenScRsp.proto

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

type GetReplayTokenScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NFFLNDEHLFJ string     `protobuf:"bytes,2,opt,name=NFFLNDEHLFJ,proto3" json:"NFFLNDEHLFJ,omitempty"`
	ReplayType  ReplayType `protobuf:"varint,10,opt,name=replay_type,json=replayType,proto3,enum=ReplayType" json:"replay_type,omitempty"`
	StageId     uint32     `protobuf:"varint,1,opt,name=stage_id,json=stageId,proto3" json:"stage_id,omitempty"`
	Retcode     uint32     `protobuf:"varint,6,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Token       string     `protobuf:"bytes,8,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetReplayTokenScRsp) Reset() {
	*x = GetReplayTokenScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetReplayTokenScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReplayTokenScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReplayTokenScRsp) ProtoMessage() {}

func (x *GetReplayTokenScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetReplayTokenScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReplayTokenScRsp.ProtoReflect.Descriptor instead.
func (*GetReplayTokenScRsp) Descriptor() ([]byte, []int) {
	return file_GetReplayTokenScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetReplayTokenScRsp) GetNFFLNDEHLFJ() string {
	if x != nil {
		return x.NFFLNDEHLFJ
	}
	return ""
}

func (x *GetReplayTokenScRsp) GetReplayType() ReplayType {
	if x != nil {
		return x.ReplayType
	}
	return ReplayType_REPLAY_TYPE_NONE
}

func (x *GetReplayTokenScRsp) GetStageId() uint32 {
	if x != nil {
		return x.StageId
	}
	return 0
}

func (x *GetReplayTokenScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetReplayTokenScRsp) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_GetReplayTokenScRsp_proto protoreflect.FileDescriptor

var file_GetReplayTokenScRsp_proto_rawDesc = []byte{
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x52, 0x65, 0x70,
	0x6c, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x46, 0x46, 0x4c, 0x4e, 0x44, 0x45,
	0x48, 0x4c, 0x46, 0x4a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4e, 0x46, 0x46, 0x4c,
	0x4e, 0x44, 0x45, 0x48, 0x4c, 0x46, 0x4a, 0x12, 0x2c, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x61,
	0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x52,
	0x65, 0x70, 0x6c, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x61,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45,
	0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_GetReplayTokenScRsp_proto_rawDescOnce sync.Once
	file_GetReplayTokenScRsp_proto_rawDescData = file_GetReplayTokenScRsp_proto_rawDesc
)

func file_GetReplayTokenScRsp_proto_rawDescGZIP() []byte {
	file_GetReplayTokenScRsp_proto_rawDescOnce.Do(func() {
		file_GetReplayTokenScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetReplayTokenScRsp_proto_rawDescData)
	})
	return file_GetReplayTokenScRsp_proto_rawDescData
}

var file_GetReplayTokenScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetReplayTokenScRsp_proto_goTypes = []interface{}{
	(*GetReplayTokenScRsp)(nil), // 0: GetReplayTokenScRsp
	(ReplayType)(0),             // 1: ReplayType
}
var file_GetReplayTokenScRsp_proto_depIdxs = []int32{
	1, // 0: GetReplayTokenScRsp.replay_type:type_name -> ReplayType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetReplayTokenScRsp_proto_init() }
func file_GetReplayTokenScRsp_proto_init() {
	if File_GetReplayTokenScRsp_proto != nil {
		return
	}
	file_ReplayType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetReplayTokenScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReplayTokenScRsp); i {
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
			RawDescriptor: file_GetReplayTokenScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetReplayTokenScRsp_proto_goTypes,
		DependencyIndexes: file_GetReplayTokenScRsp_proto_depIdxs,
		MessageInfos:      file_GetReplayTokenScRsp_proto_msgTypes,
	}.Build()
	File_GetReplayTokenScRsp_proto = out.File
	file_GetReplayTokenScRsp_proto_rawDesc = nil
	file_GetReplayTokenScRsp_proto_goTypes = nil
	file_GetReplayTokenScRsp_proto_depIdxs = nil
}

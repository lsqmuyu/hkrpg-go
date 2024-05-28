// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetCurBattleInfoScRsp.proto

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

type GetCurBattleInfoScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HMIJFMDBBEO uint32                  `protobuf:"varint,15,opt,name=HMIJFMDBBEO,proto3" json:"HMIJFMDBBEO,omitempty"`
	OJLHDMBAGAC *AetherDivideBattleInfo `protobuf:"bytes,11,opt,name=OJLHDMBAGAC,proto3" json:"OJLHDMBAGAC,omitempty"`
	Retcode     uint32                  `protobuf:"varint,9,opt,name=retcode,proto3" json:"retcode,omitempty"`
	IICMLFOBDFH BattleEndStatus         `protobuf:"varint,6,opt,name=IICMLFOBDFH,proto3,enum=BattleEndStatus" json:"IICMLFOBDFH,omitempty"`
	BattleInfo  *SceneBattleInfo        `protobuf:"bytes,7,opt,name=battle_info,json=battleInfo,proto3" json:"battle_info,omitempty"`
}

func (x *GetCurBattleInfoScRsp) Reset() {
	*x = GetCurBattleInfoScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetCurBattleInfoScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurBattleInfoScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurBattleInfoScRsp) ProtoMessage() {}

func (x *GetCurBattleInfoScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetCurBattleInfoScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurBattleInfoScRsp.ProtoReflect.Descriptor instead.
func (*GetCurBattleInfoScRsp) Descriptor() ([]byte, []int) {
	return file_GetCurBattleInfoScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetCurBattleInfoScRsp) GetHMIJFMDBBEO() uint32 {
	if x != nil {
		return x.HMIJFMDBBEO
	}
	return 0
}

func (x *GetCurBattleInfoScRsp) GetOJLHDMBAGAC() *AetherDivideBattleInfo {
	if x != nil {
		return x.OJLHDMBAGAC
	}
	return nil
}

func (x *GetCurBattleInfoScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetCurBattleInfoScRsp) GetIICMLFOBDFH() BattleEndStatus {
	if x != nil {
		return x.IICMLFOBDFH
	}
	return BattleEndStatus_BATTLE_END_NONE
}

func (x *GetCurBattleInfoScRsp) GetBattleInfo() *SceneBattleInfo {
	if x != nil {
		return x.BattleInfo
	}
	return nil
}

var File_GetCurBattleInfoScRsp_proto protoreflect.FileDescriptor

var file_GetCurBattleInfoScRsp_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x42,
	0x61, 0x74, 0x74, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x41, 0x65, 0x74,
	0x68, 0x65, 0x72, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x01, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x43, 0x75, 0x72, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63,
	0x52, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x4d, 0x49, 0x4a, 0x46, 0x4d, 0x44, 0x42, 0x42,
	0x45, 0x4f, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x48, 0x4d, 0x49, 0x4a, 0x46, 0x4d,
	0x44, 0x42, 0x42, 0x45, 0x4f, 0x12, 0x39, 0x0a, 0x0b, 0x4f, 0x4a, 0x4c, 0x48, 0x44, 0x4d, 0x42,
	0x41, 0x47, 0x41, 0x43, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x41, 0x65, 0x74,
	0x68, 0x65, 0x72, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x4f, 0x4a, 0x4c, 0x48, 0x44, 0x4d, 0x42, 0x41, 0x47, 0x41, 0x43,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x32, 0x0a, 0x0b, 0x49, 0x49,
	0x43, 0x4d, 0x4c, 0x46, 0x4f, 0x42, 0x44, 0x46, 0x48, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x10, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x45, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x0b, 0x49, 0x49, 0x43, 0x4d, 0x4c, 0x46, 0x4f, 0x42, 0x44, 0x46, 0x48, 0x12, 0x31,
	0x0a, 0x0b, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_GetCurBattleInfoScRsp_proto_rawDescOnce sync.Once
	file_GetCurBattleInfoScRsp_proto_rawDescData = file_GetCurBattleInfoScRsp_proto_rawDesc
)

func file_GetCurBattleInfoScRsp_proto_rawDescGZIP() []byte {
	file_GetCurBattleInfoScRsp_proto_rawDescOnce.Do(func() {
		file_GetCurBattleInfoScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetCurBattleInfoScRsp_proto_rawDescData)
	})
	return file_GetCurBattleInfoScRsp_proto_rawDescData
}

var file_GetCurBattleInfoScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetCurBattleInfoScRsp_proto_goTypes = []interface{}{
	(*GetCurBattleInfoScRsp)(nil),  // 0: GetCurBattleInfoScRsp
	(*AetherDivideBattleInfo)(nil), // 1: AetherDivideBattleInfo
	(BattleEndStatus)(0),           // 2: BattleEndStatus
	(*SceneBattleInfo)(nil),        // 3: SceneBattleInfo
}
var file_GetCurBattleInfoScRsp_proto_depIdxs = []int32{
	1, // 0: GetCurBattleInfoScRsp.OJLHDMBAGAC:type_name -> AetherDivideBattleInfo
	2, // 1: GetCurBattleInfoScRsp.IICMLFOBDFH:type_name -> BattleEndStatus
	3, // 2: GetCurBattleInfoScRsp.battle_info:type_name -> SceneBattleInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_GetCurBattleInfoScRsp_proto_init() }
func file_GetCurBattleInfoScRsp_proto_init() {
	if File_GetCurBattleInfoScRsp_proto != nil {
		return
	}
	file_BattleEndStatus_proto_init()
	file_SceneBattleInfo_proto_init()
	file_AetherDivideBattleInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetCurBattleInfoScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurBattleInfoScRsp); i {
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
			RawDescriptor: file_GetCurBattleInfoScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetCurBattleInfoScRsp_proto_goTypes,
		DependencyIndexes: file_GetCurBattleInfoScRsp_proto_depIdxs,
		MessageInfos:      file_GetCurBattleInfoScRsp_proto_msgTypes,
	}.Build()
	File_GetCurBattleInfoScRsp_proto = out.File
	file_GetCurBattleInfoScRsp_proto_rawDesc = nil
	file_GetCurBattleInfoScRsp_proto_goTypes = nil
	file_GetCurBattleInfoScRsp_proto_depIdxs = nil
}

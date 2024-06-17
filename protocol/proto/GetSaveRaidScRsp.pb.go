// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetSaveRaidScRsp.proto

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

type GetSaveRaidScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorldLevel  uint32  `protobuf:"varint,6,opt,name=world_level,json=worldLevel,proto3" json:"world_level,omitempty"`
	Retcode     uint32  `protobuf:"varint,3,opt,name=retcode,proto3" json:"retcode,omitempty"`
	RaidList    []*Raid `protobuf:"bytes,10,rep,name=raid_list,json=raidList,proto3" json:"raid_list,omitempty"`
	RaidId      uint32  `protobuf:"varint,13,opt,name=raid_id,json=raidId,proto3" json:"raid_id,omitempty"`
	IIJPDKFDAKN bool    `protobuf:"varint,4,opt,name=IIJPDKFDAKN,proto3" json:"IIJPDKFDAKN,omitempty"`
}

func (x *GetSaveRaidScRsp) Reset() {
	*x = GetSaveRaidScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetSaveRaidScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSaveRaidScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSaveRaidScRsp) ProtoMessage() {}

func (x *GetSaveRaidScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetSaveRaidScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSaveRaidScRsp.ProtoReflect.Descriptor instead.
func (*GetSaveRaidScRsp) Descriptor() ([]byte, []int) {
	return file_GetSaveRaidScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetSaveRaidScRsp) GetWorldLevel() uint32 {
	if x != nil {
		return x.WorldLevel
	}
	return 0
}

func (x *GetSaveRaidScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetSaveRaidScRsp) GetRaidList() []*Raid {
	if x != nil {
		return x.RaidList
	}
	return nil
}

func (x *GetSaveRaidScRsp) GetRaidId() uint32 {
	if x != nil {
		return x.RaidId
	}
	return 0
}

func (x *GetSaveRaidScRsp) GetIIJPDKFDAKN() bool {
	if x != nil {
		return x.IIJPDKFDAKN
	}
	return false
}

var File_GetSaveRaidScRsp_proto protoreflect.FileDescriptor

var file_GetSaveRaidScRsp_proto_rawDesc = []byte{
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x61, 0x76, 0x65, 0x52, 0x61, 0x69, 0x64, 0x53, 0x63, 0x52,
	0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x52, 0x61, 0x69, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x61, 0x76, 0x65,
	0x52, 0x61, 0x69, 0x64, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x09, 0x72, 0x61, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x52, 0x61, 0x69, 0x64, 0x52, 0x08,
	0x72, 0x61, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x61, 0x69, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x72, 0x61, 0x69, 0x64, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x49, 0x4a, 0x50, 0x44, 0x4b, 0x46, 0x44, 0x41, 0x4b, 0x4e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x49, 0x49, 0x4a, 0x50, 0x44, 0x4b, 0x46, 0x44,
	0x41, 0x4b, 0x4e, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa,
	0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetSaveRaidScRsp_proto_rawDescOnce sync.Once
	file_GetSaveRaidScRsp_proto_rawDescData = file_GetSaveRaidScRsp_proto_rawDesc
)

func file_GetSaveRaidScRsp_proto_rawDescGZIP() []byte {
	file_GetSaveRaidScRsp_proto_rawDescOnce.Do(func() {
		file_GetSaveRaidScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetSaveRaidScRsp_proto_rawDescData)
	})
	return file_GetSaveRaidScRsp_proto_rawDescData
}

var file_GetSaveRaidScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetSaveRaidScRsp_proto_goTypes = []interface{}{
	(*GetSaveRaidScRsp)(nil), // 0: GetSaveRaidScRsp
	(*Raid)(nil),             // 1: Raid
}
var file_GetSaveRaidScRsp_proto_depIdxs = []int32{
	1, // 0: GetSaveRaidScRsp.raid_list:type_name -> Raid
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetSaveRaidScRsp_proto_init() }
func file_GetSaveRaidScRsp_proto_init() {
	if File_GetSaveRaidScRsp_proto != nil {
		return
	}
	file_Raid_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetSaveRaidScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSaveRaidScRsp); i {
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
			RawDescriptor: file_GetSaveRaidScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetSaveRaidScRsp_proto_goTypes,
		DependencyIndexes: file_GetSaveRaidScRsp_proto_depIdxs,
		MessageInfos:      file_GetSaveRaidScRsp_proto_msgTypes,
	}.Build()
	File_GetSaveRaidScRsp_proto = out.File
	file_GetSaveRaidScRsp_proto_rawDesc = nil
	file_GetSaveRaidScRsp_proto_goTypes = nil
	file_GetSaveRaidScRsp_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PlayerHeroBasicTypeInfo.proto

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

type PlayerHeroBasicTypeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IGBEBPJJJIH    uint32             `protobuf:"varint,11,opt,name=IGBEBPJJJIH,proto3" json:"IGBEBPJJJIH,omitempty"`
	SkillTreeList  []*AvatarSkillTree `protobuf:"bytes,12,rep,name=skill_tree_list,json=skillTreeList,proto3" json:"skill_tree_list,omitempty"`
	BasicType      HeroBasicType      `protobuf:"varint,14,opt,name=basic_type,json=basicType,proto3,enum=HeroBasicType" json:"basic_type,omitempty"`
	EquipRelicList []*EquipRelic      `protobuf:"bytes,7,rep,name=equip_relic_list,json=equipRelicList,proto3" json:"equip_relic_list,omitempty"`
	Rank           uint32             `protobuf:"varint,6,opt,name=rank,proto3" json:"rank,omitempty"`
}

func (x *PlayerHeroBasicTypeInfo) Reset() {
	*x = PlayerHeroBasicTypeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerHeroBasicTypeInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerHeroBasicTypeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerHeroBasicTypeInfo) ProtoMessage() {}

func (x *PlayerHeroBasicTypeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerHeroBasicTypeInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerHeroBasicTypeInfo.ProtoReflect.Descriptor instead.
func (*PlayerHeroBasicTypeInfo) Descriptor() ([]byte, []int) {
	return file_PlayerHeroBasicTypeInfo_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerHeroBasicTypeInfo) GetIGBEBPJJJIH() uint32 {
	if x != nil {
		return x.IGBEBPJJJIH
	}
	return 0
}

func (x *PlayerHeroBasicTypeInfo) GetSkillTreeList() []*AvatarSkillTree {
	if x != nil {
		return x.SkillTreeList
	}
	return nil
}

func (x *PlayerHeroBasicTypeInfo) GetBasicType() HeroBasicType {
	if x != nil {
		return x.BasicType
	}
	return HeroBasicType_None
}

func (x *PlayerHeroBasicTypeInfo) GetEquipRelicList() []*EquipRelic {
	if x != nil {
		return x.EquipRelicList
	}
	return nil
}

func (x *PlayerHeroBasicTypeInfo) GetRank() uint32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

var File_PlayerHeroBasicTypeInfo_proto protoreflect.FileDescriptor

var file_PlayerHeroBasicTypeInfo_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x48, 0x65, 0x72, 0x6f, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x54, 0x72, 0x65, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x48, 0x65, 0x72, 0x6f, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x45, 0x71, 0x75,
	0x69, 0x70, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x01,
	0x0a, 0x17, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x48, 0x65, 0x72, 0x6f, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x47, 0x42,
	0x45, 0x42, 0x50, 0x4a, 0x4a, 0x4a, 0x49, 0x48, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x49, 0x47, 0x42, 0x45, 0x42, 0x50, 0x4a, 0x4a, 0x4a, 0x49, 0x48, 0x12, 0x38, 0x0a, 0x0f, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x74, 0x72, 0x65, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x54, 0x72, 0x65, 0x65, 0x52, 0x0d, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x54, 0x72, 0x65,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0a, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x48, 0x65, 0x72, 0x6f,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x35, 0x0a, 0x10, 0x65, 0x71, 0x75, 0x69, 0x70, 0x5f, 0x72, 0x65,
	0x6c, 0x69, 0x63, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x45, 0x71, 0x75, 0x69, 0x70, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x52, 0x0e, 0x65, 0x71, 0x75,
	0x69, 0x70, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x61, 0x6e, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x42,
	0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67,
	0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_PlayerHeroBasicTypeInfo_proto_rawDescOnce sync.Once
	file_PlayerHeroBasicTypeInfo_proto_rawDescData = file_PlayerHeroBasicTypeInfo_proto_rawDesc
)

func file_PlayerHeroBasicTypeInfo_proto_rawDescGZIP() []byte {
	file_PlayerHeroBasicTypeInfo_proto_rawDescOnce.Do(func() {
		file_PlayerHeroBasicTypeInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerHeroBasicTypeInfo_proto_rawDescData)
	})
	return file_PlayerHeroBasicTypeInfo_proto_rawDescData
}

var file_PlayerHeroBasicTypeInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerHeroBasicTypeInfo_proto_goTypes = []interface{}{
	(*PlayerHeroBasicTypeInfo)(nil), // 0: PlayerHeroBasicTypeInfo
	(*AvatarSkillTree)(nil),         // 1: AvatarSkillTree
	(HeroBasicType)(0),              // 2: HeroBasicType
	(*EquipRelic)(nil),              // 3: EquipRelic
}
var file_PlayerHeroBasicTypeInfo_proto_depIdxs = []int32{
	1, // 0: PlayerHeroBasicTypeInfo.skill_tree_list:type_name -> AvatarSkillTree
	2, // 1: PlayerHeroBasicTypeInfo.basic_type:type_name -> HeroBasicType
	3, // 2: PlayerHeroBasicTypeInfo.equip_relic_list:type_name -> EquipRelic
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_PlayerHeroBasicTypeInfo_proto_init() }
func file_PlayerHeroBasicTypeInfo_proto_init() {
	if File_PlayerHeroBasicTypeInfo_proto != nil {
		return
	}
	file_AvatarSkillTree_proto_init()
	file_HeroBasicType_proto_init()
	file_EquipRelic_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerHeroBasicTypeInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerHeroBasicTypeInfo); i {
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
			RawDescriptor: file_PlayerHeroBasicTypeInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerHeroBasicTypeInfo_proto_goTypes,
		DependencyIndexes: file_PlayerHeroBasicTypeInfo_proto_depIdxs,
		MessageInfos:      file_PlayerHeroBasicTypeInfo_proto_msgTypes,
	}.Build()
	File_PlayerHeroBasicTypeInfo_proto = out.File
	file_PlayerHeroBasicTypeInfo_proto_rawDesc = nil
	file_PlayerHeroBasicTypeInfo_proto_goTypes = nil
	file_PlayerHeroBasicTypeInfo_proto_depIdxs = nil
}

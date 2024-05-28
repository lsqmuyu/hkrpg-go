// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: BattleAvatar.proto

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

type BattleAvatar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarType    AvatarType         `protobuf:"varint,1,opt,name=avatar_type,json=avatarType,proto3,enum=AvatarType" json:"avatar_type,omitempty"`
	Id            uint32             `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Level         uint32             `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
	Rank          uint32             `protobuf:"varint,4,opt,name=rank,proto3" json:"rank,omitempty"`
	Index         uint32             `protobuf:"varint,5,opt,name=index,proto3" json:"index,omitempty"`
	SkilltreeList []*AvatarSkillTree `protobuf:"bytes,6,rep,name=skilltree_list,json=skilltreeList,proto3" json:"skilltree_list,omitempty"`
	EquipmentList []*BattleEquipment `protobuf:"bytes,7,rep,name=equipment_list,json=equipmentList,proto3" json:"equipment_list,omitempty"`
	Hp            uint32             `protobuf:"varint,8,opt,name=hp,proto3" json:"hp,omitempty"`
	Promotion     uint32             `protobuf:"varint,10,opt,name=promotion,proto3" json:"promotion,omitempty"`
	RelicList     []*BattleRelic     `protobuf:"bytes,11,rep,name=relic_list,json=relicList,proto3" json:"relic_list,omitempty"`
	WorldLevel    uint32             `protobuf:"varint,12,opt,name=world_level,json=worldLevel,proto3" json:"world_level,omitempty"`
	AssistUid     uint32             `protobuf:"varint,13,opt,name=assist_uid,json=assistUid,proto3" json:"assist_uid,omitempty"`
	EHIACJMFNJP   *BCDGIINPGNB       `protobuf:"bytes,15,opt,name=EHIACJMFNJP,proto3" json:"EHIACJMFNJP,omitempty"`
	SpBar         *SpBarInfo         `protobuf:"bytes,16,opt,name=sp_bar,json=spBar,proto3" json:"sp_bar,omitempty"`
	IKKHEKNCKAE   uint32             `protobuf:"varint,17,opt,name=IKKHEKNCKAE,proto3" json:"IKKHEKNCKAE,omitempty"`
}

func (x *BattleAvatar) Reset() {
	*x = BattleAvatar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BattleAvatar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BattleAvatar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BattleAvatar) ProtoMessage() {}

func (x *BattleAvatar) ProtoReflect() protoreflect.Message {
	mi := &file_BattleAvatar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BattleAvatar.ProtoReflect.Descriptor instead.
func (*BattleAvatar) Descriptor() ([]byte, []int) {
	return file_BattleAvatar_proto_rawDescGZIP(), []int{0}
}

func (x *BattleAvatar) GetAvatarType() AvatarType {
	if x != nil {
		return x.AvatarType
	}
	return AvatarType_AVATAR_TYPE_NONE
}

func (x *BattleAvatar) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BattleAvatar) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *BattleAvatar) GetRank() uint32 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *BattleAvatar) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *BattleAvatar) GetSkilltreeList() []*AvatarSkillTree {
	if x != nil {
		return x.SkilltreeList
	}
	return nil
}

func (x *BattleAvatar) GetEquipmentList() []*BattleEquipment {
	if x != nil {
		return x.EquipmentList
	}
	return nil
}

func (x *BattleAvatar) GetHp() uint32 {
	if x != nil {
		return x.Hp
	}
	return 0
}

func (x *BattleAvatar) GetPromotion() uint32 {
	if x != nil {
		return x.Promotion
	}
	return 0
}

func (x *BattleAvatar) GetRelicList() []*BattleRelic {
	if x != nil {
		return x.RelicList
	}
	return nil
}

func (x *BattleAvatar) GetWorldLevel() uint32 {
	if x != nil {
		return x.WorldLevel
	}
	return 0
}

func (x *BattleAvatar) GetAssistUid() uint32 {
	if x != nil {
		return x.AssistUid
	}
	return 0
}

func (x *BattleAvatar) GetEHIACJMFNJP() *BCDGIINPGNB {
	if x != nil {
		return x.EHIACJMFNJP
	}
	return nil
}

func (x *BattleAvatar) GetSpBar() *SpBarInfo {
	if x != nil {
		return x.SpBar
	}
	return nil
}

func (x *BattleAvatar) GetIKKHEKNCKAE() uint32 {
	if x != nil {
		return x.IKKHEKNCKAE
	}
	return 0
}

var File_BattleAvatar_proto protoreflect.FileDescriptor

var file_BattleAvatar_proto_rawDesc = []byte{
	0x0a, 0x12, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x53, 0x70, 0x42, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x54, 0x72, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x42, 0x43, 0x44, 0x47, 0x49, 0x49, 0x4e, 0x50, 0x47, 0x4e, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x45, 0x71, 0x75, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x04, 0x0a, 0x0c,
	0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x2c, 0x0a, 0x0b,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0b, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x72, 0x61, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x37, 0x0a, 0x0e, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x74, 0x72, 0x65, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x54, 0x72, 0x65, 0x65, 0x52, 0x0d, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x74, 0x72, 0x65, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0e, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x42, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0d, 0x65,
	0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x68, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x68, 0x70, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x0a, 0x72, 0x65,
	0x6c, 0x69, 0x63, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x52, 0x09, 0x72, 0x65,
	0x6c, 0x69, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x69,
	0x73, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x61, 0x73,
	0x73, 0x69, 0x73, 0x74, 0x55, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x0b, 0x45, 0x48, 0x49, 0x41, 0x43,
	0x4a, 0x4d, 0x46, 0x4e, 0x4a, 0x50, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x42,
	0x43, 0x44, 0x47, 0x49, 0x49, 0x4e, 0x50, 0x47, 0x4e, 0x42, 0x52, 0x0b, 0x45, 0x48, 0x49, 0x41,
	0x43, 0x4a, 0x4d, 0x46, 0x4e, 0x4a, 0x50, 0x12, 0x21, 0x0a, 0x06, 0x73, 0x70, 0x5f, 0x62, 0x61,
	0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x70, 0x42, 0x61, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x05, 0x73, 0x70, 0x42, 0x61, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x4b,
	0x4b, 0x48, 0x45, 0x4b, 0x4e, 0x43, 0x4b, 0x41, 0x45, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x49, 0x4b, 0x4b, 0x48, 0x45, 0x4b, 0x4e, 0x43, 0x4b, 0x41, 0x45, 0x42, 0x28, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BattleAvatar_proto_rawDescOnce sync.Once
	file_BattleAvatar_proto_rawDescData = file_BattleAvatar_proto_rawDesc
)

func file_BattleAvatar_proto_rawDescGZIP() []byte {
	file_BattleAvatar_proto_rawDescOnce.Do(func() {
		file_BattleAvatar_proto_rawDescData = protoimpl.X.CompressGZIP(file_BattleAvatar_proto_rawDescData)
	})
	return file_BattleAvatar_proto_rawDescData
}

var file_BattleAvatar_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BattleAvatar_proto_goTypes = []interface{}{
	(*BattleAvatar)(nil),    // 0: BattleAvatar
	(AvatarType)(0),         // 1: AvatarType
	(*AvatarSkillTree)(nil), // 2: AvatarSkillTree
	(*BattleEquipment)(nil), // 3: BattleEquipment
	(*BattleRelic)(nil),     // 4: BattleRelic
	(*BCDGIINPGNB)(nil),     // 5: BCDGIINPGNB
	(*SpBarInfo)(nil),       // 6: SpBarInfo
}
var file_BattleAvatar_proto_depIdxs = []int32{
	1, // 0: BattleAvatar.avatar_type:type_name -> AvatarType
	2, // 1: BattleAvatar.skilltree_list:type_name -> AvatarSkillTree
	3, // 2: BattleAvatar.equipment_list:type_name -> BattleEquipment
	4, // 3: BattleAvatar.relic_list:type_name -> BattleRelic
	5, // 4: BattleAvatar.EHIACJMFNJP:type_name -> BCDGIINPGNB
	6, // 5: BattleAvatar.sp_bar:type_name -> SpBarInfo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_BattleAvatar_proto_init() }
func file_BattleAvatar_proto_init() {
	if File_BattleAvatar_proto != nil {
		return
	}
	file_SpBarInfo_proto_init()
	file_AvatarType_proto_init()
	file_AvatarSkillTree_proto_init()
	file_BCDGIINPGNB_proto_init()
	file_BattleRelic_proto_init()
	file_BattleEquipment_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BattleAvatar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BattleAvatar); i {
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
			RawDescriptor: file_BattleAvatar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BattleAvatar_proto_goTypes,
		DependencyIndexes: file_BattleAvatar_proto_depIdxs,
		MessageInfos:      file_BattleAvatar_proto_msgTypes,
	}.Build()
	File_BattleAvatar_proto = out.File
	file_BattleAvatar_proto_rawDesc = nil
	file_BattleAvatar_proto_goTypes = nil
	file_BattleAvatar_proto_depIdxs = nil
}

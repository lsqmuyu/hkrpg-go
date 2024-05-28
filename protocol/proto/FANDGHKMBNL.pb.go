// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: FANDGHKMBNL.proto

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

type FANDGHKMBNL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EPPHHCPFLNC uint32                    `protobuf:"varint,6,opt,name=EPPHHCPFLNC,proto3" json:"EPPHHCPFLNC,omitempty"`
	MDNLMMAMEJD *PunkLordBattleRecordList `protobuf:"bytes,8,opt,name=MDNLMMAMEJD,proto3" json:"MDNLMMAMEJD,omitempty"`
	BasicInfo   *CLDJMHDELHN              `protobuf:"bytes,12,opt,name=basic_info,json=basicInfo,proto3" json:"basic_info,omitempty"`
	NBEJEMIOONI PunkLordAttackerStatus    `protobuf:"varint,1,opt,name=NBEJEMIOONI,proto3,enum=PunkLordAttackerStatus" json:"NBEJEMIOONI,omitempty"`
}

func (x *FANDGHKMBNL) Reset() {
	*x = FANDGHKMBNL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FANDGHKMBNL_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FANDGHKMBNL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FANDGHKMBNL) ProtoMessage() {}

func (x *FANDGHKMBNL) ProtoReflect() protoreflect.Message {
	mi := &file_FANDGHKMBNL_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FANDGHKMBNL.ProtoReflect.Descriptor instead.
func (*FANDGHKMBNL) Descriptor() ([]byte, []int) {
	return file_FANDGHKMBNL_proto_rawDescGZIP(), []int{0}
}

func (x *FANDGHKMBNL) GetEPPHHCPFLNC() uint32 {
	if x != nil {
		return x.EPPHHCPFLNC
	}
	return 0
}

func (x *FANDGHKMBNL) GetMDNLMMAMEJD() *PunkLordBattleRecordList {
	if x != nil {
		return x.MDNLMMAMEJD
	}
	return nil
}

func (x *FANDGHKMBNL) GetBasicInfo() *CLDJMHDELHN {
	if x != nil {
		return x.BasicInfo
	}
	return nil
}

func (x *FANDGHKMBNL) GetNBEJEMIOONI() PunkLordAttackerStatus {
	if x != nil {
		return x.NBEJEMIOONI
	}
	return PunkLordAttackerStatus_PUNK_LORD_ATTACKER_STATUS_NONE
}

var File_FANDGHKMBNL_proto protoreflect.FileDescriptor

var file_FANDGHKMBNL_proto_rawDesc = []byte{
	0x0a, 0x11, 0x46, 0x41, 0x4e, 0x44, 0x47, 0x48, 0x4b, 0x4d, 0x42, 0x4e, 0x4c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x43, 0x4c, 0x44, 0x4a, 0x4d, 0x48, 0x44, 0x45, 0x4c, 0x48, 0x4e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64,
	0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x01, 0x0a, 0x0b, 0x46, 0x41, 0x4e, 0x44, 0x47, 0x48, 0x4b,
	0x4d, 0x42, 0x4e, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x50, 0x50, 0x48, 0x48, 0x43, 0x50, 0x46,
	0x4c, 0x4e, 0x43, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x50, 0x50, 0x48, 0x48,
	0x43, 0x50, 0x46, 0x4c, 0x4e, 0x43, 0x12, 0x3b, 0x0a, 0x0b, 0x4d, 0x44, 0x4e, 0x4c, 0x4d, 0x4d,
	0x41, 0x4d, 0x45, 0x4a, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x50, 0x75,
	0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0b, 0x4d, 0x44, 0x4e, 0x4c, 0x4d, 0x4d, 0x41, 0x4d,
	0x45, 0x4a, 0x44, 0x12, 0x2b, 0x0a, 0x0a, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x4c, 0x44, 0x4a, 0x4d, 0x48,
	0x44, 0x45, 0x4c, 0x48, 0x4e, 0x52, 0x09, 0x62, 0x61, 0x73, 0x69, 0x63, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x39, 0x0a, 0x0b, 0x4e, 0x42, 0x45, 0x4a, 0x45, 0x4d, 0x49, 0x4f, 0x4f, 0x4e, 0x49, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x50, 0x75, 0x6e, 0x6b, 0x4c, 0x6f, 0x72, 0x64,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0b,
	0x4e, 0x42, 0x45, 0x4a, 0x45, 0x4d, 0x49, 0x4f, 0x4f, 0x4e, 0x49, 0x42, 0x28, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e,
	0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FANDGHKMBNL_proto_rawDescOnce sync.Once
	file_FANDGHKMBNL_proto_rawDescData = file_FANDGHKMBNL_proto_rawDesc
)

func file_FANDGHKMBNL_proto_rawDescGZIP() []byte {
	file_FANDGHKMBNL_proto_rawDescOnce.Do(func() {
		file_FANDGHKMBNL_proto_rawDescData = protoimpl.X.CompressGZIP(file_FANDGHKMBNL_proto_rawDescData)
	})
	return file_FANDGHKMBNL_proto_rawDescData
}

var file_FANDGHKMBNL_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FANDGHKMBNL_proto_goTypes = []interface{}{
	(*FANDGHKMBNL)(nil),              // 0: FANDGHKMBNL
	(*PunkLordBattleRecordList)(nil), // 1: PunkLordBattleRecordList
	(*CLDJMHDELHN)(nil),              // 2: CLDJMHDELHN
	(PunkLordAttackerStatus)(0),      // 3: PunkLordAttackerStatus
}
var file_FANDGHKMBNL_proto_depIdxs = []int32{
	1, // 0: FANDGHKMBNL.MDNLMMAMEJD:type_name -> PunkLordBattleRecordList
	2, // 1: FANDGHKMBNL.basic_info:type_name -> CLDJMHDELHN
	3, // 2: FANDGHKMBNL.NBEJEMIOONI:type_name -> PunkLordAttackerStatus
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_FANDGHKMBNL_proto_init() }
func file_FANDGHKMBNL_proto_init() {
	if File_FANDGHKMBNL_proto != nil {
		return
	}
	file_CLDJMHDELHN_proto_init()
	file_PunkLordBattleRecordList_proto_init()
	file_PunkLordAttackerStatus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FANDGHKMBNL_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FANDGHKMBNL); i {
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
			RawDescriptor: file_FANDGHKMBNL_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FANDGHKMBNL_proto_goTypes,
		DependencyIndexes: file_FANDGHKMBNL_proto_depIdxs,
		MessageInfos:      file_FANDGHKMBNL_proto_msgTypes,
	}.Build()
	File_FANDGHKMBNL_proto = out.File
	file_FANDGHKMBNL_proto_rawDesc = nil
	file_FANDGHKMBNL_proto_goTypes = nil
	file_FANDGHKMBNL_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SetTurnFoodSwitchScRsp.proto

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

type SetTurnFoodSwitchScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode     uint32         `protobuf:"varint,12,opt,name=retcode,proto3" json:"retcode,omitempty"`
	DBMEPFHDMBD TurnFoodSwitch `protobuf:"varint,3,opt,name=DBMEPFHDMBD,proto3,enum=TurnFoodSwitch" json:"DBMEPFHDMBD,omitempty"`
	ACJODJLOPCJ bool           `protobuf:"varint,2,opt,name=ACJODJLOPCJ,proto3" json:"ACJODJLOPCJ,omitempty"`
}

func (x *SetTurnFoodSwitchScRsp) Reset() {
	*x = SetTurnFoodSwitchScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SetTurnFoodSwitchScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTurnFoodSwitchScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTurnFoodSwitchScRsp) ProtoMessage() {}

func (x *SetTurnFoodSwitchScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_SetTurnFoodSwitchScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTurnFoodSwitchScRsp.ProtoReflect.Descriptor instead.
func (*SetTurnFoodSwitchScRsp) Descriptor() ([]byte, []int) {
	return file_SetTurnFoodSwitchScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *SetTurnFoodSwitchScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *SetTurnFoodSwitchScRsp) GetDBMEPFHDMBD() TurnFoodSwitch {
	if x != nil {
		return x.DBMEPFHDMBD
	}
	return TurnFoodSwitch_TURN_FOOD_SWITCH_NONE
}

func (x *SetTurnFoodSwitchScRsp) GetACJODJLOPCJ() bool {
	if x != nil {
		return x.ACJODJLOPCJ
	}
	return false
}

var File_SetTurnFoodSwitchScRsp_proto protoreflect.FileDescriptor

var file_SetTurnFoodSwitchScRsp_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x53, 0x65, 0x74, 0x54, 0x75, 0x72, 0x6e, 0x46, 0x6f, 0x6f, 0x64, 0x53, 0x77, 0x69,
	0x74, 0x63, 0x68, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14,
	0x54, 0x75, 0x72, 0x6e, 0x46, 0x6f, 0x6f, 0x64, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x16, 0x53, 0x65, 0x74, 0x54, 0x75, 0x72, 0x6e,
	0x46, 0x6f, 0x6f, 0x64, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x44, 0x42, 0x4d,
	0x45, 0x50, 0x46, 0x48, 0x44, 0x4d, 0x42, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f,
	0x2e, 0x54, 0x75, 0x72, 0x6e, 0x46, 0x6f, 0x6f, 0x64, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x52,
	0x0b, 0x44, 0x42, 0x4d, 0x45, 0x50, 0x46, 0x48, 0x44, 0x4d, 0x42, 0x44, 0x12, 0x20, 0x0a, 0x0b,
	0x41, 0x43, 0x4a, 0x4f, 0x44, 0x4a, 0x4c, 0x4f, 0x50, 0x43, 0x4a, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x41, 0x43, 0x4a, 0x4f, 0x44, 0x4a, 0x4c, 0x4f, 0x50, 0x43, 0x4a, 0x42, 0x28,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67,
	0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SetTurnFoodSwitchScRsp_proto_rawDescOnce sync.Once
	file_SetTurnFoodSwitchScRsp_proto_rawDescData = file_SetTurnFoodSwitchScRsp_proto_rawDesc
)

func file_SetTurnFoodSwitchScRsp_proto_rawDescGZIP() []byte {
	file_SetTurnFoodSwitchScRsp_proto_rawDescOnce.Do(func() {
		file_SetTurnFoodSwitchScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_SetTurnFoodSwitchScRsp_proto_rawDescData)
	})
	return file_SetTurnFoodSwitchScRsp_proto_rawDescData
}

var file_SetTurnFoodSwitchScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SetTurnFoodSwitchScRsp_proto_goTypes = []interface{}{
	(*SetTurnFoodSwitchScRsp)(nil), // 0: SetTurnFoodSwitchScRsp
	(TurnFoodSwitch)(0),            // 1: TurnFoodSwitch
}
var file_SetTurnFoodSwitchScRsp_proto_depIdxs = []int32{
	1, // 0: SetTurnFoodSwitchScRsp.DBMEPFHDMBD:type_name -> TurnFoodSwitch
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SetTurnFoodSwitchScRsp_proto_init() }
func file_SetTurnFoodSwitchScRsp_proto_init() {
	if File_SetTurnFoodSwitchScRsp_proto != nil {
		return
	}
	file_TurnFoodSwitch_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SetTurnFoodSwitchScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetTurnFoodSwitchScRsp); i {
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
			RawDescriptor: file_SetTurnFoodSwitchScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SetTurnFoodSwitchScRsp_proto_goTypes,
		DependencyIndexes: file_SetTurnFoodSwitchScRsp_proto_depIdxs,
		MessageInfos:      file_SetTurnFoodSwitchScRsp_proto_msgTypes,
	}.Build()
	File_SetTurnFoodSwitchScRsp_proto = out.File
	file_SetTurnFoodSwitchScRsp_proto_rawDesc = nil
	file_SetTurnFoodSwitchScRsp_proto_goTypes = nil
	file_SetTurnFoodSwitchScRsp_proto_depIdxs = nil
}

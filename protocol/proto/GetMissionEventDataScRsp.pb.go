// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetMissionEventDataScRsp.proto

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

type GetMissionEventDataScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode          uint32     `protobuf:"varint,7,opt,name=retcode,proto3" json:"retcode,omitempty"`
	ChallengeEventId uint32     `protobuf:"varint,3,opt,name=challenge_event_id,json=challengeEventId,proto3" json:"challenge_event_id,omitempty"`
	MissionEventList []*Mission `protobuf:"bytes,15,rep,name=mission_event_list,json=missionEventList,proto3" json:"mission_event_list,omitempty"`
}

func (x *GetMissionEventDataScRsp) Reset() {
	*x = GetMissionEventDataScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetMissionEventDataScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMissionEventDataScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMissionEventDataScRsp) ProtoMessage() {}

func (x *GetMissionEventDataScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetMissionEventDataScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMissionEventDataScRsp.ProtoReflect.Descriptor instead.
func (*GetMissionEventDataScRsp) Descriptor() ([]byte, []int) {
	return file_GetMissionEventDataScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetMissionEventDataScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetMissionEventDataScRsp) GetChallengeEventId() uint32 {
	if x != nil {
		return x.ChallengeEventId
	}
	return 0
}

func (x *GetMissionEventDataScRsp) GetMissionEventList() []*Mission {
	if x != nil {
		return x.MissionEventList
	}
	return nil
}

var File_GetMissionEventDataScRsp_proto protoreflect.FileDescriptor

var file_GetMissionEventDataScRsp_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0d, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9a, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65,
	0x6e, 0x67, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x10, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x12, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x28, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetMissionEventDataScRsp_proto_rawDescOnce sync.Once
	file_GetMissionEventDataScRsp_proto_rawDescData = file_GetMissionEventDataScRsp_proto_rawDesc
)

func file_GetMissionEventDataScRsp_proto_rawDescGZIP() []byte {
	file_GetMissionEventDataScRsp_proto_rawDescOnce.Do(func() {
		file_GetMissionEventDataScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetMissionEventDataScRsp_proto_rawDescData)
	})
	return file_GetMissionEventDataScRsp_proto_rawDescData
}

var file_GetMissionEventDataScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetMissionEventDataScRsp_proto_goTypes = []interface{}{
	(*GetMissionEventDataScRsp)(nil), // 0: GetMissionEventDataScRsp
	(*Mission)(nil),                  // 1: Mission
}
var file_GetMissionEventDataScRsp_proto_depIdxs = []int32{
	1, // 0: GetMissionEventDataScRsp.mission_event_list:type_name -> Mission
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetMissionEventDataScRsp_proto_init() }
func file_GetMissionEventDataScRsp_proto_init() {
	if File_GetMissionEventDataScRsp_proto != nil {
		return
	}
	file_Mission_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetMissionEventDataScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMissionEventDataScRsp); i {
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
			RawDescriptor: file_GetMissionEventDataScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetMissionEventDataScRsp_proto_goTypes,
		DependencyIndexes: file_GetMissionEventDataScRsp_proto_depIdxs,
		MessageInfos:      file_GetMissionEventDataScRsp_proto_msgTypes,
	}.Build()
	File_GetMissionEventDataScRsp_proto = out.File
	file_GetMissionEventDataScRsp_proto_rawDesc = nil
	file_GetMissionEventDataScRsp_proto_goTypes = nil
	file_GetMissionEventDataScRsp_proto_depIdxs = nil
}

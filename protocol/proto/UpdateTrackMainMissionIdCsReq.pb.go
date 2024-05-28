// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: UpdateTrackMainMissionIdCsReq.proto

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

type UpdateTrackMainMissionIdCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AIODMCNLNPE            uint32                 `protobuf:"varint,8,opt,name=AIODMCNLNPE,proto3" json:"AIODMCNLNPE,omitempty"`
	TrackMainMissionStatus TrackMainMissionUpdate `protobuf:"varint,3,opt,name=track_main_mission_status,json=trackMainMissionStatus,proto3,enum=TrackMainMissionUpdate" json:"track_main_mission_status,omitempty"`
	EKMONACBDKK            uint32                 `protobuf:"varint,9,opt,name=EKMONACBDKK,proto3" json:"EKMONACBDKK,omitempty"`
}

func (x *UpdateTrackMainMissionIdCsReq) Reset() {
	*x = UpdateTrackMainMissionIdCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UpdateTrackMainMissionIdCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTrackMainMissionIdCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTrackMainMissionIdCsReq) ProtoMessage() {}

func (x *UpdateTrackMainMissionIdCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_UpdateTrackMainMissionIdCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTrackMainMissionIdCsReq.ProtoReflect.Descriptor instead.
func (*UpdateTrackMainMissionIdCsReq) Descriptor() ([]byte, []int) {
	return file_UpdateTrackMainMissionIdCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateTrackMainMissionIdCsReq) GetAIODMCNLNPE() uint32 {
	if x != nil {
		return x.AIODMCNLNPE
	}
	return 0
}

func (x *UpdateTrackMainMissionIdCsReq) GetTrackMainMissionStatus() TrackMainMissionUpdate {
	if x != nil {
		return x.TrackMainMissionStatus
	}
	return TrackMainMissionUpdate_TRACK_MAIN_MISSION_UPDATE_NONE
}

func (x *UpdateTrackMainMissionIdCsReq) GetEKMONACBDKK() uint32 {
	if x != nil {
		return x.EKMONACBDKK
	}
	return 0
}

var File_UpdateTrackMainMissionIdCsReq_proto protoreflect.FileDescriptor

var file_UpdateTrackMainMissionIdCsReq_proto_rawDesc = []byte{
	0x0a, 0x23, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x4d, 0x61, 0x69,
	0x6e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x4d, 0x61, 0x69, 0x6e,
	0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x01, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x4d, 0x61, 0x69, 0x6e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x49, 0x4f, 0x44, 0x4d, 0x43, 0x4e,
	0x4c, 0x4e, 0x50, 0x45, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x41, 0x49, 0x4f, 0x44,
	0x4d, 0x43, 0x4e, 0x4c, 0x4e, 0x50, 0x45, 0x12, 0x52, 0x0a, 0x19, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x4d, 0x61, 0x69, 0x6e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x16, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x4d, 0x61, 0x69, 0x6e, 0x4d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x45,
	0x4b, 0x4d, 0x4f, 0x4e, 0x41, 0x43, 0x42, 0x44, 0x4b, 0x4b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x45, 0x4b, 0x4d, 0x4f, 0x4e, 0x41, 0x43, 0x42, 0x44, 0x4b, 0x4b, 0x42, 0x28, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_UpdateTrackMainMissionIdCsReq_proto_rawDescOnce sync.Once
	file_UpdateTrackMainMissionIdCsReq_proto_rawDescData = file_UpdateTrackMainMissionIdCsReq_proto_rawDesc
)

func file_UpdateTrackMainMissionIdCsReq_proto_rawDescGZIP() []byte {
	file_UpdateTrackMainMissionIdCsReq_proto_rawDescOnce.Do(func() {
		file_UpdateTrackMainMissionIdCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_UpdateTrackMainMissionIdCsReq_proto_rawDescData)
	})
	return file_UpdateTrackMainMissionIdCsReq_proto_rawDescData
}

var file_UpdateTrackMainMissionIdCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_UpdateTrackMainMissionIdCsReq_proto_goTypes = []interface{}{
	(*UpdateTrackMainMissionIdCsReq)(nil), // 0: UpdateTrackMainMissionIdCsReq
	(TrackMainMissionUpdate)(0),           // 1: TrackMainMissionUpdate
}
var file_UpdateTrackMainMissionIdCsReq_proto_depIdxs = []int32{
	1, // 0: UpdateTrackMainMissionIdCsReq.track_main_mission_status:type_name -> TrackMainMissionUpdate
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_UpdateTrackMainMissionIdCsReq_proto_init() }
func file_UpdateTrackMainMissionIdCsReq_proto_init() {
	if File_UpdateTrackMainMissionIdCsReq_proto != nil {
		return
	}
	file_TrackMainMissionUpdate_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_UpdateTrackMainMissionIdCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTrackMainMissionIdCsReq); i {
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
			RawDescriptor: file_UpdateTrackMainMissionIdCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_UpdateTrackMainMissionIdCsReq_proto_goTypes,
		DependencyIndexes: file_UpdateTrackMainMissionIdCsReq_proto_depIdxs,
		MessageInfos:      file_UpdateTrackMainMissionIdCsReq_proto_msgTypes,
	}.Build()
	File_UpdateTrackMainMissionIdCsReq_proto = out.File
	file_UpdateTrackMainMissionIdCsReq_proto_rawDesc = nil
	file_UpdateTrackMainMissionIdCsReq_proto_goTypes = nil
	file_UpdateTrackMainMissionIdCsReq_proto_depIdxs = nil
}

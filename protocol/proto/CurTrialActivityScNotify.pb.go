// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: CurTrialActivityScNotify.proto

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

type CurTrialActivityScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LEPLEBFIOBP uint32              `protobuf:"varint,9,opt,name=LEPLEBFIOBP,proto3" json:"LEPLEBFIOBP,omitempty"`
	Status      TrialActivityStatus `protobuf:"varint,6,opt,name=status,proto3,enum=TrialActivityStatus" json:"status,omitempty"`
}

func (x *CurTrialActivityScNotify) Reset() {
	*x = CurTrialActivityScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CurTrialActivityScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurTrialActivityScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurTrialActivityScNotify) ProtoMessage() {}

func (x *CurTrialActivityScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_CurTrialActivityScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurTrialActivityScNotify.ProtoReflect.Descriptor instead.
func (*CurTrialActivityScNotify) Descriptor() ([]byte, []int) {
	return file_CurTrialActivityScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *CurTrialActivityScNotify) GetLEPLEBFIOBP() uint32 {
	if x != nil {
		return x.LEPLEBFIOBP
	}
	return 0
}

func (x *CurTrialActivityScNotify) GetStatus() TrialActivityStatus {
	if x != nil {
		return x.Status
	}
	return TrialActivityStatus_TRIAL_ACTIVITY_STATUS_NONE
}

var File_CurTrialActivityScNotify_proto protoreflect.FileDescriptor

var file_CurTrialActivityScNotify_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x43, 0x75, 0x72, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x18, 0x43,
	0x75, 0x72, 0x54, 0x72, 0x69, 0x61, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53,
	0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x45, 0x50, 0x4c, 0x45,
	0x42, 0x46, 0x49, 0x4f, 0x42, 0x50, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x45,
	0x50, 0x4c, 0x45, 0x42, 0x46, 0x49, 0x4f, 0x42, 0x50, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x54, 0x72, 0x69, 0x61,
	0x6c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CurTrialActivityScNotify_proto_rawDescOnce sync.Once
	file_CurTrialActivityScNotify_proto_rawDescData = file_CurTrialActivityScNotify_proto_rawDesc
)

func file_CurTrialActivityScNotify_proto_rawDescGZIP() []byte {
	file_CurTrialActivityScNotify_proto_rawDescOnce.Do(func() {
		file_CurTrialActivityScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_CurTrialActivityScNotify_proto_rawDescData)
	})
	return file_CurTrialActivityScNotify_proto_rawDescData
}

var file_CurTrialActivityScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_CurTrialActivityScNotify_proto_goTypes = []interface{}{
	(*CurTrialActivityScNotify)(nil), // 0: CurTrialActivityScNotify
	(TrialActivityStatus)(0),         // 1: TrialActivityStatus
}
var file_CurTrialActivityScNotify_proto_depIdxs = []int32{
	1, // 0: CurTrialActivityScNotify.status:type_name -> TrialActivityStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CurTrialActivityScNotify_proto_init() }
func file_CurTrialActivityScNotify_proto_init() {
	if File_CurTrialActivityScNotify_proto != nil {
		return
	}
	file_TrialActivityStatus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CurTrialActivityScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurTrialActivityScNotify); i {
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
			RawDescriptor: file_CurTrialActivityScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CurTrialActivityScNotify_proto_goTypes,
		DependencyIndexes: file_CurTrialActivityScNotify_proto_depIdxs,
		MessageInfos:      file_CurTrialActivityScNotify_proto_msgTypes,
	}.Build()
	File_CurTrialActivityScNotify_proto = out.File
	file_CurTrialActivityScNotify_proto_rawDesc = nil
	file_CurTrialActivityScNotify_proto_goTypes = nil
	file_CurTrialActivityScNotify_proto_depIdxs = nil
}

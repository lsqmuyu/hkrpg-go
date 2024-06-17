// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: EnterMapRotationRegionCsReq.proto

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

type EnterMapRotationRegionCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MNNHPJGLOKM uint32      `protobuf:"varint,12,opt,name=MNNHPJGLOKM,proto3" json:"MNNHPJGLOKM,omitempty"`
	Motion      *MotionInfo `protobuf:"bytes,7,opt,name=motion,proto3" json:"motion,omitempty"`
	Rotation    uint32      `protobuf:"varint,13,opt,name=rotation,proto3" json:"rotation,omitempty"`
}

func (x *EnterMapRotationRegionCsReq) Reset() {
	*x = EnterMapRotationRegionCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EnterMapRotationRegionCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnterMapRotationRegionCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnterMapRotationRegionCsReq) ProtoMessage() {}

func (x *EnterMapRotationRegionCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_EnterMapRotationRegionCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnterMapRotationRegionCsReq.ProtoReflect.Descriptor instead.
func (*EnterMapRotationRegionCsReq) Descriptor() ([]byte, []int) {
	return file_EnterMapRotationRegionCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *EnterMapRotationRegionCsReq) GetMNNHPJGLOKM() uint32 {
	if x != nil {
		return x.MNNHPJGLOKM
	}
	return 0
}

func (x *EnterMapRotationRegionCsReq) GetMotion() *MotionInfo {
	if x != nil {
		return x.Motion
	}
	return nil
}

func (x *EnterMapRotationRegionCsReq) GetRotation() uint32 {
	if x != nil {
		return x.Rotation
	}
	return 0
}

var File_EnterMapRotationRegionCsReq_proto protoreflect.FileDescriptor

var file_EnterMapRotationRegionCsReq_proto_rawDesc = []byte{
	0x0a, 0x21, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x70, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x1b, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x4d,
	0x61, 0x70, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x4e, 0x4e, 0x48, 0x50, 0x4a, 0x47,
	0x4c, 0x4f, 0x4b, 0x4d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4d, 0x4e, 0x4e, 0x48,
	0x50, 0x4a, 0x47, 0x4c, 0x4f, 0x4b, 0x4d, 0x12, 0x23, 0x0a, 0x06, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44,
	0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EnterMapRotationRegionCsReq_proto_rawDescOnce sync.Once
	file_EnterMapRotationRegionCsReq_proto_rawDescData = file_EnterMapRotationRegionCsReq_proto_rawDesc
)

func file_EnterMapRotationRegionCsReq_proto_rawDescGZIP() []byte {
	file_EnterMapRotationRegionCsReq_proto_rawDescOnce.Do(func() {
		file_EnterMapRotationRegionCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_EnterMapRotationRegionCsReq_proto_rawDescData)
	})
	return file_EnterMapRotationRegionCsReq_proto_rawDescData
}

var file_EnterMapRotationRegionCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EnterMapRotationRegionCsReq_proto_goTypes = []interface{}{
	(*EnterMapRotationRegionCsReq)(nil), // 0: EnterMapRotationRegionCsReq
	(*MotionInfo)(nil),                  // 1: MotionInfo
}
var file_EnterMapRotationRegionCsReq_proto_depIdxs = []int32{
	1, // 0: EnterMapRotationRegionCsReq.motion:type_name -> MotionInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_EnterMapRotationRegionCsReq_proto_init() }
func file_EnterMapRotationRegionCsReq_proto_init() {
	if File_EnterMapRotationRegionCsReq_proto != nil {
		return
	}
	file_MotionInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EnterMapRotationRegionCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnterMapRotationRegionCsReq); i {
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
			RawDescriptor: file_EnterMapRotationRegionCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EnterMapRotationRegionCsReq_proto_goTypes,
		DependencyIndexes: file_EnterMapRotationRegionCsReq_proto_depIdxs,
		MessageInfos:      file_EnterMapRotationRegionCsReq_proto_msgTypes,
	}.Build()
	File_EnterMapRotationRegionCsReq_proto = out.File
	file_EnterMapRotationRegionCsReq_proto_rawDesc = nil
	file_EnterMapRotationRegionCsReq_proto_goTypes = nil
	file_EnterMapRotationRegionCsReq_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: EntityBuffChange.proto

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

type EntityBuffChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId uint32 `protobuf:"varint,10,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	// Types that are assignable to JFMOKMCOHDJ:
	//
	//	*EntityBuffChange_BuffInfo
	//	*EntityBuffChange_DPFJOGMODLH
	JFMOKMCOHDJ isEntityBuffChange_JFMOKMCOHDJ `protobuf_oneof:"JFMOKMCOHDJ"`
}

func (x *EntityBuffChange) Reset() {
	*x = EntityBuffChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EntityBuffChange_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityBuffChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityBuffChange) ProtoMessage() {}

func (x *EntityBuffChange) ProtoReflect() protoreflect.Message {
	mi := &file_EntityBuffChange_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityBuffChange.ProtoReflect.Descriptor instead.
func (*EntityBuffChange) Descriptor() ([]byte, []int) {
	return file_EntityBuffChange_proto_rawDescGZIP(), []int{0}
}

func (x *EntityBuffChange) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (m *EntityBuffChange) GetJFMOKMCOHDJ() isEntityBuffChange_JFMOKMCOHDJ {
	if m != nil {
		return m.JFMOKMCOHDJ
	}
	return nil
}

func (x *EntityBuffChange) GetBuffInfo() *BuffInfo {
	if x, ok := x.GetJFMOKMCOHDJ().(*EntityBuffChange_BuffInfo); ok {
		return x.BuffInfo
	}
	return nil
}

func (x *EntityBuffChange) GetDPFJOGMODLH() uint32 {
	if x, ok := x.GetJFMOKMCOHDJ().(*EntityBuffChange_DPFJOGMODLH); ok {
		return x.DPFJOGMODLH
	}
	return 0
}

type isEntityBuffChange_JFMOKMCOHDJ interface {
	isEntityBuffChange_JFMOKMCOHDJ()
}

type EntityBuffChange_BuffInfo struct {
	BuffInfo *BuffInfo `protobuf:"bytes,7,opt,name=buff_info,json=buffInfo,proto3,oneof"`
}

type EntityBuffChange_DPFJOGMODLH struct {
	DPFJOGMODLH uint32 `protobuf:"varint,2,opt,name=DPFJOGMODLH,proto3,oneof"`
}

func (*EntityBuffChange_BuffInfo) isEntityBuffChange_JFMOKMCOHDJ() {}

func (*EntityBuffChange_DPFJOGMODLH) isEntityBuffChange_JFMOKMCOHDJ() {}

var File_EntityBuffChange_proto protoreflect.FileDescriptor

var file_EntityBuffChange_proto_rawDesc = []byte{
	0x0a, 0x16, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x75, 0x66, 0x66, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x42, 0x75, 0x66, 0x66, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x10, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x42, 0x75, 0x66, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x09, 0x62, 0x75,
	0x66, 0x66, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x08, 0x62, 0x75, 0x66, 0x66,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x0a, 0x0b, 0x44, 0x50, 0x46, 0x4a, 0x4f, 0x47, 0x4d, 0x4f,
	0x44, 0x4c, 0x48, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0b, 0x44, 0x50, 0x46,
	0x4a, 0x4f, 0x47, 0x4d, 0x4f, 0x44, 0x4c, 0x48, 0x42, 0x0d, 0x0a, 0x0b, 0x4a, 0x46, 0x4d, 0x4f,
	0x4b, 0x4d, 0x43, 0x4f, 0x48, 0x44, 0x4a, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EntityBuffChange_proto_rawDescOnce sync.Once
	file_EntityBuffChange_proto_rawDescData = file_EntityBuffChange_proto_rawDesc
)

func file_EntityBuffChange_proto_rawDescGZIP() []byte {
	file_EntityBuffChange_proto_rawDescOnce.Do(func() {
		file_EntityBuffChange_proto_rawDescData = protoimpl.X.CompressGZIP(file_EntityBuffChange_proto_rawDescData)
	})
	return file_EntityBuffChange_proto_rawDescData
}

var file_EntityBuffChange_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EntityBuffChange_proto_goTypes = []interface{}{
	(*EntityBuffChange)(nil), // 0: EntityBuffChange
	(*BuffInfo)(nil),         // 1: BuffInfo
}
var file_EntityBuffChange_proto_depIdxs = []int32{
	1, // 0: EntityBuffChange.buff_info:type_name -> BuffInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_EntityBuffChange_proto_init() }
func file_EntityBuffChange_proto_init() {
	if File_EntityBuffChange_proto != nil {
		return
	}
	file_BuffInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EntityBuffChange_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityBuffChange); i {
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
	file_EntityBuffChange_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*EntityBuffChange_BuffInfo)(nil),
		(*EntityBuffChange_DPFJOGMODLH)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_EntityBuffChange_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EntityBuffChange_proto_goTypes,
		DependencyIndexes: file_EntityBuffChange_proto_depIdxs,
		MessageInfos:      file_EntityBuffChange_proto_msgTypes,
	}.Build()
	File_EntityBuffChange_proto = out.File
	file_EntityBuffChange_proto_rawDesc = nil
	file_EntityBuffChange_proto_goTypes = nil
	file_EntityBuffChange_proto_depIdxs = nil
}

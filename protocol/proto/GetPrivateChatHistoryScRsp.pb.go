// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetPrivateChatHistoryScRsp.proto

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

type GetPrivateChatHistoryScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatMessageList []*ChatMessageData `protobuf:"bytes,8,rep,name=chat_message_list,json=chatMessageList,proto3" json:"chat_message_list,omitempty"`
	OLIGKFNJKMA     uint32             `protobuf:"varint,13,opt,name=OLIGKFNJKMA,proto3" json:"OLIGKFNJKMA,omitempty"`
	ContactId       uint32             `protobuf:"varint,9,opt,name=contact_id,json=contactId,proto3" json:"contact_id,omitempty"`
	Retcode         uint32             `protobuf:"varint,4,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *GetPrivateChatHistoryScRsp) Reset() {
	*x = GetPrivateChatHistoryScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetPrivateChatHistoryScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPrivateChatHistoryScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrivateChatHistoryScRsp) ProtoMessage() {}

func (x *GetPrivateChatHistoryScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetPrivateChatHistoryScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrivateChatHistoryScRsp.ProtoReflect.Descriptor instead.
func (*GetPrivateChatHistoryScRsp) Descriptor() ([]byte, []int) {
	return file_GetPrivateChatHistoryScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetPrivateChatHistoryScRsp) GetChatMessageList() []*ChatMessageData {
	if x != nil {
		return x.ChatMessageList
	}
	return nil
}

func (x *GetPrivateChatHistoryScRsp) GetOLIGKFNJKMA() uint32 {
	if x != nil {
		return x.OLIGKFNJKMA
	}
	return 0
}

func (x *GetPrivateChatHistoryScRsp) GetContactId() uint32 {
	if x != nil {
		return x.ContactId
	}
	return 0
}

func (x *GetPrivateChatHistoryScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_GetPrivateChatHistoryScRsp_proto protoreflect.FileDescriptor

var file_GetPrivateChatHistoryScRsp_proto_rawDesc = []byte{
	0x0a, 0x20, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x15, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a, 0x1a, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x3c, 0x0a, 0x11, 0x63, 0x68, 0x61, 0x74,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0f, 0x63, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x4c, 0x49, 0x47, 0x4b, 0x46,
	0x4e, 0x4a, 0x4b, 0x4d, 0x41, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x4c, 0x49,
	0x47, 0x4b, 0x46, 0x4e, 0x4a, 0x4b, 0x4d, 0x41, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_GetPrivateChatHistoryScRsp_proto_rawDescOnce sync.Once
	file_GetPrivateChatHistoryScRsp_proto_rawDescData = file_GetPrivateChatHistoryScRsp_proto_rawDesc
)

func file_GetPrivateChatHistoryScRsp_proto_rawDescGZIP() []byte {
	file_GetPrivateChatHistoryScRsp_proto_rawDescOnce.Do(func() {
		file_GetPrivateChatHistoryScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetPrivateChatHistoryScRsp_proto_rawDescData)
	})
	return file_GetPrivateChatHistoryScRsp_proto_rawDescData
}

var file_GetPrivateChatHistoryScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetPrivateChatHistoryScRsp_proto_goTypes = []interface{}{
	(*GetPrivateChatHistoryScRsp)(nil), // 0: GetPrivateChatHistoryScRsp
	(*ChatMessageData)(nil),            // 1: ChatMessageData
}
var file_GetPrivateChatHistoryScRsp_proto_depIdxs = []int32{
	1, // 0: GetPrivateChatHistoryScRsp.chat_message_list:type_name -> ChatMessageData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetPrivateChatHistoryScRsp_proto_init() }
func file_GetPrivateChatHistoryScRsp_proto_init() {
	if File_GetPrivateChatHistoryScRsp_proto != nil {
		return
	}
	file_ChatMessageData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetPrivateChatHistoryScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPrivateChatHistoryScRsp); i {
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
			RawDescriptor: file_GetPrivateChatHistoryScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetPrivateChatHistoryScRsp_proto_goTypes,
		DependencyIndexes: file_GetPrivateChatHistoryScRsp_proto_depIdxs,
		MessageInfos:      file_GetPrivateChatHistoryScRsp_proto_msgTypes,
	}.Build()
	File_GetPrivateChatHistoryScRsp_proto = out.File
	file_GetPrivateChatHistoryScRsp_proto_rawDesc = nil
	file_GetPrivateChatHistoryScRsp_proto_goTypes = nil
	file_GetPrivateChatHistoryScRsp_proto_depIdxs = nil
}

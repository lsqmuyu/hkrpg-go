// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PlayerLoginCsReq.proto

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

type PlayerLoginCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Platform      PlatformType `protobuf:"varint,9,opt,name=platform,proto3,enum=PlatformType" json:"platform,omitempty"`
	LDFIOFJHJJA   string       `protobuf:"bytes,350,opt,name=LDFIOFJHJJA,proto3" json:"LDFIOFJHJJA,omitempty"`
	EOLCNDBAELO   *EDCFBAGKIGN `protobuf:"bytes,1094,opt,name=EOLCNDBAELO,proto3" json:"EOLCNDBAELO,omitempty"`
	PHPFKNDOLND   string       `protobuf:"bytes,1800,opt,name=PHPFKNDOLND,proto3" json:"PHPFKNDOLND,omitempty"`
	ClientVersion string       `protobuf:"bytes,5,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
	PFMLAGFCCCC   string       `protobuf:"bytes,4,opt,name=PFMLAGFCCCC,proto3" json:"PFMLAGFCCCC,omitempty"`
	EDOPMEBJIDE   string       `protobuf:"bytes,2,opt,name=EDOPMEBJIDE,proto3" json:"EDOPMEBJIDE,omitempty"`
	OBGFMGEOJCJ   string       `protobuf:"bytes,12,opt,name=OBGFMGEOJCJ,proto3" json:"OBGFMGEOJCJ,omitempty"`
	AJIBAOAICKM   string       `protobuf:"bytes,1045,opt,name=AJIBAOAICKM,proto3" json:"AJIBAOAICKM,omitempty"`
	ResVersion    uint32       `protobuf:"varint,15,opt,name=res_version,json=resVersion,proto3" json:"res_version,omitempty"`
	LoginRandom   uint64       `protobuf:"varint,10,opt,name=login_random,json=loginRandom,proto3" json:"login_random,omitempty"`
	AGDIEDBLCML   string       `protobuf:"bytes,185,opt,name=AGDIEDBLCML,proto3" json:"AGDIEDBLCML,omitempty"`
	CIGAFCKHGGI   string       `protobuf:"bytes,8,opt,name=CIGAFCKHGGI,proto3" json:"CIGAFCKHGGI,omitempty"`
	JIMGBENIKEB   bool         `protobuf:"varint,704,opt,name=JIMGBENIKEB,proto3" json:"JIMGBENIKEB,omitempty"`
	GIIJCHKENDC   string       `protobuf:"bytes,3,opt,name=GIIJCHKENDC,proto3" json:"GIIJCHKENDC,omitempty"`
	GNPDAIFMHLA   string       `protobuf:"bytes,13,opt,name=GNPDAIFMHLA,proto3" json:"GNPDAIFMHLA,omitempty"`
	AILINANGJNE   string       `protobuf:"bytes,1092,opt,name=AILINANGJNE,proto3" json:"AILINANGJNE,omitempty"`
	HFFEEAEKFMI   string       `protobuf:"bytes,14,opt,name=HFFEEAEKFMI,proto3" json:"HFFEEAEKFMI,omitempty"`
	MOIKALNPCPA   uint32       `protobuf:"varint,7,opt,name=MOIKALNPCPA,proto3" json:"MOIKALNPCPA,omitempty"`
	PLLIIPJIFOG   uint32       `protobuf:"varint,277,opt,name=PLLIIPJIFOG,proto3" json:"PLLIIPJIFOG,omitempty"`
	RogueGetInfo  string       `protobuf:"bytes,11,opt,name=rogue_get_info,json=rogueGetInfo,proto3" json:"rogue_get_info,omitempty"`
	AGPOJJJKKJI   uint32       `protobuf:"varint,174,opt,name=AGPOJJJKKJI,proto3" json:"AGPOJJJKKJI,omitempty"`
	Signature     string       `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
	Language      LanguageType `protobuf:"varint,1,opt,name=language,proto3,enum=LanguageType" json:"language,omitempty"`
}

func (x *PlayerLoginCsReq) Reset() {
	*x = PlayerLoginCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerLoginCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerLoginCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerLoginCsReq) ProtoMessage() {}

func (x *PlayerLoginCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerLoginCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerLoginCsReq.ProtoReflect.Descriptor instead.
func (*PlayerLoginCsReq) Descriptor() ([]byte, []int) {
	return file_PlayerLoginCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerLoginCsReq) GetPlatform() PlatformType {
	if x != nil {
		return x.Platform
	}
	return PlatformType_EDITOR
}

func (x *PlayerLoginCsReq) GetLDFIOFJHJJA() string {
	if x != nil {
		return x.LDFIOFJHJJA
	}
	return ""
}

func (x *PlayerLoginCsReq) GetEOLCNDBAELO() *EDCFBAGKIGN {
	if x != nil {
		return x.EOLCNDBAELO
	}
	return nil
}

func (x *PlayerLoginCsReq) GetPHPFKNDOLND() string {
	if x != nil {
		return x.PHPFKNDOLND
	}
	return ""
}

func (x *PlayerLoginCsReq) GetClientVersion() string {
	if x != nil {
		return x.ClientVersion
	}
	return ""
}

func (x *PlayerLoginCsReq) GetPFMLAGFCCCC() string {
	if x != nil {
		return x.PFMLAGFCCCC
	}
	return ""
}

func (x *PlayerLoginCsReq) GetEDOPMEBJIDE() string {
	if x != nil {
		return x.EDOPMEBJIDE
	}
	return ""
}

func (x *PlayerLoginCsReq) GetOBGFMGEOJCJ() string {
	if x != nil {
		return x.OBGFMGEOJCJ
	}
	return ""
}

func (x *PlayerLoginCsReq) GetAJIBAOAICKM() string {
	if x != nil {
		return x.AJIBAOAICKM
	}
	return ""
}

func (x *PlayerLoginCsReq) GetResVersion() uint32 {
	if x != nil {
		return x.ResVersion
	}
	return 0
}

func (x *PlayerLoginCsReq) GetLoginRandom() uint64 {
	if x != nil {
		return x.LoginRandom
	}
	return 0
}

func (x *PlayerLoginCsReq) GetAGDIEDBLCML() string {
	if x != nil {
		return x.AGDIEDBLCML
	}
	return ""
}

func (x *PlayerLoginCsReq) GetCIGAFCKHGGI() string {
	if x != nil {
		return x.CIGAFCKHGGI
	}
	return ""
}

func (x *PlayerLoginCsReq) GetJIMGBENIKEB() bool {
	if x != nil {
		return x.JIMGBENIKEB
	}
	return false
}

func (x *PlayerLoginCsReq) GetGIIJCHKENDC() string {
	if x != nil {
		return x.GIIJCHKENDC
	}
	return ""
}

func (x *PlayerLoginCsReq) GetGNPDAIFMHLA() string {
	if x != nil {
		return x.GNPDAIFMHLA
	}
	return ""
}

func (x *PlayerLoginCsReq) GetAILINANGJNE() string {
	if x != nil {
		return x.AILINANGJNE
	}
	return ""
}

func (x *PlayerLoginCsReq) GetHFFEEAEKFMI() string {
	if x != nil {
		return x.HFFEEAEKFMI
	}
	return ""
}

func (x *PlayerLoginCsReq) GetMOIKALNPCPA() uint32 {
	if x != nil {
		return x.MOIKALNPCPA
	}
	return 0
}

func (x *PlayerLoginCsReq) GetPLLIIPJIFOG() uint32 {
	if x != nil {
		return x.PLLIIPJIFOG
	}
	return 0
}

func (x *PlayerLoginCsReq) GetRogueGetInfo() string {
	if x != nil {
		return x.RogueGetInfo
	}
	return ""
}

func (x *PlayerLoginCsReq) GetAGPOJJJKKJI() uint32 {
	if x != nil {
		return x.AGPOJJJKKJI
	}
	return 0
}

func (x *PlayerLoginCsReq) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *PlayerLoginCsReq) GetLanguage() LanguageType {
	if x != nil {
		return x.Language
	}
	return LanguageType_LANGUAGE_NONE
}

var File_PlayerLoginCsReq_proto protoreflect.FileDescriptor

var file_PlayerLoginCsReq_proto_rawDesc = []byte{
	0x0a, 0x16, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x73, 0x52,
	0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x45, 0x44, 0x43, 0x46, 0x42, 0x41, 0x47, 0x4b, 0x49, 0x47, 0x4e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x06, 0x0a, 0x10, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x29, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x12, 0x21, 0x0a, 0x0b, 0x4c, 0x44, 0x46, 0x49, 0x4f, 0x46, 0x4a, 0x48, 0x4a,
	0x4a, 0x41, 0x18, 0xde, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4c, 0x44, 0x46, 0x49, 0x4f,
	0x46, 0x4a, 0x48, 0x4a, 0x4a, 0x41, 0x12, 0x2f, 0x0a, 0x0b, 0x45, 0x4f, 0x4c, 0x43, 0x4e, 0x44,
	0x42, 0x41, 0x45, 0x4c, 0x4f, 0x18, 0xc6, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45,
	0x44, 0x43, 0x46, 0x42, 0x41, 0x47, 0x4b, 0x49, 0x47, 0x4e, 0x52, 0x0b, 0x45, 0x4f, 0x4c, 0x43,
	0x4e, 0x44, 0x42, 0x41, 0x45, 0x4c, 0x4f, 0x12, 0x21, 0x0a, 0x0b, 0x50, 0x48, 0x50, 0x46, 0x4b,
	0x4e, 0x44, 0x4f, 0x4c, 0x4e, 0x44, 0x18, 0x88, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50,
	0x48, 0x50, 0x46, 0x4b, 0x4e, 0x44, 0x4f, 0x4c, 0x4e, 0x44, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x46, 0x4d, 0x4c, 0x41, 0x47, 0x46, 0x43, 0x43, 0x43, 0x43,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x50, 0x46, 0x4d, 0x4c, 0x41, 0x47, 0x46, 0x43,
	0x43, 0x43, 0x43, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x44, 0x4f, 0x50, 0x4d, 0x45, 0x42, 0x4a, 0x49,
	0x44, 0x45, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x45, 0x44, 0x4f, 0x50, 0x4d, 0x45,
	0x42, 0x4a, 0x49, 0x44, 0x45, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x42, 0x47, 0x46, 0x4d, 0x47, 0x45,
	0x4f, 0x4a, 0x43, 0x4a, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4f, 0x42, 0x47, 0x46,
	0x4d, 0x47, 0x45, 0x4f, 0x4a, 0x43, 0x4a, 0x12, 0x21, 0x0a, 0x0b, 0x41, 0x4a, 0x49, 0x42, 0x41,
	0x4f, 0x41, 0x49, 0x43, 0x4b, 0x4d, 0x18, 0x95, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41,
	0x4a, 0x49, 0x42, 0x41, 0x4f, 0x41, 0x49, 0x43, 0x4b, 0x4d, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x72, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x12, 0x21,
	0x0a, 0x0b, 0x41, 0x47, 0x44, 0x49, 0x45, 0x44, 0x42, 0x4c, 0x43, 0x4d, 0x4c, 0x18, 0xb9, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x47, 0x44, 0x49, 0x45, 0x44, 0x42, 0x4c, 0x43, 0x4d,
	0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x49, 0x47, 0x41, 0x46, 0x43, 0x4b, 0x48, 0x47, 0x47, 0x49,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x49, 0x47, 0x41, 0x46, 0x43, 0x4b, 0x48,
	0x47, 0x47, 0x49, 0x12, 0x21, 0x0a, 0x0b, 0x4a, 0x49, 0x4d, 0x47, 0x42, 0x45, 0x4e, 0x49, 0x4b,
	0x45, 0x42, 0x18, 0xc0, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4a, 0x49, 0x4d, 0x47, 0x42,
	0x45, 0x4e, 0x49, 0x4b, 0x45, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x49, 0x49, 0x4a, 0x43, 0x48,
	0x4b, 0x45, 0x4e, 0x44, 0x43, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x47, 0x49, 0x49,
	0x4a, 0x43, 0x48, 0x4b, 0x45, 0x4e, 0x44, 0x43, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x4e, 0x50, 0x44,
	0x41, 0x49, 0x46, 0x4d, 0x48, 0x4c, 0x41, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x47,
	0x4e, 0x50, 0x44, 0x41, 0x49, 0x46, 0x4d, 0x48, 0x4c, 0x41, 0x12, 0x21, 0x0a, 0x0b, 0x41, 0x49,
	0x4c, 0x49, 0x4e, 0x41, 0x4e, 0x47, 0x4a, 0x4e, 0x45, 0x18, 0xc4, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x41, 0x49, 0x4c, 0x49, 0x4e, 0x41, 0x4e, 0x47, 0x4a, 0x4e, 0x45, 0x12, 0x20, 0x0a,
	0x0b, 0x48, 0x46, 0x46, 0x45, 0x45, 0x41, 0x45, 0x4b, 0x46, 0x4d, 0x49, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x48, 0x46, 0x46, 0x45, 0x45, 0x41, 0x45, 0x4b, 0x46, 0x4d, 0x49, 0x12,
	0x20, 0x0a, 0x0b, 0x4d, 0x4f, 0x49, 0x4b, 0x41, 0x4c, 0x4e, 0x50, 0x43, 0x50, 0x41, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4d, 0x4f, 0x49, 0x4b, 0x41, 0x4c, 0x4e, 0x50, 0x43, 0x50,
	0x41, 0x12, 0x21, 0x0a, 0x0b, 0x50, 0x4c, 0x4c, 0x49, 0x49, 0x50, 0x4a, 0x49, 0x46, 0x4f, 0x47,
	0x18, 0x95, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x50, 0x4c, 0x4c, 0x49, 0x49, 0x50, 0x4a,
	0x49, 0x46, 0x4f, 0x47, 0x12, 0x24, 0x0a, 0x0e, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x67, 0x65,
	0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x6f,
	0x67, 0x75, 0x65, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x0b, 0x41, 0x47,
	0x50, 0x4f, 0x4a, 0x4a, 0x4a, 0x4b, 0x4b, 0x4a, 0x49, 0x18, 0xae, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x41, 0x47, 0x50, 0x4f, 0x4a, 0x4a, 0x4a, 0x4b, 0x4b, 0x4a, 0x49, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e,
	0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerLoginCsReq_proto_rawDescOnce sync.Once
	file_PlayerLoginCsReq_proto_rawDescData = file_PlayerLoginCsReq_proto_rawDesc
)

func file_PlayerLoginCsReq_proto_rawDescGZIP() []byte {
	file_PlayerLoginCsReq_proto_rawDescOnce.Do(func() {
		file_PlayerLoginCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerLoginCsReq_proto_rawDescData)
	})
	return file_PlayerLoginCsReq_proto_rawDescData
}

var file_PlayerLoginCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerLoginCsReq_proto_goTypes = []interface{}{
	(*PlayerLoginCsReq)(nil), // 0: PlayerLoginCsReq
	(PlatformType)(0),        // 1: PlatformType
	(*EDCFBAGKIGN)(nil),      // 2: EDCFBAGKIGN
	(LanguageType)(0),        // 3: LanguageType
}
var file_PlayerLoginCsReq_proto_depIdxs = []int32{
	1, // 0: PlayerLoginCsReq.platform:type_name -> PlatformType
	2, // 1: PlayerLoginCsReq.EOLCNDBAELO:type_name -> EDCFBAGKIGN
	3, // 2: PlayerLoginCsReq.language:type_name -> LanguageType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_PlayerLoginCsReq_proto_init() }
func file_PlayerLoginCsReq_proto_init() {
	if File_PlayerLoginCsReq_proto != nil {
		return
	}
	file_LanguageType_proto_init()
	file_PlatformType_proto_init()
	file_EDCFBAGKIGN_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerLoginCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerLoginCsReq); i {
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
			RawDescriptor: file_PlayerLoginCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerLoginCsReq_proto_goTypes,
		DependencyIndexes: file_PlayerLoginCsReq_proto_depIdxs,
		MessageInfos:      file_PlayerLoginCsReq_proto_msgTypes,
	}.Build()
	File_PlayerLoginCsReq_proto = out.File
	file_PlayerLoginCsReq_proto_rawDesc = nil
	file_PlayerLoginCsReq_proto_goTypes = nil
	file_PlayerLoginCsReq_proto_depIdxs = nil
}

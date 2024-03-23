// protos/issuer.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: protos/issuer.proto

package protos

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

type MsgIssueVC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Did   string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	Nonce string `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ReqVP string `protobuf:"bytes,3,opt,name=reqVP,proto3" json:"reqVP,omitempty"`
}

func (x *MsgIssueVC) Reset() {
	*x = MsgIssueVC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_issuer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgIssueVC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgIssueVC) ProtoMessage() {}

func (x *MsgIssueVC) ProtoReflect() protoreflect.Message {
	mi := &file_protos_issuer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgIssueVC.ProtoReflect.Descriptor instead.
func (*MsgIssueVC) Descriptor() ([]byte, []int) {
	return file_protos_issuer_proto_rawDescGZIP(), []int{0}
}

func (x *MsgIssueVC) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *MsgIssueVC) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *MsgIssueVC) GetReqVP() string {
	if x != nil {
		return x.ReqVP
	}
	return ""
}

type MsgIssueVCResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Did   string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	Nonce string `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Vc    string `protobuf:"bytes,3,opt,name=vc,proto3" json:"vc,omitempty"`
}

func (x *MsgIssueVCResponse) Reset() {
	*x = MsgIssueVCResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_issuer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgIssueVCResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgIssueVCResponse) ProtoMessage() {}

func (x *MsgIssueVCResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_issuer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgIssueVCResponse.ProtoReflect.Descriptor instead.
func (*MsgIssueVCResponse) Descriptor() ([]byte, []int) {
	return file_protos_issuer_proto_rawDescGZIP(), []int{1}
}

func (x *MsgIssueVCResponse) GetDid() string {
	if x != nil {
		return x.Did
	}
	return ""
}

func (x *MsgIssueVCResponse) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *MsgIssueVCResponse) GetVc() string {
	if x != nil {
		return x.Vc
	}
	return ""
}

var File_protos_issuer_proto protoreflect.FileDescriptor

var file_protos_issuer_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x22, 0x4a, 0x0a,
	0x0a, 0x4d, 0x73, 0x67, 0x49, 0x73, 0x73, 0x75, 0x65, 0x56, 0x43, 0x12, 0x10, 0x0a, 0x03, 0x64,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f,
	0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x71, 0x56, 0x50, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x71, 0x56, 0x50, 0x22, 0x4c, 0x0a, 0x12, 0x4d, 0x73, 0x67,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x56, 0x43, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x64, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x76, 0x63, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x76, 0x63, 0x32, 0x51, 0x0a, 0x0c, 0x53, 0x69, 0x6d, 0x70, 0x6c,
	0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x0d, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x56, 0x43, 0x12, 0x12, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x72, 0x2e, 0x4d, 0x73, 0x67, 0x49, 0x73, 0x73, 0x75, 0x65, 0x56, 0x43, 0x1a, 0x1a, 0x2e, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x72, 0x2e, 0x4d, 0x73, 0x67, 0x49, 0x73, 0x73, 0x75, 0x65, 0x56, 0x43,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x73, 0x73,
	0x69, 0x6b, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protos_issuer_proto_rawDescOnce sync.Once
	file_protos_issuer_proto_rawDescData = file_protos_issuer_proto_rawDesc
)

func file_protos_issuer_proto_rawDescGZIP() []byte {
	file_protos_issuer_proto_rawDescOnce.Do(func() {
		file_protos_issuer_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_issuer_proto_rawDescData)
	})
	return file_protos_issuer_proto_rawDescData
}

var file_protos_issuer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_issuer_proto_goTypes = []interface{}{
	(*MsgIssueVC)(nil),         // 0: issuer.MsgIssueVC
	(*MsgIssueVCResponse)(nil), // 1: issuer.MsgIssueVCResponse
}
var file_protos_issuer_proto_depIdxs = []int32{
	0, // 0: issuer.SimpleIssuer.IssueSimpleVC:input_type -> issuer.MsgIssueVC
	1, // 1: issuer.SimpleIssuer.IssueSimpleVC:output_type -> issuer.MsgIssueVCResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_issuer_proto_init() }
func file_protos_issuer_proto_init() {
	if File_protos_issuer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_issuer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgIssueVC); i {
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
		file_protos_issuer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgIssueVCResponse); i {
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
			RawDescriptor: file_protos_issuer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_issuer_proto_goTypes,
		DependencyIndexes: file_protos_issuer_proto_depIdxs,
		MessageInfos:      file_protos_issuer_proto_msgTypes,
	}.Build()
	File_protos_issuer_proto = out.File
	file_protos_issuer_proto_rawDesc = nil
	file_protos_issuer_proto_goTypes = nil
	file_protos_issuer_proto_depIdxs = nil
}

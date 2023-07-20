// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: mydemo.proto

package demo

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

type InMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *InMsg) Reset() {
	*x = InMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mydemo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InMsg) ProtoMessage() {}

func (x *InMsg) ProtoReflect() protoreflect.Message {
	mi := &file_mydemo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InMsg.ProtoReflect.Descriptor instead.
func (*InMsg) Descriptor() ([]byte, []int) {
	return file_mydemo_proto_rawDescGZIP(), []int{0}
}

func (x *InMsg) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type OutMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *OutMsg) Reset() {
	*x = OutMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mydemo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutMsg) ProtoMessage() {}

func (x *OutMsg) ProtoReflect() protoreflect.Message {
	mi := &file_mydemo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutMsg.ProtoReflect.Descriptor instead.
func (*OutMsg) Descriptor() ([]byte, []int) {
	return file_mydemo_proto_rawDescGZIP(), []int{1}
}

func (x *OutMsg) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_mydemo_proto protoreflect.FileDescriptor

var file_mydemo_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x79, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6d, 0x79, 0x64, 0x65, 0x6d, 0x6f, 0x22, 0x19, 0x0a, 0x05, 0x49, 0x6e, 0x4d, 0x73, 0x67, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x22, 0x1a, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x31, 0x0a,
	0x04, 0x44, 0x65, 0x6d, 0x6f, 0x12, 0x29, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x12, 0x0d, 0x2e, 0x6d, 0x79, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x49, 0x6e, 0x4d, 0x73, 0x67,
	0x1a, 0x0e, 0x2e, 0x6d, 0x79, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x4f, 0x75, 0x74, 0x4d, 0x73, 0x67,
	0x42, 0x0f, 0x5a, 0x0d, 0x64, 0x65, 0x6d, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x6d,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mydemo_proto_rawDescOnce sync.Once
	file_mydemo_proto_rawDescData = file_mydemo_proto_rawDesc
)

func file_mydemo_proto_rawDescGZIP() []byte {
	file_mydemo_proto_rawDescOnce.Do(func() {
		file_mydemo_proto_rawDescData = protoimpl.X.CompressGZIP(file_mydemo_proto_rawDescData)
	})
	return file_mydemo_proto_rawDescData
}

var file_mydemo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mydemo_proto_goTypes = []interface{}{
	(*InMsg)(nil),  // 0: mydemo.InMsg
	(*OutMsg)(nil), // 1: mydemo.OutMsg
}
var file_mydemo_proto_depIdxs = []int32{
	0, // 0: mydemo.Demo.SayHello:input_type -> mydemo.InMsg
	1, // 1: mydemo.Demo.SayHello:output_type -> mydemo.OutMsg
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mydemo_proto_init() }
func file_mydemo_proto_init() {
	if File_mydemo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mydemo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InMsg); i {
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
		file_mydemo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutMsg); i {
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
			RawDescriptor: file_mydemo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mydemo_proto_goTypes,
		DependencyIndexes: file_mydemo_proto_depIdxs,
		MessageInfos:      file_mydemo_proto_msgTypes,
	}.Build()
	File_mydemo_proto = out.File
	file_mydemo_proto_rawDesc = nil
	file_mydemo_proto_goTypes = nil
	file_mydemo_proto_depIdxs = nil
}

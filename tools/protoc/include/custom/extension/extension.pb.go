// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: custom/extension/extension.proto

package extension

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HttpMethodOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// if true, means is pure http method, do not generate grpc code, and force register http handler.
	// default is false, will generate both grpc and http code.
	Pure bool `protobuf:"varint,1,opt,name=pure,proto3" json:"pure,omitempty"`
}

func (x *HttpMethodOption) Reset() {
	*x = HttpMethodOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_custom_extension_extension_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpMethodOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpMethodOption) ProtoMessage() {}

func (x *HttpMethodOption) ProtoReflect() protoreflect.Message {
	mi := &file_custom_extension_extension_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpMethodOption.ProtoReflect.Descriptor instead.
func (*HttpMethodOption) Descriptor() ([]byte, []int) {
	return file_custom_extension_extension_proto_rawDescGZIP(), []int{0}
}

func (x *HttpMethodOption) GetPure() bool {
	if x != nil {
		return x.Pure
	}
	return false
}

var file_custom_extension_extension_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*HttpMethodOption)(nil),
		Field:         1001,
		Name:          "custom.extension.http",
		Tag:           "bytes,1001,opt,name=http",
		Filename:      "custom/extension/extension.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional custom.extension.HttpMethodOption http = 1001;
	E_Http = &file_custom_extension_extension_proto_extTypes[0]
)

var File_custom_extension_extension_proto protoreflect.FileDescriptor

var file_custom_extension_extension_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x10, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x75,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x70, 0x75, 0x72, 0x65, 0x3a, 0x57,
	0x0a, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x74, 0x6f,
	0x6f, 0x6c, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x69, 0x6e, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_custom_extension_extension_proto_rawDescOnce sync.Once
	file_custom_extension_extension_proto_rawDescData = file_custom_extension_extension_proto_rawDesc
)

func file_custom_extension_extension_proto_rawDescGZIP() []byte {
	file_custom_extension_extension_proto_rawDescOnce.Do(func() {
		file_custom_extension_extension_proto_rawDescData = protoimpl.X.CompressGZIP(file_custom_extension_extension_proto_rawDescData)
	})
	return file_custom_extension_extension_proto_rawDescData
}

var file_custom_extension_extension_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_custom_extension_extension_proto_goTypes = []interface{}{
	(*HttpMethodOption)(nil),           // 0: custom.extension.HttpMethodOption
	(*descriptorpb.MethodOptions)(nil), // 1: google.protobuf.MethodOptions
}
var file_custom_extension_extension_proto_depIdxs = []int32{
	1, // 0: custom.extension.http:extendee -> google.protobuf.MethodOptions
	0, // 1: custom.extension.http:type_name -> custom.extension.HttpMethodOption
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_custom_extension_extension_proto_init() }
func file_custom_extension_extension_proto_init() {
	if File_custom_extension_extension_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_custom_extension_extension_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpMethodOption); i {
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
			RawDescriptor: file_custom_extension_extension_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_custom_extension_extension_proto_goTypes,
		DependencyIndexes: file_custom_extension_extension_proto_depIdxs,
		MessageInfos:      file_custom_extension_extension_proto_msgTypes,
		ExtensionInfos:    file_custom_extension_extension_proto_extTypes,
	}.Build()
	File_custom_extension_extension_proto = out.File
	file_custom_extension_extension_proto_rawDesc = nil
	file_custom_extension_extension_proto_goTypes = nil
	file_custom_extension_extension_proto_depIdxs = nil
}
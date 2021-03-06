// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.0--rc1
// source: api/baseapp/interface/v1/baseapp_interface_error.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type BaseappInterfaceError int32

const (
	BaseappInterfaceError_INFO_NOT_FOUND              BaseappInterfaceError = 0
	BaseappInterfaceError_CONTENT_MISSING             BaseappInterfaceError = 1
	BaseappInterfaceError_INVALID_ACCOUNT_OR_PASSWORD BaseappInterfaceError = 2
	BaseappInterfaceError_UNAUTHORIZED                BaseappInterfaceError = 3
)

// Enum value maps for BaseappInterfaceError.
var (
	BaseappInterfaceError_name = map[int32]string{
		0: "INFO_NOT_FOUND",
		1: "CONTENT_MISSING",
		2: "INVALID_ACCOUNT_OR_PASSWORD",
		3: "UNAUTHORIZED",
	}
	BaseappInterfaceError_value = map[string]int32{
		"INFO_NOT_FOUND":              0,
		"CONTENT_MISSING":             1,
		"INVALID_ACCOUNT_OR_PASSWORD": 2,
		"UNAUTHORIZED":                3,
	}
)

func (x BaseappInterfaceError) Enum() *BaseappInterfaceError {
	p := new(BaseappInterfaceError)
	*p = x
	return p
}

func (x BaseappInterfaceError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BaseappInterfaceError) Descriptor() protoreflect.EnumDescriptor {
	return file_api_baseapp_interface_v1_baseapp_interface_error_proto_enumTypes[0].Descriptor()
}

func (BaseappInterfaceError) Type() protoreflect.EnumType {
	return &file_api_baseapp_interface_v1_baseapp_interface_error_proto_enumTypes[0]
}

func (x BaseappInterfaceError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BaseappInterfaceError.Descriptor instead.
func (BaseappInterfaceError) EnumDescriptor() ([]byte, []int) {
	return file_api_baseapp_interface_v1_baseapp_interface_error_proto_rawDescGZIP(), []int{0}
}

var File_api_baseapp_interface_v1_baseapp_interface_error_proto protoreflect.FileDescriptor

var file_api_baseapp_interface_v1_baseapp_interface_error_proto_rawDesc = []byte{
	0x0a, 0x36, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x61, 0x70, 0x70, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x61,
	0x70, 0x70, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61,
	0x73, 0x65, 0x61, 0x70, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x91, 0x01, 0x0a, 0x15, 0x42, 0x61, 0x73, 0x65,
	0x61, 0x70, 0x70, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x18, 0x0a, 0x0e, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f,
	0x55, 0x4e, 0x44, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x43,
	0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x01,
	0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x25, 0x0a, 0x1b, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4f, 0x52, 0x5f, 0x50, 0x41, 0x53,
	0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0xf5, 0x03, 0x12, 0x16, 0x0a,
	0x0c, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x03, 0x1a,
	0x04, 0xa8, 0x45, 0x93, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x46, 0x0a, 0x18, 0x61,
	0x70, 0x69, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x61, 0x70, 0x70, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x28, 0x62, 0x61, 0x73, 0x65, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x61, 0x70, 0x70, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_baseapp_interface_v1_baseapp_interface_error_proto_rawDescOnce sync.Once
	file_api_baseapp_interface_v1_baseapp_interface_error_proto_rawDescData = file_api_baseapp_interface_v1_baseapp_interface_error_proto_rawDesc
)

func file_api_baseapp_interface_v1_baseapp_interface_error_proto_rawDescGZIP() []byte {
	file_api_baseapp_interface_v1_baseapp_interface_error_proto_rawDescOnce.Do(func() {
		file_api_baseapp_interface_v1_baseapp_interface_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_baseapp_interface_v1_baseapp_interface_error_proto_rawDescData)
	})
	return file_api_baseapp_interface_v1_baseapp_interface_error_proto_rawDescData
}

var file_api_baseapp_interface_v1_baseapp_interface_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_baseapp_interface_v1_baseapp_interface_error_proto_goTypes = []interface{}{
	(BaseappInterfaceError)(0), // 0: api.baseapp.interface.v1.BaseappInterfaceError
}
var file_api_baseapp_interface_v1_baseapp_interface_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_baseapp_interface_v1_baseapp_interface_error_proto_init() }
func file_api_baseapp_interface_v1_baseapp_interface_error_proto_init() {
	if File_api_baseapp_interface_v1_baseapp_interface_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_baseapp_interface_v1_baseapp_interface_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_baseapp_interface_v1_baseapp_interface_error_proto_goTypes,
		DependencyIndexes: file_api_baseapp_interface_v1_baseapp_interface_error_proto_depIdxs,
		EnumInfos:         file_api_baseapp_interface_v1_baseapp_interface_error_proto_enumTypes,
	}.Build()
	File_api_baseapp_interface_v1_baseapp_interface_error_proto = out.File
	file_api_baseapp_interface_v1_baseapp_interface_error_proto_rawDesc = nil
	file_api_baseapp_interface_v1_baseapp_interface_error_proto_goTypes = nil
	file_api_baseapp_interface_v1_baseapp_interface_error_proto_depIdxs = nil
}

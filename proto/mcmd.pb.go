// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: mcmd.proto

//option go_package = "~/go/src/go_game_server/proto/mcmd";

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MessageCmd int32

const (
	MessageCmd_GET_USER_LIST MessageCmd = 0
)

// Enum value maps for MessageCmd.
var (
	MessageCmd_name = map[int32]string{
		0: "GET_USER_LIST",
	}
	MessageCmd_value = map[string]int32{
		"GET_USER_LIST": 0,
	}
)

func (x MessageCmd) Enum() *MessageCmd {
	p := new(MessageCmd)
	*p = x
	return p
}

func (x MessageCmd) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageCmd) Descriptor() protoreflect.EnumDescriptor {
	return file_mcmd_proto_enumTypes[0].Descriptor()
}

func (MessageCmd) Type() protoreflect.EnumType {
	return &file_mcmd_proto_enumTypes[0]
}

func (x MessageCmd) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageCmd.Descriptor instead.
func (MessageCmd) EnumDescriptor() ([]byte, []int) {
	return file_mcmd_proto_rawDescGZIP(), []int{0}
}

var File_mcmd_proto protoreflect.FileDescriptor

var file_mcmd_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x63, 0x6d, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0x1f, 0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6d,
	0x64, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x45, 0x54, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4c, 0x49,
	0x53, 0x54, 0x10, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mcmd_proto_rawDescOnce sync.Once
	file_mcmd_proto_rawDescData = file_mcmd_proto_rawDesc
)

func file_mcmd_proto_rawDescGZIP() []byte {
	file_mcmd_proto_rawDescOnce.Do(func() {
		file_mcmd_proto_rawDescData = protoimpl.X.CompressGZIP(file_mcmd_proto_rawDescData)
	})
	return file_mcmd_proto_rawDescData
}

var file_mcmd_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mcmd_proto_goTypes = []interface{}{
	(MessageCmd)(0), // 0: proto.MessageCmd
}
var file_mcmd_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mcmd_proto_init() }
func file_mcmd_proto_init() {
	if File_mcmd_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mcmd_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mcmd_proto_goTypes,
		DependencyIndexes: file_mcmd_proto_depIdxs,
		EnumInfos:         file_mcmd_proto_enumTypes,
	}.Build()
	File_mcmd_proto = out.File
	file_mcmd_proto_rawDesc = nil
	file_mcmd_proto_goTypes = nil
	file_mcmd_proto_depIdxs = nil
}

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mcmd.proto

//option go_package = "~/go/src/go_game_server/proto/mcmd";

package message

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MessageCmd int32

const (
	MessageCmd_GET_USER_LIST MessageCmd = 0
)

var MessageCmd_name = map[int32]string{
	0: "GET_USER_LIST",
}

var MessageCmd_value = map[string]int32{
	"GET_USER_LIST": 0,
}

func (x MessageCmd) String() string {
	return proto.EnumName(MessageCmd_name, int32(x))
}

func (MessageCmd) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_613abd029ef44132, []int{0}
}

func init() {
	proto.RegisterEnum("message.MessageCmd", MessageCmd_name, MessageCmd_value)
}

func init() { proto.RegisterFile("mcmd.proto", fileDescriptor_613abd029ef44132) }

var fileDescriptor_613abd029ef44132 = []byte{
	// 98 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x4d, 0xce, 0x4d,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5,
	0x92, 0xe7, 0xe2, 0xf2, 0x85, 0x30, 0x9d, 0x73, 0x53, 0x84, 0x04, 0xb9, 0x78, 0xdd, 0x5d, 0x43,
	0xe2, 0x43, 0x83, 0x5d, 0x83, 0xe2, 0x7d, 0x3c, 0x83, 0x43, 0x04, 0x18, 0x9c, 0x04, 0x4e, 0x3c,
	0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x19, 0x8f, 0xe5, 0x18, 0x92,
	0xd8, 0xc0, 0x46, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x47, 0x41, 0xe9, 0xd6, 0x50, 0x00,
	0x00, 0x00,
}

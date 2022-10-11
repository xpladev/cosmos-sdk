// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/msg/v1/msg.proto

package msgservice

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

var E_Service = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         11110000,
	Name:          "cosmos.msg.v1.service",
	Tag:           "varint,11110000,opt,name=service",
	Filename:      "cosmos/msg/v1/msg.proto",
}

var E_Signer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.MessageOptions)(nil),
	ExtensionType: ([]string)(nil),
	Field:         11110000,
	Name:          "cosmos.msg.v1.signer",
	Tag:           "bytes,11110000,rep,name=signer",
	Filename:      "cosmos/msg/v1/msg.proto",
}

var E_LegacyAminoName = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         11110001,
	Name:          "cosmos.msg.v1.legacy_amino_name",
	Tag:           "bytes,11110001,opt,name=legacy_amino_name",
	Filename:      "cosmos/msg/v1/msg.proto",
}

func init() {
	proto.RegisterExtension(E_Service)
	proto.RegisterExtension(E_Signer)
	proto.RegisterExtension(E_LegacyAminoName)
}

func init() { proto.RegisterFile("cosmos/msg/v1/msg.proto", fileDescriptor_5c08b83ea858d203) }

var fileDescriptor_5c08b83ea858d203 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0xcf, 0x2d, 0x4e, 0xd7, 0x2f, 0x33, 0x04, 0x51, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0xbc, 0x10, 0x09, 0x3d, 0x90, 0x48, 0x99, 0xa1, 0x94, 0x42, 0x7a, 0x7e, 0x7e, 0x7a,
	0x4e, 0xaa, 0x3e, 0x58, 0x32, 0xa9, 0x34, 0x4d, 0x3f, 0x25, 0xb5, 0x38, 0xb9, 0x28, 0xb3, 0xa0,
	0x24, 0xbf, 0x08, 0xa2, 0xc1, 0xca, 0x86, 0x8b, 0xbd, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55,
	0x48, 0x5e, 0x0f, 0xa2, 0x5a, 0x0f, 0xa6, 0x5a, 0x2f, 0x18, 0x22, 0xe3, 0x5f, 0x50, 0x92, 0x99,
	0x9f, 0x57, 0x2c, 0xf1, 0xa1, 0x67, 0x19, 0xab, 0x02, 0xa3, 0x06, 0x47, 0x10, 0x4c, 0x8b, 0x95,
	0x15, 0x17, 0x5b, 0x71, 0x66, 0x7a, 0x5e, 0x6a, 0x11, 0x16, 0xcd, 0xbe, 0xa9, 0xc5, 0xc5, 0x89,
	0xe9, 0xa8, 0x9a, 0x99, 0x35, 0x38, 0x83, 0xa0, 0x3a, 0xac, 0xfc, 0xb8, 0x04, 0x73, 0x52, 0xd3,
	0x13, 0x93, 0x2b, 0xe3, 0x13, 0x73, 0x33, 0xf3, 0xf2, 0xe3, 0xf3, 0x12, 0x73, 0x53, 0x09, 0x1b,
	0xf3, 0x11, 0xe2, 0x06, 0xce, 0x20, 0x7e, 0x88, 0x66, 0x47, 0x90, 0x5e, 0xbf, 0xc4, 0xdc, 0x54,
	0x27, 0xf7, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2,
	0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x4d, 0xcf, 0x2c,
	0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x87, 0x06, 0x1c, 0x84, 0xd2, 0x2d, 0x4e, 0xc9,
	0xd6, 0x2f, 0xa9, 0x2c, 0x48, 0x05, 0x87, 0x24, 0xd4, 0x53, 0x49, 0x6c, 0x60, 0xbb, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xf8, 0x04, 0x39, 0x65, 0x01, 0x00, 0x00,
}

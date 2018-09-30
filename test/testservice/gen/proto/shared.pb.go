// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shared.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	shared.proto
	testservice.proto

It has these top-level messages:
	CommonValues
	SayHelloRequest
	SayHelloResponse
	TestJsonRequest
	TestJsonResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type CommonValues struct {
	SomeId int64 `protobuf:"varint,1,opt,name=someId" json:"someId,omitempty"`
}

func (m *CommonValues) Reset()                    { *m = CommonValues{} }
func (m *CommonValues) String() string            { return proto1.CompactTextString(m) }
func (*CommonValues) ProtoMessage()               {}
func (*CommonValues) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CommonValues) GetSomeId() int64 {
	if m != nil {
		return m.SomeId
	}
	return 0
}

func init() {
	proto1.RegisterType((*CommonValues)(nil), "proto.CommonValues")
}

func init() { proto1.RegisterFile("shared.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 81 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xce, 0x48, 0x2c,
	0x4a, 0x4d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x6a, 0x5c, 0x3c,
	0xce, 0xf9, 0xb9, 0xb9, 0xf9, 0x79, 0x61, 0x89, 0x39, 0xa5, 0xa9, 0xc5, 0x42, 0x62, 0x5c, 0x6c,
	0xc5, 0xf9, 0xb9, 0xa9, 0x9e, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x50, 0x5e, 0x12,
	0x1b, 0x58, 0xb9, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x28, 0xbf, 0xd7, 0x45, 0x00, 0x00,
	0x00,
}

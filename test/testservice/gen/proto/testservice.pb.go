// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testservice.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SayHelloRequest struct {
	Values       *CommonValues `protobuf:"bytes,1,opt,name=values" json:"values,omitempty"`
	YourName     string        `protobuf:"bytes,2,opt,name=yourName" json:"yourName,omitempty"`
	Int64Value   int64         `protobuf:"varint,3,opt,name=int64Value" json:"int64Value,omitempty"`
	BoolValue    bool          `protobuf:"varint,4,opt,name=boolValue" json:"boolValue,omitempty"`
	Float64Value float64       `protobuf:"fixed64,5,opt,name=float64Value" json:"float64Value,omitempty"`
	Uint64Value  uint64        `protobuf:"varint,6,opt,name=uint64Value" json:"uint64Value,omitempty"`
	StringList   []string      `protobuf:"bytes,7,rep,name=stringList" json:"stringList,omitempty"`
	Int64List    []int64       `protobuf:"varint,8,rep,packed,name=int64List" json:"int64List,omitempty"`
	BoolList     []bool        `protobuf:"varint,9,rep,packed,name=boolList" json:"boolList,omitempty"`
	DoubleList   []float64     `protobuf:"fixed64,10,rep,packed,name=doubleList" json:"doubleList,omitempty"`
	Uint64List   []uint64      `protobuf:"varint,11,rep,packed,name=uint64List" json:"uint64List,omitempty"`
}

func (m *SayHelloRequest) Reset()                    { *m = SayHelloRequest{} }
func (m *SayHelloRequest) String() string            { return proto1.CompactTextString(m) }
func (*SayHelloRequest) ProtoMessage()               {}
func (*SayHelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *SayHelloRequest) GetValues() *CommonValues {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *SayHelloRequest) GetYourName() string {
	if m != nil {
		return m.YourName
	}
	return ""
}

func (m *SayHelloRequest) GetInt64Value() int64 {
	if m != nil {
		return m.Int64Value
	}
	return 0
}

func (m *SayHelloRequest) GetBoolValue() bool {
	if m != nil {
		return m.BoolValue
	}
	return false
}

func (m *SayHelloRequest) GetFloat64Value() float64 {
	if m != nil {
		return m.Float64Value
	}
	return 0
}

func (m *SayHelloRequest) GetUint64Value() uint64 {
	if m != nil {
		return m.Uint64Value
	}
	return 0
}

func (m *SayHelloRequest) GetStringList() []string {
	if m != nil {
		return m.StringList
	}
	return nil
}

func (m *SayHelloRequest) GetInt64List() []int64 {
	if m != nil {
		return m.Int64List
	}
	return nil
}

func (m *SayHelloRequest) GetBoolList() []bool {
	if m != nil {
		return m.BoolList
	}
	return nil
}

func (m *SayHelloRequest) GetDoubleList() []float64 {
	if m != nil {
		return m.DoubleList
	}
	return nil
}

func (m *SayHelloRequest) GetUint64List() []uint64 {
	if m != nil {
		return m.Uint64List
	}
	return nil
}

type SayHelloResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *SayHelloResponse) Reset()                    { *m = SayHelloResponse{} }
func (m *SayHelloResponse) String() string            { return proto1.CompactTextString(m) }
func (*SayHelloResponse) ProtoMessage()               {}
func (*SayHelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *SayHelloResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type TestJsonRequest struct {
}

func (m *TestJsonRequest) Reset()                    { *m = TestJsonRequest{} }
func (m *TestJsonRequest) String() string            { return proto1.CompactTextString(m) }
func (*TestJsonRequest) ProtoMessage()               {}
func (*TestJsonRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type TestJsonResponse struct {
}

func (m *TestJsonResponse) Reset()                    { *m = TestJsonResponse{} }
func (m *TestJsonResponse) String() string            { return proto1.CompactTextString(m) }
func (*TestJsonResponse) ProtoMessage()               {}
func (*TestJsonResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func init() {
	proto1.RegisterType((*SayHelloRequest)(nil), "proto.SayHelloRequest")
	proto1.RegisterType((*SayHelloResponse)(nil), "proto.SayHelloResponse")
	proto1.RegisterType((*TestJsonRequest)(nil), "proto.TestJsonRequest")
	proto1.RegisterType((*TestJsonResponse)(nil), "proto.TestJsonResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TestService service

type TestServiceClient interface {
	SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error)
	TestJson(ctx context.Context, in *TestJsonRequest, opts ...grpc.CallOption) (*TestJsonResponse, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error) {
	out := new(SayHelloResponse)
	err := grpc.Invoke(ctx, "/proto.TestService/sayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) TestJson(ctx context.Context, in *TestJsonRequest, opts ...grpc.CallOption) (*TestJsonResponse, error) {
	out := new(TestJsonResponse)
	err := grpc.Invoke(ctx, "/proto.TestService/testJson", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TestService service

type TestServiceServer interface {
	SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error)
	TestJson(context.Context, *TestJsonRequest) (*TestJsonResponse, error)
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TestService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).SayHello(ctx, req.(*SayHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_TestJson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestJsonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).TestJson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TestService/TestJson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).TestJson(ctx, req.(*TestJsonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sayHello",
			Handler:    _TestService_SayHello_Handler,
		},
		{
			MethodName: "testJson",
			Handler:    _TestService_TestJson_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testservice.proto",
}

func init() { proto1.RegisterFile("testservice.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xbb, 0xdd, 0xfe, 0x49, 0x26, 0x85, 0xb6, 0xfb, 0xc2, 0xfb, 0x86, 0xf2, 0x22, 0x4b,
	0x4e, 0x0b, 0x4a, 0x0f, 0x55, 0xbc, 0x79, 0xf2, 0x22, 0x22, 0x1e, 0xb6, 0xe2, 0x3d, 0xb5, 0x6b,
	0x0d, 0x24, 0xd9, 0x9a, 0xd9, 0x14, 0xfa, 0x19, 0xfc, 0xc2, 0x1e, 0x65, 0x37, 0x31, 0x89, 0xf5,
	0x14, 0xe6, 0xf7, 0xcc, 0x33, 0xf3, 0x64, 0x07, 0xe6, 0x46, 0xa1, 0x41, 0x55, 0x1c, 0x92, 0x17,
	0xb5, 0xdc, 0x17, 0xda, 0x68, 0x36, 0x74, 0x9f, 0xc5, 0x04, 0xdf, 0xe2, 0x42, 0x6d, 0x2b, 0x18,
	0x7d, 0xf6, 0x61, 0xba, 0x8e, 0x8f, 0x77, 0x2a, 0x4d, 0xb5, 0x54, 0xef, 0xa5, 0x42, 0xc3, 0xce,
	0x61, 0x74, 0x88, 0xd3, 0x52, 0x61, 0x48, 0x38, 0x11, 0xc1, 0xea, 0x4f, 0xd5, 0xbb, 0xbc, 0xd5,
	0x59, 0xa6, 0xf3, 0x67, 0x27, 0xc9, 0xba, 0x85, 0x2d, 0xc0, 0x3b, 0xea, 0xb2, 0x78, 0x8c, 0x33,
	0x15, 0xf6, 0x39, 0x11, 0xbe, 0x6c, 0x6a, 0x76, 0x06, 0x90, 0xe4, 0xe6, 0xfa, 0xca, 0x59, 0x42,
	0xca, 0x89, 0xa0, 0xb2, 0x43, 0xd8, 0x7f, 0xf0, 0x37, 0x5a, 0xa7, 0x95, 0x3c, 0xe0, 0x44, 0x78,
	0xb2, 0x05, 0x2c, 0x82, 0xc9, 0x6b, 0xaa, 0xe3, 0xc6, 0x3f, 0xe4, 0x44, 0x10, 0xf9, 0x83, 0x31,
	0x0e, 0x41, 0xd9, 0x59, 0x31, 0xe2, 0x44, 0x0c, 0x64, 0x17, 0xd9, 0x0c, 0x68, 0x8a, 0x24, 0xdf,
	0x3d, 0x24, 0x68, 0xc2, 0x31, 0xa7, 0xc2, 0x97, 0x1d, 0x62, 0x33, 0xb8, 0x6e, 0x27, 0x7b, 0x9c,
	0x0a, 0x2a, 0x5b, 0x60, 0xff, 0xce, 0x06, 0x72, 0xa2, 0xcf, 0xa9, 0xf0, 0x64, 0x53, 0xdb, 0xc9,
	0x5b, 0x5d, 0x6e, 0x52, 0xe5, 0x54, 0xe0, 0x54, 0x10, 0xd9, 0x21, 0x56, 0x2f, 0xdb, 0xd1, 0x01,
	0xa7, 0x62, 0x20, 0x3b, 0x24, 0xba, 0x80, 0x59, 0xfb, 0xf2, 0xb8, 0xd7, 0x39, 0x2a, 0x16, 0xc2,
	0x38, 0x53, 0x88, 0xf1, 0x4e, 0xb9, 0xb7, 0xf7, 0xe5, 0x77, 0x19, 0xcd, 0x61, 0xfa, 0xa4, 0xd0,
	0xdc, 0xa3, 0xce, 0xeb, 0x3b, 0x45, 0x0c, 0x66, 0x2d, 0xaa, 0x06, 0xac, 0x3e, 0x08, 0x04, 0x16,
	0xae, 0xab, 0xd3, 0xb3, 0x1b, 0xf0, 0xb0, 0x5e, 0xc2, 0xfe, 0xd6, 0x77, 0x3c, 0xb9, 0xf7, 0xe2,
	0xdf, 0x2f, 0x5e, 0x0d, 0x8b, 0x7a, 0xd6, 0x6e, 0xea, 0x15, 0x8d, 0xfd, 0x24, 0x46, 0x63, 0x3f,
	0xcd, 0x12, 0xf5, 0x36, 0x23, 0xa7, 0x5c, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x7e, 0xdc, 0xca,
	0x79, 0x8e, 0x02, 0x00, 0x00,
}

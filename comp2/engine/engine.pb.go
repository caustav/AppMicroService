// Code generated by protoc-gen-go.
// source: engine.proto
// DO NOT EDIT!

/*
Package engine is a generated protocol buffer package.

It is generated from these files:
	engine.proto

It has these top-level messages:
	EngineRequest
	EngineResponse
*/
package engine

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EngineRequest struct {
	UniqueNumber string `protobuf:"bytes,1,opt,name=uniqueNumber" json:"uniqueNumber,omitempty"`
	Url          string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
}

func (m *EngineRequest) Reset()                    { *m = EngineRequest{} }
func (m *EngineRequest) String() string            { return proto.CompactTextString(m) }
func (*EngineRequest) ProtoMessage()               {}
func (*EngineRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EngineRequest) GetUniqueNumber() string {
	if m != nil {
		return m.UniqueNumber
	}
	return ""
}

func (m *EngineRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type EngineResponse struct {
}

func (m *EngineResponse) Reset()                    { *m = EngineResponse{} }
func (m *EngineResponse) String() string            { return proto.CompactTextString(m) }
func (*EngineResponse) ProtoMessage()               {}
func (*EngineResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*EngineRequest)(nil), "engine.EngineRequest")
	proto.RegisterType((*EngineResponse)(nil), "engine.EngineResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Engine service

type EngineClient interface {
	Register(ctx context.Context, in *EngineRequest, opts ...grpc.CallOption) (*EngineResponse, error)
}

type engineClient struct {
	cc *grpc.ClientConn
}

func NewEngineClient(cc *grpc.ClientConn) EngineClient {
	return &engineClient{cc}
}

func (c *engineClient) Register(ctx context.Context, in *EngineRequest, opts ...grpc.CallOption) (*EngineResponse, error) {
	out := new(EngineResponse)
	err := grpc.Invoke(ctx, "/engine.Engine/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Engine service

type EngineServer interface {
	Register(context.Context, *EngineRequest) (*EngineResponse, error)
}

func RegisterEngineServer(s *grpc.Server, srv EngineServer) {
	s.RegisterService(&_Engine_serviceDesc, srv)
}

func _Engine_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EngineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/engine.Engine/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineServer).Register(ctx, req.(*EngineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Engine_serviceDesc = grpc.ServiceDesc{
	ServiceName: "engine.Engine",
	HandlerType: (*EngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Engine_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "engine.proto",
}

func init() { proto.RegisterFile("engine.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xcd, 0x4b, 0xcf,
	0xcc, 0x4b, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x5c, 0xb9, 0x78,
	0x5d, 0xc1, 0xac, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x25, 0x2e, 0x9e, 0xd2, 0xbc,
	0xcc, 0xc2, 0xd2, 0x54, 0xbf, 0xd2, 0xdc, 0xa4, 0xd4, 0x22, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x14, 0x31, 0x21, 0x01, 0x2e, 0xe6, 0xd2, 0xa2, 0x1c, 0x09, 0x26, 0xb0, 0x14, 0x88, 0xa9,
	0x24, 0xc0, 0xc5, 0x07, 0x33, 0xa6, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0xd5, 0xc8, 0x95, 0x8b, 0x0d,
	0x22, 0x22, 0x64, 0xcd, 0xc5, 0x11, 0x94, 0x9a, 0x9e, 0x59, 0x5c, 0x92, 0x5a, 0x24, 0x24, 0xaa,
	0x07, 0x75, 0x05, 0x8a, 0xa5, 0x52, 0x62, 0xe8, 0xc2, 0x10, 0x43, 0x94, 0x18, 0x9c, 0x94, 0xb8,
	0x04, 0x33, 0xf3, 0xf5, 0xd2, 0x8b, 0x0a, 0x92, 0xf5, 0xb2, 0x93, 0xa1, 0xaa, 0x9c, 0xb8, 0x21,
	0xca, 0x02, 0x40, 0x3e, 0x09, 0x60, 0x4c, 0x62, 0x03, 0x7b, 0xc9, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0xa7, 0xc8, 0x19, 0x98, 0xe2, 0x00, 0x00, 0x00,
}
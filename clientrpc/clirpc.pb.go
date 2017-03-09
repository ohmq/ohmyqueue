// Code generated by protoc-gen-go.
// source: clirpc.proto
// DO NOT EDIT!

/*
Package clientrpc is a generated protocol buffer package.

It is generated from these files:
	clirpc.proto

It has these top-level messages:
	Req
	Msg
	Resp
	StatusCode
*/
package clientrpc

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

type Req struct {
	Offset string `protobuf:"bytes,1,opt,name=offset" json:"offset,omitempty"`
}

func (m *Req) Reset()                    { *m = Req{} }
func (m *Req) String() string            { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()               {}
func (*Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Req) GetOffset() string {
	if m != nil {
		return m.Offset
	}
	return ""
}

type Msg struct {
	Offset string `protobuf:"bytes,1,opt,name=offset" json:"offset,omitempty"`
	Body   string `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
}

func (m *Msg) Reset()                    { *m = Msg{} }
func (m *Msg) String() string            { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()               {}
func (*Msg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Msg) GetOffset() string {
	if m != nil {
		return m.Offset
	}
	return ""
}

func (m *Msg) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type Resp struct {
	Offset string `protobuf:"bytes,1,opt,name=offset" json:"offset,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *Resp) Reset()                    { *m = Resp{} }
func (m *Resp) String() string            { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()               {}
func (*Resp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Resp) GetOffset() string {
	if m != nil {
		return m.Offset
	}
	return ""
}

func (m *Resp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type StatusCode struct {
	Code   int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Offset string `protobuf:"bytes,2,opt,name=offset" json:"offset,omitempty"`
}

func (m *StatusCode) Reset()                    { *m = StatusCode{} }
func (m *StatusCode) String() string            { return proto.CompactTextString(m) }
func (*StatusCode) ProtoMessage()               {}
func (*StatusCode) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StatusCode) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *StatusCode) GetOffset() string {
	if m != nil {
		return m.Offset
	}
	return ""
}

func init() {
	proto.RegisterType((*Req)(nil), "clientrpc.Req")
	proto.RegisterType((*Msg)(nil), "clientrpc.Msg")
	proto.RegisterType((*Resp)(nil), "clientrpc.Resp")
	proto.RegisterType((*StatusCode)(nil), "clientrpc.StatusCode")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Omq service

type OmqClient interface {
	PutMsg(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*StatusCode, error)
	Poll(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error)
}

type omqClient struct {
	cc *grpc.ClientConn
}

func NewOmqClient(cc *grpc.ClientConn) OmqClient {
	return &omqClient{cc}
}

func (c *omqClient) PutMsg(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*StatusCode, error) {
	out := new(StatusCode)
	err := grpc.Invoke(ctx, "/clientrpc.Omq/PutMsg", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omqClient) Poll(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := grpc.Invoke(ctx, "/clientrpc.Omq/poll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Omq service

type OmqServer interface {
	PutMsg(context.Context, *Msg) (*StatusCode, error)
	Poll(context.Context, *Req) (*Resp, error)
}

func RegisterOmqServer(s *grpc.Server, srv OmqServer) {
	s.RegisterService(&_Omq_serviceDesc, srv)
}

func _Omq_PutMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmqServer).PutMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clientrpc.Omq/PutMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmqServer).PutMsg(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Omq_Poll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmqServer).Poll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clientrpc.Omq/Poll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmqServer).Poll(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _Omq_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clientrpc.Omq",
	HandlerType: (*OmqServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutMsg",
			Handler:    _Omq_PutMsg_Handler,
		},
		{
			MethodName: "poll",
			Handler:    _Omq_Poll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "clirpc.proto",
}

func init() { proto.RegisterFile("clirpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x8b, 0x83, 0x30,
	0x10, 0xc5, 0xd5, 0xb8, 0x82, 0xc3, 0xb2, 0xbb, 0x04, 0xb6, 0x88, 0x50, 0x28, 0x39, 0xb5, 0x17,
	0xa9, 0xed, 0xa5, 0xf7, 0x9e, 0xa5, 0x25, 0xfd, 0x04, 0x35, 0x46, 0x11, 0x62, 0x13, 0x4d, 0x3c,
	0xf4, 0xdb, 0x17, 0x83, 0xb4, 0x7a, 0xf0, 0xf6, 0xde, 0x0c, 0xbf, 0x37, 0x7f, 0xe0, 0x9b, 0x89,
	0xba, 0x53, 0x2c, 0x51, 0x9d, 0x34, 0x12, 0x87, 0x4c, 0xd4, 0xfc, 0x61, 0x3a, 0xc5, 0xc8, 0x1a,
	0x10, 0xe5, 0x2d, 0x5e, 0x41, 0x20, 0xcb, 0x52, 0x73, 0x13, 0xb9, 0x1b, 0x77, 0x1b, 0xd2, 0xd1,
	0x91, 0x14, 0x50, 0xa6, 0xab, 0xa5, 0x36, 0xc6, 0xe0, 0xe7, 0xb2, 0x78, 0x46, 0x9e, 0xad, 0x5a,
	0x4d, 0xf6, 0xe0, 0x53, 0xae, 0xd5, 0x22, 0xf3, 0x07, 0xa8, 0xd1, 0xd5, 0x88, 0x0c, 0x92, 0x9c,
	0x00, 0x6e, 0xe6, 0x6e, 0x7a, 0x7d, 0x96, 0x05, 0x1f, 0x32, 0x99, 0x2c, 0xb8, 0xa5, 0xbe, 0xa8,
	0xd5, 0x93, 0x2c, 0x6f, 0x9a, 0x75, 0x60, 0x80, 0x2e, 0x4d, 0x8b, 0x53, 0x08, 0xae, 0xbd, 0x19,
	0x16, 0xfd, 0x49, 0xde, 0xa7, 0x25, 0x99, 0xae, 0xe2, 0xff, 0x89, 0xff, 0xcc, 0x20, 0x0e, 0xde,
	0x81, 0xaf, 0xa4, 0x10, 0x33, 0x80, 0xf2, 0x36, 0xfe, 0x9d, 0x79, 0xad, 0x88, 0x93, 0x07, 0xf6,
	0x69, 0xc7, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x71, 0x73, 0x4d, 0x44, 0x01, 0x00, 0x00,
}

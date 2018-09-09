// Code generated by protoc-gen-go. DO NOT EDIT.
// source: commands.proto

package commands

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

type Data struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Value                int64    `protobuf:"varint,2,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_commands_05adc7e371293899, []int{0}
}
func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (dst *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(dst, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Data) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Request struct {
	Value                []int64  `protobuf:"varint,1,rep,packed,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_commands_05adc7e371293899, []int{1}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetValue() []int64 {
	if m != nil {
		return m.Value
	}
	return nil
}

type Reply struct {
	Msg                  []*Data  `protobuf:"bytes,1,rep,name=Msg,proto3" json:"Msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_commands_05adc7e371293899, []int{2}
}
func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (dst *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(dst, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetMsg() []*Data {
	if m != nil {
		return m.Msg
	}
	return nil
}

func init() {
	proto.RegisterType((*Data)(nil), "commands.Data")
	proto.RegisterType((*Request)(nil), "commands.Request")
	proto.RegisterType((*Reply)(nil), "commands.Reply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommandsClient is the client API for Commands service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommandsClient interface {
	Add(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	Sleep(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
}

type commandsClient struct {
	cc *grpc.ClientConn
}

func NewCommandsClient(cc *grpc.ClientConn) CommandsClient {
	return &commandsClient{cc}
}

func (c *commandsClient) Add(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/commands.Commands/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandsClient) Sleep(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/commands.Commands/Sleep", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommandsServer is the server API for Commands service.
type CommandsServer interface {
	Add(context.Context, *Request) (*Reply, error)
	Sleep(context.Context, *Request) (*Reply, error)
}

func RegisterCommandsServer(s *grpc.Server, srv CommandsServer) {
	s.RegisterService(&_Commands_serviceDesc, srv)
}

func _Commands_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandsServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commands.Commands/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandsServer).Add(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Commands_Sleep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandsServer).Sleep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/commands.Commands/Sleep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandsServer).Sleep(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Commands_serviceDesc = grpc.ServiceDesc{
	ServiceName: "commands.Commands",
	HandlerType: (*CommandsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Commands_Add_Handler,
		},
		{
			MethodName: "Sleep",
			Handler:    _Commands_Sleep_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "commands.proto",
}

func init() { proto.RegisterFile("commands.proto", fileDescriptor_commands_05adc7e371293899) }

var fileDescriptor_commands_05adc7e371293899 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xce, 0xcf, 0xcd,
	0x4d, 0xcc, 0x4b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x0c,
	0xb8, 0x58, 0x5c, 0x12, 0x4b, 0x12, 0x85, 0x84, 0xb8, 0x58, 0xfc, 0x12, 0x73, 0x53, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x11, 0x2e, 0xd6, 0xb0, 0xc4, 0x9c, 0xd2, 0x54,
	0x09, 0x26, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x08, 0x47, 0x49, 0x9e, 0x8b, 0x3d, 0x28, 0xb5, 0xb0,
	0x34, 0xb5, 0xb8, 0x04, 0xa1, 0x80, 0x51, 0x81, 0x19, 0xa1, 0x40, 0x93, 0x8b, 0x35, 0x28, 0xb5,
	0x20, 0xa7, 0x52, 0x48, 0x81, 0x8b, 0xd9, 0xb7, 0x38, 0x1d, 0x2c, 0xc9, 0x6d, 0xc4, 0xa7, 0x07,
	0x77, 0x03, 0xc8, 0xc2, 0x20, 0x90, 0x94, 0x51, 0x1a, 0x17, 0x87, 0x33, 0x54, 0x54, 0x48, 0x9b,
	0x8b, 0xd9, 0x31, 0x25, 0x45, 0x48, 0x10, 0xa1, 0x0e, 0x6a, 0x8d, 0x14, 0x3f, 0xb2, 0x50, 0x41,
	0x4e, 0xa5, 0x12, 0x83, 0x90, 0x2e, 0x17, 0x6b, 0x70, 0x4e, 0x6a, 0x6a, 0x01, 0x71, 0xca, 0x93,
	0xd8, 0xc0, 0xde, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x5f, 0xa3, 0xb8, 0x08, 0x01,
	0x00, 0x00,
}
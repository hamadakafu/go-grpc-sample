// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sample.proto

package sample

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RequestGreet struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestGreet) Reset()         { *m = RequestGreet{} }
func (m *RequestGreet) String() string { return proto.CompactTextString(m) }
func (*RequestGreet) ProtoMessage()    {}
func (*RequestGreet) Descriptor() ([]byte, []int) {
	return fileDescriptor_2141552de9bf11d0, []int{0}
}

func (m *RequestGreet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestGreet.Unmarshal(m, b)
}
func (m *RequestGreet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestGreet.Marshal(b, m, deterministic)
}
func (m *RequestGreet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestGreet.Merge(m, src)
}
func (m *RequestGreet) XXX_Size() int {
	return xxx_messageInfo_RequestGreet.Size(m)
}
func (m *RequestGreet) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestGreet.DiscardUnknown(m)
}

var xxx_messageInfo_RequestGreet proto.InternalMessageInfo

func (m *RequestGreet) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ResponseGreet struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseGreet) Reset()         { *m = ResponseGreet{} }
func (m *ResponseGreet) String() string { return proto.CompactTextString(m) }
func (*ResponseGreet) ProtoMessage()    {}
func (*ResponseGreet) Descriptor() ([]byte, []int) {
	return fileDescriptor_2141552de9bf11d0, []int{1}
}

func (m *ResponseGreet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseGreet.Unmarshal(m, b)
}
func (m *ResponseGreet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseGreet.Marshal(b, m, deterministic)
}
func (m *ResponseGreet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseGreet.Merge(m, src)
}
func (m *ResponseGreet) XXX_Size() int {
	return xxx_messageInfo_ResponseGreet.Size(m)
}
func (m *ResponseGreet) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseGreet.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseGreet proto.InternalMessageInfo

func (m *ResponseGreet) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestGreet)(nil), "sample.RequestGreet")
	proto.RegisterType((*ResponseGreet)(nil), "sample.ResponseGreet")
}

func init() { proto.RegisterFile("sample.proto", fileDescriptor_2141552de9bf11d0) }

var fileDescriptor_2141552de9bf11d0 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0xcc, 0x2d,
	0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x34, 0xb8, 0x78,
	0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0xdc, 0x8b, 0x52, 0x53, 0x4b, 0x84, 0x24, 0xb8, 0xd8,
	0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c,
	0x25, 0x4d, 0x2e, 0xde, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x02, 0x4a, 0x8d, 0x2e,
	0x33, 0x72, 0xb1, 0x42, 0xd4, 0x58, 0x72, 0x71, 0x85, 0xe6, 0x25, 0x16, 0x55, 0x42, 0x78, 0x22,
	0x7a, 0x50, 0x37, 0x20, 0x5b, 0x29, 0x25, 0x8a, 0x10, 0x45, 0x36, 0xde, 0x89, 0x4b, 0x30, 0x38,
	0xb5, 0xa8, 0x2c, 0xb5, 0x28, 0xb8, 0xa4, 0x28, 0x35, 0x31, 0x97, 0x74, 0x13, 0x0c, 0x18, 0x85,
	0xbc, 0xb9, 0x24, 0x9c, 0x32, 0x53, 0x32, 0x8b, 0x52, 0x93, 0x4b, 0x32, 0xf3, 0xf3, 0x12, 0x73,
	0xc8, 0x35, 0x4a, 0x83, 0xd1, 0x80, 0x31, 0x89, 0x0d, 0x1c, 0x72, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x2f, 0xeb, 0x5f, 0x65, 0x49, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreetClient is the client API for Greet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetClient interface {
	UnaryGreet(ctx context.Context, in *RequestGreet, opts ...grpc.CallOption) (*ResponseGreet, error)
	ServerStreamGreet(ctx context.Context, in *RequestGreet, opts ...grpc.CallOption) (Greet_ServerStreamGreetClient, error)
	BidirectionalStreamGreet(ctx context.Context, opts ...grpc.CallOption) (Greet_BidirectionalStreamGreetClient, error)
}

type greetClient struct {
	cc *grpc.ClientConn
}

func NewGreetClient(cc *grpc.ClientConn) GreetClient {
	return &greetClient{cc}
}

func (c *greetClient) UnaryGreet(ctx context.Context, in *RequestGreet, opts ...grpc.CallOption) (*ResponseGreet, error) {
	out := new(ResponseGreet)
	err := c.cc.Invoke(ctx, "/sample.Greet/UnaryGreet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetClient) ServerStreamGreet(ctx context.Context, in *RequestGreet, opts ...grpc.CallOption) (Greet_ServerStreamGreetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greet_serviceDesc.Streams[0], "/sample.Greet/ServerStreamGreet", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServerStreamGreetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greet_ServerStreamGreetClient interface {
	Recv() (*ResponseGreet, error)
	grpc.ClientStream
}

type greetServerStreamGreetClient struct {
	grpc.ClientStream
}

func (x *greetServerStreamGreetClient) Recv() (*ResponseGreet, error) {
	m := new(ResponseGreet)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetClient) BidirectionalStreamGreet(ctx context.Context, opts ...grpc.CallOption) (Greet_BidirectionalStreamGreetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Greet_serviceDesc.Streams[1], "/sample.Greet/BidirectionalStreamGreet", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetBidirectionalStreamGreetClient{stream}
	return x, nil
}

type Greet_BidirectionalStreamGreetClient interface {
	Send(*RequestGreet) error
	Recv() (*ResponseGreet, error)
	grpc.ClientStream
}

type greetBidirectionalStreamGreetClient struct {
	grpc.ClientStream
}

func (x *greetBidirectionalStreamGreetClient) Send(m *RequestGreet) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetBidirectionalStreamGreetClient) Recv() (*ResponseGreet, error) {
	m := new(ResponseGreet)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServer is the server API for Greet service.
type GreetServer interface {
	UnaryGreet(context.Context, *RequestGreet) (*ResponseGreet, error)
	ServerStreamGreet(*RequestGreet, Greet_ServerStreamGreetServer) error
	BidirectionalStreamGreet(Greet_BidirectionalStreamGreetServer) error
}

// UnimplementedGreetServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServer struct {
}

func (*UnimplementedGreetServer) UnaryGreet(ctx context.Context, req *RequestGreet) (*ResponseGreet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryGreet not implemented")
}
func (*UnimplementedGreetServer) ServerStreamGreet(req *RequestGreet, srv Greet_ServerStreamGreetServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreamGreet not implemented")
}
func (*UnimplementedGreetServer) BidirectionalStreamGreet(srv Greet_BidirectionalStreamGreetServer) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalStreamGreet not implemented")
}

func RegisterGreetServer(s *grpc.Server, srv GreetServer) {
	s.RegisterService(&_Greet_serviceDesc, srv)
}

func _Greet_UnaryGreet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGreet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServer).UnaryGreet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sample.Greet/UnaryGreet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServer).UnaryGreet(ctx, req.(*RequestGreet))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greet_ServerStreamGreet_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestGreet)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServer).ServerStreamGreet(m, &greetServerStreamGreetServer{stream})
}

type Greet_ServerStreamGreetServer interface {
	Send(*ResponseGreet) error
	grpc.ServerStream
}

type greetServerStreamGreetServer struct {
	grpc.ServerStream
}

func (x *greetServerStreamGreetServer) Send(m *ResponseGreet) error {
	return x.ServerStream.SendMsg(m)
}

func _Greet_BidirectionalStreamGreet_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServer).BidirectionalStreamGreet(&greetBidirectionalStreamGreetServer{stream})
}

type Greet_BidirectionalStreamGreetServer interface {
	Send(*ResponseGreet) error
	Recv() (*RequestGreet, error)
	grpc.ServerStream
}

type greetBidirectionalStreamGreetServer struct {
	grpc.ServerStream
}

func (x *greetBidirectionalStreamGreetServer) Send(m *ResponseGreet) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetBidirectionalStreamGreetServer) Recv() (*RequestGreet, error) {
	m := new(RequestGreet)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Greet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sample.Greet",
	HandlerType: (*GreetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryGreet",
			Handler:    _Greet_UnaryGreet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStreamGreet",
			Handler:       _Greet_ServerStreamGreet_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BidirectionalStreamGreet",
			Handler:       _Greet_BidirectionalStreamGreet_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "sample.proto",
}

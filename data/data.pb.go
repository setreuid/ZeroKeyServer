// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data.proto

package data

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

type KeysRequest struct {
	Acts                 []*Act   `protobuf:"bytes,1,rep,name=acts,proto3" json:"acts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeysRequest) Reset()         { *m = KeysRequest{} }
func (m *KeysRequest) String() string { return proto.CompactTextString(m) }
func (*KeysRequest) ProtoMessage()    {}
func (*KeysRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{0}
}

func (m *KeysRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeysRequest.Unmarshal(m, b)
}
func (m *KeysRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeysRequest.Marshal(b, m, deterministic)
}
func (m *KeysRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeysRequest.Merge(m, src)
}
func (m *KeysRequest) XXX_Size() int {
	return xxx_messageInfo_KeysRequest.Size(m)
}
func (m *KeysRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KeysRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KeysRequest proto.InternalMessageInfo

func (m *KeysRequest) GetActs() []*Act {
	if m != nil {
		return m.Acts
	}
	return nil
}

type Act struct {
	Code                 float64  `protobuf:"fixed64,1,opt,name=code,proto3" json:"code,omitempty"`
	State                float64  `protobuf:"fixed64,2,opt,name=state,proto3" json:"state,omitempty"`
	Time                 float64  `protobuf:"fixed64,3,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Act) Reset()         { *m = Act{} }
func (m *Act) String() string { return proto.CompactTextString(m) }
func (*Act) ProtoMessage()    {}
func (*Act) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{1}
}

func (m *Act) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Act.Unmarshal(m, b)
}
func (m *Act) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Act.Marshal(b, m, deterministic)
}
func (m *Act) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Act.Merge(m, src)
}
func (m *Act) XXX_Size() int {
	return xxx_messageInfo_Act.Size(m)
}
func (m *Act) XXX_DiscardUnknown() {
	xxx_messageInfo_Act.DiscardUnknown(m)
}

var xxx_messageInfo_Act proto.InternalMessageInfo

func (m *Act) GetCode() float64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Act) GetState() float64 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *Act) GetTime() float64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type Noting struct {
	Status               float64  `protobuf:"fixed64,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Noting) Reset()         { *m = Noting{} }
func (m *Noting) String() string { return proto.CompactTextString(m) }
func (*Noting) ProtoMessage()    {}
func (*Noting) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{2}
}

func (m *Noting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Noting.Unmarshal(m, b)
}
func (m *Noting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Noting.Marshal(b, m, deterministic)
}
func (m *Noting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Noting.Merge(m, src)
}
func (m *Noting) XXX_Size() int {
	return xxx_messageInfo_Noting.Size(m)
}
func (m *Noting) XXX_DiscardUnknown() {
	xxx_messageInfo_Noting.DiscardUnknown(m)
}

var xxx_messageInfo_Noting proto.InternalMessageInfo

func (m *Noting) GetStatus() float64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*KeysRequest)(nil), "data.KeysRequest")
	proto.RegisterType((*Act)(nil), "data.Act")
	proto.RegisterType((*Noting)(nil), "data.Noting")
}

func init() { proto.RegisterFile("data.proto", fileDescriptor_871986018790d2fd) }

var fileDescriptor_871986018790d2fd = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x14, 0x45, 0x27, 0x26, 0x56, 0x7c, 0xe3, 0xc6, 0x87, 0x48, 0x18, 0x10, 0x4a, 0x74, 0x31, 0x82,
	0xcc, 0x62, 0xfc, 0x82, 0xa1, 0x4b, 0xc5, 0x45, 0xfa, 0x05, 0x31, 0x09, 0xd2, 0x45, 0x9b, 0x6a,
	0x5e, 0x05, 0x7f, 0xc9, 0xaf, 0x94, 0x24, 0x5d, 0x68, 0x85, 0xd9, 0xdd, 0x77, 0x38, 0xb9, 0x5c,
	0x02, 0xe0, 0x0c, 0x99, 0xdd, 0xf8, 0x11, 0x28, 0xa0, 0x48, 0x59, 0x3d, 0xc0, 0xfa, 0xc9, 0x7f,
	0x45, 0xed, 0xdf, 0x27, 0x1f, 0x09, 0x6f, 0x40, 0x18, 0x4b, 0x51, 0xb2, 0x9a, 0x6f, 0xd7, 0xfb,
	0xf3, 0x5d, 0xf6, 0x0f, 0x96, 0x74, 0xc6, 0xaa, 0x01, 0x7e, 0xb0, 0x84, 0x08, 0xc2, 0x06, 0xe7,
	0x25, 0xab, 0xd9, 0x96, 0xe9, 0x9c, 0xf1, 0x0a, 0x4e, 0x23, 0x19, 0xf2, 0xf2, 0x24, 0xc3, 0x72,
	0x24, 0x93, 0xba, 0xde, 0x4b, 0x5e, 0xcc, 0x94, 0x55, 0x0d, 0xd5, 0x4b, 0xa0, 0x6e, 0x78, 0xc3,
	0x6b, 0xa8, 0x92, 0x36, 0xc5, 0xb9, 0x69, 0xbe, 0xf6, 0xdf, 0x0c, 0xce, 0x9a, 0xd0, 0xf7, 0x66,
	0x70, 0x78, 0x0f, 0xa2, 0x35, 0x9f, 0x1e, 0x2f, 0xcb, 0x96, 0x5f, 0x63, 0x37, 0x17, 0x05, 0x95,
	0x32, 0xb5, 0x4a, 0xea, 0x73, 0x30, 0x0e, 0xff, 0xf0, 0xcd, 0xff, 0x87, 0x6a, 0x85, 0xb7, 0xc0,
	0xf5, 0x34, 0x2c, 0xcc, 0x65, 0xdf, 0x1d, 0x88, 0x96, 0xc2, 0x78, 0xdc, 0x7a, 0xad, 0xf2, 0x77,
	0x3e, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x2e, 0x0f, 0x1c, 0x5c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommandClient is the client API for Command service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommandClient interface {
	Save(ctx context.Context, in *KeysRequest, opts ...grpc.CallOption) (*Noting, error)
	Load(ctx context.Context, in *Noting, opts ...grpc.CallOption) (*KeysRequest, error)
	Run(ctx context.Context, in *Noting, opts ...grpc.CallOption) (*Noting, error)
	Stop(ctx context.Context, in *Noting, opts ...grpc.CallOption) (*Noting, error)
}

type commandClient struct {
	cc *grpc.ClientConn
}

func NewCommandClient(cc *grpc.ClientConn) CommandClient {
	return &commandClient{cc}
}

func (c *commandClient) Save(ctx context.Context, in *KeysRequest, opts ...grpc.CallOption) (*Noting, error) {
	out := new(Noting)
	err := c.cc.Invoke(ctx, "/data.Command/Save", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandClient) Load(ctx context.Context, in *Noting, opts ...grpc.CallOption) (*KeysRequest, error) {
	out := new(KeysRequest)
	err := c.cc.Invoke(ctx, "/data.Command/Load", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandClient) Run(ctx context.Context, in *Noting, opts ...grpc.CallOption) (*Noting, error) {
	out := new(Noting)
	err := c.cc.Invoke(ctx, "/data.Command/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandClient) Stop(ctx context.Context, in *Noting, opts ...grpc.CallOption) (*Noting, error) {
	out := new(Noting)
	err := c.cc.Invoke(ctx, "/data.Command/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommandServer is the server API for Command service.
type CommandServer interface {
	Save(context.Context, *KeysRequest) (*Noting, error)
	Load(context.Context, *Noting) (*KeysRequest, error)
	Run(context.Context, *Noting) (*Noting, error)
	Stop(context.Context, *Noting) (*Noting, error)
}

// UnimplementedCommandServer can be embedded to have forward compatible implementations.
type UnimplementedCommandServer struct {
}

func (*UnimplementedCommandServer) Save(ctx context.Context, req *KeysRequest) (*Noting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (*UnimplementedCommandServer) Load(ctx context.Context, req *Noting) (*KeysRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Load not implemented")
}
func (*UnimplementedCommandServer) Run(ctx context.Context, req *Noting) (*Noting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (*UnimplementedCommandServer) Stop(ctx context.Context, req *Noting) (*Noting, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}

func RegisterCommandServer(s *grpc.Server, srv CommandServer) {
	s.RegisterService(&_Command_serviceDesc, srv)
}

func _Command_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.Command/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServer).Save(ctx, req.(*KeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Command_Load_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Noting)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServer).Load(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.Command/Load",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServer).Load(ctx, req.(*Noting))
	}
	return interceptor(ctx, in, info, handler)
}

func _Command_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Noting)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.Command/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServer).Run(ctx, req.(*Noting))
	}
	return interceptor(ctx, in, info, handler)
}

func _Command_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Noting)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.Command/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServer).Stop(ctx, req.(*Noting))
	}
	return interceptor(ctx, in, info, handler)
}

var _Command_serviceDesc = grpc.ServiceDesc{
	ServiceName: "data.Command",
	HandlerType: (*CommandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Save",
			Handler:    _Command_Save_Handler,
		},
		{
			MethodName: "Load",
			Handler:    _Command_Load_Handler,
		},
		{
			MethodName: "Run",
			Handler:    _Command_Run_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Command_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data.proto",
}

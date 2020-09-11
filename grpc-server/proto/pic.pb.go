// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pic.proto

package pic_proto

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

type RequestPic struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestPic) Reset()         { *m = RequestPic{} }
func (m *RequestPic) String() string { return proto.CompactTextString(m) }
func (*RequestPic) ProtoMessage()    {}
func (*RequestPic) Descriptor() ([]byte, []int) {
	return fileDescriptor_27532d036b1f396d, []int{0}
}

func (m *RequestPic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestPic.Unmarshal(m, b)
}
func (m *RequestPic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestPic.Marshal(b, m, deterministic)
}
func (m *RequestPic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestPic.Merge(m, src)
}
func (m *RequestPic) XXX_Size() int {
	return xxx_messageInfo_RequestPic.Size(m)
}
func (m *RequestPic) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestPic.DiscardUnknown(m)
}

var xxx_messageInfo_RequestPic proto.InternalMessageInfo

func (m *RequestPic) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *RequestPic) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type ResposePic struct {
	Pics                 []string `protobuf:"bytes,1,rep,name=pics,proto3" json:"pics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResposePic) Reset()         { *m = ResposePic{} }
func (m *ResposePic) String() string { return proto.CompactTextString(m) }
func (*ResposePic) ProtoMessage()    {}
func (*ResposePic) Descriptor() ([]byte, []int) {
	return fileDescriptor_27532d036b1f396d, []int{1}
}

func (m *ResposePic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResposePic.Unmarshal(m, b)
}
func (m *ResposePic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResposePic.Marshal(b, m, deterministic)
}
func (m *ResposePic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResposePic.Merge(m, src)
}
func (m *ResposePic) XXX_Size() int {
	return xxx_messageInfo_ResposePic.Size(m)
}
func (m *ResposePic) XXX_DiscardUnknown() {
	xxx_messageInfo_ResposePic.DiscardUnknown(m)
}

var xxx_messageInfo_ResposePic proto.InternalMessageInfo

func (m *ResposePic) GetPics() []string {
	if m != nil {
		return m.Pics
	}
	return nil
}

func init() {
	proto.RegisterType((*RequestPic)(nil), "pic_proto.RequestPic")
	proto.RegisterType((*ResposePic)(nil), "pic_proto.ResposePic")
}

func init() { proto.RegisterFile("pic.proto", fileDescriptor_27532d036b1f396d) }

var fileDescriptor_27532d036b1f396d = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0xc8, 0x4c, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x02, 0x31, 0xe3, 0xc1, 0x4c, 0x25, 0x5b, 0x2e, 0xae, 0xa0,
	0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x92, 0x80, 0xcc, 0x64, 0x21, 0x21, 0x2e, 0x96, 0x82, 0xc4, 0xf4,
	0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x30, 0x5b, 0x48, 0x9a, 0x8b, 0x13, 0x44, 0xc7,
	0x17, 0x67, 0x56, 0xa5, 0x4a, 0x30, 0x81, 0x25, 0x38, 0x40, 0x02, 0xc1, 0x99, 0x55, 0xa9, 0x4a,
	0x0a, 0x20, 0xed, 0xc5, 0x05, 0xf9, 0xc5, 0xa9, 0x30, 0xed, 0x99, 0xc9, 0xc5, 0x12, 0x8c, 0x0a,
	0xcc, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x91, 0x1b, 0x17, 0x67, 0x40, 0x66, 0x72, 0x70, 0x6a, 0x51,
	0x59, 0x6a, 0x91, 0x90, 0x25, 0x17, 0xbb, 0x7b, 0x2a, 0xc8, 0xa6, 0x62, 0x21, 0x51, 0x3d, 0xb8,
	0x23, 0xf4, 0x10, 0x2e, 0x90, 0x42, 0x15, 0x86, 0x99, 0xac, 0xc4, 0x90, 0xc4, 0x06, 0x16, 0x33,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x06, 0x31, 0xc4, 0x4f, 0xc7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PicServerClient is the client API for PicServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PicServerClient interface {
	GetPics(ctx context.Context, in *RequestPic, opts ...grpc.CallOption) (*ResposePic, error)
}

type picServerClient struct {
	cc *grpc.ClientConn
}

func NewPicServerClient(cc *grpc.ClientConn) PicServerClient {
	return &picServerClient{cc}
}

func (c *picServerClient) GetPics(ctx context.Context, in *RequestPic, opts ...grpc.CallOption) (*ResposePic, error) {
	out := new(ResposePic)
	err := c.cc.Invoke(ctx, "/pic_proto.PicServer/GetPics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PicServerServer is the server API for PicServer service.
type PicServerServer interface {
	GetPics(context.Context, *RequestPic) (*ResposePic, error)
}

// UnimplementedPicServerServer can be embedded to have forward compatible implementations.
type UnimplementedPicServerServer struct {
}

func (*UnimplementedPicServerServer) GetPics(ctx context.Context, req *RequestPic) (*ResposePic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPics not implemented")
}

func RegisterPicServerServer(s *grpc.Server, srv PicServerServer) {
	s.RegisterService(&_PicServer_serviceDesc, srv)
}

func _PicServer_GetPics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestPic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PicServerServer).GetPics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pic_proto.PicServer/GetPics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PicServerServer).GetPics(ctx, req.(*RequestPic))
	}
	return interceptor(ctx, in, info, handler)
}

var _PicServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pic_proto.PicServer",
	HandlerType: (*PicServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPics",
			Handler:    _PicServer_GetPics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pic.proto",
}

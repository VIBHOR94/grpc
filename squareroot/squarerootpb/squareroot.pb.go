// Code generated by protoc-gen-go. DO NOT EDIT.
// source: squareroot/squarerootpb/squareroot.proto

package square_rootpb

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

type SquareRootRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquareRootRequest) Reset()         { *m = SquareRootRequest{} }
func (m *SquareRootRequest) String() string { return proto.CompactTextString(m) }
func (*SquareRootRequest) ProtoMessage()    {}
func (*SquareRootRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e9636629a13a5e0, []int{0}
}

func (m *SquareRootRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquareRootRequest.Unmarshal(m, b)
}
func (m *SquareRootRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquareRootRequest.Marshal(b, m, deterministic)
}
func (m *SquareRootRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquareRootRequest.Merge(m, src)
}
func (m *SquareRootRequest) XXX_Size() int {
	return xxx_messageInfo_SquareRootRequest.Size(m)
}
func (m *SquareRootRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SquareRootRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SquareRootRequest proto.InternalMessageInfo

func (m *SquareRootRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type SquarerootResponse struct {
	Result               float64  `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SquarerootResponse) Reset()         { *m = SquarerootResponse{} }
func (m *SquarerootResponse) String() string { return proto.CompactTextString(m) }
func (*SquarerootResponse) ProtoMessage()    {}
func (*SquarerootResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e9636629a13a5e0, []int{1}
}

func (m *SquarerootResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SquarerootResponse.Unmarshal(m, b)
}
func (m *SquarerootResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SquarerootResponse.Marshal(b, m, deterministic)
}
func (m *SquarerootResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SquarerootResponse.Merge(m, src)
}
func (m *SquarerootResponse) XXX_Size() int {
	return xxx_messageInfo_SquarerootResponse.Size(m)
}
func (m *SquarerootResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SquarerootResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SquarerootResponse proto.InternalMessageInfo

func (m *SquarerootResponse) GetResult() float64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*SquareRootRequest)(nil), "squareroot.SquareRootRequest")
	proto.RegisterType((*SquarerootResponse)(nil), "squareroot.SquarerootResponse")
}

func init() {
	proto.RegisterFile("squareroot/squarerootpb/squareroot.proto", fileDescriptor_4e9636629a13a5e0)
}

var fileDescriptor_4e9636629a13a5e0 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x28, 0x2e, 0x2c, 0x4d,
	0x2c, 0x4a, 0x2d, 0xca, 0xcf, 0x2f, 0xd1, 0x47, 0x30, 0x0b, 0x92, 0x90, 0x38, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0x5c, 0x08, 0x11, 0x25, 0x6d, 0x2e, 0xc1, 0x60, 0x30, 0x2f, 0x28, 0x3f,
	0xbf, 0x24, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x8c, 0x8b, 0x2d, 0xaf, 0x34, 0x37,
	0x29, 0xb5, 0x48, 0x82, 0x51, 0x81, 0x51, 0x83, 0x35, 0x08, 0xca, 0x53, 0xd2, 0xe1, 0x12, 0x0a,
	0x86, 0x6b, 0x0d, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x05, 0xa9, 0x2e, 0x4a, 0x2d, 0x2e,
	0xcd, 0x29, 0x01, 0xab, 0x66, 0x0c, 0x82, 0xf2, 0x8c, 0x32, 0x91, 0x8d, 0x0e, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x15, 0x0a, 0xe1, 0x12, 0x74, 0xce, 0xcf, 0x2d, 0x28, 0x2d, 0x49, 0x45, 0xc8,
	0x09, 0xc9, 0xea, 0x21, 0xb9, 0x11, 0xc3, 0x39, 0x52, 0x72, 0x98, 0xd2, 0xc8, 0x0e, 0x50, 0x62,
	0x70, 0xe2, 0x8f, 0xe2, 0x85, 0x28, 0x89, 0x87, 0xf8, 0x39, 0x89, 0x0d, 0xec, 0x53, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xba, 0x08, 0x97, 0x80, 0x15, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SquareRootServiceClient is the client API for SquareRootService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SquareRootServiceClient interface {
	// Will Send error of type INVALID_ARGUMENT if number sent is negative
	ComputeSquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquarerootResponse, error)
}

type squareRootServiceClient struct {
	cc *grpc.ClientConn
}

func NewSquareRootServiceClient(cc *grpc.ClientConn) SquareRootServiceClient {
	return &squareRootServiceClient{cc}
}

func (c *squareRootServiceClient) ComputeSquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquarerootResponse, error) {
	out := new(SquarerootResponse)
	err := c.cc.Invoke(ctx, "/squareroot.SquareRootService/ComputeSquareRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SquareRootServiceServer is the server API for SquareRootService service.
type SquareRootServiceServer interface {
	// Will Send error of type INVALID_ARGUMENT if number sent is negative
	ComputeSquareRoot(context.Context, *SquareRootRequest) (*SquarerootResponse, error)
}

// UnimplementedSquareRootServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSquareRootServiceServer struct {
}

func (*UnimplementedSquareRootServiceServer) ComputeSquareRoot(ctx context.Context, req *SquareRootRequest) (*SquarerootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComputeSquareRoot not implemented")
}

func RegisterSquareRootServiceServer(s *grpc.Server, srv SquareRootServiceServer) {
	s.RegisterService(&_SquareRootService_serviceDesc, srv)
}

func _SquareRootService_ComputeSquareRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquareRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SquareRootServiceServer).ComputeSquareRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/squareroot.SquareRootService/ComputeSquareRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SquareRootServiceServer).ComputeSquareRoot(ctx, req.(*SquareRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SquareRootService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "squareroot.SquareRootService",
	HandlerType: (*SquareRootServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ComputeSquareRoot",
			Handler:    _SquareRootService_ComputeSquareRoot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "squareroot/squarerootpb/squareroot.proto",
}

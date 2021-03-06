// Code generated by protoc-gen-go. DO NOT EDIT.
// source: number_decomposition/number_decompositionpb/number_decomposition.proto

package number_decompositionpb

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

type NumberDecompositionRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NumberDecompositionRequest) Reset()         { *m = NumberDecompositionRequest{} }
func (m *NumberDecompositionRequest) String() string { return proto.CompactTextString(m) }
func (*NumberDecompositionRequest) ProtoMessage()    {}
func (*NumberDecompositionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_264a3250a52e4d69, []int{0}
}

func (m *NumberDecompositionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NumberDecompositionRequest.Unmarshal(m, b)
}
func (m *NumberDecompositionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NumberDecompositionRequest.Marshal(b, m, deterministic)
}
func (m *NumberDecompositionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumberDecompositionRequest.Merge(m, src)
}
func (m *NumberDecompositionRequest) XXX_Size() int {
	return xxx_messageInfo_NumberDecompositionRequest.Size(m)
}
func (m *NumberDecompositionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NumberDecompositionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NumberDecompositionRequest proto.InternalMessageInfo

func (m *NumberDecompositionRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type NumberDecompositionResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NumberDecompositionResponse) Reset()         { *m = NumberDecompositionResponse{} }
func (m *NumberDecompositionResponse) String() string { return proto.CompactTextString(m) }
func (*NumberDecompositionResponse) ProtoMessage()    {}
func (*NumberDecompositionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_264a3250a52e4d69, []int{1}
}

func (m *NumberDecompositionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NumberDecompositionResponse.Unmarshal(m, b)
}
func (m *NumberDecompositionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NumberDecompositionResponse.Marshal(b, m, deterministic)
}
func (m *NumberDecompositionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumberDecompositionResponse.Merge(m, src)
}
func (m *NumberDecompositionResponse) XXX_Size() int {
	return xxx_messageInfo_NumberDecompositionResponse.Size(m)
}
func (m *NumberDecompositionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NumberDecompositionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NumberDecompositionResponse proto.InternalMessageInfo

func (m *NumberDecompositionResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*NumberDecompositionRequest)(nil), "numberdecomposition.NumberDecompositionRequest")
	proto.RegisterType((*NumberDecompositionResponse)(nil), "numberdecomposition.NumberDecompositionResponse")
}

func init() {
	proto.RegisterFile("number_decomposition/number_decompositionpb/number_decomposition.proto", fileDescriptor_264a3250a52e4d69)
}

var fileDescriptor_264a3250a52e4d69 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0xcb, 0x2b, 0xcd, 0x4d,
	0x4a, 0x2d, 0x8a, 0x4f, 0x49, 0x4d, 0xce, 0xcf, 0x2d, 0xc8, 0x2f, 0xce, 0x2c, 0xc9, 0xcc, 0xcf,
	0xd3, 0xc7, 0x26, 0x58, 0x90, 0x84, 0x55, 0x58, 0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x48, 0x18,
	0x22, 0x87, 0x22, 0xa5, 0x64, 0xc2, 0x25, 0xe5, 0x07, 0x16, 0x76, 0x41, 0x16, 0x0e, 0x4a, 0x2d,
	0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe3, 0x62, 0x83, 0x68, 0x92, 0x60, 0x54, 0x60, 0xd4, 0x60,
	0x0d, 0x82, 0xf2, 0x94, 0x4c, 0xb9, 0xa4, 0xb1, 0xea, 0x2a, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x05,
	0x69, 0x2b, 0x4a, 0x2d, 0x2e, 0xcd, 0x29, 0x01, 0x6b, 0xe3, 0x0c, 0x82, 0xf2, 0x8c, 0xa6, 0x30,
	0x62, 0xb5, 0x2d, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0xa8, 0x8c, 0x8b, 0x17, 0x45, 0x5c,
	0x48, 0x5f, 0x0f, 0x8b, 0x93, 0xf5, 0x70, 0xbb, 0x57, 0xca, 0x80, 0x78, 0x0d, 0x10, 0xa7, 0x2a,
	0x31, 0x18, 0x30, 0x3a, 0x49, 0x44, 0x89, 0x61, 0x0f, 0xcd, 0x24, 0x36, 0x70, 0xc8, 0x19, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xf7, 0xc5, 0x4e, 0x83, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NumberDecompositionServiceClient is the client API for NumberDecompositionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NumberDecompositionServiceClient interface {
	// Server Streaming
	Decomposition(ctx context.Context, in *NumberDecompositionRequest, opts ...grpc.CallOption) (NumberDecompositionService_DecompositionClient, error)
}

type numberDecompositionServiceClient struct {
	cc *grpc.ClientConn
}

func NewNumberDecompositionServiceClient(cc *grpc.ClientConn) NumberDecompositionServiceClient {
	return &numberDecompositionServiceClient{cc}
}

func (c *numberDecompositionServiceClient) Decomposition(ctx context.Context, in *NumberDecompositionRequest, opts ...grpc.CallOption) (NumberDecompositionService_DecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NumberDecompositionService_serviceDesc.Streams[0], "/numberdecomposition.NumberDecompositionService/Decomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &numberDecompositionServiceDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NumberDecompositionService_DecompositionClient interface {
	Recv() (*NumberDecompositionResponse, error)
	grpc.ClientStream
}

type numberDecompositionServiceDecompositionClient struct {
	grpc.ClientStream
}

func (x *numberDecompositionServiceDecompositionClient) Recv() (*NumberDecompositionResponse, error) {
	m := new(NumberDecompositionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NumberDecompositionServiceServer is the server API for NumberDecompositionService service.
type NumberDecompositionServiceServer interface {
	// Server Streaming
	Decomposition(*NumberDecompositionRequest, NumberDecompositionService_DecompositionServer) error
}

// UnimplementedNumberDecompositionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNumberDecompositionServiceServer struct {
}

func (*UnimplementedNumberDecompositionServiceServer) Decomposition(req *NumberDecompositionRequest, srv NumberDecompositionService_DecompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method Decomposition not implemented")
}

func RegisterNumberDecompositionServiceServer(s *grpc.Server, srv NumberDecompositionServiceServer) {
	s.RegisterService(&_NumberDecompositionService_serviceDesc, srv)
}

func _NumberDecompositionService_Decomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NumberDecompositionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NumberDecompositionServiceServer).Decomposition(m, &numberDecompositionServiceDecompositionServer{stream})
}

type NumberDecompositionService_DecompositionServer interface {
	Send(*NumberDecompositionResponse) error
	grpc.ServerStream
}

type numberDecompositionServiceDecompositionServer struct {
	grpc.ServerStream
}

func (x *numberDecompositionServiceDecompositionServer) Send(m *NumberDecompositionResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _NumberDecompositionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "numberdecomposition.NumberDecompositionService",
	HandlerType: (*NumberDecompositionServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Decomposition",
			Handler:       _NumberDecompositionService_Decomposition_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "number_decomposition/number_decompositionpb/number_decomposition.proto",
}

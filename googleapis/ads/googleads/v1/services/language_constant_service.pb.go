// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/language_constant_service.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Request message for
// [LanguageConstantService.GetLanguageConstant][google.ads.googleads.v1.services.LanguageConstantService.GetLanguageConstant].
type GetLanguageConstantRequest struct {
	// Resource name of the language constant to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLanguageConstantRequest) Reset()         { *m = GetLanguageConstantRequest{} }
func (m *GetLanguageConstantRequest) String() string { return proto.CompactTextString(m) }
func (*GetLanguageConstantRequest) ProtoMessage()    {}
func (*GetLanguageConstantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_167ba166e955dde1, []int{0}
}

func (m *GetLanguageConstantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLanguageConstantRequest.Unmarshal(m, b)
}
func (m *GetLanguageConstantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLanguageConstantRequest.Marshal(b, m, deterministic)
}
func (m *GetLanguageConstantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLanguageConstantRequest.Merge(m, src)
}
func (m *GetLanguageConstantRequest) XXX_Size() int {
	return xxx_messageInfo_GetLanguageConstantRequest.Size(m)
}
func (m *GetLanguageConstantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLanguageConstantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLanguageConstantRequest proto.InternalMessageInfo

func (m *GetLanguageConstantRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetLanguageConstantRequest)(nil), "google.ads.googleads.v1.services.GetLanguageConstantRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/language_constant_service.proto", fileDescriptor_167ba166e955dde1)
}

var fileDescriptor_167ba166e955dde1 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x4a, 0xc3, 0x40,
	0x14, 0x26, 0x11, 0x04, 0x07, 0xdd, 0xc4, 0x85, 0x12, 0xba, 0x28, 0xb5, 0xe0, 0xcf, 0x62, 0x86,
	0xd8, 0x95, 0xa3, 0x82, 0x53, 0x17, 0x75, 0x21, 0x52, 0x2a, 0x74, 0x21, 0x81, 0x32, 0x36, 0xc3,
	0x10, 0x68, 0x67, 0x6a, 0xde, 0xb4, 0x1b, 0x71, 0xa3, 0x47, 0xf0, 0x06, 0x2e, 0xbd, 0x81, 0x57,
	0x70, 0x2b, 0xde, 0xc0, 0x95, 0xa7, 0x90, 0x74, 0x32, 0x01, 0x4b, 0x43, 0x77, 0x1f, 0xf3, 0xbe,
	0x9f, 0xf7, 0xbe, 0x04, 0x5d, 0x48, 0xad, 0xe5, 0x48, 0x10, 0x9e, 0x00, 0xb1, 0x30, 0x47, 0xb3,
	0x88, 0x80, 0xc8, 0x66, 0xe9, 0x50, 0x00, 0x19, 0x71, 0x25, 0xa7, 0x5c, 0x8a, 0xc1, 0x50, 0x2b,
	0x30, 0x5c, 0x99, 0x41, 0x31, 0xc2, 0x93, 0x4c, 0x1b, 0x1d, 0xd4, 0xad, 0x0c, 0xf3, 0x04, 0x70,
	0xe9, 0x80, 0x67, 0x11, 0x76, 0x0e, 0xe1, 0x49, 0x55, 0x46, 0x26, 0x40, 0x4f, 0xb3, 0xa5, 0x21,
	0xd6, 0x3c, 0xac, 0x39, 0xe9, 0x24, 0x25, 0x5c, 0x29, 0x6d, 0xb8, 0x49, 0xb5, 0x02, 0x3b, 0x6d,
	0x30, 0x14, 0x76, 0x84, 0xb9, 0x2e, 0xb4, 0x97, 0x85, 0xb4, 0x27, 0x1e, 0xa6, 0x02, 0x4c, 0xb0,
	0x87, 0xb6, 0x5c, 0xc0, 0x40, 0xf1, 0xb1, 0xd8, 0xf5, 0xea, 0xde, 0xc1, 0x46, 0x6f, 0xd3, 0x3d,
	0xde, 0xf0, 0xb1, 0x38, 0xfe, 0xf6, 0xd0, 0xce, 0xa2, 0xc1, 0xad, 0x5d, 0x3c, 0xf8, 0xf0, 0xd0,
	0xf6, 0x12, 0xff, 0xe0, 0x0c, 0xaf, 0x3a, 0x19, 0x57, 0xaf, 0x15, 0xb6, 0x2a, 0xd5, 0x65, 0x1d,
	0x78, 0x51, 0xdb, 0x20, 0xcf, 0x5f, 0x3f, 0xaf, 0xfe, 0x61, 0xb0, 0x9f, 0xd7, 0xf6, 0xf8, 0xef,
	0xac, 0xf3, 0xd1, 0x02, 0x19, 0xc8, 0xd1, 0x53, 0xfb, 0xc5, 0x47, 0xcd, 0xa1, 0x1e, 0xaf, 0xdc,
	0xb4, 0x5d, 0xab, 0xb8, 0xbe, 0x9b, 0x37, 0xdc, 0xf5, 0xee, 0xae, 0x0a, 0x07, 0xa9, 0xf3, 0x1c,
	0xac, 0x33, 0x49, 0xa4, 0x50, 0xf3, 0xfe, 0xdd, 0xc7, 0x9c, 0xa4, 0x50, 0xfd, 0xff, 0x9c, 0x3a,
	0xf0, 0xe6, 0xaf, 0x75, 0x18, 0x7b, 0xf7, 0xeb, 0x1d, 0x6b, 0xc8, 0x12, 0xc0, 0x16, 0xe6, 0xa8,
	0x1f, 0xe1, 0x22, 0x18, 0x3e, 0x1d, 0x25, 0x66, 0x09, 0xc4, 0x25, 0x25, 0xee, 0x47, 0xb1, 0xa3,
	0xfc, 0xfa, 0x4d, 0xfb, 0x4e, 0x29, 0x4b, 0x80, 0xd2, 0x92, 0x44, 0x69, 0x3f, 0xa2, 0xd4, 0xd1,
	0xee, 0xd7, 0xe7, 0x7b, 0xb6, 0xfe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xf6, 0xfe, 0xfa, 0xe6,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LanguageConstantServiceClient is the client API for LanguageConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LanguageConstantServiceClient interface {
	// Returns the requested language constant.
	GetLanguageConstant(ctx context.Context, in *GetLanguageConstantRequest, opts ...grpc.CallOption) (*resources.LanguageConstant, error)
}

type languageConstantServiceClient struct {
	cc *grpc.ClientConn
}

func NewLanguageConstantServiceClient(cc *grpc.ClientConn) LanguageConstantServiceClient {
	return &languageConstantServiceClient{cc}
}

func (c *languageConstantServiceClient) GetLanguageConstant(ctx context.Context, in *GetLanguageConstantRequest, opts ...grpc.CallOption) (*resources.LanguageConstant, error) {
	out := new(resources.LanguageConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.LanguageConstantService/GetLanguageConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LanguageConstantServiceServer is the server API for LanguageConstantService service.
type LanguageConstantServiceServer interface {
	// Returns the requested language constant.
	GetLanguageConstant(context.Context, *GetLanguageConstantRequest) (*resources.LanguageConstant, error)
}

// UnimplementedLanguageConstantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLanguageConstantServiceServer struct {
}

func (*UnimplementedLanguageConstantServiceServer) GetLanguageConstant(ctx context.Context, req *GetLanguageConstantRequest) (*resources.LanguageConstant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLanguageConstant not implemented")
}

func RegisterLanguageConstantServiceServer(s *grpc.Server, srv LanguageConstantServiceServer) {
	s.RegisterService(&_LanguageConstantService_serviceDesc, srv)
}

func _LanguageConstantService_GetLanguageConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLanguageConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LanguageConstantServiceServer).GetLanguageConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.LanguageConstantService/GetLanguageConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanguageConstantServiceServer).GetLanguageConstant(ctx, req.(*GetLanguageConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LanguageConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.LanguageConstantService",
	HandlerType: (*LanguageConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLanguageConstant",
			Handler:    _LanguageConstantService_GetLanguageConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/language_constant_service.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/language_constant_service.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
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

// Request message for [LanguageConstantService.GetLanguageConstant][google.ads.googleads.v2.services.LanguageConstantService.GetLanguageConstant].
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
	return fileDescriptor_bc548c57933bf456, []int{0}
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
	proto.RegisterType((*GetLanguageConstantRequest)(nil), "google.ads.googleads.v2.services.GetLanguageConstantRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/language_constant_service.proto", fileDescriptor_bc548c57933bf456)
}

var fileDescriptor_bc548c57933bf456 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x4a, 0xc3, 0x40,
	0x14, 0xc6, 0x49, 0x04, 0xc1, 0xa0, 0x9b, 0xb8, 0x68, 0x89, 0x5d, 0x94, 0x5a, 0xf0, 0xcf, 0x62,
	0x06, 0xd2, 0x95, 0xa3, 0x82, 0xa9, 0x8b, 0xba, 0x10, 0x29, 0x15, 0xba, 0x90, 0x40, 0x19, 0x93,
	0x21, 0x04, 0x92, 0x99, 0x9a, 0x97, 0x76, 0x23, 0x6e, 0xf4, 0x08, 0xde, 0xc0, 0xa5, 0x37, 0xf0,
	0x0a, 0xdd, 0x7a, 0x05, 0x57, 0x2e, 0x3d, 0x81, 0xa4, 0x33, 0x93, 0x6a, 0x69, 0xe8, 0xee, 0x63,
	0xde, 0xf7, 0x7e, 0xef, 0xbd, 0x2f, 0xb1, 0x2e, 0x22, 0x21, 0xa2, 0x84, 0x61, 0x1a, 0x02, 0x96,
	0xb2, 0x50, 0x53, 0x17, 0x03, 0xcb, 0xa6, 0x71, 0xc0, 0x00, 0x27, 0x94, 0x47, 0x13, 0x1a, 0xb1,
	0x51, 0x20, 0x38, 0xe4, 0x94, 0xe7, 0x23, 0x55, 0x42, 0xe3, 0x4c, 0xe4, 0xc2, 0x6e, 0xca, 0x36,
	0x44, 0x43, 0x40, 0x25, 0x01, 0x4d, 0x5d, 0xa4, 0x09, 0xce, 0x49, 0xd5, 0x8c, 0x8c, 0x81, 0x98,
	0x64, 0x2b, 0x87, 0x48, 0xb8, 0xd3, 0xd0, 0xad, 0xe3, 0x18, 0x53, 0xce, 0x45, 0x4e, 0xf3, 0x58,
	0x70, 0x50, 0xd5, 0xda, 0x9f, 0x6a, 0x90, 0xc4, 0x4c, 0xb7, 0xb5, 0x3c, 0xcb, 0xe9, 0xb1, 0xfc,
	0x5a, 0x41, 0x2f, 0x15, 0x73, 0xc0, 0x1e, 0x26, 0x0c, 0x72, 0x7b, 0xdf, 0xda, 0xd1, 0x93, 0x47,
	0x9c, 0xa6, 0xac, 0x6e, 0x34, 0x8d, 0xc3, 0xad, 0xc1, 0xb6, 0x7e, 0xbc, 0xa1, 0x29, 0x73, 0x7f,
	0x0c, 0xab, 0xb6, 0x0c, 0xb8, 0x95, 0x17, 0xd9, 0x1f, 0x86, 0xb5, 0xbb, 0x82, 0x6f, 0x9f, 0xa1,
	0x75, 0x59, 0xa0, 0xea, 0xb5, 0x9c, 0x4e, 0x65, 0x77, 0x99, 0x13, 0x5a, 0xee, 0x6d, 0xe1, 0xe7,
	0xcf, 0xaf, 0x57, 0xf3, 0xc8, 0x3e, 0x28, 0xf2, 0x7c, 0xfc, 0x77, 0xd6, 0x79, 0xb2, 0x64, 0x06,
	0x7c, 0xfc, 0xe4, 0xec, 0xcd, 0xbc, 0xfa, 0x02, 0xae, 0xd4, 0x38, 0x06, 0x14, 0x88, 0xb4, 0xfb,
	0x62, 0x5a, 0xed, 0x40, 0xa4, 0x6b, 0xcf, 0xe8, 0x36, 0x2a, 0xa2, 0xe9, 0x17, 0xf1, 0xf7, 0x8d,
	0xbb, 0x2b, 0x45, 0x88, 0x44, 0xb1, 0x04, 0x12, 0x59, 0x84, 0x23, 0xc6, 0xe7, 0x1f, 0x07, 0x2f,
	0x66, 0x56, 0xff, 0x75, 0xa7, 0x5a, 0xbc, 0x99, 0x1b, 0x3d, 0xcf, 0x7b, 0x37, 0x9b, 0x3d, 0x09,
	0xf4, 0x42, 0x40, 0x52, 0x16, 0x6a, 0xe8, 0x22, 0x35, 0x18, 0x66, 0xda, 0xe2, 0x7b, 0x21, 0xf8,
	0xa5, 0xc5, 0x1f, 0xba, 0xbe, 0xb6, 0x7c, 0x9b, 0x6d, 0xf9, 0x4e, 0x88, 0x17, 0x02, 0x21, 0xa5,
	0x89, 0x90, 0xa1, 0x4b, 0x88, 0xb6, 0xdd, 0x6f, 0xce, 0xf7, 0xec, 0xfc, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x90, 0xf1, 0x3f, 0x1e, 0x1c, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.LanguageConstantService/GetLanguageConstant", in, out, opts...)
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
		FullMethod: "/google.ads.googleads.v2.services.LanguageConstantService/GetLanguageConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LanguageConstantServiceServer).GetLanguageConstant(ctx, req.(*GetLanguageConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LanguageConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.LanguageConstantService",
	HandlerType: (*LanguageConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLanguageConstant",
			Handler:    _LanguageConstantService_GetLanguageConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/language_constant_service.proto",
}

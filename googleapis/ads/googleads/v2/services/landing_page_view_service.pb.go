// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/landing_page_view_service.proto

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

// Request message for [LandingPageViewService.GetLandingPageView][google.ads.googleads.v2.services.LandingPageViewService.GetLandingPageView].
type GetLandingPageViewRequest struct {
	// The resource name of the landing page view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetLandingPageViewRequest) Reset()         { *m = GetLandingPageViewRequest{} }
func (m *GetLandingPageViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetLandingPageViewRequest) ProtoMessage()    {}
func (*GetLandingPageViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c82fb57b246174b, []int{0}
}

func (m *GetLandingPageViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLandingPageViewRequest.Unmarshal(m, b)
}
func (m *GetLandingPageViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLandingPageViewRequest.Marshal(b, m, deterministic)
}
func (m *GetLandingPageViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLandingPageViewRequest.Merge(m, src)
}
func (m *GetLandingPageViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetLandingPageViewRequest.Size(m)
}
func (m *GetLandingPageViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLandingPageViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLandingPageViewRequest proto.InternalMessageInfo

func (m *GetLandingPageViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetLandingPageViewRequest)(nil), "google.ads.googleads.v2.services.GetLandingPageViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/landing_page_view_service.proto", fileDescriptor_4c82fb57b246174b)
}

var fileDescriptor_4c82fb57b246174b = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xbf, 0x0a, 0xd3, 0x40,
	0x18, 0x27, 0x11, 0x04, 0x83, 0x2e, 0x19, 0xb4, 0xa6, 0x0e, 0xa5, 0x76, 0x90, 0x0e, 0x77, 0x70,
	0xba, 0x78, 0x45, 0x68, 0xba, 0xd4, 0x41, 0xa4, 0x54, 0xc8, 0x20, 0x81, 0x70, 0x26, 0x1f, 0xc7,
	0x41, 0x72, 0x17, 0x73, 0x69, 0x3a, 0x88, 0x4b, 0x5f, 0xc1, 0x37, 0x70, 0xf4, 0x2d, 0x5c, 0xbb,
	0xfa, 0x0a, 0x4e, 0xbe, 0x81, 0x9b, 0xa4, 0x97, 0x4b, 0xb5, 0x36, 0x74, 0xfb, 0x71, 0xdf, 0xef,
	0xcf, 0x7d, 0xbf, 0x3b, 0x6f, 0xc9, 0x95, 0xe2, 0x39, 0x60, 0x96, 0x69, 0x6c, 0x60, 0x8b, 0x1a,
	0x82, 0x35, 0x54, 0x8d, 0x48, 0x41, 0xe3, 0x9c, 0xc9, 0x4c, 0x48, 0x9e, 0x94, 0x8c, 0x43, 0xd2,
	0x08, 0xd8, 0x27, 0xdd, 0x08, 0x95, 0x95, 0xaa, 0x95, 0x3f, 0x31, 0x32, 0xc4, 0x32, 0x8d, 0x7a,
	0x07, 0xd4, 0x10, 0x64, 0x1d, 0x82, 0x97, 0x43, 0x19, 0x15, 0x68, 0xb5, 0xab, 0xae, 0x86, 0x18,
	0xf3, 0xe0, 0x89, 0x95, 0x96, 0x02, 0x33, 0x29, 0x55, 0xcd, 0x6a, 0xa1, 0xa4, 0xee, 0xa6, 0x8f,
	0xfe, 0x9a, 0xa6, 0xb9, 0x00, 0x59, 0x9b, 0xc1, 0x74, 0xe9, 0x3d, 0x5e, 0x43, 0xfd, 0xc6, 0x98,
	0x6e, 0x18, 0x87, 0x48, 0xc0, 0x7e, 0x0b, 0x1f, 0x77, 0xa0, 0x6b, 0xff, 0xa9, 0xf7, 0xc0, 0x06,
	0x27, 0x92, 0x15, 0x30, 0x72, 0x26, 0xce, 0xb3, 0x7b, 0xdb, 0xfb, 0xf6, 0xf0, 0x2d, 0x2b, 0x80,
	0xfc, 0x76, 0xbc, 0x87, 0x17, 0xfa, 0x77, 0x66, 0x1f, 0xff, 0xbb, 0xe3, 0xf9, 0xff, 0xbb, 0xfb,
	0x0b, 0x74, 0xab, 0x08, 0x34, 0x78, 0xa7, 0x80, 0x0c, 0x8a, 0xfb, 0x8e, 0xd0, 0x85, 0x74, 0x4a,
	0x0f, 0x3f, 0x7e, 0x7e, 0x71, 0x5f, 0xf8, 0xa4, 0xad, 0xf2, 0xd3, 0x3f, 0x2b, 0xbd, 0x4a, 0x77,
	0xba, 0x56, 0x05, 0x54, 0x1a, 0xcf, 0x6d, 0xb7, 0x56, 0xa7, 0xf1, 0xfc, 0x73, 0x30, 0x3e, 0x86,
	0xa3, 0x73, 0x4c, 0x87, 0x4a, 0xa1, 0x51, 0xaa, 0x8a, 0xd5, 0xc1, 0xf5, 0x66, 0xa9, 0x2a, 0x6e,
	0xee, 0xb3, 0x1a, 0x5f, 0x6f, 0x68, 0xd3, 0xbe, 0xc1, 0xc6, 0x79, 0xff, 0xba, 0x33, 0xe0, 0x2a,
	0x67, 0x92, 0x23, 0x55, 0x71, 0xcc, 0x41, 0x9e, 0x5e, 0x08, 0x9f, 0x23, 0x87, 0xbf, 0xde, 0xc2,
	0x82, 0xaf, 0xee, 0x9d, 0x75, 0x18, 0x7e, 0x73, 0x27, 0x6b, 0x63, 0x18, 0x66, 0x1a, 0x19, 0xd8,
	0xa2, 0x88, 0xa0, 0x2e, 0x58, 0x1f, 0x2d, 0x25, 0x0e, 0x33, 0x1d, 0xf7, 0x94, 0x38, 0x22, 0xb1,
	0xa5, 0xfc, 0x72, 0x67, 0xe6, 0x9c, 0xd2, 0x30, 0xd3, 0x94, 0xf6, 0x24, 0x4a, 0x23, 0x42, 0xa9,
	0xa5, 0x7d, 0xb8, 0x7b, 0xba, 0xe7, 0xf3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x21, 0xdd, 0xdc,
	0x37, 0x21, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LandingPageViewServiceClient is the client API for LandingPageViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LandingPageViewServiceClient interface {
	// Returns the requested landing page view in full detail.
	GetLandingPageView(ctx context.Context, in *GetLandingPageViewRequest, opts ...grpc.CallOption) (*resources.LandingPageView, error)
}

type landingPageViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewLandingPageViewServiceClient(cc *grpc.ClientConn) LandingPageViewServiceClient {
	return &landingPageViewServiceClient{cc}
}

func (c *landingPageViewServiceClient) GetLandingPageView(ctx context.Context, in *GetLandingPageViewRequest, opts ...grpc.CallOption) (*resources.LandingPageView, error) {
	out := new(resources.LandingPageView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.LandingPageViewService/GetLandingPageView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LandingPageViewServiceServer is the server API for LandingPageViewService service.
type LandingPageViewServiceServer interface {
	// Returns the requested landing page view in full detail.
	GetLandingPageView(context.Context, *GetLandingPageViewRequest) (*resources.LandingPageView, error)
}

// UnimplementedLandingPageViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLandingPageViewServiceServer struct {
}

func (*UnimplementedLandingPageViewServiceServer) GetLandingPageView(ctx context.Context, req *GetLandingPageViewRequest) (*resources.LandingPageView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLandingPageView not implemented")
}

func RegisterLandingPageViewServiceServer(s *grpc.Server, srv LandingPageViewServiceServer) {
	s.RegisterService(&_LandingPageViewService_serviceDesc, srv)
}

func _LandingPageViewService_GetLandingPageView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLandingPageViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LandingPageViewServiceServer).GetLandingPageView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.LandingPageViewService/GetLandingPageView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LandingPageViewServiceServer).GetLandingPageView(ctx, req.(*GetLandingPageViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LandingPageViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.LandingPageViewService",
	HandlerType: (*LandingPageViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLandingPageView",
			Handler:    _LandingPageViewService_GetLandingPageView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/landing_page_view_service.proto",
}

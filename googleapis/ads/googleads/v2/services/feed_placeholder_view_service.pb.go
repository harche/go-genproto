// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/feed_placeholder_view_service.proto

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

// Request message for [FeedPlaceholderViewService.GetFeedPlaceholderView][google.ads.googleads.v2.services.FeedPlaceholderViewService.GetFeedPlaceholderView].
type GetFeedPlaceholderViewRequest struct {
	// The resource name of the feed placeholder view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedPlaceholderViewRequest) Reset()         { *m = GetFeedPlaceholderViewRequest{} }
func (m *GetFeedPlaceholderViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetFeedPlaceholderViewRequest) ProtoMessage()    {}
func (*GetFeedPlaceholderViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d866ea96d2752d11, []int{0}
}

func (m *GetFeedPlaceholderViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedPlaceholderViewRequest.Unmarshal(m, b)
}
func (m *GetFeedPlaceholderViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedPlaceholderViewRequest.Marshal(b, m, deterministic)
}
func (m *GetFeedPlaceholderViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedPlaceholderViewRequest.Merge(m, src)
}
func (m *GetFeedPlaceholderViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetFeedPlaceholderViewRequest.Size(m)
}
func (m *GetFeedPlaceholderViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedPlaceholderViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedPlaceholderViewRequest proto.InternalMessageInfo

func (m *GetFeedPlaceholderViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetFeedPlaceholderViewRequest)(nil), "google.ads.googleads.v2.services.GetFeedPlaceholderViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/feed_placeholder_view_service.proto", fileDescriptor_d866ea96d2752d11)
}

var fileDescriptor_d866ea96d2752d11 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x31, 0x6b, 0xdb, 0x40,
	0x14, 0x46, 0x2a, 0x14, 0x2a, 0xda, 0x45, 0x43, 0x6b, 0xd4, 0x96, 0x1a, 0xd7, 0x43, 0xf1, 0x70,
	0x07, 0x2a, 0x98, 0x72, 0xc5, 0x2d, 0x32, 0xa6, 0xee, 0x14, 0x8c, 0x03, 0x1a, 0x82, 0x40, 0x5c,
	0x74, 0x2f, 0x8a, 0x40, 0xd2, 0x29, 0x3a, 0x59, 0x1e, 0x42, 0x96, 0x0c, 0xd9, 0x32, 0xe5, 0x1f,
	0x64, 0xcc, 0x4f, 0xf1, 0x16, 0xf2, 0x17, 0x32, 0xe5, 0x57, 0x04, 0xf9, 0x74, 0x72, 0x12, 0xac,
	0x78, 0xfb, 0xb8, 0xf7, 0xbd, 0xef, 0x7b, 0xef, 0x7b, 0x67, 0x4c, 0x42, 0xce, 0xc3, 0x18, 0x30,
	0x65, 0x02, 0x4b, 0x58, 0xa1, 0xd2, 0xc6, 0x02, 0xf2, 0x32, 0x0a, 0x40, 0xe0, 0x23, 0x00, 0xe6,
	0x67, 0x31, 0x0d, 0xe0, 0x98, 0xc7, 0x0c, 0x72, 0xbf, 0x8c, 0x60, 0xe9, 0xd7, 0x65, 0x94, 0xe5,
	0xbc, 0xe0, 0x66, 0x57, 0xb6, 0x22, 0xca, 0x04, 0x6a, 0x54, 0x50, 0x69, 0x23, 0xa5, 0x62, 0x8d,
	0xda, 0x7c, 0x72, 0x10, 0x7c, 0x91, 0xb7, 0x1a, 0x49, 0x03, 0xeb, 0x8b, 0x6a, 0xcf, 0x22, 0x4c,
	0xd3, 0x94, 0x17, 0xb4, 0x88, 0x78, 0x2a, 0xea, 0xea, 0xa7, 0x27, 0xd5, 0x20, 0x8e, 0x20, 0x2d,
	0x64, 0xa1, 0x37, 0x31, 0xbe, 0x4e, 0xa1, 0xf8, 0x07, 0xc0, 0x66, 0x1b, 0x5d, 0x37, 0x82, 0xe5,
	0x1c, 0x4e, 0x16, 0x20, 0x0a, 0xf3, 0xbb, 0xf1, 0x41, 0x0d, 0xe0, 0xa7, 0x34, 0x81, 0x8e, 0xd6,
	0xd5, 0x7e, 0xbc, 0x9b, 0xbf, 0x57, 0x8f, 0x7b, 0x34, 0x01, 0xfb, 0x52, 0x37, 0xac, 0x2d, 0x1a,
	0xfb, 0x72, 0x37, 0xf3, 0x56, 0x33, 0x3e, 0x6e, 0x77, 0x31, 0xff, 0xa2, 0x5d, 0xc1, 0xa0, 0x57,
	0xe7, 0xb3, 0x86, 0xad, 0x02, 0x4d, 0x6e, 0x68, 0x4b, 0x7b, 0xef, 0xcf, 0xf9, 0xdd, 0xfd, 0x95,
	0xfe, 0xcb, 0x1c, 0x56, 0x11, 0x9f, 0x3e, 0x5b, 0x71, 0x14, 0x2c, 0x44, 0xc1, 0x13, 0xc8, 0x05,
	0x1e, 0xac, 0x33, 0x7f, 0xd1, 0x2b, 0xf0, 0xe0, 0xcc, 0xfa, 0xbc, 0x72, 0x3a, 0x1b, 0xbb, 0x1a,
	0x65, 0x91, 0x40, 0x01, 0x4f, 0xc6, 0x17, 0xba, 0xd1, 0x0f, 0x78, 0xb2, 0x73, 0xb7, 0xf1, 0xb7,
	0xf6, 0xd4, 0x66, 0xd5, 0x7d, 0x66, 0xda, 0xc1, 0xff, 0x5a, 0x24, 0xe4, 0x31, 0x4d, 0x43, 0xc4,
	0xf3, 0x10, 0x87, 0x90, 0xae, 0xaf, 0x87, 0x37, 0xb6, 0xed, 0xdf, 0xf3, 0xb7, 0x02, 0xd7, 0xfa,
	0x9b, 0xa9, 0xe3, 0xdc, 0xe8, 0xdd, 0xa9, 0x14, 0x74, 0x98, 0x40, 0x12, 0x56, 0xc8, 0xb5, 0x51,
	0x6d, 0x2c, 0x56, 0x8a, 0xe2, 0x39, 0x4c, 0x78, 0x0d, 0xc5, 0x73, 0x6d, 0x4f, 0x51, 0x1e, 0xf4,
	0xbe, 0x7c, 0x27, 0xc4, 0x61, 0x82, 0x90, 0x86, 0x44, 0x88, 0x6b, 0x13, 0xa2, 0x68, 0x87, 0x6f,
	0xd7, 0x73, 0xfe, 0x7c, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x1b, 0xba, 0xbf, 0xcc, 0x45, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FeedPlaceholderViewServiceClient is the client API for FeedPlaceholderViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeedPlaceholderViewServiceClient interface {
	// Returns the requested feed placeholder view in full detail.
	GetFeedPlaceholderView(ctx context.Context, in *GetFeedPlaceholderViewRequest, opts ...grpc.CallOption) (*resources.FeedPlaceholderView, error)
}

type feedPlaceholderViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewFeedPlaceholderViewServiceClient(cc *grpc.ClientConn) FeedPlaceholderViewServiceClient {
	return &feedPlaceholderViewServiceClient{cc}
}

func (c *feedPlaceholderViewServiceClient) GetFeedPlaceholderView(ctx context.Context, in *GetFeedPlaceholderViewRequest, opts ...grpc.CallOption) (*resources.FeedPlaceholderView, error) {
	out := new(resources.FeedPlaceholderView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.FeedPlaceholderViewService/GetFeedPlaceholderView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedPlaceholderViewServiceServer is the server API for FeedPlaceholderViewService service.
type FeedPlaceholderViewServiceServer interface {
	// Returns the requested feed placeholder view in full detail.
	GetFeedPlaceholderView(context.Context, *GetFeedPlaceholderViewRequest) (*resources.FeedPlaceholderView, error)
}

// UnimplementedFeedPlaceholderViewServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFeedPlaceholderViewServiceServer struct {
}

func (*UnimplementedFeedPlaceholderViewServiceServer) GetFeedPlaceholderView(ctx context.Context, req *GetFeedPlaceholderViewRequest) (*resources.FeedPlaceholderView, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeedPlaceholderView not implemented")
}

func RegisterFeedPlaceholderViewServiceServer(s *grpc.Server, srv FeedPlaceholderViewServiceServer) {
	s.RegisterService(&_FeedPlaceholderViewService_serviceDesc, srv)
}

func _FeedPlaceholderViewService_GetFeedPlaceholderView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedPlaceholderViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedPlaceholderViewServiceServer).GetFeedPlaceholderView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.FeedPlaceholderViewService/GetFeedPlaceholderView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedPlaceholderViewServiceServer).GetFeedPlaceholderView(ctx, req.(*GetFeedPlaceholderViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeedPlaceholderViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.FeedPlaceholderViewService",
	HandlerType: (*FeedPlaceholderViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeedPlaceholderView",
			Handler:    _FeedPlaceholderViewService_GetFeedPlaceholderView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/feed_placeholder_view_service.proto",
}

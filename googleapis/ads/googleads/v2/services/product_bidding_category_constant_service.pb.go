// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/product_bidding_category_constant_service.proto

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

// Request message for
// [ProductBiddingCategoryService.GetProductBiddingCategory][].
type GetProductBiddingCategoryConstantRequest struct {
	// Resource name of the Product Bidding Category to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductBiddingCategoryConstantRequest) Reset() {
	*m = GetProductBiddingCategoryConstantRequest{}
}
func (m *GetProductBiddingCategoryConstantRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductBiddingCategoryConstantRequest) ProtoMessage()    {}
func (*GetProductBiddingCategoryConstantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d5e5b75993b3b91, []int{0}
}

func (m *GetProductBiddingCategoryConstantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductBiddingCategoryConstantRequest.Unmarshal(m, b)
}
func (m *GetProductBiddingCategoryConstantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductBiddingCategoryConstantRequest.Marshal(b, m, deterministic)
}
func (m *GetProductBiddingCategoryConstantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductBiddingCategoryConstantRequest.Merge(m, src)
}
func (m *GetProductBiddingCategoryConstantRequest) XXX_Size() int {
	return xxx_messageInfo_GetProductBiddingCategoryConstantRequest.Size(m)
}
func (m *GetProductBiddingCategoryConstantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductBiddingCategoryConstantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductBiddingCategoryConstantRequest proto.InternalMessageInfo

func (m *GetProductBiddingCategoryConstantRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetProductBiddingCategoryConstantRequest)(nil), "google.ads.googleads.v2.services.GetProductBiddingCategoryConstantRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/product_bidding_category_constant_service.proto", fileDescriptor_8d5e5b75993b3b91)
}

var fileDescriptor_8d5e5b75993b3b91 = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbf, 0x6a, 0xe3, 0x30,
	0x1c, 0xc7, 0xb1, 0x0f, 0x0e, 0xce, 0xdc, 0x2d, 0x5e, 0x2e, 0xf8, 0x6e, 0xc8, 0xe5, 0x52, 0x08,
	0x19, 0x24, 0x70, 0x29, 0x05, 0x95, 0x0c, 0x4e, 0x86, 0xb4, 0x1d, 0x5a, 0x93, 0x42, 0x86, 0x62,
	0x30, 0x8a, 0x25, 0x84, 0x21, 0x96, 0x5c, 0x4b, 0x09, 0x94, 0xd2, 0xa5, 0x6f, 0x50, 0xfa, 0x06,
	0x1d, 0xfb, 0x04, 0x7d, 0x86, 0xac, 0x7d, 0x85, 0x4e, 0xdd, 0xfa, 0x06, 0xc5, 0x91, 0xe5, 0xb4,
	0x43, 0xfe, 0x6c, 0x5f, 0xa4, 0xaf, 0xbe, 0x9f, 0xdf, 0x1f, 0x39, 0x21, 0x13, 0x82, 0x4d, 0x29,
	0xc4, 0x44, 0x42, 0x2d, 0x4b, 0x35, 0xf7, 0xa1, 0xa4, 0xc5, 0x3c, 0x4d, 0xa8, 0x84, 0x79, 0x21,
	0xc8, 0x2c, 0x51, 0xf1, 0x24, 0x25, 0x24, 0xe5, 0x2c, 0x4e, 0xb0, 0xa2, 0x4c, 0x14, 0xd7, 0x71,
	0x22, 0xb8, 0x54, 0x98, 0xab, 0xb8, 0xb2, 0x82, 0xbc, 0x10, 0x4a, 0xb8, 0x4d, 0x1d, 0x03, 0x30,
	0x91, 0xa0, 0x4e, 0x04, 0x73, 0x1f, 0x98, 0x44, 0xef, 0x64, 0x1d, 0xb3, 0xa0, 0x52, 0xcc, 0x8a,
	0x9d, 0xa0, 0x1a, 0xe6, 0xfd, 0x35, 0x51, 0x79, 0x0a, 0x31, 0xe7, 0x42, 0x61, 0x95, 0x0a, 0x2e,
	0xab, 0xdb, 0xdf, 0x9f, 0x6e, 0x93, 0x69, 0x4a, 0xcd, 0xb3, 0xd6, 0xb9, 0xd3, 0x19, 0x52, 0x15,
	0x6a, 0x48, 0x5f, 0x33, 0x06, 0x15, 0x62, 0x50, 0x11, 0x46, 0xf4, 0x6a, 0x46, 0xa5, 0x72, 0xff,
	0x3b, 0xbf, 0x4c, 0x5d, 0x31, 0xc7, 0x19, 0x6d, 0x58, 0x4d, 0xab, 0xf3, 0x63, 0xf4, 0xd3, 0x1c,
	0x9e, 0xe1, 0x8c, 0xfa, 0xcf, 0xb6, 0xb3, 0xb7, 0x39, 0xee, 0x42, 0x77, 0xef, 0xbe, 0x5b, 0xce,
	0xbf, 0xad, 0x6c, 0xf7, 0x14, 0x6c, 0x9b, 0x22, 0xd8, 0xb5, 0x01, 0x2f, 0x58, 0x9b, 0x55, 0xcf,
	0x1b, 0x6c, 0x4e, 0x6a, 0xf5, 0xee, 0x5e, 0x5e, 0x1f, 0xec, 0x43, 0xf7, 0xa0, 0xdc, 0xd2, 0xcd,
	0x97, 0x71, 0xf4, 0xf2, 0x8d, 0x4f, 0x25, 0xec, 0xde, 0x7a, 0x7f, 0x16, 0x41, 0x63, 0x05, 0xae,
	0x54, 0x9e, 0x4a, 0x90, 0x88, 0xac, 0x7f, 0x6f, 0x3b, 0xed, 0x44, 0x64, 0x5b, 0x1b, 0xee, 0x77,
	0x77, 0x1a, 0x70, 0x58, 0x2e, 0x38, 0xb4, 0x2e, 0x8f, 0xab, 0x3c, 0x26, 0xa6, 0x98, 0x33, 0x20,
	0x0a, 0x06, 0x19, 0xe5, 0xcb, 0xf5, 0xc3, 0x55, 0x05, 0xeb, 0xff, 0xfd, 0x91, 0x11, 0x8f, 0xf6,
	0xb7, 0x61, 0x10, 0x3c, 0xd9, 0xcd, 0xa1, 0x0e, 0x0c, 0x88, 0x04, 0x5a, 0x96, 0x6a, 0xec, 0x83,
	0x0a, 0x2c, 0x17, 0xc6, 0x12, 0x05, 0x44, 0x46, 0xb5, 0x25, 0x1a, 0xfb, 0x91, 0xb1, 0xbc, 0xd9,
	0x6d, 0x7d, 0x8e, 0x50, 0x40, 0x24, 0x42, 0xb5, 0x09, 0xa1, 0xb1, 0x8f, 0x90, 0xb1, 0x4d, 0xbe,
	0x2f, 0xeb, 0xdc, 0xff, 0x08, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x7e, 0xd5, 0x8e, 0x9e, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProductBiddingCategoryConstantServiceClient is the client API for ProductBiddingCategoryConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductBiddingCategoryConstantServiceClient interface {
	// Returns the requested Product Bidding Category in full detail.
	GetProductBiddingCategoryConstant(ctx context.Context, in *GetProductBiddingCategoryConstantRequest, opts ...grpc.CallOption) (*resources.ProductBiddingCategoryConstant, error)
}

type productBiddingCategoryConstantServiceClient struct {
	cc *grpc.ClientConn
}

func NewProductBiddingCategoryConstantServiceClient(cc *grpc.ClientConn) ProductBiddingCategoryConstantServiceClient {
	return &productBiddingCategoryConstantServiceClient{cc}
}

func (c *productBiddingCategoryConstantServiceClient) GetProductBiddingCategoryConstant(ctx context.Context, in *GetProductBiddingCategoryConstantRequest, opts ...grpc.CallOption) (*resources.ProductBiddingCategoryConstant, error) {
	out := new(resources.ProductBiddingCategoryConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.ProductBiddingCategoryConstantService/GetProductBiddingCategoryConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductBiddingCategoryConstantServiceServer is the server API for ProductBiddingCategoryConstantService service.
type ProductBiddingCategoryConstantServiceServer interface {
	// Returns the requested Product Bidding Category in full detail.
	GetProductBiddingCategoryConstant(context.Context, *GetProductBiddingCategoryConstantRequest) (*resources.ProductBiddingCategoryConstant, error)
}

// UnimplementedProductBiddingCategoryConstantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductBiddingCategoryConstantServiceServer struct {
}

func (*UnimplementedProductBiddingCategoryConstantServiceServer) GetProductBiddingCategoryConstant(ctx context.Context, req *GetProductBiddingCategoryConstantRequest) (*resources.ProductBiddingCategoryConstant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductBiddingCategoryConstant not implemented")
}

func RegisterProductBiddingCategoryConstantServiceServer(s *grpc.Server, srv ProductBiddingCategoryConstantServiceServer) {
	s.RegisterService(&_ProductBiddingCategoryConstantService_serviceDesc, srv)
}

func _ProductBiddingCategoryConstantService_GetProductBiddingCategoryConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductBiddingCategoryConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductBiddingCategoryConstantServiceServer).GetProductBiddingCategoryConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.ProductBiddingCategoryConstantService/GetProductBiddingCategoryConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductBiddingCategoryConstantServiceServer).GetProductBiddingCategoryConstant(ctx, req.(*GetProductBiddingCategoryConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductBiddingCategoryConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.ProductBiddingCategoryConstantService",
	HandlerType: (*ProductBiddingCategoryConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProductBiddingCategoryConstant",
			Handler:    _ProductBiddingCategoryConstantService_GetProductBiddingCategoryConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/product_bidding_category_constant_service.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/ad_service.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

// Request message for [AdService.GetAd][google.ads.googleads.v2.services.AdService.GetAd].
type GetAdRequest struct {
	// The resource name of the ad to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdRequest) Reset()         { *m = GetAdRequest{} }
func (m *GetAdRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdRequest) ProtoMessage()    {}
func (*GetAdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b4b1037e2d66699, []int{0}
}

func (m *GetAdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdRequest.Unmarshal(m, b)
}
func (m *GetAdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdRequest.Marshal(b, m, deterministic)
}
func (m *GetAdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdRequest.Merge(m, src)
}
func (m *GetAdRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdRequest.Size(m)
}
func (m *GetAdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdRequest proto.InternalMessageInfo

func (m *GetAdRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [AdService.MutateAds][google.ads.googleads.v2.services.AdService.MutateAds].
type MutateAdsRequest struct {
	// The ID of the customer whose ads are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual ads.
	Operations           []*AdOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MutateAdsRequest) Reset()         { *m = MutateAdsRequest{} }
func (m *MutateAdsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateAdsRequest) ProtoMessage()    {}
func (*MutateAdsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b4b1037e2d66699, []int{1}
}

func (m *MutateAdsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdsRequest.Unmarshal(m, b)
}
func (m *MutateAdsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdsRequest.Marshal(b, m, deterministic)
}
func (m *MutateAdsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdsRequest.Merge(m, src)
}
func (m *MutateAdsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateAdsRequest.Size(m)
}
func (m *MutateAdsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdsRequest proto.InternalMessageInfo

func (m *MutateAdsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateAdsRequest) GetOperations() []*AdOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

// A single update operation on an ad.
type AdOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*AdOperation_Update
	Operation            isAdOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *AdOperation) Reset()         { *m = AdOperation{} }
func (m *AdOperation) String() string { return proto.CompactTextString(m) }
func (*AdOperation) ProtoMessage()    {}
func (*AdOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b4b1037e2d66699, []int{2}
}

func (m *AdOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdOperation.Unmarshal(m, b)
}
func (m *AdOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdOperation.Marshal(b, m, deterministic)
}
func (m *AdOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdOperation.Merge(m, src)
}
func (m *AdOperation) XXX_Size() int {
	return xxx_messageInfo_AdOperation.Size(m)
}
func (m *AdOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AdOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AdOperation proto.InternalMessageInfo

func (m *AdOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isAdOperation_Operation interface {
	isAdOperation_Operation()
}

type AdOperation_Update struct {
	Update *resources.Ad `protobuf:"bytes,1,opt,name=update,proto3,oneof"`
}

func (*AdOperation_Update) isAdOperation_Operation() {}

func (m *AdOperation) GetOperation() isAdOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *AdOperation) GetUpdate() *resources.Ad {
	if x, ok := m.GetOperation().(*AdOperation_Update); ok {
		return x.Update
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdOperation_Update)(nil),
	}
}

// Response message for an ad mutate.
type MutateAdsResponse struct {
	// All results for the mutate.
	Results              []*MutateAdResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MutateAdsResponse) Reset()         { *m = MutateAdsResponse{} }
func (m *MutateAdsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateAdsResponse) ProtoMessage()    {}
func (*MutateAdsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b4b1037e2d66699, []int{3}
}

func (m *MutateAdsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdsResponse.Unmarshal(m, b)
}
func (m *MutateAdsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdsResponse.Marshal(b, m, deterministic)
}
func (m *MutateAdsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdsResponse.Merge(m, src)
}
func (m *MutateAdsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateAdsResponse.Size(m)
}
func (m *MutateAdsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdsResponse proto.InternalMessageInfo

func (m *MutateAdsResponse) GetResults() []*MutateAdResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the ad mutate.
type MutateAdResult struct {
	// The resource name returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdResult) Reset()         { *m = MutateAdResult{} }
func (m *MutateAdResult) String() string { return proto.CompactTextString(m) }
func (*MutateAdResult) ProtoMessage()    {}
func (*MutateAdResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b4b1037e2d66699, []int{4}
}

func (m *MutateAdResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdResult.Unmarshal(m, b)
}
func (m *MutateAdResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdResult.Marshal(b, m, deterministic)
}
func (m *MutateAdResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdResult.Merge(m, src)
}
func (m *MutateAdResult) XXX_Size() int {
	return xxx_messageInfo_MutateAdResult.Size(m)
}
func (m *MutateAdResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdResult proto.InternalMessageInfo

func (m *MutateAdResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdRequest)(nil), "google.ads.googleads.v2.services.GetAdRequest")
	proto.RegisterType((*MutateAdsRequest)(nil), "google.ads.googleads.v2.services.MutateAdsRequest")
	proto.RegisterType((*AdOperation)(nil), "google.ads.googleads.v2.services.AdOperation")
	proto.RegisterType((*MutateAdsResponse)(nil), "google.ads.googleads.v2.services.MutateAdsResponse")
	proto.RegisterType((*MutateAdResult)(nil), "google.ads.googleads.v2.services.MutateAdResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/ad_service.proto", fileDescriptor_5b4b1037e2d66699)
}

var fileDescriptor_5b4b1037e2d66699 = []byte{
	// 581 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xb1, 0x6f, 0xd3, 0x4e,
	0x18, 0xfd, 0xd9, 0xd5, 0xaf, 0x28, 0xe7, 0x52, 0xc1, 0x2d, 0x44, 0x06, 0x89, 0xc8, 0xb4, 0x22,
	0x8a, 0x94, 0x33, 0x38, 0x62, 0x71, 0x55, 0xa1, 0xcb, 0x40, 0x0a, 0x52, 0xa0, 0x32, 0x52, 0x06,
	0x14, 0x29, 0xba, 0xe6, 0xae, 0x96, 0xd5, 0xd8, 0x67, 0x7c, 0xe7, 0x2c, 0x55, 0x97, 0xee, 0x4c,
	0x2c, 0xcc, 0x6c, 0xf0, 0xa7, 0x74, 0xe5, 0x5f, 0xe8, 0xc4, 0xce, 0x8e, 0xec, 0xf3, 0x39, 0xee,
	0x10, 0x25, 0x6c, 0xcf, 0x77, 0xef, 0xbd, 0xef, 0xdd, 0xf7, 0x7d, 0x09, 0x78, 0x19, 0x72, 0x1e,
	0x2e, 0x98, 0x4b, 0xa8, 0x70, 0x15, 0x2c, 0xd0, 0xd2, 0x73, 0x05, 0xcb, 0x96, 0xd1, 0x9c, 0x09,
	0x97, 0xd0, 0x59, 0x85, 0x51, 0x9a, 0x71, 0xc9, 0x61, 0x47, 0xf1, 0x10, 0xa1, 0x02, 0xd5, 0x12,
	0xb4, 0xf4, 0x90, 0x96, 0xd8, 0xbd, 0x75, 0xa6, 0x19, 0x13, 0x3c, 0xcf, 0x94, 0xab, 0x72, 0xb3,
	0x9f, 0x68, 0x6e, 0x1a, 0xb9, 0x24, 0x49, 0xb8, 0x24, 0x32, 0xe2, 0x89, 0xa8, 0x6e, 0xab, 0x5a,
	0x6e, 0xf9, 0x75, 0x96, 0x9f, 0xbb, 0xe7, 0x11, 0x5b, 0xd0, 0x59, 0x4c, 0xc4, 0x45, 0xc5, 0x78,
	0xd4, 0xd0, 0xcf, 0x17, 0x11, 0x4b, 0xa4, 0xba, 0x70, 0x06, 0x60, 0x6f, 0xc4, 0x24, 0xa6, 0x01,
	0xfb, 0x9c, 0x33, 0x21, 0xe1, 0x33, 0x70, 0x5f, 0x97, 0x9f, 0x25, 0x24, 0x66, 0x6d, 0xa3, 0x63,
	0x74, 0x5b, 0xc1, 0x9e, 0x3e, 0x7c, 0x4f, 0x62, 0xe6, 0x5c, 0x1b, 0xe0, 0xc1, 0x38, 0x97, 0x44,
	0x32, 0x4c, 0x85, 0x56, 0x3e, 0x05, 0xd6, 0x3c, 0x17, 0x92, 0xc7, 0x2c, 0x9b, 0x45, 0xb4, 0xd2,
	0x01, 0x7d, 0xf4, 0x96, 0xc2, 0x31, 0x00, 0x3c, 0x65, 0x99, 0x4a, 0xde, 0x36, 0x3b, 0x3b, 0x5d,
	0xcb, 0xeb, 0xa3, 0x4d, 0x6d, 0x42, 0x98, 0x7e, 0xd0, 0xaa, 0xa0, 0x61, 0xe0, 0x7c, 0x33, 0x80,
	0xd5, 0xb8, 0x83, 0x47, 0xc0, 0xca, 0x53, 0x4a, 0x24, 0x2b, 0xdf, 0xdd, 0x36, 0x3b, 0x46, 0xd7,
	0xf2, 0x6c, 0xed, 0xaf, 0x5b, 0x83, 0xde, 0x14, 0xad, 0x19, 0x13, 0x71, 0x11, 0x00, 0x45, 0x2f,
	0x30, 0x7c, 0x0d, 0x76, 0xd5, 0x57, 0x99, 0xdb, 0xf2, 0x0e, 0xd7, 0xe6, 0xaa, 0x87, 0x83, 0x30,
	0x3d, 0xf9, 0x2f, 0xa8, 0x64, 0x43, 0x0b, 0xb4, 0xea, 0x6c, 0xce, 0x0c, 0x3c, 0x6c, 0xb4, 0x47,
	0xa4, 0x3c, 0x11, 0x0c, 0xbe, 0x03, 0xf7, 0x32, 0x26, 0xf2, 0x85, 0xd4, 0x6f, 0x7f, 0xb1, 0xf9,
	0xed, 0xda, 0x25, 0x28, 0x85, 0x81, 0x36, 0x70, 0x5e, 0x81, 0xfd, 0xbb, 0x57, 0x5b, 0xcd, 0xcd,
	0xbb, 0x35, 0x41, 0x0b, 0xd3, 0x8f, 0xca, 0x1d, 0x7e, 0x31, 0xc0, 0xff, 0xe5, 0xec, 0x21, 0xda,
	0x9c, 0xa4, 0xb9, 0x24, 0xf6, 0x76, 0xdd, 0x71, 0xfa, 0xd7, 0xbf, 0x6e, 0xbf, 0x9a, 0xcf, 0xe1,
	0x61, 0xb1, 0xd4, 0x97, 0x77, 0xe2, 0x1d, 0xeb, 0xbd, 0x10, 0x6e, 0xaf, 0xfc, 0x01, 0xf4, 0xae,
	0xe0, 0x0f, 0x03, 0xb4, 0xea, 0xb6, 0x41, 0x6f, 0xfb, 0xee, 0xe8, 0x15, 0xb4, 0x07, 0xff, 0xa4,
	0x51, 0x73, 0x71, 0x06, 0x65, 0xca, 0xbe, 0xd3, 0x2d, 0x52, 0xae, 0x62, 0x5d, 0x36, 0x96, 0xf9,
	0xb8, 0x77, 0x55, 0x84, 0xf4, 0xe3, 0x52, 0xec, 0x1b, 0x3d, 0xfb, 0xf1, 0x0d, 0x6e, 0xaf, 0x0a,
	0x54, 0x28, 0x8d, 0x04, 0x9a, 0xf3, 0x78, 0xf8, 0xc7, 0x00, 0x07, 0x73, 0x1e, 0x6f, 0x0c, 0x33,
	0xdc, 0xaf, 0x87, 0x71, 0x5a, 0xac, 0xe7, 0xa9, 0xf1, 0xe9, 0xa4, 0xd2, 0x84, 0x7c, 0x41, 0x92,
	0x10, 0xf1, 0x2c, 0x74, 0x43, 0x96, 0x94, 0xcb, 0xeb, 0xae, 0xaa, 0xac, 0xff, 0x1f, 0x3a, 0xd2,
	0xe0, 0xbb, 0xb9, 0x33, 0xc2, 0xf8, 0xa7, 0xd9, 0x19, 0x29, 0x43, 0x4c, 0x05, 0x52, 0xb0, 0x40,
	0x13, 0x0f, 0x55, 0x85, 0xc5, 0x8d, 0xa6, 0x4c, 0x31, 0x15, 0xd3, 0x9a, 0x32, 0x9d, 0x78, 0x53,
	0x4d, 0xf9, 0x6d, 0x1e, 0xa8, 0x73, 0xdf, 0xc7, 0x54, 0xf8, 0x7e, 0x4d, 0xf2, 0xfd, 0x89, 0xe7,
	0xfb, 0x9a, 0x76, 0xb6, 0x5b, 0xe6, 0x1c, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x49, 0xeb,
	0xdf, 0x2e, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdServiceClient is the client API for AdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdServiceClient interface {
	// Returns the requested ad in full detail.
	GetAd(ctx context.Context, in *GetAdRequest, opts ...grpc.CallOption) (*resources.Ad, error)
	// Updates ads. Operation statuses are returned.
	MutateAds(ctx context.Context, in *MutateAdsRequest, opts ...grpc.CallOption) (*MutateAdsResponse, error)
}

type adServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdServiceClient(cc *grpc.ClientConn) AdServiceClient {
	return &adServiceClient{cc}
}

func (c *adServiceClient) GetAd(ctx context.Context, in *GetAdRequest, opts ...grpc.CallOption) (*resources.Ad, error) {
	out := new(resources.Ad)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.AdService/GetAd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adServiceClient) MutateAds(ctx context.Context, in *MutateAdsRequest, opts ...grpc.CallOption) (*MutateAdsResponse, error) {
	out := new(MutateAdsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.AdService/MutateAds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdServiceServer is the server API for AdService service.
type AdServiceServer interface {
	// Returns the requested ad in full detail.
	GetAd(context.Context, *GetAdRequest) (*resources.Ad, error)
	// Updates ads. Operation statuses are returned.
	MutateAds(context.Context, *MutateAdsRequest) (*MutateAdsResponse, error)
}

// UnimplementedAdServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdServiceServer struct {
}

func (*UnimplementedAdServiceServer) GetAd(ctx context.Context, req *GetAdRequest) (*resources.Ad, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAd not implemented")
}
func (*UnimplementedAdServiceServer) MutateAds(ctx context.Context, req *MutateAdsRequest) (*MutateAdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MutateAds not implemented")
}

func RegisterAdServiceServer(s *grpc.Server, srv AdServiceServer) {
	s.RegisterService(&_AdService_serviceDesc, srv)
}

func _AdService_GetAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdServiceServer).GetAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.AdService/GetAd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdServiceServer).GetAd(ctx, req.(*GetAdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdService_MutateAds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdServiceServer).MutateAds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.AdService/MutateAds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdServiceServer).MutateAds(ctx, req.(*MutateAdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.AdService",
	HandlerType: (*AdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAd",
			Handler:    _AdService_GetAd_Handler,
		},
		{
			MethodName: "MutateAds",
			Handler:    _AdService_MutateAds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/ad_service.proto",
}

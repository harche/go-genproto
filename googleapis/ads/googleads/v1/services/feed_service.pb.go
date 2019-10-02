// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/feed_service.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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
// [FeedService.GetFeed][google.ads.googleads.v1.services.FeedService.GetFeed].
type GetFeedRequest struct {
	// The resource name of the feed to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedRequest) Reset()         { *m = GetFeedRequest{} }
func (m *GetFeedRequest) String() string { return proto.CompactTextString(m) }
func (*GetFeedRequest) ProtoMessage()    {}
func (*GetFeedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b158f21a9b50a59, []int{0}
}

func (m *GetFeedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedRequest.Unmarshal(m, b)
}
func (m *GetFeedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedRequest.Marshal(b, m, deterministic)
}
func (m *GetFeedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedRequest.Merge(m, src)
}
func (m *GetFeedRequest) XXX_Size() int {
	return xxx_messageInfo_GetFeedRequest.Size(m)
}
func (m *GetFeedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedRequest proto.InternalMessageInfo

func (m *GetFeedRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for
// [FeedService.MutateFeeds][google.ads.googleads.v1.services.FeedService.MutateFeeds].
type MutateFeedsRequest struct {
	// The ID of the customer whose feeds are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual feeds.
	Operations []*FeedOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateFeedsRequest) Reset()         { *m = MutateFeedsRequest{} }
func (m *MutateFeedsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateFeedsRequest) ProtoMessage()    {}
func (*MutateFeedsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b158f21a9b50a59, []int{1}
}

func (m *MutateFeedsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedsRequest.Unmarshal(m, b)
}
func (m *MutateFeedsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedsRequest.Marshal(b, m, deterministic)
}
func (m *MutateFeedsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedsRequest.Merge(m, src)
}
func (m *MutateFeedsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateFeedsRequest.Size(m)
}
func (m *MutateFeedsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedsRequest proto.InternalMessageInfo

func (m *MutateFeedsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateFeedsRequest) GetOperations() []*FeedOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateFeedsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateFeedsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on an feed.
type FeedOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*FeedOperation_Create
	//	*FeedOperation_Update
	//	*FeedOperation_Remove
	Operation            isFeedOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *FeedOperation) Reset()         { *m = FeedOperation{} }
func (m *FeedOperation) String() string { return proto.CompactTextString(m) }
func (*FeedOperation) ProtoMessage()    {}
func (*FeedOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b158f21a9b50a59, []int{2}
}

func (m *FeedOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedOperation.Unmarshal(m, b)
}
func (m *FeedOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedOperation.Marshal(b, m, deterministic)
}
func (m *FeedOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedOperation.Merge(m, src)
}
func (m *FeedOperation) XXX_Size() int {
	return xxx_messageInfo_FeedOperation.Size(m)
}
func (m *FeedOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedOperation.DiscardUnknown(m)
}

var xxx_messageInfo_FeedOperation proto.InternalMessageInfo

func (m *FeedOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isFeedOperation_Operation interface {
	isFeedOperation_Operation()
}

type FeedOperation_Create struct {
	Create *resources.Feed `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type FeedOperation_Update struct {
	Update *resources.Feed `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type FeedOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*FeedOperation_Create) isFeedOperation_Operation() {}

func (*FeedOperation_Update) isFeedOperation_Operation() {}

func (*FeedOperation_Remove) isFeedOperation_Operation() {}

func (m *FeedOperation) GetOperation() isFeedOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *FeedOperation) GetCreate() *resources.Feed {
	if x, ok := m.GetOperation().(*FeedOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *FeedOperation) GetUpdate() *resources.Feed {
	if x, ok := m.GetOperation().(*FeedOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *FeedOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*FeedOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FeedOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FeedOperation_Create)(nil),
		(*FeedOperation_Update)(nil),
		(*FeedOperation_Remove)(nil),
	}
}

// Response message for an feed mutate.
type MutateFeedsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateFeedResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MutateFeedsResponse) Reset()         { *m = MutateFeedsResponse{} }
func (m *MutateFeedsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateFeedsResponse) ProtoMessage()    {}
func (*MutateFeedsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b158f21a9b50a59, []int{3}
}

func (m *MutateFeedsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedsResponse.Unmarshal(m, b)
}
func (m *MutateFeedsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedsResponse.Marshal(b, m, deterministic)
}
func (m *MutateFeedsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedsResponse.Merge(m, src)
}
func (m *MutateFeedsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateFeedsResponse.Size(m)
}
func (m *MutateFeedsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedsResponse proto.InternalMessageInfo

func (m *MutateFeedsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateFeedsResponse) GetResults() []*MutateFeedResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the feed mutate.
type MutateFeedResult struct {
	// Returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateFeedResult) Reset()         { *m = MutateFeedResult{} }
func (m *MutateFeedResult) String() string { return proto.CompactTextString(m) }
func (*MutateFeedResult) ProtoMessage()    {}
func (*MutateFeedResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b158f21a9b50a59, []int{4}
}

func (m *MutateFeedResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateFeedResult.Unmarshal(m, b)
}
func (m *MutateFeedResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateFeedResult.Marshal(b, m, deterministic)
}
func (m *MutateFeedResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateFeedResult.Merge(m, src)
}
func (m *MutateFeedResult) XXX_Size() int {
	return xxx_messageInfo_MutateFeedResult.Size(m)
}
func (m *MutateFeedResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateFeedResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateFeedResult proto.InternalMessageInfo

func (m *MutateFeedResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetFeedRequest)(nil), "google.ads.googleads.v1.services.GetFeedRequest")
	proto.RegisterType((*MutateFeedsRequest)(nil), "google.ads.googleads.v1.services.MutateFeedsRequest")
	proto.RegisterType((*FeedOperation)(nil), "google.ads.googleads.v1.services.FeedOperation")
	proto.RegisterType((*MutateFeedsResponse)(nil), "google.ads.googleads.v1.services.MutateFeedsResponse")
	proto.RegisterType((*MutateFeedResult)(nil), "google.ads.googleads.v1.services.MutateFeedResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/feed_service.proto", fileDescriptor_3b158f21a9b50a59)
}

var fileDescriptor_3b158f21a9b50a59 = []byte{
	// 694 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0x76, 0xb7, 0xd2, 0xda, 0xd9, 0xb6, 0x96, 0x29, 0x62, 0x08, 0xa2, 0x61, 0x2d, 0xb4, 0x06,
	0xd9, 0x31, 0xa9, 0x45, 0xd8, 0xd2, 0x43, 0x0a, 0xa6, 0x15, 0xac, 0x2d, 0x5b, 0xe8, 0x41, 0x02,
	0xcb, 0x34, 0x3b, 0x09, 0x4b, 0x77, 0x77, 0xd6, 0x99, 0xd9, 0x48, 0x29, 0xbd, 0xf8, 0x17, 0xc4,
	0x3f, 0xe0, 0x51, 0x4f, 0xfe, 0x0d, 0xaf, 0x7a, 0xf1, 0xee, 0xc9, 0x3f, 0x20, 0xde, 0x64, 0x66,
	0x76, 0xd2, 0xa4, 0x50, 0xd2, 0xdc, 0xde, 0xbe, 0xf9, 0xbe, 0xef, 0x7d, 0xf3, 0xde, 0xbe, 0x01,
	0x1b, 0x7d, 0x4a, 0xfb, 0x09, 0x41, 0x38, 0xe2, 0x48, 0x87, 0x32, 0x1a, 0x34, 0x10, 0x27, 0x6c,
	0x10, 0x77, 0x09, 0x47, 0x3d, 0x42, 0xa2, 0xb0, 0xfc, 0xf2, 0x72, 0x46, 0x05, 0x85, 0x35, 0x8d,
	0xf4, 0x70, 0xc4, 0xbd, 0x21, 0xc9, 0x1b, 0x34, 0x3c, 0x43, 0xaa, 0x3e, 0xbd, 0x4e, 0x96, 0x11,
	0x4e, 0x0b, 0x66, 0x74, 0xb5, 0x5e, 0xf5, 0x81, 0x41, 0xe7, 0x31, 0xc2, 0x59, 0x46, 0x05, 0x16,
	0x31, 0xcd, 0x78, 0x79, 0x5a, 0x56, 0x43, 0xea, 0xeb, 0xa4, 0xe8, 0xa1, 0x5e, 0x4c, 0x92, 0x28,
	0x4c, 0x31, 0x3f, 0x2d, 0x11, 0x0f, 0xaf, 0x22, 0xde, 0x33, 0x9c, 0xe7, 0x84, 0x19, 0x85, 0xfb,
	0xe5, 0x39, 0xcb, 0xbb, 0x88, 0x0b, 0x2c, 0x8a, 0xf2, 0xc0, 0xdd, 0x04, 0x4b, 0xbb, 0x44, 0xb4,
	0x09, 0x89, 0x02, 0xf2, 0xae, 0x20, 0x5c, 0xc0, 0xc7, 0x60, 0xd1, 0x58, 0x0c, 0x33, 0x9c, 0x92,
	0x8a, 0x55, 0xb3, 0xd6, 0xe7, 0x83, 0x05, 0x93, 0x7c, 0x83, 0x53, 0xe2, 0xfe, 0xb4, 0x00, 0xdc,
	0x2f, 0x04, 0x16, 0x44, 0x52, 0xb9, 0xe1, 0x3e, 0x02, 0x4e, 0xb7, 0xe0, 0x82, 0xa6, 0x84, 0x85,
	0x71, 0x54, 0x32, 0x81, 0x49, 0xbd, 0x8a, 0xe0, 0x01, 0x00, 0x34, 0x27, 0x4c, 0xdf, 0xae, 0x62,
	0xd7, 0x66, 0xd6, 0x9d, 0x26, 0xf2, 0x26, 0x35, 0xd3, 0x93, 0x45, 0x0e, 0x0c, 0x2f, 0x18, 0x91,
	0x80, 0x6b, 0xe0, 0x6e, 0x8e, 0x99, 0x88, 0x71, 0x12, 0xf6, 0x70, 0x9c, 0x14, 0x8c, 0x54, 0x66,
	0x6a, 0xd6, 0xfa, 0x9d, 0x60, 0xa9, 0x4c, 0xb7, 0x75, 0x56, 0x5e, 0x6b, 0x80, 0x93, 0x38, 0xc2,
	0x82, 0x84, 0x34, 0x4b, 0xce, 0x2a, 0xb7, 0x15, 0x6c, 0xc1, 0x24, 0x0f, 0xb2, 0xe4, 0xcc, 0xfd,
	0x67, 0x81, 0xc5, 0xb1, 0x5a, 0x70, 0x0b, 0x38, 0x45, 0xae, 0x48, 0xb2, 0xdb, 0x8a, 0xe4, 0x34,
	0xab, 0xc6, 0xb1, 0x69, 0xb7, 0xd7, 0x96, 0x03, 0xd9, 0xc7, 0xfc, 0x34, 0x00, 0x1a, 0x2e, 0x63,
	0xd8, 0x02, 0xb3, 0x5d, 0x46, 0xb0, 0xd0, 0x3d, 0x74, 0x9a, 0x6b, 0xd7, 0xde, 0x74, 0xf8, 0x53,
	0xa8, 0xab, 0xee, 0xdd, 0x0a, 0x4a, 0xa2, 0x94, 0xd0, 0x82, 0x15, 0x7b, 0x6a, 0x09, 0x4d, 0x84,
	0x15, 0x30, 0xcb, 0x48, 0x4a, 0x07, 0xba, 0x33, 0xf3, 0xf2, 0x44, 0x7f, 0xef, 0x38, 0x60, 0x7e,
	0xd8, 0x4a, 0xf7, 0xab, 0x05, 0x56, 0xc6, 0x46, 0xca, 0x73, 0x9a, 0x71, 0x02, 0xdb, 0xe0, 0xde,
	0x95, 0x0e, 0x87, 0x84, 0x31, 0xca, 0x94, 0x9a, 0xd3, 0x84, 0xc6, 0x10, 0xcb, 0xbb, 0xde, 0x91,
	0xfa, 0xb5, 0x82, 0x95, 0xf1, 0xde, 0xbf, 0x94, 0x70, 0xf8, 0x1a, 0xcc, 0x31, 0xc2, 0x8b, 0x44,
	0x98, 0xb9, 0x37, 0x27, 0xcf, 0xfd, 0xd2, 0x4f, 0xa0, 0xa8, 0x81, 0x91, 0x70, 0x5f, 0x80, 0xe5,
	0xab, 0x87, 0x37, 0xfa, 0x73, 0x9b, 0xbf, 0x6c, 0xe0, 0x48, 0xce, 0x91, 0xae, 0x01, 0x3f, 0x59,
	0x60, 0xae, 0xdc, 0x00, 0xf8, 0x6c, 0xb2, 0xa3, 0xf1, 0x65, 0xa9, 0xde, 0x74, 0x1c, 0x2e, 0xfa,
	0xf0, 0xe3, 0xf7, 0x47, 0xfb, 0x09, 0x5c, 0x93, 0x4f, 0xc0, 0xf9, 0x98, 0xcd, 0x6d, 0xb3, 0x1f,
	0x1c, 0xd5, 0xd5, 0x9b, 0xc0, 0x51, 0xfd, 0x02, 0x7e, 0xb3, 0x80, 0x33, 0x32, 0x0e, 0xf8, 0x7c,
	0x9a, 0x6e, 0x99, 0x85, 0xac, 0x6e, 0x4e, 0xc9, 0xd2, 0x33, 0x77, 0x37, 0x95, 0x5b, 0xe4, 0xd6,
	0xa5, 0xdb, 0x4b, 0x7b, 0xe7, 0x23, 0xcb, 0xbd, 0x5d, 0xbf, 0xd0, 0x66, 0xfd, 0x54, 0x09, 0xf8,
	0x56, 0x7d, 0xe7, 0xaf, 0x05, 0x56, 0xbb, 0x34, 0x9d, 0x58, 0x73, 0x67, 0x79, 0x64, 0x02, 0x87,
	0x72, 0x87, 0x0e, 0xad, 0xb7, 0x7b, 0x25, 0xab, 0x4f, 0x13, 0x9c, 0xf5, 0x3d, 0xca, 0xfa, 0xa8,
	0x4f, 0x32, 0xb5, 0x61, 0xe6, 0x01, 0xcd, 0x63, 0x7e, 0xfd, 0x33, 0xbd, 0x65, 0x82, 0xcf, 0xf6,
	0xcc, 0x6e, 0xab, 0xf5, 0xc5, 0xae, 0xed, 0x6a, 0xc1, 0x56, 0xc4, 0x3d, 0x1d, 0xca, 0xe8, 0xb8,
	0xe1, 0x95, 0x85, 0xf9, 0x77, 0x03, 0xe9, 0xb4, 0x22, 0xde, 0x19, 0x42, 0x3a, 0xc7, 0x8d, 0x8e,
	0x81, 0xfc, 0xb1, 0x57, 0x75, 0xde, 0xf7, 0x5b, 0x11, 0xf7, 0xfd, 0x21, 0xc8, 0xf7, 0x8f, 0x1b,
	0xbe, 0x6f, 0x60, 0x27, 0xb3, 0xca, 0xe7, 0xc6, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x05,
	0x04, 0x92, 0x4d, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FeedServiceClient is the client API for FeedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeedServiceClient interface {
	// Returns the requested feed in full detail.
	GetFeed(ctx context.Context, in *GetFeedRequest, opts ...grpc.CallOption) (*resources.Feed, error)
	// Creates, updates, or removes feeds. Operation statuses are
	// returned.
	MutateFeeds(ctx context.Context, in *MutateFeedsRequest, opts ...grpc.CallOption) (*MutateFeedsResponse, error)
}

type feedServiceClient struct {
	cc *grpc.ClientConn
}

func NewFeedServiceClient(cc *grpc.ClientConn) FeedServiceClient {
	return &feedServiceClient{cc}
}

func (c *feedServiceClient) GetFeed(ctx context.Context, in *GetFeedRequest, opts ...grpc.CallOption) (*resources.Feed, error) {
	out := new(resources.Feed)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.FeedService/GetFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedServiceClient) MutateFeeds(ctx context.Context, in *MutateFeedsRequest, opts ...grpc.CallOption) (*MutateFeedsResponse, error) {
	out := new(MutateFeedsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.FeedService/MutateFeeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedServiceServer is the server API for FeedService service.
type FeedServiceServer interface {
	// Returns the requested feed in full detail.
	GetFeed(context.Context, *GetFeedRequest) (*resources.Feed, error)
	// Creates, updates, or removes feeds. Operation statuses are
	// returned.
	MutateFeeds(context.Context, *MutateFeedsRequest) (*MutateFeedsResponse, error)
}

// UnimplementedFeedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFeedServiceServer struct {
}

func (*UnimplementedFeedServiceServer) GetFeed(ctx context.Context, req *GetFeedRequest) (*resources.Feed, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetFeed not implemented")
}
func (*UnimplementedFeedServiceServer) MutateFeeds(ctx context.Context, req *MutateFeedsRequest) (*MutateFeedsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateFeeds not implemented")
}

func RegisterFeedServiceServer(s *grpc.Server, srv FeedServiceServer) {
	s.RegisterService(&_FeedService_serviceDesc, srv)
}

func _FeedService_GetFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).GetFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.FeedService/GetFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).GetFeed(ctx, req.(*GetFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedService_MutateFeeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateFeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedServiceServer).MutateFeeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.FeedService/MutateFeeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedServiceServer).MutateFeeds(ctx, req.(*MutateFeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeedService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.FeedService",
	HandlerType: (*FeedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFeed",
			Handler:    _FeedService_GetFeed_Handler,
		},
		{
			MethodName: "MutateFeeds",
			Handler:    _FeedService_MutateFeeds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/feed_service.proto",
}

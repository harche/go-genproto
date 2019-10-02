// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/ad_group_ad_service.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "google.golang.org/genproto/googleapis/ads/googleads/v2/common"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
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

// Request message for [AdGroupAdService.GetAdGroupAd][google.ads.googleads.v2.services.AdGroupAdService.GetAdGroupAd].
type GetAdGroupAdRequest struct {
	// The resource name of the ad to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdGroupAdRequest) Reset()         { *m = GetAdGroupAdRequest{} }
func (m *GetAdGroupAdRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdGroupAdRequest) ProtoMessage()    {}
func (*GetAdGroupAdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_756c28aeb001c0ab, []int{0}
}

func (m *GetAdGroupAdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdGroupAdRequest.Unmarshal(m, b)
}
func (m *GetAdGroupAdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdGroupAdRequest.Marshal(b, m, deterministic)
}
func (m *GetAdGroupAdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdGroupAdRequest.Merge(m, src)
}
func (m *GetAdGroupAdRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdGroupAdRequest.Size(m)
}
func (m *GetAdGroupAdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdGroupAdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdGroupAdRequest proto.InternalMessageInfo

func (m *GetAdGroupAdRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [AdGroupAdService.MutateAdGroupAds][google.ads.googleads.v2.services.AdGroupAdService.MutateAdGroupAds].
type MutateAdGroupAdsRequest struct {
	// The ID of the customer whose ads are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual ads.
	Operations []*AdGroupAdOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
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

func (m *MutateAdGroupAdsRequest) Reset()         { *m = MutateAdGroupAdsRequest{} }
func (m *MutateAdGroupAdsRequest) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupAdsRequest) ProtoMessage()    {}
func (*MutateAdGroupAdsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_756c28aeb001c0ab, []int{1}
}

func (m *MutateAdGroupAdsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupAdsRequest.Unmarshal(m, b)
}
func (m *MutateAdGroupAdsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupAdsRequest.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupAdsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupAdsRequest.Merge(m, src)
}
func (m *MutateAdGroupAdsRequest) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupAdsRequest.Size(m)
}
func (m *MutateAdGroupAdsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupAdsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupAdsRequest proto.InternalMessageInfo

func (m *MutateAdGroupAdsRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateAdGroupAdsRequest) GetOperations() []*AdGroupAdOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateAdGroupAdsRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateAdGroupAdsRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on an ad group ad.
type AdGroupAdOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// Configuration for how policies are validated.
	PolicyValidationParameter *common.PolicyValidationParameter `protobuf:"bytes,5,opt,name=policy_validation_parameter,json=policyValidationParameter,proto3" json:"policy_validation_parameter,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*AdGroupAdOperation_Create
	//	*AdGroupAdOperation_Update
	//	*AdGroupAdOperation_Remove
	Operation            isAdGroupAdOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *AdGroupAdOperation) Reset()         { *m = AdGroupAdOperation{} }
func (m *AdGroupAdOperation) String() string { return proto.CompactTextString(m) }
func (*AdGroupAdOperation) ProtoMessage()    {}
func (*AdGroupAdOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_756c28aeb001c0ab, []int{2}
}

func (m *AdGroupAdOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupAdOperation.Unmarshal(m, b)
}
func (m *AdGroupAdOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupAdOperation.Marshal(b, m, deterministic)
}
func (m *AdGroupAdOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupAdOperation.Merge(m, src)
}
func (m *AdGroupAdOperation) XXX_Size() int {
	return xxx_messageInfo_AdGroupAdOperation.Size(m)
}
func (m *AdGroupAdOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupAdOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupAdOperation proto.InternalMessageInfo

func (m *AdGroupAdOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

func (m *AdGroupAdOperation) GetPolicyValidationParameter() *common.PolicyValidationParameter {
	if m != nil {
		return m.PolicyValidationParameter
	}
	return nil
}

type isAdGroupAdOperation_Operation interface {
	isAdGroupAdOperation_Operation()
}

type AdGroupAdOperation_Create struct {
	Create *resources.AdGroupAd `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AdGroupAdOperation_Update struct {
	Update *resources.AdGroupAd `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type AdGroupAdOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*AdGroupAdOperation_Create) isAdGroupAdOperation_Operation() {}

func (*AdGroupAdOperation_Update) isAdGroupAdOperation_Operation() {}

func (*AdGroupAdOperation_Remove) isAdGroupAdOperation_Operation() {}

func (m *AdGroupAdOperation) GetOperation() isAdGroupAdOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *AdGroupAdOperation) GetCreate() *resources.AdGroupAd {
	if x, ok := m.GetOperation().(*AdGroupAdOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *AdGroupAdOperation) GetUpdate() *resources.AdGroupAd {
	if x, ok := m.GetOperation().(*AdGroupAdOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *AdGroupAdOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*AdGroupAdOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdGroupAdOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdGroupAdOperation_Create)(nil),
		(*AdGroupAdOperation_Update)(nil),
		(*AdGroupAdOperation_Remove)(nil),
	}
}

// Response message for an ad group ad mutate.
type MutateAdGroupAdsResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateAdGroupAdResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *MutateAdGroupAdsResponse) Reset()         { *m = MutateAdGroupAdsResponse{} }
func (m *MutateAdGroupAdsResponse) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupAdsResponse) ProtoMessage()    {}
func (*MutateAdGroupAdsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_756c28aeb001c0ab, []int{3}
}

func (m *MutateAdGroupAdsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupAdsResponse.Unmarshal(m, b)
}
func (m *MutateAdGroupAdsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupAdsResponse.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupAdsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupAdsResponse.Merge(m, src)
}
func (m *MutateAdGroupAdsResponse) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupAdsResponse.Size(m)
}
func (m *MutateAdGroupAdsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupAdsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupAdsResponse proto.InternalMessageInfo

func (m *MutateAdGroupAdsResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateAdGroupAdsResponse) GetResults() []*MutateAdGroupAdResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the ad mutate.
type MutateAdGroupAdResult struct {
	// The resource name returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdGroupAdResult) Reset()         { *m = MutateAdGroupAdResult{} }
func (m *MutateAdGroupAdResult) String() string { return proto.CompactTextString(m) }
func (*MutateAdGroupAdResult) ProtoMessage()    {}
func (*MutateAdGroupAdResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_756c28aeb001c0ab, []int{4}
}

func (m *MutateAdGroupAdResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdGroupAdResult.Unmarshal(m, b)
}
func (m *MutateAdGroupAdResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdGroupAdResult.Marshal(b, m, deterministic)
}
func (m *MutateAdGroupAdResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdGroupAdResult.Merge(m, src)
}
func (m *MutateAdGroupAdResult) XXX_Size() int {
	return xxx_messageInfo_MutateAdGroupAdResult.Size(m)
}
func (m *MutateAdGroupAdResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdGroupAdResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdGroupAdResult proto.InternalMessageInfo

func (m *MutateAdGroupAdResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdGroupAdRequest)(nil), "google.ads.googleads.v2.services.GetAdGroupAdRequest")
	proto.RegisterType((*MutateAdGroupAdsRequest)(nil), "google.ads.googleads.v2.services.MutateAdGroupAdsRequest")
	proto.RegisterType((*AdGroupAdOperation)(nil), "google.ads.googleads.v2.services.AdGroupAdOperation")
	proto.RegisterType((*MutateAdGroupAdsResponse)(nil), "google.ads.googleads.v2.services.MutateAdGroupAdsResponse")
	proto.RegisterType((*MutateAdGroupAdResult)(nil), "google.ads.googleads.v2.services.MutateAdGroupAdResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/ad_group_ad_service.proto", fileDescriptor_756c28aeb001c0ab)
}

var fileDescriptor_756c28aeb001c0ab = []byte{
	// 775 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x41, 0x4f, 0xdb, 0x48,
	0x14, 0x5e, 0x27, 0xbb, 0xec, 0x32, 0x61, 0x77, 0xd1, 0x20, 0x84, 0x1b, 0x2a, 0x35, 0x72, 0x91,
	0x8a, 0x52, 0x64, 0x4b, 0x86, 0xaa, 0xc2, 0x94, 0x43, 0x90, 0x4a, 0xe8, 0x81, 0x92, 0x9a, 0x2a,
	0x87, 0x2a, 0x92, 0x35, 0xd8, 0x43, 0x64, 0x61, 0x7b, 0xdc, 0x99, 0x71, 0xa4, 0x08, 0x71, 0x69,
	0x7f, 0x42, 0xff, 0x41, 0x7b, 0x6a, 0xef, 0xfd, 0x09, 0xbd, 0x70, 0xed, 0xbd, 0xa7, 0x9e, 0xfa,
	0x1b, 0x7a, 0xa8, 0xc6, 0xe3, 0x31, 0x21, 0x10, 0xa5, 0x70, 0x7b, 0x7e, 0xf3, 0x7d, 0xdf, 0xfb,
	0x66, 0xde, 0x9b, 0x31, 0x70, 0xfa, 0x84, 0xf4, 0x23, 0x6c, 0xa1, 0x80, 0x59, 0x32, 0x14, 0xd1,
	0xc0, 0xb6, 0x18, 0xa6, 0x83, 0xd0, 0xc7, 0xcc, 0x42, 0x81, 0xd7, 0xa7, 0x24, 0x4b, 0x3d, 0x14,
	0x78, 0x45, 0xd2, 0x4c, 0x29, 0xe1, 0x04, 0x36, 0x24, 0xc1, 0x44, 0x01, 0x33, 0x4b, 0xae, 0x39,
	0xb0, 0x4d, 0xc5, 0xad, 0x3f, 0x9c, 0xa4, 0xee, 0x93, 0x38, 0x26, 0x89, 0x95, 0x92, 0x28, 0xf4,
	0x87, 0x52, 0xae, 0xbe, 0x3e, 0x09, 0x4c, 0x31, 0x23, 0x19, 0x1d, 0xf3, 0x52, 0x90, 0xee, 0x2a,
	0x52, 0x1a, 0x5a, 0x28, 0x49, 0x08, 0x47, 0x3c, 0x24, 0x09, 0x2b, 0x56, 0x0b, 0x87, 0x56, 0xfe,
	0x75, 0x94, 0x1d, 0x5b, 0xc7, 0x21, 0x8e, 0x02, 0x2f, 0x46, 0xec, 0xa4, 0x40, 0x2c, 0x15, 0x08,
	0x9a, 0xfa, 0x16, 0xe3, 0x88, 0x67, 0x6c, 0x6c, 0x41, 0x08, 0xfb, 0x51, 0x88, 0x13, 0x2e, 0x17,
	0x0c, 0x07, 0x2c, 0xb4, 0x31, 0x6f, 0x05, 0x6d, 0x61, 0xa4, 0x15, 0xb8, 0xf8, 0x75, 0x86, 0x19,
	0x87, 0xf7, 0xc1, 0xbf, 0xca, 0xa7, 0x97, 0xa0, 0x18, 0xeb, 0x5a, 0x43, 0x5b, 0x9d, 0x75, 0xe7,
	0x54, 0xf2, 0x39, 0x8a, 0xb1, 0xf1, 0x4d, 0x03, 0x4b, 0xfb, 0x19, 0x47, 0x1c, 0x97, 0x7c, 0xa6,
	0x04, 0xee, 0x81, 0x9a, 0x9f, 0x31, 0x4e, 0x62, 0x4c, 0xbd, 0x30, 0x28, 0xe8, 0x40, 0xa5, 0x9e,
	0x05, 0xf0, 0x25, 0x00, 0x24, 0xc5, 0x54, 0x6e, 0x50, 0xaf, 0x34, 0xaa, 0xab, 0x35, 0x7b, 0xc3,
	0x9c, 0xd6, 0x03, 0xb3, 0xac, 0x74, 0xa0, 0xc8, 0xee, 0x88, 0x0e, 0x7c, 0x00, 0xfe, 0x4f, 0x11,
	0xe5, 0x21, 0x8a, 0xbc, 0x63, 0x14, 0x46, 0x19, 0xc5, 0x7a, 0xb5, 0xa1, 0xad, 0xfe, 0xe3, 0xfe,
	0x57, 0xa4, 0x77, 0x65, 0x56, 0x6c, 0x70, 0x80, 0xa2, 0x30, 0x40, 0x1c, 0x7b, 0x24, 0x89, 0x86,
	0xfa, 0x9f, 0x39, 0x6c, 0x4e, 0x25, 0x0f, 0x92, 0x68, 0x68, 0xbc, 0xad, 0x02, 0x78, 0xb5, 0x20,
	0xdc, 0x02, 0xb5, 0x2c, 0xcd, 0x99, 0xe2, 0xe8, 0x73, 0x66, 0xcd, 0xae, 0x2b, 0xef, 0xaa, 0x3b,
	0xe6, 0xae, 0xe8, 0xce, 0x3e, 0x62, 0x27, 0x2e, 0x90, 0x70, 0x11, 0xc3, 0x21, 0x58, 0x96, 0x73,
	0xe2, 0x15, 0xa5, 0x42, 0x92, 0x78, 0x29, 0xa2, 0x28, 0xc6, 0x1c, 0x53, 0xfd, 0xaf, 0x5c, 0x6c,
	0x73, 0xe2, 0x41, 0xc8, 0x51, 0x33, 0x3b, 0xb9, 0x44, 0xb7, 0x54, 0xe8, 0x28, 0x01, 0xf7, 0x4e,
	0x3a, 0x69, 0x09, 0xee, 0x82, 0x19, 0x9f, 0x62, 0xc4, 0x65, 0x37, 0x6b, 0xf6, 0xda, 0xc4, 0x2a,
	0xe5, 0x8c, 0x5e, 0x9c, 0xf7, 0xde, 0x1f, 0x6e, 0xc1, 0x16, 0x3a, 0x72, 0x43, 0x7a, 0xe5, 0x76,
	0x3a, 0x92, 0x0d, 0x75, 0x30, 0x43, 0x71, 0x4c, 0x06, 0xb2, 0x47, 0xb3, 0x62, 0x45, 0x7e, 0xef,
	0xd4, 0xc0, 0x6c, 0xd9, 0x54, 0xe3, 0xb3, 0x06, 0xf4, 0xab, 0x63, 0xc6, 0x52, 0x92, 0x30, 0xe1,
	0x65, 0x71, 0xac, 0xe1, 0x1e, 0xa6, 0x94, 0xd0, 0x5c, 0xb2, 0x66, 0x43, 0x65, 0x8d, 0xa6, 0xbe,
	0x79, 0x98, 0xdf, 0x08, 0x77, 0xe1, 0xf2, 0x28, 0x3c, 0x15, 0x70, 0xf8, 0x02, 0xfc, 0x4d, 0x31,
	0xcb, 0x22, 0xae, 0x66, 0xf1, 0xf1, 0xf4, 0x59, 0x1c, 0x33, 0xe5, 0xe6, 0x7c, 0x57, 0xe9, 0x18,
	0x4f, 0xc0, 0xe2, 0xb5, 0x88, 0xdf, 0xba, 0x5c, 0xf6, 0x87, 0x2a, 0x98, 0x2f, 0x89, 0x87, 0xb2,
	0x24, 0xfc, 0xa8, 0x81, 0xb9, 0xd1, 0xeb, 0x0a, 0x1f, 0x4d, 0x77, 0x79, 0xcd, 0xf5, 0xae, 0xdf,
	0xa8, 0x63, 0xc6, 0xc6, 0x9b, 0xaf, 0xdf, 0xdf, 0x55, 0x4c, 0xb8, 0x26, 0x9e, 0xaf, 0xd3, 0x4b,
	0xd6, 0xb7, 0xd5, 0x8d, 0x66, 0x56, 0xd3, 0x42, 0x65, 0x7b, 0xac, 0xe6, 0x19, 0xfc, 0xa2, 0x81,
	0xf9, 0xf1, 0xb6, 0xc1, 0xcd, 0x1b, 0x9f, 0xaa, 0x7a, 0x51, 0xea, 0xce, 0x6d, 0xa8, 0x72, 0x4a,
	0x0c, 0x27, 0xdf, 0xc1, 0x86, 0x61, 0xe5, 0xaf, 0x75, 0x69, 0xf9, 0x74, 0xe4, 0x89, 0xda, 0x6e,
	0x9e, 0x8d, 0x6c, 0xc0, 0x89, 0x73, 0x29, 0x47, 0x6b, 0xd6, 0x97, 0xcf, 0x5b, 0xfa, 0x45, 0xb9,
	0x22, 0x4a, 0x43, 0x26, 0xae, 0xe2, 0xce, 0x4f, 0x0d, 0xac, 0xf8, 0x24, 0x9e, 0x6a, 0x6d, 0x67,
	0x71, 0xbc, 0x97, 0x1d, 0xf1, 0x4c, 0x74, 0xb4, 0x57, 0x7b, 0x05, 0xb5, 0x4f, 0x22, 0x94, 0xf4,
	0x4d, 0x42, 0xfb, 0x56, 0x1f, 0x27, 0xf9, 0x23, 0x62, 0x5d, 0x14, 0x9b, 0xfc, 0x47, 0xdb, 0x52,
	0xc1, 0xfb, 0x4a, 0xb5, 0xdd, 0x6a, 0x7d, 0xaa, 0x34, 0xda, 0x52, 0xb0, 0x15, 0x30, 0x53, 0x86,
	0x22, 0xea, 0xda, 0x66, 0x51, 0x98, 0x9d, 0x2b, 0x48, 0xaf, 0x15, 0xb0, 0x5e, 0x09, 0xe9, 0x75,
	0xed, 0x9e, 0x82, 0xfc, 0xa8, 0xac, 0xc8, 0xbc, 0xe3, 0x88, 0xc3, 0x70, 0x4a, 0x90, 0xe3, 0x74,
	0x6d, 0xc7, 0x51, 0xb0, 0xa3, 0x99, 0xdc, 0xe7, 0xfa, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3d,
	0x54, 0xe0, 0x0c, 0x78, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdGroupAdServiceClient is the client API for AdGroupAdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdGroupAdServiceClient interface {
	// Returns the requested ad in full detail.
	GetAdGroupAd(ctx context.Context, in *GetAdGroupAdRequest, opts ...grpc.CallOption) (*resources.AdGroupAd, error)
	// Creates, updates, or removes ads. Operation statuses are returned.
	MutateAdGroupAds(ctx context.Context, in *MutateAdGroupAdsRequest, opts ...grpc.CallOption) (*MutateAdGroupAdsResponse, error)
}

type adGroupAdServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdGroupAdServiceClient(cc *grpc.ClientConn) AdGroupAdServiceClient {
	return &adGroupAdServiceClient{cc}
}

func (c *adGroupAdServiceClient) GetAdGroupAd(ctx context.Context, in *GetAdGroupAdRequest, opts ...grpc.CallOption) (*resources.AdGroupAd, error) {
	out := new(resources.AdGroupAd)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.AdGroupAdService/GetAdGroupAd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adGroupAdServiceClient) MutateAdGroupAds(ctx context.Context, in *MutateAdGroupAdsRequest, opts ...grpc.CallOption) (*MutateAdGroupAdsResponse, error) {
	out := new(MutateAdGroupAdsResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.AdGroupAdService/MutateAdGroupAds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdGroupAdServiceServer is the server API for AdGroupAdService service.
type AdGroupAdServiceServer interface {
	// Returns the requested ad in full detail.
	GetAdGroupAd(context.Context, *GetAdGroupAdRequest) (*resources.AdGroupAd, error)
	// Creates, updates, or removes ads. Operation statuses are returned.
	MutateAdGroupAds(context.Context, *MutateAdGroupAdsRequest) (*MutateAdGroupAdsResponse, error)
}

// UnimplementedAdGroupAdServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdGroupAdServiceServer struct {
}

func (*UnimplementedAdGroupAdServiceServer) GetAdGroupAd(ctx context.Context, req *GetAdGroupAdRequest) (*resources.AdGroupAd, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GetAdGroupAd not implemented")
}
func (*UnimplementedAdGroupAdServiceServer) MutateAdGroupAds(ctx context.Context, req *MutateAdGroupAdsRequest) (*MutateAdGroupAdsResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method MutateAdGroupAds not implemented")
}

func RegisterAdGroupAdServiceServer(s *grpc.Server, srv AdGroupAdServiceServer) {
	s.RegisterService(&_AdGroupAdService_serviceDesc, srv)
}

func _AdGroupAdService_GetAdGroupAd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdGroupAdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupAdServiceServer).GetAdGroupAd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.AdGroupAdService/GetAdGroupAd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupAdServiceServer).GetAdGroupAd(ctx, req.(*GetAdGroupAdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdGroupAdService_MutateAdGroupAds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdGroupAdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdGroupAdServiceServer).MutateAdGroupAds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.AdGroupAdService/MutateAdGroupAds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdGroupAdServiceServer).MutateAdGroupAds(ctx, req.(*MutateAdGroupAdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdGroupAdService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.AdGroupAdService",
	HandlerType: (*AdGroupAdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdGroupAd",
			Handler:    _AdGroupAdService_GetAdGroupAd_Handler,
		},
		{
			MethodName: "MutateAdGroupAds",
			Handler:    _AdGroupAdService_MutateAdGroupAds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/ad_group_ad_service.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/customer_feed_error.proto

package errors

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enum describing possible customer feed errors.
type CustomerFeedErrorEnum_CustomerFeedError int32

const (
	// Enum unspecified.
	CustomerFeedErrorEnum_UNSPECIFIED CustomerFeedErrorEnum_CustomerFeedError = 0
	// The received error code is not known in this version.
	CustomerFeedErrorEnum_UNKNOWN CustomerFeedErrorEnum_CustomerFeedError = 1
	// An active feed already exists for this customer and place holder type.
	CustomerFeedErrorEnum_FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE CustomerFeedErrorEnum_CustomerFeedError = 2
	// The specified feed is removed.
	CustomerFeedErrorEnum_CANNOT_CREATE_FOR_REMOVED_FEED CustomerFeedErrorEnum_CustomerFeedError = 3
	// The CustomerFeed already exists. Update should be used to modify the
	// existing CustomerFeed.
	CustomerFeedErrorEnum_CANNOT_CREATE_ALREADY_EXISTING_CUSTOMER_FEED CustomerFeedErrorEnum_CustomerFeedError = 4
	// Cannot update removed customer feed.
	CustomerFeedErrorEnum_CANNOT_MODIFY_REMOVED_CUSTOMER_FEED CustomerFeedErrorEnum_CustomerFeedError = 5
	// Invalid placeholder type.
	CustomerFeedErrorEnum_INVALID_PLACEHOLDER_TYPE CustomerFeedErrorEnum_CustomerFeedError = 6
	// Feed mapping for this placeholder type does not exist.
	CustomerFeedErrorEnum_MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE CustomerFeedErrorEnum_CustomerFeedError = 7
	// Placeholder not allowed at the account level.
	CustomerFeedErrorEnum_PLACEHOLDER_TYPE_NOT_ALLOWED_ON_CUSTOMER_FEED CustomerFeedErrorEnum_CustomerFeedError = 8
)

var CustomerFeedErrorEnum_CustomerFeedError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE",
	3: "CANNOT_CREATE_FOR_REMOVED_FEED",
	4: "CANNOT_CREATE_ALREADY_EXISTING_CUSTOMER_FEED",
	5: "CANNOT_MODIFY_REMOVED_CUSTOMER_FEED",
	6: "INVALID_PLACEHOLDER_TYPE",
	7: "MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE",
	8: "PLACEHOLDER_TYPE_NOT_ALLOWED_ON_CUSTOMER_FEED",
}

var CustomerFeedErrorEnum_CustomerFeedError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"FEED_ALREADY_EXISTS_FOR_PLACEHOLDER_TYPE":      2,
	"CANNOT_CREATE_FOR_REMOVED_FEED":                3,
	"CANNOT_CREATE_ALREADY_EXISTING_CUSTOMER_FEED":  4,
	"CANNOT_MODIFY_REMOVED_CUSTOMER_FEED":           5,
	"INVALID_PLACEHOLDER_TYPE":                      6,
	"MISSING_FEEDMAPPING_FOR_PLACEHOLDER_TYPE":      7,
	"PLACEHOLDER_TYPE_NOT_ALLOWED_ON_CUSTOMER_FEED": 8,
}

func (x CustomerFeedErrorEnum_CustomerFeedError) String() string {
	return proto.EnumName(CustomerFeedErrorEnum_CustomerFeedError_name, int32(x))
}

func (CustomerFeedErrorEnum_CustomerFeedError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b77f587276d28976, []int{0, 0}
}

// Container for enum describing possible customer feed errors.
type CustomerFeedErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerFeedErrorEnum) Reset()         { *m = CustomerFeedErrorEnum{} }
func (m *CustomerFeedErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CustomerFeedErrorEnum) ProtoMessage()    {}
func (*CustomerFeedErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_b77f587276d28976, []int{0}
}

func (m *CustomerFeedErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerFeedErrorEnum.Unmarshal(m, b)
}
func (m *CustomerFeedErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerFeedErrorEnum.Marshal(b, m, deterministic)
}
func (m *CustomerFeedErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerFeedErrorEnum.Merge(m, src)
}
func (m *CustomerFeedErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CustomerFeedErrorEnum.Size(m)
}
func (m *CustomerFeedErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerFeedErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerFeedErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.CustomerFeedErrorEnum_CustomerFeedError", CustomerFeedErrorEnum_CustomerFeedError_name, CustomerFeedErrorEnum_CustomerFeedError_value)
	proto.RegisterType((*CustomerFeedErrorEnum)(nil), "google.ads.googleads.v2.errors.CustomerFeedErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/customer_feed_error.proto", fileDescriptor_b77f587276d28976)
}

var fileDescriptor_b77f587276d28976 = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x69, 0x16, 0x76, 0x91, 0xf7, 0x40, 0xb1, 0x04, 0x42, 0x68, 0xd5, 0x43, 0x38, 0xc0,
	0x61, 0x71, 0x20, 0x5c, 0x50, 0x38, 0x79, 0x63, 0xa7, 0x44, 0x24, 0x76, 0x94, 0xa4, 0x59, 0x8a,
	0x2a, 0x59, 0x61, 0x13, 0xa2, 0x4a, 0xdb, 0xb8, 0x8a, 0xbb, 0xfb, 0x40, 0x1c, 0x79, 0x14, 0xde,
	0x81, 0x17, 0xe0, 0xc6, 0x89, 0x2b, 0x72, 0xdc, 0x56, 0x6a, 0x0b, 0x9c, 0xf2, 0x6b, 0xf2, 0xfd,
	0xff, 0x8c, 0x35, 0x03, 0xde, 0x36, 0x52, 0x36, 0xd7, 0xb5, 0x53, 0x56, 0xca, 0x31, 0x52, 0xab,
	0x5b, 0xd7, 0xa9, 0xbb, 0x4e, 0x76, 0xca, 0xb9, 0xba, 0x51, 0x2b, 0xb9, 0xa8, 0x3b, 0xf1, 0xa5,
	0xae, 0x2b, 0xd1, 0x17, 0xd1, 0xb2, 0x93, 0x2b, 0x09, 0x47, 0x06, 0x47, 0x65, 0xa5, 0xd0, 0xd6,
	0x89, 0x6e, 0x5d, 0x64, 0x9c, 0x4f, 0xcf, 0x36, 0xc9, 0xcb, 0xb9, 0x53, 0xb6, 0xad, 0x5c, 0x95,
	0xab, 0xb9, 0x6c, 0x95, 0x71, 0xdb, 0xbf, 0x2d, 0xf0, 0xc8, 0x5f, 0x67, 0x07, 0x75, 0x5d, 0x51,
	0x6d, 0xa2, 0xed, 0xcd, 0xc2, 0xfe, 0x61, 0x81, 0x87, 0x07, 0x7f, 0xe0, 0x03, 0x70, 0x3a, 0x61,
	0x59, 0x42, 0xfd, 0x30, 0x08, 0x29, 0x19, 0xde, 0x81, 0xa7, 0xe0, 0x64, 0xc2, 0x3e, 0x30, 0x7e,
	0xc9, 0x86, 0x03, 0x78, 0x0e, 0x5e, 0x04, 0x94, 0x12, 0x81, 0xa3, 0x94, 0x62, 0x32, 0x15, 0xf4,
	0x63, 0x98, 0xe5, 0x99, 0x08, 0x78, 0x2a, 0x92, 0x08, 0xfb, 0xf4, 0x3d, 0x8f, 0x08, 0x4d, 0x45,
	0x3e, 0x4d, 0xe8, 0xd0, 0x82, 0x36, 0x18, 0xf9, 0x98, 0x31, 0x9e, 0x0b, 0x3f, 0xa5, 0x38, 0xa7,
	0x3d, 0x97, 0xd2, 0x98, 0x17, 0x94, 0x08, 0x9d, 0x33, 0x3c, 0x82, 0xaf, 0xc0, 0xf9, 0x2e, 0xb3,
	0x13, 0x1d, 0xb2, 0xb1, 0xf0, 0x27, 0x59, 0xce, 0x63, 0x9a, 0x1a, 0xc7, 0x5d, 0xf8, 0x1c, 0x3c,
	0x5b, 0x3b, 0x62, 0x4e, 0xc2, 0x60, 0xba, 0x4d, 0xdc, 0x05, 0xef, 0xc1, 0x33, 0xf0, 0x24, 0x64,
	0x05, 0x8e, 0x42, 0x72, 0x38, 0xdc, 0xb1, 0x7e, 0x4a, 0x1c, 0x66, 0x99, 0xee, 0xa0, 0xf9, 0x18,
	0x27, 0x49, 0xaf, 0xff, 0xf6, 0x94, 0x13, 0xf8, 0x1a, 0xbc, 0xdc, 0xaf, 0x0a, 0x3d, 0x02, 0x8e,
	0x22, 0x7e, 0x49, 0x89, 0xe0, 0x6c, 0xaf, 0xfd, 0xfd, 0x8b, 0x5f, 0x03, 0x60, 0x5f, 0xc9, 0x05,
	0xfa, 0xff, 0xfa, 0x2e, 0x1e, 0x1f, 0xec, 0x20, 0xd1, 0x8b, 0x4b, 0x06, 0x9f, 0xc8, 0xda, 0xd9,
	0xc8, 0xeb, 0xb2, 0x6d, 0x90, 0xec, 0x1a, 0xa7, 0xa9, 0xdb, 0x7e, 0xad, 0x9b, 0x13, 0x5a, 0xce,
	0xd5, 0xbf, 0x2e, 0xea, 0x9d, 0xf9, 0x7c, 0xb5, 0x8e, 0xc6, 0x18, 0x7f, 0xb3, 0x46, 0x63, 0x13,
	0x86, 0x2b, 0x85, 0x8c, 0xd4, 0xaa, 0x70, 0x51, 0xdf, 0x52, 0x7d, 0xdf, 0x00, 0x33, 0x5c, 0xa9,
	0xd9, 0x16, 0x98, 0x15, 0xee, 0xcc, 0x00, 0x3f, 0x2d, 0xdb, 0x54, 0x3d, 0x0f, 0x57, 0xca, 0xf3,
	0xb6, 0x88, 0xe7, 0x15, 0xae, 0xe7, 0x19, 0xe8, 0xf3, 0x71, 0x3f, 0xdd, 0x9b, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x3b, 0xa4, 0xb3, 0x42, 0xee, 0x02, 0x00, 0x00,
}

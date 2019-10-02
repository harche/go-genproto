// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/feed_attribute_reference_error.proto

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

// Enum describing possible feed attribute reference errors.
type FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError int32

const (
	// Enum unspecified.
	FeedAttributeReferenceErrorEnum_UNSPECIFIED FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError = 0
	// The received error code is not known in this version.
	FeedAttributeReferenceErrorEnum_UNKNOWN FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError = 1
	// A feed referenced by ID has been removed.
	FeedAttributeReferenceErrorEnum_CANNOT_REFERENCE_REMOVED_FEED FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError = 2
	// There is no enabled feed with the given name.
	FeedAttributeReferenceErrorEnum_INVALID_FEED_NAME FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError = 3
	// There is no feed attribute in an enabled feed with the given name.
	FeedAttributeReferenceErrorEnum_INVALID_FEED_ATTRIBUTE_NAME FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError = 4
)

var FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CANNOT_REFERENCE_REMOVED_FEED",
	3: "INVALID_FEED_NAME",
	4: "INVALID_FEED_ATTRIBUTE_NAME",
}

var FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError_value = map[string]int32{
	"UNSPECIFIED":                   0,
	"UNKNOWN":                       1,
	"CANNOT_REFERENCE_REMOVED_FEED": 2,
	"INVALID_FEED_NAME":             3,
	"INVALID_FEED_ATTRIBUTE_NAME":   4,
}

func (x FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError) String() string {
	return proto.EnumName(FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError_name, int32(x))
}

func (FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1513e40adc63f345, []int{0, 0}
}

// Container for enum describing possible feed attribute reference errors.
type FeedAttributeReferenceErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedAttributeReferenceErrorEnum) Reset()         { *m = FeedAttributeReferenceErrorEnum{} }
func (m *FeedAttributeReferenceErrorEnum) String() string { return proto.CompactTextString(m) }
func (*FeedAttributeReferenceErrorEnum) ProtoMessage()    {}
func (*FeedAttributeReferenceErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_1513e40adc63f345, []int{0}
}

func (m *FeedAttributeReferenceErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedAttributeReferenceErrorEnum.Unmarshal(m, b)
}
func (m *FeedAttributeReferenceErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedAttributeReferenceErrorEnum.Marshal(b, m, deterministic)
}
func (m *FeedAttributeReferenceErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedAttributeReferenceErrorEnum.Merge(m, src)
}
func (m *FeedAttributeReferenceErrorEnum) XXX_Size() int {
	return xxx_messageInfo_FeedAttributeReferenceErrorEnum.Size(m)
}
func (m *FeedAttributeReferenceErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedAttributeReferenceErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedAttributeReferenceErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError", FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError_name, FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError_value)
	proto.RegisterType((*FeedAttributeReferenceErrorEnum)(nil), "google.ads.googleads.v1.errors.FeedAttributeReferenceErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/feed_attribute_reference_error.proto", fileDescriptor_1513e40adc63f345)
}

var fileDescriptor_1513e40adc63f345 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x6e, 0xa3, 0x30,
	0x14, 0x86, 0x07, 0x32, 0x9a, 0x91, 0x9c, 0xc5, 0x30, 0x48, 0xb3, 0x99, 0xb4, 0x49, 0xcb, 0x01,
	0x8c, 0x50, 0x77, 0xee, 0xca, 0x01, 0x13, 0xa1, 0x36, 0x4e, 0x44, 0x09, 0x95, 0x2a, 0x24, 0x44,
	0x82, 0x83, 0x90, 0x12, 0x1c, 0xd9, 0x24, 0x47, 0xe9, 0x01, 0xba, 0xec, 0x11, 0x7a, 0x84, 0x1e,
	0xa5, 0x57, 0xe8, 0xa6, 0x02, 0x97, 0x48, 0x5d, 0x34, 0x2b, 0xff, 0x7a, 0xfe, 0xde, 0xff, 0x3f,
	0xbd, 0x07, 0xdc, 0x82, 0xf3, 0x62, 0xc3, 0xec, 0x2c, 0x97, 0xb6, 0x92, 0x8d, 0x3a, 0x38, 0x36,
	0x13, 0x82, 0x0b, 0x69, 0xaf, 0x19, 0xcb, 0xd3, 0xac, 0xae, 0x45, 0xb9, 0xdc, 0xd7, 0x2c, 0x15,
	0x6c, 0xcd, 0x04, 0xab, 0x56, 0x2c, 0x6d, 0xff, 0xe1, 0x4e, 0xf0, 0x9a, 0x9b, 0x43, 0xd5, 0x09,
	0xb3, 0x5c, 0xc2, 0xa3, 0x09, 0x3c, 0x38, 0x50, 0x99, 0xfc, 0x3f, 0xeb, 0x42, 0x76, 0xa5, 0x9d,
	0x55, 0x15, 0xaf, 0xb3, 0xba, 0xe4, 0x95, 0x54, 0xdd, 0xd6, 0x8b, 0x06, 0x46, 0x3e, 0x63, 0x39,
	0xee, 0x52, 0xc2, 0x2e, 0x84, 0x34, 0xed, 0xa4, 0xda, 0x6f, 0xad, 0x47, 0x0d, 0x0c, 0x4e, 0x30,
	0xe6, 0x1f, 0xd0, 0x5f, 0xd0, 0xbb, 0x39, 0x71, 0x03, 0x3f, 0x20, 0x9e, 0xf1, 0xc3, 0xec, 0x83,
	0xdf, 0x0b, 0x7a, 0x43, 0x67, 0xf7, 0xd4, 0xd0, 0xcc, 0x4b, 0x70, 0xee, 0x62, 0x4a, 0x67, 0x51,
	0x1a, 0x12, 0x9f, 0x84, 0x84, 0xba, 0x24, 0x0d, 0xc9, 0x74, 0x16, 0x13, 0x2f, 0xf5, 0x09, 0xf1,
	0x0c, 0xdd, 0xfc, 0x07, 0xfe, 0x06, 0x34, 0xc6, 0xb7, 0x81, 0xaa, 0xa4, 0x14, 0x4f, 0x89, 0xd1,
	0x33, 0x47, 0x60, 0xf0, 0xa5, 0x8c, 0xa3, 0x28, 0x0c, 0xc6, 0x8b, 0x88, 0x28, 0xe0, 0xe7, 0xf8,
	0x5d, 0x03, 0xd6, 0x8a, 0x6f, 0xe1, 0xe9, 0x0d, 0x8c, 0x2f, 0x4e, 0x0c, 0x3f, 0x6f, 0xb6, 0x30,
	0xd7, 0x1e, 0xbc, 0x4f, 0x8f, 0x82, 0x6f, 0xb2, 0xaa, 0x80, 0x5c, 0x14, 0x76, 0xc1, 0xaa, 0x76,
	0x47, 0xdd, 0x69, 0x76, 0xa5, 0xfc, 0xee, 0x52, 0xd7, 0xea, 0x79, 0xd2, 0x7b, 0x13, 0x8c, 0x9f,
	0xf5, 0xe1, 0x44, 0x99, 0xe1, 0x5c, 0x42, 0x25, 0x1b, 0x15, 0x3b, 0xb0, 0x8d, 0x94, 0xaf, 0x1d,
	0x90, 0xe0, 0x5c, 0x26, 0x47, 0x20, 0x89, 0x9d, 0x44, 0x01, 0x6f, 0xba, 0xa5, 0xaa, 0x08, 0xe1,
	0x5c, 0x22, 0x74, 0x44, 0x10, 0x8a, 0x1d, 0x84, 0x14, 0xb4, 0xfc, 0xd5, 0x4e, 0x77, 0xf5, 0x11,
	0x00, 0x00, 0xff, 0xff, 0xd6, 0xd2, 0x2a, 0xb7, 0x46, 0x02, 0x00, 0x00,
}

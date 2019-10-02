// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/ad_group_ad_error.proto

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

// Enum describing possible ad group ad errors.
type AdGroupAdErrorEnum_AdGroupAdError int32

const (
	// Enum unspecified.
	AdGroupAdErrorEnum_UNSPECIFIED AdGroupAdErrorEnum_AdGroupAdError = 0
	// The received error code is not known in this version.
	AdGroupAdErrorEnum_UNKNOWN AdGroupAdErrorEnum_AdGroupAdError = 1
	// No link found between the adgroup ad and the label.
	AdGroupAdErrorEnum_AD_GROUP_AD_LABEL_DOES_NOT_EXIST AdGroupAdErrorEnum_AdGroupAdError = 2
	// The label has already been attached to the adgroup ad.
	AdGroupAdErrorEnum_AD_GROUP_AD_LABEL_ALREADY_EXISTS AdGroupAdErrorEnum_AdGroupAdError = 3
	// The specified ad was not found in the adgroup
	AdGroupAdErrorEnum_AD_NOT_UNDER_ADGROUP AdGroupAdErrorEnum_AdGroupAdError = 4
	// Removed ads may not be modified
	AdGroupAdErrorEnum_CANNOT_OPERATE_ON_REMOVED_ADGROUPAD AdGroupAdErrorEnum_AdGroupAdError = 5
	// An ad of this type is deprecated and cannot be created. Only deletions
	// are permitted.
	AdGroupAdErrorEnum_CANNOT_CREATE_DEPRECATED_ADS AdGroupAdErrorEnum_AdGroupAdError = 6
	// Text ads are deprecated and cannot be created. Use expanded text ads
	// instead.
	AdGroupAdErrorEnum_CANNOT_CREATE_TEXT_ADS AdGroupAdErrorEnum_AdGroupAdError = 7
	// A required field was not specified or is an empty string.
	AdGroupAdErrorEnum_EMPTY_FIELD AdGroupAdErrorEnum_AdGroupAdError = 8
	// An ad may only be modified once per call
	AdGroupAdErrorEnum_RESOURCE_REFERENCED_IN_MULTIPLE_OPS AdGroupAdErrorEnum_AdGroupAdError = 9
)

var AdGroupAdErrorEnum_AdGroupAdError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "AD_GROUP_AD_LABEL_DOES_NOT_EXIST",
	3: "AD_GROUP_AD_LABEL_ALREADY_EXISTS",
	4: "AD_NOT_UNDER_ADGROUP",
	5: "CANNOT_OPERATE_ON_REMOVED_ADGROUPAD",
	6: "CANNOT_CREATE_DEPRECATED_ADS",
	7: "CANNOT_CREATE_TEXT_ADS",
	8: "EMPTY_FIELD",
	9: "RESOURCE_REFERENCED_IN_MULTIPLE_OPS",
}

var AdGroupAdErrorEnum_AdGroupAdError_value = map[string]int32{
	"UNSPECIFIED":                         0,
	"UNKNOWN":                             1,
	"AD_GROUP_AD_LABEL_DOES_NOT_EXIST":    2,
	"AD_GROUP_AD_LABEL_ALREADY_EXISTS":    3,
	"AD_NOT_UNDER_ADGROUP":                4,
	"CANNOT_OPERATE_ON_REMOVED_ADGROUPAD": 5,
	"CANNOT_CREATE_DEPRECATED_ADS":        6,
	"CANNOT_CREATE_TEXT_ADS":              7,
	"EMPTY_FIELD":                         8,
	"RESOURCE_REFERENCED_IN_MULTIPLE_OPS": 9,
}

func (x AdGroupAdErrorEnum_AdGroupAdError) String() string {
	return proto.EnumName(AdGroupAdErrorEnum_AdGroupAdError_name, int32(x))
}

func (AdGroupAdErrorEnum_AdGroupAdError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4f21c135ecafa10d, []int{0, 0}
}

// Container for enum describing possible ad group ad errors.
type AdGroupAdErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupAdErrorEnum) Reset()         { *m = AdGroupAdErrorEnum{} }
func (m *AdGroupAdErrorEnum) String() string { return proto.CompactTextString(m) }
func (*AdGroupAdErrorEnum) ProtoMessage()    {}
func (*AdGroupAdErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f21c135ecafa10d, []int{0}
}

func (m *AdGroupAdErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupAdErrorEnum.Unmarshal(m, b)
}
func (m *AdGroupAdErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupAdErrorEnum.Marshal(b, m, deterministic)
}
func (m *AdGroupAdErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupAdErrorEnum.Merge(m, src)
}
func (m *AdGroupAdErrorEnum) XXX_Size() int {
	return xxx_messageInfo_AdGroupAdErrorEnum.Size(m)
}
func (m *AdGroupAdErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupAdErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupAdErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.AdGroupAdErrorEnum_AdGroupAdError", AdGroupAdErrorEnum_AdGroupAdError_name, AdGroupAdErrorEnum_AdGroupAdError_value)
	proto.RegisterType((*AdGroupAdErrorEnum)(nil), "google.ads.googleads.v2.errors.AdGroupAdErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/ad_group_ad_error.proto", fileDescriptor_4f21c135ecafa10d)
}

var fileDescriptor_4f21c135ecafa10d = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x51, 0x6b, 0xd4, 0x40,
	0x10, 0xc7, 0xbd, 0x54, 0x5b, 0xdd, 0x82, 0x86, 0x28, 0x22, 0xa5, 0x94, 0x72, 0x0a, 0xbe, 0x25,
	0x70, 0x82, 0x0f, 0xf1, 0x69, 0x2e, 0x3b, 0x77, 0x04, 0x73, 0x9b, 0x65, 0x93, 0x9c, 0xad, 0x1c,
	0x2c, 0xd1, 0x3d, 0xc2, 0x41, 0x9b, 0x3d, 0xb2, 0xd7, 0x7e, 0x20, 0x1f, 0x7d, 0xf7, 0x4b, 0x88,
	0x9f, 0xc4, 0x07, 0x3f, 0x83, 0x6c, 0xb6, 0x77, 0x70, 0xa0, 0x7d, 0xda, 0x3f, 0x33, 0xbf, 0xff,
	0xcc, 0xee, 0xce, 0x90, 0xf7, 0x8d, 0xd6, 0xcd, 0xd5, 0x32, 0xaa, 0x95, 0x89, 0x9c, 0xb4, 0xea,
	0x76, 0x14, 0x2d, 0xbb, 0x4e, 0x77, 0x26, 0xaa, 0x95, 0x6c, 0x3a, 0x7d, 0xb3, 0x96, 0xb5, 0x92,
	0x7d, 0x28, 0x5c, 0x77, 0x7a, 0xa3, 0x83, 0x33, 0x07, 0x87, 0xb5, 0x32, 0xe1, 0xce, 0x17, 0xde,
	0x8e, 0x42, 0xe7, 0x3b, 0x39, 0xdd, 0xd6, 0x5d, 0xaf, 0xa2, 0xba, 0x6d, 0xf5, 0xa6, 0xde, 0xac,
	0x74, 0x6b, 0x9c, 0x7b, 0xf8, 0xcb, 0x23, 0x01, 0xa8, 0xa9, 0x2d, 0x0c, 0x0a, 0xad, 0x03, 0xdb,
	0x9b, 0xeb, 0xe1, 0x0f, 0x8f, 0x3c, 0xdd, 0x0f, 0x07, 0xcf, 0xc8, 0x71, 0xc5, 0x0a, 0x8e, 0x49,
	0x3a, 0x49, 0x91, 0xfa, 0x0f, 0x82, 0x63, 0x72, 0x54, 0xb1, 0x8f, 0x2c, 0xff, 0xc4, 0xfc, 0x41,
	0xf0, 0x86, 0x9c, 0x03, 0x95, 0x53, 0x91, 0x57, 0x5c, 0x02, 0x95, 0x19, 0x8c, 0x31, 0x93, 0x34,
	0xc7, 0x42, 0xb2, 0xbc, 0x94, 0x78, 0x91, 0x16, 0xa5, 0xef, 0xfd, 0x9b, 0x82, 0x4c, 0x20, 0xd0,
	0x4b, 0x07, 0x15, 0xfe, 0x41, 0xf0, 0x8a, 0xbc, 0x00, 0xda, 0xfb, 0x2a, 0x46, 0x51, 0x48, 0xa0,
	0xbd, 0xc3, 0x7f, 0x18, 0xbc, 0x25, 0xaf, 0x13, 0x60, 0x36, 0x93, 0x73, 0x14, 0x50, 0xa2, 0xcc,
	0x99, 0x14, 0x38, 0xcb, 0xe7, 0x48, 0xb7, 0x18, 0x50, 0xff, 0x51, 0x70, 0x4e, 0x4e, 0xef, 0xc0,
	0x44, 0xa0, 0xe5, 0x28, 0x72, 0x81, 0x09, 0x94, 0x3d, 0x57, 0xf8, 0x87, 0xc1, 0x09, 0x79, 0xb9,
	0x4f, 0x94, 0x78, 0x51, 0xf6, 0xb9, 0x23, 0xfb, 0x54, 0x9c, 0xf1, 0xf2, 0x52, 0x4e, 0x52, 0xcc,
	0xa8, 0xff, 0xd8, 0xf6, 0x15, 0x58, 0xe4, 0x95, 0x48, 0x50, 0x0a, 0x9c, 0xa0, 0x40, 0x96, 0x20,
	0x95, 0x29, 0x93, 0xb3, 0x2a, 0x2b, 0x53, 0x9e, 0xa1, 0xcc, 0x79, 0xe1, 0x3f, 0x19, 0xff, 0x19,
	0x90, 0xe1, 0x57, 0x7d, 0x1d, 0xde, 0x3f, 0x93, 0xf1, 0xf3, 0xfd, 0xbf, 0xe5, 0x76, 0x14, 0x7c,
	0xf0, 0x99, 0xde, 0xd9, 0x1a, 0x7d, 0x55, 0xb7, 0x4d, 0xa8, 0xbb, 0x26, 0x6a, 0x96, 0x6d, 0x3f,
	0xa8, 0xed, 0x4a, 0xac, 0x57, 0xe6, 0x7f, 0x1b, 0xf2, 0xc1, 0x1d, 0xdf, 0xbc, 0x83, 0x29, 0xc0,
	0x77, 0xef, 0x6c, 0xea, 0x8a, 0x81, 0x32, 0xa1, 0x93, 0x56, 0xcd, 0x47, 0x61, 0xdf, 0xd2, 0xfc,
	0xdc, 0x02, 0x0b, 0x50, 0x66, 0xb1, 0x03, 0x16, 0xf3, 0xd1, 0xc2, 0x01, 0xbf, 0xbd, 0xa1, 0x8b,
	0xc6, 0x31, 0x28, 0x13, 0xc7, 0x3b, 0x24, 0x8e, 0xe7, 0xa3, 0x38, 0x76, 0xd0, 0x97, 0xc3, 0xfe,
	0x76, 0xef, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x09, 0x2b, 0xd4, 0x2c, 0xbe, 0x02, 0x00, 0x00,
}

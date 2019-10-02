// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/ad_group_type.proto

package enums

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

// Enum listing the possible types of an ad group.
type AdGroupTypeEnum_AdGroupType int32

const (
	// The type has not been specified.
	AdGroupTypeEnum_UNSPECIFIED AdGroupTypeEnum_AdGroupType = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	AdGroupTypeEnum_UNKNOWN AdGroupTypeEnum_AdGroupType = 1
	// The default ad group type for Search campaigns.
	AdGroupTypeEnum_SEARCH_STANDARD AdGroupTypeEnum_AdGroupType = 2
	// The default ad group type for Display campaigns.
	AdGroupTypeEnum_DISPLAY_STANDARD AdGroupTypeEnum_AdGroupType = 3
	// The ad group type for Shopping campaigns serving standard product ads.
	AdGroupTypeEnum_SHOPPING_PRODUCT_ADS AdGroupTypeEnum_AdGroupType = 4
	// The default ad group type for Hotel campaigns.
	AdGroupTypeEnum_HOTEL_ADS AdGroupTypeEnum_AdGroupType = 6
	// The type for ad groups in Smart Shopping campaigns.
	AdGroupTypeEnum_SHOPPING_SMART_ADS AdGroupTypeEnum_AdGroupType = 7
	// Short unskippable in-stream video ads.
	AdGroupTypeEnum_VIDEO_BUMPER AdGroupTypeEnum_AdGroupType = 8
	// TrueView (skippable) in-stream video ads.
	AdGroupTypeEnum_VIDEO_TRUE_VIEW_IN_STREAM AdGroupTypeEnum_AdGroupType = 9
	// TrueView in-display video ads.
	AdGroupTypeEnum_VIDEO_TRUE_VIEW_IN_DISPLAY AdGroupTypeEnum_AdGroupType = 10
	// Unskippable in-stream video ads.
	AdGroupTypeEnum_VIDEO_NON_SKIPPABLE_IN_STREAM AdGroupTypeEnum_AdGroupType = 11
	// Outstream video ads.
	AdGroupTypeEnum_VIDEO_OUTSTREAM AdGroupTypeEnum_AdGroupType = 12
	// Ad group type for Dynamic Search Ads ad groups.
	AdGroupTypeEnum_SEARCH_DYNAMIC_ADS AdGroupTypeEnum_AdGroupType = 13
	// The type for ad groups in Shopping Comparison Listing campaigns.
	AdGroupTypeEnum_SHOPPING_COMPARISON_LISTING_ADS AdGroupTypeEnum_AdGroupType = 14
)

var AdGroupTypeEnum_AdGroupType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "SEARCH_STANDARD",
	3:  "DISPLAY_STANDARD",
	4:  "SHOPPING_PRODUCT_ADS",
	6:  "HOTEL_ADS",
	7:  "SHOPPING_SMART_ADS",
	8:  "VIDEO_BUMPER",
	9:  "VIDEO_TRUE_VIEW_IN_STREAM",
	10: "VIDEO_TRUE_VIEW_IN_DISPLAY",
	11: "VIDEO_NON_SKIPPABLE_IN_STREAM",
	12: "VIDEO_OUTSTREAM",
	13: "SEARCH_DYNAMIC_ADS",
	14: "SHOPPING_COMPARISON_LISTING_ADS",
}

var AdGroupTypeEnum_AdGroupType_value = map[string]int32{
	"UNSPECIFIED":                     0,
	"UNKNOWN":                         1,
	"SEARCH_STANDARD":                 2,
	"DISPLAY_STANDARD":                3,
	"SHOPPING_PRODUCT_ADS":            4,
	"HOTEL_ADS":                       6,
	"SHOPPING_SMART_ADS":              7,
	"VIDEO_BUMPER":                    8,
	"VIDEO_TRUE_VIEW_IN_STREAM":       9,
	"VIDEO_TRUE_VIEW_IN_DISPLAY":      10,
	"VIDEO_NON_SKIPPABLE_IN_STREAM":   11,
	"VIDEO_OUTSTREAM":                 12,
	"SEARCH_DYNAMIC_ADS":              13,
	"SHOPPING_COMPARISON_LISTING_ADS": 14,
}

func (x AdGroupTypeEnum_AdGroupType) String() string {
	return proto.EnumName(AdGroupTypeEnum_AdGroupType_name, int32(x))
}

func (AdGroupTypeEnum_AdGroupType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_53ee4e19e3828476, []int{0, 0}
}

// Defines types of an ad group, specific to a particular campaign channel
// type. This type drives validations that restrict which entities can be
// added to the ad group.
type AdGroupTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupTypeEnum) Reset()         { *m = AdGroupTypeEnum{} }
func (m *AdGroupTypeEnum) String() string { return proto.CompactTextString(m) }
func (*AdGroupTypeEnum) ProtoMessage()    {}
func (*AdGroupTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_53ee4e19e3828476, []int{0}
}

func (m *AdGroupTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupTypeEnum.Unmarshal(m, b)
}
func (m *AdGroupTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupTypeEnum.Marshal(b, m, deterministic)
}
func (m *AdGroupTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupTypeEnum.Merge(m, src)
}
func (m *AdGroupTypeEnum) XXX_Size() int {
	return xxx_messageInfo_AdGroupTypeEnum.Size(m)
}
func (m *AdGroupTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.AdGroupTypeEnum_AdGroupType", AdGroupTypeEnum_AdGroupType_name, AdGroupTypeEnum_AdGroupType_value)
	proto.RegisterType((*AdGroupTypeEnum)(nil), "google.ads.googleads.v2.enums.AdGroupTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/ad_group_type.proto", fileDescriptor_53ee4e19e3828476)
}

var fileDescriptor_53ee4e19e3828476 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xc1, 0x6e, 0xd3, 0x4c,
	0x18, 0xfc, 0xe3, 0xfe, 0x6a, 0xe9, 0xa6, 0x25, 0xab, 0xa5, 0x42, 0x50, 0x11, 0x50, 0xca, 0xdd,
	0x16, 0xe6, 0x66, 0x4e, 0x6b, 0x7b, 0x49, 0x56, 0x8d, 0xd7, 0x2b, 0xaf, 0xed, 0xaa, 0x28, 0xd2,
	0xca, 0x60, 0xcb, 0x8a, 0xd4, 0x78, 0xad, 0x38, 0xa9, 0xd4, 0xd7, 0xe1, 0xc8, 0xa3, 0xf0, 0x0e,
	0x1c, 0xb9, 0x70, 0xe3, 0x0d, 0x90, 0xbd, 0x89, 0xc9, 0x01, 0xb8, 0x58, 0x9f, 0x67, 0xe6, 0x1b,
	0xcd, 0x6a, 0x3e, 0xf0, 0xa6, 0x54, 0xaa, 0xbc, 0x2b, 0xac, 0x2c, 0x6f, 0x2c, 0x3d, 0xb6, 0xd3,
	0xbd, 0x6d, 0x15, 0xd5, 0x76, 0xd5, 0x58, 0x59, 0x2e, 0xcb, 0xb5, 0xda, 0xd6, 0x72, 0xf3, 0x50,
	0x17, 0x66, 0xbd, 0x56, 0x1b, 0x85, 0xc6, 0x5a, 0x67, 0x66, 0x79, 0x63, 0xf6, 0x2b, 0xe6, 0xbd,
	0x6d, 0x76, 0x2b, 0x97, 0x2f, 0xf6, 0x8e, 0xf5, 0xd2, 0xca, 0xaa, 0x4a, 0x6d, 0xb2, 0xcd, 0x52,
	0x55, 0x8d, 0x5e, 0xbe, 0xfa, 0x69, 0x80, 0x11, 0xce, 0xa7, 0xad, 0x67, 0xfc, 0x50, 0x17, 0xa4,
	0xda, 0xae, 0xae, 0xbe, 0x19, 0x60, 0x78, 0x80, 0xa1, 0x11, 0x18, 0x26, 0x4c, 0x70, 0xe2, 0xd1,
	0xf7, 0x94, 0xf8, 0xf0, 0x3f, 0x34, 0x04, 0x27, 0x09, 0xbb, 0x66, 0xe1, 0x0d, 0x83, 0x03, 0xf4,
	0x04, 0x8c, 0x04, 0xc1, 0x91, 0x37, 0x93, 0x22, 0xc6, 0xcc, 0xc7, 0x91, 0x0f, 0x0d, 0x74, 0x01,
	0xa0, 0x4f, 0x05, 0x9f, 0xe3, 0xdb, 0xdf, 0xe8, 0x11, 0x7a, 0x06, 0x2e, 0xc4, 0x2c, 0xe4, 0x9c,
	0xb2, 0xa9, 0xe4, 0x51, 0xe8, 0x27, 0x5e, 0x2c, 0xb1, 0x2f, 0xe0, 0xff, 0xe8, 0x1c, 0x9c, 0xce,
	0xc2, 0x98, 0xcc, 0xbb, 0xdf, 0x63, 0xf4, 0x14, 0xa0, 0x5e, 0x28, 0x02, 0x1c, 0x69, 0xd9, 0x09,
	0x82, 0xe0, 0x2c, 0xa5, 0x3e, 0x09, 0xa5, 0x9b, 0x04, 0x9c, 0x44, 0xf0, 0x11, 0x1a, 0x83, 0xe7,
	0x1a, 0x89, 0xa3, 0x84, 0xc8, 0x94, 0x92, 0x1b, 0x49, 0x99, 0x14, 0x71, 0x44, 0x70, 0x00, 0x4f,
	0xd1, 0x4b, 0x70, 0xf9, 0x07, 0x7a, 0x17, 0x0d, 0x02, 0x34, 0x01, 0x63, 0xcd, 0xb3, 0x90, 0x49,
	0x71, 0x4d, 0x39, 0xc7, 0xee, 0x9c, 0x1c, 0x58, 0x0c, 0xdb, 0xf7, 0x69, 0x49, 0x98, 0xc4, 0x3b,
	0xf0, 0xac, 0x0b, 0xa8, 0x1f, 0xed, 0xdf, 0x32, 0x1c, 0x50, 0xaf, 0x0b, 0x78, 0x8e, 0x5e, 0x83,
	0x57, 0x7d, 0x70, 0x2f, 0x0c, 0x38, 0x8e, 0xa8, 0x08, 0x99, 0x9c, 0x53, 0x11, 0xb7, 0x50, 0x2b,
	0x7a, 0xec, 0x7e, 0x1f, 0x80, 0xc9, 0x27, 0xb5, 0x32, 0xff, 0xd9, 0x9b, 0x0b, 0x0f, 0x2a, 0xe0,
	0x6d, 0x57, 0x7c, 0xf0, 0xc1, 0xdd, 0xad, 0x94, 0xea, 0x2e, 0xab, 0x4a, 0x53, 0xad, 0x4b, 0xab,
	0x2c, 0xaa, 0xae, 0xc9, 0xfd, 0xb5, 0xd4, 0xcb, 0xe6, 0x2f, 0xc7, 0xf3, 0xae, 0xfb, 0x7e, 0x36,
	0x8e, 0xa6, 0x18, 0x7f, 0x31, 0xc6, 0x53, 0x6d, 0x85, 0xf3, 0xc6, 0xd4, 0x63, 0x3b, 0xa5, 0xb6,
	0xd9, 0x9e, 0x40, 0xf3, 0x75, 0xcf, 0x2f, 0x70, 0xde, 0x2c, 0x7a, 0x7e, 0x91, 0xda, 0x8b, 0x8e,
	0xff, 0x61, 0x4c, 0x34, 0xe8, 0x38, 0x38, 0x6f, 0x1c, 0xa7, 0x57, 0x38, 0x4e, 0x6a, 0x3b, 0x4e,
	0xa7, 0xf9, 0x78, 0xdc, 0x05, 0x7b, 0xfb, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x12, 0xce, 0x5a,
	0xd4, 0x02, 0x00, 0x00,
}

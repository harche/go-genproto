// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/keyword_plan_campaign_error.proto

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

// Enum describing possible errors from applying a keyword plan campaign.
type KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError int32

const (
	// Enum unspecified.
	KeywordPlanCampaignErrorEnum_UNSPECIFIED KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 0
	// The received error code is not known in this version.
	KeywordPlanCampaignErrorEnum_UNKNOWN KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 1
	// A keyword plan campaign name is missing, empty, longer than allowed limit
	// or contains invalid chars.
	KeywordPlanCampaignErrorEnum_INVALID_NAME KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 2
	// A keyword plan campaign contains one or more untargetable languages.
	KeywordPlanCampaignErrorEnum_INVALID_LANGUAGES KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 3
	// A keyword plan campaign contains one or more invalid geo targets.
	KeywordPlanCampaignErrorEnum_INVALID_GEOS KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 4
	// The keyword plan campaign name is duplicate to an existing keyword plan
	// campaign name or other keyword plan campaign name in the request.
	KeywordPlanCampaignErrorEnum_DUPLICATE_NAME KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 5
	// The number of geo targets in the keyword plan campaign exceeds limits.
	KeywordPlanCampaignErrorEnum_MAX_GEOS_EXCEEDED KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError = 6
)

var KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "INVALID_NAME",
	3: "INVALID_LANGUAGES",
	4: "INVALID_GEOS",
	5: "DUPLICATE_NAME",
	6: "MAX_GEOS_EXCEEDED",
}

var KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError_value = map[string]int32{
	"UNSPECIFIED":       0,
	"UNKNOWN":           1,
	"INVALID_NAME":      2,
	"INVALID_LANGUAGES": 3,
	"INVALID_GEOS":      4,
	"DUPLICATE_NAME":    5,
	"MAX_GEOS_EXCEEDED": 6,
}

func (x KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError) String() string {
	return proto.EnumName(KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError_name, int32(x))
}

func (KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_adc1819a9134bc15, []int{0, 0}
}

// Container for enum describing possible errors from applying a keyword plan
// campaign.
type KeywordPlanCampaignErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeywordPlanCampaignErrorEnum) Reset()         { *m = KeywordPlanCampaignErrorEnum{} }
func (m *KeywordPlanCampaignErrorEnum) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanCampaignErrorEnum) ProtoMessage()    {}
func (*KeywordPlanCampaignErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_adc1819a9134bc15, []int{0}
}

func (m *KeywordPlanCampaignErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanCampaignErrorEnum.Unmarshal(m, b)
}
func (m *KeywordPlanCampaignErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanCampaignErrorEnum.Marshal(b, m, deterministic)
}
func (m *KeywordPlanCampaignErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanCampaignErrorEnum.Merge(m, src)
}
func (m *KeywordPlanCampaignErrorEnum) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanCampaignErrorEnum.Size(m)
}
func (m *KeywordPlanCampaignErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanCampaignErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanCampaignErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError", KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError_name, KeywordPlanCampaignErrorEnum_KeywordPlanCampaignError_value)
	proto.RegisterType((*KeywordPlanCampaignErrorEnum)(nil), "google.ads.googleads.v2.errors.KeywordPlanCampaignErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/keyword_plan_campaign_error.proto", fileDescriptor_adc1819a9134bc15)
}

var fileDescriptor_adc1819a9134bc15 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x6a, 0xa3, 0x40,
	0x1c, 0xc6, 0x57, 0xb3, 0x9b, 0x85, 0xc9, 0xb2, 0xeb, 0x0a, 0x85, 0x52, 0xd2, 0x1c, 0x7c, 0x80,
	0x11, 0xec, 0x6d, 0x7a, 0xe9, 0x44, 0xa7, 0x22, 0x49, 0x8c, 0x90, 0x6a, 0x43, 0x11, 0x64, 0x1a,
	0x45, 0xa4, 0x66, 0x46, 0x9c, 0x34, 0xa5, 0x2f, 0xd3, 0x43, 0x8f, 0x7d, 0x89, 0xde, 0xfb, 0x28,
	0x7d, 0x82, 0x1e, 0x8b, 0x4e, 0x0c, 0xbd, 0xa4, 0x27, 0x3f, 0xfe, 0xfe, 0xbe, 0xef, 0xd3, 0xff,
	0x1f, 0x5c, 0xe4, 0x9c, 0xe7, 0x65, 0x66, 0xd2, 0x54, 0x98, 0x52, 0x36, 0x6a, 0x6b, 0x99, 0x59,
	0x5d, 0xf3, 0x5a, 0x98, 0x77, 0xd9, 0xe3, 0x03, 0xaf, 0xd3, 0xa4, 0x2a, 0x29, 0x4b, 0x56, 0x74,
	0x5d, 0xd1, 0x22, 0x67, 0x49, 0xfb, 0x12, 0x56, 0x35, 0xdf, 0x70, 0x7d, 0x24, 0x6d, 0x90, 0xa6,
	0x02, 0xee, 0x13, 0xe0, 0xd6, 0x82, 0x32, 0xe1, 0x64, 0xd8, 0x35, 0x54, 0x85, 0x49, 0x19, 0xe3,
	0x1b, 0xba, 0x29, 0x38, 0x13, 0xd2, 0x6d, 0xbc, 0x2a, 0x60, 0x38, 0x91, 0x1d, 0x41, 0x49, 0x99,
	0xbd, 0x6b, 0x20, 0x8d, 0x97, 0xb0, 0xfb, 0xb5, 0xf1, 0xa4, 0x80, 0xe3, 0x43, 0x80, 0xfe, 0x0f,
	0x0c, 0x42, 0x7f, 0x11, 0x10, 0xdb, 0xbb, 0xf4, 0x88, 0xa3, 0xfd, 0xd0, 0x07, 0xe0, 0x77, 0xe8,
	0x4f, 0xfc, 0xf9, 0xb5, 0xaf, 0x29, 0xba, 0x06, 0xfe, 0x78, 0x7e, 0x84, 0xa7, 0x9e, 0x93, 0xf8,
	0x78, 0x46, 0x34, 0x55, 0x3f, 0x02, 0xff, 0xbb, 0xc9, 0x14, 0xfb, 0x6e, 0x88, 0x5d, 0xb2, 0xd0,
	0x7a, 0x5f, 0x41, 0x97, 0xcc, 0x17, 0xda, 0x4f, 0x5d, 0x07, 0x7f, 0x9d, 0x30, 0x98, 0x7a, 0x36,
	0xbe, 0x22, 0xd2, 0xfc, 0xab, 0x31, 0xcf, 0xf0, 0xb2, 0x25, 0x12, 0xb2, 0xb4, 0x09, 0x71, 0x88,
	0xa3, 0xf5, 0xc7, 0x1f, 0x0a, 0x30, 0x56, 0x7c, 0x0d, 0xbf, 0x5f, 0xc3, 0xf8, 0xf4, 0xd0, 0x4f,
	0x04, 0xcd, 0x1e, 0x02, 0xe5, 0xc6, 0xd9, 0x05, 0xe4, 0xbc, 0xa4, 0x2c, 0x87, 0xbc, 0xce, 0xcd,
	0x3c, 0x63, 0xed, 0x96, 0xba, 0xcb, 0x54, 0x85, 0x38, 0x74, 0xa8, 0x73, 0xf9, 0x78, 0x56, 0x7b,
	0x2e, 0xc6, 0x2f, 0xea, 0xc8, 0x95, 0x61, 0x38, 0x15, 0x50, 0xca, 0x46, 0x45, 0x16, 0x6c, 0x2b,
	0xc5, 0x5b, 0x07, 0xc4, 0x38, 0x15, 0xf1, 0x1e, 0x88, 0x23, 0x2b, 0x96, 0xc0, 0xbb, 0x6a, 0xc8,
	0x29, 0x42, 0x38, 0x15, 0x08, 0xed, 0x11, 0x84, 0x22, 0x0b, 0x21, 0x09, 0xdd, 0xf6, 0xdb, 0xaf,
	0x3b, 0xfb, 0x0c, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x1d, 0x95, 0x58, 0x45, 0x02, 0x00, 0x00,
}

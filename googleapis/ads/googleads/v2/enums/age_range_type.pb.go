// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/age_range_type.proto

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

// The type of demographic age ranges (e.g. between 18 and 24 years old).
type AgeRangeTypeEnum_AgeRangeType int32

const (
	// Not specified.
	AgeRangeTypeEnum_UNSPECIFIED AgeRangeTypeEnum_AgeRangeType = 0
	// Used for return value only. Represents value unknown in this version.
	AgeRangeTypeEnum_UNKNOWN AgeRangeTypeEnum_AgeRangeType = 1
	// Between 18 and 24 years old.
	AgeRangeTypeEnum_AGE_RANGE_18_24 AgeRangeTypeEnum_AgeRangeType = 503001
	// Between 25 and 34 years old.
	AgeRangeTypeEnum_AGE_RANGE_25_34 AgeRangeTypeEnum_AgeRangeType = 503002
	// Between 35 and 44 years old.
	AgeRangeTypeEnum_AGE_RANGE_35_44 AgeRangeTypeEnum_AgeRangeType = 503003
	// Between 45 and 54 years old.
	AgeRangeTypeEnum_AGE_RANGE_45_54 AgeRangeTypeEnum_AgeRangeType = 503004
	// Between 55 and 64 years old.
	AgeRangeTypeEnum_AGE_RANGE_55_64 AgeRangeTypeEnum_AgeRangeType = 503005
	// 65 years old and beyond.
	AgeRangeTypeEnum_AGE_RANGE_65_UP AgeRangeTypeEnum_AgeRangeType = 503006
	// Undetermined age range.
	AgeRangeTypeEnum_AGE_RANGE_UNDETERMINED AgeRangeTypeEnum_AgeRangeType = 503999
)

var AgeRangeTypeEnum_AgeRangeType_name = map[int32]string{
	0:      "UNSPECIFIED",
	1:      "UNKNOWN",
	503001: "AGE_RANGE_18_24",
	503002: "AGE_RANGE_25_34",
	503003: "AGE_RANGE_35_44",
	503004: "AGE_RANGE_45_54",
	503005: "AGE_RANGE_55_64",
	503006: "AGE_RANGE_65_UP",
	503999: "AGE_RANGE_UNDETERMINED",
}

var AgeRangeTypeEnum_AgeRangeType_value = map[string]int32{
	"UNSPECIFIED":            0,
	"UNKNOWN":                1,
	"AGE_RANGE_18_24":        503001,
	"AGE_RANGE_25_34":        503002,
	"AGE_RANGE_35_44":        503003,
	"AGE_RANGE_45_54":        503004,
	"AGE_RANGE_55_64":        503005,
	"AGE_RANGE_65_UP":        503006,
	"AGE_RANGE_UNDETERMINED": 503999,
}

func (x AgeRangeTypeEnum_AgeRangeType) String() string {
	return proto.EnumName(AgeRangeTypeEnum_AgeRangeType_name, int32(x))
}

func (AgeRangeTypeEnum_AgeRangeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6862e2c9d484ba9a, []int{0, 0}
}

// Container for enum describing the type of demographic age ranges.
type AgeRangeTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgeRangeTypeEnum) Reset()         { *m = AgeRangeTypeEnum{} }
func (m *AgeRangeTypeEnum) String() string { return proto.CompactTextString(m) }
func (*AgeRangeTypeEnum) ProtoMessage()    {}
func (*AgeRangeTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_6862e2c9d484ba9a, []int{0}
}

func (m *AgeRangeTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgeRangeTypeEnum.Unmarshal(m, b)
}
func (m *AgeRangeTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgeRangeTypeEnum.Marshal(b, m, deterministic)
}
func (m *AgeRangeTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgeRangeTypeEnum.Merge(m, src)
}
func (m *AgeRangeTypeEnum) XXX_Size() int {
	return xxx_messageInfo_AgeRangeTypeEnum.Size(m)
}
func (m *AgeRangeTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AgeRangeTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AgeRangeTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.AgeRangeTypeEnum_AgeRangeType", AgeRangeTypeEnum_AgeRangeType_name, AgeRangeTypeEnum_AgeRangeType_value)
	proto.RegisterType((*AgeRangeTypeEnum)(nil), "google.ads.googleads.v2.enums.AgeRangeTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/age_range_type.proto", fileDescriptor_6862e2c9d484ba9a)
}

var fileDescriptor_6862e2c9d484ba9a = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x4d, 0x05, 0x85, 0xa9, 0xd0, 0x31, 0xa0, 0x07, 0x69, 0x0b, 0xed, 0x07, 0x98, 0xe0,
	0x34, 0x29, 0x32, 0x9e, 0xa6, 0x76, 0x0c, 0x45, 0x8c, 0xa1, 0x36, 0x15, 0x24, 0x30, 0x44, 0x33,
	0x0c, 0x85, 0x76, 0x26, 0x74, 0xd2, 0x42, 0xbf, 0x8e, 0x47, 0xbf, 0x84, 0x77, 0xcf, 0x5e, 0xea,
	0x3f, 0xd8, 0xcb, 0xc2, 0xee, 0xa7, 0x58, 0x92, 0x6c, 0xbb, 0xdd, 0xc0, 0xee, 0x25, 0x3c, 0xfc,
	0xde, 0xf7, 0x79, 0xc8, 0x3c, 0x2f, 0xc0, 0x52, 0x6b, 0xb9, 0x14, 0x4e, 0x92, 0x1a, 0xa7, 0x92,
	0x85, 0xda, 0x62, 0x47, 0xa8, 0xcd, 0xca, 0x38, 0x89, 0x14, 0x7c, 0x9d, 0x28, 0x29, 0x78, 0xbe,
	0xcb, 0x04, 0xca, 0xd6, 0x3a, 0xd7, 0x76, 0xa7, 0x5a, 0x44, 0x49, 0x6a, 0xd0, 0xd1, 0x83, 0xb6,
	0x18, 0x95, 0x9e, 0x17, 0xed, 0x43, 0x64, 0xb6, 0x70, 0x12, 0xa5, 0x74, 0x9e, 0xe4, 0x0b, 0xad,
	0x4c, 0x65, 0xee, 0x5f, 0x58, 0x00, 0x52, 0x29, 0xa6, 0x45, 0xe8, 0x6c, 0x97, 0x09, 0xa6, 0x36,
	0xab, 0xfe, 0x2f, 0x0b, 0x3c, 0x39, 0x85, 0x76, 0x0b, 0x34, 0xa3, 0xe0, 0x63, 0xc8, 0xde, 0x4c,
	0xde, 0x4e, 0xd8, 0x18, 0x3e, 0xb0, 0x9b, 0xe0, 0x71, 0x14, 0xbc, 0x0b, 0x3e, 0x7c, 0x0a, 0xa0,
	0x65, 0x3f, 0x03, 0x2d, 0xea, 0x33, 0x3e, 0xa5, 0x81, 0xcf, 0xf8, 0xcb, 0x57, 0x1c, 0xbb, 0x70,
	0xbf, 0xef, 0xde, 0xc6, 0xd8, 0xe3, 0x03, 0x17, 0xfe, 0xae, 0xe3, 0x81, 0xc7, 0x5d, 0x17, 0xfe,
	0xa9, 0x63, 0xd7, 0xe3, 0x9e, 0x0b, 0xff, 0xd6, 0xb1, 0xe7, 0xf1, 0xa1, 0x0b, 0xff, 0xd5, 0xf1,
	0xd0, 0xe3, 0x51, 0x08, 0xff, 0xef, 0xbb, 0x76, 0x1b, 0x3c, 0xbf, 0xc1, 0x51, 0x30, 0x66, 0x33,
	0x36, 0x7d, 0x3f, 0x09, 0xd8, 0x18, 0xfe, 0x38, 0xeb, 0x8e, 0xce, 0x2d, 0xd0, 0xfb, 0xaa, 0x57,
	0xe8, 0xde, 0xbe, 0x46, 0x4f, 0x4f, 0x5f, 0x1e, 0x16, 0x25, 0x85, 0xd6, 0xe7, 0xd1, 0xb5, 0x47,
	0xea, 0x65, 0xa2, 0x24, 0xd2, 0x6b, 0xe9, 0x48, 0xa1, 0xca, 0x0a, 0x0f, 0x77, 0xca, 0x16, 0xe6,
	0x8e, 0xb3, 0xbd, 0x2e, 0xbf, 0xdf, 0x1a, 0x0f, 0x7d, 0x4a, 0xbf, 0x37, 0x3a, 0x7e, 0x15, 0x45,
	0x53, 0x83, 0x2a, 0x59, 0xa8, 0x39, 0x46, 0x45, 0xf5, 0xe6, 0xe7, 0x61, 0x1e, 0xd3, 0xd4, 0xc4,
	0xc7, 0x79, 0x3c, 0xc7, 0x71, 0x39, 0xbf, 0x6c, 0xf4, 0x2a, 0x48, 0x08, 0x4d, 0x0d, 0x21, 0xc7,
	0x0d, 0x42, 0xe6, 0x98, 0x90, 0x72, 0xe7, 0xcb, 0xa3, 0xf2, 0xc7, 0x06, 0x57, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x3d, 0x10, 0x47, 0x8f, 0x4e, 0x02, 0x00, 0x00,
}

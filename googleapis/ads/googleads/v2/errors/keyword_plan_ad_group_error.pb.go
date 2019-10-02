// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/keyword_plan_ad_group_error.proto

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

// Enum describing possible errors from applying a keyword plan ad group.
type KeywordPlanAdGroupErrorEnum_KeywordPlanAdGroupError int32

const (
	// Enum unspecified.
	KeywordPlanAdGroupErrorEnum_UNSPECIFIED KeywordPlanAdGroupErrorEnum_KeywordPlanAdGroupError = 0
	// The received error code is not known in this version.
	KeywordPlanAdGroupErrorEnum_UNKNOWN KeywordPlanAdGroupErrorEnum_KeywordPlanAdGroupError = 1
	// The keyword plan ad group name is missing, empty, longer than allowed
	// limit or contains invalid chars.
	KeywordPlanAdGroupErrorEnum_INVALID_NAME KeywordPlanAdGroupErrorEnum_KeywordPlanAdGroupError = 2
	// The keyword plan ad group name is duplicate to an existing keyword plan
	// AdGroup name or other keyword plan AdGroup name in the request.
	KeywordPlanAdGroupErrorEnum_DUPLICATE_NAME KeywordPlanAdGroupErrorEnum_KeywordPlanAdGroupError = 3
)

var KeywordPlanAdGroupErrorEnum_KeywordPlanAdGroupError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "INVALID_NAME",
	3: "DUPLICATE_NAME",
}

var KeywordPlanAdGroupErrorEnum_KeywordPlanAdGroupError_value = map[string]int32{
	"UNSPECIFIED":    0,
	"UNKNOWN":        1,
	"INVALID_NAME":   2,
	"DUPLICATE_NAME": 3,
}

func (x KeywordPlanAdGroupErrorEnum_KeywordPlanAdGroupError) String() string {
	return proto.EnumName(KeywordPlanAdGroupErrorEnum_KeywordPlanAdGroupError_name, int32(x))
}

func (KeywordPlanAdGroupErrorEnum_KeywordPlanAdGroupError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_89f6d075052bb2c4, []int{0, 0}
}

// Container for enum describing possible errors from applying a keyword plan
// ad group.
type KeywordPlanAdGroupErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeywordPlanAdGroupErrorEnum) Reset()         { *m = KeywordPlanAdGroupErrorEnum{} }
func (m *KeywordPlanAdGroupErrorEnum) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanAdGroupErrorEnum) ProtoMessage()    {}
func (*KeywordPlanAdGroupErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_89f6d075052bb2c4, []int{0}
}

func (m *KeywordPlanAdGroupErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanAdGroupErrorEnum.Unmarshal(m, b)
}
func (m *KeywordPlanAdGroupErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanAdGroupErrorEnum.Marshal(b, m, deterministic)
}
func (m *KeywordPlanAdGroupErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanAdGroupErrorEnum.Merge(m, src)
}
func (m *KeywordPlanAdGroupErrorEnum) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanAdGroupErrorEnum.Size(m)
}
func (m *KeywordPlanAdGroupErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanAdGroupErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanAdGroupErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.KeywordPlanAdGroupErrorEnum_KeywordPlanAdGroupError", KeywordPlanAdGroupErrorEnum_KeywordPlanAdGroupError_name, KeywordPlanAdGroupErrorEnum_KeywordPlanAdGroupError_value)
	proto.RegisterType((*KeywordPlanAdGroupErrorEnum)(nil), "google.ads.googleads.v2.errors.KeywordPlanAdGroupErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/keyword_plan_ad_group_error.proto", fileDescriptor_89f6d075052bb2c4)
}

var fileDescriptor_89f6d075052bb2c4 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcf, 0x4a, 0xc3, 0x30,
	0x18, 0x77, 0x1d, 0x28, 0x64, 0xa2, 0xa5, 0x17, 0x41, 0xc7, 0x0e, 0x7d, 0x80, 0x14, 0xea, 0x2d,
	0x5e, 0xcc, 0xd6, 0x3a, 0xca, 0x66, 0x2c, 0xe8, 0x2a, 0x48, 0xa5, 0x44, 0x53, 0xc2, 0xb0, 0x4b,
	0x4a, 0xb2, 0x4d, 0x04, 0x9f, 0xc6, 0xa3, 0x8f, 0xe2, 0xa3, 0xf8, 0x02, 0x5e, 0xa5, 0xfd, 0xdc,
	0x6e, 0xf3, 0x94, 0x1f, 0xf9, 0x7e, 0x7f, 0xbe, 0xdf, 0x87, 0x2e, 0xa5, 0xd6, 0xb2, 0x2a, 0x03,
	0x2e, 0x6c, 0x00, 0xb0, 0x41, 0xeb, 0x30, 0x28, 0x8d, 0xd1, 0xc6, 0x06, 0x2f, 0xe5, 0xdb, 0xab,
	0x36, 0xa2, 0xa8, 0x2b, 0xae, 0x0a, 0x2e, 0x0a, 0x69, 0xf4, 0xaa, 0x2e, 0xda, 0x21, 0xae, 0x8d,
	0x5e, 0x6a, 0x6f, 0x00, 0x32, 0xcc, 0x85, 0xc5, 0x5b, 0x07, 0xbc, 0x0e, 0x31, 0x38, 0x9c, 0xf6,
	0x37, 0x09, 0xf5, 0x3c, 0xe0, 0x4a, 0xe9, 0x25, 0x5f, 0xce, 0xb5, 0xb2, 0xa0, 0xf6, 0xdf, 0xd1,
	0xd9, 0x04, 0x22, 0xd2, 0x8a, 0x2b, 0x2a, 0xc6, 0x8d, 0x7f, 0xdc, 0x28, 0x63, 0xb5, 0x5a, 0xf8,
	0x8f, 0xe8, 0x64, 0xc7, 0xd8, 0x3b, 0x46, 0xbd, 0x19, 0xbb, 0x4d, 0xe3, 0x51, 0x72, 0x95, 0xc4,
	0x91, 0xbb, 0xe7, 0xf5, 0xd0, 0xc1, 0x8c, 0x4d, 0xd8, 0xcd, 0x3d, 0x73, 0x3b, 0x9e, 0x8b, 0x0e,
	0x13, 0x96, 0xd1, 0x69, 0x12, 0x15, 0x8c, 0x5e, 0xc7, 0xae, 0xe3, 0x79, 0xe8, 0x28, 0x9a, 0xa5,
	0xd3, 0x64, 0x44, 0xef, 0x62, 0xf8, 0xeb, 0x0e, 0x7f, 0x3a, 0xc8, 0x7f, 0xd6, 0x0b, 0xfc, 0x7f,
	0x85, 0x61, 0x7f, 0xc7, 0x0e, 0x69, 0x53, 0x21, 0xed, 0x3c, 0x44, 0x7f, 0x7a, 0xa9, 0x2b, 0xae,
	0x24, 0xd6, 0x46, 0x06, 0xb2, 0x54, 0x6d, 0xc1, 0xcd, 0x51, 0xeb, 0xb9, 0xdd, 0x75, 0xe3, 0x0b,
	0x78, 0x3e, 0x9c, 0xee, 0x98, 0xd2, 0x4f, 0x67, 0x30, 0x06, 0x33, 0x2a, 0x2c, 0x06, 0xd8, 0xa0,
	0x2c, 0xc4, 0x6d, 0xa4, 0xfd, 0xda, 0x10, 0x72, 0x2a, 0x6c, 0xbe, 0x25, 0xe4, 0x59, 0x98, 0x03,
	0xe1, 0xdb, 0xf1, 0xe1, 0x97, 0x10, 0x2a, 0x2c, 0x21, 0x5b, 0x0a, 0x21, 0x59, 0x48, 0x08, 0x90,
	0x9e, 0xf6, 0xdb, 0xed, 0xce, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x01, 0x8f, 0x7d, 0x00,
	0x02, 0x00, 0x00,
}

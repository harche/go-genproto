// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/matching_function_context_type.proto

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

// Possible context types for an operand in a matching function.
type MatchingFunctionContextTypeEnum_MatchingFunctionContextType int32

const (
	// Not specified.
	MatchingFunctionContextTypeEnum_UNSPECIFIED MatchingFunctionContextTypeEnum_MatchingFunctionContextType = 0
	// Used for return value only. Represents value unknown in this version.
	MatchingFunctionContextTypeEnum_UNKNOWN MatchingFunctionContextTypeEnum_MatchingFunctionContextType = 1
	// Feed item id in the request context.
	MatchingFunctionContextTypeEnum_FEED_ITEM_ID MatchingFunctionContextTypeEnum_MatchingFunctionContextType = 2
	// The device being used (possible values are 'Desktop' or 'Mobile').
	MatchingFunctionContextTypeEnum_DEVICE_NAME MatchingFunctionContextTypeEnum_MatchingFunctionContextType = 3
)

var MatchingFunctionContextTypeEnum_MatchingFunctionContextType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "FEED_ITEM_ID",
	3: "DEVICE_NAME",
}

var MatchingFunctionContextTypeEnum_MatchingFunctionContextType_value = map[string]int32{
	"UNSPECIFIED":  0,
	"UNKNOWN":      1,
	"FEED_ITEM_ID": 2,
	"DEVICE_NAME":  3,
}

func (x MatchingFunctionContextTypeEnum_MatchingFunctionContextType) String() string {
	return proto.EnumName(MatchingFunctionContextTypeEnum_MatchingFunctionContextType_name, int32(x))
}

func (MatchingFunctionContextTypeEnum_MatchingFunctionContextType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d2c8b9633b2c349b, []int{0, 0}
}

// Container for context types for an operand in a matching function.
type MatchingFunctionContextTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MatchingFunctionContextTypeEnum) Reset()         { *m = MatchingFunctionContextTypeEnum{} }
func (m *MatchingFunctionContextTypeEnum) String() string { return proto.CompactTextString(m) }
func (*MatchingFunctionContextTypeEnum) ProtoMessage()    {}
func (*MatchingFunctionContextTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2c8b9633b2c349b, []int{0}
}

func (m *MatchingFunctionContextTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchingFunctionContextTypeEnum.Unmarshal(m, b)
}
func (m *MatchingFunctionContextTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchingFunctionContextTypeEnum.Marshal(b, m, deterministic)
}
func (m *MatchingFunctionContextTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchingFunctionContextTypeEnum.Merge(m, src)
}
func (m *MatchingFunctionContextTypeEnum) XXX_Size() int {
	return xxx_messageInfo_MatchingFunctionContextTypeEnum.Size(m)
}
func (m *MatchingFunctionContextTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchingFunctionContextTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_MatchingFunctionContextTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.MatchingFunctionContextTypeEnum_MatchingFunctionContextType", MatchingFunctionContextTypeEnum_MatchingFunctionContextType_name, MatchingFunctionContextTypeEnum_MatchingFunctionContextType_value)
	proto.RegisterType((*MatchingFunctionContextTypeEnum)(nil), "google.ads.googleads.v2.enums.MatchingFunctionContextTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/matching_function_context_type.proto", fileDescriptor_d2c8b9633b2c349b)
}

var fileDescriptor_d2c8b9633b2c349b = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0xb5, 0x1d, 0x28, 0x64, 0x82, 0xa5, 0x47, 0x75, 0xe8, 0xf6, 0x03, 0x52, 0xa8, 0xb7, 0x78,
	0xea, 0xd6, 0x6c, 0x14, 0x59, 0x1d, 0xb8, 0x55, 0x90, 0x62, 0xa9, 0x6d, 0x8c, 0x85, 0x35, 0x29,
	0x4b, 0x36, 0xdc, 0xd1, 0xbf, 0xe2, 0xd1, 0x9f, 0xe2, 0x4f, 0xf1, 0xee, 0x5d, 0x92, 0x6c, 0xbb,
	0xd9, 0x4b, 0x78, 0xe4, 0x7b, 0xdf, 0x7b, 0xdf, 0x7b, 0x60, 0x48, 0x39, 0xa7, 0x4b, 0xe2, 0xe5,
	0xa5, 0xf0, 0x0c, 0x54, 0x68, 0xe3, 0x7b, 0x84, 0xad, 0x6b, 0xe1, 0xd5, 0xb9, 0x2c, 0xde, 0x2a,
	0x46, 0xb3, 0xd7, 0x35, 0x2b, 0x64, 0xc5, 0x59, 0x56, 0x70, 0x26, 0xc9, 0xbb, 0xcc, 0xe4, 0xb6,
	0x21, 0xb0, 0x59, 0x71, 0xc9, 0xdd, 0x9e, 0x59, 0x84, 0x79, 0x29, 0xe0, 0x41, 0x03, 0x6e, 0x7c,
	0xa8, 0x35, 0xce, 0x2f, 0xf7, 0x16, 0x4d, 0xe5, 0xe5, 0x8c, 0x71, 0x99, 0x2b, 0x21, 0x61, 0x96,
	0x07, 0x1f, 0x16, 0xb8, 0x9a, 0xee, 0x5c, 0xc6, 0x3b, 0x93, 0x91, 0xf1, 0x98, 0x6f, 0x1b, 0x82,
	0xd9, 0xba, 0x1e, 0x3c, 0x83, 0x8b, 0x16, 0x8a, 0x7b, 0x06, 0xba, 0x8b, 0xf8, 0x61, 0x86, 0x47,
	0xd1, 0x38, 0xc2, 0xa1, 0x73, 0xe4, 0x76, 0xc1, 0xc9, 0x22, 0xbe, 0x8b, 0xef, 0x1f, 0x63, 0xc7,
	0x72, 0x1d, 0x70, 0x3a, 0xc6, 0x38, 0xcc, 0xa2, 0x39, 0x9e, 0x66, 0x51, 0xe8, 0xd8, 0x8a, 0x1f,
	0xe2, 0x24, 0x1a, 0xe1, 0x2c, 0x0e, 0xa6, 0xd8, 0xe9, 0x0c, 0x7f, 0x2d, 0xd0, 0x2f, 0x78, 0x0d,
	0x5b, 0x73, 0x0c, 0xaf, 0x5b, 0x6e, 0x98, 0xa9, 0x2c, 0x33, 0xeb, 0x69, 0x57, 0x27, 0xa4, 0x7c,
	0x99, 0x33, 0x0a, 0xf9, 0x8a, 0x7a, 0x94, 0x30, 0x9d, 0x74, 0x5f, 0x6f, 0x53, 0x89, 0x7f, 0xda,
	0xbe, 0xd5, 0xef, 0xa7, 0xdd, 0x99, 0x04, 0xc1, 0x97, 0xdd, 0x9b, 0x18, 0xa9, 0xa0, 0x14, 0xd0,
	0x40, 0x85, 0x12, 0x1f, 0xaa, 0x4a, 0xc4, 0xf7, 0x7e, 0x9e, 0x06, 0xa5, 0x48, 0x0f, 0xf3, 0x34,
	0xf1, 0x53, 0x3d, 0xff, 0xb1, 0xfb, 0xe6, 0x13, 0xa1, 0xa0, 0x14, 0x08, 0x1d, 0x18, 0x08, 0x25,
	0x3e, 0x42, 0x9a, 0xf3, 0x72, 0xac, 0x0f, 0xbb, 0xf9, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x59,
	0xfc, 0x62, 0x05, 0x02, 0x00, 0x00,
}

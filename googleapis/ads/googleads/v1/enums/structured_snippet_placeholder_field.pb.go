// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/structured_snippet_placeholder_field.proto

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

// Possible values for Structured Snippet placeholder fields.
type StructuredSnippetPlaceholderFieldEnum_StructuredSnippetPlaceholderField int32

const (
	// Not specified.
	StructuredSnippetPlaceholderFieldEnum_UNSPECIFIED StructuredSnippetPlaceholderFieldEnum_StructuredSnippetPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	StructuredSnippetPlaceholderFieldEnum_UNKNOWN StructuredSnippetPlaceholderFieldEnum_StructuredSnippetPlaceholderField = 1
	// Data Type: STRING. The category of snippet of your products/services.
	// Must match one of the predefined structured snippets headers exactly.
	// See
	// https://developers.google.com/adwords/api
	// /docs/appendix/structured-snippet-headers
	StructuredSnippetPlaceholderFieldEnum_HEADER StructuredSnippetPlaceholderFieldEnum_StructuredSnippetPlaceholderField = 2
	// Data Type: STRING_LIST. Text values that describe your products/services.
	// All text must be family safe. Special or non-ASCII characters are not
	// permitted. A snippet can be at most 25 characters.
	StructuredSnippetPlaceholderFieldEnum_SNIPPETS StructuredSnippetPlaceholderFieldEnum_StructuredSnippetPlaceholderField = 3
)

var StructuredSnippetPlaceholderFieldEnum_StructuredSnippetPlaceholderField_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "HEADER",
	3: "SNIPPETS",
}

var StructuredSnippetPlaceholderFieldEnum_StructuredSnippetPlaceholderField_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"HEADER":      2,
	"SNIPPETS":    3,
}

func (x StructuredSnippetPlaceholderFieldEnum_StructuredSnippetPlaceholderField) String() string {
	return proto.EnumName(StructuredSnippetPlaceholderFieldEnum_StructuredSnippetPlaceholderField_name, int32(x))
}

func (StructuredSnippetPlaceholderFieldEnum_StructuredSnippetPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b931266ec688ae7c, []int{0, 0}
}

// Values for Structured Snippet placeholder fields.
type StructuredSnippetPlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StructuredSnippetPlaceholderFieldEnum) Reset()         { *m = StructuredSnippetPlaceholderFieldEnum{} }
func (m *StructuredSnippetPlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*StructuredSnippetPlaceholderFieldEnum) ProtoMessage()    {}
func (*StructuredSnippetPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_b931266ec688ae7c, []int{0}
}

func (m *StructuredSnippetPlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StructuredSnippetPlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *StructuredSnippetPlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StructuredSnippetPlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (m *StructuredSnippetPlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StructuredSnippetPlaceholderFieldEnum.Merge(m, src)
}
func (m *StructuredSnippetPlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_StructuredSnippetPlaceholderFieldEnum.Size(m)
}
func (m *StructuredSnippetPlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_StructuredSnippetPlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_StructuredSnippetPlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.StructuredSnippetPlaceholderFieldEnum_StructuredSnippetPlaceholderField", StructuredSnippetPlaceholderFieldEnum_StructuredSnippetPlaceholderField_name, StructuredSnippetPlaceholderFieldEnum_StructuredSnippetPlaceholderField_value)
	proto.RegisterType((*StructuredSnippetPlaceholderFieldEnum)(nil), "google.ads.googleads.v1.enums.StructuredSnippetPlaceholderFieldEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/structured_snippet_placeholder_field.proto", fileDescriptor_b931266ec688ae7c)
}

var fileDescriptor_b931266ec688ae7c = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x50, 0xdf, 0x4a, 0xfb, 0x30,
	0x18, 0xfd, 0xb5, 0x83, 0xfd, 0x24, 0x13, 0x2c, 0xbd, 0x14, 0x77, 0xb1, 0x81, 0x5e, 0xa6, 0x14,
	0xef, 0xe2, 0x55, 0xe7, 0xba, 0x3f, 0x08, 0xb5, 0x58, 0x37, 0x41, 0x0b, 0x23, 0x2e, 0x31, 0x16,
	0xba, 0x24, 0x34, 0xe9, 0x9e, 0xc0, 0x27, 0xf1, 0xd2, 0x47, 0xf1, 0x51, 0x7c, 0x05, 0x6f, 0xa4,
	0x89, 0xab, 0x57, 0xba, 0x9b, 0x70, 0xc8, 0x77, 0xbe, 0x73, 0xbe, 0x73, 0xc0, 0x8c, 0x09, 0xc1,
	0x4a, 0x1a, 0x60, 0xa2, 0x02, 0x0b, 0x1b, 0xb4, 0x0d, 0x03, 0xca, 0xeb, 0x8d, 0x0a, 0x94, 0xae,
	0xea, 0xb5, 0xae, 0x2b, 0x4a, 0x56, 0x8a, 0x17, 0x52, 0x52, 0xbd, 0x92, 0x25, 0x5e, 0xd3, 0x67,
	0x51, 0x12, 0x5a, 0xad, 0x9e, 0x0a, 0x5a, 0x12, 0x28, 0x2b, 0xa1, 0x85, 0xdf, 0xb7, 0xeb, 0x10,
	0x13, 0x05, 0x5b, 0x25, 0xb8, 0x0d, 0xa1, 0x51, 0x3a, 0x3e, 0xd9, 0x19, 0xc9, 0x22, 0xc0, 0x9c,
	0x0b, 0x8d, 0x75, 0x21, 0xb8, 0xb2, 0xcb, 0xc3, 0x17, 0x07, 0x9c, 0x66, 0xad, 0x57, 0x66, 0xad,
	0xd2, 0x1f, 0xa7, 0x49, 0x63, 0x14, 0xf3, 0x7a, 0x33, 0x7c, 0x00, 0x83, 0xbd, 0x44, 0xff, 0x08,
	0xf4, 0x16, 0x49, 0x96, 0xc6, 0x97, 0xf3, 0xc9, 0x3c, 0x1e, 0x7b, 0xff, 0xfc, 0x1e, 0xf8, 0xbf,
	0x48, 0xae, 0x92, 0xeb, 0xbb, 0xc4, 0x73, 0x7c, 0x00, 0xba, 0xb3, 0x38, 0x1a, 0xc7, 0x37, 0x9e,
	0xeb, 0x1f, 0x82, 0x83, 0x2c, 0x99, 0xa7, 0x69, 0x7c, 0x9b, 0x79, 0x9d, 0xd1, 0xa7, 0x03, 0x06,
	0x6b, 0xb1, 0x81, 0x7f, 0x46, 0x19, 0x9d, 0xed, 0x3d, 0x20, 0x6d, 0x42, 0xa5, 0xce, 0xfd, 0xe8,
	0x5b, 0x88, 0x89, 0x12, 0x73, 0x06, 0x45, 0xc5, 0x02, 0x46, 0xb9, 0x89, 0xbc, 0x6b, 0x5b, 0x16,
	0xea, 0x97, 0xf2, 0x2f, 0xcc, 0xfb, 0xea, 0x76, 0xa6, 0x51, 0xf4, 0xe6, 0xf6, 0xa7, 0x56, 0x2a,
	0x22, 0x0a, 0x5a, 0xd8, 0xa0, 0x65, 0x08, 0x9b, 0x56, 0xd4, 0xfb, 0x6e, 0x9e, 0x47, 0x44, 0xe5,
	0xed, 0x3c, 0x5f, 0x86, 0xb9, 0x99, 0x7f, 0xb8, 0x03, 0xfb, 0x89, 0x50, 0x44, 0x14, 0x42, 0x2d,
	0x03, 0xa1, 0x65, 0x88, 0x90, 0xe1, 0x3c, 0x76, 0xcd, 0x61, 0xe7, 0x5f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xad, 0xa6, 0xff, 0xbe, 0x14, 0x02, 0x00, 0x00,
}

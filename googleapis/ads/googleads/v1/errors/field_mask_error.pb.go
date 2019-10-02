// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/field_mask_error.proto

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

// Enum describing possible field mask errors.
type FieldMaskErrorEnum_FieldMaskError int32

const (
	// Enum unspecified.
	FieldMaskErrorEnum_UNSPECIFIED FieldMaskErrorEnum_FieldMaskError = 0
	// The received error code is not known in this version.
	FieldMaskErrorEnum_UNKNOWN FieldMaskErrorEnum_FieldMaskError = 1
	// The field mask must be provided for update operations.
	FieldMaskErrorEnum_FIELD_MASK_MISSING FieldMaskErrorEnum_FieldMaskError = 5
	// The field mask must be empty for create and remove operations.
	FieldMaskErrorEnum_FIELD_MASK_NOT_ALLOWED FieldMaskErrorEnum_FieldMaskError = 4
	// The field mask contained an invalid field.
	FieldMaskErrorEnum_FIELD_NOT_FOUND FieldMaskErrorEnum_FieldMaskError = 2
	// The field mask updated a field with subfields. Fields with subfields may
	// be cleared, but not updated. To fix this, the field mask should select
	// all the subfields of the invalid field.
	FieldMaskErrorEnum_FIELD_HAS_SUBFIELDS FieldMaskErrorEnum_FieldMaskError = 3
)

var FieldMaskErrorEnum_FieldMaskError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	5: "FIELD_MASK_MISSING",
	4: "FIELD_MASK_NOT_ALLOWED",
	2: "FIELD_NOT_FOUND",
	3: "FIELD_HAS_SUBFIELDS",
}

var FieldMaskErrorEnum_FieldMaskError_value = map[string]int32{
	"UNSPECIFIED":            0,
	"UNKNOWN":                1,
	"FIELD_MASK_MISSING":     5,
	"FIELD_MASK_NOT_ALLOWED": 4,
	"FIELD_NOT_FOUND":        2,
	"FIELD_HAS_SUBFIELDS":    3,
}

func (x FieldMaskErrorEnum_FieldMaskError) String() string {
	return proto.EnumName(FieldMaskErrorEnum_FieldMaskError_name, int32(x))
}

func (FieldMaskErrorEnum_FieldMaskError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bb7b8d8152b776ca, []int{0, 0}
}

// Container for enum describing possible field mask errors.
type FieldMaskErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldMaskErrorEnum) Reset()         { *m = FieldMaskErrorEnum{} }
func (m *FieldMaskErrorEnum) String() string { return proto.CompactTextString(m) }
func (*FieldMaskErrorEnum) ProtoMessage()    {}
func (*FieldMaskErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb7b8d8152b776ca, []int{0}
}

func (m *FieldMaskErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldMaskErrorEnum.Unmarshal(m, b)
}
func (m *FieldMaskErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldMaskErrorEnum.Marshal(b, m, deterministic)
}
func (m *FieldMaskErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldMaskErrorEnum.Merge(m, src)
}
func (m *FieldMaskErrorEnum) XXX_Size() int {
	return xxx_messageInfo_FieldMaskErrorEnum.Size(m)
}
func (m *FieldMaskErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldMaskErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FieldMaskErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.FieldMaskErrorEnum_FieldMaskError", FieldMaskErrorEnum_FieldMaskError_name, FieldMaskErrorEnum_FieldMaskError_value)
	proto.RegisterType((*FieldMaskErrorEnum)(nil), "google.ads.googleads.v1.errors.FieldMaskErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/field_mask_error.proto", fileDescriptor_bb7b8d8152b776ca)
}

var fileDescriptor_bb7b8d8152b776ca = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xd1, 0x4a, 0xfb, 0x30,
	0x14, 0xc6, 0xff, 0xed, 0xfe, 0x2a, 0x64, 0xe0, 0x4a, 0x06, 0x13, 0x86, 0xec, 0xa2, 0x0f, 0x90,
	0x52, 0xc4, 0x9b, 0x78, 0x95, 0xd9, 0x76, 0x96, 0x6d, 0xe9, 0x20, 0x76, 0x03, 0x29, 0x94, 0x68,
	0x6b, 0x29, 0xdb, 0x9a, 0xd1, 0xcc, 0x3d, 0x87, 0xcf, 0xe0, 0x8d, 0xe0, 0xa3, 0xf8, 0x28, 0x5e,
	0xf8, 0x0c, 0xd2, 0xc6, 0x15, 0x77, 0xa1, 0x57, 0xf9, 0xf2, 0xe5, 0xf7, 0x9d, 0x9c, 0x73, 0xc0,
	0x65, 0x26, 0x44, 0xb6, 0x4a, 0x2d, 0x9e, 0x48, 0x4b, 0xc9, 0x4a, 0xed, 0x6c, 0x2b, 0x2d, 0x4b,
	0x51, 0x4a, 0xeb, 0x31, 0x4f, 0x57, 0x49, 0xbc, 0xe6, 0x72, 0x19, 0xd7, 0x0e, 0xda, 0x94, 0x62,
	0x2b, 0xe0, 0x40, 0xb1, 0x88, 0x27, 0x12, 0x35, 0x31, 0xb4, 0xb3, 0x91, 0x8a, 0xf5, 0xcf, 0xf7,
	0x65, 0x37, 0xb9, 0xc5, 0x8b, 0x42, 0x6c, 0xf9, 0x36, 0x17, 0x85, 0x54, 0x69, 0xf3, 0x55, 0x03,
	0xd0, 0xab, 0x0a, 0x4f, 0xb9, 0x5c, 0xba, 0x55, 0xc2, 0x2d, 0x9e, 0xd6, 0xe6, 0xb3, 0x06, 0x4e,
	0x0f, 0x6d, 0xd8, 0x01, 0xed, 0x90, 0xb2, 0x99, 0x7b, 0xed, 0x7b, 0xbe, 0xeb, 0x18, 0xff, 0x60,
	0x1b, 0x9c, 0x84, 0x74, 0x4c, 0x83, 0x05, 0x35, 0x34, 0xd8, 0x03, 0xd0, 0xf3, 0xdd, 0x89, 0x13,
	0x4f, 0x09, 0x1b, 0xc7, 0x53, 0x9f, 0x31, 0x9f, 0x8e, 0x8c, 0x23, 0xd8, 0x07, 0xbd, 0x1f, 0x3e,
	0x0d, 0x6e, 0x63, 0x32, 0x99, 0x04, 0x0b, 0xd7, 0x31, 0xfe, 0xc3, 0x2e, 0xe8, 0xa8, 0xb7, 0xca,
	0xf6, 0x82, 0x90, 0x3a, 0x86, 0x0e, 0xcf, 0x40, 0x57, 0x99, 0x37, 0x84, 0xc5, 0x2c, 0x1c, 0xd6,
	0x17, 0x66, 0xb4, 0x86, 0x9f, 0x1a, 0x30, 0x1f, 0xc4, 0x1a, 0xfd, 0x3d, 0xee, 0xb0, 0x7b, 0xd8,
	0xf6, 0xac, 0x9a, 0x72, 0xa6, 0xdd, 0x39, 0xdf, 0xb1, 0x4c, 0xac, 0x78, 0x91, 0x21, 0x51, 0x66,
	0x56, 0x96, 0x16, 0xf5, 0x0e, 0xf6, 0xcb, 0xde, 0xe4, 0xf2, 0xb7, 0xdd, 0x5f, 0xa9, 0xe3, 0x45,
	0x6f, 0x8d, 0x08, 0x79, 0xd3, 0x07, 0x23, 0x55, 0x8c, 0x24, 0x12, 0x29, 0x59, 0xa9, 0xb9, 0x8d,
	0xea, 0x2f, 0xe5, 0xfb, 0x1e, 0x88, 0x48, 0x22, 0xa3, 0x06, 0x88, 0xe6, 0x76, 0xa4, 0x80, 0x0f,
	0xdd, 0x54, 0x2e, 0xc6, 0x24, 0x91, 0x18, 0x37, 0x08, 0xc6, 0x73, 0x1b, 0x63, 0x05, 0xdd, 0x1f,
	0xd7, 0xdd, 0x5d, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0x56, 0x39, 0xab, 0xbc, 0x18, 0x02, 0x00,
	0x00,
}

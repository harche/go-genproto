// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/null_error.proto

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

// Enum describing possible null errors.
type NullErrorEnum_NullError int32

const (
	// Enum unspecified.
	NullErrorEnum_UNSPECIFIED NullErrorEnum_NullError = 0
	// The received error code is not known in this version.
	NullErrorEnum_UNKNOWN NullErrorEnum_NullError = 1
	// Specified list/container must not contain any null elements
	NullErrorEnum_NULL_CONTENT NullErrorEnum_NullError = 2
)

var NullErrorEnum_NullError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "NULL_CONTENT",
}

var NullErrorEnum_NullError_value = map[string]int32{
	"UNSPECIFIED":  0,
	"UNKNOWN":      1,
	"NULL_CONTENT": 2,
}

func (x NullErrorEnum_NullError) String() string {
	return proto.EnumName(NullErrorEnum_NullError_name, int32(x))
}

func (NullErrorEnum_NullError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1aea2f2f74eaad78, []int{0, 0}
}

// Container for enum describing possible null errors.
type NullErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NullErrorEnum) Reset()         { *m = NullErrorEnum{} }
func (m *NullErrorEnum) String() string { return proto.CompactTextString(m) }
func (*NullErrorEnum) ProtoMessage()    {}
func (*NullErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_1aea2f2f74eaad78, []int{0}
}

func (m *NullErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NullErrorEnum.Unmarshal(m, b)
}
func (m *NullErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NullErrorEnum.Marshal(b, m, deterministic)
}
func (m *NullErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NullErrorEnum.Merge(m, src)
}
func (m *NullErrorEnum) XXX_Size() int {
	return xxx_messageInfo_NullErrorEnum.Size(m)
}
func (m *NullErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_NullErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_NullErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.NullErrorEnum_NullError", NullErrorEnum_NullError_name, NullErrorEnum_NullError_value)
	proto.RegisterType((*NullErrorEnum)(nil), "google.ads.googleads.v1.errors.NullErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/null_error.proto", fileDescriptor_1aea2f2f74eaad78)
}

var fileDescriptor_1aea2f2f74eaad78 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xdf, 0x6a, 0x83, 0x30,
	0x14, 0xc6, 0xa7, 0x83, 0x8d, 0xa5, 0xfb, 0x23, 0x5e, 0x8e, 0xd1, 0x0b, 0x1f, 0x20, 0x41, 0x76,
	0x97, 0x5e, 0xd9, 0xd6, 0x95, 0x32, 0x49, 0x85, 0x55, 0x07, 0x43, 0x28, 0x6e, 0x4a, 0x10, 0xd2,
	0x1c, 0x31, 0xda, 0x07, 0xda, 0xe5, 0x1e, 0x65, 0x6f, 0xb2, 0x3d, 0xc5, 0xd0, 0xd4, 0xdc, 0xad,
	0x57, 0xf9, 0x11, 0x7e, 0xdf, 0x39, 0x1f, 0x07, 0x11, 0x0e, 0xc0, 0x45, 0x49, 0xf2, 0x42, 0x1d,
	0xb1, 0xa7, 0x83, 0x4f, 0xca, 0xa6, 0x81, 0x46, 0x11, 0xd9, 0x09, 0xb1, 0x1b, 0x18, 0xd7, 0x0d,
	0xb4, 0xe0, 0x4e, 0xb5, 0x85, 0xf3, 0x42, 0x61, 0x13, 0xc0, 0x07, 0x1f, 0xeb, 0xc0, 0xfd, 0xc3,
	0x38, 0xb0, 0xae, 0x48, 0x2e, 0x25, 0xb4, 0x79, 0x5b, 0x81, 0x54, 0x3a, 0xed, 0x45, 0xe8, 0x86,
	0x75, 0x42, 0x84, 0xbd, 0x1b, 0xca, 0x6e, 0xef, 0xcd, 0xd0, 0x95, 0xf9, 0x70, 0xef, 0xd0, 0x24,
	0x61, 0x2f, 0x71, 0xb8, 0x58, 0x3f, 0xad, 0xc3, 0xa5, 0x73, 0xe6, 0x4e, 0xd0, 0x65, 0xc2, 0x9e,
	0xd9, 0xe6, 0x95, 0x39, 0x96, 0xeb, 0xa0, 0x6b, 0x96, 0x44, 0xd1, 0x6e, 0xb1, 0x61, 0xdb, 0x90,
	0x6d, 0x1d, 0x7b, 0xfe, 0x63, 0x21, 0xef, 0x03, 0xf6, 0xf8, 0x74, 0xa5, 0xf9, 0xad, 0xd9, 0x10,
	0xf7, 0x25, 0x62, 0xeb, 0x6d, 0x79, 0x4c, 0x70, 0x10, 0xb9, 0xe4, 0x18, 0x1a, 0x4e, 0x78, 0x29,
	0x87, 0x8a, 0xe3, 0x15, 0xea, 0x4a, 0xfd, 0x77, 0x94, 0x99, 0x7e, 0x3e, 0xed, 0xf3, 0x55, 0x10,
	0x7c, 0xd9, 0xd3, 0x95, 0x1e, 0x16, 0x14, 0x0a, 0x6b, 0xec, 0x29, 0xf5, 0xf1, 0xb0, 0x52, 0x7d,
	0x8f, 0x42, 0x16, 0x14, 0x2a, 0x33, 0x42, 0x96, 0xfa, 0x99, 0x16, 0x7e, 0x6d, 0x4f, 0xff, 0x52,
	0x1a, 0x14, 0x8a, 0x52, 0xa3, 0x50, 0x9a, 0xfa, 0x94, 0x6a, 0xe9, 0xfd, 0x62, 0x68, 0xf7, 0xf8,
	0x17, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xcd, 0xb2, 0x8f, 0xb1, 0x01, 0x00, 0x00,
}

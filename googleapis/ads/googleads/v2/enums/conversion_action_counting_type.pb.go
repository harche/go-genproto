// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/conversion_action_counting_type.proto

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

// Indicates how conversions for this action will be counted. For more
// information, see https://support.google.com/google-ads/answer/3438531.
type ConversionActionCountingTypeEnum_ConversionActionCountingType int32

const (
	// Not specified.
	ConversionActionCountingTypeEnum_UNSPECIFIED ConversionActionCountingTypeEnum_ConversionActionCountingType = 0
	// Used for return value only. Represents value unknown in this version.
	ConversionActionCountingTypeEnum_UNKNOWN ConversionActionCountingTypeEnum_ConversionActionCountingType = 1
	// Count only one conversion per click.
	ConversionActionCountingTypeEnum_ONE_PER_CLICK ConversionActionCountingTypeEnum_ConversionActionCountingType = 2
	// Count all conversions per click.
	ConversionActionCountingTypeEnum_MANY_PER_CLICK ConversionActionCountingTypeEnum_ConversionActionCountingType = 3
)

var ConversionActionCountingTypeEnum_ConversionActionCountingType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ONE_PER_CLICK",
	3: "MANY_PER_CLICK",
}

var ConversionActionCountingTypeEnum_ConversionActionCountingType_value = map[string]int32{
	"UNSPECIFIED":    0,
	"UNKNOWN":        1,
	"ONE_PER_CLICK":  2,
	"MANY_PER_CLICK": 3,
}

func (x ConversionActionCountingTypeEnum_ConversionActionCountingType) String() string {
	return proto.EnumName(ConversionActionCountingTypeEnum_ConversionActionCountingType_name, int32(x))
}

func (ConversionActionCountingTypeEnum_ConversionActionCountingType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_605b93e01867841c, []int{0, 0}
}

// Container for enum describing the conversion deduplication mode for
// conversion optimizer.
type ConversionActionCountingTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConversionActionCountingTypeEnum) Reset()         { *m = ConversionActionCountingTypeEnum{} }
func (m *ConversionActionCountingTypeEnum) String() string { return proto.CompactTextString(m) }
func (*ConversionActionCountingTypeEnum) ProtoMessage()    {}
func (*ConversionActionCountingTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_605b93e01867841c, []int{0}
}

func (m *ConversionActionCountingTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionActionCountingTypeEnum.Unmarshal(m, b)
}
func (m *ConversionActionCountingTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionActionCountingTypeEnum.Marshal(b, m, deterministic)
}
func (m *ConversionActionCountingTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionActionCountingTypeEnum.Merge(m, src)
}
func (m *ConversionActionCountingTypeEnum) XXX_Size() int {
	return xxx_messageInfo_ConversionActionCountingTypeEnum.Size(m)
}
func (m *ConversionActionCountingTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionActionCountingTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionActionCountingTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.ConversionActionCountingTypeEnum_ConversionActionCountingType", ConversionActionCountingTypeEnum_ConversionActionCountingType_name, ConversionActionCountingTypeEnum_ConversionActionCountingType_value)
	proto.RegisterType((*ConversionActionCountingTypeEnum)(nil), "google.ads.googleads.v2.enums.ConversionActionCountingTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/conversion_action_counting_type.proto", fileDescriptor_605b93e01867841c)
}

var fileDescriptor_605b93e01867841c = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xdf, 0x4a, 0xf3, 0x30,
	0x1c, 0xfd, 0xda, 0xc1, 0x27, 0x64, 0xa8, 0xb3, 0x97, 0xb2, 0x81, 0xdb, 0x03, 0xa4, 0x50, 0xef,
	0xe2, 0x55, 0x16, 0xeb, 0x18, 0xd3, 0xac, 0xa8, 0x9b, 0x28, 0x85, 0x12, 0xdb, 0x10, 0x0a, 0x5b,
	0x52, 0x96, 0x6e, 0xb0, 0x27, 0xf0, 0x3d, 0xbc, 0xf4, 0x51, 0x7c, 0x14, 0x1f, 0xc0, 0x6b, 0x49,
	0xb2, 0x15, 0x6f, 0xdc, 0x4d, 0x72, 0xf8, 0xfd, 0x39, 0xe7, 0x77, 0x0e, 0x20, 0x42, 0x29, 0xb1,
	0xe0, 0x21, 0x2b, 0x74, 0xe8, 0xa0, 0x41, 0x9b, 0x28, 0xe4, 0x72, 0xbd, 0xd4, 0x61, 0xae, 0xe4,
	0x86, 0xaf, 0x74, 0xa9, 0x64, 0xc6, 0xf2, 0xda, 0x7c, 0xb9, 0x5a, 0xcb, 0xba, 0x94, 0x22, 0xab,
	0xb7, 0x15, 0x87, 0xd5, 0x4a, 0xd5, 0x2a, 0xe8, 0xb9, 0x4d, 0xc8, 0x0a, 0x0d, 0x1b, 0x12, 0xb8,
	0x89, 0xa0, 0x25, 0x39, 0xef, 0xee, 0x35, 0xaa, 0x32, 0x64, 0x52, 0xaa, 0x9a, 0x19, 0x26, 0xed,
	0x96, 0x07, 0x6f, 0x1e, 0xb8, 0x20, 0x8d, 0x0c, 0xb6, 0x2a, 0x64, 0x27, 0xf2, 0xb8, 0xad, 0x78,
	0x2c, 0xd7, 0xcb, 0x41, 0x0e, 0xba, 0x87, 0x66, 0x82, 0x53, 0xd0, 0x9e, 0xd1, 0x87, 0x24, 0x26,
	0xe3, 0x9b, 0x71, 0x7c, 0xdd, 0xf9, 0x17, 0xb4, 0xc1, 0xd1, 0x8c, 0x4e, 0xe8, 0xf4, 0x89, 0x76,
	0xbc, 0xe0, 0x0c, 0x1c, 0x4f, 0x69, 0x9c, 0x25, 0xf1, 0x7d, 0x46, 0x6e, 0xc7, 0x64, 0xd2, 0xf1,
	0x83, 0x00, 0x9c, 0xdc, 0x61, 0xfa, 0xfc, 0xab, 0xd6, 0x1a, 0x7e, 0x7b, 0xa0, 0x9f, 0xab, 0x25,
	0x3c, 0xe8, 0x66, 0xd8, 0x3f, 0x74, 0x48, 0x62, 0x2c, 0x25, 0xde, 0xcb, 0x70, 0xc7, 0x21, 0xd4,
	0x82, 0x49, 0x01, 0xd5, 0x4a, 0x84, 0x82, 0x4b, 0x6b, 0x78, 0x1f, 0x73, 0x55, 0xea, 0x3f, 0x52,
	0xbf, 0xb2, 0xef, 0xbb, 0xdf, 0x1a, 0x61, 0xfc, 0xe1, 0xf7, 0x46, 0x8e, 0x0a, 0x17, 0x1a, 0x3a,
	0x68, 0xd0, 0x3c, 0x82, 0x26, 0x18, 0xfd, 0xb9, 0xef, 0xa7, 0xb8, 0xd0, 0x69, 0xd3, 0x4f, 0xe7,
	0x51, 0x6a, 0xfb, 0x5f, 0x7e, 0xdf, 0x15, 0x11, 0xc2, 0x85, 0x46, 0xa8, 0x99, 0x40, 0x68, 0x1e,
	0x21, 0x64, 0x67, 0x5e, 0xff, 0xdb, 0xc3, 0x2e, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x82,
	0x3f, 0x56, 0x0d, 0x02, 0x00, 0x00,
}

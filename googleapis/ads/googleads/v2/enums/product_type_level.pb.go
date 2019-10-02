// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/product_type_level.proto

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

// Enum describing the level of the type of a product offer.
type ProductTypeLevelEnum_ProductTypeLevel int32

const (
	// Not specified.
	ProductTypeLevelEnum_UNSPECIFIED ProductTypeLevelEnum_ProductTypeLevel = 0
	// Used for return value only. Represents value unknown in this version.
	ProductTypeLevelEnum_UNKNOWN ProductTypeLevelEnum_ProductTypeLevel = 1
	// Level 1.
	ProductTypeLevelEnum_LEVEL1 ProductTypeLevelEnum_ProductTypeLevel = 7
	// Level 2.
	ProductTypeLevelEnum_LEVEL2 ProductTypeLevelEnum_ProductTypeLevel = 8
	// Level 3.
	ProductTypeLevelEnum_LEVEL3 ProductTypeLevelEnum_ProductTypeLevel = 9
	// Level 4.
	ProductTypeLevelEnum_LEVEL4 ProductTypeLevelEnum_ProductTypeLevel = 10
	// Level 5.
	ProductTypeLevelEnum_LEVEL5 ProductTypeLevelEnum_ProductTypeLevel = 11
)

var ProductTypeLevelEnum_ProductTypeLevel_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	7:  "LEVEL1",
	8:  "LEVEL2",
	9:  "LEVEL3",
	10: "LEVEL4",
	11: "LEVEL5",
}

var ProductTypeLevelEnum_ProductTypeLevel_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"LEVEL1":      7,
	"LEVEL2":      8,
	"LEVEL3":      9,
	"LEVEL4":      10,
	"LEVEL5":      11,
}

func (x ProductTypeLevelEnum_ProductTypeLevel) String() string {
	return proto.EnumName(ProductTypeLevelEnum_ProductTypeLevel_name, int32(x))
}

func (ProductTypeLevelEnum_ProductTypeLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_73d48975fa861d36, []int{0, 0}
}

// Level of the type of a product offer.
type ProductTypeLevelEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductTypeLevelEnum) Reset()         { *m = ProductTypeLevelEnum{} }
func (m *ProductTypeLevelEnum) String() string { return proto.CompactTextString(m) }
func (*ProductTypeLevelEnum) ProtoMessage()    {}
func (*ProductTypeLevelEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_73d48975fa861d36, []int{0}
}

func (m *ProductTypeLevelEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductTypeLevelEnum.Unmarshal(m, b)
}
func (m *ProductTypeLevelEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductTypeLevelEnum.Marshal(b, m, deterministic)
}
func (m *ProductTypeLevelEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductTypeLevelEnum.Merge(m, src)
}
func (m *ProductTypeLevelEnum) XXX_Size() int {
	return xxx_messageInfo_ProductTypeLevelEnum.Size(m)
}
func (m *ProductTypeLevelEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductTypeLevelEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ProductTypeLevelEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.ProductTypeLevelEnum_ProductTypeLevel", ProductTypeLevelEnum_ProductTypeLevel_name, ProductTypeLevelEnum_ProductTypeLevel_value)
	proto.RegisterType((*ProductTypeLevelEnum)(nil), "google.ads.googleads.v2.enums.ProductTypeLevelEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/product_type_level.proto", fileDescriptor_73d48975fa861d36)
}

var fileDescriptor_73d48975fa861d36 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4a, 0x03, 0x31,
	0x18, 0x85, 0x9d, 0x0a, 0xad, 0xa6, 0x0b, 0xc3, 0xa0, 0x1b, 0xb1, 0x8b, 0xf6, 0x00, 0x09, 0xa6,
	0xea, 0x22, 0xae, 0xa6, 0x3a, 0x96, 0x62, 0x19, 0x07, 0xb4, 0x23, 0xc8, 0x40, 0x19, 0x9b, 0x10,
	0x0a, 0xd3, 0x24, 0x34, 0xd3, 0x42, 0xf7, 0x9e, 0xc4, 0xa5, 0x47, 0xf1, 0x28, 0x3d, 0x85, 0x4c,
	0x62, 0xb3, 0x28, 0xe8, 0x26, 0x7c, 0xe4, 0xff, 0xdf, 0xe3, 0x7f, 0x0f, 0xdc, 0x08, 0xa5, 0x44,
	0xc9, 0x71, 0xc1, 0x0c, 0x76, 0x58, 0xd3, 0x9a, 0x60, 0x2e, 0x57, 0x0b, 0x83, 0xf5, 0x52, 0xb1,
	0xd5, 0xac, 0x9a, 0x56, 0x1b, 0xcd, 0xa7, 0x25, 0x5f, 0xf3, 0x12, 0xe9, 0xa5, 0xaa, 0x54, 0xd8,
	0x71, 0xcb, 0xa8, 0x60, 0x06, 0x79, 0x1d, 0x5a, 0x13, 0x64, 0x75, 0xe7, 0x17, 0x3b, 0x5b, 0x3d,
	0xc7, 0x85, 0x94, 0xaa, 0x2a, 0xaa, 0xb9, 0x92, 0xc6, 0x89, 0x7b, 0x1f, 0x01, 0x38, 0x4d, 0x9d,
	0xf3, 0xcb, 0x46, 0xf3, 0x71, 0xed, 0x1b, 0xcb, 0xd5, 0xa2, 0x57, 0x02, 0xb8, 0xff, 0x1f, 0x9e,
	0x80, 0xf6, 0x24, 0x79, 0x4e, 0xe3, 0xbb, 0xd1, 0xc3, 0x28, 0xbe, 0x87, 0x07, 0x61, 0x1b, 0xb4,
	0x26, 0xc9, 0x63, 0xf2, 0xf4, 0x9a, 0xc0, 0x20, 0x04, 0xa0, 0x39, 0x8e, 0xb3, 0x78, 0x7c, 0x09,
	0x5b, 0x9e, 0x09, 0x3c, 0xf2, 0xdc, 0x87, 0xc7, 0x9e, 0xaf, 0x20, 0xf0, 0x7c, 0x0d, 0xdb, 0x83,
	0x6d, 0x00, 0xba, 0x33, 0xb5, 0x40, 0xff, 0x46, 0x19, 0x9c, 0xed, 0x5f, 0x94, 0xd6, 0x19, 0xd2,
	0xe0, 0x6d, 0xf0, 0xab, 0x13, 0xaa, 0x2c, 0xa4, 0x40, 0x6a, 0x29, 0xb0, 0xe0, 0xd2, 0x26, 0xdc,
	0x55, 0xa9, 0xe7, 0xe6, 0x8f, 0x66, 0x6f, 0xed, 0xfb, 0xd9, 0x38, 0x1c, 0x46, 0xd1, 0x57, 0xa3,
	0x33, 0x74, 0x56, 0x11, 0x33, 0xc8, 0x61, 0x4d, 0x19, 0x41, 0x75, 0x2b, 0xe6, 0x7b, 0x37, 0xcf,
	0x23, 0x66, 0x72, 0x3f, 0xcf, 0x33, 0x92, 0xdb, 0xf9, 0xb6, 0xd1, 0x75, 0x9f, 0x94, 0x46, 0xcc,
	0x50, 0xea, 0x37, 0x28, 0xcd, 0x08, 0xa5, 0x76, 0xe7, 0xbd, 0x69, 0x0f, 0xeb, 0xff, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x5f, 0x5f, 0x05, 0xb1, 0xf1, 0x01, 0x00, 0x00,
}

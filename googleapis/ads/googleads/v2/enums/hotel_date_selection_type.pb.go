// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/hotel_date_selection_type.proto

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

// Enum describing possible hotel date selection types.
type HotelDateSelectionTypeEnum_HotelDateSelectionType int32

const (
	// Not specified.
	HotelDateSelectionTypeEnum_UNSPECIFIED HotelDateSelectionTypeEnum_HotelDateSelectionType = 0
	// Used for return value only. Represents value unknown in this version.
	HotelDateSelectionTypeEnum_UNKNOWN HotelDateSelectionTypeEnum_HotelDateSelectionType = 1
	// Dates selected by default.
	HotelDateSelectionTypeEnum_DEFAULT_SELECTION HotelDateSelectionTypeEnum_HotelDateSelectionType = 50
	// Dates selected by the user.
	HotelDateSelectionTypeEnum_USER_SELECTED HotelDateSelectionTypeEnum_HotelDateSelectionType = 51
)

var HotelDateSelectionTypeEnum_HotelDateSelectionType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	50: "DEFAULT_SELECTION",
	51: "USER_SELECTED",
}

var HotelDateSelectionTypeEnum_HotelDateSelectionType_value = map[string]int32{
	"UNSPECIFIED":       0,
	"UNKNOWN":           1,
	"DEFAULT_SELECTION": 50,
	"USER_SELECTED":     51,
}

func (x HotelDateSelectionTypeEnum_HotelDateSelectionType) String() string {
	return proto.EnumName(HotelDateSelectionTypeEnum_HotelDateSelectionType_name, int32(x))
}

func (HotelDateSelectionTypeEnum_HotelDateSelectionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_411d1f94192acf8e, []int{0, 0}
}

// Container for enum describing possible hotel date selection types
type HotelDateSelectionTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HotelDateSelectionTypeEnum) Reset()         { *m = HotelDateSelectionTypeEnum{} }
func (m *HotelDateSelectionTypeEnum) String() string { return proto.CompactTextString(m) }
func (*HotelDateSelectionTypeEnum) ProtoMessage()    {}
func (*HotelDateSelectionTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_411d1f94192acf8e, []int{0}
}

func (m *HotelDateSelectionTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HotelDateSelectionTypeEnum.Unmarshal(m, b)
}
func (m *HotelDateSelectionTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HotelDateSelectionTypeEnum.Marshal(b, m, deterministic)
}
func (m *HotelDateSelectionTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HotelDateSelectionTypeEnum.Merge(m, src)
}
func (m *HotelDateSelectionTypeEnum) XXX_Size() int {
	return xxx_messageInfo_HotelDateSelectionTypeEnum.Size(m)
}
func (m *HotelDateSelectionTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_HotelDateSelectionTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_HotelDateSelectionTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.HotelDateSelectionTypeEnum_HotelDateSelectionType", HotelDateSelectionTypeEnum_HotelDateSelectionType_name, HotelDateSelectionTypeEnum_HotelDateSelectionType_value)
	proto.RegisterType((*HotelDateSelectionTypeEnum)(nil), "google.ads.googleads.v2.enums.HotelDateSelectionTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/hotel_date_selection_type.proto", fileDescriptor_411d1f94192acf8e)
}

var fileDescriptor_411d1f94192acf8e = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xc1, 0x6a, 0xf2, 0x40,
	0x18, 0xfc, 0xf5, 0x87, 0x16, 0x56, 0x4a, 0x63, 0xa0, 0x3d, 0xd8, 0x7a, 0xd0, 0x07, 0xd8, 0x40,
	0xbc, 0x6d, 0xe9, 0x21, 0x9a, 0xd5, 0x4a, 0x25, 0x4a, 0x63, 0x2c, 0x94, 0x40, 0xba, 0x35, 0x4b,
	0x2a, 0xc4, 0xdd, 0xe0, 0xae, 0x82, 0x97, 0x3e, 0x4c, 0x8f, 0x7d, 0x94, 0x3e, 0x4a, 0x4f, 0x7d,
	0x84, 0xb2, 0xbb, 0x26, 0x27, 0xdb, 0xcb, 0xc7, 0xf0, 0xcd, 0x37, 0xc3, 0x37, 0x03, 0x6e, 0x33,
	0xce, 0xb3, 0x9c, 0x3a, 0x24, 0x15, 0x8e, 0x81, 0x0a, 0xed, 0x5c, 0x87, 0xb2, 0xed, 0x5a, 0x38,
	0xaf, 0x5c, 0xd2, 0x3c, 0x49, 0x89, 0xa4, 0x89, 0xa0, 0x39, 0x5d, 0xca, 0x15, 0x67, 0x89, 0xdc,
	0x17, 0x14, 0x16, 0x1b, 0x2e, 0xb9, 0xdd, 0x36, 0x1a, 0x48, 0x52, 0x01, 0x2b, 0x39, 0xdc, 0xb9,
	0x50, 0xcb, 0x5b, 0xd7, 0xa5, 0x7b, 0xb1, 0x72, 0x08, 0x63, 0x5c, 0x12, 0x65, 0x20, 0x8c, 0xb8,
	0xfb, 0x06, 0x5a, 0x77, 0xca, 0xdf, 0x27, 0x92, 0x86, 0xa5, 0xfb, 0x7c, 0x5f, 0x50, 0xcc, 0xb6,
	0xeb, 0xee, 0x33, 0xb8, 0x3c, 0xce, 0xda, 0xe7, 0xa0, 0x11, 0x05, 0xe1, 0x0c, 0x0f, 0xc6, 0xc3,
	0x31, 0xf6, 0xad, 0x7f, 0x76, 0x03, 0x9c, 0x46, 0xc1, 0x7d, 0x30, 0x7d, 0x0c, 0xac, 0x9a, 0x7d,
	0x01, 0x9a, 0x3e, 0x1e, 0x7a, 0xd1, 0x64, 0x9e, 0x84, 0x78, 0x82, 0x07, 0xf3, 0xf1, 0x34, 0xb0,
	0x5c, 0xbb, 0x09, 0xce, 0xa2, 0x10, 0x3f, 0x1c, 0x76, 0xd8, 0xb7, 0x7a, 0xfd, 0xef, 0x1a, 0xe8,
	0x2c, 0xf9, 0x1a, 0xfe, 0x99, 0xa1, 0x7f, 0x75, 0xfc, 0x8b, 0x99, 0x8a, 0x30, 0xab, 0x3d, 0xf5,
	0x0f, 0xea, 0x8c, 0xe7, 0x84, 0x65, 0x90, 0x6f, 0x32, 0x27, 0xa3, 0x4c, 0x07, 0x2c, 0x0b, 0x2d,
	0x56, 0xe2, 0x97, 0x7e, 0x6f, 0xf4, 0x7c, 0xaf, 0xff, 0x1f, 0x79, 0xde, 0x47, 0xbd, 0x3d, 0x32,
	0x56, 0x5e, 0x2a, 0xa0, 0x81, 0x0a, 0x2d, 0x5c, 0xa8, 0xfa, 0x10, 0x9f, 0x25, 0x1f, 0x7b, 0xa9,
	0x88, 0x2b, 0x3e, 0x5e, 0xb8, 0xb1, 0xe6, 0xbf, 0xea, 0x1d, 0xb3, 0x44, 0xc8, 0x4b, 0x05, 0x42,
	0xd5, 0x05, 0x42, 0x0b, 0x17, 0x21, 0x7d, 0xf3, 0x72, 0xa2, 0x1f, 0xeb, 0xfd, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x23, 0xab, 0x2c, 0xb2, 0xf7, 0x01, 0x00, 0x00,
}

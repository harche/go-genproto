// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/app_url_operating_system_type.proto

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

// Operating System
type AppUrlOperatingSystemTypeEnum_AppUrlOperatingSystemType int32

const (
	// Not specified.
	AppUrlOperatingSystemTypeEnum_UNSPECIFIED AppUrlOperatingSystemTypeEnum_AppUrlOperatingSystemType = 0
	// Used for return value only. Represents value unknown in this version.
	AppUrlOperatingSystemTypeEnum_UNKNOWN AppUrlOperatingSystemTypeEnum_AppUrlOperatingSystemType = 1
	// The Apple IOS operating system.
	AppUrlOperatingSystemTypeEnum_IOS AppUrlOperatingSystemTypeEnum_AppUrlOperatingSystemType = 2
	// The Android operating system.
	AppUrlOperatingSystemTypeEnum_ANDROID AppUrlOperatingSystemTypeEnum_AppUrlOperatingSystemType = 3
)

var AppUrlOperatingSystemTypeEnum_AppUrlOperatingSystemType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "IOS",
	3: "ANDROID",
}

var AppUrlOperatingSystemTypeEnum_AppUrlOperatingSystemType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"IOS":         2,
	"ANDROID":     3,
}

func (x AppUrlOperatingSystemTypeEnum_AppUrlOperatingSystemType) String() string {
	return proto.EnumName(AppUrlOperatingSystemTypeEnum_AppUrlOperatingSystemType_name, int32(x))
}

func (AppUrlOperatingSystemTypeEnum_AppUrlOperatingSystemType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4f2dc8e3a8835a20, []int{0, 0}
}

// The possible OS types for a deeplink AppUrl.
type AppUrlOperatingSystemTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppUrlOperatingSystemTypeEnum) Reset()         { *m = AppUrlOperatingSystemTypeEnum{} }
func (m *AppUrlOperatingSystemTypeEnum) String() string { return proto.CompactTextString(m) }
func (*AppUrlOperatingSystemTypeEnum) ProtoMessage()    {}
func (*AppUrlOperatingSystemTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f2dc8e3a8835a20, []int{0}
}

func (m *AppUrlOperatingSystemTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppUrlOperatingSystemTypeEnum.Unmarshal(m, b)
}
func (m *AppUrlOperatingSystemTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppUrlOperatingSystemTypeEnum.Marshal(b, m, deterministic)
}
func (m *AppUrlOperatingSystemTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppUrlOperatingSystemTypeEnum.Merge(m, src)
}
func (m *AppUrlOperatingSystemTypeEnum) XXX_Size() int {
	return xxx_messageInfo_AppUrlOperatingSystemTypeEnum.Size(m)
}
func (m *AppUrlOperatingSystemTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AppUrlOperatingSystemTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AppUrlOperatingSystemTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.AppUrlOperatingSystemTypeEnum_AppUrlOperatingSystemType", AppUrlOperatingSystemTypeEnum_AppUrlOperatingSystemType_name, AppUrlOperatingSystemTypeEnum_AppUrlOperatingSystemType_value)
	proto.RegisterType((*AppUrlOperatingSystemTypeEnum)(nil), "google.ads.googleads.v2.enums.AppUrlOperatingSystemTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/app_url_operating_system_type.proto", fileDescriptor_4f2dc8e3a8835a20)
}

var fileDescriptor_4f2dc8e3a8835a20 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x75, 0x1d, 0x38, 0xc8, 0x0e, 0x96, 0xde, 0x14, 0x27, 0x6c, 0x3f, 0x20, 0x81, 0x7a, 0x8b,
	0xa7, 0xcc, 0xcd, 0x51, 0x84, 0x76, 0x38, 0x37, 0x41, 0x0a, 0x25, 0xda, 0x10, 0x0a, 0x6d, 0x12,
	0x9a, 0x6c, 0xd0, 0xbf, 0xe3, 0xd1, 0x9f, 0xe2, 0x4f, 0xf1, 0xea, 0x1f, 0x90, 0x26, 0xb6, 0xb7,
	0x7a, 0x09, 0x8f, 0xef, 0x7b, 0xdf, 0x7b, 0x2f, 0x0f, 0x10, 0x2e, 0x25, 0x2f, 0x19, 0xa2, 0xb9,
	0x46, 0x0e, 0xb6, 0xe8, 0x14, 0x22, 0x26, 0x8e, 0x95, 0x46, 0x54, 0xa9, 0xec, 0x58, 0x97, 0x99,
	0x54, 0xac, 0xa6, 0xa6, 0x10, 0x3c, 0xd3, 0x8d, 0x36, 0xac, 0xca, 0x4c, 0xa3, 0x18, 0x54, 0xb5,
	0x34, 0x32, 0x98, 0xb9, 0x3b, 0x48, 0x73, 0x0d, 0x7b, 0x09, 0x78, 0x0a, 0xa1, 0x95, 0xb8, 0xba,
	0xee, 0x1c, 0x54, 0x81, 0xa8, 0x10, 0xd2, 0x50, 0x53, 0x48, 0xa1, 0xdd, 0xf1, 0x42, 0x81, 0x19,
	0x51, 0x6a, 0x5f, 0x97, 0x49, 0xe7, 0xb0, 0xb3, 0x06, 0xcf, 0x8d, 0x62, 0x6b, 0x71, 0xac, 0x16,
	0x09, 0xb8, 0x1c, 0x24, 0x04, 0x17, 0x60, 0xba, 0x8f, 0x77, 0xdb, 0xf5, 0x7d, 0xf4, 0x10, 0xad,
	0x57, 0xfe, 0x59, 0x30, 0x05, 0x93, 0x7d, 0xfc, 0x18, 0x27, 0x2f, 0xb1, 0x3f, 0x0a, 0x26, 0x60,
	0x1c, 0x25, 0x3b, 0xdf, 0x6b, 0xa7, 0x24, 0x5e, 0x3d, 0x25, 0xd1, 0xca, 0x1f, 0x2f, 0x7f, 0x46,
	0x60, 0xfe, 0x2e, 0x2b, 0xf8, 0x6f, 0xea, 0xe5, 0xcd, 0xa0, 0xe9, 0xb6, 0xcd, 0xbd, 0x1d, 0xbd,
	0x2e, 0xff, 0x04, 0xb8, 0x2c, 0xa9, 0xe0, 0x50, 0xd6, 0x1c, 0x71, 0x26, 0xec, 0xaf, 0xba, 0x26,
	0x55, 0xa1, 0x07, 0x8a, 0xbd, 0xb3, 0xef, 0x87, 0x37, 0xde, 0x10, 0xf2, 0xe9, 0xcd, 0x36, 0x4e,
	0x8a, 0xe4, 0x1a, 0x3a, 0xd8, 0xa2, 0x43, 0x08, 0xdb, 0x06, 0xf4, 0x57, 0xb7, 0x4f, 0x49, 0xae,
	0xd3, 0x7e, 0x9f, 0x1e, 0xc2, 0xd4, 0xee, 0xbf, 0xbd, 0xb9, 0x1b, 0x62, 0x4c, 0x72, 0x8d, 0x71,
	0xcf, 0xc0, 0xf8, 0x10, 0x62, 0x6c, 0x39, 0x6f, 0xe7, 0x36, 0xd8, 0xed, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x32, 0xfb, 0x78, 0x0c, 0xf0, 0x01, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/app_payment_model_type.proto

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

// Enum describing possible app payment models.
type AppPaymentModelTypeEnum_AppPaymentModelType int32

const (
	// Not specified.
	AppPaymentModelTypeEnum_UNSPECIFIED AppPaymentModelTypeEnum_AppPaymentModelType = 0
	// Used for return value only. Represents value unknown in this version.
	AppPaymentModelTypeEnum_UNKNOWN AppPaymentModelTypeEnum_AppPaymentModelType = 1
	// Represents paid-for apps.
	AppPaymentModelTypeEnum_PAID AppPaymentModelTypeEnum_AppPaymentModelType = 30
)

var AppPaymentModelTypeEnum_AppPaymentModelType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	30: "PAID",
}

var AppPaymentModelTypeEnum_AppPaymentModelType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"PAID":        30,
}

func (x AppPaymentModelTypeEnum_AppPaymentModelType) String() string {
	return proto.EnumName(AppPaymentModelTypeEnum_AppPaymentModelType_name, int32(x))
}

func (AppPaymentModelTypeEnum_AppPaymentModelType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_35fd4c9053283d24, []int{0, 0}
}

// Represents a criterion for targeting paid apps.
type AppPaymentModelTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppPaymentModelTypeEnum) Reset()         { *m = AppPaymentModelTypeEnum{} }
func (m *AppPaymentModelTypeEnum) String() string { return proto.CompactTextString(m) }
func (*AppPaymentModelTypeEnum) ProtoMessage()    {}
func (*AppPaymentModelTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_35fd4c9053283d24, []int{0}
}

func (m *AppPaymentModelTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppPaymentModelTypeEnum.Unmarshal(m, b)
}
func (m *AppPaymentModelTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppPaymentModelTypeEnum.Marshal(b, m, deterministic)
}
func (m *AppPaymentModelTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppPaymentModelTypeEnum.Merge(m, src)
}
func (m *AppPaymentModelTypeEnum) XXX_Size() int {
	return xxx_messageInfo_AppPaymentModelTypeEnum.Size(m)
}
func (m *AppPaymentModelTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AppPaymentModelTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AppPaymentModelTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.AppPaymentModelTypeEnum_AppPaymentModelType", AppPaymentModelTypeEnum_AppPaymentModelType_name, AppPaymentModelTypeEnum_AppPaymentModelType_value)
	proto.RegisterType((*AppPaymentModelTypeEnum)(nil), "google.ads.googleads.v2.enums.AppPaymentModelTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/app_payment_model_type.proto", fileDescriptor_35fd4c9053283d24)
}

var fileDescriptor_35fd4c9053283d24 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xd1, 0x4a, 0xf3, 0x30,
	0x1c, 0xc5, 0xbf, 0xed, 0x13, 0x95, 0xec, 0xc2, 0x51, 0x2f, 0x14, 0x71, 0xc2, 0xf6, 0x00, 0x09,
	0xd4, 0xbb, 0x88, 0x17, 0xa9, 0xab, 0xa3, 0x88, 0xb5, 0xa0, 0xab, 0x22, 0x85, 0x12, 0x4d, 0x08,
	0x85, 0x36, 0x09, 0x4b, 0x37, 0xe8, 0xeb, 0x78, 0xe9, 0xa3, 0xf8, 0x28, 0x82, 0xef, 0x20, 0x4d,
	0x6c, 0xaf, 0xa6, 0x37, 0xe1, 0x90, 0xf3, 0xff, 0x9d, 0x9c, 0xfc, 0x01, 0x16, 0x4a, 0x89, 0x92,
	0x23, 0xca, 0x0c, 0x72, 0xb2, 0x55, 0x1b, 0x1f, 0x71, 0xb9, 0xae, 0x0c, 0xa2, 0x5a, 0xe7, 0x9a,
	0x36, 0x15, 0x97, 0x75, 0x5e, 0x29, 0xc6, 0xcb, 0xbc, 0x6e, 0x34, 0x87, 0x7a, 0xa5, 0x6a, 0xe5,
	0x4d, 0x1c, 0x00, 0x29, 0x33, 0xb0, 0x67, 0xe1, 0xc6, 0x87, 0x96, 0x3d, 0x39, 0xed, 0xa2, 0x75,
	0x81, 0xa8, 0x94, 0xaa, 0xa6, 0x75, 0xa1, 0xa4, 0x71, 0xf0, 0xec, 0x09, 0x1c, 0x11, 0xad, 0x13,
	0x97, 0x7d, 0xdb, 0x46, 0x3f, 0x34, 0x9a, 0x87, 0x72, 0x5d, 0xcd, 0x2e, 0xc1, 0xe1, 0x16, 0xcb,
	0x3b, 0x00, 0xa3, 0x65, 0x7c, 0x9f, 0x84, 0x57, 0xd1, 0x75, 0x14, 0xce, 0xc7, 0xff, 0xbc, 0x11,
	0xd8, 0x5b, 0xc6, 0x37, 0xf1, 0xdd, 0x63, 0x3c, 0x1e, 0x78, 0xfb, 0x60, 0x27, 0x21, 0xd1, 0x7c,
	0x7c, 0x16, 0x7c, 0x0d, 0xc0, 0xf4, 0x55, 0x55, 0xf0, 0xcf, 0x76, 0xc1, 0xf1, 0x96, 0x27, 0x92,
	0xb6, 0x59, 0x32, 0x78, 0x0e, 0x7e, 0x50, 0xa1, 0x4a, 0x2a, 0x05, 0x54, 0x2b, 0x81, 0x04, 0x97,
	0xb6, 0x77, 0xb7, 0x24, 0x5d, 0x98, 0x5f, 0x76, 0x76, 0x61, 0xcf, 0xb7, 0xe1, 0xff, 0x05, 0x21,
	0xef, 0xc3, 0xc9, 0xc2, 0x45, 0x11, 0x66, 0xa0, 0x93, 0xad, 0x4a, 0x7d, 0xd8, 0xfe, 0xd4, 0x7c,
	0x74, 0x7e, 0x46, 0x98, 0xc9, 0x7a, 0x3f, 0x4b, 0xfd, 0xcc, 0xfa, 0x9f, 0xc3, 0xa9, 0xbb, 0xc4,
	0x98, 0x30, 0x83, 0x71, 0x3f, 0x81, 0x71, 0xea, 0x63, 0x6c, 0x67, 0x5e, 0x76, 0x6d, 0xb1, 0xf3,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x36, 0x1d, 0xdf, 0xbb, 0xcb, 0x01, 0x00, 0x00,
}

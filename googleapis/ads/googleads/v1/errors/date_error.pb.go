// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/date_error.proto

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

// Enum describing possible date errors.
type DateErrorEnum_DateError int32

const (
	// Enum unspecified.
	DateErrorEnum_UNSPECIFIED DateErrorEnum_DateError = 0
	// The received error code is not known in this version.
	DateErrorEnum_UNKNOWN DateErrorEnum_DateError = 1
	// Given field values do not correspond to a valid date.
	DateErrorEnum_INVALID_FIELD_VALUES_IN_DATE DateErrorEnum_DateError = 2
	// Given field values do not correspond to a valid date time.
	DateErrorEnum_INVALID_FIELD_VALUES_IN_DATE_TIME DateErrorEnum_DateError = 3
	// The string date's format should be yyyy-mm-dd.
	DateErrorEnum_INVALID_STRING_DATE DateErrorEnum_DateError = 4
	// The string date time's format should be yyyy-mm-dd hh:mm:ss.ssssss.
	DateErrorEnum_INVALID_STRING_DATE_TIME_MICROS DateErrorEnum_DateError = 6
	// The string date time's format should be yyyy-mm-dd hh:mm:ss.
	DateErrorEnum_INVALID_STRING_DATE_TIME_SECONDS DateErrorEnum_DateError = 11
	// The string date time's format should be yyyy-mm-dd hh:mm:ss+|-hh:mm.
	DateErrorEnum_INVALID_STRING_DATE_TIME_SECONDS_WITH_OFFSET DateErrorEnum_DateError = 12
	// Date is before allowed minimum.
	DateErrorEnum_EARLIER_THAN_MINIMUM_DATE DateErrorEnum_DateError = 7
	// Date is after allowed maximum.
	DateErrorEnum_LATER_THAN_MAXIMUM_DATE DateErrorEnum_DateError = 8
	// Date range bounds are not in order.
	DateErrorEnum_DATE_RANGE_MINIMUM_DATE_LATER_THAN_MAXIMUM_DATE DateErrorEnum_DateError = 9
	// Both dates in range are null.
	DateErrorEnum_DATE_RANGE_MINIMUM_AND_MAXIMUM_DATES_BOTH_NULL DateErrorEnum_DateError = 10
)

var DateErrorEnum_DateError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "INVALID_FIELD_VALUES_IN_DATE",
	3:  "INVALID_FIELD_VALUES_IN_DATE_TIME",
	4:  "INVALID_STRING_DATE",
	6:  "INVALID_STRING_DATE_TIME_MICROS",
	11: "INVALID_STRING_DATE_TIME_SECONDS",
	12: "INVALID_STRING_DATE_TIME_SECONDS_WITH_OFFSET",
	7:  "EARLIER_THAN_MINIMUM_DATE",
	8:  "LATER_THAN_MAXIMUM_DATE",
	9:  "DATE_RANGE_MINIMUM_DATE_LATER_THAN_MAXIMUM_DATE",
	10: "DATE_RANGE_MINIMUM_AND_MAXIMUM_DATES_BOTH_NULL",
}

var DateErrorEnum_DateError_value = map[string]int32{
	"UNSPECIFIED":                                     0,
	"UNKNOWN":                                         1,
	"INVALID_FIELD_VALUES_IN_DATE":                    2,
	"INVALID_FIELD_VALUES_IN_DATE_TIME":               3,
	"INVALID_STRING_DATE":                             4,
	"INVALID_STRING_DATE_TIME_MICROS":                 6,
	"INVALID_STRING_DATE_TIME_SECONDS":                11,
	"INVALID_STRING_DATE_TIME_SECONDS_WITH_OFFSET":    12,
	"EARLIER_THAN_MINIMUM_DATE":                       7,
	"LATER_THAN_MAXIMUM_DATE":                         8,
	"DATE_RANGE_MINIMUM_DATE_LATER_THAN_MAXIMUM_DATE": 9,
	"DATE_RANGE_MINIMUM_AND_MAXIMUM_DATES_BOTH_NULL":  10,
}

func (x DateErrorEnum_DateError) String() string {
	return proto.EnumName(DateErrorEnum_DateError_name, int32(x))
}

func (DateErrorEnum_DateError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_df8c01899fe51563, []int{0, 0}
}

// Container for enum describing possible date errors.
type DateErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DateErrorEnum) Reset()         { *m = DateErrorEnum{} }
func (m *DateErrorEnum) String() string { return proto.CompactTextString(m) }
func (*DateErrorEnum) ProtoMessage()    {}
func (*DateErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_df8c01899fe51563, []int{0}
}

func (m *DateErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DateErrorEnum.Unmarshal(m, b)
}
func (m *DateErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DateErrorEnum.Marshal(b, m, deterministic)
}
func (m *DateErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateErrorEnum.Merge(m, src)
}
func (m *DateErrorEnum) XXX_Size() int {
	return xxx_messageInfo_DateErrorEnum.Size(m)
}
func (m *DateErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_DateErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_DateErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.DateErrorEnum_DateError", DateErrorEnum_DateError_name, DateErrorEnum_DateError_value)
	proto.RegisterType((*DateErrorEnum)(nil), "google.ads.googleads.v1.errors.DateErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/date_error.proto", fileDescriptor_df8c01899fe51563)
}

var fileDescriptor_df8c01899fe51563 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xdd, 0x6e, 0xd3, 0x30,
	0x18, 0x86, 0x69, 0x8b, 0x36, 0xf6, 0x95, 0x9f, 0xc8, 0x1c, 0x4c, 0xc0, 0x18, 0xa3, 0xc0, 0x19,
	0x72, 0x08, 0x3b, 0x0b, 0x47, 0x6e, 0xe3, 0xb6, 0x16, 0x89, 0x53, 0xe5, 0x6f, 0x08, 0x55, 0xb2,
	0x02, 0x89, 0xa2, 0x4a, 0x5b, 0x5c, 0xc5, 0x61, 0x17, 0xc4, 0x21, 0x12, 0xd7, 0xc0, 0x39, 0x77,
	0x02, 0x57, 0x81, 0x1a, 0x2f, 0x41, 0x93, 0xd6, 0xee, 0x28, 0xaf, 0xfc, 0x3d, 0xcf, 0x17, 0x4b,
	0x7e, 0xc1, 0x2c, 0xa4, 0x2c, 0xce, 0x73, 0x33, 0xcd, 0xd4, 0x55, 0xdc, 0xa4, 0x4b, 0xcb, 0xcc,
	0xab, 0x4a, 0x56, 0xca, 0xcc, 0xd2, 0x3a, 0x17, 0x4d, 0xc6, 0xeb, 0x4a, 0xd6, 0x12, 0x1d, 0x6b,
	0x0a, 0xa7, 0x99, 0xc2, 0x9d, 0x80, 0x2f, 0x2d, 0xac, 0x85, 0xa7, 0x47, 0xed, 0xc2, 0xf5, 0xca,
	0x4c, 0xcb, 0x52, 0xd6, 0x69, 0xbd, 0x92, 0xa5, 0xd2, 0xf6, 0xe8, 0xd7, 0x00, 0x1e, 0x38, 0x69,
	0x9d, 0xd3, 0x0d, 0x4c, 0xcb, 0x6f, 0x17, 0xa3, 0x9f, 0x03, 0x38, 0xe8, 0x4e, 0xd0, 0x23, 0x18,
	0xc6, 0x3c, 0x5c, 0xd0, 0x09, 0x9b, 0x32, 0xea, 0x18, 0x77, 0xd0, 0x10, 0xf6, 0x63, 0xfe, 0x91,
	0xfb, 0x67, 0xdc, 0xe8, 0xa1, 0x13, 0x38, 0x62, 0x3c, 0x21, 0x2e, 0x73, 0xc4, 0x94, 0x51, 0xd7,
	0x11, 0x09, 0x71, 0x63, 0x1a, 0x0a, 0xc6, 0x85, 0x43, 0x22, 0x6a, 0xf4, 0xd1, 0x1b, 0x78, 0xb9,
	0x8b, 0x10, 0x11, 0xf3, 0xa8, 0x31, 0x40, 0x87, 0xf0, 0xb8, 0xc5, 0xc2, 0x28, 0x60, 0x7c, 0xa6,
	0xfd, 0xbb, 0xe8, 0x15, 0xbc, 0xb8, 0x61, 0xd0, 0x68, 0xc2, 0x63, 0x93, 0xc0, 0x0f, 0x8d, 0x3d,
	0xf4, 0x1a, 0x4e, 0xb6, 0x42, 0x21, 0x9d, 0xf8, 0xdc, 0x09, 0x8d, 0x21, 0x7a, 0x07, 0x6f, 0x6f,
	0xa3, 0xc4, 0x19, 0x8b, 0xe6, 0xc2, 0x9f, 0x4e, 0x43, 0x1a, 0x19, 0xf7, 0xd1, 0x73, 0x78, 0x42,
	0x49, 0xe0, 0x32, 0x1a, 0x88, 0x68, 0x4e, 0xb8, 0xf0, 0x18, 0x67, 0x5e, 0xec, 0xe9, 0xbb, 0xed,
	0xa3, 0x67, 0x70, 0xe8, 0x92, 0xa8, 0x1b, 0x92, 0x4f, 0xff, 0x87, 0xf7, 0xd0, 0x29, 0x98, 0xcd,
	0xfa, 0x80, 0xf0, 0x19, 0xbd, 0x66, 0x8a, 0x6d, 0xd2, 0x01, 0x7a, 0x0f, 0xf8, 0x06, 0x89, 0x70,
	0xe7, 0x1a, 0x18, 0x8a, 0xb1, 0x1f, 0xcd, 0x05, 0x8f, 0x5d, 0xd7, 0x80, 0xf1, 0x9f, 0x1e, 0x8c,
	0xbe, 0xca, 0x0b, 0xbc, 0xbb, 0x06, 0xe3, 0x87, 0xdd, 0x9b, 0x2e, 0x36, 0x0f, 0xbf, 0xe8, 0x7d,
	0x76, 0xae, 0x8c, 0x42, 0x9e, 0xa7, 0x65, 0x81, 0x65, 0x55, 0x98, 0x45, 0x5e, 0x36, 0xb5, 0x68,
	0x9b, 0xb7, 0x5e, 0xa9, 0x6d, 0x45, 0xfc, 0xa0, 0x3f, 0xdf, 0xfb, 0x83, 0x19, 0x21, 0x3f, 0xfa,
	0xc7, 0x33, 0xbd, 0x8c, 0x64, 0x0a, 0xeb, 0xb8, 0x49, 0x89, 0x85, 0x9b, 0x5f, 0xaa, 0xdf, 0x2d,
	0xb0, 0x24, 0x99, 0x5a, 0x76, 0xc0, 0x32, 0xb1, 0x96, 0x1a, 0xf8, 0xdb, 0x1f, 0xe9, 0x53, 0xdb,
	0x26, 0x99, 0xb2, 0xed, 0x0e, 0xb1, 0xed, 0xc4, 0xb2, 0x6d, 0x0d, 0x7d, 0xd9, 0x6b, 0x6e, 0x77,
	0xfa, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x90, 0x80, 0x3b, 0x25, 0x03, 0x00, 0x00,
}

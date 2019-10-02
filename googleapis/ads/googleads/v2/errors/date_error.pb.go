// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/date_error.proto

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
	return fileDescriptor_a4bbec016e5e2e3d, []int{0, 0}
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
	return fileDescriptor_a4bbec016e5e2e3d, []int{0}
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
	proto.RegisterEnum("google.ads.googleads.v2.errors.DateErrorEnum_DateError", DateErrorEnum_DateError_name, DateErrorEnum_DateError_value)
	proto.RegisterType((*DateErrorEnum)(nil), "google.ads.googleads.v2.errors.DateErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/date_error.proto", fileDescriptor_a4bbec016e5e2e3d)
}

var fileDescriptor_a4bbec016e5e2e3d = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x69, 0x8b, 0x36, 0xf6, 0x96, 0x3f, 0x91, 0x39, 0x4c, 0xc0, 0x18, 0xa3, 0xc0, 0x0d,
	0x39, 0x28, 0xbb, 0x85, 0x93, 0xdb, 0xb8, 0xad, 0x45, 0xe2, 0x54, 0x89, 0x93, 0x21, 0x54, 0xc9,
	0x0a, 0x24, 0x8a, 0x2a, 0x6d, 0x71, 0x15, 0x87, 0x7d, 0x20, 0x8e, 0x48, 0x7c, 0x06, 0xee, 0x7c,
	0x13, 0xf8, 0x14, 0xa8, 0xf1, 0x1a, 0x34, 0x69, 0x1d, 0xa7, 0x3c, 0xf2, 0xfb, 0xfb, 0xbd, 0xb1,
	0xe4, 0x07, 0xec, 0x52, 0xa9, 0xf2, 0xbc, 0xb0, 0xb3, 0x5c, 0x5f, 0xc5, 0x4d, 0xba, 0x74, 0xec,
	0xa2, 0xae, 0x55, 0xad, 0xed, 0x3c, 0x6b, 0x0a, 0xd9, 0x66, 0xbc, 0xae, 0x55, 0xa3, 0xd0, 0xb1,
	0xa1, 0x70, 0x96, 0x6b, 0xdc, 0x09, 0xf8, 0xd2, 0xc1, 0x46, 0x78, 0x7a, 0xb4, 0x5d, 0xb8, 0x5e,
	0xd9, 0x59, 0x55, 0xa9, 0x26, 0x6b, 0x56, 0xaa, 0xd2, 0xc6, 0x1e, 0xfd, 0x1c, 0xc0, 0x03, 0x2f,
	0x6b, 0x0a, 0xba, 0x81, 0x69, 0xf5, 0xf5, 0x62, 0xf4, 0x63, 0x00, 0x07, 0xdd, 0x09, 0x7a, 0x04,
	0xc3, 0x84, 0xc7, 0x0b, 0x3a, 0x61, 0x53, 0x46, 0x3d, 0xeb, 0x0e, 0x1a, 0xc2, 0x7e, 0xc2, 0x3f,
	0xf0, 0xf0, 0x8c, 0x5b, 0x3d, 0x74, 0x02, 0x47, 0x8c, 0xa7, 0xc4, 0x67, 0x9e, 0x9c, 0x32, 0xea,
	0x7b, 0x32, 0x25, 0x7e, 0x42, 0x63, 0xc9, 0xb8, 0xf4, 0x88, 0xa0, 0x56, 0x1f, 0xbd, 0x81, 0x97,
	0xb7, 0x11, 0x52, 0xb0, 0x80, 0x5a, 0x03, 0x74, 0x08, 0x8f, 0xb7, 0x58, 0x2c, 0x22, 0xc6, 0x67,
	0xc6, 0xbf, 0x8b, 0x5e, 0xc1, 0x8b, 0x1b, 0x06, 0xad, 0x26, 0x03, 0x36, 0x89, 0xc2, 0xd8, 0xda,
	0x43, 0xaf, 0xe1, 0x64, 0x27, 0x14, 0xd3, 0x49, 0xc8, 0xbd, 0xd8, 0x1a, 0xa2, 0x77, 0xf0, 0xf6,
	0x7f, 0x94, 0x3c, 0x63, 0x62, 0x2e, 0xc3, 0xe9, 0x34, 0xa6, 0xc2, 0xba, 0x8f, 0x9e, 0xc3, 0x13,
	0x4a, 0x22, 0x9f, 0xd1, 0x48, 0x8a, 0x39, 0xe1, 0x32, 0x60, 0x9c, 0x05, 0x49, 0x60, 0xee, 0xb6,
	0x8f, 0x9e, 0xc1, 0xa1, 0x4f, 0x44, 0x37, 0x24, 0x1f, 0xff, 0x0d, 0xef, 0xa1, 0x53, 0xb0, 0xdb,
	0xf5, 0x11, 0xe1, 0x33, 0x7a, 0xcd, 0x94, 0xbb, 0xa4, 0x03, 0xe4, 0x00, 0xbe, 0x41, 0x22, 0xdc,
	0xbb, 0x06, 0xc6, 0x72, 0x1c, 0x8a, 0xb9, 0xe4, 0x89, 0xef, 0x5b, 0x30, 0xfe, 0xdd, 0x83, 0xd1,
	0x17, 0x75, 0x81, 0x6f, 0xaf, 0xc1, 0xf8, 0x61, 0xf7, 0xa6, 0x8b, 0xcd, 0xc3, 0x2f, 0x7a, 0x9f,
	0xbc, 0x2b, 0xa3, 0x54, 0xe7, 0x59, 0x55, 0x62, 0x55, 0x97, 0x76, 0x59, 0x54, 0x6d, 0x2d, 0xb6,
	0xcd, 0x5b, 0xaf, 0xf4, 0xae, 0x22, 0xbe, 0x37, 0x9f, 0x6f, 0xfd, 0xc1, 0x8c, 0x90, 0xef, 0xfd,
	0xe3, 0x99, 0x59, 0x46, 0x72, 0x8d, 0x4d, 0xdc, 0xa4, 0xd4, 0xc1, 0xed, 0x2f, 0xf5, 0xaf, 0x2d,
	0xb0, 0x24, 0xb9, 0x5e, 0x76, 0xc0, 0x32, 0x75, 0x96, 0x06, 0xf8, 0xd3, 0x1f, 0x99, 0x53, 0xd7,
	0x25, 0xb9, 0x76, 0xdd, 0x0e, 0x71, 0xdd, 0xd4, 0x71, 0x5d, 0x03, 0x7d, 0xde, 0x6b, 0x6f, 0x77,
	0xfa, 0x37, 0x00, 0x00, 0xff, 0xff, 0x98, 0x9e, 0x2a, 0xca, 0x25, 0x03, 0x00, 0x00,
}

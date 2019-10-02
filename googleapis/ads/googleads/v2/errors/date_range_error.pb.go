// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/date_range_error.proto

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

// Enum describing possible date range errors.
type DateRangeErrorEnum_DateRangeError int32

const (
	// Enum unspecified.
	DateRangeErrorEnum_UNSPECIFIED DateRangeErrorEnum_DateRangeError = 0
	// The received error code is not known in this version.
	DateRangeErrorEnum_UNKNOWN DateRangeErrorEnum_DateRangeError = 1
	// Invalid date.
	DateRangeErrorEnum_INVALID_DATE DateRangeErrorEnum_DateRangeError = 2
	// The start date was after the end date.
	DateRangeErrorEnum_START_DATE_AFTER_END_DATE DateRangeErrorEnum_DateRangeError = 3
	// Cannot set date to past time
	DateRangeErrorEnum_CANNOT_SET_DATE_TO_PAST DateRangeErrorEnum_DateRangeError = 4
	// A date was used that is past the system "last" date.
	DateRangeErrorEnum_AFTER_MAXIMUM_ALLOWABLE_DATE DateRangeErrorEnum_DateRangeError = 5
	// Trying to change start date on a resource that has started.
	DateRangeErrorEnum_CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED DateRangeErrorEnum_DateRangeError = 6
)

var DateRangeErrorEnum_DateRangeError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "INVALID_DATE",
	3: "START_DATE_AFTER_END_DATE",
	4: "CANNOT_SET_DATE_TO_PAST",
	5: "AFTER_MAXIMUM_ALLOWABLE_DATE",
	6: "CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED",
}

var DateRangeErrorEnum_DateRangeError_value = map[string]int32{
	"UNSPECIFIED":                  0,
	"UNKNOWN":                      1,
	"INVALID_DATE":                 2,
	"START_DATE_AFTER_END_DATE":    3,
	"CANNOT_SET_DATE_TO_PAST":      4,
	"AFTER_MAXIMUM_ALLOWABLE_DATE": 5,
	"CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED": 6,
}

func (x DateRangeErrorEnum_DateRangeError) String() string {
	return proto.EnumName(DateRangeErrorEnum_DateRangeError_name, int32(x))
}

func (DateRangeErrorEnum_DateRangeError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_64ccd3dfacbc339c, []int{0, 0}
}

// Container for enum describing possible date range errors.
type DateRangeErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DateRangeErrorEnum) Reset()         { *m = DateRangeErrorEnum{} }
func (m *DateRangeErrorEnum) String() string { return proto.CompactTextString(m) }
func (*DateRangeErrorEnum) ProtoMessage()    {}
func (*DateRangeErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_64ccd3dfacbc339c, []int{0}
}

func (m *DateRangeErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DateRangeErrorEnum.Unmarshal(m, b)
}
func (m *DateRangeErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DateRangeErrorEnum.Marshal(b, m, deterministic)
}
func (m *DateRangeErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateRangeErrorEnum.Merge(m, src)
}
func (m *DateRangeErrorEnum) XXX_Size() int {
	return xxx_messageInfo_DateRangeErrorEnum.Size(m)
}
func (m *DateRangeErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_DateRangeErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_DateRangeErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.DateRangeErrorEnum_DateRangeError", DateRangeErrorEnum_DateRangeError_name, DateRangeErrorEnum_DateRangeError_value)
	proto.RegisterType((*DateRangeErrorEnum)(nil), "google.ads.googleads.v2.errors.DateRangeErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/date_range_error.proto", fileDescriptor_64ccd3dfacbc339c)
}

var fileDescriptor_64ccd3dfacbc339c = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x8e, 0x95, 0x30,
	0x14, 0x86, 0xe5, 0x8e, 0x8e, 0x49, 0xc7, 0x28, 0xa9, 0x0b, 0xa3, 0x8e, 0x13, 0xc3, 0xd6, 0xa4,
	0x24, 0x18, 0x37, 0x75, 0x75, 0xee, 0xb4, 0x4c, 0x88, 0x50, 0x08, 0x70, 0x19, 0x35, 0x24, 0x4d,
	0x15, 0x42, 0x6e, 0x32, 0x43, 0x6f, 0x28, 0xce, 0x03, 0xb9, 0xf4, 0x51, 0xdc, 0xf8, 0x16, 0x2e,
	0x5c, 0xf8, 0x0c, 0x06, 0xca, 0xbd, 0x71, 0x16, 0xba, 0xe2, 0xe7, 0xef, 0xff, 0x9d, 0xd3, 0x9e,
	0x83, 0xde, 0x74, 0x5a, 0x77, 0x57, 0xad, 0xaf, 0x1a, 0xe3, 0x5b, 0x39, 0xa9, 0x9b, 0xc0, 0x6f,
	0x87, 0x41, 0x0f, 0xc6, 0x6f, 0xd4, 0xd8, 0xca, 0x41, 0xf5, 0x5d, 0x2b, 0x67, 0x87, 0xec, 0x06,
	0x3d, 0x6a, 0x7c, 0x66, 0xb3, 0x44, 0x35, 0x86, 0x1c, 0x30, 0x72, 0x13, 0x10, 0x8b, 0x3d, 0x3b,
	0xdd, 0x97, 0xdd, 0x6d, 0x7d, 0xd5, 0xf7, 0x7a, 0x54, 0xe3, 0x56, 0xf7, 0xc6, 0xd2, 0xde, 0x4f,
	0x07, 0x61, 0xa6, 0xc6, 0x36, 0x9f, 0xea, 0xf2, 0x89, 0xe0, 0xfd, 0x97, 0x6b, 0xef, 0x87, 0x83,
	0x1e, 0xde, 0xb6, 0xf1, 0x23, 0x74, 0xb2, 0x11, 0x45, 0xc6, 0xcf, 0xa3, 0x30, 0xe2, 0xcc, 0xbd,
	0x83, 0x4f, 0xd0, 0xfd, 0x8d, 0x78, 0x27, 0xd2, 0x4b, 0xe1, 0x3a, 0xd8, 0x45, 0x0f, 0x22, 0x51,
	0x41, 0x1c, 0x31, 0xc9, 0xa0, 0xe4, 0xee, 0x0a, 0xbf, 0x40, 0x4f, 0x8b, 0x12, 0xf2, 0x72, 0xfe,
	0x97, 0x10, 0x96, 0x3c, 0x97, 0x5c, 0x2c, 0xc7, 0x47, 0xf8, 0x39, 0x7a, 0x72, 0x0e, 0x42, 0xa4,
	0xa5, 0x2c, 0xf8, 0x92, 0x29, 0x53, 0x99, 0x41, 0x51, 0xba, 0x77, 0xf1, 0x4b, 0x74, 0x6a, 0x81,
	0x04, 0xde, 0x47, 0xc9, 0x26, 0x91, 0x10, 0xc7, 0xe9, 0x25, 0xac, 0x63, 0x6e, 0xf1, 0x7b, 0xd8,
	0x47, 0xaf, 0x16, 0x3c, 0x49, 0x59, 0x14, 0x7e, 0x90, 0x7f, 0xf5, 0x8a, 0x42, 0x09, 0x71, 0xce,
	0x81, 0x2d, 0x2e, 0x67, 0xee, 0xf1, 0xfa, 0xb7, 0x83, 0xbc, 0xcf, 0xfa, 0x9a, 0xfc, 0x7f, 0x5a,
	0xeb, 0xc7, 0xb7, 0x5f, 0x9d, 0x4d, 0x43, 0xca, 0x9c, 0x8f, 0x6c, 0xc1, 0x3a, 0x7d, 0xa5, 0xfa,
	0x8e, 0xe8, 0xa1, 0xf3, 0xbb, 0xb6, 0x9f, 0x47, 0xb8, 0xdf, 0xd5, 0x6e, 0x6b, 0xfe, 0xb5, 0xba,
	0xb7, 0xf6, 0xf3, 0x75, 0x75, 0x74, 0x01, 0xf0, 0x6d, 0x75, 0x76, 0x61, 0x8b, 0x41, 0x63, 0x88,
	0x95, 0x93, 0xaa, 0x02, 0x32, 0xb7, 0x34, 0xdf, 0xf7, 0x81, 0x1a, 0x1a, 0x53, 0x1f, 0x02, 0x75,
	0x15, 0xd4, 0x36, 0xf0, 0x6b, 0xe5, 0x59, 0x97, 0x52, 0x68, 0x0c, 0xa5, 0x87, 0x08, 0xa5, 0x55,
	0x40, 0xa9, 0x0d, 0x7d, 0x3a, 0x9e, 0x6f, 0xf7, 0xfa, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x27,
	0x03, 0x3f, 0xa4, 0x57, 0x02, 0x00, 0x00,
}

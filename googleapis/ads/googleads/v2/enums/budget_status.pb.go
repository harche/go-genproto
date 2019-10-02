// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/budget_status.proto

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

// Possible statuses of a Budget.
type BudgetStatusEnum_BudgetStatus int32

const (
	// Not specified.
	BudgetStatusEnum_UNSPECIFIED BudgetStatusEnum_BudgetStatus = 0
	// Used for return value only. Represents value unknown in this version.
	BudgetStatusEnum_UNKNOWN BudgetStatusEnum_BudgetStatus = 1
	// Budget is enabled.
	BudgetStatusEnum_ENABLED BudgetStatusEnum_BudgetStatus = 2
	// Budget is removed.
	BudgetStatusEnum_REMOVED BudgetStatusEnum_BudgetStatus = 3
)

var BudgetStatusEnum_BudgetStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "REMOVED",
}

var BudgetStatusEnum_BudgetStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"REMOVED":     3,
}

func (x BudgetStatusEnum_BudgetStatus) String() string {
	return proto.EnumName(BudgetStatusEnum_BudgetStatus_name, int32(x))
}

func (BudgetStatusEnum_BudgetStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1f54f072cd971115, []int{0, 0}
}

// Message describing a Budget status
type BudgetStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BudgetStatusEnum) Reset()         { *m = BudgetStatusEnum{} }
func (m *BudgetStatusEnum) String() string { return proto.CompactTextString(m) }
func (*BudgetStatusEnum) ProtoMessage()    {}
func (*BudgetStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f54f072cd971115, []int{0}
}

func (m *BudgetStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BudgetStatusEnum.Unmarshal(m, b)
}
func (m *BudgetStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BudgetStatusEnum.Marshal(b, m, deterministic)
}
func (m *BudgetStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BudgetStatusEnum.Merge(m, src)
}
func (m *BudgetStatusEnum) XXX_Size() int {
	return xxx_messageInfo_BudgetStatusEnum.Size(m)
}
func (m *BudgetStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_BudgetStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_BudgetStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.BudgetStatusEnum_BudgetStatus", BudgetStatusEnum_BudgetStatus_name, BudgetStatusEnum_BudgetStatus_value)
	proto.RegisterType((*BudgetStatusEnum)(nil), "google.ads.googleads.v2.enums.BudgetStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/budget_status.proto", fileDescriptor_1f54f072cd971115)
}

var fileDescriptor_1f54f072cd971115 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x4f, 0x4b, 0xfb, 0x30,
	0x18, 0xfe, 0xad, 0x83, 0x9f, 0x90, 0x09, 0xd6, 0x1e, 0xc5, 0x1d, 0xb6, 0x0f, 0x90, 0x60, 0xbd,
	0xc5, 0x53, 0x6a, 0xb3, 0x31, 0xd4, 0xae, 0x38, 0x56, 0x61, 0x14, 0x24, 0x33, 0x25, 0x14, 0xd6,
	0xa4, 0x2c, 0xe9, 0x3e, 0x90, 0x47, 0x3f, 0x8a, 0x1f, 0xc4, 0x83, 0x9f, 0x42, 0x92, 0xae, 0x65,
	0x17, 0xbd, 0x94, 0xe7, 0x7d, 0x9f, 0x3f, 0x7d, 0xf2, 0x82, 0x1b, 0xa1, 0x94, 0xd8, 0x15, 0x88,
	0x71, 0x8d, 0x5a, 0x68, 0xd1, 0x21, 0x44, 0x85, 0x6c, 0x2a, 0x8d, 0xb6, 0x0d, 0x17, 0x85, 0x79,
	0xd5, 0x86, 0x99, 0x46, 0xc3, 0x7a, 0xaf, 0x8c, 0x0a, 0xc6, 0xad, 0x0e, 0x32, 0xae, 0x61, 0x6f,
	0x81, 0x87, 0x10, 0x3a, 0xcb, 0xd5, 0x75, 0x97, 0x58, 0x97, 0x88, 0x49, 0xa9, 0x0c, 0x33, 0xa5,
	0x92, 0x47, 0xf3, 0x74, 0x03, 0xfc, 0xc8, 0x65, 0xae, 0x5c, 0x24, 0x95, 0x4d, 0x35, 0x9d, 0x81,
	0xf3, 0xd3, 0x5d, 0x70, 0x01, 0x46, 0xeb, 0x64, 0x95, 0xd2, 0xfb, 0xc5, 0x6c, 0x41, 0x63, 0xff,
	0x5f, 0x30, 0x02, 0x67, 0xeb, 0xe4, 0x21, 0x59, 0xbe, 0x24, 0xfe, 0xc0, 0x0e, 0x34, 0x21, 0xd1,
	0x23, 0x8d, 0x7d, 0xcf, 0x0e, 0xcf, 0xf4, 0x69, 0x99, 0xd1, 0xd8, 0x1f, 0x46, 0x5f, 0x03, 0x30,
	0x79, 0x53, 0x15, 0xfc, 0xb3, 0x5f, 0x74, 0x79, 0xfa, 0xaf, 0xd4, 0x96, 0x4a, 0x07, 0x9b, 0xe8,
	0xe8, 0x11, 0x6a, 0xc7, 0xa4, 0x80, 0x6a, 0x2f, 0x90, 0x28, 0xa4, 0xab, 0xdc, 0x9d, 0xa5, 0x2e,
	0xf5, 0x2f, 0x57, 0xba, 0x73, 0xdf, 0x77, 0x6f, 0x38, 0x27, 0xe4, 0xc3, 0x1b, 0xcf, 0xdb, 0x28,
	0xc2, 0x35, 0x6c, 0xa1, 0x45, 0x59, 0x08, 0xed, 0x5b, 0xf5, 0x67, 0xc7, 0xe7, 0x84, 0xeb, 0xbc,
	0xe7, 0xf3, 0x2c, 0xcc, 0x1d, 0xff, 0xed, 0x4d, 0xda, 0x25, 0xc6, 0x84, 0x6b, 0x8c, 0x7b, 0x05,
	0xc6, 0x59, 0x88, 0xb1, 0xd3, 0x6c, 0xff, 0xbb, 0x62, 0xb7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x66, 0x92, 0xf0, 0x7e, 0xbd, 0x01, 0x00, 0x00,
}

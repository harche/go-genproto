// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/campaign_shared_set_error.proto

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

// Enum describing possible campaign shared set errors.
type CampaignSharedSetErrorEnum_CampaignSharedSetError int32

const (
	// Enum unspecified.
	CampaignSharedSetErrorEnum_UNSPECIFIED CampaignSharedSetErrorEnum_CampaignSharedSetError = 0
	// The received error code is not known in this version.
	CampaignSharedSetErrorEnum_UNKNOWN CampaignSharedSetErrorEnum_CampaignSharedSetError = 1
	// The shared set belongs to another customer and permission isn't granted.
	CampaignSharedSetErrorEnum_SHARED_SET_ACCESS_DENIED CampaignSharedSetErrorEnum_CampaignSharedSetError = 2
)

var CampaignSharedSetErrorEnum_CampaignSharedSetError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "SHARED_SET_ACCESS_DENIED",
}

var CampaignSharedSetErrorEnum_CampaignSharedSetError_value = map[string]int32{
	"UNSPECIFIED":              0,
	"UNKNOWN":                  1,
	"SHARED_SET_ACCESS_DENIED": 2,
}

func (x CampaignSharedSetErrorEnum_CampaignSharedSetError) String() string {
	return proto.EnumName(CampaignSharedSetErrorEnum_CampaignSharedSetError_name, int32(x))
}

func (CampaignSharedSetErrorEnum_CampaignSharedSetError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5901f69bd03ba308, []int{0, 0}
}

// Container for enum describing possible campaign shared set errors.
type CampaignSharedSetErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignSharedSetErrorEnum) Reset()         { *m = CampaignSharedSetErrorEnum{} }
func (m *CampaignSharedSetErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CampaignSharedSetErrorEnum) ProtoMessage()    {}
func (*CampaignSharedSetErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_5901f69bd03ba308, []int{0}
}

func (m *CampaignSharedSetErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignSharedSetErrorEnum.Unmarshal(m, b)
}
func (m *CampaignSharedSetErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignSharedSetErrorEnum.Marshal(b, m, deterministic)
}
func (m *CampaignSharedSetErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignSharedSetErrorEnum.Merge(m, src)
}
func (m *CampaignSharedSetErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CampaignSharedSetErrorEnum.Size(m)
}
func (m *CampaignSharedSetErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignSharedSetErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignSharedSetErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.CampaignSharedSetErrorEnum_CampaignSharedSetError", CampaignSharedSetErrorEnum_CampaignSharedSetError_name, CampaignSharedSetErrorEnum_CampaignSharedSetError_value)
	proto.RegisterType((*CampaignSharedSetErrorEnum)(nil), "google.ads.googleads.v2.errors.CampaignSharedSetErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/campaign_shared_set_error.proto", fileDescriptor_5901f69bd03ba308)
}

var fileDescriptor_5901f69bd03ba308 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xd1, 0x4a, 0xf3, 0x30,
	0x14, 0xc7, 0xbf, 0xf5, 0x03, 0x85, 0xec, 0xc2, 0xd1, 0x0b, 0x91, 0x39, 0x76, 0xd1, 0x07, 0x48,
	0xa1, 0xde, 0x45, 0x10, 0xb2, 0x36, 0xce, 0x21, 0xd4, 0x61, 0xb6, 0x09, 0x52, 0x08, 0x71, 0x09,
	0xb1, 0xb0, 0x25, 0x25, 0xa9, 0x7b, 0x20, 0x2f, 0x7d, 0x14, 0x1f, 0xc5, 0x07, 0xf0, 0x5a, 0xda,
	0x6c, 0xbd, 0x9a, 0x5e, 0xf5, 0x4f, 0xcf, 0xef, 0x77, 0xce, 0xc9, 0x01, 0x37, 0xca, 0x18, 0xb5,
	0x91, 0x31, 0x17, 0x2e, 0xf6, 0xb1, 0x49, 0xbb, 0x24, 0x96, 0xd6, 0x1a, 0xeb, 0xe2, 0x35, 0xdf,
	0x56, 0xbc, 0x54, 0x9a, 0xb9, 0x57, 0x6e, 0xa5, 0x60, 0x4e, 0xd6, 0xac, 0x2d, 0xc1, 0xca, 0x9a,
	0xda, 0x84, 0x63, 0x2f, 0x41, 0x2e, 0x1c, 0xec, 0x7c, 0xb8, 0x4b, 0xa0, 0xf7, 0x87, 0xa3, 0x43,
	0xff, 0xaa, 0x8c, 0xb9, 0xd6, 0xa6, 0xe6, 0x75, 0x69, 0xb4, 0xf3, 0x76, 0x64, 0xc1, 0x30, 0xdd,
	0x0f, 0xa0, 0x6d, 0x7f, 0x2a, 0x6b, 0xd2, 0x88, 0x44, 0xbf, 0x6d, 0xa3, 0x05, 0x38, 0x3f, 0x5e,
	0x0d, 0xcf, 0x40, 0x7f, 0x99, 0xd3, 0x39, 0x49, 0x67, 0xb7, 0x33, 0x92, 0x0d, 0xfe, 0x85, 0x7d,
	0x70, 0xba, 0xcc, 0xef, 0xf3, 0x87, 0xa7, 0x7c, 0xd0, 0x0b, 0x47, 0xe0, 0x82, 0xde, 0xe1, 0x47,
	0x92, 0x31, 0x4a, 0x16, 0x0c, 0xa7, 0x29, 0xa1, 0x94, 0x65, 0x24, 0x6f, 0xd0, 0x60, 0xf2, 0xdd,
	0x03, 0xd1, 0xda, 0x6c, 0xe1, 0xdf, 0x8b, 0x4f, 0x2e, 0x8f, 0x8f, 0x9e, 0x37, 0x7b, 0xcf, 0x7b,
	0xcf, 0xd9, 0x5e, 0x57, 0x66, 0xc3, 0xb5, 0x82, 0xc6, 0xaa, 0x58, 0x49, 0xdd, 0xbe, 0xea, 0x70,
	0xc7, 0xaa, 0x74, 0xbf, 0x9d, 0xf5, 0xda, 0x7f, 0xde, 0x83, 0xff, 0x53, 0x8c, 0x3f, 0x82, 0xf1,
	0xd4, 0x37, 0xc3, 0xc2, 0x41, 0x1f, 0x9b, 0xb4, 0x4a, 0x60, 0x3b, 0xd2, 0x7d, 0x1e, 0x80, 0x02,
	0x0b, 0x57, 0x74, 0x40, 0xb1, 0x4a, 0x0a, 0x0f, 0x7c, 0x05, 0x91, 0xff, 0x8b, 0x10, 0x16, 0x0e,
	0xa1, 0x0e, 0x41, 0x68, 0x95, 0x20, 0xe4, 0xa1, 0x97, 0x93, 0x76, 0xbb, 0xab, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x52, 0xf0, 0x7d, 0x5d, 0xf3, 0x01, 0x00, 0x00,
}

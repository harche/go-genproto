// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/common/dates.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A date range.
type DateRange struct {
	// The start date, in yyyy-mm-dd format.
	StartDate *wrappers.StringValue `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// The end date, in yyyy-mm-dd format.
	EndDate              *wrappers.StringValue `protobuf:"bytes,2,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DateRange) Reset()         { *m = DateRange{} }
func (m *DateRange) String() string { return proto.CompactTextString(m) }
func (*DateRange) ProtoMessage()    {}
func (*DateRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_141c54e845fb6763, []int{0}
}

func (m *DateRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DateRange.Unmarshal(m, b)
}
func (m *DateRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DateRange.Marshal(b, m, deterministic)
}
func (m *DateRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateRange.Merge(m, src)
}
func (m *DateRange) XXX_Size() int {
	return xxx_messageInfo_DateRange.Size(m)
}
func (m *DateRange) XXX_DiscardUnknown() {
	xxx_messageInfo_DateRange.DiscardUnknown(m)
}

var xxx_messageInfo_DateRange proto.InternalMessageInfo

func (m *DateRange) GetStartDate() *wrappers.StringValue {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *DateRange) GetEndDate() *wrappers.StringValue {
	if m != nil {
		return m.EndDate
	}
	return nil
}

func init() {
	proto.RegisterType((*DateRange)(nil), "google.ads.googleads.v1.common.DateRange")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/common/dates.proto", fileDescriptor_141c54e845fb6763)
}

var fileDescriptor_141c54e845fb6763 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x69, 0x05, 0x75, 0xf1, 0xd6, 0x93, 0x8c, 0x31, 0xa4, 0x27, 0xf1, 0x90, 0x50, 0x3d,
	0x08, 0xd9, 0xa9, 0x73, 0xb0, 0xeb, 0x98, 0xd0, 0x83, 0x14, 0xe4, 0xdd, 0x12, 0x43, 0xa1, 0xcd,
	0x5b, 0x92, 0x6c, 0x9e, 0xfd, 0x2a, 0x1e, 0xfd, 0x28, 0x7e, 0x0f, 0x2f, 0x7e, 0x0a, 0x49, 0xd3,
	0xf6, 0xa6, 0x78, 0xca, 0x93, 0xe4, 0xf7, 0x3c, 0xef, 0x1f, 0x72, 0xa3, 0x10, 0x55, 0x2d, 0x19,
	0x08, 0xcb, 0x82, 0xf4, 0xea, 0x98, 0xb1, 0x3d, 0x36, 0x0d, 0x6a, 0x26, 0xc0, 0x49, 0x4b, 0x5b,
	0x83, 0x0e, 0x93, 0x79, 0x00, 0x28, 0x08, 0x4b, 0x47, 0x96, 0x1e, 0x33, 0x1a, 0xd8, 0xe9, 0x6c,
	0xc8, 0x6a, 0x2b, 0x06, 0x5a, 0xa3, 0x03, 0x57, 0xa1, 0xee, 0xdd, 0xd3, 0xde, 0xcd, 0xba, 0xdb,
	0xee, 0xf0, 0xc2, 0x5e, 0x0d, 0xb4, 0xad, 0x34, 0xfd, 0x7f, 0xfa, 0x16, 0x91, 0xc9, 0x0a, 0x9c,
	0xdc, 0x82, 0x56, 0x32, 0x59, 0x10, 0x62, 0x1d, 0x18, 0xf7, 0xec, 0x1b, 0xb8, 0x8c, 0xae, 0xa2,
	0xeb, 0x8b, 0xdb, 0x59, 0x5f, 0x95, 0x0e, 0x11, 0xf4, 0xd1, 0x99, 0x4a, 0xab, 0x02, 0xea, 0x83,
	0xdc, 0x4e, 0x3a, 0xde, 0x27, 0x24, 0xf7, 0xe4, 0x5c, 0x6a, 0x11, 0xac, 0xf1, 0x3f, 0xac, 0x67,
	0x52, 0x0b, 0x6f, 0x5c, 0x7e, 0x45, 0x24, 0xdd, 0x63, 0x43, 0xff, 0x1e, 0x74, 0x49, 0x3c, 0x6c,
	0x37, 0x3e, 0x6a, 0x13, 0x3d, 0xad, 0x7a, 0x5a, 0x61, 0x0d, 0x5a, 0x51, 0x34, 0x8a, 0x29, 0xa9,
	0xbb, 0x42, 0xc3, 0x4a, 0xdb, 0xca, 0xfe, 0xb6, 0xe1, 0x45, 0x38, 0xde, 0xe3, 0x93, 0x75, 0x9e,
	0x7f, 0xc4, 0xf3, 0x75, 0x08, 0xcb, 0x85, 0xa5, 0x41, 0x7a, 0x55, 0x64, 0xf4, 0xa1, 0xc3, 0x3e,
	0x07, 0xa0, 0xcc, 0x85, 0x2d, 0x47, 0xa0, 0x2c, 0xb2, 0x32, 0x00, 0xdf, 0x71, 0x1a, 0x5e, 0x39,
	0xcf, 0x85, 0xe5, 0x7c, 0x44, 0x38, 0x2f, 0x32, 0xce, 0x03, 0xb4, 0x3b, 0xed, 0xba, 0xbb, 0xfb,
	0x09, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xeb, 0x9e, 0x78, 0xfe, 0x01, 0x00, 0x00,
}

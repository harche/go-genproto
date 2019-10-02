// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/common/click_location.proto

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

// Location criteria associated with a click.
type ClickLocation struct {
	// The city location criterion associated with the impression.
	City *wrappers.StringValue `protobuf:"bytes,1,opt,name=city,proto3" json:"city,omitempty"`
	// The country location criterion associated with the impression.
	Country *wrappers.StringValue `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	// The metro location criterion associated with the impression.
	Metro *wrappers.StringValue `protobuf:"bytes,3,opt,name=metro,proto3" json:"metro,omitempty"`
	// The most specific location criterion associated with the impression.
	MostSpecific *wrappers.StringValue `protobuf:"bytes,4,opt,name=most_specific,json=mostSpecific,proto3" json:"most_specific,omitempty"`
	// The region location criterion associated with the impression.
	Region               *wrappers.StringValue `protobuf:"bytes,5,opt,name=region,proto3" json:"region,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ClickLocation) Reset()         { *m = ClickLocation{} }
func (m *ClickLocation) String() string { return proto.CompactTextString(m) }
func (*ClickLocation) ProtoMessage()    {}
func (*ClickLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1caa69f1bff59a2, []int{0}
}

func (m *ClickLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClickLocation.Unmarshal(m, b)
}
func (m *ClickLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClickLocation.Marshal(b, m, deterministic)
}
func (m *ClickLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClickLocation.Merge(m, src)
}
func (m *ClickLocation) XXX_Size() int {
	return xxx_messageInfo_ClickLocation.Size(m)
}
func (m *ClickLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_ClickLocation.DiscardUnknown(m)
}

var xxx_messageInfo_ClickLocation proto.InternalMessageInfo

func (m *ClickLocation) GetCity() *wrappers.StringValue {
	if m != nil {
		return m.City
	}
	return nil
}

func (m *ClickLocation) GetCountry() *wrappers.StringValue {
	if m != nil {
		return m.Country
	}
	return nil
}

func (m *ClickLocation) GetMetro() *wrappers.StringValue {
	if m != nil {
		return m.Metro
	}
	return nil
}

func (m *ClickLocation) GetMostSpecific() *wrappers.StringValue {
	if m != nil {
		return m.MostSpecific
	}
	return nil
}

func (m *ClickLocation) GetRegion() *wrappers.StringValue {
	if m != nil {
		return m.Region
	}
	return nil
}

func init() {
	proto.RegisterType((*ClickLocation)(nil), "google.ads.googleads.v2.common.ClickLocation")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/common/click_location.proto", fileDescriptor_a1caa69f1bff59a2)
}

var fileDescriptor_a1caa69f1bff59a2 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbd, 0x6a, 0xf3, 0x30,
	0x14, 0x86, 0xb1, 0xf3, 0xf3, 0x81, 0xbe, 0x66, 0xf1, 0x64, 0x42, 0x08, 0x25, 0x53, 0x27, 0xa9,
	0x38, 0xa5, 0x83, 0x3a, 0x39, 0x29, 0x64, 0xe9, 0x10, 0x1a, 0xf0, 0x50, 0x0c, 0xc1, 0x91, 0x1d,
	0x21, 0x6a, 0xeb, 0x18, 0x49, 0x49, 0xc9, 0xed, 0x74, 0xe8, 0xd0, 0x4b, 0xe9, 0xa5, 0x14, 0x7a,
	0x0f, 0xc5, 0x96, 0x1d, 0xe8, 0xd0, 0xe2, 0xc9, 0x07, 0xeb, 0x79, 0xde, 0xd7, 0x58, 0x07, 0xcd,
	0x39, 0x00, 0xcf, 0x33, 0x92, 0xa4, 0x9a, 0xd8, 0xb1, 0x9a, 0x8e, 0x01, 0x61, 0x50, 0x14, 0x20,
	0x09, 0xcb, 0x05, 0x7b, 0xde, 0xe6, 0xc0, 0x12, 0x23, 0x40, 0xe2, 0x52, 0x81, 0x01, 0x6f, 0x6a,
	0x49, 0x9c, 0xa4, 0x1a, 0x9f, 0x25, 0x7c, 0x0c, 0xb0, 0x95, 0xc6, 0xcd, 0x39, 0xa9, 0xe9, 0xdd,
	0x61, 0x4f, 0x5e, 0x54, 0x52, 0x96, 0x99, 0xd2, 0xd6, 0x1f, 0x4f, 0xda, 0xd2, 0x52, 0x90, 0x44,
	0x4a, 0x30, 0x75, 0x78, 0x73, 0x3a, 0x7b, 0x73, 0xd1, 0x68, 0x59, 0xd5, 0x3e, 0x34, 0xad, 0xde,
	0x35, 0xea, 0x33, 0x61, 0x4e, 0xbe, 0x73, 0xe9, 0x5c, 0xfd, 0x0f, 0x26, 0x4d, 0x27, 0x6e, 0xe3,
	0xf1, 0xc6, 0x28, 0x21, 0x79, 0x94, 0xe4, 0x87, 0xec, 0xb1, 0x26, 0xbd, 0x5b, 0xf4, 0x8f, 0xc1,
	0x41, 0x1a, 0x75, 0xf2, 0xdd, 0x0e, 0x52, 0x0b, 0x7b, 0x01, 0x1a, 0x14, 0x99, 0x51, 0xe0, 0xf7,
	0x3a, 0x58, 0x16, 0xf5, 0x42, 0x34, 0x2a, 0x40, 0x9b, 0xad, 0x2e, 0x33, 0x26, 0xf6, 0x82, 0xf9,
	0xfd, 0x0e, 0xee, 0x45, 0xa5, 0x6c, 0x1a, 0xc3, 0xbb, 0x41, 0x43, 0x95, 0x71, 0x01, 0xd2, 0x1f,
	0x74, 0x70, 0x1b, 0x76, 0xf1, 0xe5, 0xa0, 0x19, 0x83, 0x02, 0xff, 0x7d, 0x1b, 0x0b, 0xef, 0xc7,
	0xcf, 0x5c, 0x57, 0x89, 0x6b, 0xe7, 0xe9, 0xbe, 0xb1, 0x38, 0xe4, 0x89, 0xe4, 0x18, 0x14, 0x27,
	0x3c, 0x93, 0x75, 0x5f, 0xbb, 0x08, 0xa5, 0xd0, 0xbf, 0xed, 0xc5, 0x9d, 0x7d, 0xbc, 0xba, 0xbd,
	0x55, 0x18, 0xbe, 0xbb, 0xd3, 0x95, 0x0d, 0x0b, 0x53, 0x8d, 0xed, 0x58, 0x4d, 0x51, 0x80, 0x97,
	0x35, 0xf6, 0xd1, 0x02, 0x71, 0x98, 0xea, 0xf8, 0x0c, 0xc4, 0x51, 0x10, 0x5b, 0xe0, 0xd3, 0x9d,
	0xd9, 0xb7, 0x94, 0x86, 0xa9, 0xa6, 0xf4, 0x8c, 0x50, 0x1a, 0x05, 0x94, 0x5a, 0x68, 0x37, 0xac,
	0xbf, 0x6e, 0xfe, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x82, 0xc5, 0x77, 0x80, 0xb4, 0x02, 0x00, 0x00,
}

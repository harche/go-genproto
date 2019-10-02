// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/reach_plan_ad_length.proto

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

// Possible ad length values.
type ReachPlanAdLengthEnum_ReachPlanAdLength int32

const (
	// Not specified.
	ReachPlanAdLengthEnum_UNSPECIFIED ReachPlanAdLengthEnum_ReachPlanAdLength = 0
	// The value is unknown in this version.
	ReachPlanAdLengthEnum_UNKNOWN ReachPlanAdLengthEnum_ReachPlanAdLength = 1
	// 6 seconds long ad.
	ReachPlanAdLengthEnum_SIX_SECONDS ReachPlanAdLengthEnum_ReachPlanAdLength = 2
	// 15 or 20 seconds long ad.
	ReachPlanAdLengthEnum_FIFTEEN_OR_TWENTY_SECONDS ReachPlanAdLengthEnum_ReachPlanAdLength = 3
	// More than 20 seconds long ad.
	ReachPlanAdLengthEnum_TWENTY_SECONDS_OR_MORE ReachPlanAdLengthEnum_ReachPlanAdLength = 4
)

var ReachPlanAdLengthEnum_ReachPlanAdLength_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "SIX_SECONDS",
	3: "FIFTEEN_OR_TWENTY_SECONDS",
	4: "TWENTY_SECONDS_OR_MORE",
}

var ReachPlanAdLengthEnum_ReachPlanAdLength_value = map[string]int32{
	"UNSPECIFIED":               0,
	"UNKNOWN":                   1,
	"SIX_SECONDS":               2,
	"FIFTEEN_OR_TWENTY_SECONDS": 3,
	"TWENTY_SECONDS_OR_MORE":    4,
}

func (x ReachPlanAdLengthEnum_ReachPlanAdLength) String() string {
	return proto.EnumName(ReachPlanAdLengthEnum_ReachPlanAdLength_name, int32(x))
}

func (ReachPlanAdLengthEnum_ReachPlanAdLength) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8ff9071174e50cfb, []int{0, 0}
}

// Message describing length of a plannable video ad.
type ReachPlanAdLengthEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReachPlanAdLengthEnum) Reset()         { *m = ReachPlanAdLengthEnum{} }
func (m *ReachPlanAdLengthEnum) String() string { return proto.CompactTextString(m) }
func (*ReachPlanAdLengthEnum) ProtoMessage()    {}
func (*ReachPlanAdLengthEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ff9071174e50cfb, []int{0}
}

func (m *ReachPlanAdLengthEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReachPlanAdLengthEnum.Unmarshal(m, b)
}
func (m *ReachPlanAdLengthEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReachPlanAdLengthEnum.Marshal(b, m, deterministic)
}
func (m *ReachPlanAdLengthEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReachPlanAdLengthEnum.Merge(m, src)
}
func (m *ReachPlanAdLengthEnum) XXX_Size() int {
	return xxx_messageInfo_ReachPlanAdLengthEnum.Size(m)
}
func (m *ReachPlanAdLengthEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ReachPlanAdLengthEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ReachPlanAdLengthEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.ReachPlanAdLengthEnum_ReachPlanAdLength", ReachPlanAdLengthEnum_ReachPlanAdLength_name, ReachPlanAdLengthEnum_ReachPlanAdLength_value)
	proto.RegisterType((*ReachPlanAdLengthEnum)(nil), "google.ads.googleads.v2.enums.ReachPlanAdLengthEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/reach_plan_ad_length.proto", fileDescriptor_8ff9071174e50cfb)
}

var fileDescriptor_8ff9071174e50cfb = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x6a, 0xe3, 0x30,
	0x18, 0xc5, 0xc7, 0xce, 0x30, 0x03, 0xca, 0x62, 0x3c, 0x86, 0x06, 0x1a, 0x9a, 0x45, 0x72, 0x00,
	0x19, 0xdc, 0x4d, 0x51, 0x57, 0x4e, 0xa2, 0x04, 0xd3, 0xd6, 0x36, 0x71, 0xfe, 0xb4, 0xc5, 0x60,
	0xd4, 0xc8, 0x28, 0x01, 0x47, 0x32, 0x91, 0x93, 0x5d, 0xaf, 0xd1, 0x03, 0x74, 0xd9, 0xa3, 0xf4,
	0x28, 0xed, 0x25, 0x8a, 0xe4, 0xc6, 0x50, 0x42, 0xbb, 0x11, 0x0f, 0xfd, 0xde, 0x7b, 0x7c, 0xdf,
	0x07, 0x2e, 0x98, 0x10, 0x2c, 0xcf, 0x1c, 0x42, 0xa5, 0x53, 0x49, 0xa5, 0xf6, 0xae, 0x93, 0xf1,
	0xdd, 0x46, 0x3a, 0xdb, 0x8c, 0x2c, 0x57, 0x69, 0x91, 0x13, 0x9e, 0x12, 0x9a, 0xe6, 0x19, 0x67,
	0xe5, 0x0a, 0x16, 0x5b, 0x51, 0x0a, 0xbb, 0x53, 0xd9, 0x21, 0xa1, 0x12, 0xd6, 0x49, 0xb8, 0x77,
	0xa1, 0x4e, 0xb6, 0xcf, 0x0e, 0xc5, 0xc5, 0xda, 0x21, 0x9c, 0x8b, 0x92, 0x94, 0x6b, 0xc1, 0x65,
	0x15, 0xee, 0x3d, 0x19, 0xe0, 0x64, 0xa2, 0xba, 0xa3, 0x9c, 0x70, 0x8f, 0x5e, 0xeb, 0x62, 0xcc,
	0x77, 0x9b, 0xde, 0x23, 0xf8, 0x7f, 0x04, 0xec, 0x7f, 0xa0, 0x39, 0x0b, 0xe2, 0x08, 0x0f, 0xfc,
	0x91, 0x8f, 0x87, 0xd6, 0x2f, 0xbb, 0x09, 0xfe, 0xce, 0x82, 0xab, 0x20, 0x5c, 0x04, 0x96, 0xa1,
	0x68, 0xec, 0xdf, 0xa6, 0x31, 0x1e, 0x84, 0xc1, 0x30, 0xb6, 0x4c, 0xbb, 0x03, 0x4e, 0x47, 0xfe,
	0x68, 0x8a, 0x71, 0x90, 0x86, 0x93, 0x74, 0xba, 0xc0, 0xc1, 0xf4, 0xae, 0xc6, 0x0d, 0xbb, 0x0d,
	0x5a, 0x5f, 0xff, 0x94, 0xeb, 0x26, 0x9c, 0x60, 0xeb, 0x77, 0xff, 0xdd, 0x00, 0xdd, 0xa5, 0xd8,
	0xc0, 0x1f, 0x97, 0xeb, 0xb7, 0x8e, 0x46, 0x8c, 0xd4, 0x5a, 0x91, 0x71, 0xdf, 0xff, 0x0c, 0x32,
	0x91, 0x13, 0xce, 0xa0, 0xd8, 0x32, 0x87, 0x65, 0x5c, 0x2f, 0x7d, 0xb8, 0x6f, 0xb1, 0x96, 0xdf,
	0x9c, 0xfb, 0x52, 0xbf, 0xcf, 0x66, 0x63, 0xec, 0x79, 0x2f, 0x66, 0x67, 0x5c, 0x55, 0x79, 0x54,
	0xc2, 0x4a, 0x2a, 0x35, 0x77, 0xa1, 0xba, 0x93, 0x7c, 0x3d, 0xf0, 0xc4, 0xa3, 0x32, 0xa9, 0x79,
	0x32, 0x77, 0x13, 0xcd, 0xdf, 0xcc, 0x6e, 0xf5, 0x89, 0x90, 0x47, 0x25, 0x42, 0xb5, 0x03, 0xa1,
	0xb9, 0x8b, 0x90, 0xf6, 0x3c, 0xfc, 0xd1, 0x83, 0x9d, 0x7f, 0x04, 0x00, 0x00, 0xff, 0xff, 0x47,
	0x31, 0x6b, 0x3f, 0x06, 0x02, 0x00, 0x00,
}

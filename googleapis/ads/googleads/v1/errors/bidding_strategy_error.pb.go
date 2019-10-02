// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/bidding_strategy_error.proto

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

// Enum describing possible bidding strategy errors.
type BiddingStrategyErrorEnum_BiddingStrategyError int32

const (
	// Enum unspecified.
	BiddingStrategyErrorEnum_UNSPECIFIED BiddingStrategyErrorEnum_BiddingStrategyError = 0
	// The received error code is not known in this version.
	BiddingStrategyErrorEnum_UNKNOWN BiddingStrategyErrorEnum_BiddingStrategyError = 1
	// Each bidding strategy must have a unique name.
	BiddingStrategyErrorEnum_DUPLICATE_NAME BiddingStrategyErrorEnum_BiddingStrategyError = 2
	// Bidding strategy type is immutable.
	BiddingStrategyErrorEnum_CANNOT_CHANGE_BIDDING_STRATEGY_TYPE BiddingStrategyErrorEnum_BiddingStrategyError = 3
	// Only bidding strategies not linked to campaigns, adgroups or adgroup
	// criteria can be removed.
	BiddingStrategyErrorEnum_CANNOT_REMOVE_ASSOCIATED_STRATEGY BiddingStrategyErrorEnum_BiddingStrategyError = 4
	// The specified bidding strategy is not supported.
	BiddingStrategyErrorEnum_BIDDING_STRATEGY_NOT_SUPPORTED BiddingStrategyErrorEnum_BiddingStrategyError = 5
	// The bidding strategy is incompatible with the campaign's bidding
	// strategy goal type.
	BiddingStrategyErrorEnum_INCOMPATIBLE_BIDDING_STRATEGY_AND_BIDDING_STRATEGY_GOAL_TYPE BiddingStrategyErrorEnum_BiddingStrategyError = 6
)

var BiddingStrategyErrorEnum_BiddingStrategyError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "DUPLICATE_NAME",
	3: "CANNOT_CHANGE_BIDDING_STRATEGY_TYPE",
	4: "CANNOT_REMOVE_ASSOCIATED_STRATEGY",
	5: "BIDDING_STRATEGY_NOT_SUPPORTED",
	6: "INCOMPATIBLE_BIDDING_STRATEGY_AND_BIDDING_STRATEGY_GOAL_TYPE",
}

var BiddingStrategyErrorEnum_BiddingStrategyError_value = map[string]int32{
	"UNSPECIFIED":                         0,
	"UNKNOWN":                             1,
	"DUPLICATE_NAME":                      2,
	"CANNOT_CHANGE_BIDDING_STRATEGY_TYPE": 3,
	"CANNOT_REMOVE_ASSOCIATED_STRATEGY":   4,
	"BIDDING_STRATEGY_NOT_SUPPORTED":      5,
	"INCOMPATIBLE_BIDDING_STRATEGY_AND_BIDDING_STRATEGY_GOAL_TYPE": 6,
}

func (x BiddingStrategyErrorEnum_BiddingStrategyError) String() string {
	return proto.EnumName(BiddingStrategyErrorEnum_BiddingStrategyError_name, int32(x))
}

func (BiddingStrategyErrorEnum_BiddingStrategyError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4903c61cc975ddf7, []int{0, 0}
}

// Container for enum describing possible bidding strategy errors.
type BiddingStrategyErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BiddingStrategyErrorEnum) Reset()         { *m = BiddingStrategyErrorEnum{} }
func (m *BiddingStrategyErrorEnum) String() string { return proto.CompactTextString(m) }
func (*BiddingStrategyErrorEnum) ProtoMessage()    {}
func (*BiddingStrategyErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_4903c61cc975ddf7, []int{0}
}

func (m *BiddingStrategyErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BiddingStrategyErrorEnum.Unmarshal(m, b)
}
func (m *BiddingStrategyErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BiddingStrategyErrorEnum.Marshal(b, m, deterministic)
}
func (m *BiddingStrategyErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BiddingStrategyErrorEnum.Merge(m, src)
}
func (m *BiddingStrategyErrorEnum) XXX_Size() int {
	return xxx_messageInfo_BiddingStrategyErrorEnum.Size(m)
}
func (m *BiddingStrategyErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_BiddingStrategyErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_BiddingStrategyErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.BiddingStrategyErrorEnum_BiddingStrategyError", BiddingStrategyErrorEnum_BiddingStrategyError_name, BiddingStrategyErrorEnum_BiddingStrategyError_value)
	proto.RegisterType((*BiddingStrategyErrorEnum)(nil), "google.ads.googleads.v1.errors.BiddingStrategyErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/bidding_strategy_error.proto", fileDescriptor_4903c61cc975ddf7)
}

var fileDescriptor_4903c61cc975ddf7 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x8b, 0xd4, 0x30,
	0x1c, 0xc5, 0x9d, 0xae, 0xae, 0x90, 0x05, 0x2d, 0xc1, 0x83, 0x8a, 0x0c, 0x58, 0x11, 0x6f, 0x29,
	0xc5, 0x5b, 0xd6, 0x83, 0x69, 0x13, 0x6b, 0x70, 0x26, 0x0d, 0xd3, 0xcc, 0xc8, 0xca, 0x40, 0xe8,
	0xda, 0x21, 0x14, 0x76, 0x9b, 0xa1, 0xa9, 0x0b, 0x7e, 0x16, 0x6f, 0x1e, 0xfd, 0x28, 0x7e, 0x14,
	0xcf, 0x9e, 0x45, 0x3a, 0x99, 0xe9, 0x65, 0x47, 0x4f, 0x7d, 0xfc, 0xfb, 0x7e, 0x2f, 0x7f, 0xde,
	0x1f, 0x9c, 0x1b, 0x6b, 0xcd, 0xd5, 0x26, 0xae, 0x6a, 0x17, 0x7b, 0x39, 0xa8, 0x9b, 0x24, 0xde,
	0x74, 0x9d, 0xed, 0x5c, 0x7c, 0xd9, 0xd4, 0x75, 0xd3, 0x1a, 0xed, 0xfa, 0xae, 0xea, 0x37, 0xe6,
	0xab, 0xde, 0xcd, 0xd1, 0xb6, 0xb3, 0xbd, 0x85, 0x53, 0x4f, 0xa0, 0xaa, 0x76, 0x68, 0x84, 0xd1,
	0x4d, 0x82, 0x3c, 0xfc, 0xf4, 0xd9, 0x21, 0x7c, 0xdb, 0xc4, 0x55, 0xdb, 0xda, 0xbe, 0xea, 0x1b,
	0xdb, 0x3a, 0x4f, 0x47, 0xdf, 0x02, 0xf0, 0x38, 0xf5, 0xf1, 0xe5, 0x3e, 0x9d, 0x0d, 0x1c, 0x6b,
	0xbf, 0x5c, 0x47, 0x7f, 0x26, 0xe0, 0xd1, 0xb1, 0x9f, 0xf0, 0x21, 0x38, 0x5b, 0x8a, 0x52, 0xb2,
	0x8c, 0xbf, 0xe3, 0x8c, 0x86, 0x77, 0xe0, 0x19, 0xb8, 0xbf, 0x14, 0x1f, 0x44, 0xf1, 0x51, 0x84,
	0x13, 0x08, 0xc1, 0x03, 0xba, 0x94, 0x33, 0x9e, 0x11, 0xc5, 0xb4, 0x20, 0x73, 0x16, 0x06, 0xf0,
	0x15, 0x78, 0x91, 0x11, 0x21, 0x0a, 0xa5, 0xb3, 0xf7, 0x44, 0xe4, 0x4c, 0xa7, 0x9c, 0x52, 0x2e,
	0x72, 0x5d, 0xaa, 0x05, 0x51, 0x2c, 0xbf, 0xd0, 0xea, 0x42, 0xb2, 0xf0, 0x04, 0xbe, 0x04, 0xcf,
	0xf7, 0xc6, 0x05, 0x9b, 0x17, 0x2b, 0xa6, 0x49, 0x59, 0x16, 0x19, 0x27, 0x8a, 0xd1, 0xd1, 0x1b,
	0xde, 0x85, 0x11, 0x98, 0xde, 0x4a, 0x18, 0xa0, 0x72, 0x29, 0x65, 0xb1, 0x50, 0x8c, 0x86, 0xf7,
	0xe0, 0x5b, 0xf0, 0x86, 0x8b, 0xac, 0x98, 0x4b, 0xa2, 0x78, 0x3a, 0x3b, 0xf2, 0x24, 0x11, 0xf4,
	0xf6, 0x30, 0x2f, 0xc8, 0xcc, 0x2f, 0x73, 0x9a, 0xfe, 0x9e, 0x80, 0xe8, 0xb3, 0xbd, 0x46, 0xff,
	0xaf, 0x38, 0x7d, 0x72, 0xac, 0x24, 0x39, 0xf4, 0x2b, 0x27, 0x9f, 0xe8, 0x1e, 0x36, 0xf6, 0xaa,
	0x6a, 0x0d, 0xb2, 0x9d, 0x89, 0xcd, 0xa6, 0xdd, 0xb5, 0x7f, 0x38, 0xf6, 0xb6, 0x71, 0xff, 0xba,
	0xfd, 0xb9, 0xff, 0x7c, 0x0f, 0x4e, 0x72, 0x42, 0x7e, 0x04, 0xd3, 0xdc, 0x87, 0x91, 0xda, 0x21,
	0x2f, 0x07, 0xb5, 0x4a, 0xd0, 0xee, 0x49, 0xf7, 0xf3, 0x60, 0x58, 0x93, 0xda, 0xad, 0x47, 0xc3,
	0x7a, 0x95, 0xac, 0xbd, 0xe1, 0x57, 0x10, 0xf9, 0x29, 0xc6, 0xa4, 0x76, 0x18, 0x8f, 0x16, 0x8c,
	0x57, 0x09, 0xc6, 0xde, 0x74, 0x79, 0xba, 0xdb, 0xee, 0xf5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x22, 0xb7, 0x07, 0xda, 0x98, 0x02, 0x00, 0x00,
}

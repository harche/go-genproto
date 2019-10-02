// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/bidding_strategy_type.proto

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

// Enum describing possible bidding strategy types.
type BiddingStrategyTypeEnum_BiddingStrategyType int32

const (
	// Not specified.
	BiddingStrategyTypeEnum_UNSPECIFIED BiddingStrategyTypeEnum_BiddingStrategyType = 0
	// Used for return value only. Represents value unknown in this version.
	BiddingStrategyTypeEnum_UNKNOWN BiddingStrategyTypeEnum_BiddingStrategyType = 1
	// Commission is an automatic bidding strategy in which the advertiser pays
	// a certain portion of the conversion value.
	BiddingStrategyTypeEnum_COMMISSION BiddingStrategyTypeEnum_BiddingStrategyType = 16
	// Enhanced CPC is a bidding strategy that raises bids for clicks
	// that seem more likely to lead to a conversion and lowers
	// them for clicks where they seem less likely.
	BiddingStrategyTypeEnum_ENHANCED_CPC BiddingStrategyTypeEnum_BiddingStrategyType = 2
	// Manual click based bidding where user pays per click.
	BiddingStrategyTypeEnum_MANUAL_CPC BiddingStrategyTypeEnum_BiddingStrategyType = 3
	// Manual impression based bidding
	// where user pays per thousand impressions.
	BiddingStrategyTypeEnum_MANUAL_CPM BiddingStrategyTypeEnum_BiddingStrategyType = 4
	// A bidding strategy that pays a configurable amount per video view.
	BiddingStrategyTypeEnum_MANUAL_CPV BiddingStrategyTypeEnum_BiddingStrategyType = 13
	// A bidding strategy that automatically maximizes number of conversions
	// given a daily budget.
	BiddingStrategyTypeEnum_MAXIMIZE_CONVERSIONS BiddingStrategyTypeEnum_BiddingStrategyType = 10
	// An automated bidding strategy that automatically sets bids to maximize
	// revenue while spending your budget.
	BiddingStrategyTypeEnum_MAXIMIZE_CONVERSION_VALUE BiddingStrategyTypeEnum_BiddingStrategyType = 11
	// Page-One Promoted bidding scheme, which sets max cpc bids to
	// target impressions on page one or page one promoted slots on google.com.
	// This enum value is deprecated.
	BiddingStrategyTypeEnum_PAGE_ONE_PROMOTED BiddingStrategyTypeEnum_BiddingStrategyType = 5
	// Percent Cpc is bidding strategy where bids are a fraction of the
	// advertised price for some good or service.
	BiddingStrategyTypeEnum_PERCENT_CPC BiddingStrategyTypeEnum_BiddingStrategyType = 12
	// Target CPA is an automated bid strategy that sets bids
	// to help get as many conversions as possible
	// at the target cost-per-acquisition (CPA) you set.
	BiddingStrategyTypeEnum_TARGET_CPA BiddingStrategyTypeEnum_BiddingStrategyType = 6
	// Target CPM is an automated bid strategy that sets bids to help get
	// as many impressions as possible at the target cost per one thousand
	// impressions (CPM) you set.
	BiddingStrategyTypeEnum_TARGET_CPM BiddingStrategyTypeEnum_BiddingStrategyType = 14
	// An automated bidding strategy that sets bids so that a certain percentage
	// of search ads are shown at the top of the first page (or other targeted
	// location).
	BiddingStrategyTypeEnum_TARGET_IMPRESSION_SHARE BiddingStrategyTypeEnum_BiddingStrategyType = 15
	// Target Outrank Share is an automated bidding strategy that sets bids
	// based on the target fraction of auctions where the advertiser
	// should outrank a specific competitor.
	// This enum value is deprecated.
	BiddingStrategyTypeEnum_TARGET_OUTRANK_SHARE BiddingStrategyTypeEnum_BiddingStrategyType = 7
	// Target ROAS is an automated bidding strategy
	// that helps you maximize revenue while averaging
	// a specific target Return On Average Spend (ROAS).
	BiddingStrategyTypeEnum_TARGET_ROAS BiddingStrategyTypeEnum_BiddingStrategyType = 8
	// Target Spend is an automated bid strategy that sets your bids
	// to help get as many clicks as possible within your budget.
	BiddingStrategyTypeEnum_TARGET_SPEND BiddingStrategyTypeEnum_BiddingStrategyType = 9
)

var BiddingStrategyTypeEnum_BiddingStrategyType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	16: "COMMISSION",
	2:  "ENHANCED_CPC",
	3:  "MANUAL_CPC",
	4:  "MANUAL_CPM",
	13: "MANUAL_CPV",
	10: "MAXIMIZE_CONVERSIONS",
	11: "MAXIMIZE_CONVERSION_VALUE",
	5:  "PAGE_ONE_PROMOTED",
	12: "PERCENT_CPC",
	6:  "TARGET_CPA",
	14: "TARGET_CPM",
	15: "TARGET_IMPRESSION_SHARE",
	7:  "TARGET_OUTRANK_SHARE",
	8:  "TARGET_ROAS",
	9:  "TARGET_SPEND",
}

var BiddingStrategyTypeEnum_BiddingStrategyType_value = map[string]int32{
	"UNSPECIFIED":               0,
	"UNKNOWN":                   1,
	"COMMISSION":                16,
	"ENHANCED_CPC":              2,
	"MANUAL_CPC":                3,
	"MANUAL_CPM":                4,
	"MANUAL_CPV":                13,
	"MAXIMIZE_CONVERSIONS":      10,
	"MAXIMIZE_CONVERSION_VALUE": 11,
	"PAGE_ONE_PROMOTED":         5,
	"PERCENT_CPC":               12,
	"TARGET_CPA":                6,
	"TARGET_CPM":                14,
	"TARGET_IMPRESSION_SHARE":   15,
	"TARGET_OUTRANK_SHARE":      7,
	"TARGET_ROAS":               8,
	"TARGET_SPEND":              9,
}

func (x BiddingStrategyTypeEnum_BiddingStrategyType) String() string {
	return proto.EnumName(BiddingStrategyTypeEnum_BiddingStrategyType_name, int32(x))
}

func (BiddingStrategyTypeEnum_BiddingStrategyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_39550eb3746deb7d, []int{0, 0}
}

// Container for enum describing possible bidding strategy types.
type BiddingStrategyTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BiddingStrategyTypeEnum) Reset()         { *m = BiddingStrategyTypeEnum{} }
func (m *BiddingStrategyTypeEnum) String() string { return proto.CompactTextString(m) }
func (*BiddingStrategyTypeEnum) ProtoMessage()    {}
func (*BiddingStrategyTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_39550eb3746deb7d, []int{0}
}

func (m *BiddingStrategyTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BiddingStrategyTypeEnum.Unmarshal(m, b)
}
func (m *BiddingStrategyTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BiddingStrategyTypeEnum.Marshal(b, m, deterministic)
}
func (m *BiddingStrategyTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BiddingStrategyTypeEnum.Merge(m, src)
}
func (m *BiddingStrategyTypeEnum) XXX_Size() int {
	return xxx_messageInfo_BiddingStrategyTypeEnum.Size(m)
}
func (m *BiddingStrategyTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_BiddingStrategyTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_BiddingStrategyTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.BiddingStrategyTypeEnum_BiddingStrategyType", BiddingStrategyTypeEnum_BiddingStrategyType_name, BiddingStrategyTypeEnum_BiddingStrategyType_value)
	proto.RegisterType((*BiddingStrategyTypeEnum)(nil), "google.ads.googleads.v2.enums.BiddingStrategyTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/bidding_strategy_type.proto", fileDescriptor_39550eb3746deb7d)
}

var fileDescriptor_39550eb3746deb7d = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0x69, 0x06, 0x1b, 0xb8, 0x63, 0x33, 0x01, 0xb4, 0xf1, 0xd1, 0x8b, 0xed, 0x01, 0x1c,
	0xa9, 0x5c, 0x11, 0xae, 0xdc, 0xd4, 0x74, 0xd1, 0xe6, 0x0f, 0xe5, 0x0b, 0x34, 0x55, 0x8a, 0x32,
	0x12, 0x45, 0x91, 0xd6, 0x38, 0xaa, 0xb3, 0x49, 0x7d, 0x1d, 0x2e, 0x79, 0x14, 0x9e, 0x82, 0x6b,
	0x24, 0xde, 0x80, 0x0b, 0x64, 0x27, 0xad, 0x54, 0xa9, 0xec, 0x26, 0x3a, 0x7f, 0xff, 0xce, 0xf9,
	0xfb, 0xc4, 0xe7, 0x80, 0x8f, 0xa5, 0x94, 0xe5, 0x6d, 0xe1, 0x64, 0xb9, 0x72, 0xba, 0x50, 0x47,
	0xf7, 0x63, 0xa7, 0xa8, 0xef, 0x16, 0xca, 0xb9, 0xa9, 0xf2, 0xbc, 0xaa, 0xcb, 0x54, 0xb5, 0xcb,
	0xac, 0x2d, 0xca, 0x55, 0xda, 0xae, 0x9a, 0x02, 0x35, 0x4b, 0xd9, 0x4a, 0x7b, 0xd4, 0xe5, 0xa3,
	0x2c, 0x57, 0x68, 0x53, 0x8a, 0xee, 0xc7, 0xc8, 0x94, 0xbe, 0x7d, 0xbf, 0x76, 0x6e, 0x2a, 0x27,
	0xab, 0x6b, 0xd9, 0x66, 0x6d, 0x25, 0x6b, 0xd5, 0x15, 0x9f, 0xff, 0xb5, 0xc0, 0xc9, 0xa4, 0x33,
	0x0f, 0x7b, 0xef, 0x68, 0xd5, 0x14, 0xa4, 0xbe, 0x5b, 0x9c, 0xff, 0xb2, 0xc0, 0xcb, 0x1d, 0xcc,
	0x3e, 0x06, 0xc3, 0x98, 0x85, 0x82, 0x78, 0xfe, 0x67, 0x9f, 0x4c, 0xe1, 0x23, 0x7b, 0x08, 0x0e,
	0x62, 0x76, 0xc9, 0xf8, 0x17, 0x06, 0x07, 0xf6, 0x11, 0x00, 0x1e, 0xa7, 0xd4, 0x0f, 0x43, 0x9f,
	0x33, 0x08, 0x6d, 0x08, 0x0e, 0x09, 0xbb, 0xc0, 0xcc, 0x23, 0xd3, 0xd4, 0x13, 0x1e, 0xb4, 0x74,
	0x06, 0xc5, 0x2c, 0xc6, 0x57, 0x46, 0xef, 0x6d, 0x69, 0x0a, 0x1f, 0x6f, 0xe9, 0x04, 0x3e, 0xb7,
	0x4f, 0xc1, 0x2b, 0x8a, 0xbf, 0xfa, 0xd4, 0xbf, 0x26, 0xa9, 0xc7, 0x59, 0x42, 0x02, 0x6d, 0x1d,
	0x42, 0x60, 0x8f, 0xc0, 0x9b, 0x1d, 0x24, 0x4d, 0xf0, 0x55, 0x4c, 0xe0, 0xd0, 0x7e, 0x0d, 0x5e,
	0x08, 0x3c, 0x23, 0x29, 0x67, 0x24, 0x15, 0x01, 0xa7, 0x3c, 0x22, 0x53, 0xf8, 0x44, 0xf7, 0x2f,
	0x48, 0xe0, 0x11, 0x16, 0x99, 0x06, 0x0e, 0xf5, 0x85, 0x11, 0x0e, 0x66, 0x44, 0x6b, 0x0c, 0xf7,
	0xb7, 0x34, 0x85, 0x47, 0xf6, 0x3b, 0x70, 0xd2, 0x6b, 0x9f, 0x8a, 0x80, 0x98, 0x3f, 0x4b, 0xc3,
	0x0b, 0x1c, 0x10, 0x78, 0xac, 0xbb, 0xeb, 0x21, 0x8f, 0xa3, 0x00, 0xb3, 0xcb, 0x9e, 0x1c, 0xe8,
	0x7b, 0x7a, 0x12, 0x70, 0x1c, 0xc2, 0xa7, 0xfa, 0x29, 0xfa, 0x83, 0x50, 0x10, 0x36, 0x85, 0xcf,
	0x26, 0x7f, 0x06, 0xe0, 0xec, 0x9b, 0x5c, 0xa0, 0x07, 0x47, 0x38, 0x39, 0xdd, 0x31, 0x05, 0xa1,
	0xc7, 0x27, 0x06, 0xd7, 0x93, 0xbe, 0xb4, 0x94, 0xb7, 0x59, 0x5d, 0x22, 0xb9, 0x2c, 0x9d, 0xb2,
	0xa8, 0xcd, 0x70, 0xd7, 0x8b, 0xd4, 0x54, 0xea, 0x3f, 0x7b, 0xf5, 0xc9, 0x7c, 0xbf, 0x5b, 0x7b,
	0x33, 0x8c, 0x7f, 0x58, 0xa3, 0x59, 0x67, 0x85, 0x73, 0x85, 0xba, 0x50, 0x47, 0xc9, 0x18, 0xe9,
	0x6d, 0x50, 0x3f, 0xd7, 0x7c, 0x8e, 0x73, 0x35, 0xdf, 0xf0, 0x79, 0x32, 0x9e, 0x1b, 0xfe, 0xdb,
	0x3a, 0xeb, 0x0e, 0x5d, 0x17, 0xe7, 0xca, 0x75, 0x37, 0x19, 0xae, 0x9b, 0x8c, 0x5d, 0xd7, 0xe4,
	0xdc, 0xec, 0x9b, 0xc6, 0x3e, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x25, 0x30, 0x46, 0xd8, 0xef,
	0x02, 0x00, 0x00,
}

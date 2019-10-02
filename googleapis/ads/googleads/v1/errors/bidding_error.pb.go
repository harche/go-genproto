// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/bidding_error.proto

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

// Enum describing possible bidding errors.
type BiddingErrorEnum_BiddingError int32

const (
	// Enum unspecified.
	BiddingErrorEnum_UNSPECIFIED BiddingErrorEnum_BiddingError = 0
	// The received error code is not known in this version.
	BiddingErrorEnum_UNKNOWN BiddingErrorEnum_BiddingError = 1
	// Cannot transition to new bidding strategy.
	BiddingErrorEnum_BIDDING_STRATEGY_TRANSITION_NOT_ALLOWED BiddingErrorEnum_BiddingError = 2
	// Cannot attach bidding strategy to campaign.
	BiddingErrorEnum_CANNOT_ATTACH_BIDDING_STRATEGY_TO_CAMPAIGN BiddingErrorEnum_BiddingError = 7
	// Bidding strategy is not supported or cannot be used as anonymous.
	BiddingErrorEnum_INVALID_ANONYMOUS_BIDDING_STRATEGY_TYPE BiddingErrorEnum_BiddingError = 10
	// The type does not match the named strategy's type.
	BiddingErrorEnum_INVALID_BIDDING_STRATEGY_TYPE BiddingErrorEnum_BiddingError = 14
	// The bid is invalid.
	BiddingErrorEnum_INVALID_BID BiddingErrorEnum_BiddingError = 17
	// Bidding strategy is not available for the account type.
	BiddingErrorEnum_BIDDING_STRATEGY_NOT_AVAILABLE_FOR_ACCOUNT_TYPE BiddingErrorEnum_BiddingError = 18
	// Conversion tracking is not enabled for the campaign for VBB transition.
	BiddingErrorEnum_CONVERSION_TRACKING_NOT_ENABLED BiddingErrorEnum_BiddingError = 19
	// Not enough conversions tracked for VBB transitions.
	BiddingErrorEnum_NOT_ENOUGH_CONVERSIONS BiddingErrorEnum_BiddingError = 20
	// Campaign can not be created with given bidding strategy. It can be
	// transitioned to the strategy, once eligible.
	BiddingErrorEnum_CANNOT_CREATE_CAMPAIGN_WITH_BIDDING_STRATEGY BiddingErrorEnum_BiddingError = 21
	// Cannot target content network only as campaign uses Page One Promoted
	// bidding strategy.
	BiddingErrorEnum_CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_CAMPAIGN_LEVEL_POP_BIDDING_STRATEGY BiddingErrorEnum_BiddingError = 23
	// Budget Optimizer and Target Spend bidding strategies are not supported
	// for campaigns with AdSchedule targeting.
	BiddingErrorEnum_BIDDING_STRATEGY_NOT_SUPPORTED_WITH_AD_SCHEDULE BiddingErrorEnum_BiddingError = 24
	// Pay per conversion is not available to all the customer, only few
	// whitelisted customers can use this.
	BiddingErrorEnum_PAY_PER_CONVERSION_NOT_AVAILABLE_FOR_CUSTOMER BiddingErrorEnum_BiddingError = 25
	// Pay per conversion is not allowed with Target CPA.
	BiddingErrorEnum_PAY_PER_CONVERSION_NOT_ALLOWED_WITH_TARGET_CPA BiddingErrorEnum_BiddingError = 26
	// Cannot set bidding strategy to Manual CPM for search network only
	// campaigns.
	BiddingErrorEnum_BIDDING_STRATEGY_NOT_ALLOWED_FOR_SEARCH_ONLY_CAMPAIGNS BiddingErrorEnum_BiddingError = 27
	// The bidding strategy is not supported for use in drafts or experiments.
	BiddingErrorEnum_BIDDING_STRATEGY_NOT_SUPPORTED_IN_DRAFTS_OR_EXPERIMENTS BiddingErrorEnum_BiddingError = 28
	// Bidding strategy type does not support product type ad group criterion.
	BiddingErrorEnum_BIDDING_STRATEGY_TYPE_DOES_NOT_SUPPORT_PRODUCT_TYPE_ADGROUP_CRITERION BiddingErrorEnum_BiddingError = 29
	// Bid amount is too small.
	BiddingErrorEnum_BID_TOO_SMALL BiddingErrorEnum_BiddingError = 30
	// Bid amount is too big.
	BiddingErrorEnum_BID_TOO_BIG BiddingErrorEnum_BiddingError = 31
	// Bid has too many fractional digit precision.
	BiddingErrorEnum_BID_TOO_MANY_FRACTIONAL_DIGITS BiddingErrorEnum_BiddingError = 32
	// Invalid domain name specified.
	BiddingErrorEnum_INVALID_DOMAIN_NAME BiddingErrorEnum_BiddingError = 33
	// The field is not compatible with the payment mode.
	BiddingErrorEnum_NOT_COMPATIBLE_WITH_PAYMENT_MODE BiddingErrorEnum_BiddingError = 34
	// The field is not compatible with the budget type.
	BiddingErrorEnum_NOT_COMPATIBLE_WITH_BUDGET_TYPE BiddingErrorEnum_BiddingError = 35
	// The field is not compatible with the bidding strategy type.
	BiddingErrorEnum_NOT_COMPATIBLE_WITH_BIDDING_STRATEGY_TYPE BiddingErrorEnum_BiddingError = 36
)

var BiddingErrorEnum_BiddingError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "BIDDING_STRATEGY_TRANSITION_NOT_ALLOWED",
	7:  "CANNOT_ATTACH_BIDDING_STRATEGY_TO_CAMPAIGN",
	10: "INVALID_ANONYMOUS_BIDDING_STRATEGY_TYPE",
	14: "INVALID_BIDDING_STRATEGY_TYPE",
	17: "INVALID_BID",
	18: "BIDDING_STRATEGY_NOT_AVAILABLE_FOR_ACCOUNT_TYPE",
	19: "CONVERSION_TRACKING_NOT_ENABLED",
	20: "NOT_ENOUGH_CONVERSIONS",
	21: "CANNOT_CREATE_CAMPAIGN_WITH_BIDDING_STRATEGY",
	23: "CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_CAMPAIGN_LEVEL_POP_BIDDING_STRATEGY",
	24: "BIDDING_STRATEGY_NOT_SUPPORTED_WITH_AD_SCHEDULE",
	25: "PAY_PER_CONVERSION_NOT_AVAILABLE_FOR_CUSTOMER",
	26: "PAY_PER_CONVERSION_NOT_ALLOWED_WITH_TARGET_CPA",
	27: "BIDDING_STRATEGY_NOT_ALLOWED_FOR_SEARCH_ONLY_CAMPAIGNS",
	28: "BIDDING_STRATEGY_NOT_SUPPORTED_IN_DRAFTS_OR_EXPERIMENTS",
	29: "BIDDING_STRATEGY_TYPE_DOES_NOT_SUPPORT_PRODUCT_TYPE_ADGROUP_CRITERION",
	30: "BID_TOO_SMALL",
	31: "BID_TOO_BIG",
	32: "BID_TOO_MANY_FRACTIONAL_DIGITS",
	33: "INVALID_DOMAIN_NAME",
	34: "NOT_COMPATIBLE_WITH_PAYMENT_MODE",
	35: "NOT_COMPATIBLE_WITH_BUDGET_TYPE",
	36: "NOT_COMPATIBLE_WITH_BIDDING_STRATEGY_TYPE",
}

var BiddingErrorEnum_BiddingError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"BIDDING_STRATEGY_TRANSITION_NOT_ALLOWED":                                     2,
	"CANNOT_ATTACH_BIDDING_STRATEGY_TO_CAMPAIGN":                                  7,
	"INVALID_ANONYMOUS_BIDDING_STRATEGY_TYPE":                                     10,
	"INVALID_BIDDING_STRATEGY_TYPE":                                               14,
	"INVALID_BID":                                                                 17,
	"BIDDING_STRATEGY_NOT_AVAILABLE_FOR_ACCOUNT_TYPE":                             18,
	"CONVERSION_TRACKING_NOT_ENABLED":                                             19,
	"NOT_ENOUGH_CONVERSIONS":                                                      20,
	"CANNOT_CREATE_CAMPAIGN_WITH_BIDDING_STRATEGY":                                21,
	"CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_CAMPAIGN_LEVEL_POP_BIDDING_STRATEGY": 23,
	"BIDDING_STRATEGY_NOT_SUPPORTED_WITH_AD_SCHEDULE":                             24,
	"PAY_PER_CONVERSION_NOT_AVAILABLE_FOR_CUSTOMER":                               25,
	"PAY_PER_CONVERSION_NOT_ALLOWED_WITH_TARGET_CPA":                              26,
	"BIDDING_STRATEGY_NOT_ALLOWED_FOR_SEARCH_ONLY_CAMPAIGNS":                      27,
	"BIDDING_STRATEGY_NOT_SUPPORTED_IN_DRAFTS_OR_EXPERIMENTS":                     28,
	"BIDDING_STRATEGY_TYPE_DOES_NOT_SUPPORT_PRODUCT_TYPE_ADGROUP_CRITERION":       29,
	"BID_TOO_SMALL":                             30,
	"BID_TOO_BIG":                               31,
	"BID_TOO_MANY_FRACTIONAL_DIGITS":            32,
	"INVALID_DOMAIN_NAME":                       33,
	"NOT_COMPATIBLE_WITH_PAYMENT_MODE":          34,
	"NOT_COMPATIBLE_WITH_BUDGET_TYPE":           35,
	"NOT_COMPATIBLE_WITH_BIDDING_STRATEGY_TYPE": 36,
}

func (x BiddingErrorEnum_BiddingError) String() string {
	return proto.EnumName(BiddingErrorEnum_BiddingError_name, int32(x))
}

func (BiddingErrorEnum_BiddingError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a54920be9d4893fa, []int{0, 0}
}

// Container for enum describing possible bidding errors.
type BiddingErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BiddingErrorEnum) Reset()         { *m = BiddingErrorEnum{} }
func (m *BiddingErrorEnum) String() string { return proto.CompactTextString(m) }
func (*BiddingErrorEnum) ProtoMessage()    {}
func (*BiddingErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_a54920be9d4893fa, []int{0}
}

func (m *BiddingErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BiddingErrorEnum.Unmarshal(m, b)
}
func (m *BiddingErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BiddingErrorEnum.Marshal(b, m, deterministic)
}
func (m *BiddingErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BiddingErrorEnum.Merge(m, src)
}
func (m *BiddingErrorEnum) XXX_Size() int {
	return xxx_messageInfo_BiddingErrorEnum.Size(m)
}
func (m *BiddingErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_BiddingErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_BiddingErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.BiddingErrorEnum_BiddingError", BiddingErrorEnum_BiddingError_name, BiddingErrorEnum_BiddingError_value)
	proto.RegisterType((*BiddingErrorEnum)(nil), "google.ads.googleads.v1.errors.BiddingErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/bidding_error.proto", fileDescriptor_a54920be9d4893fa)
}

var fileDescriptor_a54920be9d4893fa = []byte{
	// 739 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x5d, 0x8f, 0x1b, 0x35,
	0x14, 0x65, 0x17, 0x89, 0x56, 0x2e, 0x1f, 0x5e, 0x17, 0x28, 0x2c, 0xed, 0xb6, 0x4d, 0x2b, 0x21,
	0x3e, 0x3a, 0x21, 0xad, 0x04, 0xd2, 0xf4, 0xe9, 0xce, 0xd8, 0x99, 0x58, 0x99, 0xb1, 0x2d, 0xdb,
	0x93, 0x10, 0x14, 0xc9, 0x4a, 0xc9, 0x2a, 0x8a, 0xd4, 0x66, 0x56, 0x99, 0xa5, 0x3f, 0x87, 0x07,
	0x1e, 0x79, 0xe5, 0x5f, 0xf0, 0x53, 0x10, 0x3f, 0x02, 0x79, 0x9c, 0x09, 0x91, 0x32, 0x0b, 0x4f,
	0xb9, 0xb9, 0x3a, 0xe7, 0xd8, 0xe7, 0xce, 0xf1, 0x45, 0xcf, 0x57, 0x55, 0xb5, 0x7a, 0x7d, 0xd9,
	0x5f, 0x2c, 0xeb, 0x7e, 0x28, 0x7d, 0xf5, 0x76, 0xd0, 0xbf, 0xdc, 0x6e, 0xab, 0x6d, 0xdd, 0x7f,
	0xb5, 0x5e, 0x2e, 0xd7, 0x9b, 0x95, 0x6b, 0xfe, 0x46, 0x57, 0xdb, 0xea, 0xba, 0x22, 0x17, 0x01,
	0x18, 0x2d, 0x96, 0x75, 0xb4, 0xe7, 0x44, 0x6f, 0x07, 0x51, 0xe0, 0x9c, 0xdf, 0x6f, 0x35, 0xaf,
	0xd6, 0xfd, 0xc5, 0x66, 0x53, 0x5d, 0x2f, 0xae, 0xd7, 0xd5, 0xa6, 0x0e, 0xec, 0xde, 0x1f, 0xb7,
	0x11, 0x4e, 0x82, 0x2a, 0xf3, 0x78, 0xb6, 0xf9, 0xe5, 0x4d, 0xef, 0xd7, 0xdb, 0xe8, 0xfd, 0xc3,
	0x26, 0xf9, 0x08, 0xdd, 0x29, 0x85, 0x51, 0x2c, 0xe5, 0x43, 0xce, 0x28, 0x7e, 0x87, 0xdc, 0x41,
	0xb7, 0x4a, 0x31, 0x16, 0x72, 0x2a, 0xf0, 0x09, 0xf9, 0x06, 0x7d, 0x99, 0x70, 0x4a, 0xb9, 0xc8,
	0x9c, 0xb1, 0x1a, 0x2c, 0xcb, 0x66, 0xce, 0x6a, 0x10, 0x86, 0x5b, 0x2e, 0x85, 0x13, 0xd2, 0x3a,
	0xc8, 0x73, 0x39, 0x65, 0x14, 0x9f, 0x92, 0x08, 0x7d, 0x9d, 0x82, 0x68, 0x7a, 0xd6, 0x42, 0x3a,
	0x72, 0xc7, 0x54, 0xe9, 0x52, 0x28, 0x14, 0xf0, 0x4c, 0xe0, 0x5b, 0x5e, 0x9c, 0x8b, 0x09, 0xe4,
	0x9c, 0x3a, 0x10, 0x52, 0xcc, 0x0a, 0x59, 0x9a, 0x0e, 0xce, 0x4c, 0x31, 0x8c, 0xc8, 0x63, 0xf4,
	0xa0, 0x05, 0x77, 0x43, 0x3e, 0xf4, 0x56, 0x0e, 0x20, 0xf8, 0x8c, 0xbc, 0x40, 0xfd, 0x23, 0x6c,
	0x73, 0xbd, 0x09, 0xf0, 0x1c, 0x92, 0x9c, 0xb9, 0xa1, 0xd4, 0x0e, 0xd2, 0x54, 0x96, 0xc2, 0x06,
	0x15, 0x42, 0x9e, 0xa0, 0x87, 0xa9, 0x14, 0x13, 0xa6, 0x8d, 0x77, 0x68, 0x35, 0xa4, 0x63, 0x2f,
	0xe0, 0x79, 0x4c, 0x78, 0x12, 0xc5, 0x77, 0xc9, 0x39, 0xfa, 0x34, 0x34, 0x64, 0x99, 0x8d, 0xdc,
	0xbf, 0x78, 0x83, 0x3f, 0x26, 0xdf, 0xa1, 0x6f, 0x77, 0x63, 0x48, 0x35, 0x03, 0xcb, 0xf6, 0x96,
	0xdd, 0x94, 0xdb, 0xe3, 0xa1, 0xe0, 0x4f, 0x88, 0x44, 0xe3, 0x1d, 0xc3, 0x82, 0xce, 0x98, 0xf5,
	0x82, 0x96, 0x09, 0xeb, 0x04, 0xb3, 0x53, 0xa9, 0xc7, 0x4e, 0x8a, 0x7c, 0x16, 0xd8, 0x7b, 0xad,
	0x9c, 0x4d, 0x58, 0xee, 0x94, 0x54, 0xc7, 0x82, 0xf7, 0x6e, 0x34, 0x6e, 0x4a, 0xa5, 0xa4, 0xb6,
	0x8c, 0x06, 0x31, 0xa0, 0xce, 0xa4, 0x23, 0x46, 0xcb, 0x9c, 0xe1, 0xcf, 0xc8, 0x00, 0x3d, 0x53,
	0x30, 0x73, 0x8a, 0xe9, 0x03, 0x43, 0x1d, 0xf3, 0x4a, 0x4b, 0x63, 0x65, 0xc1, 0x34, 0xfe, 0x9c,
	0x3c, 0x47, 0xd1, 0x4d, 0x94, 0x90, 0x8a, 0x70, 0x4e, 0xeb, 0x4a, 0x01, 0x3e, 0x27, 0x31, 0xfa,
	0xbe, 0xfb, 0xa3, 0xec, 0x18, 0xfe, 0x08, 0xc3, 0x40, 0xa7, 0xa3, 0x60, 0xbb, 0x75, 0x6c, 0xf0,
	0x17, 0xe4, 0x25, 0xfa, 0xe1, 0x7f, 0x7c, 0x71, 0xe1, 0xa8, 0x86, 0xa1, 0x35, 0x4e, 0x6a, 0xc7,
	0x7e, 0x54, 0x4c, 0xf3, 0x82, 0x09, 0x6b, 0xf0, 0x7d, 0xc2, 0x11, 0xeb, 0x4c, 0x8e, 0xa3, 0x92,
	0x99, 0x43, 0x19, 0xa7, 0xb4, 0xa4, 0x65, 0x1a, 0x12, 0xe1, 0x80, 0x66, 0x5a, 0x96, 0xca, 0xa5,
	0x9a, 0x5b, 0xa6, 0xb9, 0x14, 0xf8, 0x01, 0x39, 0x43, 0x1f, 0x24, 0x9c, 0x3a, 0x2b, 0xa5, 0x33,
	0x05, 0xe4, 0x39, 0xbe, 0xf0, 0xe1, 0x6b, 0x5b, 0x09, 0xcf, 0xf0, 0x43, 0xd2, 0x43, 0x17, 0x6d,
	0xa3, 0x00, 0x31, 0x73, 0x43, 0x0d, 0xa9, 0x7f, 0x34, 0x90, 0x3b, 0xca, 0x33, 0x6e, 0x0d, 0x7e,
	0x44, 0xee, 0xa1, 0xbb, 0x6d, 0x62, 0xa9, 0x2c, 0x80, 0x0b, 0x27, 0xa0, 0x60, 0xf8, 0x31, 0x79,
	0x8a, 0x1e, 0x35, 0x01, 0x92, 0x85, 0x02, 0xcb, 0xfd, 0xe4, 0x9b, 0x41, 0x2a, 0x98, 0x79, 0x37,
	0xae, 0x90, 0x94, 0xe1, 0x9e, 0x8f, 0x6a, 0x17, 0x2a, 0x29, 0xa9, 0x1f, 0x77, 0x93, 0xe7, 0x27,
	0xe4, 0x19, 0xfa, 0xaa, 0x13, 0xd4, 0xf9, 0x88, 0x9e, 0x26, 0x7f, 0x9f, 0xa0, 0xde, 0xcf, 0xd5,
	0x9b, 0xe8, 0xbf, 0x57, 0x4f, 0x72, 0x76, 0xb8, 0x44, 0x94, 0xdf, 0x37, 0xea, 0xe4, 0x27, 0xba,
	0x23, 0xad, 0xaa, 0xd7, 0x8b, 0xcd, 0x2a, 0xaa, 0xb6, 0xab, 0xfe, 0xea, 0x72, 0xd3, 0x6c, 0xa3,
	0x76, 0xe7, 0x5d, 0xad, 0xeb, 0x9b, 0x56, 0xe0, 0xcb, 0xf0, 0xf3, 0xdb, 0xe9, 0xbb, 0x19, 0xc0,
	0xef, 0xa7, 0x17, 0x59, 0x10, 0x83, 0x65, 0x1d, 0x85, 0xd2, 0x57, 0x93, 0x41, 0xd4, 0x1c, 0x59,
	0xff, 0xd9, 0x02, 0xe6, 0xb0, 0xac, 0xe7, 0x7b, 0xc0, 0x7c, 0x32, 0x98, 0x07, 0xc0, 0x5f, 0xa7,
	0xbd, 0xd0, 0x8d, 0x63, 0x58, 0xd6, 0x71, 0xbc, 0x87, 0xc4, 0xf1, 0x64, 0x10, 0xc7, 0x01, 0xf4,
	0xea, 0xbd, 0xe6, 0x76, 0x2f, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x4b, 0xa3, 0xc4, 0x9f,
	0x05, 0x00, 0x00,
}

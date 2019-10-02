// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/campaign_error.proto

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

// Enum describing possible campaign errors.
type CampaignErrorEnum_CampaignError int32

const (
	// Enum unspecified.
	CampaignErrorEnum_UNSPECIFIED CampaignErrorEnum_CampaignError = 0
	// The received error code is not known in this version.
	CampaignErrorEnum_UNKNOWN CampaignErrorEnum_CampaignError = 1
	// Cannot target content network.
	CampaignErrorEnum_CANNOT_TARGET_CONTENT_NETWORK CampaignErrorEnum_CampaignError = 3
	// Cannot target search network.
	CampaignErrorEnum_CANNOT_TARGET_SEARCH_NETWORK CampaignErrorEnum_CampaignError = 4
	// Cannot cover search network without google search network.
	CampaignErrorEnum_CANNOT_TARGET_SEARCH_NETWORK_WITHOUT_GOOGLE_SEARCH CampaignErrorEnum_CampaignError = 5
	// Cannot target Google Search network for a CPM campaign.
	CampaignErrorEnum_CANNOT_TARGET_GOOGLE_SEARCH_FOR_CPM_CAMPAIGN CampaignErrorEnum_CampaignError = 6
	// Must target at least one network.
	CampaignErrorEnum_CAMPAIGN_MUST_TARGET_AT_LEAST_ONE_NETWORK CampaignErrorEnum_CampaignError = 7
	// Only some Google partners are allowed to target partner search network.
	CampaignErrorEnum_CANNOT_TARGET_PARTNER_SEARCH_NETWORK CampaignErrorEnum_CampaignError = 8
	// Cannot target content network only as campaign has criteria-level bidding
	// strategy.
	CampaignErrorEnum_CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_CRITERIA_LEVEL_BIDDING_STRATEGY CampaignErrorEnum_CampaignError = 9
	// Cannot modify the start or end date such that the campaign duration would
	// not contain the durations of all runnable trials.
	CampaignErrorEnum_CAMPAIGN_DURATION_MUST_CONTAIN_ALL_RUNNABLE_TRIALS CampaignErrorEnum_CampaignError = 10
	// Cannot modify dates, budget or campaign name of a trial campaign.
	CampaignErrorEnum_CANNOT_MODIFY_FOR_TRIAL_CAMPAIGN CampaignErrorEnum_CampaignError = 11
	// Trying to modify the name of an active or paused campaign, where the name
	// is already assigned to another active or paused campaign.
	CampaignErrorEnum_DUPLICATE_CAMPAIGN_NAME CampaignErrorEnum_CampaignError = 12
	// Two fields are in conflicting modes.
	CampaignErrorEnum_INCOMPATIBLE_CAMPAIGN_FIELD CampaignErrorEnum_CampaignError = 13
	// Campaign name cannot be used.
	CampaignErrorEnum_INVALID_CAMPAIGN_NAME CampaignErrorEnum_CampaignError = 14
	// Given status is invalid.
	CampaignErrorEnum_INVALID_AD_SERVING_OPTIMIZATION_STATUS CampaignErrorEnum_CampaignError = 15
	// Error in the campaign level tracking url.
	CampaignErrorEnum_INVALID_TRACKING_URL CampaignErrorEnum_CampaignError = 16
	// Cannot set both tracking url template and tracking setting. An user has
	// to clear legacy tracking setting in order to add tracking url template.
	CampaignErrorEnum_CANNOT_SET_BOTH_TRACKING_URL_TEMPLATE_AND_TRACKING_SETTING CampaignErrorEnum_CampaignError = 17
	// The maximum number of impressions for Frequency Cap should be an integer
	// greater than 0.
	CampaignErrorEnum_MAX_IMPRESSIONS_NOT_IN_RANGE CampaignErrorEnum_CampaignError = 18
	// Only the Day, Week and Month time units are supported.
	CampaignErrorEnum_TIME_UNIT_NOT_SUPPORTED CampaignErrorEnum_CampaignError = 19
	// Operation not allowed on a campaign whose serving status has ended
	CampaignErrorEnum_INVALID_OPERATION_IF_SERVING_STATUS_HAS_ENDED CampaignErrorEnum_CampaignError = 20
	// This budget is exclusively linked to a Campaign that is using experiments
	// so it cannot be shared.
	CampaignErrorEnum_BUDGET_CANNOT_BE_SHARED CampaignErrorEnum_CampaignError = 21
	// Campaigns using experiments cannot use a shared budget.
	CampaignErrorEnum_CAMPAIGN_CANNOT_USE_SHARED_BUDGET CampaignErrorEnum_CampaignError = 22
	// A different budget cannot be assigned to a campaign when there are
	// running or scheduled trials.
	CampaignErrorEnum_CANNOT_CHANGE_BUDGET_ON_CAMPAIGN_WITH_TRIALS CampaignErrorEnum_CampaignError = 23
	// No link found between the campaign and the label.
	CampaignErrorEnum_CAMPAIGN_LABEL_DOES_NOT_EXIST CampaignErrorEnum_CampaignError = 24
	// The label has already been attached to the campaign.
	CampaignErrorEnum_CAMPAIGN_LABEL_ALREADY_EXISTS CampaignErrorEnum_CampaignError = 25
	// A ShoppingSetting was not found when creating a shopping campaign.
	CampaignErrorEnum_MISSING_SHOPPING_SETTING CampaignErrorEnum_CampaignError = 26
	// The country in shopping setting is not an allowed country.
	CampaignErrorEnum_INVALID_SHOPPING_SALES_COUNTRY CampaignErrorEnum_CampaignError = 27
	// A Campaign with channel sub type UNIVERSAL_APP_CAMPAIGN must have a
	// UniversalAppCampaignSetting specified.
	CampaignErrorEnum_MISSING_UNIVERSAL_APP_CAMPAIGN_SETTING CampaignErrorEnum_CampaignError = 30
	// The requested channel type is not available according to the customer's
	// account setting.
	CampaignErrorEnum_ADVERTISING_CHANNEL_TYPE_NOT_AVAILABLE_FOR_ACCOUNT_TYPE CampaignErrorEnum_CampaignError = 31
	// The AdvertisingChannelSubType is not a valid subtype of the primary
	// channel type.
	CampaignErrorEnum_INVALID_ADVERTISING_CHANNEL_SUB_TYPE CampaignErrorEnum_CampaignError = 32
	// At least one conversion must be selected.
	CampaignErrorEnum_AT_LEAST_ONE_CONVERSION_MUST_BE_SELECTED CampaignErrorEnum_CampaignError = 33
	// Setting ad rotation mode for a campaign is not allowed. Ad rotation mode
	// at campaign is deprecated.
	CampaignErrorEnum_CANNOT_SET_AD_ROTATION_MODE CampaignErrorEnum_CampaignError = 34
	// Trying to change start date on a campaign that has started.
	CampaignErrorEnum_CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED CampaignErrorEnum_CampaignError = 35
	// Trying to modify a date into the past.
	CampaignErrorEnum_CANNOT_SET_DATE_TO_PAST CampaignErrorEnum_CampaignError = 36
	// Hotel center id in the hotel setting does not match any customer links.
	CampaignErrorEnum_MISSING_HOTEL_CUSTOMER_LINK CampaignErrorEnum_CampaignError = 37
	// Hotel center id in the hotel setting must match an active customer link.
	CampaignErrorEnum_INVALID_HOTEL_CUSTOMER_LINK CampaignErrorEnum_CampaignError = 38
	// Hotel setting was not found when creating a hotel ads campaign.
	CampaignErrorEnum_MISSING_HOTEL_SETTING CampaignErrorEnum_CampaignError = 39
	// A Campaign cannot use shared campaign budgets and be part of a campaign
	// group.
	CampaignErrorEnum_CANNOT_USE_SHARED_CAMPAIGN_BUDGET_WHILE_PART_OF_CAMPAIGN_GROUP CampaignErrorEnum_CampaignError = 40
	// The app ID was not found.
	CampaignErrorEnum_APP_NOT_FOUND CampaignErrorEnum_CampaignError = 41
	// Campaign.shopping_setting.enable_local is not supported for the specified
	// campaign type.
	CampaignErrorEnum_SHOPPING_ENABLE_LOCAL_NOT_SUPPORTED_FOR_CAMPAIGN_TYPE CampaignErrorEnum_CampaignError = 42
	// The merchant does not support the creation of campaigns for Shopping
	// Comparison Listing Ads.
	CampaignErrorEnum_MERCHANT_NOT_ALLOWED_FOR_COMPARISON_LISTING_ADS CampaignErrorEnum_CampaignError = 43
	// The App campaign for engagement cannot be created because there aren't
	// enough installs.
	CampaignErrorEnum_INSUFFICIENT_APP_INSTALLS_COUNT CampaignErrorEnum_CampaignError = 44
)

var CampaignErrorEnum_CampaignError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	3:  "CANNOT_TARGET_CONTENT_NETWORK",
	4:  "CANNOT_TARGET_SEARCH_NETWORK",
	5:  "CANNOT_TARGET_SEARCH_NETWORK_WITHOUT_GOOGLE_SEARCH",
	6:  "CANNOT_TARGET_GOOGLE_SEARCH_FOR_CPM_CAMPAIGN",
	7:  "CAMPAIGN_MUST_TARGET_AT_LEAST_ONE_NETWORK",
	8:  "CANNOT_TARGET_PARTNER_SEARCH_NETWORK",
	9:  "CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_CRITERIA_LEVEL_BIDDING_STRATEGY",
	10: "CAMPAIGN_DURATION_MUST_CONTAIN_ALL_RUNNABLE_TRIALS",
	11: "CANNOT_MODIFY_FOR_TRIAL_CAMPAIGN",
	12: "DUPLICATE_CAMPAIGN_NAME",
	13: "INCOMPATIBLE_CAMPAIGN_FIELD",
	14: "INVALID_CAMPAIGN_NAME",
	15: "INVALID_AD_SERVING_OPTIMIZATION_STATUS",
	16: "INVALID_TRACKING_URL",
	17: "CANNOT_SET_BOTH_TRACKING_URL_TEMPLATE_AND_TRACKING_SETTING",
	18: "MAX_IMPRESSIONS_NOT_IN_RANGE",
	19: "TIME_UNIT_NOT_SUPPORTED",
	20: "INVALID_OPERATION_IF_SERVING_STATUS_HAS_ENDED",
	21: "BUDGET_CANNOT_BE_SHARED",
	22: "CAMPAIGN_CANNOT_USE_SHARED_BUDGET",
	23: "CANNOT_CHANGE_BUDGET_ON_CAMPAIGN_WITH_TRIALS",
	24: "CAMPAIGN_LABEL_DOES_NOT_EXIST",
	25: "CAMPAIGN_LABEL_ALREADY_EXISTS",
	26: "MISSING_SHOPPING_SETTING",
	27: "INVALID_SHOPPING_SALES_COUNTRY",
	30: "MISSING_UNIVERSAL_APP_CAMPAIGN_SETTING",
	31: "ADVERTISING_CHANNEL_TYPE_NOT_AVAILABLE_FOR_ACCOUNT_TYPE",
	32: "INVALID_ADVERTISING_CHANNEL_SUB_TYPE",
	33: "AT_LEAST_ONE_CONVERSION_MUST_BE_SELECTED",
	34: "CANNOT_SET_AD_ROTATION_MODE",
	35: "CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED",
	36: "CANNOT_SET_DATE_TO_PAST",
	37: "MISSING_HOTEL_CUSTOMER_LINK",
	38: "INVALID_HOTEL_CUSTOMER_LINK",
	39: "MISSING_HOTEL_SETTING",
	40: "CANNOT_USE_SHARED_CAMPAIGN_BUDGET_WHILE_PART_OF_CAMPAIGN_GROUP",
	41: "APP_NOT_FOUND",
	42: "SHOPPING_ENABLE_LOCAL_NOT_SUPPORTED_FOR_CAMPAIGN_TYPE",
	43: "MERCHANT_NOT_ALLOWED_FOR_COMPARISON_LISTING_ADS",
	44: "INSUFFICIENT_APP_INSTALLS_COUNT",
}

var CampaignErrorEnum_CampaignError_value = map[string]int32{
	"UNSPECIFIED":                   0,
	"UNKNOWN":                       1,
	"CANNOT_TARGET_CONTENT_NETWORK": 3,
	"CANNOT_TARGET_SEARCH_NETWORK":  4,
	"CANNOT_TARGET_SEARCH_NETWORK_WITHOUT_GOOGLE_SEARCH":                      5,
	"CANNOT_TARGET_GOOGLE_SEARCH_FOR_CPM_CAMPAIGN":                            6,
	"CAMPAIGN_MUST_TARGET_AT_LEAST_ONE_NETWORK":                               7,
	"CANNOT_TARGET_PARTNER_SEARCH_NETWORK":                                    8,
	"CANNOT_TARGET_CONTENT_NETWORK_ONLY_WITH_CRITERIA_LEVEL_BIDDING_STRATEGY": 9,
	"CAMPAIGN_DURATION_MUST_CONTAIN_ALL_RUNNABLE_TRIALS":                      10,
	"CANNOT_MODIFY_FOR_TRIAL_CAMPAIGN":                                        11,
	"DUPLICATE_CAMPAIGN_NAME":                                                 12,
	"INCOMPATIBLE_CAMPAIGN_FIELD":                                             13,
	"INVALID_CAMPAIGN_NAME":                                                   14,
	"INVALID_AD_SERVING_OPTIMIZATION_STATUS":                                  15,
	"INVALID_TRACKING_URL":                                                    16,
	"CANNOT_SET_BOTH_TRACKING_URL_TEMPLATE_AND_TRACKING_SETTING":              17,
	"MAX_IMPRESSIONS_NOT_IN_RANGE":                                            18,
	"TIME_UNIT_NOT_SUPPORTED":                                                 19,
	"INVALID_OPERATION_IF_SERVING_STATUS_HAS_ENDED":                           20,
	"BUDGET_CANNOT_BE_SHARED":                                                 21,
	"CAMPAIGN_CANNOT_USE_SHARED_BUDGET":                                       22,
	"CANNOT_CHANGE_BUDGET_ON_CAMPAIGN_WITH_TRIALS":                            23,
	"CAMPAIGN_LABEL_DOES_NOT_EXIST":                                           24,
	"CAMPAIGN_LABEL_ALREADY_EXISTS":                                           25,
	"MISSING_SHOPPING_SETTING":                                                26,
	"INVALID_SHOPPING_SALES_COUNTRY":                                          27,
	"MISSING_UNIVERSAL_APP_CAMPAIGN_SETTING":                                  30,
	"ADVERTISING_CHANNEL_TYPE_NOT_AVAILABLE_FOR_ACCOUNT_TYPE":                 31,
	"INVALID_ADVERTISING_CHANNEL_SUB_TYPE":                                    32,
	"AT_LEAST_ONE_CONVERSION_MUST_BE_SELECTED":                                33,
	"CANNOT_SET_AD_ROTATION_MODE":                                             34,
	"CANNOT_MODIFY_START_DATE_IF_ALREADY_STARTED":                             35,
	"CANNOT_SET_DATE_TO_PAST":                                                 36,
	"MISSING_HOTEL_CUSTOMER_LINK":                                             37,
	"INVALID_HOTEL_CUSTOMER_LINK":                                             38,
	"MISSING_HOTEL_SETTING":                                                   39,
	"CANNOT_USE_SHARED_CAMPAIGN_BUDGET_WHILE_PART_OF_CAMPAIGN_GROUP":          40,
	"APP_NOT_FOUND":                                                           41,
	"SHOPPING_ENABLE_LOCAL_NOT_SUPPORTED_FOR_CAMPAIGN_TYPE":                   42,
	"MERCHANT_NOT_ALLOWED_FOR_COMPARISON_LISTING_ADS":                         43,
	"INSUFFICIENT_APP_INSTALLS_COUNT":                                         44,
}

func (x CampaignErrorEnum_CampaignError) String() string {
	return proto.EnumName(CampaignErrorEnum_CampaignError_name, int32(x))
}

func (CampaignErrorEnum_CampaignError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8032436b204b8568, []int{0, 0}
}

// Container for enum describing possible campaign errors.
type CampaignErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignErrorEnum) Reset()         { *m = CampaignErrorEnum{} }
func (m *CampaignErrorEnum) String() string { return proto.CompactTextString(m) }
func (*CampaignErrorEnum) ProtoMessage()    {}
func (*CampaignErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_8032436b204b8568, []int{0}
}

func (m *CampaignErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignErrorEnum.Unmarshal(m, b)
}
func (m *CampaignErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignErrorEnum.Marshal(b, m, deterministic)
}
func (m *CampaignErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignErrorEnum.Merge(m, src)
}
func (m *CampaignErrorEnum) XXX_Size() int {
	return xxx_messageInfo_CampaignErrorEnum.Size(m)
}
func (m *CampaignErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.CampaignErrorEnum_CampaignError", CampaignErrorEnum_CampaignError_name, CampaignErrorEnum_CampaignError_value)
	proto.RegisterType((*CampaignErrorEnum)(nil), "google.ads.googleads.v2.errors.CampaignErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/campaign_error.proto", fileDescriptor_8032436b204b8568)
}

var fileDescriptor_8032436b204b8568 = []byte{
	// 1062 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0xeb, 0x8e, 0x1b, 0x35,
	0x14, 0xa6, 0x2d, 0xb4, 0xe0, 0xb2, 0xd4, 0x35, 0x2d, 0xbd, 0xb2, 0x6d, 0xd3, 0x0b, 0xbd, 0x6c,
	0x13, 0xc8, 0x0a, 0x10, 0xa9, 0x54, 0xe9, 0x64, 0x7c, 0x32, 0xb1, 0xd6, 0x63, 0x8f, 0x6c, 0x4f,
	0xb6, 0xa9, 0x56, 0xb2, 0x42, 0x77, 0x15, 0xad, 0xd4, 0x4d, 0x56, 0x9b, 0xa5, 0x0f, 0xc4, 0x2f,
	0xc4, 0x43, 0xf0, 0x00, 0x3c, 0x0a, 0x12, 0xef, 0x80, 0x3c, 0x9e, 0x99, 0x6c, 0xda, 0xb2, 0xbf,
	0xc6, 0xe3, 0xf3, 0x9d, 0xcf, 0xe7, 0xf2, 0x1d, 0x9b, 0x6c, 0x4e, 0xe7, 0xf3, 0xe9, 0xdb, 0xbd,
	0xce, 0x64, 0x77, 0xd1, 0x89, 0xcb, 0xb0, 0x7a, 0xd7, 0xed, 0xec, 0x1d, 0x1d, 0xcd, 0x8f, 0x16,
	0x9d, 0x37, 0x93, 0x83, 0xc3, 0xc9, 0xfe, 0x74, 0xe6, 0xcb, 0xff, 0xf6, 0xe1, 0xd1, 0xfc, 0x78,
	0xce, 0xd6, 0x23, 0xb2, 0x3d, 0xd9, 0x5d, 0xb4, 0x1b, 0xa7, 0xf6, 0xbb, 0x6e, 0x3b, 0x3a, 0xdd,
	0xbc, 0x5d, 0x93, 0x1e, 0xee, 0x77, 0x26, 0xb3, 0xd9, 0xfc, 0x78, 0x72, 0xbc, 0x3f, 0x9f, 0x2d,
	0xa2, 0x77, 0xeb, 0xaf, 0x35, 0x72, 0x39, 0xa9, 0x68, 0x31, 0x38, 0xe0, 0xec, 0xb7, 0x83, 0xd6,
	0x1f, 0x6b, 0x64, 0x6d, 0x65, 0x97, 0x5d, 0x22, 0x17, 0x0b, 0x65, 0x73, 0x4c, 0xc4, 0x40, 0x20,
	0xa7, 0x9f, 0xb0, 0x8b, 0xe4, 0x42, 0xa1, 0xb6, 0x94, 0xde, 0x56, 0xf4, 0x0c, 0xbb, 0x47, 0xbe,
	0x4d, 0x40, 0x29, 0xed, 0xbc, 0x03, 0x93, 0xa2, 0xf3, 0x89, 0x56, 0x0e, 0x95, 0xf3, 0x0a, 0xdd,
	0xb6, 0x36, 0x5b, 0xf4, 0x1c, 0xbb, 0x4b, 0x6e, 0xaf, 0x42, 0x2c, 0x82, 0x49, 0x86, 0x0d, 0xe2,
	0x53, 0xf6, 0x13, 0xe9, 0x9e, 0x86, 0xf0, 0xdb, 0xc2, 0x0d, 0x75, 0xe1, 0x7c, 0xaa, 0x75, 0x2a,
	0xb1, 0xb2, 0xd2, 0xcf, 0xd8, 0xf7, 0x64, 0x63, 0xd5, 0x6f, 0x05, 0xe0, 0x07, 0xda, 0xf8, 0x24,
	0xcf, 0x7c, 0x02, 0x59, 0x0e, 0x22, 0x55, 0xf4, 0x3c, 0x7b, 0x4e, 0x9e, 0xd4, 0x7f, 0x3e, 0x2b,
	0x6c, 0xe3, 0x08, 0xce, 0x4b, 0x04, 0xeb, 0xbc, 0x56, 0xd8, 0x04, 0x76, 0x81, 0x3d, 0x26, 0x0f,
	0x56, 0x0f, 0xc8, 0xc1, 0x38, 0x85, 0xe6, 0xfd, 0x14, 0x3e, 0x67, 0x5b, 0x24, 0x3d, 0xb5, 0x0e,
	0x5e, 0x2b, 0x39, 0x2e, 0x13, 0xf1, 0x89, 0x11, 0x0e, 0x8d, 0x00, 0x2f, 0x71, 0x84, 0xd2, 0xf7,
	0x05, 0xe7, 0x42, 0xa5, 0xde, 0x3a, 0x03, 0x0e, 0xd3, 0x31, 0xfd, 0x22, 0xd6, 0xa3, 0x8a, 0x92,
	0x17, 0x06, 0x9c, 0xd0, 0x55, 0xb8, 0x81, 0x15, 0x84, 0xf2, 0x20, 0xa5, 0x37, 0x85, 0x52, 0xd0,
	0x97, 0xe8, 0x9d, 0x11, 0x20, 0x2d, 0x25, 0xec, 0x01, 0xb9, 0x5b, 0x05, 0x91, 0x69, 0x2e, 0x06,
	0xe3, 0xb2, 0x02, 0xa5, 0x75, 0x59, 0x83, 0x8b, 0xec, 0x16, 0xb9, 0xc6, 0x8b, 0x5c, 0x8a, 0x04,
	0x1c, 0x36, 0xfb, 0x5e, 0x41, 0x86, 0xf4, 0x4b, 0x76, 0x87, 0xdc, 0x12, 0x2a, 0xd1, 0x59, 0x0e,
	0x4e, 0x04, 0xee, 0xc6, 0x3e, 0x10, 0x28, 0x39, 0x5d, 0x63, 0x37, 0xc8, 0x55, 0xa1, 0x46, 0x20,
	0x05, 0x7f, 0xcf, 0xf7, 0x2b, 0xf6, 0x94, 0x3c, 0xaa, 0x4d, 0xc0, 0xbd, 0x45, 0x33, 0x0a, 0x79,
	0xe9, 0xdc, 0x89, 0x4c, 0xbc, 0x8e, 0x49, 0x58, 0x07, 0xae, 0xb0, 0xf4, 0x12, 0xbb, 0x4e, 0xae,
	0xd4, 0x58, 0x67, 0x20, 0xd9, 0x0a, 0xc8, 0xc2, 0x48, 0x4a, 0xd9, 0x4b, 0xd2, 0xab, 0x92, 0xb0,
	0xe8, 0x7c, 0x5f, 0xbb, 0xe1, 0x0a, 0xc2, 0x3b, 0xcc, 0x72, 0x19, 0x42, 0x07, 0x75, 0xc2, 0xd7,
	0xa2, 0x73, 0x42, 0xa5, 0xf4, 0x72, 0x90, 0x5b, 0x06, 0xaf, 0xbc, 0xc8, 0x72, 0x83, 0xd6, 0x0a,
	0xad, 0xac, 0x0f, 0x64, 0x42, 0x79, 0x03, 0x2a, 0x45, 0xca, 0x42, 0x01, 0x9c, 0xc8, 0xd0, 0x17,
	0x4a, 0xb8, 0xd2, 0x66, 0x8b, 0x3c, 0xd7, 0xc6, 0x21, 0xa7, 0x5f, 0xb3, 0x1f, 0xc8, 0xf3, 0x3a,
	0x30, 0x9d, 0x63, 0x55, 0x7b, 0x31, 0x68, 0xd2, 0x89, 0x19, 0xf8, 0x21, 0x58, 0x8f, 0x8a, 0x23,
	0xa7, 0x57, 0x02, 0x5f, 0xbf, 0xe0, 0x65, 0xd3, 0x63, 0xe0, 0x7d, 0xf4, 0x76, 0x08, 0x06, 0x39,
	0xbd, 0xca, 0x1e, 0x92, 0x7b, 0x4d, 0x9d, 0x2a, 0x73, 0x61, 0x6b, 0xbb, 0x8f, 0x7e, 0xf4, 0x9b,
	0x13, 0x52, 0x4e, 0x86, 0x21, 0xcc, 0xca, 0xe2, 0xb5, 0x5a, 0x96, 0xb9, 0x54, 0x4f, 0xd5, 0xec,
	0x6b, 0x71, 0xf2, 0x2a, 0x8b, 0x84, 0x3e, 0x4a, 0xcf, 0x35, 0xc6, 0x5c, 0xf1, 0x95, 0xb0, 0x8e,
	0x5e, 0xff, 0x08, 0x04, 0xa4, 0x41, 0xe0, 0xe3, 0x88, 0xb0, 0xf4, 0x06, 0xbb, 0x4d, 0xae, 0x67,
	0xc2, 0xda, 0x32, 0xb3, 0xa1, 0xce, 0xf3, 0x93, 0xb5, 0xbc, 0xc9, 0x5a, 0x64, 0xbd, 0x2e, 0xc6,
	0xd2, 0x0a, 0x12, 0xad, 0x4f, 0x74, 0xa1, 0x9c, 0x19, 0xd3, 0x5b, 0xa1, 0xeb, 0x35, 0x43, 0xa1,
	0xc4, 0x08, 0x8d, 0x05, 0xe9, 0x21, 0xcf, 0x97, 0x71, 0xd7, 0x7c, 0xeb, 0xec, 0x05, 0xf9, 0x19,
	0xf8, 0x08, 0x8d, 0x13, 0x25, 0x3e, 0xa4, 0xaa, 0x50, 0x7a, 0x37, 0xce, 0xb1, 0x0c, 0x1c, 0x46,
	0x20, 0x64, 0xa9, 0xe9, 0xa0, 0x5c, 0x48, 0xca, 0x53, 0x4a, 0x33, 0xbd, 0x13, 0x86, 0x71, 0x29,
	0xaf, 0x0f, 0x49, 0x6c, 0xd1, 0x8f, 0xc8, 0xbb, 0x6c, 0x83, 0x3c, 0x5e, 0x19, 0xe8, 0x44, 0xab,
	0x10, 0x57, 0x33, 0x45, 0xa1, 0x3f, 0x28, 0x31, 0x09, 0x1d, 0xbf, 0x17, 0x24, 0x7f, 0x42, 0x70,
	0xc0, 0xbd, 0xd1, 0xae, 0x1a, 0x39, 0xcd, 0x91, 0xb6, 0x58, 0x87, 0x3c, 0x5b, 0x1d, 0x2b, 0xeb,
	0xc0, 0x38, 0xcf, 0x83, 0x08, 0xc5, 0xa0, 0xa9, 0x69, 0xb9, 0x8b, 0x9c, 0xde, 0x0f, 0x82, 0x38,
	0xc1, 0x58, 0xe2, 0x9c, 0xf6, 0x39, 0x58, 0x47, 0x1f, 0x84, 0xe3, 0xea, 0x7a, 0x0d, 0xb5, 0x43,
	0xe9, 0x93, 0xc2, 0x3a, 0x9d, 0xa1, 0xf1, 0x52, 0xa8, 0x2d, 0xfa, 0x30, 0x8e, 0x60, 0xcc, 0xf3,
	0x63, 0x80, 0x47, 0x61, 0x04, 0x57, 0x19, 0xea, 0x02, 0x7f, 0xc7, 0xfa, 0xe4, 0xe5, 0x87, 0x22,
	0x6b, 0x1a, 0x51, 0x69, 0x6a, 0x7b, 0x28, 0x24, 0x96, 0x57, 0x99, 0xd7, 0x83, 0xa5, 0x35, 0x35,
	0xba, 0xc8, 0xe9, 0x63, 0x76, 0x99, 0xac, 0x85, 0xf6, 0x05, 0x92, 0x81, 0x2e, 0x14, 0xa7, 0x4f,
	0xd8, 0x2f, 0xe4, 0xc7, 0xa6, 0xff, 0x18, 0x2f, 0x1d, 0xa9, 0x13, 0x90, 0xab, 0xd3, 0x13, 0x2f,
	0xdc, 0x9a, 0xae, 0xec, 0xc5, 0x53, 0xb6, 0x49, 0x3a, 0x19, 0x9a, 0xd0, 0xa4, 0x38, 0x6b, 0x20,
	0xa5, 0xde, 0xae, 0xb1, 0xe1, 0x9e, 0x31, 0xc2, 0x6a, 0xe5, 0xa5, 0xb0, 0x21, 0x05, 0x0f, 0xdc,
	0xd2, 0x67, 0xec, 0x3e, 0xb9, 0x23, 0x94, 0x2d, 0x06, 0x03, 0x91, 0x88, 0x70, 0x89, 0x86, 0x78,
	0x84, 0xb2, 0x0e, 0xa4, 0xac, 0xa4, 0x47, 0x37, 0xfa, 0xff, 0x9e, 0x21, 0xad, 0x37, 0xf3, 0x83,
	0xf6, 0xe9, 0xaf, 0x60, 0x9f, 0xad, 0x3c, 0x67, 0x79, 0x78, 0xfb, 0xf2, 0x33, 0xaf, 0x79, 0xe5,
	0x35, 0x9d, 0xbf, 0x9d, 0xcc, 0xa6, 0xed, 0xf9, 0xd1, 0xb4, 0x33, 0xdd, 0x9b, 0x95, 0x2f, 0x63,
	0xfd, 0x00, 0x1f, 0xee, 0x2f, 0xfe, 0xef, 0x3d, 0x7e, 0x11, 0x3f, 0xbf, 0x9f, 0x3d, 0x97, 0x02,
	0xfc, 0x79, 0x76, 0x3d, 0x8d, 0x64, 0xb0, 0xbb, 0x68, 0xc7, 0x65, 0x58, 0x8d, 0xba, 0xed, 0xf2,
	0xc8, 0xc5, 0xdf, 0x35, 0x60, 0x07, 0x76, 0x17, 0x3b, 0x0d, 0x60, 0x67, 0xd4, 0xdd, 0x89, 0x80,
	0x7f, 0xce, 0xb6, 0xe2, 0x6e, 0xaf, 0x07, 0xbb, 0x8b, 0x5e, 0xaf, 0x81, 0xf4, 0x7a, 0xa3, 0x6e,
	0xaf, 0x17, 0x41, 0xbf, 0x9e, 0x2f, 0xa3, 0xdb, 0xfc, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x91,
	0x43, 0x9d, 0x2c, 0x08, 0x00, 0x00,
}

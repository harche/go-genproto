// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/campaign_criterion.proto

package resources

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v2/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
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

// A campaign criterion.
type CampaignCriterion struct {
	// The resource name of the campaign criterion.
	// Campaign criterion resource names have the form:
	//
	// `customers/{customer_id}/campaignCriteria/{campaign_id}~{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The campaign to which the criterion belongs.
	Campaign *wrappers.StringValue `protobuf:"bytes,4,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// The ID of the criterion.
	//
	// This field is ignored during mutate.
	CriterionId *wrappers.Int64Value `protobuf:"bytes,5,opt,name=criterion_id,json=criterionId,proto3" json:"criterion_id,omitempty"`
	// The modifier for the bids when the criterion matches. The modifier must be
	// in the range: 0.1 - 10.0. Most targetable criteria types support modifiers.
	// Use 0 to opt out of a Device type.
	BidModifier *wrappers.FloatValue `protobuf:"bytes,14,opt,name=bid_modifier,json=bidModifier,proto3" json:"bid_modifier,omitempty"`
	// Whether to target (`false`) or exclude (`true`) the criterion.
	Negative *wrappers.BoolValue `protobuf:"bytes,7,opt,name=negative,proto3" json:"negative,omitempty"`
	// The type of the criterion.
	Type enums.CriterionTypeEnum_CriterionType `protobuf:"varint,6,opt,name=type,proto3,enum=google.ads.googleads.v2.enums.CriterionTypeEnum_CriterionType" json:"type,omitempty"`
	// The status of the criterion.
	Status enums.CampaignCriterionStatusEnum_CampaignCriterionStatus `protobuf:"varint,35,opt,name=status,proto3,enum=google.ads.googleads.v2.enums.CampaignCriterionStatusEnum_CampaignCriterionStatus" json:"status,omitempty"`
	// The campaign criterion.
	//
	// Exactly one must be set.
	//
	// Types that are valid to be assigned to Criterion:
	//	*CampaignCriterion_Keyword
	//	*CampaignCriterion_Placement
	//	*CampaignCriterion_MobileAppCategory
	//	*CampaignCriterion_MobileApplication
	//	*CampaignCriterion_Location
	//	*CampaignCriterion_Device
	//	*CampaignCriterion_AdSchedule
	//	*CampaignCriterion_AgeRange
	//	*CampaignCriterion_Gender
	//	*CampaignCriterion_IncomeRange
	//	*CampaignCriterion_ParentalStatus
	//	*CampaignCriterion_UserList
	//	*CampaignCriterion_YoutubeVideo
	//	*CampaignCriterion_YoutubeChannel
	//	*CampaignCriterion_Proximity
	//	*CampaignCriterion_Topic
	//	*CampaignCriterion_ListingScope
	//	*CampaignCriterion_Language
	//	*CampaignCriterion_IpBlock
	//	*CampaignCriterion_ContentLabel
	//	*CampaignCriterion_Carrier
	//	*CampaignCriterion_UserInterest
	//	*CampaignCriterion_Webpage
	//	*CampaignCriterion_OperatingSystemVersion
	//	*CampaignCriterion_MobileDevice
	//	*CampaignCriterion_LocationGroup
	//	*CampaignCriterion_CustomAffinity
	Criterion            isCampaignCriterion_Criterion `protobuf_oneof:"criterion"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CampaignCriterion) Reset()         { *m = CampaignCriterion{} }
func (m *CampaignCriterion) String() string { return proto.CompactTextString(m) }
func (*CampaignCriterion) ProtoMessage()    {}
func (*CampaignCriterion) Descriptor() ([]byte, []int) {
	return fileDescriptor_395bdd66d6311d9f, []int{0}
}

func (m *CampaignCriterion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignCriterion.Unmarshal(m, b)
}
func (m *CampaignCriterion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignCriterion.Marshal(b, m, deterministic)
}
func (m *CampaignCriterion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignCriterion.Merge(m, src)
}
func (m *CampaignCriterion) XXX_Size() int {
	return xxx_messageInfo_CampaignCriterion.Size(m)
}
func (m *CampaignCriterion) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignCriterion.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignCriterion proto.InternalMessageInfo

func (m *CampaignCriterion) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignCriterion) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *CampaignCriterion) GetCriterionId() *wrappers.Int64Value {
	if m != nil {
		return m.CriterionId
	}
	return nil
}

func (m *CampaignCriterion) GetBidModifier() *wrappers.FloatValue {
	if m != nil {
		return m.BidModifier
	}
	return nil
}

func (m *CampaignCriterion) GetNegative() *wrappers.BoolValue {
	if m != nil {
		return m.Negative
	}
	return nil
}

func (m *CampaignCriterion) GetType() enums.CriterionTypeEnum_CriterionType {
	if m != nil {
		return m.Type
	}
	return enums.CriterionTypeEnum_UNSPECIFIED
}

func (m *CampaignCriterion) GetStatus() enums.CampaignCriterionStatusEnum_CampaignCriterionStatus {
	if m != nil {
		return m.Status
	}
	return enums.CampaignCriterionStatusEnum_UNSPECIFIED
}

type isCampaignCriterion_Criterion interface {
	isCampaignCriterion_Criterion()
}

type CampaignCriterion_Keyword struct {
	Keyword *common.KeywordInfo `protobuf:"bytes,8,opt,name=keyword,proto3,oneof"`
}

type CampaignCriterion_Placement struct {
	Placement *common.PlacementInfo `protobuf:"bytes,9,opt,name=placement,proto3,oneof"`
}

type CampaignCriterion_MobileAppCategory struct {
	MobileAppCategory *common.MobileAppCategoryInfo `protobuf:"bytes,10,opt,name=mobile_app_category,json=mobileAppCategory,proto3,oneof"`
}

type CampaignCriterion_MobileApplication struct {
	MobileApplication *common.MobileApplicationInfo `protobuf:"bytes,11,opt,name=mobile_application,json=mobileApplication,proto3,oneof"`
}

type CampaignCriterion_Location struct {
	Location *common.LocationInfo `protobuf:"bytes,12,opt,name=location,proto3,oneof"`
}

type CampaignCriterion_Device struct {
	Device *common.DeviceInfo `protobuf:"bytes,13,opt,name=device,proto3,oneof"`
}

type CampaignCriterion_AdSchedule struct {
	AdSchedule *common.AdScheduleInfo `protobuf:"bytes,15,opt,name=ad_schedule,json=adSchedule,proto3,oneof"`
}

type CampaignCriterion_AgeRange struct {
	AgeRange *common.AgeRangeInfo `protobuf:"bytes,16,opt,name=age_range,json=ageRange,proto3,oneof"`
}

type CampaignCriterion_Gender struct {
	Gender *common.GenderInfo `protobuf:"bytes,17,opt,name=gender,proto3,oneof"`
}

type CampaignCriterion_IncomeRange struct {
	IncomeRange *common.IncomeRangeInfo `protobuf:"bytes,18,opt,name=income_range,json=incomeRange,proto3,oneof"`
}

type CampaignCriterion_ParentalStatus struct {
	ParentalStatus *common.ParentalStatusInfo `protobuf:"bytes,19,opt,name=parental_status,json=parentalStatus,proto3,oneof"`
}

type CampaignCriterion_UserList struct {
	UserList *common.UserListInfo `protobuf:"bytes,22,opt,name=user_list,json=userList,proto3,oneof"`
}

type CampaignCriterion_YoutubeVideo struct {
	YoutubeVideo *common.YouTubeVideoInfo `protobuf:"bytes,20,opt,name=youtube_video,json=youtubeVideo,proto3,oneof"`
}

type CampaignCriterion_YoutubeChannel struct {
	YoutubeChannel *common.YouTubeChannelInfo `protobuf:"bytes,21,opt,name=youtube_channel,json=youtubeChannel,proto3,oneof"`
}

type CampaignCriterion_Proximity struct {
	Proximity *common.ProximityInfo `protobuf:"bytes,23,opt,name=proximity,proto3,oneof"`
}

type CampaignCriterion_Topic struct {
	Topic *common.TopicInfo `protobuf:"bytes,24,opt,name=topic,proto3,oneof"`
}

type CampaignCriterion_ListingScope struct {
	ListingScope *common.ListingScopeInfo `protobuf:"bytes,25,opt,name=listing_scope,json=listingScope,proto3,oneof"`
}

type CampaignCriterion_Language struct {
	Language *common.LanguageInfo `protobuf:"bytes,26,opt,name=language,proto3,oneof"`
}

type CampaignCriterion_IpBlock struct {
	IpBlock *common.IpBlockInfo `protobuf:"bytes,27,opt,name=ip_block,json=ipBlock,proto3,oneof"`
}

type CampaignCriterion_ContentLabel struct {
	ContentLabel *common.ContentLabelInfo `protobuf:"bytes,28,opt,name=content_label,json=contentLabel,proto3,oneof"`
}

type CampaignCriterion_Carrier struct {
	Carrier *common.CarrierInfo `protobuf:"bytes,29,opt,name=carrier,proto3,oneof"`
}

type CampaignCriterion_UserInterest struct {
	UserInterest *common.UserInterestInfo `protobuf:"bytes,30,opt,name=user_interest,json=userInterest,proto3,oneof"`
}

type CampaignCriterion_Webpage struct {
	Webpage *common.WebpageInfo `protobuf:"bytes,31,opt,name=webpage,proto3,oneof"`
}

type CampaignCriterion_OperatingSystemVersion struct {
	OperatingSystemVersion *common.OperatingSystemVersionInfo `protobuf:"bytes,32,opt,name=operating_system_version,json=operatingSystemVersion,proto3,oneof"`
}

type CampaignCriterion_MobileDevice struct {
	MobileDevice *common.MobileDeviceInfo `protobuf:"bytes,33,opt,name=mobile_device,json=mobileDevice,proto3,oneof"`
}

type CampaignCriterion_LocationGroup struct {
	LocationGroup *common.LocationGroupInfo `protobuf:"bytes,34,opt,name=location_group,json=locationGroup,proto3,oneof"`
}

type CampaignCriterion_CustomAffinity struct {
	CustomAffinity *common.CustomAffinityInfo `protobuf:"bytes,36,opt,name=custom_affinity,json=customAffinity,proto3,oneof"`
}

func (*CampaignCriterion_Keyword) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Placement) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_MobileAppCategory) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_MobileApplication) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Location) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Device) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_AdSchedule) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_AgeRange) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Gender) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_IncomeRange) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_ParentalStatus) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_UserList) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_YoutubeVideo) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_YoutubeChannel) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Proximity) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Topic) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_ListingScope) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Language) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_IpBlock) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_ContentLabel) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Carrier) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_UserInterest) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Webpage) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_OperatingSystemVersion) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_MobileDevice) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_LocationGroup) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_CustomAffinity) isCampaignCriterion_Criterion() {}

func (m *CampaignCriterion) GetCriterion() isCampaignCriterion_Criterion {
	if m != nil {
		return m.Criterion
	}
	return nil
}

func (m *CampaignCriterion) GetKeyword() *common.KeywordInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Keyword); ok {
		return x.Keyword
	}
	return nil
}

func (m *CampaignCriterion) GetPlacement() *common.PlacementInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Placement); ok {
		return x.Placement
	}
	return nil
}

func (m *CampaignCriterion) GetMobileAppCategory() *common.MobileAppCategoryInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_MobileAppCategory); ok {
		return x.MobileAppCategory
	}
	return nil
}

func (m *CampaignCriterion) GetMobileApplication() *common.MobileApplicationInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_MobileApplication); ok {
		return x.MobileApplication
	}
	return nil
}

func (m *CampaignCriterion) GetLocation() *common.LocationInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Location); ok {
		return x.Location
	}
	return nil
}

func (m *CampaignCriterion) GetDevice() *common.DeviceInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Device); ok {
		return x.Device
	}
	return nil
}

func (m *CampaignCriterion) GetAdSchedule() *common.AdScheduleInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_AdSchedule); ok {
		return x.AdSchedule
	}
	return nil
}

func (m *CampaignCriterion) GetAgeRange() *common.AgeRangeInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_AgeRange); ok {
		return x.AgeRange
	}
	return nil
}

func (m *CampaignCriterion) GetGender() *common.GenderInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Gender); ok {
		return x.Gender
	}
	return nil
}

func (m *CampaignCriterion) GetIncomeRange() *common.IncomeRangeInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_IncomeRange); ok {
		return x.IncomeRange
	}
	return nil
}

func (m *CampaignCriterion) GetParentalStatus() *common.ParentalStatusInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_ParentalStatus); ok {
		return x.ParentalStatus
	}
	return nil
}

func (m *CampaignCriterion) GetUserList() *common.UserListInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_UserList); ok {
		return x.UserList
	}
	return nil
}

func (m *CampaignCriterion) GetYoutubeVideo() *common.YouTubeVideoInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_YoutubeVideo); ok {
		return x.YoutubeVideo
	}
	return nil
}

func (m *CampaignCriterion) GetYoutubeChannel() *common.YouTubeChannelInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_YoutubeChannel); ok {
		return x.YoutubeChannel
	}
	return nil
}

func (m *CampaignCriterion) GetProximity() *common.ProximityInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Proximity); ok {
		return x.Proximity
	}
	return nil
}

func (m *CampaignCriterion) GetTopic() *common.TopicInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Topic); ok {
		return x.Topic
	}
	return nil
}

func (m *CampaignCriterion) GetListingScope() *common.ListingScopeInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_ListingScope); ok {
		return x.ListingScope
	}
	return nil
}

func (m *CampaignCriterion) GetLanguage() *common.LanguageInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Language); ok {
		return x.Language
	}
	return nil
}

func (m *CampaignCriterion) GetIpBlock() *common.IpBlockInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_IpBlock); ok {
		return x.IpBlock
	}
	return nil
}

func (m *CampaignCriterion) GetContentLabel() *common.ContentLabelInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_ContentLabel); ok {
		return x.ContentLabel
	}
	return nil
}

func (m *CampaignCriterion) GetCarrier() *common.CarrierInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Carrier); ok {
		return x.Carrier
	}
	return nil
}

func (m *CampaignCriterion) GetUserInterest() *common.UserInterestInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_UserInterest); ok {
		return x.UserInterest
	}
	return nil
}

func (m *CampaignCriterion) GetWebpage() *common.WebpageInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Webpage); ok {
		return x.Webpage
	}
	return nil
}

func (m *CampaignCriterion) GetOperatingSystemVersion() *common.OperatingSystemVersionInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_OperatingSystemVersion); ok {
		return x.OperatingSystemVersion
	}
	return nil
}

func (m *CampaignCriterion) GetMobileDevice() *common.MobileDeviceInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_MobileDevice); ok {
		return x.MobileDevice
	}
	return nil
}

func (m *CampaignCriterion) GetLocationGroup() *common.LocationGroupInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_LocationGroup); ok {
		return x.LocationGroup
	}
	return nil
}

func (m *CampaignCriterion) GetCustomAffinity() *common.CustomAffinityInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_CustomAffinity); ok {
		return x.CustomAffinity
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignCriterion) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignCriterion_Keyword)(nil),
		(*CampaignCriterion_Placement)(nil),
		(*CampaignCriterion_MobileAppCategory)(nil),
		(*CampaignCriterion_MobileApplication)(nil),
		(*CampaignCriterion_Location)(nil),
		(*CampaignCriterion_Device)(nil),
		(*CampaignCriterion_AdSchedule)(nil),
		(*CampaignCriterion_AgeRange)(nil),
		(*CampaignCriterion_Gender)(nil),
		(*CampaignCriterion_IncomeRange)(nil),
		(*CampaignCriterion_ParentalStatus)(nil),
		(*CampaignCriterion_UserList)(nil),
		(*CampaignCriterion_YoutubeVideo)(nil),
		(*CampaignCriterion_YoutubeChannel)(nil),
		(*CampaignCriterion_Proximity)(nil),
		(*CampaignCriterion_Topic)(nil),
		(*CampaignCriterion_ListingScope)(nil),
		(*CampaignCriterion_Language)(nil),
		(*CampaignCriterion_IpBlock)(nil),
		(*CampaignCriterion_ContentLabel)(nil),
		(*CampaignCriterion_Carrier)(nil),
		(*CampaignCriterion_UserInterest)(nil),
		(*CampaignCriterion_Webpage)(nil),
		(*CampaignCriterion_OperatingSystemVersion)(nil),
		(*CampaignCriterion_MobileDevice)(nil),
		(*CampaignCriterion_LocationGroup)(nil),
		(*CampaignCriterion_CustomAffinity)(nil),
	}
}

func init() {
	proto.RegisterType((*CampaignCriterion)(nil), "google.ads.googleads.v2.resources.CampaignCriterion")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/campaign_criterion.proto", fileDescriptor_395bdd66d6311d9f)
}

var fileDescriptor_395bdd66d6311d9f = []byte{
	// 1141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x97, 0x6f, 0x73, 0xdb, 0x34,
	0x18, 0xc0, 0x69, 0xd9, 0xba, 0x54, 0x49, 0x3b, 0xa6, 0xb1, 0x22, 0xba, 0x32, 0xda, 0x8d, 0xdd,
	0x95, 0x3f, 0x73, 0x20, 0xc0, 0x8e, 0x0b, 0xc7, 0xee, 0xd2, 0x0c, 0xba, 0x6c, 0x2d, 0x94, 0xb4,
	0x64, 0xc7, 0xae, 0x9c, 0x4f, 0xb1, 0x9f, 0x78, 0x62, 0xb6, 0xe4, 0x93, 0xe5, 0x94, 0xbc, 0xe6,
	0x9b, 0xf0, 0x92, 0x8f, 0xc2, 0xd7, 0xe0, 0x1d, 0x9f, 0x82, 0x93, 0x2c, 0xb9, 0xe9, 0xda, 0x62,
	0xf3, 0xce, 0x7e, 0xfc, 0xfc, 0x7e, 0x7a, 0xf4, 0xc4, 0x92, 0x1c, 0xd4, 0x8d, 0x84, 0x88, 0x62,
	0x68, 0xd3, 0x30, 0x6b, 0x17, 0x97, 0xfa, 0x6a, 0xda, 0x69, 0x4b, 0xc8, 0x44, 0x2e, 0x03, 0xc8,
	0xda, 0x01, 0x4d, 0x52, 0xca, 0x22, 0xee, 0x07, 0x92, 0x29, 0x90, 0x4c, 0x70, 0x2f, 0x95, 0x42,
	0x09, 0xbc, 0x55, 0x00, 0x1e, 0x0d, 0x33, 0xaf, 0x64, 0xbd, 0x69, 0xc7, 0x2b, 0xd9, 0xf5, 0x07,
	0x97, 0xe9, 0x03, 0x91, 0x24, 0x82, 0xb7, 0xad, 0x92, 0x16, 0xc6, 0xf5, 0x6f, 0x2e, 0x4b, 0x07,
	0x9e, 0x27, 0x17, 0x55, 0xe2, 0x67, 0x8a, 0xaa, 0x3c, 0xb3, 0x78, 0xa7, 0x02, 0x2f, 0x29, 0x35,
	0x4b, 0xc1, 0x32, 0x77, 0x2c, 0x63, 0xee, 0xc6, 0xf9, 0xa4, 0x7d, 0x22, 0x69, 0x9a, 0x82, 0x74,
	0xce, 0x0d, 0xe7, 0x4c, 0x59, 0x9b, 0x72, 0x2e, 0x14, 0x55, 0x4c, 0x70, 0xfb, 0xf4, 0xee, 0xdf,
	0xb7, 0xd0, 0x8d, 0xbe, 0xad, 0xaa, 0xef, 0xf4, 0xf8, 0x1e, 0x5a, 0x71, 0x2d, 0xf0, 0x39, 0x4d,
	0x80, 0x2c, 0x6c, 0x2e, 0x6c, 0x2f, 0x0f, 0x5b, 0x2e, 0xf8, 0x3d, 0x4d, 0x00, 0x7f, 0x85, 0x1a,
	0x6e, 0x3e, 0xe4, 0xca, 0xe6, 0xc2, 0x76, 0xb3, 0xb3, 0x61, 0xbb, 0xe8, 0xb9, 0x5a, 0xbc, 0x43,
	0x25, 0x19, 0x8f, 0x46, 0x34, 0xce, 0x61, 0x58, 0x66, 0xe3, 0x47, 0xa8, 0x75, 0x3a, 0x15, 0x16,
	0x92, 0xab, 0x86, 0xbe, 0x7d, 0x8e, 0x1e, 0x70, 0xf5, 0xf0, 0x8b, 0x02, 0x6e, 0x96, 0xc0, 0x20,
	0xd4, 0xfc, 0x98, 0x85, 0x7e, 0x22, 0x42, 0x36, 0x61, 0x20, 0xc9, 0xea, 0x25, 0xfc, 0x77, 0xb1,
	0xa0, 0xca, 0xf2, 0x63, 0x16, 0xee, 0xdb, 0x7c, 0xfc, 0x10, 0x35, 0x38, 0x44, 0x54, 0xb1, 0x29,
	0x90, 0x6b, 0x86, 0x5d, 0x3f, 0xc7, 0xee, 0x08, 0x11, 0xdb, 0xba, 0x5d, 0x2e, 0x1e, 0xa2, 0x2b,
	0xba, 0xf1, 0x64, 0x69, 0x73, 0x61, 0x7b, 0xb5, 0xf3, 0xc8, 0xbb, 0xec, 0xf5, 0x31, 0xbf, 0x96,
	0x57, 0xb6, 0xf3, 0x68, 0x96, 0xc2, 0xb7, 0x3c, 0x4f, 0xce, 0x46, 0x86, 0xc6, 0x85, 0x7f, 0x45,
	0x4b, 0xc5, 0x2b, 0x40, 0xee, 0x19, 0xeb, 0xb0, 0xca, 0xfa, 0xfa, 0x8f, 0x75, 0x68, 0xe8, 0xc2,
	0x7f, 0xf1, 0xb3, 0xa1, 0x1d, 0x01, 0xef, 0xa2, 0x6b, 0xaf, 0x60, 0x76, 0x22, 0x64, 0x48, 0x1a,
	0x66, 0xda, 0x1f, 0x5f, 0x3a, 0x58, 0xf1, 0x7a, 0x7b, 0xcf, 0x8a, 0xf4, 0x01, 0x9f, 0x88, 0x27,
	0x6f, 0x0c, 0x1d, 0x8d, 0xf7, 0xd1, 0x72, 0x1a, 0xd3, 0x00, 0x12, 0xe0, 0x8a, 0x2c, 0x1b, 0xd5,
	0x83, 0x2a, 0xd5, 0x81, 0x03, 0xac, 0xec, 0xd4, 0x80, 0x23, 0x74, 0x33, 0x11, 0x63, 0x16, 0x83,
	0x4f, 0xd3, 0xd4, 0x0f, 0xa8, 0x82, 0x48, 0xc8, 0x19, 0x41, 0x46, 0xfc, 0x65, 0x95, 0x78, 0xdf,
	0xa0, 0xbd, 0x34, 0xed, 0x5b, 0xd0, 0x0e, 0x70, 0x23, 0x79, 0xfd, 0x01, 0x9e, 0x20, 0x7c, 0x3a,
	0x50, 0xcc, 0x02, 0xb3, 0x14, 0x48, 0xf3, 0x7f, 0x8e, 0xe3, 0xc0, 0x73, 0xe3, 0xb8, 0x07, 0xf8,
	0x29, 0x6a, 0xc4, 0xc2, 0xda, 0x5b, 0xc6, 0xfe, 0x49, 0x95, 0x7d, 0x4f, 0x9c, 0x91, 0x96, 0x3c,
	0x7e, 0x8c, 0x96, 0x42, 0x98, 0xb2, 0x00, 0xc8, 0x8a, 0x31, 0x7d, 0x54, 0x65, 0x7a, 0x6c, 0xb2,
	0xad, 0xc7, 0xb2, 0xf8, 0x47, 0xd4, 0xa4, 0xa1, 0x9f, 0x05, 0x2f, 0x21, 0xcc, 0x63, 0x20, 0xd7,
	0x8d, 0xca, 0xab, 0x52, 0xf5, 0xc2, 0x43, 0x4b, 0x58, 0x1d, 0xa2, 0x65, 0x04, 0x3f, 0x43, 0xcb,
	0x34, 0x02, 0x5f, 0x52, 0x1e, 0x01, 0x79, 0xab, 0xde, 0x2c, 0x7b, 0x11, 0x0c, 0x75, 0xbe, 0x9b,
	0x25, 0xb5, 0xf7, 0x7a, 0x96, 0x11, 0xf0, 0x10, 0x24, 0xb9, 0x51, 0x6f, 0x96, 0xbb, 0x26, 0xdb,
	0xcd, 0xb2, 0x60, 0xf1, 0x11, 0x6a, 0x31, 0x1e, 0x88, 0xc4, 0x55, 0x85, 0x8d, 0xab, 0x5d, 0xe5,
	0x1a, 0x18, 0x66, 0xbe, 0xb0, 0x26, 0x3b, 0x0d, 0xe1, 0x5f, 0xd0, 0xf5, 0x94, 0x4a, 0xe0, 0x8a,
	0xc6, 0x76, 0xbb, 0x26, 0x37, 0x8d, 0xb8, 0x53, 0xf9, 0xce, 0x5b, 0xac, 0x58, 0x87, 0xd6, 0xbd,
	0x9a, 0x9e, 0x89, 0xea, 0x3e, 0xe6, 0x19, 0x48, 0x3f, 0x66, 0x99, 0x22, 0x6b, 0xf5, 0xfa, 0xf8,
	0x53, 0x06, 0x72, 0x8f, 0x65, 0x6e, 0x2d, 0x35, 0x72, 0x7b, 0x8f, 0x9f, 0xa3, 0x95, 0x99, 0xc8,
	0x55, 0x3e, 0x06, 0x7f, 0xca, 0x42, 0x10, 0xe4, 0x6d, 0x23, 0xfc, 0xb4, 0x4a, 0xf8, 0xb3, 0xc8,
	0x8f, 0xf2, 0x31, 0x8c, 0x34, 0x63, 0xa5, 0x2d, 0x2b, 0x32, 0x31, 0xdd, 0x04, 0x27, 0x0e, 0x5e,
	0x52, 0xce, 0x21, 0x26, 0xb7, 0xea, 0x35, 0xc1, 0xaa, 0xfb, 0x05, 0xe5, 0x9a, 0x60, 0x65, 0x36,
	0x6a, 0x76, 0x14, 0x29, 0x7e, 0x63, 0x09, 0x53, 0x33, 0xf2, 0x4e, 0xcd, 0x1d, 0xc5, 0x01, 0xe5,
	0x8e, 0xe2, 0x02, 0xb8, 0x87, 0xae, 0x2a, 0x91, 0xb2, 0x80, 0x10, 0xa3, 0xfa, 0xb0, 0x4a, 0x75,
	0xa4, 0x93, 0xad, 0xa6, 0x20, 0x75, 0x27, 0xf5, 0x2f, 0xc2, 0x78, 0xe4, 0x67, 0x81, 0x48, 0x81,
	0xbc, 0x5b, 0xaf, 0x93, 0x7b, 0x05, 0x74, 0xa8, 0x19, 0xd7, 0xc9, 0x78, 0x2e, 0x66, 0x36, 0x07,
	0xca, 0xa3, 0x9c, 0x46, 0x40, 0xd6, 0x6b, 0x6e, 0x0e, 0x36, 0xbf, 0xdc, 0x1c, 0xec, 0x3d, 0x7e,
	0x82, 0x1a, 0x2c, 0xf5, 0xc7, 0xb1, 0x08, 0x5e, 0x91, 0xdb, 0xf5, 0xb6, 0xf4, 0x41, 0xba, 0xa3,
	0xd3, 0xdd, 0x96, 0xce, 0x8a, 0x5b, 0x3d, 0xdd, 0x40, 0x70, 0x05, 0x5c, 0xf9, 0x31, 0x1d, 0x43,
	0x4c, 0x36, 0xea, 0x4d, 0xb7, 0x5f, 0x40, 0x7b, 0x9a, 0x71, 0xd3, 0x0d, 0xe6, 0x62, 0xfa, 0xd0,
	0x09, 0xa8, 0x94, 0xfa, 0x9c, 0x7e, 0xaf, 0x5e, 0x85, 0xfd, 0x22, 0xdd, 0x55, 0x68, 0x69, 0x5d,
	0xa1, 0x59, 0x27, 0x8c, 0x2b, 0x90, 0x90, 0x29, 0x72, 0xa7, 0x5e, 0x85, 0x7a, 0xad, 0x0c, 0x2c,
	0xe3, 0x2a, 0xcc, 0xe7, 0x62, 0xba, 0xc2, 0x13, 0x18, 0xa7, 0xfa, 0xf7, 0x78, 0xbf, 0x5e, 0x85,
	0xcf, 0x8b, 0x74, 0x57, 0xa1, 0xa5, 0xf1, 0x14, 0x11, 0x91, 0x82, 0xa4, 0xc5, 0x4b, 0x33, 0xcb,
	0x14, 0x24, 0xfe, 0x14, 0x64, 0xa6, 0x8f, 0x81, 0x4d, 0x63, 0xee, 0x56, 0x99, 0x7f, 0x70, 0xfc,
	0xa1, 0xc1, 0x47, 0x05, 0x6d, 0x07, 0x5a, 0x13, 0x17, 0x3e, 0xd5, 0x9d, 0xb1, 0xc7, 0x9a, 0x3d,
	0x29, 0xb6, 0xea, 0x75, 0xa6, 0x38, 0xd1, 0xce, 0x9c, 0x17, 0xad, 0x64, 0x2e, 0x86, 0x5f, 0xa0,
	0x55, 0x77, 0x0e, 0xf9, 0x91, 0x14, 0x79, 0x4a, 0xee, 0x1a, 0xf3, 0x67, 0x75, 0x4f, 0xb3, 0x5d,
	0x0d, 0x59, 0xf5, 0x4a, 0x3c, 0x1f, 0xd4, 0x1b, 0x4a, 0x90, 0x67, 0x4a, 0x24, 0x3e, 0x9d, 0x4c,
	0x18, 0xd7, 0xeb, 0xfe, 0x83, 0x7a, 0x1b, 0x4a, 0xdf, 0x60, 0x3d, 0x4b, 0xb9, 0x0d, 0x25, 0x38,
	0x13, 0xdd, 0x69, 0xa2, 0xe5, 0xf2, 0x93, 0x71, 0xe7, 0xf7, 0x45, 0x74, 0x3f, 0x10, 0x89, 0x57,
	0xf9, 0xbd, 0xbf, 0xb3, 0x76, 0xee, 0x1b, 0xea, 0x40, 0x7f, 0x11, 0x1e, 0x2c, 0xbc, 0x78, 0x6a,
	0xe1, 0x48, 0xe8, 0xd5, 0xe7, 0x09, 0x19, 0xb5, 0x23, 0xe0, 0xe6, 0x7b, 0xd1, 0x7d, 0xab, 0xa7,
	0x2c, 0xfb, 0x8f, 0xff, 0x21, 0x5f, 0x97, 0x57, 0x7f, 0x2c, 0xbe, 0xb9, 0xdb, 0xeb, 0xfd, 0xb9,
	0xb8, 0xb5, 0x5b, 0x28, 0x7b, 0x61, 0xe6, 0x15, 0x97, 0xfa, 0x6a, 0xd4, 0xf1, 0x86, 0x2e, 0xf3,
	0x2f, 0x97, 0x73, 0xdc, 0x0b, 0xb3, 0xe3, 0x32, 0xe7, 0x78, 0xd4, 0x39, 0x2e, 0x73, 0xfe, 0x59,
	0xbc, 0x5f, 0x3c, 0xe8, 0x76, 0x7b, 0x61, 0xd6, 0xed, 0x96, 0x59, 0xdd, 0xee, 0xa8, 0xd3, 0xed,
	0x96, 0x79, 0xe3, 0x25, 0x53, 0xec, 0xe7, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x62, 0xbd, 0x4a,
	0xa0, 0x33, 0x0d, 0x00, 0x00,
}

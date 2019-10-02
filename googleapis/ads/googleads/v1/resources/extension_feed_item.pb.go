// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/extension_feed_item.proto

package resources

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
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

// An extension feed item.
type ExtensionFeedItem struct {
	// The resource name of the extension feed item.
	// Extension feed item resource names have the form:
	//
	// `customers/{customer_id}/extensionFeedItems/{feed_item_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The extension type of the extension feed item.
	// This field is read-only.
	ExtensionType enums.ExtensionTypeEnum_ExtensionType `protobuf:"varint,13,opt,name=extension_type,json=extensionType,proto3,enum=google.ads.googleads.v1.enums.ExtensionTypeEnum_ExtensionType" json:"extension_type,omitempty"`
	// Start time in which this feed item is effective and can begin serving.
	// The format is "YYYY-MM-DD HH:MM:SS".
	// Examples: "2018-03-05 09:15:00" or "2018-02-01 14:34:30"
	StartDateTime *wrappers.StringValue `protobuf:"bytes,5,opt,name=start_date_time,json=startDateTime,proto3" json:"start_date_time,omitempty"`
	// End time in which this feed item is no longer effective and will stop
	// serving.
	// The format is "YYYY-MM-DD HH:MM:SS".
	// Examples: "2018-03-05 09:15:00" or "2018-02-01 14:34:30"
	EndDateTime *wrappers.StringValue `protobuf:"bytes,6,opt,name=end_date_time,json=endDateTime,proto3" json:"end_date_time,omitempty"`
	// List of non-overlapping schedules specifying all time intervals
	// for which the feed item may serve. There can be a maximum of 6 schedules
	// per day.
	AdSchedules []*common.AdScheduleInfo `protobuf:"bytes,16,rep,name=ad_schedules,json=adSchedules,proto3" json:"ad_schedules,omitempty"`
	// The targeted device.
	Device enums.FeedItemTargetDeviceEnum_FeedItemTargetDevice `protobuf:"varint,17,opt,name=device,proto3,enum=google.ads.googleads.v1.enums.FeedItemTargetDeviceEnum_FeedItemTargetDevice" json:"device,omitempty"`
	// The targeted geo target constant.
	TargetedGeoTargetConstant *wrappers.StringValue `protobuf:"bytes,20,opt,name=targeted_geo_target_constant,json=targetedGeoTargetConstant,proto3" json:"targeted_geo_target_constant,omitempty"`
	// Status of the feed item.
	// This field is read-only.
	Status enums.FeedItemStatusEnum_FeedItemStatus `protobuf:"varint,4,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.FeedItemStatusEnum_FeedItemStatus" json:"status,omitempty"`
	// Extension type.
	//
	// Types that are valid to be assigned to Extension:
	//	*ExtensionFeedItem_SitelinkFeedItem
	//	*ExtensionFeedItem_StructuredSnippetFeedItem
	//	*ExtensionFeedItem_AppFeedItem
	//	*ExtensionFeedItem_CallFeedItem
	//	*ExtensionFeedItem_CalloutFeedItem
	//	*ExtensionFeedItem_TextMessageFeedItem
	//	*ExtensionFeedItem_PriceFeedItem
	//	*ExtensionFeedItem_PromotionFeedItem
	//	*ExtensionFeedItem_LocationFeedItem
	//	*ExtensionFeedItem_AffiliateLocationFeedItem
	Extension isExtensionFeedItem_Extension `protobuf_oneof:"extension"`
	// Targeting at either the campaign or ad group level. Feed items that target
	// a campaign or ad group will only serve with that resource.
	//
	// Types that are valid to be assigned to ServingResourceTargeting:
	//	*ExtensionFeedItem_TargetedCampaign
	//	*ExtensionFeedItem_TargetedAdGroup
	ServingResourceTargeting isExtensionFeedItem_ServingResourceTargeting `protobuf_oneof:"serving_resource_targeting"`
	XXX_NoUnkeyedLiteral     struct{}                                     `json:"-"`
	XXX_unrecognized         []byte                                       `json:"-"`
	XXX_sizecache            int32                                        `json:"-"`
}

func (m *ExtensionFeedItem) Reset()         { *m = ExtensionFeedItem{} }
func (m *ExtensionFeedItem) String() string { return proto.CompactTextString(m) }
func (*ExtensionFeedItem) ProtoMessage()    {}
func (*ExtensionFeedItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_800f8a839b1b4536, []int{0}
}

func (m *ExtensionFeedItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtensionFeedItem.Unmarshal(m, b)
}
func (m *ExtensionFeedItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtensionFeedItem.Marshal(b, m, deterministic)
}
func (m *ExtensionFeedItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionFeedItem.Merge(m, src)
}
func (m *ExtensionFeedItem) XXX_Size() int {
	return xxx_messageInfo_ExtensionFeedItem.Size(m)
}
func (m *ExtensionFeedItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionFeedItem.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionFeedItem proto.InternalMessageInfo

func (m *ExtensionFeedItem) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ExtensionFeedItem) GetExtensionType() enums.ExtensionTypeEnum_ExtensionType {
	if m != nil {
		return m.ExtensionType
	}
	return enums.ExtensionTypeEnum_UNSPECIFIED
}

func (m *ExtensionFeedItem) GetStartDateTime() *wrappers.StringValue {
	if m != nil {
		return m.StartDateTime
	}
	return nil
}

func (m *ExtensionFeedItem) GetEndDateTime() *wrappers.StringValue {
	if m != nil {
		return m.EndDateTime
	}
	return nil
}

func (m *ExtensionFeedItem) GetAdSchedules() []*common.AdScheduleInfo {
	if m != nil {
		return m.AdSchedules
	}
	return nil
}

func (m *ExtensionFeedItem) GetDevice() enums.FeedItemTargetDeviceEnum_FeedItemTargetDevice {
	if m != nil {
		return m.Device
	}
	return enums.FeedItemTargetDeviceEnum_UNSPECIFIED
}

func (m *ExtensionFeedItem) GetTargetedGeoTargetConstant() *wrappers.StringValue {
	if m != nil {
		return m.TargetedGeoTargetConstant
	}
	return nil
}

func (m *ExtensionFeedItem) GetStatus() enums.FeedItemStatusEnum_FeedItemStatus {
	if m != nil {
		return m.Status
	}
	return enums.FeedItemStatusEnum_UNSPECIFIED
}

type isExtensionFeedItem_Extension interface {
	isExtensionFeedItem_Extension()
}

type ExtensionFeedItem_SitelinkFeedItem struct {
	SitelinkFeedItem *common.SitelinkFeedItem `protobuf:"bytes,2,opt,name=sitelink_feed_item,json=sitelinkFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_StructuredSnippetFeedItem struct {
	StructuredSnippetFeedItem *common.StructuredSnippetFeedItem `protobuf:"bytes,3,opt,name=structured_snippet_feed_item,json=structuredSnippetFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_AppFeedItem struct {
	AppFeedItem *common.AppFeedItem `protobuf:"bytes,7,opt,name=app_feed_item,json=appFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_CallFeedItem struct {
	CallFeedItem *common.CallFeedItem `protobuf:"bytes,8,opt,name=call_feed_item,json=callFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_CalloutFeedItem struct {
	CalloutFeedItem *common.CalloutFeedItem `protobuf:"bytes,9,opt,name=callout_feed_item,json=calloutFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_TextMessageFeedItem struct {
	TextMessageFeedItem *common.TextMessageFeedItem `protobuf:"bytes,10,opt,name=text_message_feed_item,json=textMessageFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_PriceFeedItem struct {
	PriceFeedItem *common.PriceFeedItem `protobuf:"bytes,11,opt,name=price_feed_item,json=priceFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_PromotionFeedItem struct {
	PromotionFeedItem *common.PromotionFeedItem `protobuf:"bytes,12,opt,name=promotion_feed_item,json=promotionFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_LocationFeedItem struct {
	LocationFeedItem *common.LocationFeedItem `protobuf:"bytes,14,opt,name=location_feed_item,json=locationFeedItem,proto3,oneof"`
}

type ExtensionFeedItem_AffiliateLocationFeedItem struct {
	AffiliateLocationFeedItem *common.AffiliateLocationFeedItem `protobuf:"bytes,15,opt,name=affiliate_location_feed_item,json=affiliateLocationFeedItem,proto3,oneof"`
}

func (*ExtensionFeedItem_SitelinkFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_StructuredSnippetFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_AppFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_CallFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_CalloutFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_TextMessageFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_PriceFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_PromotionFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_LocationFeedItem) isExtensionFeedItem_Extension() {}

func (*ExtensionFeedItem_AffiliateLocationFeedItem) isExtensionFeedItem_Extension() {}

func (m *ExtensionFeedItem) GetExtension() isExtensionFeedItem_Extension {
	if m != nil {
		return m.Extension
	}
	return nil
}

func (m *ExtensionFeedItem) GetSitelinkFeedItem() *common.SitelinkFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_SitelinkFeedItem); ok {
		return x.SitelinkFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetStructuredSnippetFeedItem() *common.StructuredSnippetFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_StructuredSnippetFeedItem); ok {
		return x.StructuredSnippetFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetAppFeedItem() *common.AppFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_AppFeedItem); ok {
		return x.AppFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetCallFeedItem() *common.CallFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_CallFeedItem); ok {
		return x.CallFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetCalloutFeedItem() *common.CalloutFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_CalloutFeedItem); ok {
		return x.CalloutFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetTextMessageFeedItem() *common.TextMessageFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_TextMessageFeedItem); ok {
		return x.TextMessageFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetPriceFeedItem() *common.PriceFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_PriceFeedItem); ok {
		return x.PriceFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetPromotionFeedItem() *common.PromotionFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_PromotionFeedItem); ok {
		return x.PromotionFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetLocationFeedItem() *common.LocationFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_LocationFeedItem); ok {
		return x.LocationFeedItem
	}
	return nil
}

func (m *ExtensionFeedItem) GetAffiliateLocationFeedItem() *common.AffiliateLocationFeedItem {
	if x, ok := m.GetExtension().(*ExtensionFeedItem_AffiliateLocationFeedItem); ok {
		return x.AffiliateLocationFeedItem
	}
	return nil
}

type isExtensionFeedItem_ServingResourceTargeting interface {
	isExtensionFeedItem_ServingResourceTargeting()
}

type ExtensionFeedItem_TargetedCampaign struct {
	TargetedCampaign *wrappers.StringValue `protobuf:"bytes,18,opt,name=targeted_campaign,json=targetedCampaign,proto3,oneof"`
}

type ExtensionFeedItem_TargetedAdGroup struct {
	TargetedAdGroup *wrappers.StringValue `protobuf:"bytes,19,opt,name=targeted_ad_group,json=targetedAdGroup,proto3,oneof"`
}

func (*ExtensionFeedItem_TargetedCampaign) isExtensionFeedItem_ServingResourceTargeting() {}

func (*ExtensionFeedItem_TargetedAdGroup) isExtensionFeedItem_ServingResourceTargeting() {}

func (m *ExtensionFeedItem) GetServingResourceTargeting() isExtensionFeedItem_ServingResourceTargeting {
	if m != nil {
		return m.ServingResourceTargeting
	}
	return nil
}

func (m *ExtensionFeedItem) GetTargetedCampaign() *wrappers.StringValue {
	if x, ok := m.GetServingResourceTargeting().(*ExtensionFeedItem_TargetedCampaign); ok {
		return x.TargetedCampaign
	}
	return nil
}

func (m *ExtensionFeedItem) GetTargetedAdGroup() *wrappers.StringValue {
	if x, ok := m.GetServingResourceTargeting().(*ExtensionFeedItem_TargetedAdGroup); ok {
		return x.TargetedAdGroup
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExtensionFeedItem) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExtensionFeedItem_SitelinkFeedItem)(nil),
		(*ExtensionFeedItem_StructuredSnippetFeedItem)(nil),
		(*ExtensionFeedItem_AppFeedItem)(nil),
		(*ExtensionFeedItem_CallFeedItem)(nil),
		(*ExtensionFeedItem_CalloutFeedItem)(nil),
		(*ExtensionFeedItem_TextMessageFeedItem)(nil),
		(*ExtensionFeedItem_PriceFeedItem)(nil),
		(*ExtensionFeedItem_PromotionFeedItem)(nil),
		(*ExtensionFeedItem_LocationFeedItem)(nil),
		(*ExtensionFeedItem_AffiliateLocationFeedItem)(nil),
		(*ExtensionFeedItem_TargetedCampaign)(nil),
		(*ExtensionFeedItem_TargetedAdGroup)(nil),
	}
}

func init() {
	proto.RegisterType((*ExtensionFeedItem)(nil), "google.ads.googleads.v1.resources.ExtensionFeedItem")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/extension_feed_item.proto", fileDescriptor_800f8a839b1b4536)
}

var fileDescriptor_800f8a839b1b4536 = []byte{
	// 895 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0x51, 0x6f, 0x1c, 0x35,
	0x10, 0xe6, 0x12, 0x08, 0xc4, 0x97, 0xbb, 0xcb, 0x39, 0xa8, 0xda, 0x44, 0x27, 0x94, 0x82, 0x2a,
	0x45, 0x82, 0xee, 0x72, 0x29, 0x2f, 0x5c, 0x24, 0xd4, 0x4b, 0x52, 0xd2, 0x96, 0x82, 0xd2, 0xbb,
	0x53, 0x40, 0x28, 0xd5, 0xe2, 0xae, 0xe7, 0x16, 0xc3, 0xae, 0x6d, 0xad, 0xbd, 0x21, 0x95, 0x78,
	0xe3, 0x9f, 0xf0, 0xc8, 0x4f, 0xe1, 0x85, 0xff, 0xc1, 0xaf, 0xa8, 0xd6, 0xbb, 0xde, 0xfa, 0xd2,
	0x5c, 0xf7, 0xde, 0xec, 0x99, 0xf9, 0xbe, 0x6f, 0x66, 0x76, 0x6c, 0x2f, 0x3a, 0x8a, 0x85, 0x88,
	0x13, 0x08, 0x08, 0x55, 0x41, 0xb9, 0x2c, 0x56, 0x57, 0xc3, 0x20, 0x03, 0x25, 0xf2, 0x2c, 0x02,
	0x15, 0xc0, 0xb5, 0x06, 0xae, 0x98, 0xe0, 0xe1, 0x1c, 0x80, 0x86, 0x4c, 0x43, 0xea, 0xcb, 0x4c,
	0x68, 0x81, 0xef, 0x96, 0x08, 0x9f, 0x50, 0xe5, 0xd7, 0x60, 0xff, 0x6a, 0xe8, 0xd7, 0xe0, 0xbd,
	0xfb, 0xcb, 0xf8, 0x23, 0x91, 0xa6, 0x82, 0x07, 0x51, 0xc6, 0x34, 0x64, 0x8c, 0x94, 0x8c, 0x7b,
	0x41, 0x43, 0x78, 0x9d, 0x8b, 0xaa, 0x00, 0x87, 0xcb, 0x00, 0xc0, 0xf3, 0xd4, 0xcd, 0x5d, 0xbf,
	0x92, 0x50, 0x61, 0xbe, 0x7a, 0x37, 0xa6, 0xae, 0x32, 0x54, 0x9a, 0xe8, 0xdc, 0x2a, 0x1d, 0xad,
	0x8a, 0xd2, 0x24, 0x8b, 0x41, 0x87, 0x14, 0xae, 0x58, 0x64, 0x25, 0x07, 0x16, 0x2c, 0x59, 0x40,
	0x38, 0x17, 0x9a, 0x68, 0xa7, 0x88, 0x4f, 0x2a, 0xaf, 0xd9, 0xbd, 0xcc, 0xe7, 0xc1, 0x1f, 0x19,
	0x91, 0x12, 0xb2, 0xca, 0xff, 0xe9, 0x7f, 0x5d, 0xd4, 0x7f, 0x64, 0x2b, 0xf9, 0x16, 0x80, 0x3e,
	0xd1, 0x90, 0xe2, 0xcf, 0x50, 0xc7, 0xf6, 0x39, 0xe4, 0x24, 0x05, 0xaf, 0xb5, 0xdf, 0x3a, 0xd8,
	0x9c, 0x6c, 0x59, 0xe3, 0x0f, 0x24, 0x05, 0x0c, 0xa8, 0xbb, 0xd8, 0x03, 0xaf, 0xb3, 0xdf, 0x3a,
	0xe8, 0x1e, 0x7e, 0xe3, 0x2f, 0xfb, 0x76, 0xa6, 0x1c, 0xbf, 0x96, 0x9b, 0xbd, 0x92, 0xf0, 0x88,
	0xe7, 0xe9, 0xa2, 0x65, 0xd2, 0x01, 0x77, 0x8b, 0x4f, 0x51, 0x4f, 0x69, 0x92, 0xe9, 0x90, 0x12,
	0x0d, 0xa1, 0x66, 0x29, 0x78, 0x1f, 0xec, 0xb7, 0x0e, 0xda, 0x87, 0x03, 0xab, 0x63, 0x6b, 0xf3,
	0xa7, 0x3a, 0x63, 0x3c, 0xbe, 0x20, 0x49, 0x0e, 0x93, 0x8e, 0x01, 0x9d, 0x12, 0x0d, 0x33, 0x96,
	0x02, 0x7e, 0x88, 0x3a, 0xc0, 0xa9, 0xc3, 0xb1, 0xb1, 0x02, 0x47, 0x1b, 0x38, 0xad, 0x19, 0x9e,
	0xa3, 0x2d, 0x42, 0x43, 0x15, 0xfd, 0x0a, 0x34, 0x4f, 0x40, 0x79, 0xdb, 0xfb, 0xeb, 0x07, 0xed,
	0x43, 0x7f, 0x69, 0xb1, 0xe5, 0x58, 0xf9, 0x63, 0x3a, 0xad, 0x20, 0x4f, 0xf8, 0x5c, 0x4c, 0xda,
	0xa4, 0xde, 0x2b, 0x4c, 0xd1, 0x46, 0xf9, 0x29, 0xbd, 0xbe, 0xe9, 0xdc, 0xb3, 0x86, 0xce, 0xd9,
	0xef, 0x33, 0x33, 0x63, 0x70, 0x6a, 0xa0, 0xa6, 0x81, 0xb7, 0x39, 0x26, 0x15, 0x37, 0x7e, 0x81,
	0x06, 0xe5, 0xdc, 0x00, 0x0d, 0x63, 0x10, 0x76, 0x88, 0x22, 0xc1, 0x95, 0x26, 0x5c, 0x7b, 0x1f,
	0xaf, 0xd0, 0x89, 0x5d, 0xcb, 0x70, 0x06, 0xa2, 0x14, 0x39, 0xa9, 0xe0, 0xf8, 0x27, 0xb4, 0x51,
	0x0e, 0xb3, 0xf7, 0xbe, 0x29, 0xe2, 0xe1, 0x8a, 0x45, 0x4c, 0x0d, 0x68, 0x21, 0xfd, 0xd2, 0x34,
	0xa9, 0xf8, 0xf0, 0x2f, 0x08, 0x2b, 0xa6, 0x21, 0x61, 0xfc, 0xf7, 0x37, 0xf7, 0x83, 0xb7, 0x66,
	0xd2, 0xfd, 0xb2, 0xa9, 0xef, 0xd3, 0x0a, 0x69, 0xb9, 0x1f, 0xbf, 0x37, 0xd9, 0x56, 0x37, 0x6c,
	0xf8, 0x4f, 0x34, 0x50, 0x3a, 0xcb, 0x23, 0x9d, 0x67, 0x40, 0x43, 0xc5, 0x99, 0x94, 0xa0, 0x1d,
	0xad, 0x75, 0xa3, 0xf5, 0x75, 0xa3, 0x56, 0xcd, 0x31, 0x2d, 0x29, 0x1c, 0xd1, 0x5d, 0xb5, 0xcc,
	0x89, 0x9f, 0xa3, 0x0e, 0x91, 0xd2, 0x91, 0xfb, 0xd0, 0xc8, 0x7d, 0xde, 0x38, 0x52, 0x52, 0x3a,
	0x02, 0x6d, 0xf2, 0x66, 0x8b, 0x67, 0xa8, 0x1b, 0x91, 0x24, 0x71, 0x38, 0x3f, 0x32, 0x9c, 0x5f,
	0x34, 0x71, 0x9e, 0x90, 0x24, 0x71, 0x48, 0xb7, 0x22, 0x67, 0x8f, 0x5f, 0xa0, 0x7e, 0xb1, 0x17,
	0xb9, 0xdb, 0x9b, 0x4d, 0x43, 0x1c, 0xac, 0x42, 0x2c, 0x72, 0xb7, 0x23, 0xbd, 0x68, 0xd1, 0x84,
	0x7f, 0x43, 0x77, 0x34, 0x5c, 0xeb, 0x30, 0x05, 0xa5, 0x48, 0x0c, 0x8e, 0x06, 0x32, 0x1a, 0x0f,
	0x9a, 0x34, 0x66, 0x70, 0xad, 0xbf, 0x2f, 0xc1, 0x8e, 0xce, 0x8e, 0x7e, 0xdb, 0x8c, 0x7f, 0x44,
	0x3d, 0x99, 0xb1, 0xc8, 0x15, 0x69, 0x1b, 0x91, 0xfb, 0x4d, 0x22, 0xe7, 0x05, 0xcc, 0xa1, 0xef,
	0x48, 0xd7, 0x80, 0x23, 0xb4, 0x23, 0x33, 0x91, 0x0a, 0xbd, 0xf0, 0x9a, 0x79, 0x5b, 0x86, 0x7c,
	0xd8, 0x4c, 0x5e, 0x41, 0x1d, 0x81, 0xbe, 0xbc, 0x69, 0x2c, 0x4e, 0x44, 0x22, 0x22, 0x72, 0x43,
	0xa3, 0xbb, 0xda, 0x89, 0x78, 0x56, 0x21, 0xdd, 0x13, 0x91, 0xdc, 0xb0, 0x15, 0x27, 0x82, 0xcc,
	0xe7, 0x2c, 0x61, 0xc5, 0x45, 0x79, 0x8b, 0x56, 0x6f, 0xb5, 0x13, 0x31, 0xb6, 0x1c, 0xb7, 0x88,
	0xee, 0x92, 0x65, 0x4e, 0xfc, 0x1d, 0xea, 0xd7, 0x57, 0x55, 0x44, 0x52, 0x49, 0x58, 0xcc, 0x3d,
	0xdc, 0x7c, 0x3f, 0x3d, 0x6e, 0x4d, 0xb6, 0x2d, 0xf0, 0xa4, 0xc2, 0xe1, 0xa7, 0x0e, 0x19, 0xa1,
	0x61, 0x9c, 0x89, 0x5c, 0x7a, 0x3b, 0x2b, 0x91, 0xf5, 0x2c, 0x70, 0x4c, 0xcf, 0x0a, 0xd8, 0x71,
	0x1b, 0x6d, 0xd6, 0xaf, 0xd2, 0xf1, 0x00, 0xed, 0x29, 0xc8, 0xae, 0x18, 0x8f, 0xc3, 0xfa, 0x95,
	0x2c, 0x01, 0x8c, 0xc7, 0xc7, 0x7f, 0xad, 0xa1, 0x7b, 0x91, 0x48, 0xfd, 0xc6, 0x1f, 0x98, 0xe3,
	0x3b, 0x6f, 0x3d, 0xbc, 0xe7, 0x45, 0x3a, 0xe7, 0xad, 0x9f, 0x9f, 0x56, 0xe0, 0x58, 0x24, 0x84,
	0xc7, 0xbe, 0xc8, 0xe2, 0x20, 0x06, 0x6e, 0x92, 0xb5, 0x3f, 0x08, 0x92, 0xa9, 0x77, 0xfc, 0x59,
	0x1d, 0xd5, 0xab, 0xbf, 0xd7, 0xd6, 0xcf, 0xc6, 0xe3, 0x7f, 0xd6, 0xee, 0x9e, 0x95, 0x94, 0x63,
	0xaa, 0xfc, 0x72, 0x59, 0xac, 0x2e, 0x86, 0xfe, 0xc4, 0x46, 0xfe, 0x6b, 0x63, 0x2e, 0xc7, 0x54,
	0x5d, 0xd6, 0x31, 0x97, 0x17, 0xc3, 0xcb, 0x3a, 0xe6, 0xff, 0xb5, 0x7b, 0xa5, 0x63, 0x34, 0x1a,
	0x53, 0x35, 0x1a, 0xd5, 0x51, 0xa3, 0xd1, 0xc5, 0x70, 0x34, 0xaa, 0xe3, 0x5e, 0x6e, 0x98, 0x64,
	0x1f, 0xbc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x39, 0xc2, 0x43, 0xc5, 0x05, 0x0a, 0x00, 0x00,
}

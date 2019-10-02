// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/feed_item_target.proto

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

// A feed item target.
type FeedItemTarget struct {
	// The resource name of the feed item target.
	// Feed item target resource names have the form:
	// `customers/{customer_id}/feedItemTargets/{feed_id}~{feed_item_id}~{feed_item_target_type}~{feed_item_target_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The feed item to which this feed item target belongs.
	FeedItem *wrappers.StringValue `protobuf:"bytes,2,opt,name=feed_item,json=feedItem,proto3" json:"feed_item,omitempty"`
	// The target type of this feed item target. This field is read-only.
	FeedItemTargetType enums.FeedItemTargetTypeEnum_FeedItemTargetType `protobuf:"varint,3,opt,name=feed_item_target_type,json=feedItemTargetType,proto3,enum=google.ads.googleads.v2.enums.FeedItemTargetTypeEnum_FeedItemTargetType" json:"feed_item_target_type,omitempty"`
	// The ID of the targeted resource. This field is read-only.
	FeedItemTargetId *wrappers.Int64Value `protobuf:"bytes,6,opt,name=feed_item_target_id,json=feedItemTargetId,proto3" json:"feed_item_target_id,omitempty"`
	// The targeted resource.
	//
	// Types that are valid to be assigned to Target:
	//	*FeedItemTarget_Campaign
	//	*FeedItemTarget_AdGroup
	//	*FeedItemTarget_Keyword
	//	*FeedItemTarget_GeoTargetConstant
	//	*FeedItemTarget_Device
	//	*FeedItemTarget_AdSchedule
	Target               isFeedItemTarget_Target `protobuf_oneof:"target"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FeedItemTarget) Reset()         { *m = FeedItemTarget{} }
func (m *FeedItemTarget) String() string { return proto.CompactTextString(m) }
func (*FeedItemTarget) ProtoMessage()    {}
func (*FeedItemTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_08619793bc0ce69c, []int{0}
}

func (m *FeedItemTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedItemTarget.Unmarshal(m, b)
}
func (m *FeedItemTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedItemTarget.Marshal(b, m, deterministic)
}
func (m *FeedItemTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedItemTarget.Merge(m, src)
}
func (m *FeedItemTarget) XXX_Size() int {
	return xxx_messageInfo_FeedItemTarget.Size(m)
}
func (m *FeedItemTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedItemTarget.DiscardUnknown(m)
}

var xxx_messageInfo_FeedItemTarget proto.InternalMessageInfo

func (m *FeedItemTarget) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *FeedItemTarget) GetFeedItem() *wrappers.StringValue {
	if m != nil {
		return m.FeedItem
	}
	return nil
}

func (m *FeedItemTarget) GetFeedItemTargetType() enums.FeedItemTargetTypeEnum_FeedItemTargetType {
	if m != nil {
		return m.FeedItemTargetType
	}
	return enums.FeedItemTargetTypeEnum_UNSPECIFIED
}

func (m *FeedItemTarget) GetFeedItemTargetId() *wrappers.Int64Value {
	if m != nil {
		return m.FeedItemTargetId
	}
	return nil
}

type isFeedItemTarget_Target interface {
	isFeedItemTarget_Target()
}

type FeedItemTarget_Campaign struct {
	Campaign *wrappers.StringValue `protobuf:"bytes,4,opt,name=campaign,proto3,oneof"`
}

type FeedItemTarget_AdGroup struct {
	AdGroup *wrappers.StringValue `protobuf:"bytes,5,opt,name=ad_group,json=adGroup,proto3,oneof"`
}

type FeedItemTarget_Keyword struct {
	Keyword *common.KeywordInfo `protobuf:"bytes,7,opt,name=keyword,proto3,oneof"`
}

type FeedItemTarget_GeoTargetConstant struct {
	GeoTargetConstant *wrappers.StringValue `protobuf:"bytes,8,opt,name=geo_target_constant,json=geoTargetConstant,proto3,oneof"`
}

type FeedItemTarget_Device struct {
	Device enums.FeedItemTargetDeviceEnum_FeedItemTargetDevice `protobuf:"varint,9,opt,name=device,proto3,enum=google.ads.googleads.v2.enums.FeedItemTargetDeviceEnum_FeedItemTargetDevice,oneof"`
}

type FeedItemTarget_AdSchedule struct {
	AdSchedule *common.AdScheduleInfo `protobuf:"bytes,10,opt,name=ad_schedule,json=adSchedule,proto3,oneof"`
}

func (*FeedItemTarget_Campaign) isFeedItemTarget_Target() {}

func (*FeedItemTarget_AdGroup) isFeedItemTarget_Target() {}

func (*FeedItemTarget_Keyword) isFeedItemTarget_Target() {}

func (*FeedItemTarget_GeoTargetConstant) isFeedItemTarget_Target() {}

func (*FeedItemTarget_Device) isFeedItemTarget_Target() {}

func (*FeedItemTarget_AdSchedule) isFeedItemTarget_Target() {}

func (m *FeedItemTarget) GetTarget() isFeedItemTarget_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *FeedItemTarget) GetCampaign() *wrappers.StringValue {
	if x, ok := m.GetTarget().(*FeedItemTarget_Campaign); ok {
		return x.Campaign
	}
	return nil
}

func (m *FeedItemTarget) GetAdGroup() *wrappers.StringValue {
	if x, ok := m.GetTarget().(*FeedItemTarget_AdGroup); ok {
		return x.AdGroup
	}
	return nil
}

func (m *FeedItemTarget) GetKeyword() *common.KeywordInfo {
	if x, ok := m.GetTarget().(*FeedItemTarget_Keyword); ok {
		return x.Keyword
	}
	return nil
}

func (m *FeedItemTarget) GetGeoTargetConstant() *wrappers.StringValue {
	if x, ok := m.GetTarget().(*FeedItemTarget_GeoTargetConstant); ok {
		return x.GeoTargetConstant
	}
	return nil
}

func (m *FeedItemTarget) GetDevice() enums.FeedItemTargetDeviceEnum_FeedItemTargetDevice {
	if x, ok := m.GetTarget().(*FeedItemTarget_Device); ok {
		return x.Device
	}
	return enums.FeedItemTargetDeviceEnum_UNSPECIFIED
}

func (m *FeedItemTarget) GetAdSchedule() *common.AdScheduleInfo {
	if x, ok := m.GetTarget().(*FeedItemTarget_AdSchedule); ok {
		return x.AdSchedule
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FeedItemTarget) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FeedItemTarget_Campaign)(nil),
		(*FeedItemTarget_AdGroup)(nil),
		(*FeedItemTarget_Keyword)(nil),
		(*FeedItemTarget_GeoTargetConstant)(nil),
		(*FeedItemTarget_Device)(nil),
		(*FeedItemTarget_AdSchedule)(nil),
	}
}

func init() {
	proto.RegisterType((*FeedItemTarget)(nil), "google.ads.googleads.v2.resources.FeedItemTarget")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/feed_item_target.proto", fileDescriptor_08619793bc0ce69c)
}

var fileDescriptor_08619793bc0ce69c = []byte{
	// 580 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0xee, 0x6e, 0x75, 0xbb, 0x9d, 0x6a, 0xd1, 0x29, 0x42, 0xa8, 0x45, 0x5a, 0xa5, 0x50, 0x10,
	0x27, 0x10, 0x45, 0x6c, 0x7a, 0x95, 0xfa, 0xb3, 0xdd, 0x2a, 0xa5, 0xa6, 0x65, 0x2f, 0x64, 0x21,
	0x4c, 0x33, 0x27, 0x63, 0x70, 0x33, 0x13, 0x26, 0x93, 0x96, 0xe2, 0x8d, 0x4f, 0xe0, 0x43, 0x78,
	0xe9, 0xa3, 0xf8, 0x28, 0x3e, 0x85, 0x6c, 0x66, 0x12, 0x58, 0x76, 0xd7, 0x6d, 0xef, 0xce, 0x9e,
	0xf3, 0x7d, 0x5f, 0xbe, 0xf3, 0x9d, 0x6c, 0xd0, 0x1b, 0x2e, 0x25, 0x1f, 0x81, 0x4b, 0x59, 0xe1,
	0x9a, 0x72, 0x5c, 0x5d, 0x7a, 0xae, 0x82, 0x42, 0x96, 0x2a, 0x86, 0xc2, 0x4d, 0x00, 0x58, 0x94,
	0x6a, 0xc8, 0x22, 0x4d, 0x15, 0x07, 0x4d, 0x72, 0x25, 0xb5, 0xc4, 0x3b, 0x06, 0x4e, 0x28, 0x2b,
	0x48, 0xc3, 0x24, 0x97, 0x1e, 0x69, 0x98, 0x9b, 0x2f, 0xe6, 0x89, 0xc7, 0x32, 0xcb, 0xa4, 0x70,
	0x63, 0x95, 0x6a, 0x50, 0x29, 0x35, 0x8a, 0x9b, 0x07, 0xf3, 0xe0, 0x20, 0xca, 0x6c, 0xda, 0x47,
	0xc4, 0xe0, 0x32, 0x8d, 0xc1, 0x92, 0xf7, 0x6f, 0x49, 0xd6, 0xd7, 0x79, 0x4d, 0x7d, 0x62, 0xa9,
	0xd5, 0xaf, 0x8b, 0x32, 0x71, 0xaf, 0x14, 0xcd, 0x73, 0x50, 0x85, 0x9d, 0x6f, 0xd5, 0xd2, 0x79,
	0xea, 0x52, 0x21, 0xa4, 0xa6, 0x3a, 0x95, 0xc2, 0x4e, 0x9f, 0xfe, 0xec, 0xa0, 0xf5, 0x0f, 0x00,
	0xac, 0xaf, 0x21, 0x3b, 0xaf, 0xb4, 0xf1, 0x33, 0x74, 0xbf, 0x0e, 0x21, 0x12, 0x34, 0x03, 0xa7,
	0xb5, 0xdd, 0xda, 0x5b, 0x0d, 0xef, 0xd5, 0xcd, 0x13, 0x9a, 0x01, 0xde, 0x47, 0xab, 0x8d, 0x29,
	0xa7, 0xbd, 0xdd, 0xda, 0x5b, 0xf3, 0xb6, 0x6c, 0x90, 0xa4, 0x76, 0x42, 0xce, 0xb4, 0x4a, 0x05,
	0x1f, 0xd0, 0x51, 0x09, 0x61, 0x37, 0xb1, 0x4f, 0xc1, 0xdf, 0xd1, 0xa3, 0x99, 0xfb, 0x38, 0xcb,
	0xdb, 0xad, 0xbd, 0x75, 0xef, 0x88, 0xcc, 0x3b, 0x4d, 0x95, 0x05, 0x99, 0x74, 0x7b, 0x7e, 0x9d,
	0xc3, 0x7b, 0x51, 0x66, 0x33, 0xda, 0x21, 0x4e, 0xa6, 0x7a, 0xf8, 0x18, 0x6d, 0x4c, 0x3d, 0x3c,
	0x65, 0x4e, 0xa7, 0xda, 0xe0, 0xf1, 0xd4, 0x06, 0x7d, 0xa1, 0x5f, 0xbf, 0x32, 0x0b, 0x3c, 0x98,
	0x54, 0xeb, 0x33, 0xec, 0xa3, 0x6e, 0x4c, 0xb3, 0x9c, 0xa6, 0x5c, 0x38, 0x77, 0x16, 0x47, 0x70,
	0xb4, 0x14, 0x36, 0x78, 0xbc, 0x8f, 0xba, 0x94, 0x45, 0x5c, 0xc9, 0x32, 0x77, 0xee, 0xde, 0x88,
	0xbb, 0x42, 0x59, 0x6f, 0x0c, 0xc7, 0x3d, 0xb4, 0xf2, 0x0d, 0xae, 0xaf, 0xa4, 0x62, 0xce, 0x4a,
	0xc5, 0x7c, 0x3e, 0x37, 0x31, 0xf3, 0xa6, 0x92, 0x8f, 0x06, 0xde, 0x17, 0x89, 0x1c, 0x0b, 0x59,
	0x36, 0x3e, 0x41, 0x1b, 0x1c, 0x64, 0x9d, 0x42, 0x2c, 0x45, 0xa1, 0xa9, 0xd0, 0x4e, 0xf7, 0x46,
	0x76, 0x1e, 0x72, 0x90, 0x26, 0x89, 0xb7, 0x96, 0x88, 0x13, 0xd4, 0x31, 0x2f, 0xb5, 0xb3, 0x5a,
	0x5d, 0xf2, 0xd3, 0xad, 0x2e, 0xf9, 0xae, 0xa2, 0xce, 0xb8, 0xa5, 0x19, 0x1c, 0x2d, 0x85, 0x56,
	0x1d, 0x7f, 0x46, 0x6b, 0x94, 0x45, 0x45, 0xfc, 0x15, 0x58, 0x39, 0x02, 0x07, 0x55, 0x7e, 0xc9,
	0xa2, 0x10, 0x02, 0x76, 0x66, 0x19, 0x36, 0x07, 0x44, 0x9b, 0xce, 0x61, 0x17, 0x75, 0x4c, 0x0c,
	0x87, 0x3f, 0xda, 0x68, 0x37, 0x96, 0x19, 0x59, 0xf8, 0x7d, 0x38, 0xdc, 0x98, 0xb4, 0x79, 0x3a,
	0xce, 0xe9, 0xb4, 0xf5, 0xe5, 0xd8, 0x32, 0xb9, 0x1c, 0x51, 0xc1, 0x89, 0x54, 0xdc, 0xe5, 0x20,
	0xaa, 0x14, 0xeb, 0xbf, 0x76, 0x9e, 0x16, 0xff, 0xf9, 0x64, 0x1d, 0x34, 0xd5, 0xaf, 0xf6, 0x72,
	0x2f, 0x08, 0x7e, 0xb7, 0x77, 0x7a, 0x46, 0x32, 0x60, 0x05, 0x31, 0xe5, 0xb8, 0x1a, 0x78, 0x24,
	0xac, 0x91, 0x7f, 0x6a, 0xcc, 0x30, 0x60, 0xc5, 0xb0, 0xc1, 0x0c, 0x07, 0xde, 0xb0, 0xc1, 0xfc,
	0x6d, 0xef, 0x9a, 0x81, 0xef, 0x07, 0xac, 0xf0, 0xfd, 0x06, 0xe5, 0xfb, 0x03, 0xcf, 0xf7, 0x1b,
	0xdc, 0x45, 0xa7, 0x32, 0xfb, 0xf2, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0x2e, 0x07, 0xc0,
	0x5e, 0x05, 0x00, 0x00,
}

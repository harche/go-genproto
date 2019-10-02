// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/feed_item_target.proto

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

// A feed item target.
type FeedItemTarget struct {
	// The resource name of the feed item target.
	// Feed item target resource names have the form:
	//
	// `customers/{customer_id}/feedItemTargets/{feed_id}~{feed_item_id}~{feed_item_target_type}~{feed_item_target_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The feed item to which this feed item target belongs.
	FeedItem *wrappers.StringValue `protobuf:"bytes,2,opt,name=feed_item,json=feedItem,proto3" json:"feed_item,omitempty"`
	// The target type of this feed item target. This field is read-only.
	FeedItemTargetType enums.FeedItemTargetTypeEnum_FeedItemTargetType `protobuf:"varint,3,opt,name=feed_item_target_type,json=feedItemTargetType,proto3,enum=google.ads.googleads.v1.enums.FeedItemTargetTypeEnum_FeedItemTargetType" json:"feed_item_target_type,omitempty"`
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
	return fileDescriptor_82fa673cb0c842f7, []int{0}
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
	Device enums.FeedItemTargetDeviceEnum_FeedItemTargetDevice `protobuf:"varint,9,opt,name=device,proto3,enum=google.ads.googleads.v1.enums.FeedItemTargetDeviceEnum_FeedItemTargetDevice,oneof"`
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
	proto.RegisterType((*FeedItemTarget)(nil), "google.ads.googleads.v1.resources.FeedItemTarget")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/feed_item_target.proto", fileDescriptor_82fa673cb0c842f7)
}

var fileDescriptor_82fa673cb0c842f7 = []byte{
	// 580 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x5e, 0x3b, 0xe8, 0x3a, 0x0f, 0x26, 0xf0, 0x84, 0x14, 0x8d, 0x09, 0x6d, 0xa0, 0x49, 0x93,
	0x10, 0x8e, 0x3a, 0x10, 0x62, 0xd9, 0x55, 0xc6, 0x4f, 0xd7, 0x81, 0xa6, 0x91, 0x4d, 0xbd, 0x40,
	0x95, 0x22, 0x2f, 0x3e, 0x31, 0x11, 0x8d, 0x1d, 0x39, 0xce, 0xa6, 0x8a, 0x1b, 0x9e, 0x80, 0x87,
	0xe0, 0x92, 0x47, 0xe1, 0x51, 0x78, 0x0a, 0xd4, 0xd8, 0x89, 0x54, 0xb5, 0xa5, 0xdb, 0xdd, 0xc9,
	0x39, 0xdf, 0x77, 0xfc, 0x9d, 0xef, 0x38, 0x46, 0x6f, 0xb8, 0x94, 0x7c, 0x08, 0x2e, 0x65, 0xb9,
	0x6b, 0xc2, 0x71, 0x74, 0xd5, 0x71, 0x15, 0xe4, 0xb2, 0x50, 0x11, 0xe4, 0x6e, 0x0c, 0xc0, 0xc2,
	0x44, 0x43, 0x1a, 0x6a, 0xaa, 0x38, 0x68, 0x92, 0x29, 0xa9, 0x25, 0xde, 0x31, 0x70, 0x42, 0x59,
	0x4e, 0x6a, 0x26, 0xb9, 0xea, 0x90, 0x9a, 0xb9, 0xf9, 0x62, 0x5e, 0xf3, 0x48, 0xa6, 0xa9, 0x14,
	0x6e, 0xa4, 0x12, 0x0d, 0x2a, 0xa1, 0xa6, 0xe3, 0xe6, 0xe1, 0x3c, 0x38, 0x88, 0x22, 0x9d, 0xd6,
	0x11, 0x32, 0xb8, 0x4a, 0x22, 0xb0, 0xe4, 0x83, 0x5b, 0x92, 0xf5, 0x28, 0xab, 0xa8, 0x5b, 0x15,
	0x35, 0x4b, 0x5c, 0x2a, 0x84, 0xd4, 0x54, 0x27, 0x52, 0xe4, 0xb6, 0xfa, 0xc4, 0x56, 0xcb, 0xaf,
	0xcb, 0x22, 0x76, 0xaf, 0x15, 0xcd, 0x32, 0x50, 0xb6, 0xfe, 0xf4, 0x67, 0x0b, 0xad, 0x7f, 0x00,
	0x60, 0x3d, 0x0d, 0xe9, 0x45, 0xd9, 0x1b, 0x3f, 0x43, 0xf7, 0x2b, 0x13, 0x42, 0x41, 0x53, 0x70,
	0x1a, 0xdb, 0x8d, 0xbd, 0xd5, 0xe0, 0x5e, 0x95, 0x3c, 0xa5, 0x29, 0xe0, 0x03, 0xb4, 0x5a, 0x8b,
	0x72, 0x9a, 0xdb, 0x8d, 0xbd, 0xb5, 0xfd, 0x2d, 0x6b, 0x24, 0xa9, 0xce, 0x22, 0xe7, 0x5a, 0x25,
	0x82, 0xf7, 0xe9, 0xb0, 0x80, 0xa0, 0x1d, 0xdb, 0x53, 0xf0, 0x77, 0xf4, 0x68, 0xe6, 0x3c, 0xce,
	0xf2, 0x76, 0x63, 0x6f, 0x7d, 0xff, 0x98, 0xcc, 0x5b, 0x4d, 0xe9, 0x05, 0x99, 0x54, 0x7b, 0x31,
	0xca, 0xe0, 0xbd, 0x28, 0xd2, 0x19, 0xe9, 0x00, 0xc7, 0x53, 0x39, 0x7c, 0x82, 0x36, 0xa6, 0x0e,
	0x4f, 0x98, 0xd3, 0x2a, 0x27, 0x78, 0x3c, 0x35, 0x41, 0x4f, 0xe8, 0xd7, 0xaf, 0xcc, 0x00, 0x0f,
	0x26, 0xbb, 0xf5, 0x18, 0xf6, 0x50, 0x3b, 0xa2, 0x69, 0x46, 0x13, 0x2e, 0x9c, 0x3b, 0x8b, 0x2d,
	0x38, 0x5e, 0x0a, 0x6a, 0x3c, 0x3e, 0x40, 0x6d, 0xca, 0x42, 0xae, 0x64, 0x91, 0x39, 0x77, 0x6f,
	0xc4, 0x5d, 0xa1, 0xac, 0x3b, 0x86, 0xe3, 0x2e, 0x5a, 0xf9, 0x06, 0xa3, 0x6b, 0xa9, 0x98, 0xb3,
	0x52, 0x32, 0x9f, 0xcf, 0x75, 0xcc, 0xdc, 0x54, 0xf2, 0xd1, 0xc0, 0x7b, 0x22, 0x96, 0xe3, 0x46,
	0x96, 0x8d, 0x4f, 0xd1, 0x06, 0x07, 0x59, 0xb9, 0x10, 0x49, 0x91, 0x6b, 0x2a, 0xb4, 0xd3, 0xbe,
	0x91, 0x9c, 0x87, 0x1c, 0xa4, 0x71, 0xe2, 0xad, 0x25, 0xe2, 0x18, 0xb5, 0xcc, 0xa5, 0x76, 0x56,
	0xcb, 0x4d, 0x7e, 0xba, 0xd5, 0x26, 0xdf, 0x95, 0xd4, 0x19, 0xbb, 0x34, 0x85, 0xe3, 0xa5, 0xc0,
	0x76, 0xc7, 0x9f, 0xd1, 0x1a, 0x65, 0x61, 0x1e, 0x7d, 0x05, 0x56, 0x0c, 0xc1, 0x41, 0xa5, 0x5e,
	0xb2, 0xc8, 0x04, 0x9f, 0x9d, 0x5b, 0x86, 0xf5, 0x01, 0xd1, 0x3a, 0x73, 0xd4, 0x46, 0x2d, 0x63,
	0xc3, 0xd1, 0x8f, 0x26, 0xda, 0x8d, 0x64, 0x4a, 0x16, 0xbe, 0x0f, 0x47, 0x1b, 0x93, 0x32, 0xcf,
	0xc6, 0x3e, 0x9d, 0x35, 0xbe, 0x9c, 0x58, 0x26, 0x97, 0x43, 0x2a, 0x38, 0x91, 0x8a, 0xbb, 0x1c,
	0x44, 0xe9, 0x62, 0xf5, 0x6b, 0x67, 0x49, 0xfe, 0x9f, 0x27, 0xeb, 0xb0, 0x8e, 0x7e, 0x35, 0x97,
	0xbb, 0xbe, 0xff, 0xbb, 0xb9, 0xd3, 0x35, 0x2d, 0x7d, 0x96, 0x13, 0x13, 0x8e, 0xa3, 0x7e, 0x87,
	0x04, 0x15, 0xf2, 0x4f, 0x85, 0x19, 0xf8, 0x2c, 0x1f, 0xd4, 0x98, 0x41, 0xbf, 0x33, 0xa8, 0x31,
	0x7f, 0x9b, 0xbb, 0xa6, 0xe0, 0x79, 0x3e, 0xcb, 0x3d, 0xaf, 0x46, 0x79, 0x5e, 0xbf, 0xe3, 0x79,
	0x35, 0xee, 0xb2, 0x55, 0x8a, 0x7d, 0xf9, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x6f, 0xf4, 0x7c,
	0x5e, 0x05, 0x00, 0x00,
}

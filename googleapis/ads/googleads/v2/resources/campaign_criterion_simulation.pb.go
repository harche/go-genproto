// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/campaign_criterion_simulation.proto

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

// A campaign criterion simulation. Supported combinations of advertising
// channel type, criterion ids, simulation type and simulation modification
// method is detailed below respectively.
//
// SEARCH     30000,30001,30002  BID_MODIFIER  UNIFORM
// SHOPPING   30000,30001,30002  BID_MODIFIER  UNIFORM
// DISPLAY    30001              BID_MODIFIER  UNIFORM
type CampaignCriterionSimulation struct {
	// The resource name of the campaign criterion simulation.
	// Campaign criterion simulation resource names have the form:
	//
	// `customers/{customer_id}/campaignCriterionSimulations/{campaign_id}~{criterion_id}~{type}~{modification_method}~{start_date}~{end_date}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Campaign ID of the simulation.
	CampaignId *wrappers.Int64Value `protobuf:"bytes,2,opt,name=campaign_id,json=campaignId,proto3" json:"campaign_id,omitempty"`
	// Criterion ID of the simulation.
	CriterionId *wrappers.Int64Value `protobuf:"bytes,3,opt,name=criterion_id,json=criterionId,proto3" json:"criterion_id,omitempty"`
	// The field that the simulation modifies.
	Type enums.SimulationTypeEnum_SimulationType `protobuf:"varint,4,opt,name=type,proto3,enum=google.ads.googleads.v2.enums.SimulationTypeEnum_SimulationType" json:"type,omitempty"`
	// How the simulation modifies the field.
	ModificationMethod enums.SimulationModificationMethodEnum_SimulationModificationMethod `protobuf:"varint,5,opt,name=modification_method,json=modificationMethod,proto3,enum=google.ads.googleads.v2.enums.SimulationModificationMethodEnum_SimulationModificationMethod" json:"modification_method,omitempty"`
	// First day on which the simulation is based, in YYYY-MM-DD format.
	StartDate *wrappers.StringValue `protobuf:"bytes,6,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	// Last day on which the simulation is based, in YYYY-MM-DD format.
	EndDate *wrappers.StringValue `protobuf:"bytes,7,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	// List of simulation points.
	//
	// Types that are valid to be assigned to PointList:
	//	*CampaignCriterionSimulation_BidModifierPointList
	PointList            isCampaignCriterionSimulation_PointList `protobuf_oneof:"point_list"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *CampaignCriterionSimulation) Reset()         { *m = CampaignCriterionSimulation{} }
func (m *CampaignCriterionSimulation) String() string { return proto.CompactTextString(m) }
func (*CampaignCriterionSimulation) ProtoMessage()    {}
func (*CampaignCriterionSimulation) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5832e91cef223b5, []int{0}
}

func (m *CampaignCriterionSimulation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignCriterionSimulation.Unmarshal(m, b)
}
func (m *CampaignCriterionSimulation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignCriterionSimulation.Marshal(b, m, deterministic)
}
func (m *CampaignCriterionSimulation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignCriterionSimulation.Merge(m, src)
}
func (m *CampaignCriterionSimulation) XXX_Size() int {
	return xxx_messageInfo_CampaignCriterionSimulation.Size(m)
}
func (m *CampaignCriterionSimulation) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignCriterionSimulation.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignCriterionSimulation proto.InternalMessageInfo

func (m *CampaignCriterionSimulation) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignCriterionSimulation) GetCampaignId() *wrappers.Int64Value {
	if m != nil {
		return m.CampaignId
	}
	return nil
}

func (m *CampaignCriterionSimulation) GetCriterionId() *wrappers.Int64Value {
	if m != nil {
		return m.CriterionId
	}
	return nil
}

func (m *CampaignCriterionSimulation) GetType() enums.SimulationTypeEnum_SimulationType {
	if m != nil {
		return m.Type
	}
	return enums.SimulationTypeEnum_UNSPECIFIED
}

func (m *CampaignCriterionSimulation) GetModificationMethod() enums.SimulationModificationMethodEnum_SimulationModificationMethod {
	if m != nil {
		return m.ModificationMethod
	}
	return enums.SimulationModificationMethodEnum_UNSPECIFIED
}

func (m *CampaignCriterionSimulation) GetStartDate() *wrappers.StringValue {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *CampaignCriterionSimulation) GetEndDate() *wrappers.StringValue {
	if m != nil {
		return m.EndDate
	}
	return nil
}

type isCampaignCriterionSimulation_PointList interface {
	isCampaignCriterionSimulation_PointList()
}

type CampaignCriterionSimulation_BidModifierPointList struct {
	BidModifierPointList *common.BidModifierSimulationPointList `protobuf:"bytes,8,opt,name=bid_modifier_point_list,json=bidModifierPointList,proto3,oneof"`
}

func (*CampaignCriterionSimulation_BidModifierPointList) isCampaignCriterionSimulation_PointList() {}

func (m *CampaignCriterionSimulation) GetPointList() isCampaignCriterionSimulation_PointList {
	if m != nil {
		return m.PointList
	}
	return nil
}

func (m *CampaignCriterionSimulation) GetBidModifierPointList() *common.BidModifierSimulationPointList {
	if x, ok := m.GetPointList().(*CampaignCriterionSimulation_BidModifierPointList); ok {
		return x.BidModifierPointList
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignCriterionSimulation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignCriterionSimulation_BidModifierPointList)(nil),
	}
}

func init() {
	proto.RegisterType((*CampaignCriterionSimulation)(nil), "google.ads.googleads.v2.resources.CampaignCriterionSimulation")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/campaign_criterion_simulation.proto", fileDescriptor_a5832e91cef223b5)
}

var fileDescriptor_a5832e91cef223b5 = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdf, 0x6a, 0xdb, 0x30,
	0x14, 0xc6, 0xe7, 0xf4, 0xbf, 0xda, 0xed, 0xc2, 0x1b, 0xcc, 0xb4, 0x65, 0xa4, 0x1b, 0x85, 0x5c,
	0xc9, 0xe0, 0x8e, 0x0d, 0xdc, 0x51, 0x96, 0x74, 0xa5, 0xcb, 0x58, 0x46, 0x70, 0x4b, 0x2e, 0x46,
	0xc0, 0x28, 0x96, 0xea, 0x09, 0x22, 0xc9, 0x48, 0x72, 0x4b, 0x1f, 0xa2, 0x2f, 0xb1, 0xdd, 0xed,
	0x51, 0xf6, 0x28, 0x7b, 0x8a, 0x11, 0x59, 0x56, 0x4a, 0xb2, 0xb4, 0xb9, 0x3b, 0x3e, 0xfa, 0xbe,
	0xdf, 0xb1, 0x3e, 0x59, 0x06, 0x67, 0xb9, 0x10, 0xf9, 0x98, 0x84, 0x08, 0xab, 0xb0, 0x2a, 0x27,
	0xd5, 0x75, 0x14, 0x4a, 0xa2, 0x44, 0x29, 0x33, 0xa2, 0xc2, 0x0c, 0xb1, 0x02, 0xd1, 0x9c, 0xa7,
	0x99, 0xa4, 0x9a, 0x48, 0x2a, 0x78, 0xaa, 0x28, 0x2b, 0xc7, 0x48, 0x53, 0xc1, 0x61, 0x21, 0x85,
	0x16, 0xfe, 0x41, 0xe5, 0x85, 0x08, 0x2b, 0xe8, 0x30, 0xf0, 0x3a, 0x82, 0x0e, 0xb3, 0x1b, 0x2e,
	0x9a, 0x94, 0x09, 0xc6, 0x04, 0x0f, 0x67, 0x99, 0xbb, 0x9d, 0x45, 0x06, 0xc2, 0x4b, 0xa6, 0xee,
	0xe9, 0x53, 0x26, 0x30, 0xbd, 0xa2, 0x99, 0x7d, 0x20, 0xfa, 0x87, 0xc0, 0x96, 0x71, 0xb4, 0x34,
	0x43, 0xdf, 0x16, 0xc4, 0x9a, 0x5e, 0x59, 0x93, 0x79, 0x1a, 0x95, 0x57, 0xe1, 0x8d, 0x44, 0x45,
	0x41, 0xa4, 0xb2, 0xeb, 0xfb, 0x35, 0xb4, 0xa0, 0x21, 0xe2, 0x5c, 0x68, 0x43, 0xb0, 0xab, 0xaf,
	0x7f, 0xad, 0x81, 0xbd, 0x53, 0x1b, 0xd9, 0x69, 0x9d, 0xd8, 0x85, 0x1b, 0xe4, 0xbf, 0x01, 0x4f,
	0xeb, 0x50, 0x52, 0x8e, 0x18, 0x09, 0xbc, 0xa6, 0xd7, 0xda, 0x4a, 0x76, 0xea, 0xe6, 0x37, 0xc4,
	0x88, 0xff, 0x01, 0x6c, 0xbb, 0xd8, 0x29, 0x0e, 0x1a, 0x4d, 0xaf, 0xb5, 0x1d, 0xed, 0xd9, 0x68,
	0x61, 0xfd, 0x62, 0xb0, 0xcb, 0xf5, 0xbb, 0xb7, 0x03, 0x34, 0x2e, 0x49, 0x02, 0x6a, 0x7d, 0x17,
	0xfb, 0x27, 0x60, 0x67, 0x7a, 0x56, 0x14, 0x07, 0x2b, 0x8f, 0xdb, 0xb7, 0x9d, 0xa1, 0x8b, 0xfd,
	0x4b, 0xb0, 0x3a, 0x89, 0x23, 0x58, 0x6d, 0x7a, 0xad, 0x67, 0xd1, 0x47, 0xb8, 0xe8, 0x70, 0x4d,
	0x88, 0x70, 0xba, 0xb7, 0xcb, 0xdb, 0x82, 0x9c, 0xf1, 0x92, 0xcd, 0xb4, 0x12, 0x43, 0xf3, 0xef,
	0x3c, 0xf0, 0xfc, 0x3f, 0x27, 0x15, 0xac, 0x99, 0x29, 0xc3, 0xa5, 0xa7, 0xf4, 0xee, 0x31, 0x7a,
	0x06, 0x31, 0x33, 0x73, 0x5e, 0x90, 0xf8, 0x6c, 0xae, 0xe7, 0x1f, 0x03, 0xa0, 0x34, 0x92, 0x3a,
	0xc5, 0x48, 0x93, 0x60, 0xdd, 0x64, 0xb4, 0x3f, 0x97, 0xd1, 0x85, 0x96, 0x94, 0xe7, 0x55, 0x48,
	0x5b, 0x46, 0xff, 0x09, 0x69, 0xe2, 0xbf, 0x07, 0x9b, 0x84, 0xe3, 0xca, 0xba, 0xb1, 0x84, 0x75,
	0x83, 0x70, 0x6c, 0x8c, 0x37, 0xe0, 0xe5, 0x88, 0x62, 0xfb, 0xc9, 0x12, 0x99, 0x16, 0x82, 0x72,
	0x9d, 0x8e, 0xa9, 0xd2, 0xc1, 0xa6, 0xe1, 0x9c, 0x2c, 0x0c, 0xa2, 0xba, 0x28, 0xb0, 0x43, 0x71,
	0xcf, 0xba, 0xa7, 0x7b, 0xee, 0x4f, 0x30, 0x5f, 0xa9, 0xd2, 0x9f, 0x9f, 0x24, 0x2f, 0x46, 0x53,
	0x85, 0xeb, 0x77, 0x76, 0x00, 0x98, 0xce, 0xea, 0xdc, 0x35, 0xc0, 0x61, 0x26, 0x18, 0x7c, 0xf4,
	0xde, 0x76, 0x9a, 0x0f, 0x7c, 0xcc, 0xfd, 0xc9, 0x56, 0xfb, 0xde, 0xf7, 0x2f, 0x16, 0x93, 0x8b,
	0x31, 0xe2, 0x39, 0x14, 0x32, 0x0f, 0x73, 0xc2, 0x4d, 0x10, 0xf5, 0xb5, 0x2b, 0xa8, 0x7a, 0xe0,
	0x27, 0x73, 0xec, 0xaa, 0x9f, 0x8d, 0x95, 0xf3, 0x76, 0xfb, 0x77, 0xe3, 0xe0, 0xbc, 0x42, 0xb6,
	0xb1, 0x82, 0x55, 0x39, 0xa9, 0x06, 0x11, 0x4c, 0x6a, 0xe5, 0x9f, 0x5a, 0x33, 0x6c, 0x63, 0x35,
	0x74, 0x9a, 0xe1, 0x20, 0x1a, 0x3a, 0xcd, 0xdf, 0xc6, 0x61, 0xb5, 0x10, 0xc7, 0x6d, 0xac, 0xe2,
	0xd8, 0xa9, 0xe2, 0x78, 0x10, 0xc5, 0xb1, 0xd3, 0x8d, 0xd6, 0xcd, 0xcb, 0x1e, 0xfd, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0xfe, 0x81, 0xb6, 0xa9, 0x10, 0x05, 0x00, 0x00,
}

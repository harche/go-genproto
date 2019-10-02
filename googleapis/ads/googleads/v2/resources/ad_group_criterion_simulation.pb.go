// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/ad_group_criterion_simulation.proto

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

// An ad group criterion simulation. Supported combinations of advertising
// channel type, criterion type, simulation type, and simulation modification
// method are detailed below respectively.
//
// DISPLAY   KEYWORD        CPC_BID  UNIFORM
// SEARCH    KEYWORD        CPC_BID  UNIFORM
// SHOPPING  LISTING_GROUP  CPC_BID  UNIFORM
type AdGroupCriterionSimulation struct {
	// The resource name of the ad group criterion simulation.
	// Ad group criterion simulation resource names have the form:
	//
	// `customers/{customer_id}/adGroupCriterionSimulations/{ad_group_id}~{criterion_id}~{type}~{modification_method}~{start_date}~{end_date}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// AdGroup ID of the simulation.
	AdGroupId *wrappers.Int64Value `protobuf:"bytes,2,opt,name=ad_group_id,json=adGroupId,proto3" json:"ad_group_id,omitempty"`
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
	//	*AdGroupCriterionSimulation_CpcBidPointList
	PointList            isAdGroupCriterionSimulation_PointList `protobuf_oneof:"point_list"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *AdGroupCriterionSimulation) Reset()         { *m = AdGroupCriterionSimulation{} }
func (m *AdGroupCriterionSimulation) String() string { return proto.CompactTextString(m) }
func (*AdGroupCriterionSimulation) ProtoMessage()    {}
func (*AdGroupCriterionSimulation) Descriptor() ([]byte, []int) {
	return fileDescriptor_b92456ddc38bf208, []int{0}
}

func (m *AdGroupCriterionSimulation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupCriterionSimulation.Unmarshal(m, b)
}
func (m *AdGroupCriterionSimulation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupCriterionSimulation.Marshal(b, m, deterministic)
}
func (m *AdGroupCriterionSimulation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupCriterionSimulation.Merge(m, src)
}
func (m *AdGroupCriterionSimulation) XXX_Size() int {
	return xxx_messageInfo_AdGroupCriterionSimulation.Size(m)
}
func (m *AdGroupCriterionSimulation) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupCriterionSimulation.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupCriterionSimulation proto.InternalMessageInfo

func (m *AdGroupCriterionSimulation) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroupCriterionSimulation) GetAdGroupId() *wrappers.Int64Value {
	if m != nil {
		return m.AdGroupId
	}
	return nil
}

func (m *AdGroupCriterionSimulation) GetCriterionId() *wrappers.Int64Value {
	if m != nil {
		return m.CriterionId
	}
	return nil
}

func (m *AdGroupCriterionSimulation) GetType() enums.SimulationTypeEnum_SimulationType {
	if m != nil {
		return m.Type
	}
	return enums.SimulationTypeEnum_UNSPECIFIED
}

func (m *AdGroupCriterionSimulation) GetModificationMethod() enums.SimulationModificationMethodEnum_SimulationModificationMethod {
	if m != nil {
		return m.ModificationMethod
	}
	return enums.SimulationModificationMethodEnum_UNSPECIFIED
}

func (m *AdGroupCriterionSimulation) GetStartDate() *wrappers.StringValue {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *AdGroupCriterionSimulation) GetEndDate() *wrappers.StringValue {
	if m != nil {
		return m.EndDate
	}
	return nil
}

type isAdGroupCriterionSimulation_PointList interface {
	isAdGroupCriterionSimulation_PointList()
}

type AdGroupCriterionSimulation_CpcBidPointList struct {
	CpcBidPointList *common.CpcBidSimulationPointList `protobuf:"bytes,8,opt,name=cpc_bid_point_list,json=cpcBidPointList,proto3,oneof"`
}

func (*AdGroupCriterionSimulation_CpcBidPointList) isAdGroupCriterionSimulation_PointList() {}

func (m *AdGroupCriterionSimulation) GetPointList() isAdGroupCriterionSimulation_PointList {
	if m != nil {
		return m.PointList
	}
	return nil
}

func (m *AdGroupCriterionSimulation) GetCpcBidPointList() *common.CpcBidSimulationPointList {
	if x, ok := m.GetPointList().(*AdGroupCriterionSimulation_CpcBidPointList); ok {
		return x.CpcBidPointList
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdGroupCriterionSimulation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdGroupCriterionSimulation_CpcBidPointList)(nil),
	}
}

func init() {
	proto.RegisterType((*AdGroupCriterionSimulation)(nil), "google.ads.googleads.v2.resources.AdGroupCriterionSimulation")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/ad_group_criterion_simulation.proto", fileDescriptor_b92456ddc38bf208)
}

var fileDescriptor_b92456ddc38bf208 = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xd3, 0x3e,
	0x14, 0xff, 0xa7, 0xfb, 0xf6, 0xf6, 0x07, 0x29, 0xdc, 0x44, 0x65, 0x82, 0x0e, 0x34, 0xa9, 0x57,
	0x8e, 0x94, 0x21, 0x10, 0xa9, 0x84, 0x48, 0xc7, 0x54, 0x8a, 0x18, 0xaa, 0xb2, 0xa9, 0x17, 0xa8,
	0x52, 0xe4, 0xc6, 0x5e, 0x66, 0xa9, 0xb1, 0x2d, 0xdb, 0x19, 0xda, 0x3b, 0xc0, 0x0d, 0x8f, 0xc0,
	0x25, 0x8f, 0xc2, 0xa3, 0xf0, 0x14, 0x28, 0x4e, 0xe2, 0x4e, 0x2d, 0x65, 0xbd, 0x3b, 0x39, 0xe7,
	0xf7, 0xe1, 0xf3, 0x73, 0x5d, 0x70, 0x96, 0x71, 0x9e, 0xcd, 0x88, 0x8f, 0xb0, 0xf2, 0xab, 0xb2,
	0xac, 0x6e, 0x02, 0x5f, 0x12, 0xc5, 0x0b, 0x99, 0x12, 0xe5, 0x23, 0x9c, 0x64, 0x92, 0x17, 0x22,
	0x49, 0x25, 0xd5, 0x44, 0x52, 0xce, 0x12, 0x45, 0xf3, 0x62, 0x86, 0x34, 0xe5, 0x0c, 0x0a, 0xc9,
	0x35, 0x77, 0x8f, 0x2a, 0x2e, 0x44, 0x58, 0x41, 0x2b, 0x03, 0x6f, 0x02, 0x68, 0x65, 0xda, 0xfe,
	0x2a, 0xa7, 0x94, 0xe7, 0x39, 0x67, 0xfe, 0xa2, 0x66, 0xbb, 0xbf, 0x8a, 0x40, 0x58, 0x91, 0xab,
	0x3b, 0xf8, 0x24, 0xe7, 0x98, 0x5e, 0xd1, 0xb4, 0xfe, 0x20, 0xfa, 0x9a, 0xe3, 0x5a, 0xe3, 0x64,
	0x6d, 0x0d, 0x7d, 0x2b, 0x48, 0x4d, 0x7a, 0x52, 0x93, 0xcc, 0xd7, 0xb4, 0xb8, 0xf2, 0xbf, 0x48,
	0x24, 0x04, 0x91, 0xaa, 0x9e, 0x1f, 0x36, 0xa2, 0x82, 0xfa, 0x88, 0x31, 0xae, 0x8d, 0x42, 0x3d,
	0x7d, 0xf6, 0x7d, 0x0b, 0xb4, 0x23, 0x3c, 0x28, 0x13, 0x3b, 0x6d, 0x02, 0xbb, 0xb0, 0x3e, 0xee,
	0x73, 0xf0, 0x7f, 0x93, 0x49, 0xc2, 0x50, 0x4e, 0x3c, 0xa7, 0xe3, 0x74, 0xf7, 0xe2, 0x83, 0xa6,
	0xf9, 0x09, 0xe5, 0xc4, 0xed, 0x81, 0x7d, 0x9b, 0x3a, 0xc5, 0x5e, 0xab, 0xe3, 0x74, 0xf7, 0x83,
	0xc7, 0x75, 0xb2, 0xb0, 0x39, 0x17, 0x1c, 0x32, 0xfd, 0xf2, 0xc5, 0x18, 0xcd, 0x0a, 0x12, 0xef,
	0xa1, 0xca, 0x72, 0x88, 0xdd, 0x37, 0xe0, 0x60, 0x7e, 0x53, 0x14, 0x7b, 0x1b, 0xf7, 0xb3, 0xf7,
	0x2d, 0x61, 0x88, 0xdd, 0x4b, 0xb0, 0x59, 0x86, 0xe1, 0x6d, 0x76, 0x9c, 0xee, 0x83, 0xe0, 0x2d,
	0x5c, 0x75, 0xb5, 0x26, 0x42, 0x38, 0x5f, 0xed, 0xf2, 0x56, 0x90, 0x33, 0x56, 0xe4, 0x0b, 0xad,
	0xd8, 0xa8, 0xb9, 0xdf, 0x1c, 0xf0, 0xe8, 0x2f, 0xf7, 0xe4, 0x6d, 0x19, 0x97, 0xc9, 0xda, 0x2e,
	0xe7, 0x77, 0x34, 0xce, 0x8d, 0xc4, 0x82, 0xe7, 0x32, 0x20, 0x76, 0xf3, 0xa5, 0x9e, 0xdb, 0x03,
	0x40, 0x69, 0x24, 0x75, 0x82, 0x91, 0x26, 0xde, 0xb6, 0xc9, 0xe8, 0x70, 0x29, 0xa3, 0x0b, 0x2d,
	0x29, 0xcb, 0xea, 0x88, 0x0d, 0xfe, 0x1d, 0xd2, 0xc4, 0x7d, 0x05, 0x76, 0x09, 0xc3, 0x15, 0x75,
	0x67, 0x0d, 0xea, 0x0e, 0x61, 0xd8, 0x10, 0xaf, 0x81, 0x9b, 0x8a, 0x34, 0x99, 0x52, 0x9c, 0x08,
	0x4e, 0x99, 0x4e, 0x66, 0x54, 0x69, 0x6f, 0xd7, 0x48, 0xbc, 0x5e, 0x99, 0x41, 0xf5, 0x42, 0xe0,
	0xa9, 0x48, 0xfb, 0x14, 0xcf, 0x37, 0x1d, 0x95, 0x0a, 0x1f, 0xa9, 0xd2, 0xef, 0xff, 0x8b, 0x1f,
	0xa6, 0x66, 0x68, 0x5b, 0xfd, 0x03, 0x00, 0xe6, 0x0e, 0xfd, 0xaf, 0x2d, 0x70, 0x9c, 0xf2, 0x1c,
	0xde, 0xfb, 0x4c, 0xfb, 0x4f, 0x57, 0xff, 0x76, 0x47, 0xe5, 0x6a, 0x23, 0xe7, 0xf3, 0x87, 0x5a,
	0x25, 0xe3, 0x33, 0xc4, 0x32, 0xc8, 0x65, 0xe6, 0x67, 0x84, 0x99, 0xc5, 0x9b, 0x47, 0x26, 0xa8,
	0xfa, 0xc7, 0x5f, 0x4a, 0xcf, 0x56, 0x3f, 0x5a, 0x1b, 0x83, 0x28, 0xfa, 0xd9, 0x3a, 0x1a, 0x54,
	0x92, 0x11, 0x56, 0xb0, 0x2a, 0xcb, 0x6a, 0x1c, 0xc0, 0xb8, 0x41, 0xfe, 0x6a, 0x30, 0x93, 0x08,
	0xab, 0x89, 0xc5, 0x4c, 0xc6, 0xc1, 0xc4, 0x62, 0x7e, 0xb7, 0x8e, 0xab, 0x41, 0x18, 0x46, 0x58,
	0x85, 0xa1, 0x45, 0x85, 0xe1, 0x38, 0x08, 0x43, 0x8b, 0x9b, 0x6e, 0x9b, 0xc3, 0x9e, 0xfc, 0x09,
	0x00, 0x00, 0xff, 0xff, 0xfe, 0xd0, 0x0e, 0xc6, 0xfe, 0x04, 0x00, 0x00,
}

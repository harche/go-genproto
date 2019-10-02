// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/common/simulation.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A container for simulation points for simulations of type BID_MODIFIER.
type BidModifierSimulationPointList struct {
	// Projected metrics for a series of bid modifier amounts.
	Points               []*BidModifierSimulationPoint `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *BidModifierSimulationPointList) Reset()         { *m = BidModifierSimulationPointList{} }
func (m *BidModifierSimulationPointList) String() string { return proto.CompactTextString(m) }
func (*BidModifierSimulationPointList) ProtoMessage()    {}
func (*BidModifierSimulationPointList) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4802a06f80f525c, []int{0}
}

func (m *BidModifierSimulationPointList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BidModifierSimulationPointList.Unmarshal(m, b)
}
func (m *BidModifierSimulationPointList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BidModifierSimulationPointList.Marshal(b, m, deterministic)
}
func (m *BidModifierSimulationPointList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidModifierSimulationPointList.Merge(m, src)
}
func (m *BidModifierSimulationPointList) XXX_Size() int {
	return xxx_messageInfo_BidModifierSimulationPointList.Size(m)
}
func (m *BidModifierSimulationPointList) XXX_DiscardUnknown() {
	xxx_messageInfo_BidModifierSimulationPointList.DiscardUnknown(m)
}

var xxx_messageInfo_BidModifierSimulationPointList proto.InternalMessageInfo

func (m *BidModifierSimulationPointList) GetPoints() []*BidModifierSimulationPoint {
	if m != nil {
		return m.Points
	}
	return nil
}

// A container for simulation points for simulations of type CPC_BID.
type CpcBidSimulationPointList struct {
	// Projected metrics for a series of CPC bid amounts.
	Points               []*CpcBidSimulationPoint `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CpcBidSimulationPointList) Reset()         { *m = CpcBidSimulationPointList{} }
func (m *CpcBidSimulationPointList) String() string { return proto.CompactTextString(m) }
func (*CpcBidSimulationPointList) ProtoMessage()    {}
func (*CpcBidSimulationPointList) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4802a06f80f525c, []int{1}
}

func (m *CpcBidSimulationPointList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CpcBidSimulationPointList.Unmarshal(m, b)
}
func (m *CpcBidSimulationPointList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CpcBidSimulationPointList.Marshal(b, m, deterministic)
}
func (m *CpcBidSimulationPointList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CpcBidSimulationPointList.Merge(m, src)
}
func (m *CpcBidSimulationPointList) XXX_Size() int {
	return xxx_messageInfo_CpcBidSimulationPointList.Size(m)
}
func (m *CpcBidSimulationPointList) XXX_DiscardUnknown() {
	xxx_messageInfo_CpcBidSimulationPointList.DiscardUnknown(m)
}

var xxx_messageInfo_CpcBidSimulationPointList proto.InternalMessageInfo

func (m *CpcBidSimulationPointList) GetPoints() []*CpcBidSimulationPoint {
	if m != nil {
		return m.Points
	}
	return nil
}

// A container for simulation points for simulations of type CPV_BID.
type CpvBidSimulationPointList struct {
	// Projected metrics for a series of CPV bid amounts.
	Points               []*CpvBidSimulationPoint `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *CpvBidSimulationPointList) Reset()         { *m = CpvBidSimulationPointList{} }
func (m *CpvBidSimulationPointList) String() string { return proto.CompactTextString(m) }
func (*CpvBidSimulationPointList) ProtoMessage()    {}
func (*CpvBidSimulationPointList) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4802a06f80f525c, []int{2}
}

func (m *CpvBidSimulationPointList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CpvBidSimulationPointList.Unmarshal(m, b)
}
func (m *CpvBidSimulationPointList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CpvBidSimulationPointList.Marshal(b, m, deterministic)
}
func (m *CpvBidSimulationPointList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CpvBidSimulationPointList.Merge(m, src)
}
func (m *CpvBidSimulationPointList) XXX_Size() int {
	return xxx_messageInfo_CpvBidSimulationPointList.Size(m)
}
func (m *CpvBidSimulationPointList) XXX_DiscardUnknown() {
	xxx_messageInfo_CpvBidSimulationPointList.DiscardUnknown(m)
}

var xxx_messageInfo_CpvBidSimulationPointList proto.InternalMessageInfo

func (m *CpvBidSimulationPointList) GetPoints() []*CpvBidSimulationPoint {
	if m != nil {
		return m.Points
	}
	return nil
}

// A container for simulation points for simulations of type TARGET_CPA.
type TargetCpaSimulationPointList struct {
	// Projected metrics for a series of target CPA amounts.
	Points               []*TargetCpaSimulationPoint `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *TargetCpaSimulationPointList) Reset()         { *m = TargetCpaSimulationPointList{} }
func (m *TargetCpaSimulationPointList) String() string { return proto.CompactTextString(m) }
func (*TargetCpaSimulationPointList) ProtoMessage()    {}
func (*TargetCpaSimulationPointList) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4802a06f80f525c, []int{3}
}

func (m *TargetCpaSimulationPointList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetCpaSimulationPointList.Unmarshal(m, b)
}
func (m *TargetCpaSimulationPointList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetCpaSimulationPointList.Marshal(b, m, deterministic)
}
func (m *TargetCpaSimulationPointList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetCpaSimulationPointList.Merge(m, src)
}
func (m *TargetCpaSimulationPointList) XXX_Size() int {
	return xxx_messageInfo_TargetCpaSimulationPointList.Size(m)
}
func (m *TargetCpaSimulationPointList) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetCpaSimulationPointList.DiscardUnknown(m)
}

var xxx_messageInfo_TargetCpaSimulationPointList proto.InternalMessageInfo

func (m *TargetCpaSimulationPointList) GetPoints() []*TargetCpaSimulationPoint {
	if m != nil {
		return m.Points
	}
	return nil
}

// Projected metrics for a specific bid modifier amount.
type BidModifierSimulationPoint struct {
	// The simulated bid modifier upon which projected metrics are based.
	BidModifier *wrappers.DoubleValue `protobuf:"bytes,1,opt,name=bid_modifier,json=bidModifier,proto3" json:"bid_modifier,omitempty"`
	// Projected number of biddable conversions.
	BiddableConversions *wrappers.DoubleValue `protobuf:"bytes,2,opt,name=biddable_conversions,json=biddableConversions,proto3" json:"biddable_conversions,omitempty"`
	// Projected total value of biddable conversions.
	BiddableConversionsValue *wrappers.DoubleValue `protobuf:"bytes,3,opt,name=biddable_conversions_value,json=biddableConversionsValue,proto3" json:"biddable_conversions_value,omitempty"`
	// Projected number of clicks.
	Clicks *wrappers.Int64Value `protobuf:"bytes,4,opt,name=clicks,proto3" json:"clicks,omitempty"`
	// Projected cost in micros.
	CostMicros *wrappers.Int64Value `protobuf:"bytes,5,opt,name=cost_micros,json=costMicros,proto3" json:"cost_micros,omitempty"`
	// Projected number of impressions.
	Impressions *wrappers.Int64Value `protobuf:"bytes,6,opt,name=impressions,proto3" json:"impressions,omitempty"`
	// Projected number of top slot impressions.
	TopSlotImpressions *wrappers.Int64Value `protobuf:"bytes,7,opt,name=top_slot_impressions,json=topSlotImpressions,proto3" json:"top_slot_impressions,omitempty"`
	// Projected number of biddable conversions for the parent resource.
	ParentBiddableConversions *wrappers.DoubleValue `protobuf:"bytes,8,opt,name=parent_biddable_conversions,json=parentBiddableConversions,proto3" json:"parent_biddable_conversions,omitempty"`
	// Projected total value of biddable conversions for the parent resource.
	ParentBiddableConversionsValue *wrappers.DoubleValue `protobuf:"bytes,9,opt,name=parent_biddable_conversions_value,json=parentBiddableConversionsValue,proto3" json:"parent_biddable_conversions_value,omitempty"`
	// Projected number of clicks for the parent resource.
	ParentClicks *wrappers.Int64Value `protobuf:"bytes,10,opt,name=parent_clicks,json=parentClicks,proto3" json:"parent_clicks,omitempty"`
	// Projected cost in micros for the parent resource.
	ParentCostMicros *wrappers.Int64Value `protobuf:"bytes,11,opt,name=parent_cost_micros,json=parentCostMicros,proto3" json:"parent_cost_micros,omitempty"`
	// Projected number of impressions for the parent resource.
	ParentImpressions *wrappers.Int64Value `protobuf:"bytes,12,opt,name=parent_impressions,json=parentImpressions,proto3" json:"parent_impressions,omitempty"`
	// Projected number of top slot impressions for the parent resource.
	ParentTopSlotImpressions *wrappers.Int64Value `protobuf:"bytes,13,opt,name=parent_top_slot_impressions,json=parentTopSlotImpressions,proto3" json:"parent_top_slot_impressions,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}             `json:"-"`
	XXX_unrecognized         []byte               `json:"-"`
	XXX_sizecache            int32                `json:"-"`
}

func (m *BidModifierSimulationPoint) Reset()         { *m = BidModifierSimulationPoint{} }
func (m *BidModifierSimulationPoint) String() string { return proto.CompactTextString(m) }
func (*BidModifierSimulationPoint) ProtoMessage()    {}
func (*BidModifierSimulationPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4802a06f80f525c, []int{4}
}

func (m *BidModifierSimulationPoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BidModifierSimulationPoint.Unmarshal(m, b)
}
func (m *BidModifierSimulationPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BidModifierSimulationPoint.Marshal(b, m, deterministic)
}
func (m *BidModifierSimulationPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BidModifierSimulationPoint.Merge(m, src)
}
func (m *BidModifierSimulationPoint) XXX_Size() int {
	return xxx_messageInfo_BidModifierSimulationPoint.Size(m)
}
func (m *BidModifierSimulationPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_BidModifierSimulationPoint.DiscardUnknown(m)
}

var xxx_messageInfo_BidModifierSimulationPoint proto.InternalMessageInfo

func (m *BidModifierSimulationPoint) GetBidModifier() *wrappers.DoubleValue {
	if m != nil {
		return m.BidModifier
	}
	return nil
}

func (m *BidModifierSimulationPoint) GetBiddableConversions() *wrappers.DoubleValue {
	if m != nil {
		return m.BiddableConversions
	}
	return nil
}

func (m *BidModifierSimulationPoint) GetBiddableConversionsValue() *wrappers.DoubleValue {
	if m != nil {
		return m.BiddableConversionsValue
	}
	return nil
}

func (m *BidModifierSimulationPoint) GetClicks() *wrappers.Int64Value {
	if m != nil {
		return m.Clicks
	}
	return nil
}

func (m *BidModifierSimulationPoint) GetCostMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CostMicros
	}
	return nil
}

func (m *BidModifierSimulationPoint) GetImpressions() *wrappers.Int64Value {
	if m != nil {
		return m.Impressions
	}
	return nil
}

func (m *BidModifierSimulationPoint) GetTopSlotImpressions() *wrappers.Int64Value {
	if m != nil {
		return m.TopSlotImpressions
	}
	return nil
}

func (m *BidModifierSimulationPoint) GetParentBiddableConversions() *wrappers.DoubleValue {
	if m != nil {
		return m.ParentBiddableConversions
	}
	return nil
}

func (m *BidModifierSimulationPoint) GetParentBiddableConversionsValue() *wrappers.DoubleValue {
	if m != nil {
		return m.ParentBiddableConversionsValue
	}
	return nil
}

func (m *BidModifierSimulationPoint) GetParentClicks() *wrappers.Int64Value {
	if m != nil {
		return m.ParentClicks
	}
	return nil
}

func (m *BidModifierSimulationPoint) GetParentCostMicros() *wrappers.Int64Value {
	if m != nil {
		return m.ParentCostMicros
	}
	return nil
}

func (m *BidModifierSimulationPoint) GetParentImpressions() *wrappers.Int64Value {
	if m != nil {
		return m.ParentImpressions
	}
	return nil
}

func (m *BidModifierSimulationPoint) GetParentTopSlotImpressions() *wrappers.Int64Value {
	if m != nil {
		return m.ParentTopSlotImpressions
	}
	return nil
}

// Projected metrics for a specific CPC bid amount.
type CpcBidSimulationPoint struct {
	// The simulated CPC bid upon which projected metrics are based.
	CpcBidMicros *wrappers.Int64Value `protobuf:"bytes,1,opt,name=cpc_bid_micros,json=cpcBidMicros,proto3" json:"cpc_bid_micros,omitempty"`
	// Projected number of biddable conversions.
	BiddableConversions *wrappers.DoubleValue `protobuf:"bytes,2,opt,name=biddable_conversions,json=biddableConversions,proto3" json:"biddable_conversions,omitempty"`
	// Projected total value of biddable conversions.
	BiddableConversionsValue *wrappers.DoubleValue `protobuf:"bytes,3,opt,name=biddable_conversions_value,json=biddableConversionsValue,proto3" json:"biddable_conversions_value,omitempty"`
	// Projected number of clicks.
	Clicks *wrappers.Int64Value `protobuf:"bytes,4,opt,name=clicks,proto3" json:"clicks,omitempty"`
	// Projected cost in micros.
	CostMicros *wrappers.Int64Value `protobuf:"bytes,5,opt,name=cost_micros,json=costMicros,proto3" json:"cost_micros,omitempty"`
	// Projected number of impressions.
	Impressions *wrappers.Int64Value `protobuf:"bytes,6,opt,name=impressions,proto3" json:"impressions,omitempty"`
	// Projected number of top slot impressions.
	// Display network does not support this field at the ad group level.
	TopSlotImpressions   *wrappers.Int64Value `protobuf:"bytes,7,opt,name=top_slot_impressions,json=topSlotImpressions,proto3" json:"top_slot_impressions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CpcBidSimulationPoint) Reset()         { *m = CpcBidSimulationPoint{} }
func (m *CpcBidSimulationPoint) String() string { return proto.CompactTextString(m) }
func (*CpcBidSimulationPoint) ProtoMessage()    {}
func (*CpcBidSimulationPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4802a06f80f525c, []int{5}
}

func (m *CpcBidSimulationPoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CpcBidSimulationPoint.Unmarshal(m, b)
}
func (m *CpcBidSimulationPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CpcBidSimulationPoint.Marshal(b, m, deterministic)
}
func (m *CpcBidSimulationPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CpcBidSimulationPoint.Merge(m, src)
}
func (m *CpcBidSimulationPoint) XXX_Size() int {
	return xxx_messageInfo_CpcBidSimulationPoint.Size(m)
}
func (m *CpcBidSimulationPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_CpcBidSimulationPoint.DiscardUnknown(m)
}

var xxx_messageInfo_CpcBidSimulationPoint proto.InternalMessageInfo

func (m *CpcBidSimulationPoint) GetCpcBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidMicros
	}
	return nil
}

func (m *CpcBidSimulationPoint) GetBiddableConversions() *wrappers.DoubleValue {
	if m != nil {
		return m.BiddableConversions
	}
	return nil
}

func (m *CpcBidSimulationPoint) GetBiddableConversionsValue() *wrappers.DoubleValue {
	if m != nil {
		return m.BiddableConversionsValue
	}
	return nil
}

func (m *CpcBidSimulationPoint) GetClicks() *wrappers.Int64Value {
	if m != nil {
		return m.Clicks
	}
	return nil
}

func (m *CpcBidSimulationPoint) GetCostMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CostMicros
	}
	return nil
}

func (m *CpcBidSimulationPoint) GetImpressions() *wrappers.Int64Value {
	if m != nil {
		return m.Impressions
	}
	return nil
}

func (m *CpcBidSimulationPoint) GetTopSlotImpressions() *wrappers.Int64Value {
	if m != nil {
		return m.TopSlotImpressions
	}
	return nil
}

// Projected metrics for a specific CPV bid amount.
type CpvBidSimulationPoint struct {
	// The simulated CPV bid upon which projected metrics are based.
	CpvBidMicros *wrappers.Int64Value `protobuf:"bytes,1,opt,name=cpv_bid_micros,json=cpvBidMicros,proto3" json:"cpv_bid_micros,omitempty"`
	// Projected cost in micros.
	CostMicros *wrappers.Int64Value `protobuf:"bytes,2,opt,name=cost_micros,json=costMicros,proto3" json:"cost_micros,omitempty"`
	// Projected number of impressions.
	Impressions          *wrappers.Int64Value `protobuf:"bytes,3,opt,name=impressions,proto3" json:"impressions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CpvBidSimulationPoint) Reset()         { *m = CpvBidSimulationPoint{} }
func (m *CpvBidSimulationPoint) String() string { return proto.CompactTextString(m) }
func (*CpvBidSimulationPoint) ProtoMessage()    {}
func (*CpvBidSimulationPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4802a06f80f525c, []int{6}
}

func (m *CpvBidSimulationPoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CpvBidSimulationPoint.Unmarshal(m, b)
}
func (m *CpvBidSimulationPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CpvBidSimulationPoint.Marshal(b, m, deterministic)
}
func (m *CpvBidSimulationPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CpvBidSimulationPoint.Merge(m, src)
}
func (m *CpvBidSimulationPoint) XXX_Size() int {
	return xxx_messageInfo_CpvBidSimulationPoint.Size(m)
}
func (m *CpvBidSimulationPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_CpvBidSimulationPoint.DiscardUnknown(m)
}

var xxx_messageInfo_CpvBidSimulationPoint proto.InternalMessageInfo

func (m *CpvBidSimulationPoint) GetCpvBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpvBidMicros
	}
	return nil
}

func (m *CpvBidSimulationPoint) GetCostMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CostMicros
	}
	return nil
}

func (m *CpvBidSimulationPoint) GetImpressions() *wrappers.Int64Value {
	if m != nil {
		return m.Impressions
	}
	return nil
}

// Projected metrics for a specific target CPA amount.
type TargetCpaSimulationPoint struct {
	// The simulated target CPA upon which projected metrics are based.
	TargetCpaMicros *wrappers.Int64Value `protobuf:"bytes,1,opt,name=target_cpa_micros,json=targetCpaMicros,proto3" json:"target_cpa_micros,omitempty"`
	// Projected number of biddable conversions.
	BiddableConversions *wrappers.DoubleValue `protobuf:"bytes,2,opt,name=biddable_conversions,json=biddableConversions,proto3" json:"biddable_conversions,omitempty"`
	// Projected total value of biddable conversions.
	BiddableConversionsValue *wrappers.DoubleValue `protobuf:"bytes,3,opt,name=biddable_conversions_value,json=biddableConversionsValue,proto3" json:"biddable_conversions_value,omitempty"`
	// Projected number of clicks.
	Clicks *wrappers.Int64Value `protobuf:"bytes,4,opt,name=clicks,proto3" json:"clicks,omitempty"`
	// Projected cost in micros.
	CostMicros *wrappers.Int64Value `protobuf:"bytes,5,opt,name=cost_micros,json=costMicros,proto3" json:"cost_micros,omitempty"`
	// Projected number of impressions.
	Impressions *wrappers.Int64Value `protobuf:"bytes,6,opt,name=impressions,proto3" json:"impressions,omitempty"`
	// Projected number of top slot impressions.
	// Display network does not support this field at the ad group level.
	TopSlotImpressions   *wrappers.Int64Value `protobuf:"bytes,7,opt,name=top_slot_impressions,json=topSlotImpressions,proto3" json:"top_slot_impressions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TargetCpaSimulationPoint) Reset()         { *m = TargetCpaSimulationPoint{} }
func (m *TargetCpaSimulationPoint) String() string { return proto.CompactTextString(m) }
func (*TargetCpaSimulationPoint) ProtoMessage()    {}
func (*TargetCpaSimulationPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4802a06f80f525c, []int{7}
}

func (m *TargetCpaSimulationPoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetCpaSimulationPoint.Unmarshal(m, b)
}
func (m *TargetCpaSimulationPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetCpaSimulationPoint.Marshal(b, m, deterministic)
}
func (m *TargetCpaSimulationPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetCpaSimulationPoint.Merge(m, src)
}
func (m *TargetCpaSimulationPoint) XXX_Size() int {
	return xxx_messageInfo_TargetCpaSimulationPoint.Size(m)
}
func (m *TargetCpaSimulationPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetCpaSimulationPoint.DiscardUnknown(m)
}

var xxx_messageInfo_TargetCpaSimulationPoint proto.InternalMessageInfo

func (m *TargetCpaSimulationPoint) GetTargetCpaMicros() *wrappers.Int64Value {
	if m != nil {
		return m.TargetCpaMicros
	}
	return nil
}

func (m *TargetCpaSimulationPoint) GetBiddableConversions() *wrappers.DoubleValue {
	if m != nil {
		return m.BiddableConversions
	}
	return nil
}

func (m *TargetCpaSimulationPoint) GetBiddableConversionsValue() *wrappers.DoubleValue {
	if m != nil {
		return m.BiddableConversionsValue
	}
	return nil
}

func (m *TargetCpaSimulationPoint) GetClicks() *wrappers.Int64Value {
	if m != nil {
		return m.Clicks
	}
	return nil
}

func (m *TargetCpaSimulationPoint) GetCostMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CostMicros
	}
	return nil
}

func (m *TargetCpaSimulationPoint) GetImpressions() *wrappers.Int64Value {
	if m != nil {
		return m.Impressions
	}
	return nil
}

func (m *TargetCpaSimulationPoint) GetTopSlotImpressions() *wrappers.Int64Value {
	if m != nil {
		return m.TopSlotImpressions
	}
	return nil
}

func init() {
	proto.RegisterType((*BidModifierSimulationPointList)(nil), "google.ads.googleads.v1.common.BidModifierSimulationPointList")
	proto.RegisterType((*CpcBidSimulationPointList)(nil), "google.ads.googleads.v1.common.CpcBidSimulationPointList")
	proto.RegisterType((*CpvBidSimulationPointList)(nil), "google.ads.googleads.v1.common.CpvBidSimulationPointList")
	proto.RegisterType((*TargetCpaSimulationPointList)(nil), "google.ads.googleads.v1.common.TargetCpaSimulationPointList")
	proto.RegisterType((*BidModifierSimulationPoint)(nil), "google.ads.googleads.v1.common.BidModifierSimulationPoint")
	proto.RegisterType((*CpcBidSimulationPoint)(nil), "google.ads.googleads.v1.common.CpcBidSimulationPoint")
	proto.RegisterType((*CpvBidSimulationPoint)(nil), "google.ads.googleads.v1.common.CpvBidSimulationPoint")
	proto.RegisterType((*TargetCpaSimulationPoint)(nil), "google.ads.googleads.v1.common.TargetCpaSimulationPoint")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/common/simulation.proto", fileDescriptor_a4802a06f80f525c)
}

var fileDescriptor_a4802a06f80f525c = []byte{
	// 699 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x96, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x95, 0x76, 0x94, 0xe1, 0x6e, 0x8c, 0x99, 0x21, 0x65, 0xdd, 0x54, 0x8d, 0x5c, 0xed,
	0x2a, 0x51, 0x37, 0x40, 0x28, 0x80, 0xa0, 0xed, 0xa4, 0x69, 0x88, 0x8a, 0x69, 0x9b, 0x7a, 0x31,
	0x55, 0x8a, 0xf2, 0x6f, 0x91, 0x21, 0x89, 0xad, 0xd8, 0x0d, 0x8f, 0x00, 0xb7, 0xbc, 0x02, 0x97,
	0x3c, 0x0a, 0x2f, 0xc0, 0x3b, 0xec, 0x29, 0x50, 0x62, 0x27, 0x0d, 0x5b, 0xb2, 0xb4, 0x82, 0x9b,
	0x49, 0xbb, 0xaa, 0x53, 0x9f, 0xef, 0xf7, 0xa5, 0xe7, 0x3b, 0x72, 0x0d, 0x34, 0x0f, 0x63, 0xcf,
	0x77, 0x35, 0xd3, 0xa1, 0x62, 0x99, 0xac, 0xe2, 0x9e, 0x66, 0xe3, 0x20, 0xc0, 0xa1, 0x46, 0x51,
	0x30, 0xf5, 0x4d, 0x86, 0x70, 0xa8, 0x92, 0x08, 0x33, 0x0c, 0xbb, 0xbc, 0x4a, 0x35, 0x1d, 0xaa,
	0xe6, 0x02, 0x35, 0xee, 0xa9, 0x5c, 0xd0, 0xd9, 0xce, 0x80, 0x04, 0x69, 0x66, 0x18, 0x62, 0x96,
	0x8a, 0x29, 0x57, 0x77, 0x84, 0x5a, 0x4b, 0x9f, 0xac, 0xe9, 0x85, 0xf6, 0x25, 0x32, 0x09, 0x71,
	0x23, 0xb1, 0xaf, 0x30, 0xd0, 0x1d, 0x20, 0x67, 0x84, 0x1d, 0x74, 0x81, 0xdc, 0xe8, 0x34, 0x37,
	0x3f, 0xc6, 0x28, 0x64, 0x1f, 0x10, 0x65, 0xf0, 0x04, 0xb4, 0x48, 0xf2, 0x40, 0x65, 0x69, 0xa7,
	0xb9, 0xdb, 0xde, 0xd3, 0xd5, 0x9b, 0x5f, 0x48, 0xad, 0xe6, 0x9d, 0x08, 0x92, 0xf2, 0x09, 0x6c,
	0x0e, 0x89, 0x3d, 0x40, 0x4e, 0x99, 0xe1, 0xe8, 0x8a, 0xe1, 0xf3, 0x3a, 0xc3, 0x52, 0xd4, 0xdf,
	0x5e, 0xf1, 0xff, 0xf3, 0x8a, 0x6f, 0xf0, 0x22, 0x60, 0xfb, 0xcc, 0x8c, 0x3c, 0x97, 0x0d, 0x89,
	0x59, 0x66, 0x77, 0x7c, 0xc5, 0xee, 0x65, 0x9d, 0x5d, 0x15, 0x2d, 0x77, 0xfc, 0xba, 0x0c, 0x3a,
	0xd5, 0x0d, 0x87, 0x6f, 0xc1, 0x8a, 0x85, 0x1c, 0x23, 0x10, 0xdb, 0xb2, 0xb4, 0x23, 0xed, 0xb6,
	0xf7, 0xb6, 0x33, 0xdb, 0x6c, 0x2a, 0xd4, 0x03, 0x3c, 0xb5, 0x7c, 0x77, 0x6c, 0xfa, 0x53, 0xf7,
	0xa4, 0x6d, 0xcd, 0x78, 0xf0, 0x23, 0xd8, 0xb0, 0x90, 0xe3, 0x98, 0x96, 0xef, 0x1a, 0x36, 0x0e,
	0x63, 0x37, 0xa2, 0xc9, 0x74, 0xc9, 0x8d, 0x39, 0x40, 0x8f, 0x33, 0xe5, 0x70, 0x26, 0x84, 0xe7,
	0xa0, 0x53, 0x06, 0x34, 0xe2, 0x44, 0x22, 0x37, 0xe7, 0xc0, 0xca, 0x25, 0xd8, 0x74, 0x07, 0xee,
	0x83, 0x96, 0xed, 0x23, 0xfb, 0x33, 0x95, 0x97, 0x52, 0xce, 0xd6, 0x35, 0xce, 0x51, 0xc8, 0x5e,
	0x3c, 0xe3, 0x18, 0x51, 0x0a, 0x5f, 0x83, 0xb6, 0x8d, 0x29, 0x33, 0x02, 0x64, 0x47, 0x98, 0xca,
	0xf7, 0xea, 0x95, 0x20, 0xa9, 0x1f, 0xa5, 0xe5, 0xf0, 0x0d, 0x68, 0xa3, 0x80, 0x44, 0x2e, 0xe5,
	0x6d, 0x69, 0xd5, 0xab, 0x8b, 0xf5, 0x70, 0x04, 0x36, 0x18, 0x26, 0x06, 0xf5, 0x31, 0x33, 0x8a,
	0x9c, 0xfb, 0xf5, 0x1c, 0xc8, 0x30, 0x39, 0xf5, 0x31, 0x3b, 0x2a, 0xe0, 0x26, 0x60, 0x8b, 0x98,
	0x91, 0x1b, 0x32, 0xa3, 0x34, 0xb4, 0xe5, 0x39, 0xba, 0xbb, 0xc9, 0x01, 0x83, 0x92, 0xe8, 0x3c,
	0xf0, 0xf4, 0x06, 0xba, 0x48, 0xf0, 0xc1, 0x1c, 0x1e, 0xdd, 0x4a, 0x0f, 0x9e, 0xe3, 0x3b, 0xb0,
	0x2a, 0x8c, 0x44, 0x9c, 0xa0, 0xbe, 0x1d, 0x2b, 0x5c, 0x31, 0xe4, 0xa1, 0x1e, 0x01, 0x98, 0x11,
	0x0a, 0xd9, 0xb6, 0xeb, 0x31, 0x8f, 0x04, 0x66, 0x96, 0xf0, 0xfb, 0x1c, 0x55, 0x0c, 0x68, 0xa5,
	0x1e, 0xb5, 0xce, 0x65, 0xc5, 0x7c, 0xce, 0xf3, 0x7c, 0x4a, 0x53, 0x5f, 0xad, 0x87, 0xca, 0x5c,
	0x7f, 0x76, 0x2d, 0x7b, 0xe5, 0xdb, 0x12, 0x78, 0x52, 0x7a, 0x12, 0xc2, 0x3e, 0x78, 0x68, 0x13,
	0xdb, 0x48, 0x0f, 0x02, 0xde, 0x08, 0x69, 0x8e, 0x7e, 0xda, 0x29, 0x4b, 0x34, 0xe1, 0xee, 0x18,
	0xb8, 0x45, 0xc7, 0x80, 0xf2, 0x5b, 0x4a, 0x46, 0x21, 0xae, 0x1a, 0x85, 0x78, 0xf1, 0x51, 0x88,
	0x67, 0xa3, 0x70, 0xa5, 0x51, 0x8d, 0x7f, 0x6a, 0x54, 0x73, 0xb1, 0x46, 0x29, 0xdf, 0x97, 0x80,
	0x5c, 0xf5, 0x9f, 0x08, 0x0f, 0xc1, 0x3a, 0x4b, 0xf7, 0x0c, 0x9b, 0x98, 0x0b, 0xfc, 0xbe, 0x35,
	0x96, 0x11, 0xef, 0xa6, 0xfd, 0xf6, 0x4d, 0xfb, 0xe0, 0x52, 0x02, 0x8a, 0x8d, 0x83, 0x9a, 0xab,
	0xd4, 0x60, 0xad, 0x30, 0x2e, 0x09, 0xf9, 0x58, 0x3a, 0x3f, 0x10, 0x12, 0x0f, 0xfb, 0x66, 0xe8,
	0xa9, 0x38, 0xf2, 0x34, 0xcf, 0x0d, 0x53, 0xdf, 0xec, 0x6e, 0x4e, 0x10, 0xad, 0xba, 0xaa, 0xbf,
	0xe2, 0x1f, 0x3f, 0x1a, 0xcd, 0xc3, 0x7e, 0xff, 0x67, 0xa3, 0x7b, 0xc8, 0x61, 0x7d, 0x87, 0xaa,
	0x7c, 0x99, 0xac, 0xc6, 0x3d, 0x75, 0x98, 0x96, 0xfd, 0xca, 0x0a, 0x26, 0x7d, 0x87, 0x4e, 0xf2,
	0x82, 0xc9, 0xb8, 0x37, 0xe1, 0x05, 0x97, 0x0d, 0x85, 0x7f, 0xab, 0xeb, 0x7d, 0x87, 0xea, 0x7a,
	0x5e, 0xa2, 0xeb, 0xe3, 0x9e, 0xae, 0xf3, 0x22, 0xab, 0x95, 0xbe, 0xdd, 0xfe, 0x9f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x77, 0xd2, 0x2f, 0xaa, 0x47, 0x0c, 0x00, 0x00,
}

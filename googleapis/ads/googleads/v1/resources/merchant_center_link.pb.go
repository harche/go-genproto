// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/merchant_center_link.proto

package resources

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A data sharing connection, proposed or in use,
// between a Google Ads Customer and a Merchant Center account.
type MerchantCenterLink struct {
	// The resource name of the merchant center link.
	// Merchant center link resource names have the form:
	//
	// `customers/{customer_id}/merchantCenterLinks/{merchant_center_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the Merchant Center account.
	// This field is readonly.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the Merchant Center account.
	// This field is readonly.
	MerchantCenterAccountName *wrappers.StringValue `protobuf:"bytes,4,opt,name=merchant_center_account_name,json=merchantCenterAccountName,proto3" json:"merchant_center_account_name,omitempty"`
	// The status of the link.
	Status               enums.MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                    `json:"-"`
	XXX_unrecognized     []byte                                                      `json:"-"`
	XXX_sizecache        int32                                                       `json:"-"`
}

func (m *MerchantCenterLink) Reset()         { *m = MerchantCenterLink{} }
func (m *MerchantCenterLink) String() string { return proto.CompactTextString(m) }
func (*MerchantCenterLink) ProtoMessage()    {}
func (*MerchantCenterLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_068c27e743addf15, []int{0}
}

func (m *MerchantCenterLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantCenterLink.Unmarshal(m, b)
}
func (m *MerchantCenterLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantCenterLink.Marshal(b, m, deterministic)
}
func (m *MerchantCenterLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantCenterLink.Merge(m, src)
}
func (m *MerchantCenterLink) XXX_Size() int {
	return xxx_messageInfo_MerchantCenterLink.Size(m)
}
func (m *MerchantCenterLink) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantCenterLink.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantCenterLink proto.InternalMessageInfo

func (m *MerchantCenterLink) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *MerchantCenterLink) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *MerchantCenterLink) GetMerchantCenterAccountName() *wrappers.StringValue {
	if m != nil {
		return m.MerchantCenterAccountName
	}
	return nil
}

func (m *MerchantCenterLink) GetStatus() enums.MerchantCenterLinkStatusEnum_MerchantCenterLinkStatus {
	if m != nil {
		return m.Status
	}
	return enums.MerchantCenterLinkStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*MerchantCenterLink)(nil), "google.ads.googleads.v1.resources.MerchantCenterLink")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/merchant_center_link.proto", fileDescriptor_068c27e743addf15)
}

var fileDescriptor_068c27e743addf15 = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcf, 0x6a, 0xdb, 0x30,
	0x18, 0xc7, 0xce, 0x16, 0x98, 0xf7, 0xe7, 0xe0, 0xcb, 0xb2, 0x2c, 0x8c, 0x64, 0x23, 0x10, 0x18,
	0xc8, 0x38, 0x1b, 0x3b, 0x68, 0x83, 0xe1, 0x8c, 0x11, 0x36, 0xb6, 0x12, 0x9c, 0xe2, 0x43, 0x71,
	0x31, 0x8a, 0xad, 0xba, 0x22, 0xb6, 0x64, 0x24, 0x39, 0x7d, 0x81, 0x3e, 0x49, 0x8f, 0xbd, 0xf6,
	0x2d, 0xfa, 0x28, 0x7d, 0x8a, 0x12, 0xc9, 0x32, 0xb4, 0x69, 0xda, 0xdb, 0x97, 0xfc, 0xfe, 0x7d,
	0xbf, 0xcf, 0x72, 0x7e, 0xe4, 0x8c, 0xe5, 0x05, 0xf6, 0x50, 0x26, 0x3c, 0x3d, 0x6e, 0xa7, 0x8d,
	0xef, 0x71, 0x2c, 0x58, 0xcd, 0x53, 0x2c, 0xbc, 0x12, 0xf3, 0xf4, 0x14, 0x51, 0x99, 0xa4, 0x98,
	0x4a, 0xcc, 0x93, 0x82, 0xd0, 0x35, 0xa8, 0x38, 0x93, 0xcc, 0x1d, 0x69, 0x09, 0x40, 0x99, 0x00,
	0xad, 0x1a, 0x6c, 0x7c, 0xd0, 0xaa, 0xfb, 0x3f, 0xf7, 0x05, 0x60, 0x5a, 0x97, 0x0f, 0x9b, 0x27,
	0x42, 0x22, 0x59, 0x0b, 0x9d, 0xd1, 0x1f, 0x18, 0x83, 0x8a, 0x78, 0x88, 0x52, 0x26, 0x91, 0x24,
	0x8c, 0x1a, 0xf4, 0x43, 0x83, 0xaa, 0x5f, 0xab, 0xfa, 0xc4, 0x3b, 0xe3, 0xa8, 0xaa, 0x30, 0x6f,
	0xf0, 0x8f, 0x57, 0xb6, 0xe3, 0xfe, 0x6f, 0x32, 0x7e, 0xa9, 0x88, 0x7f, 0x84, 0xae, 0xdd, 0x4f,
	0xce, 0x6b, 0xb3, 0x62, 0x42, 0x51, 0x89, 0x7b, 0xd6, 0xd0, 0x9a, 0xbc, 0x08, 0x5f, 0x99, 0x3f,
	0x0f, 0x50, 0x89, 0xdd, 0xcf, 0x8e, 0x4d, 0xb2, 0x5e, 0x67, 0x68, 0x4d, 0x5e, 0x4e, 0xdf, 0x37,
	0xfd, 0x80, 0x09, 0x02, 0x7f, 0xa8, 0xfc, 0xf6, 0x35, 0x42, 0x45, 0x8d, 0x43, 0x9b, 0x64, 0xee,
	0xb1, 0x33, 0xb8, 0xdf, 0x05, 0xa5, 0x29, 0xab, 0xa9, 0xd4, 0x01, 0xcf, 0x94, 0xcd, 0x60, 0xc7,
	0x66, 0x29, 0x39, 0xa1, 0xb9, 0xf6, 0x79, 0x57, 0xde, 0xd9, 0x34, 0xd0, 0x7a, 0xb5, 0x4b, 0xe1,
	0x74, 0xf5, 0x55, 0x7a, 0xcf, 0x87, 0xd6, 0xe4, 0xcd, 0xf4, 0x10, 0xec, 0x3b, 0xbd, 0xba, 0x2b,
	0xd8, 0xed, 0xbc, 0x54, 0xf2, 0xdf, 0xb4, 0x2e, 0xf7, 0x82, 0x61, 0x93, 0x31, 0x3b, 0xb7, 0x9d,
	0x71, 0xca, 0x4a, 0xf0, 0xe4, 0xe7, 0x9d, 0xbd, 0xdd, 0xf5, 0x5a, 0x6c, 0xab, 0x2d, 0xac, 0xa3,
	0xbf, 0x8d, 0x3a, 0x67, 0x05, 0xa2, 0x39, 0x60, 0x3c, 0xf7, 0x72, 0x4c, 0x55, 0x71, 0xf3, 0x12,
	0x2a, 0x22, 0x1e, 0x79, 0x79, 0xdf, 0xdb, 0xe9, 0xc2, 0xee, 0xcc, 0x83, 0xe0, 0xd2, 0x1e, 0xcd,
	0xb5, 0x65, 0x90, 0x09, 0xa0, 0xc7, 0xed, 0x14, 0xf9, 0x20, 0x34, 0xcc, 0x6b, 0xc3, 0x89, 0x83,
	0x4c, 0xc4, 0x2d, 0x27, 0x8e, 0xfc, 0xb8, 0xe5, 0xdc, 0xd8, 0x63, 0x0d, 0x40, 0x18, 0x64, 0x02,
	0xc2, 0x96, 0x05, 0x61, 0xe4, 0x43, 0xd8, 0xf2, 0x56, 0x5d, 0xb5, 0xec, 0x97, 0xdb, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x3e, 0x9d, 0xc5, 0x55, 0x25, 0x03, 0x00, 0x00,
}

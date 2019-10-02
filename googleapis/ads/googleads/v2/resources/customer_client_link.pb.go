// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/customer_client_link.proto

package resources

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// Represents customer client link relationship.
type CustomerClientLink struct {
	// Name of the resource.
	// CustomerClientLink resource names have the form:
	// `customers/{customer_id}/customerClientLinks/{client_customer_id}~{manager_link_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The client customer linked to this customer.
	ClientCustomer *wrappers.StringValue `protobuf:"bytes,3,opt,name=client_customer,json=clientCustomer,proto3" json:"client_customer,omitempty"`
	// This is uniquely identifies a customer client link. Read only.
	ManagerLinkId *wrappers.Int64Value `protobuf:"bytes,4,opt,name=manager_link_id,json=managerLinkId,proto3" json:"manager_link_id,omitempty"`
	// This is the status of the link between client and manager.
	Status enums.ManagerLinkStatusEnum_ManagerLinkStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v2.enums.ManagerLinkStatusEnum_ManagerLinkStatus" json:"status,omitempty"`
	// The visibility of the link. Users can choose whether or not to see hidden
	// links in the AdWords UI.
	// Default value is false
	Hidden               *wrappers.BoolValue `protobuf:"bytes,6,opt,name=hidden,proto3" json:"hidden,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CustomerClientLink) Reset()         { *m = CustomerClientLink{} }
func (m *CustomerClientLink) String() string { return proto.CompactTextString(m) }
func (*CustomerClientLink) ProtoMessage()    {}
func (*CustomerClientLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_f87ffe96f69e22f8, []int{0}
}

func (m *CustomerClientLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerClientLink.Unmarshal(m, b)
}
func (m *CustomerClientLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerClientLink.Marshal(b, m, deterministic)
}
func (m *CustomerClientLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerClientLink.Merge(m, src)
}
func (m *CustomerClientLink) XXX_Size() int {
	return xxx_messageInfo_CustomerClientLink.Size(m)
}
func (m *CustomerClientLink) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerClientLink.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerClientLink proto.InternalMessageInfo

func (m *CustomerClientLink) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CustomerClientLink) GetClientCustomer() *wrappers.StringValue {
	if m != nil {
		return m.ClientCustomer
	}
	return nil
}

func (m *CustomerClientLink) GetManagerLinkId() *wrappers.Int64Value {
	if m != nil {
		return m.ManagerLinkId
	}
	return nil
}

func (m *CustomerClientLink) GetStatus() enums.ManagerLinkStatusEnum_ManagerLinkStatus {
	if m != nil {
		return m.Status
	}
	return enums.ManagerLinkStatusEnum_UNSPECIFIED
}

func (m *CustomerClientLink) GetHidden() *wrappers.BoolValue {
	if m != nil {
		return m.Hidden
	}
	return nil
}

func init() {
	proto.RegisterType((*CustomerClientLink)(nil), "google.ads.googleads.v2.resources.CustomerClientLink")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/customer_client_link.proto", fileDescriptor_f87ffe96f69e22f8)
}

var fileDescriptor_f87ffe96f69e22f8 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xdd, 0x6a, 0x14, 0x31,
	0x18, 0x65, 0xa6, 0xba, 0x60, 0xb4, 0x2d, 0xcc, 0x8d, 0xc3, 0x5a, 0x64, 0xab, 0x14, 0xf6, 0x2a,
	0x81, 0x28, 0x0a, 0xd1, 0x9b, 0xd9, 0xa5, 0x96, 0x8a, 0x4a, 0xd9, 0xc2, 0x5e, 0xc8, 0xe2, 0x92,
	0x6e, 0x62, 0x0c, 0x9d, 0x49, 0x86, 0x24, 0x53, 0x5f, 0xc0, 0x27, 0xf1, 0xd2, 0x47, 0xf1, 0x0d,
	0x7c, 0x05, 0x9f, 0x42, 0x36, 0x3f, 0x23, 0x32, 0xac, 0xbd, 0x3b, 0xc9, 0x77, 0xce, 0xf9, 0xce,
	0xf7, 0x25, 0xe0, 0xb5, 0xd0, 0x5a, 0xd4, 0x1c, 0x51, 0x66, 0x51, 0x80, 0x5b, 0x74, 0x83, 0x91,
	0xe1, 0x56, 0x77, 0x66, 0xc3, 0x2d, 0xda, 0x74, 0xd6, 0xe9, 0x86, 0x9b, 0xf5, 0xa6, 0x96, 0x5c,
	0xb9, 0x75, 0x2d, 0xd5, 0x35, 0x6c, 0x8d, 0x76, 0xba, 0x38, 0x0e, 0x12, 0x48, 0x99, 0x85, 0xbd,
	0x1a, 0xde, 0x60, 0xd8, 0xab, 0xc7, 0x2f, 0x77, 0x35, 0xe0, 0xaa, 0x6b, 0x2c, 0x6a, 0xa8, 0xa2,
	0x82, 0x1b, 0x6f, 0xba, 0xb6, 0x8e, 0xba, 0xce, 0x06, 0xef, 0xf1, 0xe3, 0x28, 0xf4, 0xa7, 0xab,
	0xee, 0x33, 0xfa, 0x6a, 0x68, 0xdb, 0x72, 0x93, 0xea, 0x47, 0xc9, 0xb8, 0x95, 0x88, 0x2a, 0xa5,
	0x1d, 0x75, 0x52, 0xab, 0x58, 0x7d, 0xf2, 0x2b, 0x07, 0xc5, 0x3c, 0x06, 0x9f, 0xfb, 0xdc, 0xef,
	0xa4, 0xba, 0x2e, 0x9e, 0x82, 0xfd, 0x14, 0x6d, 0xad, 0x68, 0xc3, 0xcb, 0x6c, 0x92, 0x4d, 0xef,
	0x2d, 0x1e, 0xa4, 0xcb, 0x0f, 0xb4, 0xe1, 0xc5, 0x29, 0x38, 0x8c, 0xa3, 0xa6, 0xd1, 0xcb, 0xbd,
	0x49, 0x36, 0xbd, 0x8f, 0x8f, 0xe2, 0x90, 0x30, 0x65, 0x82, 0x97, 0xce, 0x48, 0x25, 0x96, 0xb4,
	0xee, 0xf8, 0xe2, 0x20, 0x88, 0x52, 0xd7, 0x62, 0x0e, 0x0e, 0xff, 0x99, 0x4e, 0xb2, 0xf2, 0x8e,
	0xb7, 0x79, 0x34, 0xb0, 0x39, 0x57, 0xee, 0xc5, 0xf3, 0xe0, 0xb2, 0x1f, 0x35, 0xdb, 0xb8, 0xe7,
	0xac, 0xf8, 0x04, 0x46, 0x61, 0x2b, 0xe5, 0xdd, 0x49, 0x36, 0x3d, 0xc0, 0x6f, 0xe0, 0xae, 0x95,
	0xfb, 0x7d, 0xc2, 0xf7, 0x7f, 0xd5, 0x97, 0x5e, 0x77, 0xaa, 0xba, 0x66, 0x78, 0xbb, 0x88, 0xae,
	0x05, 0x06, 0xa3, 0x2f, 0x92, 0x31, 0xae, 0xca, 0x91, 0xcf, 0x36, 0x1e, 0x64, 0x9b, 0x69, 0x5d,
	0x87, 0x68, 0x91, 0x39, 0xfb, 0x96, 0x83, 0x93, 0x8d, 0x6e, 0xe0, 0xad, 0x8f, 0x3f, 0x7b, 0x38,
	0x7c, 0x82, 0x8b, 0xad, 0xef, 0x45, 0xf6, 0xf1, 0x6d, 0x54, 0x0b, 0x5d, 0x53, 0x25, 0xa0, 0x36,
	0x02, 0x09, 0xae, 0x7c, 0xd7, 0xf4, 0x4f, 0x5a, 0x69, 0xff, 0xf3, 0x2f, 0x5f, 0xf5, 0xe8, 0x7b,
	0xbe, 0x77, 0x56, 0x55, 0x3f, 0xf2, 0xe3, 0xb3, 0x60, 0x59, 0x31, 0x0b, 0x03, 0xdc, 0xa2, 0x25,
	0x86, 0x8b, 0xc4, 0xfc, 0x99, 0x38, 0xab, 0x8a, 0xd9, 0x55, 0xcf, 0x59, 0x2d, 0xf1, 0xaa, 0xe7,
	0xfc, 0xce, 0x4f, 0x42, 0x81, 0x90, 0x8a, 0x59, 0x42, 0x7a, 0x16, 0x21, 0x4b, 0x4c, 0x48, 0xcf,
	0xbb, 0x1a, 0xf9, 0xb0, 0xcf, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xde, 0x5d, 0x15, 0x90, 0x43,
	0x03, 0x00, 0x00,
}

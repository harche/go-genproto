// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/customer_client.proto

package resources

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

// A link between the given customer and a client customer. CustomerClients only
// exist for manager customers. All direct and indirect client customers are
// included, as well as the manager itself.
type CustomerClient struct {
	// The resource name of the customer client.
	// CustomerClient resource names have the form:
	// `customers/{customer_id}/customerClients/{client_customer_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The resource name of the client-customer which is linked to
	// the given customer. Read only.
	ClientCustomer *wrappers.StringValue `protobuf:"bytes,3,opt,name=client_customer,json=clientCustomer,proto3" json:"client_customer,omitempty"`
	// Specifies whether this is a
	// [hidden account](https://support.google.com/google-ads/answer/7519830).
	// Read only.
	Hidden *wrappers.BoolValue `protobuf:"bytes,4,opt,name=hidden,proto3" json:"hidden,omitempty"`
	// Distance between given customer and client. For self link, the level value
	// will be 0. Read only.
	Level *wrappers.Int64Value `protobuf:"bytes,5,opt,name=level,proto3" json:"level,omitempty"`
	// Common Locale Data Repository (CLDR) string representation of the
	// time zone of the client, e.g. America/Los_Angeles. Read only.
	TimeZone *wrappers.StringValue `protobuf:"bytes,6,opt,name=time_zone,json=timeZone,proto3" json:"time_zone,omitempty"`
	// Identifies if the client is a test account. Read only.
	TestAccount *wrappers.BoolValue `protobuf:"bytes,7,opt,name=test_account,json=testAccount,proto3" json:"test_account,omitempty"`
	// Identifies if the client is a manager. Read only.
	Manager *wrappers.BoolValue `protobuf:"bytes,8,opt,name=manager,proto3" json:"manager,omitempty"`
	// Descriptive name for the client. Read only.
	DescriptiveName *wrappers.StringValue `protobuf:"bytes,9,opt,name=descriptive_name,json=descriptiveName,proto3" json:"descriptive_name,omitempty"`
	// Currency code (e.g. 'USD', 'EUR') for the client. Read only.
	CurrencyCode *wrappers.StringValue `protobuf:"bytes,10,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// The ID of the client customer. Read only.
	Id                   *wrappers.Int64Value `protobuf:"bytes,11,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CustomerClient) Reset()         { *m = CustomerClient{} }
func (m *CustomerClient) String() string { return proto.CompactTextString(m) }
func (*CustomerClient) ProtoMessage()    {}
func (*CustomerClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_463452e1d1f19c18, []int{0}
}

func (m *CustomerClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerClient.Unmarshal(m, b)
}
func (m *CustomerClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerClient.Marshal(b, m, deterministic)
}
func (m *CustomerClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerClient.Merge(m, src)
}
func (m *CustomerClient) XXX_Size() int {
	return xxx_messageInfo_CustomerClient.Size(m)
}
func (m *CustomerClient) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerClient.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerClient proto.InternalMessageInfo

func (m *CustomerClient) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CustomerClient) GetClientCustomer() *wrappers.StringValue {
	if m != nil {
		return m.ClientCustomer
	}
	return nil
}

func (m *CustomerClient) GetHidden() *wrappers.BoolValue {
	if m != nil {
		return m.Hidden
	}
	return nil
}

func (m *CustomerClient) GetLevel() *wrappers.Int64Value {
	if m != nil {
		return m.Level
	}
	return nil
}

func (m *CustomerClient) GetTimeZone() *wrappers.StringValue {
	if m != nil {
		return m.TimeZone
	}
	return nil
}

func (m *CustomerClient) GetTestAccount() *wrappers.BoolValue {
	if m != nil {
		return m.TestAccount
	}
	return nil
}

func (m *CustomerClient) GetManager() *wrappers.BoolValue {
	if m != nil {
		return m.Manager
	}
	return nil
}

func (m *CustomerClient) GetDescriptiveName() *wrappers.StringValue {
	if m != nil {
		return m.DescriptiveName
	}
	return nil
}

func (m *CustomerClient) GetCurrencyCode() *wrappers.StringValue {
	if m != nil {
		return m.CurrencyCode
	}
	return nil
}

func (m *CustomerClient) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func init() {
	proto.RegisterType((*CustomerClient)(nil), "google.ads.googleads.v2.resources.CustomerClient")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/customer_client.proto", fileDescriptor_463452e1d1f19c18)
}

var fileDescriptor_463452e1d1f19c18 = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0x8b, 0xd4, 0x30,
	0x18, 0xc6, 0x99, 0xee, 0xdf, 0xc9, 0xcc, 0xee, 0x4a, 0xbc, 0x94, 0x71, 0x91, 0x5d, 0x65, 0x61,
	0x41, 0x48, 0xb1, 0x2e, 0x8a, 0x15, 0x0f, 0x9d, 0x41, 0x06, 0x3d, 0xc8, 0x32, 0xc2, 0x1c, 0x96,
	0x81, 0x92, 0x4d, 0x5e, 0x6b, 0xa0, 0x4d, 0x4a, 0x92, 0x8e, 0xe8, 0xc9, 0xcf, 0xe2, 0xd1, 0x8f,
	0x22, 0x7e, 0x12, 0x3f, 0x85, 0xb4, 0x69, 0x8a, 0xb2, 0xe0, 0xcc, 0xed, 0x65, 0xfa, 0xfb, 0x3d,
	0xef, 0xc3, 0x24, 0x41, 0x2f, 0x72, 0xa5, 0xf2, 0x02, 0x22, 0xca, 0x4d, 0xe4, 0xc6, 0x66, 0x5a,
	0xc7, 0x91, 0x06, 0xa3, 0x6a, 0xcd, 0xc0, 0x44, 0xac, 0x36, 0x56, 0x95, 0xa0, 0x33, 0x56, 0x08,
	0x90, 0x96, 0x54, 0x5a, 0x59, 0x85, 0xcf, 0x1d, 0x4d, 0x28, 0x37, 0xa4, 0x17, 0xc9, 0x3a, 0x26,
	0xbd, 0x38, 0x79, 0xd8, 0x65, 0xb7, 0xc2, 0x6d, 0xfd, 0x31, 0xfa, 0xac, 0x69, 0x55, 0x81, 0x36,
	0x2e, 0x62, 0x72, 0xea, 0x77, 0x57, 0x22, 0xa2, 0x52, 0x2a, 0x4b, 0xad, 0x50, 0xb2, 0xfb, 0xfa,
	0xe8, 0xd7, 0x2e, 0x3a, 0x9e, 0x75, 0xab, 0x67, 0xed, 0x66, 0xfc, 0x18, 0x1d, 0xf9, 0xf4, 0x4c,
	0xd2, 0x12, 0xc2, 0xc1, 0xd9, 0xe0, 0x72, 0xb8, 0x18, 0xfb, 0x1f, 0xdf, 0xd3, 0x12, 0xf0, 0x1b,
	0x74, 0xe2, 0x8a, 0x66, 0xbe, 0x78, 0xb8, 0x73, 0x36, 0xb8, 0x1c, 0xc5, 0xa7, 0x5d, 0x4f, 0xe2,
	0xfb, 0x90, 0x0f, 0x56, 0x0b, 0x99, 0x2f, 0x69, 0x51, 0xc3, 0xe2, 0xd8, 0x49, 0x7e, 0x23, 0x8e,
	0xd1, 0xfe, 0x27, 0xc1, 0x39, 0xc8, 0x70, 0xb7, 0xb5, 0x27, 0x77, 0xec, 0xa9, 0x52, 0x85, 0x73,
	0x3b, 0x12, 0x3f, 0x45, 0x7b, 0x05, 0xac, 0xa1, 0x08, 0xf7, 0x5a, 0xe5, 0xc1, 0x1d, 0xe5, 0xad,
	0xb4, 0xcf, 0xaf, 0x9c, 0xe3, 0x48, 0xfc, 0x12, 0x0d, 0xad, 0x28, 0x21, 0xfb, 0xaa, 0x24, 0x84,
	0xfb, 0x5b, 0xf4, 0x3c, 0x6c, 0xf0, 0x1b, 0x25, 0x01, 0xbf, 0x46, 0x63, 0x0b, 0xc6, 0x66, 0x94,
	0x31, 0x55, 0x4b, 0x1b, 0x1e, 0x6c, 0xec, 0x39, 0x6a, 0xf8, 0xd4, 0xe1, 0xf8, 0x0a, 0x1d, 0x94,
	0x54, 0xd2, 0x1c, 0x74, 0x78, 0xb8, 0xd1, 0xf4, 0x28, 0x9e, 0xa3, 0x7b, 0x1c, 0x0c, 0xd3, 0xa2,
	0xb2, 0x62, 0xdd, 0x9d, 0xc2, 0x70, 0x8b, 0xda, 0x27, 0x7f, 0x59, 0xed, 0x31, 0xa5, 0xe8, 0x88,
	0xd5, 0x5a, 0x83, 0x64, 0x5f, 0x32, 0xa6, 0x38, 0x84, 0x68, 0x8b, 0x94, 0xb1, 0x57, 0x66, 0x8a,
	0x03, 0x7e, 0x82, 0x02, 0xc1, 0xc3, 0xd1, 0xe6, 0xff, 0x3a, 0x10, 0x7c, 0xfa, 0x2d, 0x40, 0x17,
	0x4c, 0x95, 0x64, 0xe3, 0xb5, 0x9d, 0xde, 0xff, 0xf7, 0xd6, 0x5d, 0x37, 0x81, 0xd7, 0x83, 0x9b,
	0x77, 0x9d, 0x99, 0xab, 0x82, 0xca, 0x9c, 0x28, 0x9d, 0x47, 0x39, 0xc8, 0x76, 0x9d, 0x7f, 0x39,
	0x95, 0x30, 0xff, 0x79, 0x48, 0xaf, 0xfa, 0xe9, 0x7b, 0xb0, 0x33, 0x4f, 0xd3, 0x1f, 0xc1, 0xf9,
	0xdc, 0x45, 0xa6, 0xdc, 0x10, 0x37, 0x36, 0xd3, 0x32, 0x26, 0x0b, 0x4f, 0xfe, 0xf4, 0xcc, 0x2a,
	0xe5, 0x66, 0xd5, 0x33, 0xab, 0x65, 0xbc, 0xea, 0x99, 0xdf, 0xc1, 0x85, 0xfb, 0x90, 0x24, 0x29,
	0x37, 0x49, 0xd2, 0x53, 0x49, 0xb2, 0x8c, 0x93, 0xa4, 0xe7, 0x6e, 0xf7, 0xdb, 0xb2, 0xcf, 0xfe,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x41, 0x12, 0xdf, 0x6d, 0xf4, 0x03, 0x00, 0x00,
}

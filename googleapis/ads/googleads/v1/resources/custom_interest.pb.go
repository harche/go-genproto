// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/custom_interest.proto

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

// A custom interest. This is a list of users by interest.
type CustomInterest struct {
	// The resource name of the custom interest.
	// Custom interest resource names have the form:
	//
	// `customers/{customer_id}/customInterests/{custom_interest_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Id of the custom interest.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Status of this custom interest. Indicates whether the custom interest is
	// enabled or removed.
	Status enums.CustomInterestStatusEnum_CustomInterestStatus `protobuf:"varint,3,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.CustomInterestStatusEnum_CustomInterestStatus" json:"status,omitempty"`
	// Name of the custom interest. It should be unique across the same custom
	// affinity audience.
	// This field is required for create operations.
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Type of the custom interest, CUSTOM_AFFINITY or CUSTOM_INTENT.
	// By default the type is set to CUSTOM_AFFINITY.
	Type enums.CustomInterestTypeEnum_CustomInterestType `protobuf:"varint,5,opt,name=type,proto3,enum=google.ads.googleads.v1.enums.CustomInterestTypeEnum_CustomInterestType" json:"type,omitempty"`
	// Description of this custom interest audience.
	Description *wrappers.StringValue `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// List of custom interest members that this custom interest is composed of.
	// Members can be added during CustomInterest creation. If members are
	// presented in UPDATE operation, existing members will be overridden.
	Members              []*CustomInterestMember `protobuf:"bytes,7,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CustomInterest) Reset()         { *m = CustomInterest{} }
func (m *CustomInterest) String() string { return proto.CompactTextString(m) }
func (*CustomInterest) ProtoMessage()    {}
func (*CustomInterest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a831eeefd2f4541b, []int{0}
}

func (m *CustomInterest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomInterest.Unmarshal(m, b)
}
func (m *CustomInterest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomInterest.Marshal(b, m, deterministic)
}
func (m *CustomInterest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomInterest.Merge(m, src)
}
func (m *CustomInterest) XXX_Size() int {
	return xxx_messageInfo_CustomInterest.Size(m)
}
func (m *CustomInterest) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomInterest.DiscardUnknown(m)
}

var xxx_messageInfo_CustomInterest proto.InternalMessageInfo

func (m *CustomInterest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CustomInterest) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *CustomInterest) GetStatus() enums.CustomInterestStatusEnum_CustomInterestStatus {
	if m != nil {
		return m.Status
	}
	return enums.CustomInterestStatusEnum_UNSPECIFIED
}

func (m *CustomInterest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CustomInterest) GetType() enums.CustomInterestTypeEnum_CustomInterestType {
	if m != nil {
		return m.Type
	}
	return enums.CustomInterestTypeEnum_UNSPECIFIED
}

func (m *CustomInterest) GetDescription() *wrappers.StringValue {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *CustomInterest) GetMembers() []*CustomInterestMember {
	if m != nil {
		return m.Members
	}
	return nil
}

// A member of custom interest audience. A member can be a keyword or url.
// It is immutable, that is, it can only be created or removed but not changed.
type CustomInterestMember struct {
	// The type of custom interest member, KEYWORD or URL.
	MemberType enums.CustomInterestMemberTypeEnum_CustomInterestMemberType `protobuf:"varint,1,opt,name=member_type,json=memberType,proto3,enum=google.ads.googleads.v1.enums.CustomInterestMemberTypeEnum_CustomInterestMemberType" json:"member_type,omitempty"`
	// Keyword text when member_type is KEYWORD or URL string when
	// member_type is URL.
	Parameter            *wrappers.StringValue `protobuf:"bytes,2,opt,name=parameter,proto3" json:"parameter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CustomInterestMember) Reset()         { *m = CustomInterestMember{} }
func (m *CustomInterestMember) String() string { return proto.CompactTextString(m) }
func (*CustomInterestMember) ProtoMessage()    {}
func (*CustomInterestMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_a831eeefd2f4541b, []int{1}
}

func (m *CustomInterestMember) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomInterestMember.Unmarshal(m, b)
}
func (m *CustomInterestMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomInterestMember.Marshal(b, m, deterministic)
}
func (m *CustomInterestMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomInterestMember.Merge(m, src)
}
func (m *CustomInterestMember) XXX_Size() int {
	return xxx_messageInfo_CustomInterestMember.Size(m)
}
func (m *CustomInterestMember) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomInterestMember.DiscardUnknown(m)
}

var xxx_messageInfo_CustomInterestMember proto.InternalMessageInfo

func (m *CustomInterestMember) GetMemberType() enums.CustomInterestMemberTypeEnum_CustomInterestMemberType {
	if m != nil {
		return m.MemberType
	}
	return enums.CustomInterestMemberTypeEnum_UNSPECIFIED
}

func (m *CustomInterestMember) GetParameter() *wrappers.StringValue {
	if m != nil {
		return m.Parameter
	}
	return nil
}

func init() {
	proto.RegisterType((*CustomInterest)(nil), "google.ads.googleads.v1.resources.CustomInterest")
	proto.RegisterType((*CustomInterestMember)(nil), "google.ads.googleads.v1.resources.CustomInterestMember")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/custom_interest.proto", fileDescriptor_a831eeefd2f4541b)
}

var fileDescriptor_a831eeefd2f4541b = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x55, 0xd2, 0xd2, 0x69, 0x2e, 0xf4, 0xc1, 0xf0, 0x10, 0x8d, 0x09, 0x75, 0x43, 0x93, 0x2a,
	0x21, 0x39, 0xb4, 0x20, 0x86, 0x8c, 0x04, 0xca, 0x10, 0x1a, 0x43, 0x80, 0x46, 0x36, 0xf5, 0x01,
	0x45, 0xaa, 0xdc, 0xc6, 0x44, 0x91, 0x6a, 0x3b, 0xb2, 0x9d, 0xa1, 0xbd, 0xf1, 0x2d, 0x3c, 0xf2,
	0x27, 0xf0, 0x29, 0xf0, 0x13, 0x28, 0x76, 0x9c, 0x75, 0xac, 0xdd, 0xe8, 0xdb, 0xad, 0xef, 0x39,
	0xe7, 0x1e, 0x1f, 0xdf, 0x06, 0xec, 0x67, 0x42, 0x64, 0x73, 0x1a, 0x92, 0x54, 0x85, 0xb6, 0xac,
	0xaa, 0xb3, 0x61, 0x28, 0xa9, 0x12, 0xa5, 0x9c, 0x51, 0x15, 0xce, 0x4a, 0xa5, 0x05, 0x9b, 0xe4,
	0x5c, 0x53, 0x49, 0x95, 0x46, 0x85, 0x14, 0x5a, 0xc0, 0x1d, 0x8b, 0x46, 0x24, 0x55, 0xa8, 0x21,
	0xa2, 0xb3, 0x21, 0x6a, 0x88, 0x5b, 0xaf, 0x56, 0x69, 0x53, 0x5e, 0xb2, 0x2b, 0xba, 0x13, 0x46,
	0xd9, 0x94, 0xca, 0x89, 0x3e, 0x2f, 0xa8, 0x9d, 0xb1, 0x85, 0xd7, 0x13, 0x50, 0x9a, 0xe8, 0x52,
	0xd5, 0xdc, 0xe7, 0xeb, 0x71, 0x17, 0xa6, 0x6e, 0x3b, 0x66, 0x91, 0x87, 0x84, 0x73, 0xa1, 0x89,
	0xce, 0x05, 0x77, 0xba, 0x0f, 0xea, 0xae, 0xf9, 0x35, 0x2d, 0xbf, 0x84, 0x5f, 0x25, 0x29, 0x0a,
	0x2a, 0xeb, 0xfe, 0xee, 0x9f, 0x16, 0xe8, 0xbd, 0x36, 0xe2, 0x47, 0xb5, 0x36, 0x7c, 0x08, 0xee,
	0xb8, 0x50, 0x26, 0x9c, 0x30, 0x1a, 0x78, 0x7d, 0x6f, 0xb0, 0x19, 0xdf, 0x76, 0x87, 0x1f, 0x09,
	0xa3, 0xf0, 0x11, 0xf0, 0xf3, 0x34, 0xf0, 0xfb, 0xde, 0xa0, 0x3b, 0xba, 0x5f, 0x27, 0x8a, 0xdc,
	0x10, 0x74, 0xc4, 0xf5, 0xb3, 0xa7, 0x63, 0x32, 0x2f, 0x69, 0xec, 0xe7, 0x29, 0x4c, 0x41, 0xc7,
	0x5e, 0x36, 0x68, 0xf5, 0xbd, 0x41, 0x6f, 0xf4, 0x1e, 0xad, 0x7a, 0x0d, 0x73, 0x5b, 0x74, 0xd9,
	0xd0, 0x89, 0xa1, 0xbe, 0xe1, 0x25, 0x5b, 0xda, 0x88, 0x6b, 0x6d, 0xf8, 0x18, 0xb4, 0x8d, 0xdd,
	0xb6, 0x31, 0xb5, 0x7d, 0xc5, 0xd4, 0x89, 0x96, 0x39, 0xcf, 0xac, 0x2b, 0x83, 0x84, 0x09, 0x68,
	0x57, 0x41, 0x06, 0xb7, 0x8c, 0xab, 0xb7, 0x6b, 0xb9, 0x3a, 0x3d, 0x2f, 0xe8, 0x12, 0x4f, 0xd5,
	0x71, 0x6c, 0x54, 0xe1, 0x4b, 0xd0, 0x4d, 0xa9, 0x9a, 0xc9, 0xbc, 0xa8, 0x1e, 0x24, 0xe8, 0xfc,
	0x87, 0xad, 0x45, 0x02, 0xfc, 0x04, 0x36, 0xec, 0x8e, 0xa9, 0x60, 0xa3, 0xdf, 0x1a, 0x74, 0x47,
	0xfb, 0xe8, 0xc6, 0x25, 0xfe, 0xc7, 0xcd, 0x07, 0xc3, 0x8f, 0x9d, 0xce, 0xee, 0x4f, 0x0f, 0xdc,
	0x5b, 0x86, 0x80, 0x25, 0xe8, 0x2e, 0xec, 0xb3, 0x79, 0xf1, 0xde, 0xe8, 0x74, 0xad, 0x40, 0xac,
	0xd2, 0x8a, 0x58, 0x2e, 0x9a, 0x31, 0x60, 0x4d, 0x0d, 0x31, 0xd8, 0x2c, 0x88, 0x24, 0x8c, 0x6a,
	0x2a, 0xeb, 0x65, 0xba, 0x3e, 0xa0, 0x0b, 0xf8, 0xc1, 0x37, 0x1f, 0xec, 0xcd, 0x04, 0xbb, 0x39,
	0x93, 0x83, 0xbb, 0x97, 0xbd, 0x1c, 0x57, 0xc2, 0xc7, 0xde, 0xe7, 0x77, 0x35, 0x33, 0x13, 0x73,
	0xc2, 0x33, 0x24, 0x64, 0x16, 0x66, 0x94, 0x9b, 0xb1, 0xee, 0x2f, 0x58, 0xe4, 0xea, 0x9a, 0x4f,
	0xcd, 0x8b, 0xa6, 0xfa, 0xee, 0xb7, 0x0e, 0xa3, 0xe8, 0x87, 0xbf, 0x73, 0x68, 0x25, 0xa3, 0x54,
	0x21, 0x5b, 0x56, 0xd5, 0x78, 0x88, 0x62, 0x87, 0xfc, 0xe5, 0x30, 0x49, 0x94, 0xaa, 0xa4, 0xc1,
	0x24, 0xe3, 0x61, 0xd2, 0x60, 0x7e, 0xfb, 0x7b, 0xb6, 0x81, 0x71, 0x94, 0x2a, 0x8c, 0x1b, 0x14,
	0xc6, 0xe3, 0x21, 0xc6, 0x0d, 0x6e, 0xda, 0x31, 0x66, 0x9f, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0x96, 0x2a, 0x9e, 0xe6, 0x16, 0x05, 0x00, 0x00,
}

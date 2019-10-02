// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/topic_view.proto

package resources

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// A topic view.
type TopicView struct {
	// The resource name of the topic view.
	// Topic view resource names have the form:
	//
	// `customers/{customer_id}/topicViews/{ad_group_id}~{criterion_id}`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicView) Reset()         { *m = TopicView{} }
func (m *TopicView) String() string { return proto.CompactTextString(m) }
func (*TopicView) ProtoMessage()    {}
func (*TopicView) Descriptor() ([]byte, []int) {
	return fileDescriptor_19e9fa6c2dd72e26, []int{0}
}

func (m *TopicView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicView.Unmarshal(m, b)
}
func (m *TopicView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicView.Marshal(b, m, deterministic)
}
func (m *TopicView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicView.Merge(m, src)
}
func (m *TopicView) XXX_Size() int {
	return xxx_messageInfo_TopicView.Size(m)
}
func (m *TopicView) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicView.DiscardUnknown(m)
}

var xxx_messageInfo_TopicView proto.InternalMessageInfo

func (m *TopicView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*TopicView)(nil), "google.ads.googleads.v2.resources.TopicView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/topic_view.proto", fileDescriptor_19e9fa6c2dd72e26)
}

var fileDescriptor_19e9fa6c2dd72e26 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0x4b, 0xc4, 0x30,
	0x14, 0xc7, 0x69, 0x05, 0xe1, 0x8a, 0x3a, 0xdc, 0x24, 0xe2, 0xe0, 0x29, 0x07, 0x4e, 0x2f, 0x12,
	0xb7, 0x38, 0xe5, 0x96, 0x03, 0x07, 0x39, 0x0e, 0xe9, 0x20, 0x85, 0x23, 0x36, 0x21, 0x04, 0xae,
	0x79, 0xa5, 0xaf, 0xf6, 0xbe, 0x8f, 0xa3, 0x1f, 0xc5, 0x8f, 0xe2, 0x57, 0x70, 0x91, 0x36, 0x26,
	0xa3, 0x6e, 0x7f, 0x92, 0xdf, 0xff, 0xf7, 0x5e, 0x52, 0x70, 0x8b, 0x68, 0xf7, 0x86, 0x29, 0x4d,
	0x2c, 0xc4, 0x31, 0x0d, 0x9c, 0x75, 0x86, 0xf0, 0xad, 0xab, 0x0d, 0xb1, 0x1e, 0x5b, 0x57, 0xef,
	0x06, 0x67, 0x0e, 0xd0, 0x76, 0xd8, 0xe3, 0x7c, 0x11, 0x40, 0x50, 0x9a, 0x20, 0x75, 0x60, 0xe0,
	0x90, 0x3a, 0x17, 0x97, 0x51, 0xdb, 0x3a, 0xa6, 0xbc, 0xc7, 0x5e, 0xf5, 0x0e, 0x3d, 0x05, 0xc1,
	0xf5, 0x5d, 0x31, 0x7b, 0x1e, 0xa5, 0xa5, 0x33, 0x87, 0xf9, 0x4d, 0x71, 0x1a, 0x7b, 0x3b, 0xaf,
	0x1a, 0x73, 0x9e, 0x5d, 0x65, 0xb7, 0xb3, 0xed, 0x49, 0x3c, 0x7c, 0x52, 0x8d, 0x59, 0x7d, 0x67,
	0xc5, 0xb2, 0xc6, 0x06, 0xfe, 0x9d, 0xbc, 0x3a, 0x4b, 0xe6, 0xcd, 0x38, 0x6b, 0x93, 0xbd, 0x3c,
	0xfe, 0x96, 0x2c, 0xee, 0x95, 0xb7, 0x80, 0x9d, 0x65, 0xd6, 0xf8, 0x69, 0x93, 0xf8, 0xe4, 0xd6,
	0xd1, 0x1f, 0x3f, 0xf0, 0x90, 0xd2, 0x7b, 0x7e, 0xb4, 0x96, 0xf2, 0x23, 0x5f, 0xac, 0x83, 0x52,
	0x6a, 0x82, 0x10, 0xc7, 0x54, 0x72, 0xd8, 0x46, 0xf2, 0x33, 0x32, 0x95, 0xd4, 0x54, 0x25, 0xa6,
	0x2a, 0x79, 0x95, 0x98, 0xaf, 0x7c, 0x19, 0x2e, 0x84, 0x90, 0x9a, 0x84, 0x48, 0x94, 0x10, 0x25,
	0x17, 0x22, 0x71, 0xaf, 0xc7, 0xd3, 0xb2, 0xf7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x84, 0x81,
	0xa1, 0xfa, 0xad, 0x01, 0x00, 0x00,
}

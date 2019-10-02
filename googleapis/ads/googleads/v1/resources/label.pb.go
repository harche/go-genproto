// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/label.proto

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

// A label.
type Label struct {
	// Name of the resource.
	// Label resource names have the form:
	// `customers/{customer_id}/labels/{label_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Id of the label. Read only.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the label.
	//
	// This field is required and should not be empty when creating a new label.
	//
	// The length of this string should be between 1 and 80, inclusive.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Status of the label. Read only.
	Status enums.LabelStatusEnum_LabelStatus `protobuf:"varint,4,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.LabelStatusEnum_LabelStatus" json:"status,omitempty"`
	// A type of label displaying text on a colored background.
	TextLabel            *common.TextLabel `protobuf:"bytes,5,opt,name=text_label,json=textLabel,proto3" json:"text_label,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Label) Reset()         { *m = Label{} }
func (m *Label) String() string { return proto.CompactTextString(m) }
func (*Label) ProtoMessage()    {}
func (*Label) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4f50cd20a4c405a, []int{0}
}

func (m *Label) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Label.Unmarshal(m, b)
}
func (m *Label) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Label.Marshal(b, m, deterministic)
}
func (m *Label) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Label.Merge(m, src)
}
func (m *Label) XXX_Size() int {
	return xxx_messageInfo_Label.Size(m)
}
func (m *Label) XXX_DiscardUnknown() {
	xxx_messageInfo_Label.DiscardUnknown(m)
}

var xxx_messageInfo_Label proto.InternalMessageInfo

func (m *Label) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Label) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Label) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Label) GetStatus() enums.LabelStatusEnum_LabelStatus {
	if m != nil {
		return m.Status
	}
	return enums.LabelStatusEnum_UNSPECIFIED
}

func (m *Label) GetTextLabel() *common.TextLabel {
	if m != nil {
		return m.TextLabel
	}
	return nil
}

func init() {
	proto.RegisterType((*Label)(nil), "google.ads.googleads.v1.resources.Label")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/label.proto", fileDescriptor_b4f50cd20a4c405a)
}

var fileDescriptor_b4f50cd20a4c405a = []byte{
	// 411 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcb, 0x8a, 0xdb, 0x30,
	0x18, 0x85, 0xb1, 0x73, 0x81, 0xa8, 0x97, 0x85, 0x57, 0x26, 0x0d, 0x25, 0x69, 0x09, 0xa4, 0x94,
	0x4a, 0x71, 0x5a, 0xba, 0x50, 0x57, 0x0e, 0x94, 0xb4, 0xa5, 0x94, 0xe0, 0x14, 0x2f, 0x8a, 0x21,
	0x28, 0xb1, 0x6a, 0x0c, 0xb6, 0x64, 0x2c, 0x39, 0xcd, 0xf3, 0x74, 0xd7, 0x3e, 0xca, 0x3c, 0xca,
	0xbc, 0xc0, 0x6c, 0x07, 0x4b, 0x96, 0x27, 0x30, 0x78, 0x66, 0x77, 0x94, 0xff, 0x3b, 0x27, 0xff,
	0xc5, 0xe0, 0x5d, 0xc2, 0x79, 0x92, 0x51, 0x44, 0x62, 0x81, 0xb4, 0xac, 0xd5, 0xc9, 0x43, 0x25,
	0x15, 0xbc, 0x2a, 0x8f, 0x54, 0xa0, 0x8c, 0x1c, 0x68, 0x06, 0x8b, 0x92, 0x4b, 0xee, 0xcc, 0x34,
	0x03, 0x49, 0x2c, 0x60, 0x8b, 0xc3, 0x93, 0x07, 0x5b, 0x7c, 0x8c, 0xba, 0x12, 0x8f, 0x3c, 0xcf,
	0x39, 0x43, 0x92, 0x9e, 0xe5, 0xfe, 0x22, 0x73, 0xbc, 0xec, 0x32, 0x50, 0x56, 0xe5, 0xcd, 0xdf,
	0xef, 0x85, 0x24, 0xb2, 0x12, 0x8d, 0x63, 0x62, 0x1c, 0x45, 0x8a, 0x08, 0x63, 0x5c, 0x12, 0x99,
	0x72, 0x66, 0xaa, 0x2f, 0x9b, 0xaa, 0x7a, 0x1d, 0xaa, 0xdf, 0xe8, 0x4f, 0x49, 0x8a, 0x82, 0x96,
	0x4d, 0xfd, 0xd5, 0x3f, 0x1b, 0x0c, 0xbe, 0xd7, 0xa1, 0xce, 0x6b, 0xf0, 0xcc, 0xf4, 0xbd, 0x67,
	0x24, 0xa7, 0xae, 0x35, 0xb5, 0x16, 0xa3, 0xe0, 0xa9, 0xf9, 0xf1, 0x07, 0xc9, 0xa9, 0xf3, 0x16,
	0xd8, 0x69, 0xec, 0xda, 0x53, 0x6b, 0xf1, 0x64, 0xf5, 0xa2, 0x19, 0x1a, 0x9a, 0x6c, 0xf8, 0x95,
	0xc9, 0x8f, 0x1f, 0x42, 0x92, 0x55, 0x34, 0xb0, 0xd3, 0xd8, 0x59, 0x82, 0xbe, 0x0a, 0xea, 0x29,
	0x7c, 0x72, 0x0f, 0xdf, 0xc9, 0x32, 0x65, 0x89, 0xe6, 0x15, 0xe9, 0x04, 0x60, 0xa8, 0x67, 0x73,
	0xfb, 0x53, 0x6b, 0xf1, 0x7c, 0x85, 0x61, 0xd7, 0x8a, 0xd5, 0x3a, 0xa0, 0xea, 0x7c, 0xa7, 0x1c,
	0x9f, 0x59, 0x95, 0x5f, 0xbe, 0x83, 0x26, 0xc9, 0xf9, 0x02, 0xc0, 0xdd, 0x96, 0xdd, 0x81, 0xea,
	0xe5, 0x4d, 0x67, 0xae, 0xbe, 0x0b, 0xfc, 0x49, 0xcf, 0x52, 0x85, 0x05, 0x23, 0x69, 0xe4, 0xfa,
	0xc6, 0x02, 0xf3, 0x23, 0xcf, 0xe1, 0xa3, 0x67, 0x5f, 0x03, 0x65, 0xd8, 0xd6, 0x83, 0x6e, 0xad,
	0x5f, 0xdf, 0x1a, 0x43, 0xc2, 0x33, 0xc2, 0x12, 0xc8, 0xcb, 0x04, 0x25, 0x94, 0xa9, 0x35, 0x98,
	0x1b, 0x17, 0xa9, 0x78, 0xe0, 0xab, 0xfb, 0xd4, 0xaa, 0xbf, 0x76, 0x6f, 0xe3, 0xfb, 0xff, 0xed,
	0xd9, 0x46, 0x47, 0xfa, 0xb1, 0x80, 0x5a, 0xd6, 0x2a, 0xf4, 0x60, 0x60, 0xc8, 0x2b, 0xc3, 0x44,
	0x7e, 0x2c, 0xa2, 0x96, 0x89, 0x42, 0x2f, 0x6a, 0x99, 0x6b, 0x7b, 0xae, 0x0b, 0x18, 0xfb, 0xb1,
	0xc0, 0xb8, 0xa5, 0x30, 0x0e, 0x3d, 0x8c, 0x5b, 0xee, 0x30, 0x54, 0xcd, 0xbe, 0xbf, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0x86, 0xb3, 0x62, 0x11, 0x21, 0x03, 0x00, 0x00,
}

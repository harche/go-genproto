// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1/data_items.proto

package automl

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/struct"
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

// A representation of a text snippet.
type TextSnippet struct {
	// Required. The content of the text snippet as a string. Up to 250000
	// characters long.
	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// Optional. The format of [content][google.cloud.automl.v1.TextSnippet.content]. Currently the only two allowed
	// values are "text/html" and "text/plain". If left blank, the format is
	// automatically determined from the type of the uploaded [content][google.cloud.automl.v1.TextSnippet.content].
	MimeType string `protobuf:"bytes,2,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// Output only. HTTP URI where you can download the content.
	ContentUri           string   `protobuf:"bytes,4,opt,name=content_uri,json=contentUri,proto3" json:"content_uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextSnippet) Reset()         { *m = TextSnippet{} }
func (m *TextSnippet) String() string { return proto.CompactTextString(m) }
func (*TextSnippet) ProtoMessage()    {}
func (*TextSnippet) Descriptor() ([]byte, []int) {
	return fileDescriptor_9aa6620a22bddfe1, []int{0}
}

func (m *TextSnippet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextSnippet.Unmarshal(m, b)
}
func (m *TextSnippet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextSnippet.Marshal(b, m, deterministic)
}
func (m *TextSnippet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextSnippet.Merge(m, src)
}
func (m *TextSnippet) XXX_Size() int {
	return xxx_messageInfo_TextSnippet.Size(m)
}
func (m *TextSnippet) XXX_DiscardUnknown() {
	xxx_messageInfo_TextSnippet.DiscardUnknown(m)
}

var xxx_messageInfo_TextSnippet proto.InternalMessageInfo

func (m *TextSnippet) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *TextSnippet) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

func (m *TextSnippet) GetContentUri() string {
	if m != nil {
		return m.ContentUri
	}
	return ""
}

// Example data used for training or prediction.
type ExamplePayload struct {
	// Required. Input only. The example data.
	//
	// Types that are valid to be assigned to Payload:
	//	*ExamplePayload_TextSnippet
	Payload              isExamplePayload_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ExamplePayload) Reset()         { *m = ExamplePayload{} }
func (m *ExamplePayload) String() string { return proto.CompactTextString(m) }
func (*ExamplePayload) ProtoMessage()    {}
func (*ExamplePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_9aa6620a22bddfe1, []int{1}
}

func (m *ExamplePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExamplePayload.Unmarshal(m, b)
}
func (m *ExamplePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExamplePayload.Marshal(b, m, deterministic)
}
func (m *ExamplePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExamplePayload.Merge(m, src)
}
func (m *ExamplePayload) XXX_Size() int {
	return xxx_messageInfo_ExamplePayload.Size(m)
}
func (m *ExamplePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_ExamplePayload.DiscardUnknown(m)
}

var xxx_messageInfo_ExamplePayload proto.InternalMessageInfo

type isExamplePayload_Payload interface {
	isExamplePayload_Payload()
}

type ExamplePayload_TextSnippet struct {
	TextSnippet *TextSnippet `protobuf:"bytes,2,opt,name=text_snippet,json=textSnippet,proto3,oneof"`
}

func (*ExamplePayload_TextSnippet) isExamplePayload_Payload() {}

func (m *ExamplePayload) GetPayload() isExamplePayload_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ExamplePayload) GetTextSnippet() *TextSnippet {
	if x, ok := m.GetPayload().(*ExamplePayload_TextSnippet); ok {
		return x.TextSnippet
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExamplePayload) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExamplePayload_TextSnippet)(nil),
	}
}

func init() {
	proto.RegisterType((*TextSnippet)(nil), "google.cloud.automl.v1.TextSnippet")
	proto.RegisterType((*ExamplePayload)(nil), "google.cloud.automl.v1.ExamplePayload")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1/data_items.proto", fileDescriptor_9aa6620a22bddfe1)
}

var fileDescriptor_9aa6620a22bddfe1 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4e, 0xe3, 0x30,
	0x10, 0x86, 0x37, 0xd5, 0x6a, 0xbb, 0x75, 0x56, 0x7b, 0xc8, 0xa1, 0x4a, 0xbb, 0x2b, 0x8a, 0xca,
	0x01, 0x4e, 0xb6, 0x02, 0xb7, 0xc0, 0x85, 0x22, 0x44, 0x0f, 0x20, 0x55, 0xa5, 0xf4, 0x80, 0x2a,
	0x45, 0x6e, 0x62, 0x22, 0x4b, 0x89, 0xc7, 0x4a, 0x26, 0x55, 0xf3, 0x4a, 0x7d, 0x14, 0x1e, 0x85,
	0xa7, 0x40, 0xd8, 0xa9, 0x40, 0x11, 0xbd, 0x79, 0xe6, 0xfb, 0xff, 0x7f, 0x46, 0x63, 0x72, 0x9a,
	0x02, 0xa4, 0x99, 0x60, 0x71, 0x06, 0x55, 0xc2, 0x78, 0x85, 0x90, 0x67, 0x6c, 0x13, 0xb0, 0x84,
	0x23, 0x8f, 0x24, 0x8a, 0xbc, 0xa4, 0xba, 0x00, 0x04, 0xaf, 0x6f, 0x85, 0xd4, 0x08, 0xa9, 0x15,
	0xd2, 0x4d, 0x30, 0x1c, 0x1d, 0x08, 0x90, 0x60, 0x8d, 0xc3, 0x41, 0x23, 0x30, 0xd5, 0xba, 0x7a,
	0x61, 0x5c, 0xd5, 0x0d, 0x3a, 0x6a, 0xa3, 0xa4, 0x2a, 0x38, 0x4a, 0x50, 0x0d, 0xff, 0xdf, 0xe6,
	0x25, 0x16, 0x55, 0x8c, 0x2d, 0xca, 0xb5, 0x64, 0x5c, 0x29, 0x40, 0x63, 0x6d, 0xf6, 0x1d, 0x0b,
	0xe2, 0x2e, 0xc4, 0x16, 0x1f, 0x95, 0xd4, 0x5a, 0xa0, 0xe7, 0x93, 0x6e, 0x0c, 0x0a, 0x85, 0x42,
	0xdf, 0x39, 0x76, 0xce, 0x7a, 0xf3, 0x7d, 0xe9, 0xfd, 0x23, 0xbd, 0x5c, 0xe6, 0x22, 0xc2, 0x5a,
	0x0b, 0xbf, 0x63, 0xd8, 0xef, 0x8f, 0xc6, 0xa2, 0xd6, 0xc2, 0x1b, 0x11, 0xb7, 0xd1, 0x45, 0x55,
	0x21, 0xfd, 0x9f, 0x06, 0x93, 0xa6, 0xf5, 0x54, 0xc8, 0xb1, 0x20, 0x7f, 0x6f, 0xb7, 0x3c, 0xd7,
	0x99, 0x98, 0xf1, 0x3a, 0x03, 0x9e, 0x78, 0x53, 0xf2, 0x07, 0xc5, 0x16, 0xa3, 0xd2, 0x4e, 0x36,
	0x91, 0xee, 0xf9, 0x09, 0xfd, 0xfe, 0x7e, 0xf4, 0xcb, 0x92, 0xd3, 0x1f, 0x73, 0x17, 0x3f, 0xcb,
	0x49, 0x8f, 0x74, 0xb5, 0x0d, 0x9d, 0xec, 0x1c, 0x32, 0x8c, 0x21, 0x3f, 0x10, 0x32, 0x73, 0x9e,
	0xaf, 0x1a, 0x92, 0x42, 0xc6, 0x55, 0x4a, 0xa1, 0x48, 0x59, 0x2a, 0x94, 0x39, 0x05, 0xb3, 0x88,
	0x6b, 0x59, 0xb6, 0x7f, 0xe9, 0xd2, 0xbe, 0x76, 0x9d, 0xfe, 0x9d, 0xb5, 0xdf, 0x98, 0xe0, 0xeb,
	0x0a, 0xe1, 0xe1, 0x9e, 0x2e, 0x83, 0xd7, 0x3d, 0x58, 0x19, 0xb0, 0xb2, 0x60, 0xb5, 0x0c, 0xde,
	0x3a, 0x03, 0x0b, 0xc2, 0xd0, 0x90, 0x30, 0xb4, 0x28, 0x0c, 0x97, 0xc1, 0xfa, 0x97, 0x19, 0x7b,
	0xf1, 0x1e, 0x00, 0x00, 0xff, 0xff, 0xa7, 0xf3, 0x63, 0xbe, 0x5c, 0x02, 0x00, 0x00,
}

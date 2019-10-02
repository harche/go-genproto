// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/text_segment.proto

package automl

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

// A contiguous part of a text (string), assuming it has an UTF-8 NFC encoding.
type TextSegment struct {
	// Output only. The content of the TextSegment.
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	// Required. Zero-based character index of the first character of the text
	// segment (counting characters from the beginning of the text).
	StartOffset int64 `protobuf:"varint,1,opt,name=start_offset,json=startOffset,proto3" json:"start_offset,omitempty"`
	// Required. Zero-based character index of the first character past the end of
	// the text segment (counting character from the beginning of the text).
	// The character at the end_offset is NOT included in the text segment.
	EndOffset            int64    `protobuf:"varint,2,opt,name=end_offset,json=endOffset,proto3" json:"end_offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextSegment) Reset()         { *m = TextSegment{} }
func (m *TextSegment) String() string { return proto.CompactTextString(m) }
func (*TextSegment) ProtoMessage()    {}
func (*TextSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbd0421c131b44e6, []int{0}
}

func (m *TextSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextSegment.Unmarshal(m, b)
}
func (m *TextSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextSegment.Marshal(b, m, deterministic)
}
func (m *TextSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextSegment.Merge(m, src)
}
func (m *TextSegment) XXX_Size() int {
	return xxx_messageInfo_TextSegment.Size(m)
}
func (m *TextSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_TextSegment.DiscardUnknown(m)
}

var xxx_messageInfo_TextSegment proto.InternalMessageInfo

func (m *TextSegment) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *TextSegment) GetStartOffset() int64 {
	if m != nil {
		return m.StartOffset
	}
	return 0
}

func (m *TextSegment) GetEndOffset() int64 {
	if m != nil {
		return m.EndOffset
	}
	return 0
}

func init() {
	proto.RegisterType((*TextSegment)(nil), "google.cloud.automl.v1beta1.TextSegment")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/text_segment.proto", fileDescriptor_bbd0421c131b44e6)
}

var fileDescriptor_bbd0421c131b44e6 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x49, 0x0a, 0x4a, 0xb7, 0x1e, 0x24, 0xa7, 0x60, 0xfd, 0x53, 0x3d, 0xf5, 0xb4, 0x4b,
	0xf1, 0xb6, 0x9e, 0x5a, 0x0f, 0x5e, 0x14, 0x4b, 0x15, 0x0f, 0x12, 0x28, 0xdb, 0x64, 0xba, 0x04,
	0x92, 0x99, 0x90, 0x9d, 0x48, 0xdf, 0xcc, 0x77, 0xf0, 0x51, 0x7c, 0x0a, 0xe9, 0x6e, 0x0a, 0x1e,
	0xc4, 0xe3, 0xcc, 0xef, 0x37, 0xdf, 0x7e, 0xac, 0x90, 0x96, 0xc8, 0x56, 0xa0, 0xf2, 0x8a, 0xba,
	0x42, 0x99, 0x8e, 0xa9, 0xae, 0xd4, 0xc7, 0x6c, 0x03, 0x6c, 0x66, 0x8a, 0x61, 0xc7, 0x6b, 0x07,
	0xb6, 0x06, 0x64, 0xd9, 0xb4, 0xc4, 0x94, 0x8c, 0x83, 0x2f, 0xbd, 0x2f, 0x83, 0x2f, 0x7b, 0xff,
	0xec, 0xbc, 0x0f, 0x33, 0x4d, 0xa9, 0x0c, 0x22, 0xb1, 0xe1, 0x92, 0xd0, 0x85, 0xd3, 0x9b, 0x52,
	0x8c, 0x5e, 0x61, 0xc7, 0x2f, 0x21, 0x2f, 0x49, 0xc5, 0x71, 0x4e, 0xc8, 0x80, 0x9c, 0x0e, 0x26,
	0xd1, 0x74, 0xb8, 0x3a, 0x8c, 0xc9, 0xb5, 0x38, 0x71, 0x6c, 0x5a, 0x5e, 0xd3, 0x76, 0xeb, 0x80,
	0xd3, 0x68, 0x12, 0x4d, 0x07, 0xab, 0x91, 0xdf, 0x3d, 0xfb, 0x55, 0x72, 0x21, 0x04, 0x60, 0x71,
	0x10, 0x62, 0x2f, 0x0c, 0x01, 0x8b, 0x80, 0x17, 0x9f, 0x91, 0xb8, 0xca, 0xa9, 0x96, 0xff, 0x94,
	0x5d, 0x9c, 0xfe, 0x2a, 0xb3, 0xdc, 0x17, 0x5c, 0x46, 0xef, 0xf3, 0xfe, 0xc0, 0x52, 0x65, 0xd0,
	0x4a, 0x6a, 0xad, 0xb2, 0x80, 0xbe, 0xbe, 0x0a, 0xc8, 0x34, 0xa5, 0xfb, 0xf3, 0xb3, 0xee, 0xc2,
	0xf8, 0x15, 0x8f, 0x1f, 0xbc, 0x98, 0xdd, 0xef, 0xa5, 0x6c, 0xde, 0x31, 0x3d, 0x55, 0xd9, 0x5b,
	0x90, 0xbe, 0xe3, 0xcb, 0x40, 0xb5, 0xf6, 0x58, 0x6b, 0xcf, 0x1f, 0xb5, 0xee, 0x85, 0xcd, 0x91,
	0x7f, 0xec, 0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x97, 0x0e, 0x20, 0x23, 0x98, 0x01, 0x00, 0x00,
}

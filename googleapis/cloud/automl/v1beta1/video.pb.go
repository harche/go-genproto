// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/video.proto

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

// Dataset metadata specific to video classification.
// All Video Classification datasets are treated as multi label.
type VideoClassificationDatasetMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VideoClassificationDatasetMetadata) Reset()         { *m = VideoClassificationDatasetMetadata{} }
func (m *VideoClassificationDatasetMetadata) String() string { return proto.CompactTextString(m) }
func (*VideoClassificationDatasetMetadata) ProtoMessage()    {}
func (*VideoClassificationDatasetMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec8c6b74d94d2916, []int{0}
}

func (m *VideoClassificationDatasetMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoClassificationDatasetMetadata.Unmarshal(m, b)
}
func (m *VideoClassificationDatasetMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoClassificationDatasetMetadata.Marshal(b, m, deterministic)
}
func (m *VideoClassificationDatasetMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoClassificationDatasetMetadata.Merge(m, src)
}
func (m *VideoClassificationDatasetMetadata) XXX_Size() int {
	return xxx_messageInfo_VideoClassificationDatasetMetadata.Size(m)
}
func (m *VideoClassificationDatasetMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoClassificationDatasetMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_VideoClassificationDatasetMetadata proto.InternalMessageInfo

// Dataset metadata specific to video object tracking.
type VideoObjectTrackingDatasetMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VideoObjectTrackingDatasetMetadata) Reset()         { *m = VideoObjectTrackingDatasetMetadata{} }
func (m *VideoObjectTrackingDatasetMetadata) String() string { return proto.CompactTextString(m) }
func (*VideoObjectTrackingDatasetMetadata) ProtoMessage()    {}
func (*VideoObjectTrackingDatasetMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec8c6b74d94d2916, []int{1}
}

func (m *VideoObjectTrackingDatasetMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoObjectTrackingDatasetMetadata.Unmarshal(m, b)
}
func (m *VideoObjectTrackingDatasetMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoObjectTrackingDatasetMetadata.Marshal(b, m, deterministic)
}
func (m *VideoObjectTrackingDatasetMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoObjectTrackingDatasetMetadata.Merge(m, src)
}
func (m *VideoObjectTrackingDatasetMetadata) XXX_Size() int {
	return xxx_messageInfo_VideoObjectTrackingDatasetMetadata.Size(m)
}
func (m *VideoObjectTrackingDatasetMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoObjectTrackingDatasetMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_VideoObjectTrackingDatasetMetadata proto.InternalMessageInfo

// Model metadata specific to video classification.
type VideoClassificationModelMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VideoClassificationModelMetadata) Reset()         { *m = VideoClassificationModelMetadata{} }
func (m *VideoClassificationModelMetadata) String() string { return proto.CompactTextString(m) }
func (*VideoClassificationModelMetadata) ProtoMessage()    {}
func (*VideoClassificationModelMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec8c6b74d94d2916, []int{2}
}

func (m *VideoClassificationModelMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoClassificationModelMetadata.Unmarshal(m, b)
}
func (m *VideoClassificationModelMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoClassificationModelMetadata.Marshal(b, m, deterministic)
}
func (m *VideoClassificationModelMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoClassificationModelMetadata.Merge(m, src)
}
func (m *VideoClassificationModelMetadata) XXX_Size() int {
	return xxx_messageInfo_VideoClassificationModelMetadata.Size(m)
}
func (m *VideoClassificationModelMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoClassificationModelMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_VideoClassificationModelMetadata proto.InternalMessageInfo

// Model metadata specific to video object tracking.
type VideoObjectTrackingModelMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VideoObjectTrackingModelMetadata) Reset()         { *m = VideoObjectTrackingModelMetadata{} }
func (m *VideoObjectTrackingModelMetadata) String() string { return proto.CompactTextString(m) }
func (*VideoObjectTrackingModelMetadata) ProtoMessage()    {}
func (*VideoObjectTrackingModelMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec8c6b74d94d2916, []int{3}
}

func (m *VideoObjectTrackingModelMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoObjectTrackingModelMetadata.Unmarshal(m, b)
}
func (m *VideoObjectTrackingModelMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoObjectTrackingModelMetadata.Marshal(b, m, deterministic)
}
func (m *VideoObjectTrackingModelMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoObjectTrackingModelMetadata.Merge(m, src)
}
func (m *VideoObjectTrackingModelMetadata) XXX_Size() int {
	return xxx_messageInfo_VideoObjectTrackingModelMetadata.Size(m)
}
func (m *VideoObjectTrackingModelMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoObjectTrackingModelMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_VideoObjectTrackingModelMetadata proto.InternalMessageInfo

func init() {
	proto.RegisterType((*VideoClassificationDatasetMetadata)(nil), "google.cloud.automl.v1beta1.VideoClassificationDatasetMetadata")
	proto.RegisterType((*VideoObjectTrackingDatasetMetadata)(nil), "google.cloud.automl.v1beta1.VideoObjectTrackingDatasetMetadata")
	proto.RegisterType((*VideoClassificationModelMetadata)(nil), "google.cloud.automl.v1beta1.VideoClassificationModelMetadata")
	proto.RegisterType((*VideoObjectTrackingModelMetadata)(nil), "google.cloud.automl.v1beta1.VideoObjectTrackingModelMetadata")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/video.proto", fileDescriptor_ec8c6b74d94d2916)
}

var fileDescriptor_ec8c6b74d94d2916 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd0, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x06, 0x60, 0xb5, 0x03, 0x83, 0x47, 0xc6, 0x16, 0x01, 0x8a, 0x90, 0xd8, 0x6c, 0x2a, 0x36,
	0x33, 0xb5, 0x45, 0x62, 0x21, 0xa2, 0x03, 0xea, 0x80, 0xb2, 0x5c, 0x1c, 0x63, 0x19, 0x5c, 0x5f,
	0x14, 0x5f, 0xfa, 0x5e, 0xbc, 0x06, 0x8f, 0xc2, 0x53, 0xa0, 0xd8, 0x16, 0x52, 0x68, 0xd5, 0xd1,
	0xfa, 0x3f, 0xff, 0xa7, 0x3b, 0x76, 0x6b, 0x10, 0x8d, 0xd3, 0x42, 0x39, 0xec, 0x1b, 0x01, 0x3d,
	0xe1, 0xce, 0x89, 0xfd, 0xa2, 0xd6, 0x04, 0x0b, 0xb1, 0xb7, 0x8d, 0x46, 0xde, 0x76, 0x48, 0x78,
	0x3e, 0x4f, 0x90, 0x47, 0xc8, 0x13, 0xe4, 0x19, 0xce, 0xee, 0x4e, 0xb5, 0x28, 0x07, 0x21, 0xd8,
	0x77, 0xab, 0x80, 0x2c, 0xfa, 0x54, 0x37, 0xbb, 0xc8, 0x3f, 0xa0, 0xb5, 0x02, 0xbc, 0x47, 0x8a,
	0x61, 0x48, 0x69, 0x71, 0xc3, 0x8a, 0xed, 0x30, 0x7b, 0x3d, 0xfa, 0xfa, 0x08, 0x04, 0x41, 0x53,
	0xa9, 0x09, 0x1a, 0x20, 0xf8, 0x53, 0x2f, 0xf5, 0x87, 0x56, 0xf4, 0xda, 0x81, 0xfa, 0xb4, 0xde,
	0xfc, 0x57, 0x05, 0xbb, 0x3e, 0xd2, 0x55, 0x62, 0xa3, 0xdd, 0x81, 0x19, 0x37, 0x8d, 0xcc, 0xea,
	0x6b, 0xc2, 0xae, 0x14, 0xee, 0xf8, 0x89, 0x3b, 0xac, 0x58, 0x6c, 0xd9, 0x0c, 0x3b, 0x6c, 0x26,
	0x6f, 0xcb, 0x4c, 0x0d, 0x3a, 0xf0, 0x86, 0x63, 0x67, 0x84, 0xd1, 0x3e, 0x6e, 0x28, 0x52, 0x04,
	0xad, 0x0d, 0x47, 0x8f, 0xf6, 0x90, 0x9e, 0xdf, 0xd3, 0xf9, 0x53, 0x84, 0xd5, 0x7a, 0x40, 0xd5,
	0xb2, 0x27, 0x2c, 0x5d, 0xb5, 0x4d, 0xe8, 0x67, 0x7a, 0x99, 0x52, 0x29, 0x63, 0x2c, 0x65, 0xcc,
	0x9f, 0xa5, 0xcc, 0xa0, 0x3e, 0x8b, 0xc3, 0xee, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x60,
	0x79, 0x44, 0xe6, 0x01, 0x00, 0x00,
}

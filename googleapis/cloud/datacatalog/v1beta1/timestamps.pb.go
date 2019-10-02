// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/datacatalog/v1beta1/timestamps.proto

package datacatalog

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Timestamps about this resource according to a particular system.
type SystemTimestamps struct {
	// The creation time of the resource within the given system.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The last-modified time of the resource within the given system.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. The expiration time of the resource within the given system.
	// Currently only apllicable to BigQuery resources.
	ExpireTime           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SystemTimestamps) Reset()         { *m = SystemTimestamps{} }
func (m *SystemTimestamps) String() string { return proto.CompactTextString(m) }
func (*SystemTimestamps) ProtoMessage()    {}
func (*SystemTimestamps) Descriptor() ([]byte, []int) {
	return fileDescriptor_a178faa5e507582b, []int{0}
}

func (m *SystemTimestamps) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SystemTimestamps.Unmarshal(m, b)
}
func (m *SystemTimestamps) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SystemTimestamps.Marshal(b, m, deterministic)
}
func (m *SystemTimestamps) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SystemTimestamps.Merge(m, src)
}
func (m *SystemTimestamps) XXX_Size() int {
	return xxx_messageInfo_SystemTimestamps.Size(m)
}
func (m *SystemTimestamps) XXX_DiscardUnknown() {
	xxx_messageInfo_SystemTimestamps.DiscardUnknown(m)
}

var xxx_messageInfo_SystemTimestamps proto.InternalMessageInfo

func (m *SystemTimestamps) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *SystemTimestamps) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *SystemTimestamps) GetExpireTime() *timestamp.Timestamp {
	if m != nil {
		return m.ExpireTime
	}
	return nil
}

func init() {
	proto.RegisterType((*SystemTimestamps)(nil), "google.cloud.datacatalog.v1beta1.SystemTimestamps")
}

func init() {
	proto.RegisterFile("google/cloud/datacatalog/v1beta1/timestamps.proto", fileDescriptor_a178faa5e507582b)
}

var fileDescriptor_a178faa5e507582b = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xd1, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x06, 0x60, 0x99, 0x48, 0x0c, 0xce, 0x82, 0x3a, 0xa1, 0x0a, 0x89, 0x8a, 0x89, 0xc9, 0x56,
	0x60, 0xec, 0x02, 0x5d, 0x59, 0x50, 0x61, 0x62, 0xa9, 0x2e, 0xc9, 0xd5, 0x58, 0x72, 0x7a, 0x96,
	0x73, 0xa9, 0xe0, 0x25, 0x79, 0x06, 0x1e, 0x83, 0x11, 0xc5, 0x36, 0xa1, 0x4b, 0xd5, 0x31, 0x77,
	0xff, 0xf7, 0x5f, 0x94, 0xc8, 0xca, 0x10, 0x19, 0x87, 0xba, 0x71, 0x34, 0xb4, 0xba, 0x05, 0x86,
	0x06, 0x18, 0x1c, 0x19, 0xbd, 0xaf, 0x6a, 0x64, 0xa8, 0x34, 0xdb, 0x0e, 0x7b, 0x86, 0xce, 0xf7,
	0xca, 0x07, 0x62, 0x9a, 0x2d, 0x12, 0x51, 0x91, 0xa8, 0x03, 0xa2, 0x32, 0x99, 0x5f, 0xe7, 0x52,
	0xf0, 0x56, 0x6f, 0x2d, 0xba, 0x76, 0x53, 0xe3, 0x3b, 0xec, 0x2d, 0x85, 0x54, 0x31, 0x05, 0xe2,
	0x53, 0x3d, 0x6c, 0xff, 0x8f, 0xa4, 0xc0, 0xcd, 0x97, 0x90, 0x17, 0x2f, 0x9f, 0x3d, 0x63, 0xf7,
	0x3a, 0x9d, 0x9f, 0x2d, 0x65, 0xd9, 0x04, 0x04, 0xc6, 0xcd, 0x18, 0xbf, 0x14, 0x0b, 0x71, 0x5b,
	0xde, 0xcd, 0x55, 0x7e, 0x9d, 0xbf, 0x2e, 0x35, 0x89, 0xb5, 0x4c, 0xf1, 0x71, 0x30, 0xe2, 0xc1,
	0xb7, 0x13, 0x3e, 0x3b, 0x8d, 0x53, 0x3c, 0xe2, 0x07, 0x59, 0xe2, 0x87, 0xb7, 0x21, 0xe3, 0xe2,
	0x14, 0x5e, 0x15, 0xdf, 0x8f, 0xc5, 0x5a, 0x26, 0x33, 0x4e, 0x57, 0x5e, 0x5e, 0x35, 0xd4, 0xa9,
	0x63, 0x9f, 0xee, 0x59, 0xbc, 0x3d, 0xe5, 0x9d, 0x21, 0x07, 0x3b, 0xa3, 0x28, 0x18, 0x6d, 0x70,
	0x17, 0xbb, 0x75, 0x5a, 0x81, 0xb7, 0xfd, 0xf1, 0x1f, 0xb5, 0x3c, 0x98, 0xfd, 0x08, 0x51, 0x9f,
	0x47, 0x7a, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x38, 0xaa, 0x03, 0x23, 0xe2, 0x01, 0x00, 0x00,
}

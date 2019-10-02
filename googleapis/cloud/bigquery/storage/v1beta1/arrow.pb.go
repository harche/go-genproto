// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/bigquery/storage/v1beta1/arrow.proto

package storage

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Arrow schema.
type ArrowSchema struct {
	// IPC serialized Arrow schema.
	SerializedSchema     []byte   `protobuf:"bytes,1,opt,name=serialized_schema,json=serializedSchema,proto3" json:"serialized_schema,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArrowSchema) Reset()         { *m = ArrowSchema{} }
func (m *ArrowSchema) String() string { return proto.CompactTextString(m) }
func (*ArrowSchema) ProtoMessage()    {}
func (*ArrowSchema) Descriptor() ([]byte, []int) {
	return fileDescriptor_b85627c9be97e962, []int{0}
}

func (m *ArrowSchema) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrowSchema.Unmarshal(m, b)
}
func (m *ArrowSchema) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrowSchema.Marshal(b, m, deterministic)
}
func (m *ArrowSchema) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrowSchema.Merge(m, src)
}
func (m *ArrowSchema) XXX_Size() int {
	return xxx_messageInfo_ArrowSchema.Size(m)
}
func (m *ArrowSchema) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrowSchema.DiscardUnknown(m)
}

var xxx_messageInfo_ArrowSchema proto.InternalMessageInfo

func (m *ArrowSchema) GetSerializedSchema() []byte {
	if m != nil {
		return m.SerializedSchema
	}
	return nil
}

// Arrow RecordBatch.
type ArrowRecordBatch struct {
	// IPC serialized Arrow RecordBatch.
	SerializedRecordBatch []byte `protobuf:"bytes,1,opt,name=serialized_record_batch,json=serializedRecordBatch,proto3" json:"serialized_record_batch,omitempty"`
	// The count of rows in the returning block.
	RowCount             int64    `protobuf:"varint,2,opt,name=row_count,json=rowCount,proto3" json:"row_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ArrowRecordBatch) Reset()         { *m = ArrowRecordBatch{} }
func (m *ArrowRecordBatch) String() string { return proto.CompactTextString(m) }
func (*ArrowRecordBatch) ProtoMessage()    {}
func (*ArrowRecordBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_b85627c9be97e962, []int{1}
}

func (m *ArrowRecordBatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArrowRecordBatch.Unmarshal(m, b)
}
func (m *ArrowRecordBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArrowRecordBatch.Marshal(b, m, deterministic)
}
func (m *ArrowRecordBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArrowRecordBatch.Merge(m, src)
}
func (m *ArrowRecordBatch) XXX_Size() int {
	return xxx_messageInfo_ArrowRecordBatch.Size(m)
}
func (m *ArrowRecordBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_ArrowRecordBatch.DiscardUnknown(m)
}

var xxx_messageInfo_ArrowRecordBatch proto.InternalMessageInfo

func (m *ArrowRecordBatch) GetSerializedRecordBatch() []byte {
	if m != nil {
		return m.SerializedRecordBatch
	}
	return nil
}

func (m *ArrowRecordBatch) GetRowCount() int64 {
	if m != nil {
		return m.RowCount
	}
	return 0
}

func init() {
	proto.RegisterType((*ArrowSchema)(nil), "google.cloud.bigquery.storage.v1beta1.ArrowSchema")
	proto.RegisterType((*ArrowRecordBatch)(nil), "google.cloud.bigquery.storage.v1beta1.ArrowRecordBatch")
}

func init() {
	proto.RegisterFile("google/cloud/bigquery/storage/v1beta1/arrow.proto", fileDescriptor_b85627c9be97e962)
}

var fileDescriptor_b85627c9be97e962 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd0, 0xb1, 0x4b, 0xc4, 0x30,
	0x14, 0x06, 0x70, 0xaa, 0x20, 0x1a, 0x1d, 0xce, 0x82, 0x78, 0xe0, 0x72, 0x1c, 0x08, 0x27, 0x42,
	0x42, 0x11, 0x1c, 0x74, 0xb2, 0xae, 0x0e, 0x52, 0x37, 0x97, 0xf2, 0x9a, 0x3e, 0xde, 0x15, 0x7a,
	0x7d, 0xe7, 0x6b, 0x6a, 0xd1, 0xdd, 0xff, 0x5b, 0x9a, 0x46, 0xae, 0x9b, 0x37, 0xbe, 0x7c, 0xf9,
	0x3d, 0x92, 0x4f, 0x25, 0xc4, 0x4c, 0x35, 0x1a, 0x5b, 0x73, 0x57, 0x9a, 0xa2, 0xa2, 0x8f, 0x0e,
	0xe5, 0xcb, 0xb4, 0x8e, 0x05, 0x08, 0xcd, 0x67, 0x52, 0xa0, 0x83, 0xc4, 0x80, 0x08, 0xf7, 0x7a,
	0x2b, 0xec, 0x38, 0xbe, 0x1e, 0x89, 0xf6, 0x44, 0xff, 0x11, 0x1d, 0x88, 0x0e, 0x64, 0xf9, 0xa0,
	0x4e, 0x9f, 0x06, 0xf5, 0x66, 0xd7, 0xb8, 0x81, 0xf8, 0x56, 0x9d, 0xb7, 0x28, 0x15, 0xd4, 0xd5,
	0x37, 0x96, 0x79, 0xeb, 0x0f, 0xe7, 0xd1, 0x22, 0x5a, 0x9d, 0x65, 0xb3, 0x5d, 0x30, 0x5e, 0x5e,
	0x92, 0x9a, 0x79, 0x9b, 0xa1, 0x65, 0x29, 0x53, 0x70, 0x76, 0x1d, 0xdf, 0xab, 0xcb, 0xc9, 0x02,
	0xf1, 0x49, 0x5e, 0x0c, 0x51, 0x58, 0x73, 0xb1, 0x8b, 0xa7, 0xee, 0x4a, 0x9d, 0x08, 0xf7, 0xb9,
	0xe5, 0xae, 0x71, 0xf3, 0x83, 0x45, 0xb4, 0x3a, 0xcc, 0x8e, 0x85, 0xfb, 0xe7, 0x61, 0x4e, 0x7f,
	0x22, 0x75, 0x63, 0x79, 0xa3, 0xf7, 0xfa, 0x52, 0xaa, 0xfc, 0xa3, 0x5e, 0x87, 0x16, 0xde, 0x5f,
	0x02, 0x21, 0xae, 0xa1, 0x21, 0xcd, 0x42, 0x86, 0xb0, 0xf1, 0x0d, 0x99, 0x31, 0x82, 0x6d, 0xd5,
	0xfe, 0xd3, 0xeb, 0x63, 0x98, 0x8b, 0x23, 0x0f, 0xef, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x61,
	0xf1, 0x2a, 0xa7, 0x8f, 0x01, 0x00, 0x00,
}

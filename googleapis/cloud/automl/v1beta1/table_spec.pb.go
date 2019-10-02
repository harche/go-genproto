// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/table_spec.proto

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

// A specification of a relational table.
// The table's schema is represented via its child column specs. It is
// pre-populated as part of ImportData by schema inference algorithm, the
// version of which is a required parameter of ImportData InputConfig.
// Note: While working with a table, at times the schema may be
// inconsistent with the data in the table (e.g. string in a FLOAT64 column).
// The consistency validation is done upon creation of a model.
// Used by:
//   *   Tables
type TableSpec struct {
	// Output only. The resource name of the table spec.
	// Form:
	//
	// `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/tableSpecs/{table_spec_id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// column_spec_id of the time column. Only used if the parent dataset's
	// ml_use_column_spec_id is not set. Used to split rows into TRAIN, VALIDATE
	// and TEST sets such that oldest rows go to TRAIN set, newest to TEST, and
	// those in between to VALIDATE.
	// Required type: TIMESTAMP.
	// If both this column and ml_use_column are not set, then ML use of all rows
	// will be assigned by AutoML. NOTE: Updates of this field will instantly
	// affect any other users concurrently working with the dataset.
	TimeColumnSpecId string `protobuf:"bytes,2,opt,name=time_column_spec_id,json=timeColumnSpecId,proto3" json:"time_column_spec_id,omitempty"`
	// Output only. The number of rows (i.e. examples) in the table.
	RowCount int64 `protobuf:"varint,3,opt,name=row_count,json=rowCount,proto3" json:"row_count,omitempty"`
	// Output only. The number of valid rows (i.e. without values that don't match
	// DataType-s of their columns).
	ValidRowCount int64 `protobuf:"varint,4,opt,name=valid_row_count,json=validRowCount,proto3" json:"valid_row_count,omitempty"`
	// Output only. The number of columns of the table. That is, the number of
	// child ColumnSpec-s.
	ColumnCount int64 `protobuf:"varint,7,opt,name=column_count,json=columnCount,proto3" json:"column_count,omitempty"`
	// Output only. Input configs via which data currently residing in the table
	// had been imported.
	InputConfigs []*InputConfig `protobuf:"bytes,5,rep,name=input_configs,json=inputConfigs,proto3" json:"input_configs,omitempty"`
	// Used to perform consistent read-modify-write updates. If not set, a blind
	// "overwrite" update happens.
	Etag                 string   `protobuf:"bytes,6,opt,name=etag,proto3" json:"etag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TableSpec) Reset()         { *m = TableSpec{} }
func (m *TableSpec) String() string { return proto.CompactTextString(m) }
func (*TableSpec) ProtoMessage()    {}
func (*TableSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_67819f165062ce17, []int{0}
}

func (m *TableSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TableSpec.Unmarshal(m, b)
}
func (m *TableSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TableSpec.Marshal(b, m, deterministic)
}
func (m *TableSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableSpec.Merge(m, src)
}
func (m *TableSpec) XXX_Size() int {
	return xxx_messageInfo_TableSpec.Size(m)
}
func (m *TableSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_TableSpec.DiscardUnknown(m)
}

var xxx_messageInfo_TableSpec proto.InternalMessageInfo

func (m *TableSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TableSpec) GetTimeColumnSpecId() string {
	if m != nil {
		return m.TimeColumnSpecId
	}
	return ""
}

func (m *TableSpec) GetRowCount() int64 {
	if m != nil {
		return m.RowCount
	}
	return 0
}

func (m *TableSpec) GetValidRowCount() int64 {
	if m != nil {
		return m.ValidRowCount
	}
	return 0
}

func (m *TableSpec) GetColumnCount() int64 {
	if m != nil {
		return m.ColumnCount
	}
	return 0
}

func (m *TableSpec) GetInputConfigs() []*InputConfig {
	if m != nil {
		return m.InputConfigs
	}
	return nil
}

func (m *TableSpec) GetEtag() string {
	if m != nil {
		return m.Etag
	}
	return ""
}

func init() {
	proto.RegisterType((*TableSpec)(nil), "google.cloud.automl.v1beta1.TableSpec")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/table_spec.proto", fileDescriptor_67819f165062ce17)
}

var fileDescriptor_67819f165062ce17 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd1, 0xd1, 0x4a, 0xeb, 0x30,
	0x18, 0x07, 0x70, 0xda, 0xed, 0xec, 0x9c, 0x65, 0x1b, 0x47, 0xea, 0x4d, 0xd9, 0x44, 0xa7, 0x88,
	0xf4, 0x42, 0x5b, 0xa6, 0x77, 0xf5, 0x6a, 0xf6, 0x42, 0x06, 0x0e, 0xa4, 0x8a, 0x17, 0x32, 0x28,
	0x59, 0x1a, 0x43, 0x20, 0xcd, 0x57, 0xda, 0x74, 0x7b, 0x16, 0x1f, 0xc0, 0x97, 0xf1, 0x51, 0x7c,
	0x0a, 0x49, 0x52, 0xd9, 0x8d, 0xec, 0x2e, 0x5f, 0xfe, 0xbf, 0x7c, 0x5f, 0x93, 0xa2, 0x4b, 0x06,
	0xc0, 0x04, 0x8d, 0x88, 0x80, 0x26, 0x8f, 0x70, 0xa3, 0xa0, 0x10, 0xd1, 0x66, 0xb6, 0xa6, 0x0a,
	0xcf, 0x22, 0x85, 0xd7, 0x82, 0x66, 0x75, 0x49, 0x49, 0x58, 0x56, 0xa0, 0xc0, 0x9b, 0x58, 0x1d,
	0x1a, 0x1d, 0x5a, 0x1d, 0xb6, 0x7a, 0x7c, 0xbe, 0xaf, 0x15, 0x07, 0xdb, 0x62, 0x7c, 0xd4, 0x2a,
	0x5c, 0xf2, 0x08, 0x4b, 0x09, 0x0a, 0x2b, 0x0e, 0xb2, 0xb6, 0xe9, 0xd9, 0xbb, 0x8b, 0xfa, 0xcf,
	0x7a, 0xea, 0x53, 0x49, 0x89, 0xe7, 0xa1, 0xae, 0xc4, 0x05, 0xf5, 0x9d, 0xa9, 0x13, 0xf4, 0x53,
	0xb3, 0xf6, 0xae, 0xd0, 0xa1, 0xe2, 0x05, 0xcd, 0x08, 0x88, 0xa6, 0x90, 0xe6, 0xe3, 0x32, 0x9e,
	0xfb, 0xae, 0x21, 0x07, 0x3a, 0x4a, 0x4c, 0xa2, 0x1b, 0x2c, 0x72, 0x6f, 0x82, 0xfa, 0x15, 0x6c,
	0x33, 0x02, 0x8d, 0x54, 0x7e, 0x67, 0xea, 0x04, 0x9d, 0xf4, 0x5f, 0x05, 0xdb, 0x44, 0xd7, 0xde,
	0x05, 0xfa, 0xbf, 0xc1, 0x82, 0xe7, 0xd9, 0x8e, 0x74, 0x0d, 0x19, 0x99, 0xed, 0xf4, 0xc7, 0x9d,
	0xa2, 0x61, 0x3b, 0xce, 0xa2, 0xbf, 0x06, 0x0d, 0xec, 0x9e, 0x25, 0x4b, 0x34, 0xe2, 0xb2, 0x6c,
	0x54, 0x46, 0x40, 0xbe, 0x71, 0x56, 0xfb, 0x7f, 0xa6, 0x9d, 0x60, 0x70, 0x1d, 0x84, 0x7b, 0x5e,
	0x2c, 0x5c, 0xe8, 0x13, 0x89, 0x39, 0x90, 0x0e, 0xf9, 0xae, 0xa8, 0xf5, 0xcd, 0xa9, 0xc2, 0xcc,
	0xef, 0xd9, 0x9b, 0xeb, 0xf5, 0xdd, 0x87, 0x83, 0x4e, 0x08, 0x14, 0xfb, 0x3a, 0x3e, 0x3a, 0xaf,
	0xf3, 0x36, 0x66, 0x20, 0xb0, 0x64, 0x21, 0x54, 0x2c, 0x62, 0x54, 0x9a, 0xd7, 0x8d, 0x6c, 0x84,
	0x4b, 0x5e, 0xff, 0xfa, 0x93, 0x6e, 0x6d, 0xf9, 0xe9, 0x4e, 0xee, 0x0d, 0x5c, 0x25, 0x1a, 0xad,
	0xe6, 0x8d, 0x82, 0xa5, 0x58, 0xbd, 0x58, 0xf4, 0xe5, 0x1e, 0xdb, 0x34, 0x8e, 0x4d, 0x1c, 0xc7,
	0x26, 0x7f, 0x88, 0xe3, 0x16, 0xac, 0x7b, 0x66, 0xd8, 0xcd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x1c, 0x57, 0x2f, 0xa6, 0x5b, 0x02, 0x00, 0x00,
}

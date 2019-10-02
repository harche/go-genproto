// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/spanner/v1/keys.proto

package spanner

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// KeyRange represents a range of rows in a table or index.
//
// A range has a start key and an end key. These keys can be open or
// closed, indicating if the range includes rows with that key.
//
// Keys are represented by lists, where the ith value in the list
// corresponds to the ith component of the table or index primary key.
// Individual values are encoded as described
// [here][google.spanner.v1.TypeCode].
//
// For example, consider the following table definition:
//
//     CREATE TABLE UserEvents (
//       UserName STRING(MAX),
//       EventDate STRING(10)
//     ) PRIMARY KEY(UserName, EventDate);
//
// The following keys name rows in this table:
//
//     ["Bob", "2014-09-23"]
//     ["Alfred", "2015-06-12"]
//
// Since the `UserEvents` table's `PRIMARY KEY` clause names two
// columns, each `UserEvents` key has two elements; the first is the
// `UserName`, and the second is the `EventDate`.
//
// Key ranges with multiple components are interpreted
// lexicographically by component using the table or index key's declared
// sort order. For example, the following range returns all events for
// user `"Bob"` that occurred in the year 2015:
//
//     "start_closed": ["Bob", "2015-01-01"]
//     "end_closed": ["Bob", "2015-12-31"]
//
// Start and end keys can omit trailing key components. This affects the
// inclusion and exclusion of rows that exactly match the provided key
// components: if the key is closed, then rows that exactly match the
// provided components are included; if the key is open, then rows
// that exactly match are not included.
//
// For example, the following range includes all events for `"Bob"` that
// occurred during and after the year 2000:
//
//     "start_closed": ["Bob", "2000-01-01"]
//     "end_closed": ["Bob"]
//
// The next example retrieves all events for `"Bob"`:
//
//     "start_closed": ["Bob"]
//     "end_closed": ["Bob"]
//
// To retrieve events before the year 2000:
//
//     "start_closed": ["Bob"]
//     "end_open": ["Bob", "2000-01-01"]
//
// The following range includes all rows in the table:
//
//     "start_closed": []
//     "end_closed": []
//
// This range returns all users whose `UserName` begins with any
// character from A to C:
//
//     "start_closed": ["A"]
//     "end_open": ["D"]
//
// This range returns all users whose `UserName` begins with B:
//
//     "start_closed": ["B"]
//     "end_open": ["C"]
//
// Key ranges honor column sort order. For example, suppose a table is
// defined as follows:
//
//     CREATE TABLE DescendingSortedTable {
//       Key INT64,
//       ...
//     ) PRIMARY KEY(Key DESC);
//
// The following range retrieves all rows with key values between 1
// and 100 inclusive:
//
//     "start_closed": ["100"]
//     "end_closed": ["1"]
//
// Note that 100 is passed as the start, and 1 is passed as the end,
// because `Key` is a descending column in the schema.
type KeyRange struct {
	// The start key must be provided. It can be either closed or open.
	//
	// Types that are valid to be assigned to StartKeyType:
	//	*KeyRange_StartClosed
	//	*KeyRange_StartOpen
	StartKeyType isKeyRange_StartKeyType `protobuf_oneof:"start_key_type"`
	// The end key must be provided. It can be either closed or open.
	//
	// Types that are valid to be assigned to EndKeyType:
	//	*KeyRange_EndClosed
	//	*KeyRange_EndOpen
	EndKeyType           isKeyRange_EndKeyType `protobuf_oneof:"end_key_type"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *KeyRange) Reset()         { *m = KeyRange{} }
func (m *KeyRange) String() string { return proto.CompactTextString(m) }
func (*KeyRange) ProtoMessage()    {}
func (*KeyRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_558e8abca28348dd, []int{0}
}

func (m *KeyRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyRange.Unmarshal(m, b)
}
func (m *KeyRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyRange.Marshal(b, m, deterministic)
}
func (m *KeyRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyRange.Merge(m, src)
}
func (m *KeyRange) XXX_Size() int {
	return xxx_messageInfo_KeyRange.Size(m)
}
func (m *KeyRange) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyRange.DiscardUnknown(m)
}

var xxx_messageInfo_KeyRange proto.InternalMessageInfo

type isKeyRange_StartKeyType interface {
	isKeyRange_StartKeyType()
}

type KeyRange_StartClosed struct {
	StartClosed *_struct.ListValue `protobuf:"bytes,1,opt,name=start_closed,json=startClosed,proto3,oneof"`
}

type KeyRange_StartOpen struct {
	StartOpen *_struct.ListValue `protobuf:"bytes,2,opt,name=start_open,json=startOpen,proto3,oneof"`
}

func (*KeyRange_StartClosed) isKeyRange_StartKeyType() {}

func (*KeyRange_StartOpen) isKeyRange_StartKeyType() {}

func (m *KeyRange) GetStartKeyType() isKeyRange_StartKeyType {
	if m != nil {
		return m.StartKeyType
	}
	return nil
}

func (m *KeyRange) GetStartClosed() *_struct.ListValue {
	if x, ok := m.GetStartKeyType().(*KeyRange_StartClosed); ok {
		return x.StartClosed
	}
	return nil
}

func (m *KeyRange) GetStartOpen() *_struct.ListValue {
	if x, ok := m.GetStartKeyType().(*KeyRange_StartOpen); ok {
		return x.StartOpen
	}
	return nil
}

type isKeyRange_EndKeyType interface {
	isKeyRange_EndKeyType()
}

type KeyRange_EndClosed struct {
	EndClosed *_struct.ListValue `protobuf:"bytes,3,opt,name=end_closed,json=endClosed,proto3,oneof"`
}

type KeyRange_EndOpen struct {
	EndOpen *_struct.ListValue `protobuf:"bytes,4,opt,name=end_open,json=endOpen,proto3,oneof"`
}

func (*KeyRange_EndClosed) isKeyRange_EndKeyType() {}

func (*KeyRange_EndOpen) isKeyRange_EndKeyType() {}

func (m *KeyRange) GetEndKeyType() isKeyRange_EndKeyType {
	if m != nil {
		return m.EndKeyType
	}
	return nil
}

func (m *KeyRange) GetEndClosed() *_struct.ListValue {
	if x, ok := m.GetEndKeyType().(*KeyRange_EndClosed); ok {
		return x.EndClosed
	}
	return nil
}

func (m *KeyRange) GetEndOpen() *_struct.ListValue {
	if x, ok := m.GetEndKeyType().(*KeyRange_EndOpen); ok {
		return x.EndOpen
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*KeyRange) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*KeyRange_StartClosed)(nil),
		(*KeyRange_StartOpen)(nil),
		(*KeyRange_EndClosed)(nil),
		(*KeyRange_EndOpen)(nil),
	}
}

// `KeySet` defines a collection of Cloud Spanner keys and/or key ranges. All
// the keys are expected to be in the same table or index. The keys need
// not be sorted in any particular way.
//
// If the same key is specified multiple times in the set (for example
// if two ranges, two keys, or a key and a range overlap), Cloud Spanner
// behaves as if the key were only specified once.
type KeySet struct {
	// A list of specific keys. Entries in `keys` should have exactly as
	// many elements as there are columns in the primary or index key
	// with which this `KeySet` is used.  Individual key values are
	// encoded as described [here][google.spanner.v1.TypeCode].
	Keys []*_struct.ListValue `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	// A list of key ranges. See [KeyRange][google.spanner.v1.KeyRange] for more
	// information about key range specifications.
	Ranges []*KeyRange `protobuf:"bytes,2,rep,name=ranges,proto3" json:"ranges,omitempty"`
	// For convenience `all` can be set to `true` to indicate that this
	// `KeySet` matches all keys in the table or index. Note that any keys
	// specified in `keys` or `ranges` are only yielded once.
	All                  bool     `protobuf:"varint,3,opt,name=all,proto3" json:"all,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeySet) Reset()         { *m = KeySet{} }
func (m *KeySet) String() string { return proto.CompactTextString(m) }
func (*KeySet) ProtoMessage()    {}
func (*KeySet) Descriptor() ([]byte, []int) {
	return fileDescriptor_558e8abca28348dd, []int{1}
}

func (m *KeySet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeySet.Unmarshal(m, b)
}
func (m *KeySet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeySet.Marshal(b, m, deterministic)
}
func (m *KeySet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeySet.Merge(m, src)
}
func (m *KeySet) XXX_Size() int {
	return xxx_messageInfo_KeySet.Size(m)
}
func (m *KeySet) XXX_DiscardUnknown() {
	xxx_messageInfo_KeySet.DiscardUnknown(m)
}

var xxx_messageInfo_KeySet proto.InternalMessageInfo

func (m *KeySet) GetKeys() []*_struct.ListValue {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *KeySet) GetRanges() []*KeyRange {
	if m != nil {
		return m.Ranges
	}
	return nil
}

func (m *KeySet) GetAll() bool {
	if m != nil {
		return m.All
	}
	return false
}

func init() {
	proto.RegisterType((*KeyRange)(nil), "google.spanner.v1.KeyRange")
	proto.RegisterType((*KeySet)(nil), "google.spanner.v1.KeySet")
}

func init() { proto.RegisterFile("google/spanner/v1/keys.proto", fileDescriptor_558e8abca28348dd) }

var fileDescriptor_558e8abca28348dd = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6b, 0xea, 0x30,
	0x1c, 0xc7, 0x5f, 0xab, 0xf8, 0x34, 0x8a, 0xf8, 0x0a, 0x8f, 0x57, 0x7c, 0x3b, 0x88, 0xa7, 0x9d,
	0x52, 0x3a, 0x0f, 0x1b, 0x78, 0x18, 0xd4, 0xc3, 0x06, 0x0e, 0x26, 0x15, 0x3c, 0x0c, 0x41, 0xa2,
	0xfd, 0xad, 0x14, 0xb3, 0x24, 0x34, 0xa9, 0xd0, 0xd3, 0xfe, 0x87, 0xfd, 0x05, 0x3b, 0xef, 0x4f,
	0xd9, 0x5f, 0x35, 0x92, 0xa6, 0x63, 0x20, 0x6c, 0xde, 0x12, 0x3e, 0xbf, 0xcf, 0xf7, 0x9b, 0x26,
	0x45, 0x67, 0x29, 0xe7, 0x29, 0x85, 0x40, 0x0a, 0xc2, 0x18, 0xe4, 0xc1, 0x21, 0x0c, 0xf6, 0x50,
	0x4a, 0x2c, 0x72, 0xae, 0xb8, 0xf7, 0xa7, 0xa2, 0xd8, 0x52, 0x7c, 0x08, 0x87, 0xb5, 0x40, 0x44,
	0x16, 0x10, 0xc6, 0xb8, 0x22, 0x2a, 0xe3, 0xcc, 0x0a, 0x9f, 0xd4, 0xec, 0xb6, 0xc5, 0x63, 0x20,
	0x55, 0x5e, 0xec, 0x54, 0x45, 0xc7, 0xaf, 0x2e, 0x6a, 0xcf, 0xa1, 0x8c, 0x09, 0x4b, 0xc1, 0xbb,
	0x46, 0x3d, 0xa9, 0x48, 0xae, 0x36, 0x3b, 0xca, 0x25, 0x24, 0xbe, 0x33, 0x72, 0xce, 0xbb, 0x17,
	0x43, 0x6c, 0x2b, 0xeb, 0x04, 0x7c, 0x97, 0x49, 0xb5, 0x22, 0xb4, 0x80, 0xdb, 0x5f, 0x71, 0xd7,
	0x18, 0x33, 0x23, 0x78, 0x53, 0x84, 0xaa, 0x00, 0x2e, 0x80, 0xf9, 0xee, 0x09, 0x7a, 0xc7, 0xcc,
	0xdf, 0x0b, 0x60, 0x5a, 0x06, 0x96, 0xd4, 0xdd, 0x8d, 0x1f, 0x65, 0x27, 0xee, 0x00, 0x4b, 0x6c,
	0xf3, 0x25, 0x6a, 0x6b, 0xd9, 0xf4, 0x36, 0x4f, 0x50, 0x7f, 0x03, 0x4b, 0x74, 0x6b, 0x34, 0x40,
	0xfd, 0xea, 0xc8, 0x7b, 0x28, 0x37, 0xaa, 0x14, 0x10, 0xf5, 0x51, 0x4f, 0x47, 0xd5, 0xfb, 0xf1,
	0x33, 0x6a, 0xcd, 0xa1, 0x5c, 0x82, 0xf2, 0x30, 0x6a, 0xea, 0x97, 0xf0, 0x9d, 0x51, 0xe3, 0xfb,
	0x82, 0xd8, 0xcc, 0x79, 0x13, 0xd4, 0xca, 0xf5, 0xc5, 0x4a, 0xdf, 0x35, 0xc6, 0x7f, 0x7c, 0xf4,
	0x78, 0xb8, 0xbe, 0xfc, 0xd8, 0x8e, 0x7a, 0x03, 0xd4, 0x20, 0x94, 0x9a, 0xef, 0x6f, 0xc7, 0x7a,
	0x19, 0xbd, 0x38, 0xe8, 0xef, 0x8e, 0x3f, 0x1d, 0xcb, 0x51, 0x67, 0x0e, 0xa5, 0x5c, 0xe8, 0xfa,
	0x85, 0xf3, 0x70, 0x65, 0x79, 0xca, 0x29, 0x61, 0x29, 0xe6, 0x79, 0x1a, 0xa4, 0xc0, 0xcc, 0xe1,
	0x82, 0x0a, 0x11, 0x91, 0xc9, 0x2f, 0xbf, 0xd5, 0xd4, 0x2e, 0xdf, 0xdc, 0x7f, 0x37, 0x95, 0x3a,
	0xa3, 0xbc, 0x48, 0xf0, 0xd2, 0x16, 0xac, 0xc2, 0xf7, 0x9a, 0xac, 0x0d, 0x59, 0x5b, 0xb2, 0x5e,
	0x85, 0xdb, 0x96, 0x09, 0x9e, 0x7c, 0x04, 0x00, 0x00, 0xff, 0xff, 0x27, 0x88, 0xea, 0x11, 0xae,
	0x02, 0x00, 0x00,
}

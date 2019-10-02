// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/firestore/v1/query.proto

package firestore

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A sort direction.
type StructuredQuery_Direction int32

const (
	// Unspecified.
	StructuredQuery_DIRECTION_UNSPECIFIED StructuredQuery_Direction = 0
	// Ascending.
	StructuredQuery_ASCENDING StructuredQuery_Direction = 1
	// Descending.
	StructuredQuery_DESCENDING StructuredQuery_Direction = 2
)

var StructuredQuery_Direction_name = map[int32]string{
	0: "DIRECTION_UNSPECIFIED",
	1: "ASCENDING",
	2: "DESCENDING",
}

var StructuredQuery_Direction_value = map[string]int32{
	"DIRECTION_UNSPECIFIED": 0,
	"ASCENDING":             1,
	"DESCENDING":            2,
}

func (x StructuredQuery_Direction) String() string {
	return proto.EnumName(StructuredQuery_Direction_name, int32(x))
}

func (StructuredQuery_Direction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_45ff6dfb507a4c5e, []int{0, 0}
}

// A composite filter operator.
type StructuredQuery_CompositeFilter_Operator int32

const (
	// Unspecified. This value must not be used.
	StructuredQuery_CompositeFilter_OPERATOR_UNSPECIFIED StructuredQuery_CompositeFilter_Operator = 0
	// The results are required to satisfy each of the combined filters.
	StructuredQuery_CompositeFilter_AND StructuredQuery_CompositeFilter_Operator = 1
)

var StructuredQuery_CompositeFilter_Operator_name = map[int32]string{
	0: "OPERATOR_UNSPECIFIED",
	1: "AND",
}

var StructuredQuery_CompositeFilter_Operator_value = map[string]int32{
	"OPERATOR_UNSPECIFIED": 0,
	"AND":                  1,
}

func (x StructuredQuery_CompositeFilter_Operator) String() string {
	return proto.EnumName(StructuredQuery_CompositeFilter_Operator_name, int32(x))
}

func (StructuredQuery_CompositeFilter_Operator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_45ff6dfb507a4c5e, []int{0, 2, 0}
}

// A field filter operator.
type StructuredQuery_FieldFilter_Operator int32

const (
	// Unspecified. This value must not be used.
	StructuredQuery_FieldFilter_OPERATOR_UNSPECIFIED StructuredQuery_FieldFilter_Operator = 0
	// Less than. Requires that the field come first in `order_by`.
	StructuredQuery_FieldFilter_LESS_THAN StructuredQuery_FieldFilter_Operator = 1
	// Less than or equal. Requires that the field come first in `order_by`.
	StructuredQuery_FieldFilter_LESS_THAN_OR_EQUAL StructuredQuery_FieldFilter_Operator = 2
	// Greater than. Requires that the field come first in `order_by`.
	StructuredQuery_FieldFilter_GREATER_THAN StructuredQuery_FieldFilter_Operator = 3
	// Greater than or equal. Requires that the field come first in
	// `order_by`.
	StructuredQuery_FieldFilter_GREATER_THAN_OR_EQUAL StructuredQuery_FieldFilter_Operator = 4
	// Equal.
	StructuredQuery_FieldFilter_EQUAL StructuredQuery_FieldFilter_Operator = 5
	// Contains. Requires that the field is an array.
	StructuredQuery_FieldFilter_ARRAY_CONTAINS StructuredQuery_FieldFilter_Operator = 7
)

var StructuredQuery_FieldFilter_Operator_name = map[int32]string{
	0: "OPERATOR_UNSPECIFIED",
	1: "LESS_THAN",
	2: "LESS_THAN_OR_EQUAL",
	3: "GREATER_THAN",
	4: "GREATER_THAN_OR_EQUAL",
	5: "EQUAL",
	7: "ARRAY_CONTAINS",
}

var StructuredQuery_FieldFilter_Operator_value = map[string]int32{
	"OPERATOR_UNSPECIFIED":  0,
	"LESS_THAN":             1,
	"LESS_THAN_OR_EQUAL":    2,
	"GREATER_THAN":          3,
	"GREATER_THAN_OR_EQUAL": 4,
	"EQUAL":                 5,
	"ARRAY_CONTAINS":        7,
}

func (x StructuredQuery_FieldFilter_Operator) String() string {
	return proto.EnumName(StructuredQuery_FieldFilter_Operator_name, int32(x))
}

func (StructuredQuery_FieldFilter_Operator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_45ff6dfb507a4c5e, []int{0, 3, 0}
}

// A unary operator.
type StructuredQuery_UnaryFilter_Operator int32

const (
	// Unspecified. This value must not be used.
	StructuredQuery_UnaryFilter_OPERATOR_UNSPECIFIED StructuredQuery_UnaryFilter_Operator = 0
	// Test if a field is equal to NaN.
	StructuredQuery_UnaryFilter_IS_NAN StructuredQuery_UnaryFilter_Operator = 2
	// Test if an expression evaluates to Null.
	StructuredQuery_UnaryFilter_IS_NULL StructuredQuery_UnaryFilter_Operator = 3
)

var StructuredQuery_UnaryFilter_Operator_name = map[int32]string{
	0: "OPERATOR_UNSPECIFIED",
	2: "IS_NAN",
	3: "IS_NULL",
}

var StructuredQuery_UnaryFilter_Operator_value = map[string]int32{
	"OPERATOR_UNSPECIFIED": 0,
	"IS_NAN":               2,
	"IS_NULL":              3,
}

func (x StructuredQuery_UnaryFilter_Operator) String() string {
	return proto.EnumName(StructuredQuery_UnaryFilter_Operator_name, int32(x))
}

func (StructuredQuery_UnaryFilter_Operator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_45ff6dfb507a4c5e, []int{0, 4, 0}
}

// A Firestore query.
type StructuredQuery struct {
	// The projection to return.
	Select *StructuredQuery_Projection `protobuf:"bytes,1,opt,name=select,proto3" json:"select,omitempty"`
	// The collections to query.
	From []*StructuredQuery_CollectionSelector `protobuf:"bytes,2,rep,name=from,proto3" json:"from,omitempty"`
	// The filter to apply.
	Where *StructuredQuery_Filter `protobuf:"bytes,3,opt,name=where,proto3" json:"where,omitempty"`
	// The order to apply to the query results.
	//
	// Firestore guarantees a stable ordering through the following rules:
	//
	//  * Any field required to appear in `order_by`, that is not already
	//    specified in `order_by`, is appended to the order in field name order
	//    by default.
	//  * If an order on `__name__` is not specified, it is appended by default.
	//
	// Fields are appended with the same sort direction as the last order
	// specified, or 'ASCENDING' if no order was specified. For example:
	//
	//  * `SELECT * FROM Foo ORDER BY A` becomes
	//    `SELECT * FROM Foo ORDER BY A, __name__`
	//  * `SELECT * FROM Foo ORDER BY A DESC` becomes
	//    `SELECT * FROM Foo ORDER BY A DESC, __name__ DESC`
	//  * `SELECT * FROM Foo WHERE A > 1` becomes
	//    `SELECT * FROM Foo WHERE A > 1 ORDER BY A, __name__`
	OrderBy []*StructuredQuery_Order `protobuf:"bytes,4,rep,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
	// A starting point for the query results.
	StartAt *Cursor `protobuf:"bytes,7,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	// A end point for the query results.
	EndAt *Cursor `protobuf:"bytes,8,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	// The number of results to skip.
	//
	// Applies before limit, but after all other constraints. Must be >= 0 if
	// specified.
	Offset int32 `protobuf:"varint,6,opt,name=offset,proto3" json:"offset,omitempty"`
	// The maximum number of results to return.
	//
	// Applies after all other constraints.
	// Must be >= 0 if specified.
	Limit                *wrappers.Int32Value `protobuf:"bytes,5,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StructuredQuery) Reset()         { *m = StructuredQuery{} }
func (m *StructuredQuery) String() string { return proto.CompactTextString(m) }
func (*StructuredQuery) ProtoMessage()    {}
func (*StructuredQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_45ff6dfb507a4c5e, []int{0}
}

func (m *StructuredQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StructuredQuery.Unmarshal(m, b)
}
func (m *StructuredQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StructuredQuery.Marshal(b, m, deterministic)
}
func (m *StructuredQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StructuredQuery.Merge(m, src)
}
func (m *StructuredQuery) XXX_Size() int {
	return xxx_messageInfo_StructuredQuery.Size(m)
}
func (m *StructuredQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_StructuredQuery.DiscardUnknown(m)
}

var xxx_messageInfo_StructuredQuery proto.InternalMessageInfo

func (m *StructuredQuery) GetSelect() *StructuredQuery_Projection {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *StructuredQuery) GetFrom() []*StructuredQuery_CollectionSelector {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *StructuredQuery) GetWhere() *StructuredQuery_Filter {
	if m != nil {
		return m.Where
	}
	return nil
}

func (m *StructuredQuery) GetOrderBy() []*StructuredQuery_Order {
	if m != nil {
		return m.OrderBy
	}
	return nil
}

func (m *StructuredQuery) GetStartAt() *Cursor {
	if m != nil {
		return m.StartAt
	}
	return nil
}

func (m *StructuredQuery) GetEndAt() *Cursor {
	if m != nil {
		return m.EndAt
	}
	return nil
}

func (m *StructuredQuery) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *StructuredQuery) GetLimit() *wrappers.Int32Value {
	if m != nil {
		return m.Limit
	}
	return nil
}

// A selection of a collection, such as `messages as m1`.
type StructuredQuery_CollectionSelector struct {
	// The collection ID.
	// When set, selects only collections with this ID.
	CollectionId string `protobuf:"bytes,2,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	// When false, selects only collections that are immediate children of
	// the `parent` specified in the containing `RunQueryRequest`.
	// When true, selects all descendant collections.
	AllDescendants       bool     `protobuf:"varint,3,opt,name=all_descendants,json=allDescendants,proto3" json:"all_descendants,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StructuredQuery_CollectionSelector) Reset()         { *m = StructuredQuery_CollectionSelector{} }
func (m *StructuredQuery_CollectionSelector) String() string { return proto.CompactTextString(m) }
func (*StructuredQuery_CollectionSelector) ProtoMessage()    {}
func (*StructuredQuery_CollectionSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_45ff6dfb507a4c5e, []int{0, 0}
}

func (m *StructuredQuery_CollectionSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StructuredQuery_CollectionSelector.Unmarshal(m, b)
}
func (m *StructuredQuery_CollectionSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StructuredQuery_CollectionSelector.Marshal(b, m, deterministic)
}
func (m *StructuredQuery_CollectionSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StructuredQuery_CollectionSelector.Merge(m, src)
}
func (m *StructuredQuery_CollectionSelector) XXX_Size() int {
	return xxx_messageInfo_StructuredQuery_CollectionSelector.Size(m)
}
func (m *StructuredQuery_CollectionSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_StructuredQuery_CollectionSelector.DiscardUnknown(m)
}

var xxx_messageInfo_StructuredQuery_CollectionSelector proto.InternalMessageInfo

func (m *StructuredQuery_CollectionSelector) GetCollectionId() string {
	if m != nil {
		return m.CollectionId
	}
	return ""
}

func (m *StructuredQuery_CollectionSelector) GetAllDescendants() bool {
	if m != nil {
		return m.AllDescendants
	}
	return false
}

// A filter.
type StructuredQuery_Filter struct {
	// The type of filter.
	//
	// Types that are valid to be assigned to FilterType:
	//	*StructuredQuery_Filter_CompositeFilter
	//	*StructuredQuery_Filter_FieldFilter
	//	*StructuredQuery_Filter_UnaryFilter
	FilterType           isStructuredQuery_Filter_FilterType `protobuf_oneof:"filter_type"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *StructuredQuery_Filter) Reset()         { *m = StructuredQuery_Filter{} }
func (m *StructuredQuery_Filter) String() string { return proto.CompactTextString(m) }
func (*StructuredQuery_Filter) ProtoMessage()    {}
func (*StructuredQuery_Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_45ff6dfb507a4c5e, []int{0, 1}
}

func (m *StructuredQuery_Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StructuredQuery_Filter.Unmarshal(m, b)
}
func (m *StructuredQuery_Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StructuredQuery_Filter.Marshal(b, m, deterministic)
}
func (m *StructuredQuery_Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StructuredQuery_Filter.Merge(m, src)
}
func (m *StructuredQuery_Filter) XXX_Size() int {
	return xxx_messageInfo_StructuredQuery_Filter.Size(m)
}
func (m *StructuredQuery_Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_StructuredQuery_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_StructuredQuery_Filter proto.InternalMessageInfo

type isStructuredQuery_Filter_FilterType interface {
	isStructuredQuery_Filter_FilterType()
}

type StructuredQuery_Filter_CompositeFilter struct {
	CompositeFilter *StructuredQuery_CompositeFilter `protobuf:"bytes,1,opt,name=composite_filter,json=compositeFilter,proto3,oneof"`
}

type StructuredQuery_Filter_FieldFilter struct {
	FieldFilter *StructuredQuery_FieldFilter `protobuf:"bytes,2,opt,name=field_filter,json=fieldFilter,proto3,oneof"`
}

type StructuredQuery_Filter_UnaryFilter struct {
	UnaryFilter *StructuredQuery_UnaryFilter `protobuf:"bytes,3,opt,name=unary_filter,json=unaryFilter,proto3,oneof"`
}

func (*StructuredQuery_Filter_CompositeFilter) isStructuredQuery_Filter_FilterType() {}

func (*StructuredQuery_Filter_FieldFilter) isStructuredQuery_Filter_FilterType() {}

func (*StructuredQuery_Filter_UnaryFilter) isStructuredQuery_Filter_FilterType() {}

func (m *StructuredQuery_Filter) GetFilterType() isStructuredQuery_Filter_FilterType {
	if m != nil {
		return m.FilterType
	}
	return nil
}

func (m *StructuredQuery_Filter) GetCompositeFilter() *StructuredQuery_CompositeFilter {
	if x, ok := m.GetFilterType().(*StructuredQuery_Filter_CompositeFilter); ok {
		return x.CompositeFilter
	}
	return nil
}

func (m *StructuredQuery_Filter) GetFieldFilter() *StructuredQuery_FieldFilter {
	if x, ok := m.GetFilterType().(*StructuredQuery_Filter_FieldFilter); ok {
		return x.FieldFilter
	}
	return nil
}

func (m *StructuredQuery_Filter) GetUnaryFilter() *StructuredQuery_UnaryFilter {
	if x, ok := m.GetFilterType().(*StructuredQuery_Filter_UnaryFilter); ok {
		return x.UnaryFilter
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StructuredQuery_Filter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StructuredQuery_Filter_CompositeFilter)(nil),
		(*StructuredQuery_Filter_FieldFilter)(nil),
		(*StructuredQuery_Filter_UnaryFilter)(nil),
	}
}

// A filter that merges multiple other filters using the given operator.
type StructuredQuery_CompositeFilter struct {
	// The operator for combining multiple filters.
	Op StructuredQuery_CompositeFilter_Operator `protobuf:"varint,1,opt,name=op,proto3,enum=google.firestore.v1.StructuredQuery_CompositeFilter_Operator" json:"op,omitempty"`
	// The list of filters to combine.
	// Must contain at least one filter.
	Filters              []*StructuredQuery_Filter `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *StructuredQuery_CompositeFilter) Reset()         { *m = StructuredQuery_CompositeFilter{} }
func (m *StructuredQuery_CompositeFilter) String() string { return proto.CompactTextString(m) }
func (*StructuredQuery_CompositeFilter) ProtoMessage()    {}
func (*StructuredQuery_CompositeFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_45ff6dfb507a4c5e, []int{0, 2}
}

func (m *StructuredQuery_CompositeFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StructuredQuery_CompositeFilter.Unmarshal(m, b)
}
func (m *StructuredQuery_CompositeFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StructuredQuery_CompositeFilter.Marshal(b, m, deterministic)
}
func (m *StructuredQuery_CompositeFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StructuredQuery_CompositeFilter.Merge(m, src)
}
func (m *StructuredQuery_CompositeFilter) XXX_Size() int {
	return xxx_messageInfo_StructuredQuery_CompositeFilter.Size(m)
}
func (m *StructuredQuery_CompositeFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_StructuredQuery_CompositeFilter.DiscardUnknown(m)
}

var xxx_messageInfo_StructuredQuery_CompositeFilter proto.InternalMessageInfo

func (m *StructuredQuery_CompositeFilter) GetOp() StructuredQuery_CompositeFilter_Operator {
	if m != nil {
		return m.Op
	}
	return StructuredQuery_CompositeFilter_OPERATOR_UNSPECIFIED
}

func (m *StructuredQuery_CompositeFilter) GetFilters() []*StructuredQuery_Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

// A filter on a specific field.
type StructuredQuery_FieldFilter struct {
	// The field to filter by.
	Field *StructuredQuery_FieldReference `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	// The operator to filter by.
	Op StructuredQuery_FieldFilter_Operator `protobuf:"varint,2,opt,name=op,proto3,enum=google.firestore.v1.StructuredQuery_FieldFilter_Operator" json:"op,omitempty"`
	// The value to compare to.
	Value                *Value   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StructuredQuery_FieldFilter) Reset()         { *m = StructuredQuery_FieldFilter{} }
func (m *StructuredQuery_FieldFilter) String() string { return proto.CompactTextString(m) }
func (*StructuredQuery_FieldFilter) ProtoMessage()    {}
func (*StructuredQuery_FieldFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_45ff6dfb507a4c5e, []int{0, 3}
}

func (m *StructuredQuery_FieldFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StructuredQuery_FieldFilter.Unmarshal(m, b)
}
func (m *StructuredQuery_FieldFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StructuredQuery_FieldFilter.Marshal(b, m, deterministic)
}
func (m *StructuredQuery_FieldFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StructuredQuery_FieldFilter.Merge(m, src)
}
func (m *StructuredQuery_FieldFilter) XXX_Size() int {
	return xxx_messageInfo_StructuredQuery_FieldFilter.Size(m)
}
func (m *StructuredQuery_FieldFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_StructuredQuery_FieldFilter.DiscardUnknown(m)
}

var xxx_messageInfo_StructuredQuery_FieldFilter proto.InternalMessageInfo

func (m *StructuredQuery_FieldFilter) GetField() *StructuredQuery_FieldReference {
	if m != nil {
		return m.Field
	}
	return nil
}

func (m *StructuredQuery_FieldFilter) GetOp() StructuredQuery_FieldFilter_Operator {
	if m != nil {
		return m.Op
	}
	return StructuredQuery_FieldFilter_OPERATOR_UNSPECIFIED
}

func (m *StructuredQuery_FieldFilter) GetValue() *Value {
	if m != nil {
		return m.Value
	}
	return nil
}

// A filter with a single operand.
type StructuredQuery_UnaryFilter struct {
	// The unary operator to apply.
	Op StructuredQuery_UnaryFilter_Operator `protobuf:"varint,1,opt,name=op,proto3,enum=google.firestore.v1.StructuredQuery_UnaryFilter_Operator" json:"op,omitempty"`
	// The argument to the filter.
	//
	// Types that are valid to be assigned to OperandType:
	//	*StructuredQuery_UnaryFilter_Field
	OperandType          isStructuredQuery_UnaryFilter_OperandType `protobuf_oneof:"operand_type"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *StructuredQuery_UnaryFilter) Reset()         { *m = StructuredQuery_UnaryFilter{} }
func (m *StructuredQuery_UnaryFilter) String() string { return proto.CompactTextString(m) }
func (*StructuredQuery_UnaryFilter) ProtoMessage()    {}
func (*StructuredQuery_UnaryFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_45ff6dfb507a4c5e, []int{0, 4}
}

func (m *StructuredQuery_UnaryFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StructuredQuery_UnaryFilter.Unmarshal(m, b)
}
func (m *StructuredQuery_UnaryFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StructuredQuery_UnaryFilter.Marshal(b, m, deterministic)
}
func (m *StructuredQuery_UnaryFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StructuredQuery_UnaryFilter.Merge(m, src)
}
func (m *StructuredQuery_UnaryFilter) XXX_Size() int {
	return xxx_messageInfo_StructuredQuery_UnaryFilter.Size(m)
}
func (m *StructuredQuery_UnaryFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_StructuredQuery_UnaryFilter.DiscardUnknown(m)
}

var xxx_messageInfo_StructuredQuery_UnaryFilter proto.InternalMessageInfo

func (m *StructuredQuery_UnaryFilter) GetOp() StructuredQuery_UnaryFilter_Operator {
	if m != nil {
		return m.Op
	}
	return StructuredQuery_UnaryFilter_OPERATOR_UNSPECIFIED
}

type isStructuredQuery_UnaryFilter_OperandType interface {
	isStructuredQuery_UnaryFilter_OperandType()
}

type StructuredQuery_UnaryFilter_Field struct {
	Field *StructuredQuery_FieldReference `protobuf:"bytes,2,opt,name=field,proto3,oneof"`
}

func (*StructuredQuery_UnaryFilter_Field) isStructuredQuery_UnaryFilter_OperandType() {}

func (m *StructuredQuery_UnaryFilter) GetOperandType() isStructuredQuery_UnaryFilter_OperandType {
	if m != nil {
		return m.OperandType
	}
	return nil
}

func (m *StructuredQuery_UnaryFilter) GetField() *StructuredQuery_FieldReference {
	if x, ok := m.GetOperandType().(*StructuredQuery_UnaryFilter_Field); ok {
		return x.Field
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StructuredQuery_UnaryFilter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StructuredQuery_UnaryFilter_Field)(nil),
	}
}

// An order on a field.
type StructuredQuery_Order struct {
	// The field to order by.
	Field *StructuredQuery_FieldReference `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	// The direction to order by. Defaults to `ASCENDING`.
	Direction            StructuredQuery_Direction `protobuf:"varint,2,opt,name=direction,proto3,enum=google.firestore.v1.StructuredQuery_Direction" json:"direction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *StructuredQuery_Order) Reset()         { *m = StructuredQuery_Order{} }
func (m *StructuredQuery_Order) String() string { return proto.CompactTextString(m) }
func (*StructuredQuery_Order) ProtoMessage()    {}
func (*StructuredQuery_Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_45ff6dfb507a4c5e, []int{0, 5}
}

func (m *StructuredQuery_Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StructuredQuery_Order.Unmarshal(m, b)
}
func (m *StructuredQuery_Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StructuredQuery_Order.Marshal(b, m, deterministic)
}
func (m *StructuredQuery_Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StructuredQuery_Order.Merge(m, src)
}
func (m *StructuredQuery_Order) XXX_Size() int {
	return xxx_messageInfo_StructuredQuery_Order.Size(m)
}
func (m *StructuredQuery_Order) XXX_DiscardUnknown() {
	xxx_messageInfo_StructuredQuery_Order.DiscardUnknown(m)
}

var xxx_messageInfo_StructuredQuery_Order proto.InternalMessageInfo

func (m *StructuredQuery_Order) GetField() *StructuredQuery_FieldReference {
	if m != nil {
		return m.Field
	}
	return nil
}

func (m *StructuredQuery_Order) GetDirection() StructuredQuery_Direction {
	if m != nil {
		return m.Direction
	}
	return StructuredQuery_DIRECTION_UNSPECIFIED
}

// A reference to a field, such as `max(messages.time) as max_time`.
type StructuredQuery_FieldReference struct {
	FieldPath            string   `protobuf:"bytes,2,opt,name=field_path,json=fieldPath,proto3" json:"field_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StructuredQuery_FieldReference) Reset()         { *m = StructuredQuery_FieldReference{} }
func (m *StructuredQuery_FieldReference) String() string { return proto.CompactTextString(m) }
func (*StructuredQuery_FieldReference) ProtoMessage()    {}
func (*StructuredQuery_FieldReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_45ff6dfb507a4c5e, []int{0, 6}
}

func (m *StructuredQuery_FieldReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StructuredQuery_FieldReference.Unmarshal(m, b)
}
func (m *StructuredQuery_FieldReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StructuredQuery_FieldReference.Marshal(b, m, deterministic)
}
func (m *StructuredQuery_FieldReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StructuredQuery_FieldReference.Merge(m, src)
}
func (m *StructuredQuery_FieldReference) XXX_Size() int {
	return xxx_messageInfo_StructuredQuery_FieldReference.Size(m)
}
func (m *StructuredQuery_FieldReference) XXX_DiscardUnknown() {
	xxx_messageInfo_StructuredQuery_FieldReference.DiscardUnknown(m)
}

var xxx_messageInfo_StructuredQuery_FieldReference proto.InternalMessageInfo

func (m *StructuredQuery_FieldReference) GetFieldPath() string {
	if m != nil {
		return m.FieldPath
	}
	return ""
}

// The projection of document's fields to return.
type StructuredQuery_Projection struct {
	// The fields to return.
	//
	// If empty, all fields are returned. To only return the name
	// of the document, use `['__name__']`.
	Fields               []*StructuredQuery_FieldReference `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *StructuredQuery_Projection) Reset()         { *m = StructuredQuery_Projection{} }
func (m *StructuredQuery_Projection) String() string { return proto.CompactTextString(m) }
func (*StructuredQuery_Projection) ProtoMessage()    {}
func (*StructuredQuery_Projection) Descriptor() ([]byte, []int) {
	return fileDescriptor_45ff6dfb507a4c5e, []int{0, 7}
}

func (m *StructuredQuery_Projection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StructuredQuery_Projection.Unmarshal(m, b)
}
func (m *StructuredQuery_Projection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StructuredQuery_Projection.Marshal(b, m, deterministic)
}
func (m *StructuredQuery_Projection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StructuredQuery_Projection.Merge(m, src)
}
func (m *StructuredQuery_Projection) XXX_Size() int {
	return xxx_messageInfo_StructuredQuery_Projection.Size(m)
}
func (m *StructuredQuery_Projection) XXX_DiscardUnknown() {
	xxx_messageInfo_StructuredQuery_Projection.DiscardUnknown(m)
}

var xxx_messageInfo_StructuredQuery_Projection proto.InternalMessageInfo

func (m *StructuredQuery_Projection) GetFields() []*StructuredQuery_FieldReference {
	if m != nil {
		return m.Fields
	}
	return nil
}

// A position in a query result set.
type Cursor struct {
	// The values that represent a position, in the order they appear in
	// the order by clause of a query.
	//
	// Can contain fewer values than specified in the order by clause.
	Values []*Value `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	// If the position is just before or just after the given values, relative
	// to the sort order defined by the query.
	Before               bool     `protobuf:"varint,2,opt,name=before,proto3" json:"before,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cursor) Reset()         { *m = Cursor{} }
func (m *Cursor) String() string { return proto.CompactTextString(m) }
func (*Cursor) ProtoMessage()    {}
func (*Cursor) Descriptor() ([]byte, []int) {
	return fileDescriptor_45ff6dfb507a4c5e, []int{1}
}

func (m *Cursor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cursor.Unmarshal(m, b)
}
func (m *Cursor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cursor.Marshal(b, m, deterministic)
}
func (m *Cursor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cursor.Merge(m, src)
}
func (m *Cursor) XXX_Size() int {
	return xxx_messageInfo_Cursor.Size(m)
}
func (m *Cursor) XXX_DiscardUnknown() {
	xxx_messageInfo_Cursor.DiscardUnknown(m)
}

var xxx_messageInfo_Cursor proto.InternalMessageInfo

func (m *Cursor) GetValues() []*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Cursor) GetBefore() bool {
	if m != nil {
		return m.Before
	}
	return false
}

func init() {
	proto.RegisterEnum("google.firestore.v1.StructuredQuery_Direction", StructuredQuery_Direction_name, StructuredQuery_Direction_value)
	proto.RegisterEnum("google.firestore.v1.StructuredQuery_CompositeFilter_Operator", StructuredQuery_CompositeFilter_Operator_name, StructuredQuery_CompositeFilter_Operator_value)
	proto.RegisterEnum("google.firestore.v1.StructuredQuery_FieldFilter_Operator", StructuredQuery_FieldFilter_Operator_name, StructuredQuery_FieldFilter_Operator_value)
	proto.RegisterEnum("google.firestore.v1.StructuredQuery_UnaryFilter_Operator", StructuredQuery_UnaryFilter_Operator_name, StructuredQuery_UnaryFilter_Operator_value)
	proto.RegisterType((*StructuredQuery)(nil), "google.firestore.v1.StructuredQuery")
	proto.RegisterType((*StructuredQuery_CollectionSelector)(nil), "google.firestore.v1.StructuredQuery.CollectionSelector")
	proto.RegisterType((*StructuredQuery_Filter)(nil), "google.firestore.v1.StructuredQuery.Filter")
	proto.RegisterType((*StructuredQuery_CompositeFilter)(nil), "google.firestore.v1.StructuredQuery.CompositeFilter")
	proto.RegisterType((*StructuredQuery_FieldFilter)(nil), "google.firestore.v1.StructuredQuery.FieldFilter")
	proto.RegisterType((*StructuredQuery_UnaryFilter)(nil), "google.firestore.v1.StructuredQuery.UnaryFilter")
	proto.RegisterType((*StructuredQuery_Order)(nil), "google.firestore.v1.StructuredQuery.Order")
	proto.RegisterType((*StructuredQuery_FieldReference)(nil), "google.firestore.v1.StructuredQuery.FieldReference")
	proto.RegisterType((*StructuredQuery_Projection)(nil), "google.firestore.v1.StructuredQuery.Projection")
	proto.RegisterType((*Cursor)(nil), "google.firestore.v1.Cursor")
}

func init() { proto.RegisterFile("google/firestore/v1/query.proto", fileDescriptor_45ff6dfb507a4c5e) }

var fileDescriptor_45ff6dfb507a4c5e = []byte{
	// 976 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x6e, 0xe3, 0xc4,
	0x17, 0xaf, 0x9d, 0xe6, 0xeb, 0xa4, 0x4d, 0xad, 0xf9, 0xff, 0x59, 0xb2, 0xe6, 0xab, 0x0a, 0x17,
	0x54, 0x20, 0x9c, 0x6d, 0x8b, 0x40, 0x08, 0xf6, 0xc2, 0x4d, 0xdc, 0xd6, 0x6a, 0x71, 0xb2, 0x93,
	0xb4, 0xd2, 0xa2, 0x4a, 0x96, 0x6b, 0x8f, 0x53, 0x23, 0xd7, 0x63, 0xc6, 0xe3, 0xae, 0xfa, 0x24,
	0xdc, 0x22, 0xc4, 0x05, 0xe2, 0x86, 0xf7, 0xe0, 0x01, 0x78, 0x0d, 0x6e, 0x78, 0x00, 0xe4, 0xf1,
	0x24, 0x69, 0x76, 0x2b, 0x48, 0x2b, 0xee, 0x72, 0x3e, 0x7e, 0x3f, 0x9f, 0xf9, 0x9d, 0x73, 0x66,
	0x02, 0x1f, 0x4c, 0x29, 0x9d, 0xc6, 0xa4, 0x17, 0x46, 0x8c, 0x64, 0x9c, 0x32, 0xd2, 0xbb, 0xd9,
	0xed, 0x7d, 0x9f, 0x13, 0x76, 0x6b, 0xa4, 0x8c, 0x72, 0x8a, 0xfe, 0x57, 0x26, 0x18, 0xf3, 0x04,
	0xe3, 0x66, 0x57, 0xef, 0xde, 0x87, 0x0a, 0xa8, 0x9f, 0x5f, 0x93, 0x84, 0x97, 0x40, 0xfd, 0x7d,
	0x99, 0x23, 0xac, 0xcb, 0x3c, 0xec, 0xbd, 0x62, 0x5e, 0x9a, 0x12, 0x96, 0xc9, 0xf8, 0xbb, 0x32,
	0xee, 0xa5, 0x51, 0xcf, 0x4b, 0x12, 0xca, 0x3d, 0x1e, 0xd1, 0x44, 0x46, 0xbb, 0xbf, 0x69, 0xb0,
	0x35, 0xe6, 0x2c, 0xf7, 0x79, 0xce, 0x48, 0xf0, 0xa2, 0x28, 0x08, 0x1d, 0x41, 0x2d, 0x23, 0x31,
	0xf1, 0x79, 0x47, 0xd9, 0x56, 0x76, 0x5a, 0x7b, 0x3d, 0xe3, 0x9e, 0xda, 0x8c, 0xd7, 0x50, 0xc6,
	0x88, 0xd1, 0xef, 0x88, 0x5f, 0x70, 0x63, 0x09, 0x47, 0x27, 0xb0, 0x1e, 0x32, 0x7a, 0xdd, 0x51,
	0xb7, 0x2b, 0x3b, 0xad, 0xbd, 0x2f, 0x56, 0xa2, 0xe9, 0xd3, 0x38, 0x2e, 0x69, 0xc6, 0x82, 0x84,
	0x32, 0x2c, 0x48, 0x90, 0x09, 0xd5, 0x57, 0x57, 0x84, 0x91, 0x4e, 0x45, 0x14, 0xf5, 0xc9, 0x4a,
	0x6c, 0x87, 0x51, 0xcc, 0x09, 0xc3, 0x25, 0x12, 0x59, 0xd0, 0xa0, 0x2c, 0x20, 0xcc, 0xbd, 0xbc,
	0xed, 0xac, 0x8b, 0x9a, 0x3e, 0x5e, 0x89, 0x65, 0x58, 0x80, 0x70, 0x5d, 0x60, 0x0f, 0x6e, 0xd1,
	0xe7, 0xd0, 0xc8, 0xb8, 0xc7, 0xb8, 0xeb, 0xf1, 0x4e, 0x5d, 0x14, 0xf3, 0xce, 0xbd, 0x34, 0xfd,
	0x9c, 0x65, 0x94, 0xe1, 0xba, 0x48, 0x36, 0x39, 0xda, 0x83, 0x1a, 0x49, 0x82, 0x02, 0xd5, 0xf8,
	0x77, 0x54, 0x95, 0x24, 0x81, 0xc9, 0xd1, 0x13, 0xa8, 0xd1, 0x30, 0xcc, 0x08, 0xef, 0xd4, 0xb6,
	0x95, 0x9d, 0x2a, 0x96, 0x16, 0xda, 0x85, 0x6a, 0x1c, 0x5d, 0x47, 0xbc, 0x53, 0x5d, 0xa6, 0x9a,
	0x4d, 0x81, 0x61, 0x27, 0x7c, 0x7f, 0xef, 0xdc, 0x8b, 0x73, 0x82, 0xcb, 0x4c, 0xfd, 0x12, 0xd0,
	0x9b, 0xe2, 0xa2, 0x0f, 0x61, 0xd3, 0x9f, 0x7b, 0xdd, 0x28, 0xe8, 0xa8, 0xdb, 0xca, 0x4e, 0x13,
	0x6f, 0x2c, 0x9c, 0x76, 0x80, 0x3e, 0x82, 0x2d, 0x2f, 0x8e, 0xdd, 0x80, 0x64, 0x3e, 0x49, 0x02,
	0x2f, 0xe1, 0x99, 0xe8, 0x42, 0x03, 0xb7, 0xbd, 0x38, 0x1e, 0x2c, 0xbc, 0xfa, 0x2f, 0x2a, 0xd4,
	0x4a, 0xcd, 0x91, 0x07, 0x9a, 0x4f, 0xaf, 0x53, 0x9a, 0x45, 0x9c, 0xb8, 0xa1, 0xf0, 0xc9, 0x79,
	0xfa, 0x6c, 0xc5, 0x41, 0x90, 0xe0, 0x92, 0xef, 0x78, 0x0d, 0x6f, 0xf9, 0xcb, 0x2e, 0x74, 0x06,
	0x1b, 0x61, 0x44, 0xe2, 0x60, 0x46, 0xaf, 0x0a, 0xfa, 0x67, 0x2b, 0x4e, 0x06, 0x89, 0x83, 0x39,
	0x75, 0x2b, 0x5c, 0x98, 0x05, 0x6d, 0x9e, 0x78, 0xec, 0x76, 0x46, 0x5b, 0x79, 0x00, 0xed, 0x59,
	0x01, 0x5c, 0xd0, 0xe6, 0x0b, 0xf3, 0x60, 0x13, 0x5a, 0x25, 0xa1, 0xcb, 0x6f, 0x53, 0xa2, 0xff,
	0xa1, 0xc0, 0xd6, 0x6b, 0x67, 0x44, 0xdf, 0x80, 0x4a, 0x53, 0xa1, 0x52, 0x7b, 0xef, 0xf9, 0x63,
	0x54, 0x32, 0x86, 0x29, 0x61, 0x5e, 0xb1, 0x34, 0x2a, 0x4d, 0x91, 0x05, 0xf5, 0xf2, 0x8b, 0x99,
	0x5c, 0xc1, 0x07, 0x2d, 0xcd, 0x0c, 0xdb, 0xfd, 0x14, 0x1a, 0x33, 0x5a, 0xd4, 0x81, 0xff, 0x0f,
	0x47, 0x16, 0x36, 0x27, 0x43, 0xec, 0x9e, 0x39, 0xe3, 0x91, 0xd5, 0xb7, 0x0f, 0x6d, 0x6b, 0xa0,
	0xad, 0xa1, 0x3a, 0x54, 0x4c, 0x67, 0xa0, 0x29, 0xfa, 0x9f, 0x2a, 0xb4, 0xee, 0xa8, 0x8b, 0x6c,
	0xa8, 0x0a, 0x75, 0x65, 0xf7, 0xf7, 0x57, 0x6f, 0x0f, 0x26, 0x21, 0x61, 0x24, 0xf1, 0x09, 0x2e,
	0x19, 0x90, 0x2d, 0xf4, 0x51, 0x85, 0x3e, 0x5f, 0x3e, 0xb4, 0xcd, 0xcb, 0xda, 0x3c, 0x83, 0xea,
	0x4d, 0xb1, 0x1d, 0xb2, 0xbb, 0xfa, 0xbd, 0x6c, 0x72, 0x7f, 0x44, 0x62, 0xf7, 0x07, 0x65, 0x25,
	0x1d, 0x36, 0xa1, 0x79, 0x6a, 0x8d, 0xc7, 0xee, 0xe4, 0xd8, 0x74, 0x34, 0x05, 0x3d, 0x01, 0x34,
	0x37, 0xdd, 0x21, 0x76, 0xad, 0x17, 0x67, 0xe6, 0xa9, 0xa6, 0x22, 0x0d, 0x36, 0x8e, 0xb0, 0x65,
	0x4e, 0x2c, 0x5c, 0x66, 0x56, 0xd0, 0x53, 0x78, 0xeb, 0xae, 0x67, 0x91, 0xbc, 0x8e, 0x9a, 0x50,
	0x2d, 0x7f, 0x56, 0x11, 0x82, 0xb6, 0x89, 0xb1, 0xf9, 0xd2, 0xed, 0x0f, 0x9d, 0x89, 0x69, 0x3b,
	0x63, 0xad, 0xae, 0xff, 0xa5, 0x40, 0xeb, 0xce, 0xe0, 0x49, 0x99, 0x94, 0x07, 0xc8, 0x74, 0x07,
	0xbd, 0x2c, 0xd3, 0xc9, 0xac, 0x79, 0xea, 0xa3, 0x9b, 0x77, 0xbc, 0x26, 0xdb, 0xd7, 0x7d, 0xbe,
	0x92, 0x80, 0x00, 0x35, 0x7b, 0xec, 0x3a, 0xa6, 0xa3, 0xa9, 0xa8, 0x05, 0xf5, 0xe2, 0xf7, 0xd9,
	0xe9, 0xa9, 0x56, 0x39, 0x68, 0xc3, 0x06, 0x2d, 0xe0, 0x49, 0x50, 0x6e, 0xd0, 0x8f, 0x0a, 0x54,
	0xc5, 0xd5, 0xfc, 0x5f, 0x8e, 0xd8, 0x29, 0x34, 0x83, 0x88, 0x95, 0x37, 0x9f, 0x9c, 0x34, 0x63,
	0x25, 0xba, 0xc1, 0x0c, 0x85, 0x17, 0x04, 0x7a, 0x0f, 0xda, 0xcb, 0x9f, 0x41, 0xef, 0x01, 0x94,
	0x77, 0x56, 0xea, 0xf1, 0x2b, 0x79, 0xd9, 0x36, 0x85, 0x67, 0xe4, 0xf1, 0x2b, 0xfd, 0x25, 0xc0,
	0xe2, 0x21, 0x45, 0x27, 0x50, 0x13, 0xa1, 0xd9, 0xfe, 0x3e, 0xea, 0x60, 0x92, 0xa2, 0x6b, 0x41,
	0x73, 0x5e, 0x63, 0x31, 0x6c, 0x03, 0x1b, 0x5b, 0xfd, 0x89, 0x3d, 0x74, 0xde, 0x1c, 0x60, 0x73,
	0xdc, 0xb7, 0x9c, 0x81, 0xed, 0x1c, 0x69, 0x0a, 0x6a, 0x03, 0x0c, 0xac, 0xb9, 0xad, 0x76, 0x27,
	0x50, 0x2b, 0x9f, 0xa8, 0xe2, 0x3d, 0x13, 0x9b, 0x91, 0x75, 0x14, 0x51, 0xdd, 0x3f, 0xed, 0x90,
	0xcc, 0x2c, 0xde, 0xb3, 0x4b, 0x12, 0x52, 0x46, 0xc4, 0xd1, 0x1b, 0x58, 0x5a, 0x07, 0x3f, 0x2b,
	0xf0, 0xb6, 0x4f, 0xaf, 0xef, 0x63, 0x38, 0x00, 0x71, 0xac, 0x51, 0xf1, 0xb2, 0x8d, 0x94, 0x6f,
	0xbf, 0x96, 0x29, 0x53, 0x1a, 0x7b, 0xc9, 0xd4, 0xa0, 0x6c, 0xda, 0x9b, 0x92, 0x44, 0xbc, 0x7b,
	0xbd, 0x32, 0xe4, 0xa5, 0x51, 0xb6, 0xf4, 0x97, 0xe9, 0xab, 0xb9, 0xf1, 0x93, 0xba, 0x7e, 0xd4,
	0x3f, 0x1c, 0xff, 0xaa, 0x3e, 0x3d, 0x2a, 0x59, 0xfa, 0x31, 0xcd, 0x03, 0xe3, 0x70, 0xfe, 0xb9,
	0xf3, 0xdd, 0xdf, 0x67, 0xb1, 0x0b, 0x11, 0xbb, 0x98, 0xc7, 0x2e, 0xce, 0x77, 0x2f, 0x6b, 0xe2,
	0x3b, 0xfb, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x1c, 0x23, 0x29, 0xcf, 0x09, 0x00, 0x00,
}

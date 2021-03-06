// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: google/bigtable/admin/table/v1/bigtable_table_data.proto

package table

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	longrunning "google.golang.org/genproto/googleapis/longrunning"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Table_TimestampGranularity int32

const (
	Table_MILLIS Table_TimestampGranularity = 0
)

// Enum value maps for Table_TimestampGranularity.
var (
	Table_TimestampGranularity_name = map[int32]string{
		0: "MILLIS",
	}
	Table_TimestampGranularity_value = map[string]int32{
		"MILLIS": 0,
	}
)

func (x Table_TimestampGranularity) Enum() *Table_TimestampGranularity {
	p := new(Table_TimestampGranularity)
	*p = x
	return p
}

func (x Table_TimestampGranularity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Table_TimestampGranularity) Descriptor() protoreflect.EnumDescriptor {
	return file_google_bigtable_admin_table_v1_bigtable_table_data_proto_enumTypes[0].Descriptor()
}

func (Table_TimestampGranularity) Type() protoreflect.EnumType {
	return &file_google_bigtable_admin_table_v1_bigtable_table_data_proto_enumTypes[0]
}

func (x Table_TimestampGranularity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Table_TimestampGranularity.Descriptor instead.
func (Table_TimestampGranularity) EnumDescriptor() ([]byte, []int) {
	return file_google_bigtable_admin_table_v1_bigtable_table_data_proto_rawDescGZIP(), []int{0, 0}
}

// A collection of user data indexed by row, column, and timestamp.
// Each table is served using the resources of its parent cluster.
type Table struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique identifier of the form
	// <cluster_name>/tables/[_a-zA-Z0-9][-_.a-zA-Z0-9]*
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// If this Table is in the process of being created, the Operation used to
	// track its progress. As long as this operation is present, the Table will
	// not accept any Table Admin or Read/Write requests.
	CurrentOperation *longrunning.Operation `protobuf:"bytes,2,opt,name=current_operation,json=currentOperation,proto3" json:"current_operation,omitempty"`
	// The column families configured for this table, mapped by column family id.
	ColumnFamilies map[string]*ColumnFamily `protobuf:"bytes,3,rep,name=column_families,json=columnFamilies,proto3" json:"column_families,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The granularity (e.g. MILLIS, MICROS) at which timestamps are stored in
	// this table. Timestamps not matching the granularity will be rejected.
	// Cannot be changed once the table is created.
	Granularity Table_TimestampGranularity `protobuf:"varint,4,opt,name=granularity,proto3,enum=google.bigtable.admin.table.v1.Table_TimestampGranularity" json:"granularity,omitempty"`
}

func (x *Table) Reset() {
	*x = Table{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_bigtable_admin_table_v1_bigtable_table_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Table) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Table) ProtoMessage() {}

func (x *Table) ProtoReflect() protoreflect.Message {
	mi := &file_google_bigtable_admin_table_v1_bigtable_table_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Table.ProtoReflect.Descriptor instead.
func (*Table) Descriptor() ([]byte, []int) {
	return file_google_bigtable_admin_table_v1_bigtable_table_data_proto_rawDescGZIP(), []int{0}
}

func (x *Table) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Table) GetCurrentOperation() *longrunning.Operation {
	if x != nil {
		return x.CurrentOperation
	}
	return nil
}

func (x *Table) GetColumnFamilies() map[string]*ColumnFamily {
	if x != nil {
		return x.ColumnFamilies
	}
	return nil
}

func (x *Table) GetGranularity() Table_TimestampGranularity {
	if x != nil {
		return x.Granularity
	}
	return Table_MILLIS
}

// A set of columns within a table which share a common configuration.
type ColumnFamily struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique identifier of the form <table_name>/columnFamilies/[-_.a-zA-Z0-9]+
	// The last segment is the same as the "name" field in
	// google.bigtable.v1.Family.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Garbage collection expression specified by the following grammar:
	//   GC = EXPR
	//      | "" ;
	//   EXPR = EXPR, "||", EXPR              (* lowest precedence *)
	//        | EXPR, "&&", EXPR
	//        | "(", EXPR, ")"                (* highest precedence *)
	//        | PROP ;
	//   PROP = "version() >", NUM32
	//        | "age() >", NUM64, [ UNIT ] ;
	//   NUM32 = non-zero-digit { digit } ;    (* # NUM32 <= 2^32 - 1 *)
	//   NUM64 = non-zero-digit { digit } ;    (* # NUM64 <= 2^63 - 1 *)
	//   UNIT =  "d" | "h" | "m"  (* d=days, h=hours, m=minutes, else micros *)
	// GC expressions can be up to 500 characters in length
	//
	// The different types of PROP are defined as follows:
	//   version() - cell index, counting from most recent and starting at 1
	//   age() - age of the cell (current time minus cell timestamp)
	//
	// Example: "version() > 3 || (age() > 3d && version() > 1)"
	//   drop cells beyond the most recent three, and drop cells older than three
	//   days unless they're the most recent cell in the row/column
	//
	// Garbage collection executes opportunistically in the background, and so
	// it's possible for reads to return a cell even if it matches the active GC
	// expression for its family.
	GcExpression string `protobuf:"bytes,2,opt,name=gc_expression,json=gcExpression,proto3" json:"gc_expression,omitempty"`
	// Garbage collection rule specified as a protobuf.
	// Supersedes `gc_expression`.
	// Must serialize to at most 500 bytes.
	//
	// NOTE: Garbage collection executes opportunistically in the background, and
	// so it's possible for reads to return a cell even if it matches the active
	// GC expression for its family.
	GcRule *GcRule `protobuf:"bytes,3,opt,name=gc_rule,json=gcRule,proto3" json:"gc_rule,omitempty"`
}

func (x *ColumnFamily) Reset() {
	*x = ColumnFamily{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_bigtable_admin_table_v1_bigtable_table_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ColumnFamily) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColumnFamily) ProtoMessage() {}

func (x *ColumnFamily) ProtoReflect() protoreflect.Message {
	mi := &file_google_bigtable_admin_table_v1_bigtable_table_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColumnFamily.ProtoReflect.Descriptor instead.
func (*ColumnFamily) Descriptor() ([]byte, []int) {
	return file_google_bigtable_admin_table_v1_bigtable_table_data_proto_rawDescGZIP(), []int{1}
}

func (x *ColumnFamily) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ColumnFamily) GetGcExpression() string {
	if x != nil {
		return x.GcExpression
	}
	return ""
}

func (x *ColumnFamily) GetGcRule() *GcRule {
	if x != nil {
		return x.GcRule
	}
	return nil
}

// Rule for determining which cells to delete during garbage collection.
type GcRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Rule:
	//	*GcRule_MaxNumVersions
	//	*GcRule_MaxAge
	//	*GcRule_Intersection_
	//	*GcRule_Union_
	Rule isGcRule_Rule `protobuf_oneof:"rule"`
}

func (x *GcRule) Reset() {
	*x = GcRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_bigtable_admin_table_v1_bigtable_table_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcRule) ProtoMessage() {}

func (x *GcRule) ProtoReflect() protoreflect.Message {
	mi := &file_google_bigtable_admin_table_v1_bigtable_table_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcRule.ProtoReflect.Descriptor instead.
func (*GcRule) Descriptor() ([]byte, []int) {
	return file_google_bigtable_admin_table_v1_bigtable_table_data_proto_rawDescGZIP(), []int{2}
}

func (m *GcRule) GetRule() isGcRule_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (x *GcRule) GetMaxNumVersions() int32 {
	if x, ok := x.GetRule().(*GcRule_MaxNumVersions); ok {
		return x.MaxNumVersions
	}
	return 0
}

func (x *GcRule) GetMaxAge() *durationpb.Duration {
	if x, ok := x.GetRule().(*GcRule_MaxAge); ok {
		return x.MaxAge
	}
	return nil
}

func (x *GcRule) GetIntersection() *GcRule_Intersection {
	if x, ok := x.GetRule().(*GcRule_Intersection_); ok {
		return x.Intersection
	}
	return nil
}

func (x *GcRule) GetUnion() *GcRule_Union {
	if x, ok := x.GetRule().(*GcRule_Union_); ok {
		return x.Union
	}
	return nil
}

type isGcRule_Rule interface {
	isGcRule_Rule()
}

type GcRule_MaxNumVersions struct {
	// Delete all cells in a column except the most recent N.
	MaxNumVersions int32 `protobuf:"varint,1,opt,name=max_num_versions,json=maxNumVersions,proto3,oneof"`
}

type GcRule_MaxAge struct {
	// Delete cells in a column older than the given age.
	// Values must be at least one millisecond, and will be truncated to
	// microsecond granularity.
	MaxAge *durationpb.Duration `protobuf:"bytes,2,opt,name=max_age,json=maxAge,proto3,oneof"`
}

type GcRule_Intersection_ struct {
	// Delete cells that would be deleted by every nested rule.
	Intersection *GcRule_Intersection `protobuf:"bytes,3,opt,name=intersection,proto3,oneof"`
}

type GcRule_Union_ struct {
	// Delete cells that would be deleted by any nested rule.
	Union *GcRule_Union `protobuf:"bytes,4,opt,name=union,proto3,oneof"`
}

func (*GcRule_MaxNumVersions) isGcRule_Rule() {}

func (*GcRule_MaxAge) isGcRule_Rule() {}

func (*GcRule_Intersection_) isGcRule_Rule() {}

func (*GcRule_Union_) isGcRule_Rule() {}

// A GcRule which deletes cells matching all of the given rules.
type GcRule_Intersection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Only delete cells which would be deleted by every element of `rules`.
	Rules []*GcRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *GcRule_Intersection) Reset() {
	*x = GcRule_Intersection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_bigtable_admin_table_v1_bigtable_table_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcRule_Intersection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcRule_Intersection) ProtoMessage() {}

func (x *GcRule_Intersection) ProtoReflect() protoreflect.Message {
	mi := &file_google_bigtable_admin_table_v1_bigtable_table_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcRule_Intersection.ProtoReflect.Descriptor instead.
func (*GcRule_Intersection) Descriptor() ([]byte, []int) {
	return file_google_bigtable_admin_table_v1_bigtable_table_data_proto_rawDescGZIP(), []int{2, 0}
}

func (x *GcRule_Intersection) GetRules() []*GcRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

// A GcRule which deletes cells matching any of the given rules.
type GcRule_Union struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Delete cells which would be deleted by any element of `rules`.
	Rules []*GcRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
}

func (x *GcRule_Union) Reset() {
	*x = GcRule_Union{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_bigtable_admin_table_v1_bigtable_table_data_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcRule_Union) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcRule_Union) ProtoMessage() {}

func (x *GcRule_Union) ProtoReflect() protoreflect.Message {
	mi := &file_google_bigtable_admin_table_v1_bigtable_table_data_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcRule_Union.ProtoReflect.Descriptor instead.
func (*GcRule_Union) Descriptor() ([]byte, []int) {
	return file_google_bigtable_admin_table_v1_bigtable_table_data_proto_rawDescGZIP(), []int{2, 1}
}

func (x *GcRule_Union) GetRules() []*GcRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

var File_google_bigtable_admin_table_v1_bigtable_table_data_proto protoreflect.FileDescriptor

var file_google_bigtable_admin_table_v1_bigtable_table_data_proto_rawDesc = []byte{
	0x0a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xbe, 0x03, 0x0a, 0x05, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4a, 0x0a,
	0x11, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x6c, 0x6f, 0x6e, 0x67, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x62, 0x0a, 0x0f, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x5f, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x46, 0x61, 0x6d, 0x69, 0x6c, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x69, 0x65, 0x73, 0x12, 0x5c, 0x0a,
	0x0b, 0x67, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x47, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x52, 0x0b,
	0x67, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61, 0x72, 0x69, 0x74, 0x79, 0x1a, 0x6f, 0x0a, 0x13, 0x43,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x42, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x46, 0x61, 0x6d, 0x69, 0x6c,
	0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x22, 0x0a, 0x14,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x47, 0x72, 0x61, 0x6e, 0x75, 0x6c, 0x61,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x53, 0x10, 0x00,
	0x22, 0x88, 0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x46, 0x61, 0x6d, 0x69, 0x6c,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x67, 0x63, 0x5f, 0x65, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x67, 0x63,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x07, 0x67, 0x63,
	0x5f, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x63, 0x52,
	0x75, 0x6c, 0x65, 0x52, 0x06, 0x67, 0x63, 0x52, 0x75, 0x6c, 0x65, 0x22, 0xa8, 0x03, 0x0a, 0x06,
	0x47, 0x63, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x5f, 0x6e, 0x75,
	0x6d, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x00, 0x52, 0x0e, 0x6d, 0x61, 0x78, 0x4e, 0x75, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x06, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x12, 0x59, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x63, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x63, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x55, 0x6e, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x05, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x1a, 0x4c, 0x0a, 0x0c, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x05, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x63, 0x52, 0x75, 0x6c, 0x65,
	0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x1a, 0x45, 0x0a, 0x05, 0x55, 0x6e, 0x69, 0x6f, 0x6e,
	0x12, 0x3c, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x63, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x42, 0x06,
	0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x42, 0x83, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x62, 0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x16, 0x42,
	0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x62,
	0x69, 0x67, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_bigtable_admin_table_v1_bigtable_table_data_proto_rawDescOnce sync.Once
	file_google_bigtable_admin_table_v1_bigtable_table_data_proto_rawDescData = file_google_bigtable_admin_table_v1_bigtable_table_data_proto_rawDesc
)

func file_google_bigtable_admin_table_v1_bigtable_table_data_proto_rawDescGZIP() []byte {
	file_google_bigtable_admin_table_v1_bigtable_table_data_proto_rawDescOnce.Do(func() {
		file_google_bigtable_admin_table_v1_bigtable_table_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_bigtable_admin_table_v1_bigtable_table_data_proto_rawDescData)
	})
	return file_google_bigtable_admin_table_v1_bigtable_table_data_proto_rawDescData
}

var file_google_bigtable_admin_table_v1_bigtable_table_data_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_bigtable_admin_table_v1_bigtable_table_data_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_bigtable_admin_table_v1_bigtable_table_data_proto_goTypes = []interface{}{
	(Table_TimestampGranularity)(0), // 0: google.bigtable.admin.table.v1.Table.TimestampGranularity
	(*Table)(nil),                   // 1: google.bigtable.admin.table.v1.Table
	(*ColumnFamily)(nil),            // 2: google.bigtable.admin.table.v1.ColumnFamily
	(*GcRule)(nil),                  // 3: google.bigtable.admin.table.v1.GcRule
	nil,                             // 4: google.bigtable.admin.table.v1.Table.ColumnFamiliesEntry
	(*GcRule_Intersection)(nil),     // 5: google.bigtable.admin.table.v1.GcRule.Intersection
	(*GcRule_Union)(nil),            // 6: google.bigtable.admin.table.v1.GcRule.Union
	(*longrunning.Operation)(nil),   // 7: google.longrunning.Operation
	(*durationpb.Duration)(nil),     // 8: google.protobuf.Duration
}
var file_google_bigtable_admin_table_v1_bigtable_table_data_proto_depIdxs = []int32{
	7,  // 0: google.bigtable.admin.table.v1.Table.current_operation:type_name -> google.longrunning.Operation
	4,  // 1: google.bigtable.admin.table.v1.Table.column_families:type_name -> google.bigtable.admin.table.v1.Table.ColumnFamiliesEntry
	0,  // 2: google.bigtable.admin.table.v1.Table.granularity:type_name -> google.bigtable.admin.table.v1.Table.TimestampGranularity
	3,  // 3: google.bigtable.admin.table.v1.ColumnFamily.gc_rule:type_name -> google.bigtable.admin.table.v1.GcRule
	8,  // 4: google.bigtable.admin.table.v1.GcRule.max_age:type_name -> google.protobuf.Duration
	5,  // 5: google.bigtable.admin.table.v1.GcRule.intersection:type_name -> google.bigtable.admin.table.v1.GcRule.Intersection
	6,  // 6: google.bigtable.admin.table.v1.GcRule.union:type_name -> google.bigtable.admin.table.v1.GcRule.Union
	2,  // 7: google.bigtable.admin.table.v1.Table.ColumnFamiliesEntry.value:type_name -> google.bigtable.admin.table.v1.ColumnFamily
	3,  // 8: google.bigtable.admin.table.v1.GcRule.Intersection.rules:type_name -> google.bigtable.admin.table.v1.GcRule
	3,  // 9: google.bigtable.admin.table.v1.GcRule.Union.rules:type_name -> google.bigtable.admin.table.v1.GcRule
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_google_bigtable_admin_table_v1_bigtable_table_data_proto_init() }
func file_google_bigtable_admin_table_v1_bigtable_table_data_proto_init() {
	if File_google_bigtable_admin_table_v1_bigtable_table_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_bigtable_admin_table_v1_bigtable_table_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Table); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_bigtable_admin_table_v1_bigtable_table_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ColumnFamily); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_bigtable_admin_table_v1_bigtable_table_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcRule); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_bigtable_admin_table_v1_bigtable_table_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcRule_Intersection); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_bigtable_admin_table_v1_bigtable_table_data_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcRule_Union); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_google_bigtable_admin_table_v1_bigtable_table_data_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*GcRule_MaxNumVersions)(nil),
		(*GcRule_MaxAge)(nil),
		(*GcRule_Intersection_)(nil),
		(*GcRule_Union_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_bigtable_admin_table_v1_bigtable_table_data_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_bigtable_admin_table_v1_bigtable_table_data_proto_goTypes,
		DependencyIndexes: file_google_bigtable_admin_table_v1_bigtable_table_data_proto_depIdxs,
		EnumInfos:         file_google_bigtable_admin_table_v1_bigtable_table_data_proto_enumTypes,
		MessageInfos:      file_google_bigtable_admin_table_v1_bigtable_table_data_proto_msgTypes,
	}.Build()
	File_google_bigtable_admin_table_v1_bigtable_table_data_proto = out.File
	file_google_bigtable_admin_table_v1_bigtable_table_data_proto_rawDesc = nil
	file_google_bigtable_admin_table_v1_bigtable_table_data_proto_goTypes = nil
	file_google_bigtable_admin_table_v1_bigtable_table_data_proto_depIdxs = nil
}

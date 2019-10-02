// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/resultstore/v2/common.proto

package resultstore

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// These correspond to the prefix of the rule name. Eg cc_test has language CC.
type Language int32

const (
	// Language unspecified or not listed here.
	Language_LANGUAGE_UNSPECIFIED Language = 0
	// Not related to any particular language
	Language_NONE Language = 1
	// Android
	Language_ANDROID Language = 2
	// ActionScript (Flash)
	Language_AS Language = 3
	// C++ or C
	Language_CC Language = 4
	// Cascading-Style-Sheets
	Language_CSS Language = 5
	// Dart
	Language_DART Language = 6
	// Go
	Language_GO Language = 7
	// Google-Web-Toolkit
	Language_GWT Language = 8
	// Haskell
	Language_HASKELL Language = 9
	// Java
	Language_JAVA Language = 10
	// Javascript
	Language_JS Language = 11
	// Lisp
	Language_LISP Language = 12
	// Objective-C
	Language_OBJC Language = 13
	// Python
	Language_PY Language = 14
	// Shell (Typically Bash)
	Language_SH Language = 15
	// Swift
	Language_SWIFT Language = 16
	// Typescript
	Language_TS Language = 18
	// Webtesting
	Language_WEB Language = 19
	// Scala
	Language_SCALA Language = 20
	// Protocol Buffer
	Language_PROTO Language = 21
)

var Language_name = map[int32]string{
	0:  "LANGUAGE_UNSPECIFIED",
	1:  "NONE",
	2:  "ANDROID",
	3:  "AS",
	4:  "CC",
	5:  "CSS",
	6:  "DART",
	7:  "GO",
	8:  "GWT",
	9:  "HASKELL",
	10: "JAVA",
	11: "JS",
	12: "LISP",
	13: "OBJC",
	14: "PY",
	15: "SH",
	16: "SWIFT",
	18: "TS",
	19: "WEB",
	20: "SCALA",
	21: "PROTO",
}

var Language_value = map[string]int32{
	"LANGUAGE_UNSPECIFIED": 0,
	"NONE":                 1,
	"ANDROID":              2,
	"AS":                   3,
	"CC":                   4,
	"CSS":                  5,
	"DART":                 6,
	"GO":                   7,
	"GWT":                  8,
	"HASKELL":              9,
	"JAVA":                 10,
	"JS":                   11,
	"LISP":                 12,
	"OBJC":                 13,
	"PY":                   14,
	"SH":                   15,
	"SWIFT":                16,
	"TS":                   18,
	"WEB":                  19,
	"SCALA":                20,
	"PROTO":                21,
}

func (x Language) String() string {
	return proto.EnumName(Language_name, int32(x))
}

func (Language) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ff56b05a77242216, []int{0}
}

// Status of a resource.
type Status int32

const (
	// The implicit default enum value. Should never be set.
	Status_STATUS_UNSPECIFIED Status = 0
	// Displays as "Building". Means the target is compiling, linking, etc.
	Status_BUILDING Status = 1
	// Displays as "Built". Means the target was built successfully.
	// If testing was requested, it should never reach this status: it should go
	// straight from BUILDING to TESTING.
	Status_BUILT Status = 2
	// Displays as "Broken". Means build failure such as compile error.
	Status_FAILED_TO_BUILD Status = 3
	// Displays as "Testing". Means the test is running.
	Status_TESTING Status = 4
	// Displays as "Passed". Means the test was run and passed.
	Status_PASSED Status = 5
	// Displays as "Failed". Means the test was run and failed.
	Status_FAILED Status = 6
	// Displays as "Timed out". Means the test didn't finish in time.
	Status_TIMED_OUT Status = 7
	// Displays as "Cancelled". Means the build or test was cancelled.
	// E.g. User hit control-C.
	Status_CANCELLED Status = 8
	// Displays as "Tool Failed". Means the build or test had internal tool
	// failure.
	Status_TOOL_FAILED Status = 9
	// Displays as "Incomplete". Means the build or test did not complete.  This
	// might happen when a build breakage or test failure causes the tool to stop
	// trying to build anything more or run any more tests, with the default
	// bazel --nokeep_going option or the --notest_keep_going option.
	Status_INCOMPLETE Status = 10
	// Displays as "Flaky". Means the aggregate status contains some runs that
	// were successful, and some that were not.
	Status_FLAKY Status = 11
	// Displays as "Unknown". Means the tool uploading to the server died
	// mid-upload or does not know the state.
	Status_UNKNOWN Status = 12
	// Displays as "Skipped". Means building and testing were skipped.
	// (E.g. Restricted to a different configuration.)
	Status_SKIPPED Status = 13
)

var Status_name = map[int32]string{
	0:  "STATUS_UNSPECIFIED",
	1:  "BUILDING",
	2:  "BUILT",
	3:  "FAILED_TO_BUILD",
	4:  "TESTING",
	5:  "PASSED",
	6:  "FAILED",
	7:  "TIMED_OUT",
	8:  "CANCELLED",
	9:  "TOOL_FAILED",
	10: "INCOMPLETE",
	11: "FLAKY",
	12: "UNKNOWN",
	13: "SKIPPED",
}

var Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"BUILDING":           1,
	"BUILT":              2,
	"FAILED_TO_BUILD":    3,
	"TESTING":            4,
	"PASSED":             5,
	"FAILED":             6,
	"TIMED_OUT":          7,
	"CANCELLED":          8,
	"TOOL_FAILED":        9,
	"INCOMPLETE":         10,
	"FLAKY":              11,
	"UNKNOWN":            12,
	"SKIPPED":            13,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ff56b05a77242216, []int{1}
}

// Indicates the upload status of the invocation, whether it is
// post-processing, or immutable, etc.
type UploadStatus int32

const (
	// The implicit default enum value. Should never be set.
	UploadStatus_UPLOAD_STATUS_UNSPECIFIED UploadStatus = 0
	// The invocation is still uploading to the ResultStore.
	UploadStatus_UPLOADING UploadStatus = 1
	// The invocation upload is complete. The ResultStore is still post-processing
	// the invocation.
	UploadStatus_POST_PROCESSING UploadStatus = 2
	// All post-processing is complete, and the invocation is now immutable.
	UploadStatus_IMMUTABLE UploadStatus = 3
)

var UploadStatus_name = map[int32]string{
	0: "UPLOAD_STATUS_UNSPECIFIED",
	1: "UPLOADING",
	2: "POST_PROCESSING",
	3: "IMMUTABLE",
}

var UploadStatus_value = map[string]int32{
	"UPLOAD_STATUS_UNSPECIFIED": 0,
	"UPLOADING":                 1,
	"POST_PROCESSING":           2,
	"IMMUTABLE":                 3,
}

func (x UploadStatus) String() string {
	return proto.EnumName(UploadStatus_name, int32(x))
}

func (UploadStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ff56b05a77242216, []int{2}
}

// Describes the status of a resource in both enum and string form.
// Only use description when conveying additional info not captured in the enum
// name.
type StatusAttributes struct {
	// Enum representation of the status.
	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=google.devtools.resultstore.v2.Status" json:"status,omitempty"`
	// A longer description about the status.
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusAttributes) Reset()         { *m = StatusAttributes{} }
func (m *StatusAttributes) String() string { return proto.CompactTextString(m) }
func (*StatusAttributes) ProtoMessage()    {}
func (*StatusAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff56b05a77242216, []int{0}
}

func (m *StatusAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusAttributes.Unmarshal(m, b)
}
func (m *StatusAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusAttributes.Marshal(b, m, deterministic)
}
func (m *StatusAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusAttributes.Merge(m, src)
}
func (m *StatusAttributes) XXX_Size() int {
	return xxx_messageInfo_StatusAttributes.Size(m)
}
func (m *StatusAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_StatusAttributes proto.InternalMessageInfo

func (m *StatusAttributes) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_STATUS_UNSPECIFIED
}

func (m *StatusAttributes) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// A generic key-value property definition.
type Property struct {
	// The key.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The value.
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Property) Reset()         { *m = Property{} }
func (m *Property) String() string { return proto.CompactTextString(m) }
func (*Property) ProtoMessage()    {}
func (*Property) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff56b05a77242216, []int{1}
}

func (m *Property) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Property.Unmarshal(m, b)
}
func (m *Property) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Property.Marshal(b, m, deterministic)
}
func (m *Property) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Property.Merge(m, src)
}
func (m *Property) XXX_Size() int {
	return xxx_messageInfo_Property.Size(m)
}
func (m *Property) XXX_DiscardUnknown() {
	xxx_messageInfo_Property.DiscardUnknown(m)
}

var xxx_messageInfo_Property proto.InternalMessageInfo

func (m *Property) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Property) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// The timing of a particular Invocation, Action, etc. The start_time is
// specified, stop time can be calculated by adding duration to start_time.
type Timing struct {
	// The time the resource started running. This is in UTC Epoch time.
	StartTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// The duration for which the resource ran.
	Duration             *duration.Duration `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Timing) Reset()         { *m = Timing{} }
func (m *Timing) String() string { return proto.CompactTextString(m) }
func (*Timing) ProtoMessage()    {}
func (*Timing) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff56b05a77242216, []int{2}
}

func (m *Timing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timing.Unmarshal(m, b)
}
func (m *Timing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timing.Marshal(b, m, deterministic)
}
func (m *Timing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timing.Merge(m, src)
}
func (m *Timing) XXX_Size() int {
	return xxx_messageInfo_Timing.Size(m)
}
func (m *Timing) XXX_DiscardUnknown() {
	xxx_messageInfo_Timing.DiscardUnknown(m)
}

var xxx_messageInfo_Timing proto.InternalMessageInfo

func (m *Timing) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Timing) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

// Represents a dependency of a resource on another resource. This can be used
// to define a graph or a workflow paradigm through resources.
type Dependency struct {
	// The resource depended upon. It may be a Target, ConfiguredTarget, or
	// Action.
	//
	// Types that are valid to be assigned to Resource:
	//	*Dependency_Target
	//	*Dependency_ConfiguredTarget
	//	*Dependency_Action
	Resource isDependency_Resource `protobuf_oneof:"resource"`
	// A label describing this dependency.
	// The label "Root Cause" is handled specially. It is used to point to the
	// exact resource that caused a resource to fail.
	Label                string   `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dependency) Reset()         { *m = Dependency{} }
func (m *Dependency) String() string { return proto.CompactTextString(m) }
func (*Dependency) ProtoMessage()    {}
func (*Dependency) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff56b05a77242216, []int{3}
}

func (m *Dependency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dependency.Unmarshal(m, b)
}
func (m *Dependency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dependency.Marshal(b, m, deterministic)
}
func (m *Dependency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dependency.Merge(m, src)
}
func (m *Dependency) XXX_Size() int {
	return xxx_messageInfo_Dependency.Size(m)
}
func (m *Dependency) XXX_DiscardUnknown() {
	xxx_messageInfo_Dependency.DiscardUnknown(m)
}

var xxx_messageInfo_Dependency proto.InternalMessageInfo

type isDependency_Resource interface {
	isDependency_Resource()
}

type Dependency_Target struct {
	Target string `protobuf:"bytes,1,opt,name=target,proto3,oneof"`
}

type Dependency_ConfiguredTarget struct {
	ConfiguredTarget string `protobuf:"bytes,2,opt,name=configured_target,json=configuredTarget,proto3,oneof"`
}

type Dependency_Action struct {
	Action string `protobuf:"bytes,3,opt,name=action,proto3,oneof"`
}

func (*Dependency_Target) isDependency_Resource() {}

func (*Dependency_ConfiguredTarget) isDependency_Resource() {}

func (*Dependency_Action) isDependency_Resource() {}

func (m *Dependency) GetResource() isDependency_Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *Dependency) GetTarget() string {
	if x, ok := m.GetResource().(*Dependency_Target); ok {
		return x.Target
	}
	return ""
}

func (m *Dependency) GetConfiguredTarget() string {
	if x, ok := m.GetResource().(*Dependency_ConfiguredTarget); ok {
		return x.ConfiguredTarget
	}
	return ""
}

func (m *Dependency) GetAction() string {
	if x, ok := m.GetResource().(*Dependency_Action); ok {
		return x.Action
	}
	return ""
}

func (m *Dependency) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Dependency) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Dependency_Target)(nil),
		(*Dependency_ConfiguredTarget)(nil),
		(*Dependency_Action)(nil),
	}
}

func init() {
	proto.RegisterEnum("google.devtools.resultstore.v2.Language", Language_name, Language_value)
	proto.RegisterEnum("google.devtools.resultstore.v2.Status", Status_name, Status_value)
	proto.RegisterEnum("google.devtools.resultstore.v2.UploadStatus", UploadStatus_name, UploadStatus_value)
	proto.RegisterType((*StatusAttributes)(nil), "google.devtools.resultstore.v2.StatusAttributes")
	proto.RegisterType((*Property)(nil), "google.devtools.resultstore.v2.Property")
	proto.RegisterType((*Timing)(nil), "google.devtools.resultstore.v2.Timing")
	proto.RegisterType((*Dependency)(nil), "google.devtools.resultstore.v2.Dependency")
}

func init() {
	proto.RegisterFile("google/devtools/resultstore/v2/common.proto", fileDescriptor_ff56b05a77242216)
}

var fileDescriptor_ff56b05a77242216 = []byte{
	// 755 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xdf, 0x6e, 0xe3, 0x44,
	0x14, 0xc6, 0xd7, 0x49, 0xeb, 0x26, 0x27, 0xfd, 0x73, 0x98, 0x16, 0x94, 0xad, 0x04, 0x54, 0xb9,
	0x40, 0xab, 0x22, 0x12, 0x29, 0x88, 0x0b, 0x84, 0x84, 0x34, 0xb1, 0xdd, 0xd4, 0xad, 0x6b, 0x5b,
	0x9e, 0x31, 0xd5, 0x72, 0x63, 0xdc, 0x64, 0xd6, 0x8a, 0x48, 0x3c, 0xc1, 0x1e, 0x47, 0x2a, 0x6f,
	0xc1, 0xd3, 0xf1, 0x02, 0x3c, 0x08, 0x1a, 0xdb, 0x81, 0x6a, 0xf9, 0xb3, 0x57, 0x9e, 0xef, 0x7c,
	0xbf, 0x73, 0x7c, 0xbe, 0x51, 0x1c, 0xf8, 0x32, 0x93, 0x32, 0x5b, 0x8b, 0xc9, 0x52, 0xec, 0x94,
	0x94, 0xeb, 0x72, 0x52, 0x88, 0xb2, 0x5a, 0xab, 0x52, 0xc9, 0x42, 0x4c, 0x76, 0xd3, 0xc9, 0x42,
	0x6e, 0x36, 0x32, 0x1f, 0x6f, 0x0b, 0xa9, 0x24, 0xf9, 0xac, 0x81, 0xc7, 0x7b, 0x78, 0xfc, 0x02,
	0x1e, 0xef, 0xa6, 0x97, 0xad, 0x3f, 0xa9, 0xe9, 0xa7, 0xea, 0xdd, 0x64, 0x59, 0x15, 0xa9, 0x5a,
	0xed, 0xfb, 0x2f, 0x3f, 0x7f, 0xdf, 0x57, 0xab, 0x8d, 0x28, 0x55, 0xba, 0xd9, 0x36, 0xc0, 0x48,
	0x01, 0x32, 0x95, 0xaa, 0xaa, 0xa4, 0x4a, 0x15, 0xab, 0xa7, 0x4a, 0x89, 0x92, 0x7c, 0x0f, 0x66,
	0x59, 0xd7, 0x86, 0xc6, 0x95, 0xf1, 0xe6, 0x74, 0xfa, 0xc5, 0xf8, 0xff, 0xb7, 0x18, 0x37, 0x13,
	0xa2, 0xb6, 0x8b, 0x5c, 0xc1, 0x60, 0x29, 0xca, 0x45, 0xb1, 0xda, 0xea, 0x4d, 0x86, 0x9d, 0x2b,
	0xe3, 0x4d, 0x3f, 0x7a, 0x59, 0x1a, 0x4d, 0xa1, 0x17, 0x16, 0x72, 0x2b, 0x0a, 0xf5, 0x4c, 0x10,
	0xba, 0x3f, 0x8b, 0xe7, 0xfa, 0x55, 0xfd, 0x48, 0x1f, 0xc9, 0x05, 0x1c, 0xee, 0xd2, 0x75, 0x25,
	0xda, 0xce, 0x46, 0x8c, 0x7e, 0x05, 0x93, 0xaf, 0x36, 0xab, 0x3c, 0x23, 0xdf, 0x02, 0x94, 0x2a,
	0x2d, 0x54, 0xa2, 0xc3, 0xd4, 0x8d, 0x83, 0xe9, 0xe5, 0x7e, 0xc7, 0x7d, 0xd2, 0x31, 0xdf, 0x27,
	0x8d, 0xfa, 0x35, 0xad, 0x35, 0xf9, 0x06, 0x7a, 0xfb, 0x1b, 0xaa, 0xa7, 0x0f, 0xa6, 0xaf, 0xff,
	0xd1, 0x68, 0xb7, 0x40, 0xf4, 0x17, 0x3a, 0xfa, 0xcd, 0x00, 0xb0, 0xc5, 0x56, 0xe4, 0x4b, 0x91,
	0x2f, 0x9e, 0xc9, 0x10, 0x4c, 0x95, 0x16, 0x99, 0x50, 0xcd, 0xd6, 0xb7, 0xaf, 0xa2, 0x56, 0x93,
	0xaf, 0xe0, 0xa3, 0x85, 0xcc, 0xdf, 0xad, 0xb2, 0xaa, 0x10, 0xcb, 0xa4, 0x85, 0x3a, 0x2d, 0x84,
	0x7f, 0x5b, 0xbc, 0xc1, 0x87, 0x60, 0xa6, 0x8b, 0x7a, 0x99, 0xee, 0x7e, 0x50, 0xa3, 0xf5, 0x1d,
	0xac, 0xd3, 0x27, 0xb1, 0x1e, 0x1e, 0x34, 0x77, 0x50, 0x8b, 0x19, 0x40, 0xaf, 0x10, 0xa5, 0xac,
	0x8a, 0x85, 0xb8, 0xfe, 0xc3, 0x80, 0x9e, 0x97, 0xe6, 0x59, 0x95, 0x66, 0x82, 0x0c, 0xe1, 0xc2,
	0xa3, 0xfe, 0x3c, 0xa6, 0x73, 0x27, 0x89, 0x7d, 0x16, 0x3a, 0x96, 0x7b, 0xe3, 0x3a, 0x36, 0xbe,
	0x22, 0x3d, 0x38, 0xf0, 0x03, 0xdf, 0x41, 0x83, 0x0c, 0xe0, 0x88, 0xfa, 0x76, 0x14, 0xb8, 0x36,
	0x76, 0x88, 0x09, 0x1d, 0xca, 0xb0, 0xab, 0x9f, 0x96, 0x85, 0x07, 0xe4, 0x08, 0xba, 0x16, 0x63,
	0x78, 0xa8, 0x79, 0x9b, 0x46, 0x1c, 0x4d, 0x6d, 0xcd, 0x03, 0x3c, 0xd2, 0xd6, 0xfc, 0x91, 0x63,
	0x4f, 0x0f, 0xb8, 0xa5, 0xec, 0xde, 0xf1, 0x3c, 0xec, 0x6b, 0xee, 0x8e, 0xfe, 0x40, 0x11, 0x34,
	0x77, 0xc7, 0x70, 0xa0, 0x2b, 0x9e, 0xcb, 0x42, 0x3c, 0xd6, 0xa7, 0x60, 0x76, 0x67, 0xe1, 0x89,
	0xf6, 0xc2, 0xb7, 0x78, 0xaa, 0x9f, 0xec, 0x16, 0xcf, 0x48, 0x1f, 0x0e, 0xd9, 0xa3, 0x7b, 0xc3,
	0x11, 0x75, 0x89, 0x33, 0x24, 0x7a, 0xfc, 0xa3, 0x33, 0xc3, 0xf3, 0xda, 0xb3, 0xa8, 0x47, 0xf1,
	0x42, 0x1f, 0xc3, 0x28, 0xe0, 0x01, 0x7e, 0x7c, 0xfd, 0xbb, 0x01, 0x66, 0xf3, 0xfb, 0x22, 0x9f,
	0x00, 0x61, 0x9c, 0xf2, 0x98, 0xbd, 0x17, 0xf1, 0x18, 0x7a, 0xb3, 0xd8, 0xf5, 0x6c, 0xd7, 0x9f,
	0xa3, 0xa1, 0x7b, 0xb5, 0xe2, 0xd8, 0x21, 0xe7, 0x70, 0x76, 0x43, 0x5d, 0xcf, 0xb1, 0x13, 0x1e,
	0x24, 0x35, 0x82, 0x5d, 0x9d, 0x82, 0x3b, 0x8c, 0x6b, 0xf8, 0x80, 0x00, 0x98, 0x21, 0x65, 0xcc,
	0xb1, 0xf1, 0x50, 0x9f, 0x1b, 0x1a, 0x4d, 0x72, 0x02, 0x7d, 0xee, 0x3e, 0x38, 0x76, 0x12, 0xc4,
	0x1c, 0x8f, 0xb4, 0xb4, 0xa8, 0x6f, 0x39, 0x9e, 0x76, 0x7b, 0xe4, 0x0c, 0x06, 0x3c, 0x08, 0xbc,
	0xa4, 0xc5, 0xfb, 0xe4, 0x14, 0xc0, 0xf5, 0xad, 0xe0, 0x21, 0xf4, 0x1c, 0xee, 0x20, 0xe8, 0x1d,
	0x6e, 0x3c, 0x7a, 0xff, 0x16, 0x07, 0xfa, 0x75, 0xb1, 0x7f, 0xef, 0x07, 0x8f, 0x3e, 0x1e, 0x6b,
	0xc1, 0xee, 0xdd, 0x30, 0x74, 0x6c, 0x3c, 0xb9, 0xfe, 0x09, 0x8e, 0xe3, 0xed, 0x5a, 0xa6, 0xcb,
	0x36, 0xde, 0xa7, 0xf0, 0x3a, 0x0e, 0xbd, 0x80, 0xda, 0xc9, 0xbf, 0xa6, 0x3c, 0x81, 0x7e, 0x63,
	0x37, 0x31, 0xcf, 0xe1, 0x2c, 0x0c, 0x18, 0x4f, 0xc2, 0x28, 0xb0, 0x1c, 0xc6, 0x74, 0xb1, 0xa3,
	0x19, 0xf7, 0xe1, 0x21, 0xe6, 0x74, 0xe6, 0x39, 0xd8, 0x9d, 0xfd, 0x02, 0xa3, 0x85, 0xdc, 0x7c,
	0xe0, 0xeb, 0x0d, 0x8d, 0x1f, 0xdd, 0x96, 0xc8, 0xe4, 0x3a, 0xcd, 0xb3, 0xb1, 0x2c, 0xb2, 0x49,
	0x26, 0xf2, 0xfa, 0x83, 0x98, 0x34, 0x56, 0xba, 0x5d, 0x95, 0xff, 0xf5, 0x8f, 0xf5, 0xdd, 0x0b,
	0xf9, 0x64, 0xd6, 0x5d, 0x5f, 0xff, 0x19, 0x00, 0x00, 0xff, 0xff, 0x76, 0xbe, 0x13, 0x5f, 0xe6,
	0x04, 0x00, 0x00,
}

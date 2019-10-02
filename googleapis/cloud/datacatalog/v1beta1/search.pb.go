// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/datacatalog/v1beta1/search.proto

package datacatalog

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

// The different types of resources that can be returned in search.
type SearchResultType int32

const (
	// Default unknown type.
	SearchResultType_SEARCH_RESULT_TYPE_UNSPECIFIED SearchResultType = 0
	// An [Entry][google.cloud.datacatalog.v1beta1.Entry].
	SearchResultType_ENTRY SearchResultType = 1
	// A [TagTemplate][google.cloud.datacatalog.v1beta1.TagTemplate].
	SearchResultType_TAG_TEMPLATE SearchResultType = 2
	// An [EntryGroup][google.cloud.datacatalog.v1beta1.EntryGroup].
	SearchResultType_ENTRY_GROUP SearchResultType = 3
)

var SearchResultType_name = map[int32]string{
	0: "SEARCH_RESULT_TYPE_UNSPECIFIED",
	1: "ENTRY",
	2: "TAG_TEMPLATE",
	3: "ENTRY_GROUP",
}

var SearchResultType_value = map[string]int32{
	"SEARCH_RESULT_TYPE_UNSPECIFIED": 0,
	"ENTRY":                          1,
	"TAG_TEMPLATE":                   2,
	"ENTRY_GROUP":                    3,
}

func (x SearchResultType) String() string {
	return proto.EnumName(SearchResultType_name, int32(x))
}

func (SearchResultType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_075a004a2049c613, []int{0}
}

// A result that appears in the response of a search request. Each result
// captures details of one entry that matches the search.
type SearchCatalogResult struct {
	// Type of the search result. This field can be used to determine which Get
	// method to call to fetch the full resource.
	SearchResultType SearchResultType `protobuf:"varint,1,opt,name=search_result_type,json=searchResultType,proto3,enum=google.cloud.datacatalog.v1beta1.SearchResultType" json:"search_result_type,omitempty"`
	// Sub-type of the search result. This is a dot-delimited description of the
	// resource's full type, and is the same as the value callers would provide in
	// the "type" search facet.  Examples: `entry.table`, `entry.dataStream`,
	// `tagTemplate`.
	SearchResultSubtype string `protobuf:"bytes,2,opt,name=search_result_subtype,json=searchResultSubtype,proto3" json:"search_result_subtype,omitempty"`
	// The relative resource name of the resource in URL format.
	// Examples:
	//
	//  * `projects/{project_id}/locations/{location_id}/entryGroups/{entry_group_id}/entries/{entry_id}`
	//  * `projects/{project_id}/tagTemplates/{tag_template_id}`
	RelativeResourceName string `protobuf:"bytes,3,opt,name=relative_resource_name,json=relativeResourceName,proto3" json:"relative_resource_name,omitempty"`
	// The full name of the cloud resource the entry belongs to. See:
	// https://cloud.google.com/apis/design/resource_names#full_resource_name.
	// Example:
	//
	//  * `//bigquery.googleapis.com/projects/projectId/datasets/datasetId/tables/tableId`
	LinkedResource       string   `protobuf:"bytes,4,opt,name=linked_resource,json=linkedResource,proto3" json:"linked_resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchCatalogResult) Reset()         { *m = SearchCatalogResult{} }
func (m *SearchCatalogResult) String() string { return proto.CompactTextString(m) }
func (*SearchCatalogResult) ProtoMessage()    {}
func (*SearchCatalogResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_075a004a2049c613, []int{0}
}

func (m *SearchCatalogResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchCatalogResult.Unmarshal(m, b)
}
func (m *SearchCatalogResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchCatalogResult.Marshal(b, m, deterministic)
}
func (m *SearchCatalogResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchCatalogResult.Merge(m, src)
}
func (m *SearchCatalogResult) XXX_Size() int {
	return xxx_messageInfo_SearchCatalogResult.Size(m)
}
func (m *SearchCatalogResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchCatalogResult.DiscardUnknown(m)
}

var xxx_messageInfo_SearchCatalogResult proto.InternalMessageInfo

func (m *SearchCatalogResult) GetSearchResultType() SearchResultType {
	if m != nil {
		return m.SearchResultType
	}
	return SearchResultType_SEARCH_RESULT_TYPE_UNSPECIFIED
}

func (m *SearchCatalogResult) GetSearchResultSubtype() string {
	if m != nil {
		return m.SearchResultSubtype
	}
	return ""
}

func (m *SearchCatalogResult) GetRelativeResourceName() string {
	if m != nil {
		return m.RelativeResourceName
	}
	return ""
}

func (m *SearchCatalogResult) GetLinkedResource() string {
	if m != nil {
		return m.LinkedResource
	}
	return ""
}

func init() {
	proto.RegisterEnum("google.cloud.datacatalog.v1beta1.SearchResultType", SearchResultType_name, SearchResultType_value)
	proto.RegisterType((*SearchCatalogResult)(nil), "google.cloud.datacatalog.v1beta1.SearchCatalogResult")
}

func init() {
	proto.RegisterFile("google/cloud/datacatalog/v1beta1/search.proto", fileDescriptor_075a004a2049c613)
}

var fileDescriptor_075a004a2049c613 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4f, 0xf2, 0x30,
	0x00, 0xc5, 0xbf, 0xc1, 0xf7, 0x7d, 0x09, 0xd5, 0xc0, 0x52, 0xd4, 0x10, 0x63, 0x94, 0x70, 0x91,
	0x98, 0xb8, 0x05, 0xf4, 0xe6, 0x09, 0xb1, 0x22, 0x11, 0x71, 0xe9, 0xc6, 0x01, 0x2f, 0xb5, 0xdb,
	0xea, 0x5c, 0xdc, 0xd6, 0x65, 0xeb, 0x48, 0xf8, 0x17, 0xfc, 0x8b, 0x3d, 0x1a, 0xdb, 0x61, 0x90,
	0x84, 0x78, 0x7d, 0xef, 0xf7, 0xde, 0x6b, 0xda, 0x82, 0xf3, 0x80, 0xf3, 0x20, 0x62, 0xa6, 0x17,
	0xf1, 0xc2, 0x37, 0x7d, 0x2a, 0xa8, 0x47, 0x05, 0x8d, 0x78, 0x60, 0x2e, 0x7a, 0x2e, 0x13, 0xb4,
	0x67, 0xe6, 0x8c, 0x66, 0xde, 0xab, 0x91, 0x66, 0x5c, 0x70, 0xd8, 0x56, 0xb8, 0x21, 0x71, 0x63,
	0x0d, 0x37, 0x4a, 0xfc, 0xf0, 0xa4, 0x2c, 0x94, 0xbc, 0x5b, 0xbc, 0x98, 0x22, 0x8c, 0x59, 0x2e,
	0x68, 0x9c, 0xaa, 0x8a, 0xce, 0x7b, 0x05, 0x34, 0x6d, 0xd9, 0x39, 0x54, 0x51, 0xcc, 0xf2, 0x22,
	0x12, 0xf0, 0x19, 0x40, 0x35, 0x45, 0x32, 0x29, 0x10, 0xb1, 0x4c, 0x59, 0x4b, 0x6b, 0x6b, 0xdd,
	0x7a, 0xbf, 0x6f, 0xfc, 0xb6, 0x6b, 0xa8, 0x4a, 0xd5, 0xe5, 0x2c, 0x53, 0x86, 0xf5, 0x7c, 0x43,
	0x81, 0x7d, 0xb0, 0xff, 0x73, 0x21, 0x2f, 0x5c, 0x39, 0x52, 0x69, 0x6b, 0xdd, 0x1a, 0x6e, 0xae,
	0x07, 0x6c, 0x65, 0xc1, 0x4b, 0x70, 0x90, 0xb1, 0x88, 0x8a, 0x70, 0xc1, 0xbe, 0x52, 0xbc, 0xc8,
	0x3c, 0x46, 0x12, 0x1a, 0xb3, 0x56, 0x55, 0x86, 0xf6, 0x56, 0x2e, 0x2e, 0xcd, 0x29, 0x8d, 0x19,
	0x3c, 0x05, 0x8d, 0x28, 0x4c, 0xde, 0x98, 0xff, 0x9d, 0x69, 0xfd, 0x95, 0x78, 0x5d, 0xc9, 0x2b,
	0xf8, 0xcc, 0x07, 0xfa, 0xe6, 0xc1, 0x61, 0x07, 0x1c, 0xdb, 0x68, 0x80, 0x87, 0x77, 0x04, 0x23,
	0x7b, 0x36, 0x71, 0x88, 0x33, 0xb7, 0x10, 0x99, 0x4d, 0x6d, 0x0b, 0x0d, 0xc7, 0xb7, 0x63, 0x74,
	0xa3, 0xff, 0x81, 0x35, 0xf0, 0x0f, 0x4d, 0x1d, 0x3c, 0xd7, 0x35, 0xa8, 0x83, 0x5d, 0x67, 0x30,
	0x22, 0x0e, 0x7a, 0xb0, 0x26, 0x03, 0x07, 0xe9, 0x15, 0xd8, 0x00, 0x3b, 0xd2, 0x24, 0x23, 0xfc,
	0x38, 0xb3, 0xf4, 0xea, 0x75, 0x0a, 0x8e, 0x3c, 0x1e, 0x6f, 0xbd, 0x43, 0x4b, 0x7b, 0xba, 0x2f,
	0xbd, 0x80, 0x47, 0x34, 0x09, 0x0c, 0x9e, 0x05, 0x66, 0xc0, 0x12, 0xf9, 0x60, 0xa6, 0xb2, 0x68,
	0x1a, 0xe6, 0xdb, 0x7f, 0xc9, 0xd5, 0x9a, 0xf6, 0xa1, 0x69, 0xee, 0x7f, 0x19, 0xbd, 0xf8, 0x0c,
	0x00, 0x00, 0xff, 0xff, 0xaa, 0xbf, 0x92, 0x1a, 0x5f, 0x02, 0x00, 0x00,
}

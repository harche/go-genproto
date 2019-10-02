// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/websecurityscanner/v1beta/crawled_url.proto

package websecurityscanner

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

// A CrawledUrl resource represents a URL that was crawled during a ScanRun. Web
// Security Scanner Service crawls the web applications, following all links
// within the scope of sites, to find the URLs to test against.
type CrawledUrl struct {
	// Output only.
	// The http method of the request that was used to visit the URL, in
	// uppercase.
	HttpMethod string `protobuf:"bytes,1,opt,name=http_method,json=httpMethod,proto3" json:"http_method,omitempty"`
	// Output only.
	// The URL that was crawled.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// Output only.
	// The body of the request that was used to visit the URL.
	Body                 string   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CrawledUrl) Reset()         { *m = CrawledUrl{} }
func (m *CrawledUrl) String() string { return proto.CompactTextString(m) }
func (*CrawledUrl) ProtoMessage()    {}
func (*CrawledUrl) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d117cd3bfcd0897, []int{0}
}

func (m *CrawledUrl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CrawledUrl.Unmarshal(m, b)
}
func (m *CrawledUrl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CrawledUrl.Marshal(b, m, deterministic)
}
func (m *CrawledUrl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrawledUrl.Merge(m, src)
}
func (m *CrawledUrl) XXX_Size() int {
	return xxx_messageInfo_CrawledUrl.Size(m)
}
func (m *CrawledUrl) XXX_DiscardUnknown() {
	xxx_messageInfo_CrawledUrl.DiscardUnknown(m)
}

var xxx_messageInfo_CrawledUrl proto.InternalMessageInfo

func (m *CrawledUrl) GetHttpMethod() string {
	if m != nil {
		return m.HttpMethod
	}
	return ""
}

func (m *CrawledUrl) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *CrawledUrl) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*CrawledUrl)(nil), "google.cloud.websecurityscanner.v1beta.CrawledUrl")
}

func init() {
	proto.RegisterFile("google/cloud/websecurityscanner/v1beta/crawled_url.proto", fileDescriptor_4d117cd3bfcd0897)
}

var fileDescriptor_4d117cd3bfcd0897 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x3f, 0x4b, 0x04, 0x31,
	0x10, 0xc5, 0xd9, 0x3b, 0x11, 0x8c, 0x85, 0x92, 0x6a, 0x11, 0x41, 0xb1, 0x38, 0xc4, 0x22, 0x41,
	0x6c, 0x04, 0xbb, 0xbb, 0xc2, 0x4a, 0x38, 0x5c, 0xfc, 0x83, 0x2c, 0x1c, 0x49, 0x76, 0xc8, 0x2d,
	0xe4, 0x32, 0x4b, 0x36, 0xeb, 0x71, 0x5f, 0xcf, 0xd2, 0x4f, 0x25, 0x99, 0x2c, 0x58, 0x5c, 0xa1,
	0xdd, 0xf0, 0xde, 0xfc, 0xe6, 0x3d, 0x86, 0xdd, 0x5b, 0x44, 0xeb, 0x40, 0x1a, 0x87, 0x43, 0x23,
	0xb7, 0xa0, 0x7b, 0x30, 0x43, 0x68, 0xe3, 0xae, 0x37, 0xca, 0x7b, 0x08, 0xf2, 0xf3, 0x56, 0x43,
	0x54, 0xd2, 0x04, 0xb5, 0x75, 0xd0, 0xac, 0x86, 0xe0, 0x44, 0x17, 0x30, 0x22, 0x9f, 0x65, 0x52,
	0x10, 0x29, 0xf6, 0x49, 0x91, 0xc9, 0xb3, 0xf3, 0x31, 0x41, 0x75, 0xad, 0x54, 0xde, 0x63, 0x54,
	0xb1, 0x45, 0xdf, 0xe7, 0x2b, 0x57, 0x15, 0x63, 0x8b, 0x7c, 0xfa, 0x25, 0x38, 0x7e, 0xc1, 0x8e,
	0xd7, 0x31, 0x76, 0xab, 0x0d, 0xc4, 0x35, 0x36, 0x65, 0x71, 0x59, 0x5c, 0x1f, 0x3d, 0xb3, 0x24,
	0x3d, 0x91, 0xc2, 0x4f, 0xd9, 0x74, 0x08, 0xae, 0x9c, 0x90, 0x91, 0x46, 0xce, 0xd9, 0x81, 0xc6,
	0x66, 0x57, 0x4e, 0x49, 0xa2, 0x79, 0xfe, 0x55, 0xb0, 0x1b, 0x83, 0x1b, 0xf1, 0xbf, 0x86, 0xf3,
	0x93, 0xdf, 0x06, 0xcb, 0x54, 0x6a, 0x59, 0x7c, 0xbc, 0x8f, 0xa8, 0x45, 0xa7, 0xbc, 0x15, 0x18,
	0xac, 0xb4, 0xe0, 0xa9, 0xb2, 0xcc, 0x96, 0xea, 0xda, 0xfe, 0xaf, 0xaf, 0x3d, 0xec, 0x3b, 0xdf,
	0x93, 0xd9, 0x23, 0xf1, 0xf5, 0x22, 0xb1, 0xf5, 0x1b, 0xe8, 0x6a, 0xdc, 0xa8, 0xf2, 0x46, 0xfd,
	0x4a, 0xac, 0x3e, 0xa4, 0xb4, 0xbb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x03, 0x6c, 0xc3,
	0xa2, 0x01, 0x00, 0x00,
}

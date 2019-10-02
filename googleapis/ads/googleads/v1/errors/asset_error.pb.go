// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/asset_error.proto

package errors

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

// Enum describing possible asset errors.
type AssetErrorEnum_AssetError int32

const (
	// Enum unspecified.
	AssetErrorEnum_UNSPECIFIED AssetErrorEnum_AssetError = 0
	// The received error code is not known in this version.
	AssetErrorEnum_UNKNOWN AssetErrorEnum_AssetError = 1
	// The customer is not whitelisted for this asset type.
	AssetErrorEnum_CUSTOMER_NOT_WHITELISTED_FOR_ASSET_TYPE AssetErrorEnum_AssetError = 2
	// Assets are duplicated across operations.
	AssetErrorEnum_DUPLICATE_ASSET AssetErrorEnum_AssetError = 3
	// The asset name is duplicated, either across operations or with an
	// existing asset.
	AssetErrorEnum_DUPLICATE_ASSET_NAME AssetErrorEnum_AssetError = 4
	// The Asset.asset_data oneof is empty.
	AssetErrorEnum_ASSET_DATA_IS_MISSING AssetErrorEnum_AssetError = 5
	// The asset has a name which is different from an existing duplicate that
	// represents the same content.
	AssetErrorEnum_CANNOT_MODIFY_ASSET_NAME AssetErrorEnum_AssetError = 6
)

var AssetErrorEnum_AssetError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CUSTOMER_NOT_WHITELISTED_FOR_ASSET_TYPE",
	3: "DUPLICATE_ASSET",
	4: "DUPLICATE_ASSET_NAME",
	5: "ASSET_DATA_IS_MISSING",
	6: "CANNOT_MODIFY_ASSET_NAME",
}

var AssetErrorEnum_AssetError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"CUSTOMER_NOT_WHITELISTED_FOR_ASSET_TYPE": 2,
	"DUPLICATE_ASSET":                         3,
	"DUPLICATE_ASSET_NAME":                    4,
	"ASSET_DATA_IS_MISSING":                   5,
	"CANNOT_MODIFY_ASSET_NAME":                6,
}

func (x AssetErrorEnum_AssetError) String() string {
	return proto.EnumName(AssetErrorEnum_AssetError_name, int32(x))
}

func (AssetErrorEnum_AssetError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d1e46d0937f553c0, []int{0, 0}
}

// Container for enum describing possible asset errors.
type AssetErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssetErrorEnum) Reset()         { *m = AssetErrorEnum{} }
func (m *AssetErrorEnum) String() string { return proto.CompactTextString(m) }
func (*AssetErrorEnum) ProtoMessage()    {}
func (*AssetErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e46d0937f553c0, []int{0}
}

func (m *AssetErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetErrorEnum.Unmarshal(m, b)
}
func (m *AssetErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetErrorEnum.Marshal(b, m, deterministic)
}
func (m *AssetErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetErrorEnum.Merge(m, src)
}
func (m *AssetErrorEnum) XXX_Size() int {
	return xxx_messageInfo_AssetErrorEnum.Size(m)
}
func (m *AssetErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AssetErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.AssetErrorEnum_AssetError", AssetErrorEnum_AssetError_name, AssetErrorEnum_AssetError_value)
	proto.RegisterType((*AssetErrorEnum)(nil), "google.ads.googleads.v1.errors.AssetErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/asset_error.proto", fileDescriptor_d1e46d0937f553c0)
}

var fileDescriptor_d1e46d0937f553c0 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x49, 0x16, 0x16, 0xc9, 0x2b, 0xd1, 0xc8, 0x80, 0xb4, 0xa0, 0xd5, 0x1e, 0x72, 0xe1,
	0x80, 0xe4, 0x10, 0x71, 0x33, 0x27, 0x6f, 0xe2, 0x16, 0x8b, 0x8d, 0x13, 0xd5, 0x4e, 0xab, 0xa2,
	0x48, 0x56, 0x20, 0x51, 0x54, 0xa9, 0x8d, 0xab, 0x38, 0xf4, 0x81, 0x38, 0xf2, 0x12, 0xdc, 0x39,
	0xf2, 0x18, 0x7d, 0x0a, 0x94, 0x98, 0xb6, 0x08, 0x09, 0x4e, 0xfe, 0x67, 0xfc, 0xfd, 0x33, 0xf6,
	0x0c, 0x78, 0xd3, 0x68, 0xdd, 0x6c, 0xea, 0xa0, 0xac, 0x4c, 0x60, 0xe5, 0xa0, 0xf6, 0x61, 0x50,
	0x77, 0x9d, 0xee, 0x4c, 0x50, 0x1a, 0x53, 0xf7, 0x6a, 0x0c, 0xd0, 0xae, 0xd3, 0xbd, 0x86, 0xb7,
	0x16, 0x43, 0x65, 0x65, 0xd0, 0xc9, 0x81, 0xf6, 0x21, 0xb2, 0x8e, 0x97, 0x37, 0xc7, 0x8a, 0xbb,
	0x75, 0x50, 0xb6, 0xad, 0xee, 0xcb, 0x7e, 0xad, 0x5b, 0x63, 0xdd, 0xfe, 0x4f, 0x07, 0x3c, 0x21,
	0x43, 0x4d, 0x3a, 0xd0, 0xb4, 0xfd, 0xb2, 0xf5, 0xbf, 0x3b, 0x00, 0x9c, 0x53, 0x70, 0x02, 0xae,
	0x72, 0x2e, 0x32, 0x1a, 0xb1, 0x29, 0xa3, 0xb1, 0xf7, 0x00, 0x5e, 0x81, 0xc7, 0x39, 0xff, 0xc0,
	0xd3, 0x25, 0xf7, 0x1c, 0xf8, 0x1a, 0xbc, 0x8a, 0x72, 0x21, 0xd3, 0x84, 0xce, 0x15, 0x4f, 0xa5,
	0x5a, 0xbe, 0x67, 0x92, 0xde, 0x33, 0x21, 0x69, 0xac, 0xa6, 0xe9, 0x5c, 0x11, 0x21, 0xa8, 0x54,
	0x72, 0x95, 0x51, 0xcf, 0x85, 0x4f, 0xc1, 0x24, 0xce, 0xb3, 0x7b, 0x16, 0x11, 0x49, 0xed, 0x8d,
	0x77, 0x01, 0xaf, 0xc1, 0xb3, 0xbf, 0x92, 0x8a, 0x93, 0x84, 0x7a, 0x0f, 0xe1, 0x0b, 0xf0, 0xdc,
	0xc6, 0x31, 0x91, 0x44, 0x31, 0xa1, 0x12, 0x26, 0x04, 0xe3, 0x33, 0xef, 0x11, 0xbc, 0x01, 0xd7,
	0x11, 0xe1, 0x43, 0xc3, 0x24, 0x8d, 0xd9, 0x74, 0xf5, 0xa7, 0xf1, 0xf2, 0xee, 0xe0, 0x00, 0xff,
	0xb3, 0xde, 0xa2, 0xff, 0x4f, 0xe6, 0x6e, 0x72, 0xfe, 0x65, 0x36, 0x0c, 0x23, 0x73, 0x3e, 0xc6,
	0xbf, 0x2d, 0x8d, 0xde, 0x94, 0x6d, 0x83, 0x74, 0xd7, 0x04, 0x4d, 0xdd, 0x8e, 0xa3, 0x3a, 0xae,
	0x63, 0xb7, 0x36, 0xff, 0xda, 0xce, 0x3b, 0x7b, 0x7c, 0x75, 0x2f, 0x66, 0x84, 0x7c, 0x73, 0x6f,
	0x67, 0xb6, 0x18, 0xa9, 0x0c, 0xb2, 0x72, 0x50, 0x8b, 0x10, 0x8d, 0x2d, 0xcd, 0x8f, 0x23, 0x50,
	0x90, 0xca, 0x14, 0x27, 0xa0, 0x58, 0x84, 0x85, 0x05, 0x0e, 0xae, 0x6f, 0xb3, 0x18, 0x93, 0xca,
	0x60, 0x7c, 0x42, 0x30, 0x5e, 0x84, 0x18, 0x5b, 0xe8, 0xd3, 0xe5, 0xf8, 0xba, 0xb7, 0xbf, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xcc, 0x3d, 0xa0, 0x27, 0x3a, 0x02, 0x00, 0x00,
}

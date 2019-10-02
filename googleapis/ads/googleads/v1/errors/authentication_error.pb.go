// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/errors/authentication_error.proto

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

// Enum describing possible authentication errors.
type AuthenticationErrorEnum_AuthenticationError int32

const (
	// Enum unspecified.
	AuthenticationErrorEnum_UNSPECIFIED AuthenticationErrorEnum_AuthenticationError = 0
	// The received error code is not known in this version.
	AuthenticationErrorEnum_UNKNOWN AuthenticationErrorEnum_AuthenticationError = 1
	// Authentication of the request failed.
	AuthenticationErrorEnum_AUTHENTICATION_ERROR AuthenticationErrorEnum_AuthenticationError = 2
	// Client Customer Id is not a number.
	AuthenticationErrorEnum_CLIENT_CUSTOMER_ID_INVALID AuthenticationErrorEnum_AuthenticationError = 5
	// No customer found for the provided customer id.
	AuthenticationErrorEnum_CUSTOMER_NOT_FOUND AuthenticationErrorEnum_AuthenticationError = 8
	// Client's Google Account is deleted.
	AuthenticationErrorEnum_GOOGLE_ACCOUNT_DELETED AuthenticationErrorEnum_AuthenticationError = 9
	// Google account login token in the cookie is invalid.
	AuthenticationErrorEnum_GOOGLE_ACCOUNT_COOKIE_INVALID AuthenticationErrorEnum_AuthenticationError = 10
	// A problem occurred during Google account authentication.
	AuthenticationErrorEnum_GOOGLE_ACCOUNT_AUTHENTICATION_FAILED AuthenticationErrorEnum_AuthenticationError = 25
	// The user in the google account login token does not match the UserId in
	// the cookie.
	AuthenticationErrorEnum_GOOGLE_ACCOUNT_USER_AND_ADS_USER_MISMATCH AuthenticationErrorEnum_AuthenticationError = 12
	// Login cookie is required for authentication.
	AuthenticationErrorEnum_LOGIN_COOKIE_REQUIRED AuthenticationErrorEnum_AuthenticationError = 13
	// User in the cookie is not a valid Ads user.
	AuthenticationErrorEnum_NOT_ADS_USER AuthenticationErrorEnum_AuthenticationError = 14
	// Oauth token in the header is not valid.
	AuthenticationErrorEnum_OAUTH_TOKEN_INVALID AuthenticationErrorEnum_AuthenticationError = 15
	// Oauth token in the header has expired.
	AuthenticationErrorEnum_OAUTH_TOKEN_EXPIRED AuthenticationErrorEnum_AuthenticationError = 16
	// Oauth token in the header has been disabled.
	AuthenticationErrorEnum_OAUTH_TOKEN_DISABLED AuthenticationErrorEnum_AuthenticationError = 17
	// Oauth token in the header has been revoked.
	AuthenticationErrorEnum_OAUTH_TOKEN_REVOKED AuthenticationErrorEnum_AuthenticationError = 18
	// Oauth token HTTP header is malformed.
	AuthenticationErrorEnum_OAUTH_TOKEN_HEADER_INVALID AuthenticationErrorEnum_AuthenticationError = 19
	// Login cookie is not valid.
	AuthenticationErrorEnum_LOGIN_COOKIE_INVALID AuthenticationErrorEnum_AuthenticationError = 20
	// User Id in the header is not a valid id.
	AuthenticationErrorEnum_USER_ID_INVALID AuthenticationErrorEnum_AuthenticationError = 22
	// An account administrator changed this account's authentication settings.
	// To access this Google Ads account, enable 2-Step Verification in your
	// Google account at https://www.google.com/landing/2step.
	AuthenticationErrorEnum_TWO_STEP_VERIFICATION_NOT_ENROLLED AuthenticationErrorEnum_AuthenticationError = 23
	// An account administrator changed this account's authentication settings.
	// To access this Google Ads account, enable Advanced Protection in your
	// Google account at https://landing.google.com/advancedprotection.
	AuthenticationErrorEnum_ADVANCED_PROTECTION_NOT_ENROLLED AuthenticationErrorEnum_AuthenticationError = 24
)

var AuthenticationErrorEnum_AuthenticationError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "AUTHENTICATION_ERROR",
	5:  "CLIENT_CUSTOMER_ID_INVALID",
	8:  "CUSTOMER_NOT_FOUND",
	9:  "GOOGLE_ACCOUNT_DELETED",
	10: "GOOGLE_ACCOUNT_COOKIE_INVALID",
	25: "GOOGLE_ACCOUNT_AUTHENTICATION_FAILED",
	12: "GOOGLE_ACCOUNT_USER_AND_ADS_USER_MISMATCH",
	13: "LOGIN_COOKIE_REQUIRED",
	14: "NOT_ADS_USER",
	15: "OAUTH_TOKEN_INVALID",
	16: "OAUTH_TOKEN_EXPIRED",
	17: "OAUTH_TOKEN_DISABLED",
	18: "OAUTH_TOKEN_REVOKED",
	19: "OAUTH_TOKEN_HEADER_INVALID",
	20: "LOGIN_COOKIE_INVALID",
	22: "USER_ID_INVALID",
	23: "TWO_STEP_VERIFICATION_NOT_ENROLLED",
	24: "ADVANCED_PROTECTION_NOT_ENROLLED",
}

var AuthenticationErrorEnum_AuthenticationError_value = map[string]int32{
	"UNSPECIFIED":                               0,
	"UNKNOWN":                                   1,
	"AUTHENTICATION_ERROR":                      2,
	"CLIENT_CUSTOMER_ID_INVALID":                5,
	"CUSTOMER_NOT_FOUND":                        8,
	"GOOGLE_ACCOUNT_DELETED":                    9,
	"GOOGLE_ACCOUNT_COOKIE_INVALID":             10,
	"GOOGLE_ACCOUNT_AUTHENTICATION_FAILED":      25,
	"GOOGLE_ACCOUNT_USER_AND_ADS_USER_MISMATCH": 12,
	"LOGIN_COOKIE_REQUIRED":                     13,
	"NOT_ADS_USER":                              14,
	"OAUTH_TOKEN_INVALID":                       15,
	"OAUTH_TOKEN_EXPIRED":                       16,
	"OAUTH_TOKEN_DISABLED":                      17,
	"OAUTH_TOKEN_REVOKED":                       18,
	"OAUTH_TOKEN_HEADER_INVALID":                19,
	"LOGIN_COOKIE_INVALID":                      20,
	"USER_ID_INVALID":                           22,
	"TWO_STEP_VERIFICATION_NOT_ENROLLED":        23,
	"ADVANCED_PROTECTION_NOT_ENROLLED":          24,
}

func (x AuthenticationErrorEnum_AuthenticationError) String() string {
	return proto.EnumName(AuthenticationErrorEnum_AuthenticationError_name, int32(x))
}

func (AuthenticationErrorEnum_AuthenticationError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_22e1e89d89e15d35, []int{0, 0}
}

// Container for enum describing possible authentication errors.
type AuthenticationErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticationErrorEnum) Reset()         { *m = AuthenticationErrorEnum{} }
func (m *AuthenticationErrorEnum) String() string { return proto.CompactTextString(m) }
func (*AuthenticationErrorEnum) ProtoMessage()    {}
func (*AuthenticationErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_22e1e89d89e15d35, []int{0}
}

func (m *AuthenticationErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationErrorEnum.Unmarshal(m, b)
}
func (m *AuthenticationErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationErrorEnum.Marshal(b, m, deterministic)
}
func (m *AuthenticationErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationErrorEnum.Merge(m, src)
}
func (m *AuthenticationErrorEnum) XXX_Size() int {
	return xxx_messageInfo_AuthenticationErrorEnum.Size(m)
}
func (m *AuthenticationErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.errors.AuthenticationErrorEnum_AuthenticationError", AuthenticationErrorEnum_AuthenticationError_name, AuthenticationErrorEnum_AuthenticationError_value)
	proto.RegisterType((*AuthenticationErrorEnum)(nil), "google.ads.googleads.v1.errors.AuthenticationErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/errors/authentication_error.proto", fileDescriptor_22e1e89d89e15d35)
}

var fileDescriptor_22e1e89d89e15d35 = []byte{
	// 558 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x66, 0x65, 0xfc, 0x79, 0x83, 0x19, 0x77, 0xec, 0x4f, 0x30, 0x41, 0x35, 0x21, 0xb8, 0x20,
	0x51, 0xc5, 0x15, 0xe1, 0xca, 0x8b, 0x4f, 0x3b, 0xab, 0x99, 0x1d, 0x12, 0x27, 0x43, 0xa8, 0x92,
	0x15, 0xd6, 0x2a, 0x54, 0xda, 0x92, 0xa9, 0xe9, 0xf6, 0x40, 0x5c, 0xf2, 0x28, 0x88, 0x07, 0x41,
	0xdc, 0xf2, 0x02, 0xc8, 0xc9, 0x12, 0x6d, 0x61, 0x70, 0xd5, 0x53, 0x7f, 0x3f, 0xe7, 0xf3, 0xc9,
	0x31, 0x7a, 0x97, 0xe6, 0x79, 0x7a, 0x32, 0xb5, 0x93, 0x49, 0x61, 0x57, 0xa5, 0xa9, 0x2e, 0xfa,
	0xf6, 0x74, 0x3e, 0xcf, 0xe7, 0x85, 0x9d, 0x9c, 0x2f, 0xbe, 0x4c, 0xb3, 0xc5, 0xec, 0x38, 0x59,
	0xcc, 0xf2, 0x4c, 0x97, 0xa7, 0xd6, 0xd9, 0x3c, 0x5f, 0xe4, 0x64, 0xb7, 0xe2, 0x5b, 0xc9, 0xa4,
	0xb0, 0x1a, 0xa9, 0x75, 0xd1, 0xb7, 0x2a, 0xe9, 0xce, 0xd3, 0xda, 0xfa, 0x6c, 0x66, 0x27, 0x59,
	0x96, 0x2f, 0x4a, 0x8b, 0xa2, 0x52, 0xf7, 0x7e, 0x2e, 0xa3, 0x4d, 0x7a, 0xcd, 0x1c, 0x8c, 0x0c,
	0xb2, 0xf3, 0xd3, 0xde, 0x8f, 0x65, 0xd4, 0xbd, 0x01, 0x23, 0x6b, 0x68, 0x25, 0x12, 0xa1, 0x0f,
	0x2e, 0x1f, 0x70, 0x60, 0xf8, 0x16, 0x59, 0x41, 0xf7, 0x22, 0x31, 0x12, 0xf2, 0x48, 0xe0, 0x25,
	0xb2, 0x85, 0xd6, 0x69, 0xa4, 0x0e, 0x40, 0x28, 0xee, 0x52, 0xc5, 0xa5, 0xd0, 0x10, 0x04, 0x32,
	0xc0, 0x1d, 0xb2, 0x8b, 0x76, 0x5c, 0x8f, 0x83, 0x50, 0xda, 0x8d, 0x42, 0x25, 0x0f, 0x21, 0xd0,
	0x9c, 0x69, 0x2e, 0x62, 0xea, 0x71, 0x86, 0xef, 0x90, 0x0d, 0x44, 0x1a, 0x40, 0x48, 0xa5, 0x07,
	0x32, 0x12, 0x0c, 0xdf, 0x27, 0x3b, 0x68, 0x63, 0x28, 0xe5, 0xd0, 0x03, 0x4d, 0x5d, 0x57, 0x46,
	0x42, 0x69, 0x06, 0x1e, 0x28, 0x60, 0xf8, 0x01, 0x79, 0x81, 0x9e, 0xb5, 0x30, 0x57, 0xca, 0x11,
	0x87, 0xc6, 0x16, 0x91, 0x57, 0x68, 0xaf, 0x45, 0x69, 0xe5, 0x1b, 0x50, 0xee, 0x01, 0xc3, 0xdb,
	0xe4, 0x0d, 0x7a, 0xdd, 0x62, 0x46, 0x21, 0x04, 0x9a, 0x0a, 0xa6, 0x29, 0x0b, 0xab, 0x3f, 0x87,
	0x3c, 0x3c, 0xa4, 0xca, 0x3d, 0xc0, 0xab, 0x64, 0x1b, 0x3d, 0xf1, 0xe4, 0x90, 0x8b, 0xba, 0x65,
	0x00, 0x1f, 0x22, 0x1e, 0x00, 0xc3, 0x0f, 0x09, 0x46, 0xab, 0xe6, 0x06, 0xb5, 0x0a, 0x3f, 0x22,
	0x9b, 0xa8, 0x2b, 0x4d, 0x5f, 0xad, 0xe4, 0x08, 0x44, 0x13, 0x6f, 0xad, 0x0d, 0xc0, 0x47, 0xbf,
	0xf4, 0xc0, 0x66, 0x90, 0x57, 0x01, 0xc6, 0x43, 0xba, 0x6f, 0x72, 0x3e, 0x6e, 0x4b, 0x02, 0x88,
	0xe5, 0x08, 0x18, 0x26, 0x66, 0xc2, 0x57, 0x81, 0x03, 0xa0, 0xcc, 0x0c, 0xf9, 0xb2, 0x57, 0xd7,
	0x58, 0x5e, 0x4b, 0x5c, 0x23, 0xeb, 0xa4, 0x8b, 0xd6, 0xca, 0xeb, 0x5d, 0xf9, 0x20, 0x1b, 0xe4,
	0x25, 0xea, 0xa9, 0x23, 0xa9, 0x43, 0x05, 0xbe, 0x8e, 0x21, 0xe0, 0x83, 0x7a, 0x62, 0xe6, 0x6e,
	0x20, 0x02, 0xe9, 0x99, 0x3c, 0x9b, 0x64, 0x0f, 0x3d, 0xa7, 0x2c, 0xa6, 0xc2, 0x05, 0xa6, 0xfd,
	0x40, 0x2a, 0x70, 0xff, 0x66, 0x6d, 0xed, 0xff, 0x5e, 0x42, 0xbd, 0xe3, 0xfc, 0xd4, 0xfa, 0xff,
	0xbe, 0xee, 0x6f, 0xdd, 0xb0, 0x72, 0xbe, 0xd9, 0x55, 0x7f, 0xe9, 0x13, 0xbb, 0xd4, 0xa6, 0xf9,
	0x49, 0x92, 0xa5, 0x56, 0x3e, 0x4f, 0xed, 0x74, 0x9a, 0x95, 0x9b, 0x5c, 0x3f, 0x9b, 0xb3, 0x59,
	0xf1, 0xaf, 0x57, 0xf4, 0xbe, 0xfa, 0xf9, 0xda, 0xb9, 0x3d, 0xa4, 0xf4, 0x5b, 0x67, 0x77, 0x58,
	0x99, 0xd1, 0x49, 0x61, 0x55, 0xa5, 0xa9, 0xe2, 0xbe, 0x55, 0xb6, 0x2c, 0xbe, 0xd7, 0x84, 0x31,
	0x9d, 0x14, 0xe3, 0x86, 0x30, 0x8e, 0xfb, 0xe3, 0x8a, 0xf0, 0xab, 0xd3, 0xab, 0x4e, 0x1d, 0x87,
	0x4e, 0x0a, 0xc7, 0x69, 0x28, 0x8e, 0x13, 0xf7, 0x1d, 0xa7, 0x22, 0x7d, 0xbe, 0x5b, 0xa6, 0x7b,
	0xfb, 0x27, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xb7, 0xd1, 0x26, 0xe2, 0x03, 0x00, 0x00,
}

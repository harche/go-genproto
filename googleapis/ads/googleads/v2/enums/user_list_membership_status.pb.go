// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/user_list_membership_status.proto

package enums

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

// Enum containing possible user list membership statuses.
type UserListMembershipStatusEnum_UserListMembershipStatus int32

const (
	// Not specified.
	UserListMembershipStatusEnum_UNSPECIFIED UserListMembershipStatusEnum_UserListMembershipStatus = 0
	// Used for return value only. Represents value unknown in this version.
	UserListMembershipStatusEnum_UNKNOWN UserListMembershipStatusEnum_UserListMembershipStatus = 1
	// Open status - List is accruing members and can be targeted to.
	UserListMembershipStatusEnum_OPEN UserListMembershipStatusEnum_UserListMembershipStatus = 2
	// Closed status - No new members being added. Cannot be used for targeting.
	UserListMembershipStatusEnum_CLOSED UserListMembershipStatusEnum_UserListMembershipStatus = 3
)

var UserListMembershipStatusEnum_UserListMembershipStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "OPEN",
	3: "CLOSED",
}

var UserListMembershipStatusEnum_UserListMembershipStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"OPEN":        2,
	"CLOSED":      3,
}

func (x UserListMembershipStatusEnum_UserListMembershipStatus) String() string {
	return proto.EnumName(UserListMembershipStatusEnum_UserListMembershipStatus_name, int32(x))
}

func (UserListMembershipStatusEnum_UserListMembershipStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9479d24cfe2a2de6, []int{0, 0}
}

// Membership status of this user list. Indicates whether a user list is open
// or active. Only open user lists can accumulate more users and can be used for
// targeting.
type UserListMembershipStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListMembershipStatusEnum) Reset()         { *m = UserListMembershipStatusEnum{} }
func (m *UserListMembershipStatusEnum) String() string { return proto.CompactTextString(m) }
func (*UserListMembershipStatusEnum) ProtoMessage()    {}
func (*UserListMembershipStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_9479d24cfe2a2de6, []int{0}
}

func (m *UserListMembershipStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListMembershipStatusEnum.Unmarshal(m, b)
}
func (m *UserListMembershipStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListMembershipStatusEnum.Marshal(b, m, deterministic)
}
func (m *UserListMembershipStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListMembershipStatusEnum.Merge(m, src)
}
func (m *UserListMembershipStatusEnum) XXX_Size() int {
	return xxx_messageInfo_UserListMembershipStatusEnum.Size(m)
}
func (m *UserListMembershipStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListMembershipStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_UserListMembershipStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.UserListMembershipStatusEnum_UserListMembershipStatus", UserListMembershipStatusEnum_UserListMembershipStatus_name, UserListMembershipStatusEnum_UserListMembershipStatus_value)
	proto.RegisterType((*UserListMembershipStatusEnum)(nil), "google.ads.googleads.v2.enums.UserListMembershipStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/user_list_membership_status.proto", fileDescriptor_9479d24cfe2a2de6)
}

var fileDescriptor_9479d24cfe2a2de6 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4a, 0xc3, 0x40,
	0x14, 0x86, 0x6d, 0x2a, 0x55, 0xa6, 0x0b, 0x43, 0x56, 0x22, 0xed, 0xa2, 0x3d, 0xc0, 0x04, 0xe2,
	0x6e, 0x5c, 0x48, 0xda, 0xc6, 0x52, 0xac, 0xd3, 0x40, 0x69, 0x05, 0x09, 0x94, 0xa9, 0x19, 0xc6,
	0x81, 0x66, 0x26, 0xe4, 0x4d, 0x7a, 0x20, 0x97, 0x1e, 0xc5, 0xa3, 0xb8, 0xf4, 0x04, 0x92, 0x49,
	0x9b, 0x5d, 0xdc, 0x84, 0x9f, 0xfc, 0xef, 0xff, 0xe6, 0x7f, 0x0f, 0x3d, 0x0a, 0xad, 0xc5, 0x81,
	0xfb, 0x2c, 0x05, 0xbf, 0x96, 0x95, 0x3a, 0x06, 0x3e, 0x57, 0x65, 0x06, 0x7e, 0x09, 0xbc, 0xd8,
	0x1d, 0x24, 0x98, 0x5d, 0xc6, 0xb3, 0x3d, 0x2f, 0xe0, 0x43, 0xe6, 0x3b, 0x30, 0xcc, 0x94, 0x80,
	0xf3, 0x42, 0x1b, 0xed, 0x0d, 0xeb, 0x14, 0x66, 0x29, 0xe0, 0x06, 0x80, 0x8f, 0x01, 0xb6, 0x80,
	0xbb, 0xc1, 0x99, 0x9f, 0x4b, 0x9f, 0x29, 0xa5, 0x0d, 0x33, 0x52, 0xab, 0x53, 0x78, 0xac, 0xd0,
	0x60, 0x03, 0xbc, 0x58, 0x4a, 0x30, 0x2f, 0x0d, 0x7f, 0x6d, 0xf1, 0x91, 0x2a, 0xb3, 0x31, 0x45,
	0xb7, 0x6d, 0xbe, 0x77, 0x83, 0xfa, 0x1b, 0xba, 0x8e, 0xa3, 0xe9, 0xe2, 0x69, 0x11, 0xcd, 0xdc,
	0x0b, 0xaf, 0x8f, 0xae, 0x36, 0xf4, 0x99, 0xae, 0x5e, 0xa9, 0xdb, 0xf1, 0xae, 0xd1, 0xe5, 0x2a,
	0x8e, 0xa8, 0xeb, 0x78, 0x08, 0xf5, 0xa6, 0xcb, 0xd5, 0x3a, 0x9a, 0xb9, 0xdd, 0xc9, 0x6f, 0x07,
	0x8d, 0xde, 0x75, 0x86, 0xff, 0xed, 0x3c, 0x19, 0xb6, 0xbd, 0x19, 0x57, 0xa5, 0xe3, 0xce, 0xdb,
	0xe4, 0x94, 0x17, 0xfa, 0xc0, 0x94, 0xc0, 0xba, 0x10, 0xbe, 0xe0, 0xca, 0xae, 0x74, 0x3e, 0x62,
	0x2e, 0xa1, 0xe5, 0xa6, 0x0f, 0xf6, 0xfb, 0xe9, 0x74, 0xe7, 0x61, 0xf8, 0xe5, 0x0c, 0xe7, 0x35,
	0x2a, 0x4c, 0x01, 0xd7, 0xb2, 0x52, 0xdb, 0x00, 0x57, 0xfb, 0xc3, 0xf7, 0xd9, 0x4f, 0xc2, 0x14,
	0x92, 0xc6, 0x4f, 0xb6, 0x41, 0x62, 0xfd, 0x1f, 0x67, 0x54, 0xff, 0x24, 0x24, 0x4c, 0x81, 0x90,
	0x66, 0x82, 0x90, 0x6d, 0x40, 0x88, 0x9d, 0xd9, 0xf7, 0x6c, 0xb1, 0xfb, 0xbf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x19, 0x1e, 0xee, 0x12, 0xeb, 0x01, 0x00, 0x00,
}

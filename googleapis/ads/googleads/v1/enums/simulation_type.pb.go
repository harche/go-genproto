// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/simulation_type.proto

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

// Enum describing the field a simulation modifies.
type SimulationTypeEnum_SimulationType int32

const (
	// Not specified.
	SimulationTypeEnum_UNSPECIFIED SimulationTypeEnum_SimulationType = 0
	// Used for return value only. Represents value unknown in this version.
	SimulationTypeEnum_UNKNOWN SimulationTypeEnum_SimulationType = 1
	// The simulation is for a cpc bid.
	SimulationTypeEnum_CPC_BID SimulationTypeEnum_SimulationType = 2
	// The simulation is for a cpv bid.
	SimulationTypeEnum_CPV_BID SimulationTypeEnum_SimulationType = 3
	// The simulation is for a cpa target.
	SimulationTypeEnum_TARGET_CPA SimulationTypeEnum_SimulationType = 4
	// The simulation is for a bid modifier.
	SimulationTypeEnum_BID_MODIFIER SimulationTypeEnum_SimulationType = 5
)

var SimulationTypeEnum_SimulationType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CPC_BID",
	3: "CPV_BID",
	4: "TARGET_CPA",
	5: "BID_MODIFIER",
}

var SimulationTypeEnum_SimulationType_value = map[string]int32{
	"UNSPECIFIED":  0,
	"UNKNOWN":      1,
	"CPC_BID":      2,
	"CPV_BID":      3,
	"TARGET_CPA":   4,
	"BID_MODIFIER": 5,
}

func (x SimulationTypeEnum_SimulationType) String() string {
	return proto.EnumName(SimulationTypeEnum_SimulationType_name, int32(x))
}

func (SimulationTypeEnum_SimulationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bfbb0271e7521728, []int{0, 0}
}

// Container for enum describing the field a simulation modifies.
type SimulationTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimulationTypeEnum) Reset()         { *m = SimulationTypeEnum{} }
func (m *SimulationTypeEnum) String() string { return proto.CompactTextString(m) }
func (*SimulationTypeEnum) ProtoMessage()    {}
func (*SimulationTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfbb0271e7521728, []int{0}
}

func (m *SimulationTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimulationTypeEnum.Unmarshal(m, b)
}
func (m *SimulationTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimulationTypeEnum.Marshal(b, m, deterministic)
}
func (m *SimulationTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimulationTypeEnum.Merge(m, src)
}
func (m *SimulationTypeEnum) XXX_Size() int {
	return xxx_messageInfo_SimulationTypeEnum.Size(m)
}
func (m *SimulationTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_SimulationTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_SimulationTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v1.enums.SimulationTypeEnum_SimulationType", SimulationTypeEnum_SimulationType_name, SimulationTypeEnum_SimulationType_value)
	proto.RegisterType((*SimulationTypeEnum)(nil), "google.ads.googleads.v1.enums.SimulationTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/simulation_type.proto", fileDescriptor_bfbb0271e7521728)
}

var fileDescriptor_bfbb0271e7521728 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0xfd, 0x92, 0x7e, 0x2a, 0x4c, 0xa5, 0x86, 0xb8, 0x13, 0xbb, 0x68, 0x1f, 0x60, 0x86, 0xd0,
	0xdd, 0xb8, 0x9a, 0xfc, 0x58, 0x82, 0x98, 0x86, 0xfe, 0x44, 0x90, 0x40, 0x88, 0x26, 0x0c, 0x91,
	0x66, 0x26, 0x74, 0xd2, 0x42, 0x77, 0x3e, 0x8b, 0x4b, 0x1f, 0xc5, 0x27, 0x11, 0x9f, 0x42, 0x66,
	0xc6, 0x06, 0xba, 0xd0, 0xcd, 0x70, 0xee, 0xdc, 0x73, 0x0e, 0xe7, 0x1e, 0x30, 0xa1, 0x9c, 0xd3,
	0x75, 0x89, 0xf2, 0x42, 0x20, 0x0d, 0x25, 0xda, 0x39, 0xa8, 0x64, 0xdb, 0x5a, 0x20, 0x51, 0xd5,
	0xdb, 0x75, 0xde, 0x56, 0x9c, 0x65, 0xed, 0xbe, 0x29, 0x61, 0xb3, 0xe1, 0x2d, 0xb7, 0x87, 0x9a,
	0x09, 0xf3, 0x42, 0xc0, 0x4e, 0x04, 0x77, 0x0e, 0x54, 0xa2, 0xab, 0xeb, 0x83, 0x67, 0x53, 0xa1,
	0x9c, 0x31, 0xde, 0x2a, 0x03, 0xa1, 0xc5, 0xe3, 0x57, 0x03, 0xd8, 0x8b, 0xce, 0x76, 0xb9, 0x6f,
	0xca, 0x80, 0x6d, 0xeb, 0xf1, 0x0b, 0x18, 0x1c, 0xff, 0xda, 0x17, 0xa0, 0xbf, 0x8a, 0x16, 0x71,
	0xe0, 0x85, 0xb7, 0x61, 0xe0, 0x5b, 0xff, 0xec, 0x3e, 0x38, 0x5b, 0x45, 0x77, 0xd1, 0xec, 0x21,
	0xb2, 0x0c, 0x39, 0x78, 0xb1, 0x97, 0xb9, 0xa1, 0x6f, 0x99, 0x7a, 0x48, 0xd4, 0xd0, 0xb3, 0x07,
	0x00, 0x2c, 0xc9, 0x7c, 0x1a, 0x2c, 0x33, 0x2f, 0x26, 0xd6, 0x7f, 0xdb, 0x02, 0xe7, 0x6e, 0xe8,
	0x67, 0xf7, 0x33, 0x5f, 0x1a, 0xcd, 0xad, 0x13, 0xf7, 0xd3, 0x00, 0xa3, 0x67, 0x5e, 0xc3, 0x3f,
	0xcf, 0x70, 0x2f, 0x8f, 0xf3, 0xc4, 0x32, 0x7d, 0x6c, 0x3c, 0xba, 0x3f, 0x2a, 0xca, 0xd7, 0x39,
	0xa3, 0x90, 0x6f, 0x28, 0xa2, 0x25, 0x53, 0xb7, 0x1d, 0x1a, 0x6c, 0x2a, 0xf1, 0x4b, 0xa1, 0x37,
	0xea, 0x7d, 0x33, 0x7b, 0x53, 0x42, 0xde, 0xcd, 0xe1, 0x54, 0x5b, 0x91, 0x42, 0x40, 0x0d, 0x25,
	0x4a, 0x1c, 0x28, 0x1b, 0x11, 0x1f, 0x87, 0x7d, 0x4a, 0x0a, 0x91, 0x76, 0xfb, 0x34, 0x71, 0x52,
	0xb5, 0xff, 0x32, 0x47, 0xfa, 0x13, 0x63, 0x52, 0x08, 0x8c, 0x3b, 0x06, 0xc6, 0x89, 0x83, 0xb1,
	0xe2, 0x3c, 0x9d, 0xaa, 0x60, 0x93, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x3b, 0x1b, 0xfa,
	0xe8, 0x01, 0x00, 0x00,
}

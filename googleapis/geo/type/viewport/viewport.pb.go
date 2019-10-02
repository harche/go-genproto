// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/geo/type/viewport.proto

package viewport

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
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

// A latitude-longitude viewport, represented as two diagonally opposite `low`
// and `high` points. A viewport is considered a closed region, i.e. it includes
// its boundary. The latitude bounds must range between -90 to 90 degrees
// inclusive, and the longitude bounds must range between -180 to 180 degrees
// inclusive. Various cases include:
//
//  - If `low` = `high`, the viewport consists of that single point.
//
//  - If `low.longitude` > `high.longitude`, the longitude range is inverted
//    (the viewport crosses the 180 degree longitude line).
//
//  - If `low.longitude` = -180 degrees and `high.longitude` = 180 degrees,
//    the viewport includes all longitudes.
//
//  - If `low.longitude` = 180 degrees and `high.longitude` = -180 degrees,
//    the longitude range is empty.
//
//  - If `low.latitude` > `high.latitude`, the latitude range is empty.
//
// Both `low` and `high` must be populated, and the represented box cannot be
// empty (as specified by the definitions above). An empty viewport will result
// in an error.
//
// For example, this viewport fully encloses New York City:
//
// {
//     "low": {
//         "latitude": 40.477398,
//         "longitude": -74.259087
//     },
//     "high": {
//         "latitude": 40.91618,
//         "longitude": -73.70018
//     }
// }
type Viewport struct {
	// Required. The low point of the viewport.
	Low *latlng.LatLng `protobuf:"bytes,1,opt,name=low,proto3" json:"low,omitempty"`
	// Required. The high point of the viewport.
	High                 *latlng.LatLng `protobuf:"bytes,2,opt,name=high,proto3" json:"high,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Viewport) Reset()         { *m = Viewport{} }
func (m *Viewport) String() string { return proto.CompactTextString(m) }
func (*Viewport) ProtoMessage()    {}
func (*Viewport) Descriptor() ([]byte, []int) {
	return fileDescriptor_52240f0f30767909, []int{0}
}

func (m *Viewport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Viewport.Unmarshal(m, b)
}
func (m *Viewport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Viewport.Marshal(b, m, deterministic)
}
func (m *Viewport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Viewport.Merge(m, src)
}
func (m *Viewport) XXX_Size() int {
	return xxx_messageInfo_Viewport.Size(m)
}
func (m *Viewport) XXX_DiscardUnknown() {
	xxx_messageInfo_Viewport.DiscardUnknown(m)
}

var xxx_messageInfo_Viewport proto.InternalMessageInfo

func (m *Viewport) GetLow() *latlng.LatLng {
	if m != nil {
		return m.Low
	}
	return nil
}

func (m *Viewport) GetHigh() *latlng.LatLng {
	if m != nil {
		return m.High
	}
	return nil
}

func init() {
	proto.RegisterType((*Viewport)(nil), "google.geo.type.Viewport")
}

func init() { proto.RegisterFile("google/geo/type/viewport.proto", fileDescriptor_52240f0f30767909) }

var fileDescriptor_52240f0f30767909 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4f, 0xcd, 0xd7, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x2f, 0xcb, 0x4c, 0x2d,
	0x2f, 0xc8, 0x2f, 0x2a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x87, 0xc8, 0xeb, 0xa5,
	0xa7, 0xe6, 0xeb, 0x81, 0xe4, 0xa5, 0x24, 0xa0, 0x1a, 0xc0, 0x8a, 0x73, 0x12, 0x4b, 0x72, 0xf2,
	0xd2, 0x21, 0x4a, 0x95, 0xa2, 0xb8, 0x38, 0xc2, 0xa0, 0x9a, 0x85, 0x54, 0xb9, 0x98, 0x73, 0xf2,
	0xcb, 0x25, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x84, 0xf5, 0xa0, 0x86, 0x80, 0xf4, 0xe8, 0xf9,
	0x24, 0x96, 0xf8, 0xe4, 0xa5, 0x07, 0x81, 0xe4, 0x85, 0xd4, 0xb9, 0x58, 0x32, 0x32, 0xd3, 0x33,
	0x24, 0x98, 0x70, 0xab, 0x03, 0x2b, 0x70, 0xca, 0xe7, 0x12, 0x4e, 0xce, 0xcf, 0xd5, 0x43, 0x73,
	0x8c, 0x13, 0x2f, 0xcc, 0xc2, 0x00, 0x90, 0x0b, 0x02, 0x18, 0xa3, 0x1c, 0x60, 0x2a, 0xf2, 0x73,
	0x12, 0xf3, 0xd2, 0xf5, 0xf2, 0x8b, 0xd2, 0xf5, 0xd3, 0x53, 0xf3, 0xc0, 0xee, 0xd3, 0x87, 0x48,
	0x25, 0x16, 0x64, 0x16, 0x63, 0xfa, 0xd6, 0x1a, 0xc6, 0x58, 0xc4, 0xc4, 0xe2, 0xee, 0x1e, 0x12,
	0x90, 0xc4, 0x06, 0xd6, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x9c, 0x8b, 0x3a, 0x20,
	0x01, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/phishingprotection/v1beta1/phishingprotection.proto

package phishingprotection

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// The ReportPhishing request message.
type ReportPhishingRequest struct {
	// Required. The name of the project for which the report will be created,
	// in the format "projects/{project_number}".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The URI that is being reported for phishing content to be analyzed.
	Uri                  string   `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportPhishingRequest) Reset()         { *m = ReportPhishingRequest{} }
func (m *ReportPhishingRequest) String() string { return proto.CompactTextString(m) }
func (*ReportPhishingRequest) ProtoMessage()    {}
func (*ReportPhishingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa2769de3b7011f5, []int{0}
}

func (m *ReportPhishingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportPhishingRequest.Unmarshal(m, b)
}
func (m *ReportPhishingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportPhishingRequest.Marshal(b, m, deterministic)
}
func (m *ReportPhishingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportPhishingRequest.Merge(m, src)
}
func (m *ReportPhishingRequest) XXX_Size() int {
	return xxx_messageInfo_ReportPhishingRequest.Size(m)
}
func (m *ReportPhishingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportPhishingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReportPhishingRequest proto.InternalMessageInfo

func (m *ReportPhishingRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ReportPhishingRequest) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

// The ReportPhishing (empty) response message.
type ReportPhishingResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportPhishingResponse) Reset()         { *m = ReportPhishingResponse{} }
func (m *ReportPhishingResponse) String() string { return proto.CompactTextString(m) }
func (*ReportPhishingResponse) ProtoMessage()    {}
func (*ReportPhishingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa2769de3b7011f5, []int{1}
}

func (m *ReportPhishingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportPhishingResponse.Unmarshal(m, b)
}
func (m *ReportPhishingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportPhishingResponse.Marshal(b, m, deterministic)
}
func (m *ReportPhishingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportPhishingResponse.Merge(m, src)
}
func (m *ReportPhishingResponse) XXX_Size() int {
	return xxx_messageInfo_ReportPhishingResponse.Size(m)
}
func (m *ReportPhishingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportPhishingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReportPhishingResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ReportPhishingRequest)(nil), "google.cloud.phishingprotection.v1beta1.ReportPhishingRequest")
	proto.RegisterType((*ReportPhishingResponse)(nil), "google.cloud.phishingprotection.v1beta1.ReportPhishingResponse")
}

func init() {
	proto.RegisterFile("google/cloud/phishingprotection/v1beta1/phishingprotection.proto", fileDescriptor_fa2769de3b7011f5)
}

var fileDescriptor_fa2769de3b7011f5 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcf, 0x6b, 0x13, 0x41,
	0x14, 0x66, 0xb7, 0x52, 0x70, 0x0e, 0x1e, 0x16, 0x6a, 0x63, 0x10, 0xac, 0x01, 0xa9, 0xd4, 0x76,
	0x86, 0xd8, 0x5b, 0xc4, 0x1f, 0x9b, 0x0a, 0x3d, 0x78, 0x59, 0x22, 0x16, 0x94, 0x80, 0x4c, 0xb6,
	0xaf, 0xbb, 0x23, 0x9b, 0x79, 0xe3, 0xcc, 0x6c, 0x72, 0x28, 0x5e, 0xfc, 0x17, 0xfc, 0x0f, 0x3c,
	0xfa, 0xa7, 0xf4, 0x68, 0x0f, 0x82, 0xa7, 0x1e, 0xfc, 0x23, 0xc4, 0x93, 0xec, 0xce, 0x24, 0xa6,
	0xc9, 0x8a, 0xd2, 0xdb, 0xec, 0x7e, 0xdf, 0xfb, 0xbe, 0xf7, 0xbe, 0x79, 0x43, 0x9e, 0x65, 0x88,
	0x59, 0x01, 0x2c, 0x2d, 0xb0, 0x3c, 0x66, 0x2a, 0x17, 0x26, 0x17, 0x32, 0x53, 0x1a, 0x2d, 0xa4,
	0x56, 0xa0, 0x64, 0x93, 0xee, 0x08, 0x2c, 0xef, 0x36, 0x40, 0xb4, 0x3a, 0x62, 0xb4, 0xed, 0x14,
	0x68, 0xad, 0x40, 0x1b, 0x68, 0x5e, 0xa1, 0x7d, 0xdb, 0x5b, 0x71, 0x25, 0x18, 0x97, 0x12, 0x2d,
	0xaf, 0x60, 0xe3, 0x64, 0xda, 0x9b, 0x0b, 0x68, 0x5a, 0x08, 0x90, 0xd6, 0x03, 0x77, 0x16, 0x80,
	0x13, 0x01, 0xc5, 0xf1, 0xdb, 0x11, 0xe4, 0x7c, 0x22, 0x50, 0x7b, 0xc2, 0xad, 0x05, 0x82, 0x06,
	0x83, 0xa5, 0x4e, 0xc1, 0x41, 0x9d, 0x53, 0xb2, 0x31, 0x00, 0x85, 0xda, 0x26, 0xbe, 0xad, 0x01,
	0xbc, 0x2f, 0xc1, 0xd8, 0xe8, 0x05, 0x59, 0x57, 0x5c, 0x83, 0xb4, 0xad, 0x60, 0x2b, 0xb8, 0x7f,
	0xbd, 0xbf, 0x7f, 0x11, 0x87, 0xbf, 0xe2, 0x3d, 0xf2, 0xa0, 0x9e, 0x62, 0xa6, 0x32, 0xe6, 0x92,
	0x67, 0xa0, 0xa9, 0x33, 0xe0, 0x4a, 0x18, 0x9a, 0xe2, 0x98, 0x25, 0x1a, 0xdf, 0x41, 0x6a, 0x07,
	0x5e, 0x22, 0xda, 0x20, 0x6b, 0xa5, 0x16, 0xad, 0xb0, 0x56, 0x5a, 0xbb, 0x88, 0xc3, 0x41, 0xf5,
	0xdd, 0x69, 0x91, 0x9b, 0xcb, 0xe6, 0x46, 0xa1, 0x34, 0xf0, 0xf0, 0x3c, 0x24, 0x5b, 0xb3, 0x9f,
	0xc9, 0x3c, 0xa8, 0x97, 0xa0, 0x27, 0x22, 0x85, 0xa3, 0x6e, 0xbf, 0x8a, 0x2b, 0xfa, 0x16, 0x90,
	0x1b, 0x97, 0xeb, 0xa3, 0x27, 0xf4, 0x3f, 0xb3, 0xa6, 0x8d, 0x53, 0xb7, 0x9f, 0x5e, 0xb9, 0xde,
	0x35, 0xde, 0x79, 0xfe, 0x3d, 0x26, 0x6e, 0xe8, 0xdd, 0x52, 0x8b, 0x8f, 0xe7, 0x3f, 0x3e, 0x85,
	0xdd, 0xce, 0xee, 0x7c, 0x47, 0x4e, 0x1d, 0xf6, 0x58, 0xb9, 0x80, 0x0c, 0xdb, 0xf9, 0x30, 0xdf,
	0x9b, 0x9e, 0xae, 0x25, 0x7b, 0xc1, 0x4e, 0xfb, 0xd5, 0x59, 0x7c, 0xb7, 0xc1, 0xfb, 0x72, 0xca,
	0x5f, 0x63, 0x9a, 0x5b, 0xab, 0x4c, 0x8f, 0xb1, 0xe9, 0x74, 0xba, 0x7c, 0x05, 0xbc, 0xb4, 0xb9,
	0xdb, 0xd9, 0x3d, 0x55, 0x70, 0x7b, 0x82, 0x7a, 0xdc, 0xff, 0x19, 0x90, 0x7b, 0x29, 0x8e, 0x67,
	0x33, 0xfe, 0x7d, 0xba, 0xfe, 0xe6, 0x6a, 0xf8, 0xd5, 0x09, 0x93, 0xe0, 0xcd, 0x6b, 0x5f, 0x9d,
	0x61, 0xc1, 0x65, 0x46, 0x51, 0x67, 0x2c, 0x03, 0x59, 0x6f, 0x13, 0xfb, 0xd3, 0xc3, 0x3f, 0x9f,
	0xcb, 0xa3, 0x55, 0xe8, 0x73, 0x78, 0xed, 0xf0, 0x20, 0x49, 0xbe, 0x84, 0xdb, 0x87, 0xce, 0xe2,
	0xa0, 0xbe, 0x84, 0xd5, 0x56, 0xa8, 0xdf, 0x80, 0xb3, 0x19, 0x73, 0x58, 0x33, 0x87, 0xab, 0xcc,
	0xe1, 0x91, 0x73, 0x1b, 0xad, 0xd7, 0x0d, 0xee, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x25, 0xcc,
	0xbb, 0x4b, 0xce, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PhishingProtectionServiceV1Beta1Client is the client API for PhishingProtectionServiceV1Beta1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PhishingProtectionServiceV1Beta1Client interface {
	// Reports a URI suspected of containing phishing content to be reviewed. Once
	// the report review is complete, its result can be found in the Cloud
	// Security Command Center findings dashboard for Phishing Protection. If the
	// result verifies the existence of malicious phishing content, the site will
	// be added the to [Google's Social Engineering
	// lists](https://support.google.com/webmasters/answer/6350487/) in order to
	// protect users that could get exposed to this threat in the future.
	ReportPhishing(ctx context.Context, in *ReportPhishingRequest, opts ...grpc.CallOption) (*ReportPhishingResponse, error)
}

type phishingProtectionServiceV1Beta1Client struct {
	cc *grpc.ClientConn
}

func NewPhishingProtectionServiceV1Beta1Client(cc *grpc.ClientConn) PhishingProtectionServiceV1Beta1Client {
	return &phishingProtectionServiceV1Beta1Client{cc}
}

func (c *phishingProtectionServiceV1Beta1Client) ReportPhishing(ctx context.Context, in *ReportPhishingRequest, opts ...grpc.CallOption) (*ReportPhishingResponse, error) {
	out := new(ReportPhishingResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.phishingprotection.v1beta1.PhishingProtectionServiceV1Beta1/ReportPhishing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PhishingProtectionServiceV1Beta1Server is the server API for PhishingProtectionServiceV1Beta1 service.
type PhishingProtectionServiceV1Beta1Server interface {
	// Reports a URI suspected of containing phishing content to be reviewed. Once
	// the report review is complete, its result can be found in the Cloud
	// Security Command Center findings dashboard for Phishing Protection. If the
	// result verifies the existence of malicious phishing content, the site will
	// be added the to [Google's Social Engineering
	// lists](https://support.google.com/webmasters/answer/6350487/) in order to
	// protect users that could get exposed to this threat in the future.
	ReportPhishing(context.Context, *ReportPhishingRequest) (*ReportPhishingResponse, error)
}

// UnimplementedPhishingProtectionServiceV1Beta1Server can be embedded to have forward compatible implementations.
type UnimplementedPhishingProtectionServiceV1Beta1Server struct {
}

func (*UnimplementedPhishingProtectionServiceV1Beta1Server) ReportPhishing(ctx context.Context, req *ReportPhishingRequest) (*ReportPhishingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportPhishing not implemented")
}

func RegisterPhishingProtectionServiceV1Beta1Server(s *grpc.Server, srv PhishingProtectionServiceV1Beta1Server) {
	s.RegisterService(&_PhishingProtectionServiceV1Beta1_serviceDesc, srv)
}

func _PhishingProtectionServiceV1Beta1_ReportPhishing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportPhishingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhishingProtectionServiceV1Beta1Server).ReportPhishing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.phishingprotection.v1beta1.PhishingProtectionServiceV1Beta1/ReportPhishing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhishingProtectionServiceV1Beta1Server).ReportPhishing(ctx, req.(*ReportPhishingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PhishingProtectionServiceV1Beta1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.phishingprotection.v1beta1.PhishingProtectionServiceV1Beta1",
	HandlerType: (*PhishingProtectionServiceV1Beta1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReportPhishing",
			Handler:    _PhishingProtectionServiceV1Beta1_ReportPhishing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/phishingprotection/v1beta1/phishingprotection.proto",
}

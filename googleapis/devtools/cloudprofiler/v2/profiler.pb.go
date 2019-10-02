// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/devtools/cloudprofiler/v2/profiler.proto

package cloudprofiler

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

// ProfileType is type of profiling data.
// NOTE: the enumeration member names are used (in lowercase) as unique string
// identifiers of profile types, so they must not be renamed.
type ProfileType int32

const (
	// Unspecified profile type.
	ProfileType_PROFILE_TYPE_UNSPECIFIED ProfileType = 0
	// Thread CPU time sampling.
	ProfileType_CPU ProfileType = 1
	// Wallclock time sampling. More expensive as stops all threads.
	ProfileType_WALL ProfileType = 2
	// In-use heap profile. Represents a snapshot of the allocations that are
	// live at the time of the profiling.
	ProfileType_HEAP ProfileType = 3
	// Single-shot collection of all thread stacks.
	ProfileType_THREADS ProfileType = 4
	// Synchronization contention profile.
	ProfileType_CONTENTION ProfileType = 5
	// Peak heap profile.
	ProfileType_PEAK_HEAP ProfileType = 6
	// Heap allocation profile. It represents the aggregation of all allocations
	// made over the duration of the profile. All allocations are included,
	// including those that might have been freed by the end of the profiling
	// interval. The profile is in particular useful for garbage collecting
	// languages to understand which parts of the code create most of the garbage
	// collection pressure to see if those can be optimized.
	ProfileType_HEAP_ALLOC ProfileType = 7
)

var ProfileType_name = map[int32]string{
	0: "PROFILE_TYPE_UNSPECIFIED",
	1: "CPU",
	2: "WALL",
	3: "HEAP",
	4: "THREADS",
	5: "CONTENTION",
	6: "PEAK_HEAP",
	7: "HEAP_ALLOC",
}

var ProfileType_value = map[string]int32{
	"PROFILE_TYPE_UNSPECIFIED": 0,
	"CPU":                      1,
	"WALL":                     2,
	"HEAP":                     3,
	"THREADS":                  4,
	"CONTENTION":               5,
	"PEAK_HEAP":                6,
	"HEAP_ALLOC":               7,
}

func (x ProfileType) String() string {
	return proto.EnumName(ProfileType_name, int32(x))
}

func (ProfileType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_74a10a5851161432, []int{0}
}

// CreateProfileRequest describes a profile resource online creation request.
// The deployment field must be populated. The profile_type specifies the list
// of profile types supported by the agent. The creation call will hang until a
// profile of one of these types needs to be collected.
type CreateProfileRequest struct {
	// Parent project to create the profile in.
	Parent string `protobuf:"bytes,4,opt,name=parent,proto3" json:"parent,omitempty"`
	// Deployment details.
	Deployment *Deployment `protobuf:"bytes,1,opt,name=deployment,proto3" json:"deployment,omitempty"`
	// One or more profile types that the agent is capable of providing.
	ProfileType          []ProfileType `protobuf:"varint,2,rep,packed,name=profile_type,json=profileType,proto3,enum=google.devtools.cloudprofiler.v2.ProfileType" json:"profile_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CreateProfileRequest) Reset()         { *m = CreateProfileRequest{} }
func (m *CreateProfileRequest) String() string { return proto.CompactTextString(m) }
func (*CreateProfileRequest) ProtoMessage()    {}
func (*CreateProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a10a5851161432, []int{0}
}

func (m *CreateProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProfileRequest.Unmarshal(m, b)
}
func (m *CreateProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProfileRequest.Marshal(b, m, deterministic)
}
func (m *CreateProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProfileRequest.Merge(m, src)
}
func (m *CreateProfileRequest) XXX_Size() int {
	return xxx_messageInfo_CreateProfileRequest.Size(m)
}
func (m *CreateProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProfileRequest proto.InternalMessageInfo

func (m *CreateProfileRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateProfileRequest) GetDeployment() *Deployment {
	if m != nil {
		return m.Deployment
	}
	return nil
}

func (m *CreateProfileRequest) GetProfileType() []ProfileType {
	if m != nil {
		return m.ProfileType
	}
	return nil
}

// CreateOfflineProfileRequest describes a profile resource offline creation
// request. Profile field must be set.
type CreateOfflineProfileRequest struct {
	// Parent project to create the profile in.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Contents of the profile to create.
	Profile              *Profile `protobuf:"bytes,2,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOfflineProfileRequest) Reset()         { *m = CreateOfflineProfileRequest{} }
func (m *CreateOfflineProfileRequest) String() string { return proto.CompactTextString(m) }
func (*CreateOfflineProfileRequest) ProtoMessage()    {}
func (*CreateOfflineProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a10a5851161432, []int{1}
}

func (m *CreateOfflineProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOfflineProfileRequest.Unmarshal(m, b)
}
func (m *CreateOfflineProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOfflineProfileRequest.Marshal(b, m, deterministic)
}
func (m *CreateOfflineProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOfflineProfileRequest.Merge(m, src)
}
func (m *CreateOfflineProfileRequest) XXX_Size() int {
	return xxx_messageInfo_CreateOfflineProfileRequest.Size(m)
}
func (m *CreateOfflineProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOfflineProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOfflineProfileRequest proto.InternalMessageInfo

func (m *CreateOfflineProfileRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateOfflineProfileRequest) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

// UpdateProfileRequest contains the profile to update.
type UpdateProfileRequest struct {
	// Profile to update
	Profile *Profile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	// Field mask used to specify the fields to be overwritten. Currently only
	// profile_bytes and labels fields are supported by UpdateProfile, so only
	// those fields can be specified in the mask. When no mask is provided, all
	// fields are overwritten.
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateProfileRequest) Reset()         { *m = UpdateProfileRequest{} }
func (m *UpdateProfileRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateProfileRequest) ProtoMessage()    {}
func (*UpdateProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a10a5851161432, []int{2}
}

func (m *UpdateProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateProfileRequest.Unmarshal(m, b)
}
func (m *UpdateProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateProfileRequest.Marshal(b, m, deterministic)
}
func (m *UpdateProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateProfileRequest.Merge(m, src)
}
func (m *UpdateProfileRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateProfileRequest.Size(m)
}
func (m *UpdateProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateProfileRequest proto.InternalMessageInfo

func (m *UpdateProfileRequest) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func (m *UpdateProfileRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// Profile resource.
type Profile struct {
	// Output only. Opaque, server-assigned, unique ID for this profile.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Type of profile.
	// For offline mode, this must be specified when creating the profile. For
	// online mode it is assigned and returned by the server.
	ProfileType ProfileType `protobuf:"varint,2,opt,name=profile_type,json=profileType,proto3,enum=google.devtools.cloudprofiler.v2.ProfileType" json:"profile_type,omitempty"`
	// Deployment this profile corresponds to.
	Deployment *Deployment `protobuf:"bytes,3,opt,name=deployment,proto3" json:"deployment,omitempty"`
	// Duration of the profiling session.
	// Input (for the offline mode) or output (for the online mode).
	// The field represents requested profiling duration. It may slightly differ
	// from the effective profiling duration, which is recorded in the profile
	// data, in case the profiling can't be stopped immediately (e.g. in case
	// stopping the profiling is handled asynchronously).
	Duration *duration.Duration `protobuf:"bytes,4,opt,name=duration,proto3" json:"duration,omitempty"`
	// Input only. Profile bytes, as a gzip compressed serialized proto, the
	// format is https://github.com/google/pprof/blob/master/proto/profile.proto.
	ProfileBytes []byte `protobuf:"bytes,5,opt,name=profile_bytes,json=profileBytes,proto3" json:"profile_bytes,omitempty"`
	// Input only. Labels associated to this specific profile. These labels will
	// get merged with the deployment labels for the final data set.  See
	// documentation on deployment labels for validation rules and limits.
	Labels               map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a10a5851161432, []int{3}
}

func (m *Profile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Profile.Unmarshal(m, b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
}
func (m *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(m, src)
}
func (m *Profile) XXX_Size() int {
	return xxx_messageInfo_Profile.Size(m)
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

func (m *Profile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Profile) GetProfileType() ProfileType {
	if m != nil {
		return m.ProfileType
	}
	return ProfileType_PROFILE_TYPE_UNSPECIFIED
}

func (m *Profile) GetDeployment() *Deployment {
	if m != nil {
		return m.Deployment
	}
	return nil
}

func (m *Profile) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *Profile) GetProfileBytes() []byte {
	if m != nil {
		return m.ProfileBytes
	}
	return nil
}

func (m *Profile) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// Deployment contains the deployment identification information.
type Deployment struct {
	// Project ID is the ID of a cloud project.
	// Validation regex: `^[a-z][-a-z0-9:.]{4,61}[a-z0-9]$`.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Target is the service name used to group related deployments:
	// * Service name for GAE Flex / Standard.
	// * Cluster and container name for GKE.
	// * User-specified string for direct GCE profiling (e.g. Java).
	// * Job name for Dataflow.
	// Validation regex: `^[a-z]([-a-z0-9_.]{0,253}[a-z0-9])?$`.
	Target string `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	// Labels identify the deployment within the user universe and same target.
	// Validation regex for label names: `^[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$`.
	// Value for an individual label must be <= 512 bytes, the total
	// size of all label names and values must be <= 1024 bytes.
	//
	// Label named "language" can be used to record the programming language of
	// the profiled deployment. The standard choices for the value include "java",
	// "go", "python", "ruby", "nodejs", "php", "dotnet".
	//
	// For deployments running on Google Cloud Platform, "zone" or "region" label
	// should be present describing the deployment location. An example of a zone
	// is "us-central1-a", an example of a region is "us-central1" or
	// "us-central".
	Labels               map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Deployment) Reset()         { *m = Deployment{} }
func (m *Deployment) String() string { return proto.CompactTextString(m) }
func (*Deployment) ProtoMessage()    {}
func (*Deployment) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a10a5851161432, []int{4}
}

func (m *Deployment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deployment.Unmarshal(m, b)
}
func (m *Deployment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deployment.Marshal(b, m, deterministic)
}
func (m *Deployment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deployment.Merge(m, src)
}
func (m *Deployment) XXX_Size() int {
	return xxx_messageInfo_Deployment.Size(m)
}
func (m *Deployment) XXX_DiscardUnknown() {
	xxx_messageInfo_Deployment.DiscardUnknown(m)
}

var xxx_messageInfo_Deployment proto.InternalMessageInfo

func (m *Deployment) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *Deployment) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *Deployment) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.devtools.cloudprofiler.v2.ProfileType", ProfileType_name, ProfileType_value)
	proto.RegisterType((*CreateProfileRequest)(nil), "google.devtools.cloudprofiler.v2.CreateProfileRequest")
	proto.RegisterType((*CreateOfflineProfileRequest)(nil), "google.devtools.cloudprofiler.v2.CreateOfflineProfileRequest")
	proto.RegisterType((*UpdateProfileRequest)(nil), "google.devtools.cloudprofiler.v2.UpdateProfileRequest")
	proto.RegisterType((*Profile)(nil), "google.devtools.cloudprofiler.v2.Profile")
	proto.RegisterMapType((map[string]string)(nil), "google.devtools.cloudprofiler.v2.Profile.LabelsEntry")
	proto.RegisterType((*Deployment)(nil), "google.devtools.cloudprofiler.v2.Deployment")
	proto.RegisterMapType((map[string]string)(nil), "google.devtools.cloudprofiler.v2.Deployment.LabelsEntry")
}

func init() {
	proto.RegisterFile("google/devtools/cloudprofiler/v2/profiler.proto", fileDescriptor_74a10a5851161432)
}

var fileDescriptor_74a10a5851161432 = []byte{
	// 786 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x5f, 0x6f, 0xda, 0x56,
	0x14, 0xdf, 0xc5, 0x04, 0xc2, 0x71, 0xc8, 0xac, 0xab, 0x68, 0x62, 0x2c, 0xdb, 0x2c, 0x4f, 0x93,
	0x18, 0xdb, 0x6c, 0xc9, 0x51, 0xa6, 0xfc, 0x51, 0x34, 0x11, 0x70, 0x14, 0x34, 0x02, 0x96, 0x43,
	0x34, 0x6d, 0x2f, 0xc8, 0xe0, 0x0b, 0x72, 0x63, 0x6c, 0xd7, 0x36, 0x48, 0xb4, 0xca, 0x4b, 0xd5,
	0x6f, 0xd0, 0x97, 0xbe, 0xf5, 0xa1, 0x0f, 0x7d, 0xea, 0x87, 0xe8, 0x17, 0xa8, 0x2a, 0xf5, 0x2b,
	0xf4, 0x83, 0x54, 0xb6, 0xaf, 0x09, 0x24, 0x44, 0x90, 0xa6, 0x6f, 0xe7, 0xde, 0x7b, 0xce, 0xcf,
	0xbf, 0xdf, 0xb9, 0xc7, 0x3f, 0x1b, 0xa4, 0x81, 0xe3, 0x0c, 0x2c, 0x22, 0x19, 0x64, 0x1c, 0x38,
	0x8e, 0xe5, 0x4b, 0x3d, 0xcb, 0x19, 0x19, 0xae, 0xe7, 0xf4, 0x4d, 0x8b, 0x78, 0xd2, 0x58, 0x96,
	0x92, 0x58, 0x74, 0x3d, 0x27, 0x70, 0x30, 0x1f, 0x17, 0x88, 0x49, 0x81, 0x38, 0x57, 0x20, 0x8e,
	0xe5, 0xe2, 0x36, 0x85, 0xd4, 0x5d, 0x53, 0xd2, 0x6d, 0xdb, 0x09, 0xf4, 0xc0, 0x74, 0x6c, 0x3f,
	0xae, 0x2f, 0xfe, 0x44, 0x4f, 0xa3, 0x55, 0x77, 0xd4, 0x97, 0x8c, 0x91, 0x17, 0x25, 0xd0, 0x73,
	0xfe, 0xe6, 0x79, 0xdf, 0x24, 0x96, 0xd1, 0x19, 0xea, 0xfe, 0x25, 0xcd, 0xf8, 0xf9, 0x66, 0x46,
	0x60, 0x0e, 0x89, 0x1f, 0xe8, 0x43, 0x37, 0x4e, 0x10, 0xde, 0x23, 0xd8, 0xaa, 0x7a, 0x44, 0x0f,
	0x88, 0x1a, 0xd3, 0xd2, 0xc8, 0xe3, 0x11, 0xf1, 0x03, 0xfc, 0x1d, 0x64, 0x5c, 0xdd, 0x23, 0x76,
	0x50, 0x48, 0xf3, 0xa8, 0x94, 0xd3, 0xe8, 0x0a, 0x37, 0x00, 0x0c, 0xe2, 0x5a, 0xce, 0x64, 0x18,
	0x9e, 0x21, 0x1e, 0x95, 0x58, 0xf9, 0x0f, 0x71, 0x99, 0x50, 0xb1, 0x36, 0xad, 0xd1, 0x66, 0xea,
	0xb1, 0x0a, 0x1b, 0x34, 0xab, 0x13, 0x4c, 0x5c, 0x52, 0x48, 0xf1, 0x4c, 0x69, 0x53, 0xfe, 0x73,
	0x39, 0x1e, 0x65, 0xdb, 0x9e, 0xb8, 0x44, 0x63, 0xdd, 0xeb, 0x85, 0xf0, 0x04, 0x7e, 0x88, 0xf5,
	0xb4, 0xfa, 0x7d, 0xcb, 0xb4, 0xef, 0x96, 0x85, 0xe6, 0x64, 0x55, 0x21, 0x4b, 0x51, 0x0a, 0xa9,
	0x48, 0xd3, 0x6f, 0x2b, 0x73, 0xd0, 0x92, 0x4a, 0xe1, 0x25, 0x82, 0xad, 0x0b, 0xd7, 0xb8, 0xdd,
	0xcc, 0x19, 0x74, 0xf4, 0xa5, 0xe8, 0xf8, 0x10, 0xd8, 0x51, 0x04, 0x1e, 0x5d, 0x30, 0xa5, 0x59,
	0x4c, 0x80, 0x92, 0x1b, 0x16, 0x4f, 0xc2, 0x19, 0x38, 0xd3, 0xfd, 0x4b, 0x0d, 0xe2, 0xf4, 0x30,
	0x16, 0x5e, 0x33, 0x90, 0xa5, 0x88, 0x18, 0x43, 0xda, 0xd6, 0x87, 0x84, 0x76, 0x20, 0x8a, 0x17,
	0x5c, 0x04, 0x7a, 0xd8, 0x45, 0xdc, 0x18, 0x14, 0xe6, 0x81, 0x83, 0xb2, 0x0b, 0xeb, 0xc9, 0xf0,
	0x47, 0x03, 0xc9, 0xca, 0xdf, 0xdf, 0x52, 0x5e, 0xa3, 0x09, 0xda, 0x34, 0x15, 0xff, 0x02, 0xf9,
	0x44, 0x56, 0x77, 0x12, 0x10, 0xbf, 0xb0, 0xc6, 0xa3, 0xd2, 0x86, 0x96, 0x68, 0x3d, 0x0e, 0xf7,
	0xf0, 0x19, 0x64, 0x2c, 0xbd, 0x4b, 0x2c, 0xbf, 0x90, 0xe1, 0x99, 0x12, 0x2b, 0xef, 0xae, 0xac,
	0x5a, 0x6c, 0x44, 0x75, 0x8a, 0x1d, 0x78, 0x13, 0x8d, 0x82, 0x14, 0xf7, 0x81, 0x9d, 0xd9, 0xc6,
	0x1c, 0x30, 0x97, 0x64, 0x42, 0x9b, 0x1d, 0x86, 0x78, 0x0b, 0xd6, 0xc6, 0xba, 0x35, 0x8a, 0x9b,
	0x9c, 0xd3, 0xe2, 0xc5, 0x41, 0x6a, 0x0f, 0x09, 0x1f, 0x10, 0xc0, 0x75, 0x03, 0xf0, 0x8f, 0x00,
	0xae, 0xe7, 0x3c, 0x22, 0xbd, 0xa0, 0x63, 0x1a, 0x14, 0x21, 0x47, 0x77, 0xea, 0x46, 0x38, 0xcb,
	0x81, 0xee, 0x0d, 0x48, 0x40, 0x81, 0xe8, 0x0a, 0xab, 0x53, 0x3d, 0x4c, 0xa4, 0x67, 0xef, 0x3e,
	0x5d, 0xff, 0xca, 0x92, 0xca, 0xcf, 0x11, 0xb0, 0x33, 0x33, 0x82, 0xb7, 0xa1, 0xa0, 0x6a, 0xad,
	0x93, 0x7a, 0x43, 0xe9, 0xb4, 0xff, 0x53, 0x95, 0xce, 0x45, 0xf3, 0x5c, 0x55, 0xaa, 0xf5, 0x93,
	0xba, 0x52, 0xe3, 0xbe, 0xc1, 0x59, 0x60, 0xaa, 0xea, 0x05, 0x87, 0xf0, 0x3a, 0xa4, 0xff, 0xad,
	0x34, 0x1a, 0x5c, 0x2a, 0x8c, 0x4e, 0x95, 0x8a, 0xca, 0x31, 0x98, 0x85, 0x6c, 0xfb, 0x54, 0x53,
	0x2a, 0xb5, 0x73, 0x2e, 0x8d, 0x37, 0x01, 0xaa, 0xad, 0x66, 0x5b, 0x69, 0xb6, 0xeb, 0xad, 0x26,
	0xb7, 0x86, 0xf3, 0x90, 0x53, 0x95, 0xca, 0x3f, 0x9d, 0x28, 0x37, 0x13, 0x1e, 0x87, 0x51, 0xa7,
	0xd2, 0x68, 0xb4, 0xaa, 0x5c, 0x56, 0x7e, 0x93, 0x86, 0x6f, 0x29, 0x0d, 0xef, 0x9c, 0x78, 0x63,
	0xb3, 0x47, 0xf0, 0x2b, 0x04, 0xf9, 0x39, 0xef, 0xc3, 0x7f, 0x2d, 0xef, 0xd4, 0x22, 0xb3, 0x2c,
	0xae, 0xfe, 0x3a, 0x0b, 0xbf, 0x3f, 0xfb, 0xf8, 0xe9, 0x45, 0xea, 0x57, 0x81, 0x0f, 0xbf, 0x17,
	0x4f, 0x63, 0xf7, 0x39, 0xa2, 0x77, 0xea, 0x4b, 0xe5, 0xab, 0xe4, 0x1b, 0xe2, 0x1f, 0xa0, 0x32,
	0x7e, 0x37, 0x75, 0xe7, 0x79, 0x37, 0xc3, 0x47, 0xab, 0x12, 0x5d, 0xe8, 0x82, 0xf7, 0xe1, 0xfb,
	0x77, 0xc4, 0x77, 0x5f, 0x10, 0x97, 0xf2, 0xed, 0xcd, 0x3e, 0xf0, 0x60, 0x6a, 0x5b, 0x6f, 0x11,
	0xe4, 0xe7, 0x4c, 0x71, 0x95, 0x2e, 0x2f, 0x72, 0xd1, 0xfb, 0xb0, 0xde, 0x8f, 0x58, 0xef, 0xc8,
	0xa5, 0x98, 0x35, 0x7d, 0x5b, 0x43, 0xa7, 0x9b, 0xe1, 0x3e, 0xa5, 0x2e, 0x95, 0xaf, 0xa6, 0x7c,
	0x8f, 0x5b, 0xff, 0x9f, 0xd1, 0xc7, 0x0c, 0x1c, 0x4b, 0xb7, 0x07, 0xa2, 0xe3, 0x0d, 0xa4, 0x01,
	0xb1, 0x23, 0x9b, 0xa1, 0xbf, 0x00, 0xba, 0x6b, 0xfa, 0x77, 0xff, 0x06, 0x1c, 0xce, 0x6d, 0x74,
	0x33, 0x51, 0xe5, 0xce, 0xe7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x08, 0x2b, 0x60, 0x3f, 0x08,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ProfilerServiceClient is the client API for ProfilerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProfilerServiceClient interface {
	// CreateProfile creates a new profile resource in the online mode.
	//
	// The server ensures that the new profiles are created at a constant rate per
	// deployment, so the creation request may hang for some time until the next
	// profile session is available.
	//
	// The request may fail with ABORTED error if the creation is not available
	// within ~1m, the response will indicate the duration of the backoff the
	// client should take before attempting creating a profile again. The backoff
	// duration is returned in google.rpc.RetryInfo extension on the response
	// status. To a gRPC client, the extension will be return as a
	// binary-serialized proto in the trailing metadata item named
	// "google.rpc.retryinfo-bin".
	CreateProfile(ctx context.Context, in *CreateProfileRequest, opts ...grpc.CallOption) (*Profile, error)
	// CreateOfflineProfile creates a new profile resource in the offline mode.
	// The client provides the profile to create along with the profile bytes, the
	// server records it.
	CreateOfflineProfile(ctx context.Context, in *CreateOfflineProfileRequest, opts ...grpc.CallOption) (*Profile, error)
	// UpdateProfile updates the profile bytes and labels on the profile resource
	// created in the online mode. Updating the bytes for profiles created in the
	// offline mode is currently not supported: the profile content must be
	// provided at the time of the profile creation.
	UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*Profile, error)
}

type profilerServiceClient struct {
	cc *grpc.ClientConn
}

func NewProfilerServiceClient(cc *grpc.ClientConn) ProfilerServiceClient {
	return &profilerServiceClient{cc}
}

func (c *profilerServiceClient) CreateProfile(ctx context.Context, in *CreateProfileRequest, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := c.cc.Invoke(ctx, "/google.devtools.cloudprofiler.v2.ProfilerService/CreateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilerServiceClient) CreateOfflineProfile(ctx context.Context, in *CreateOfflineProfileRequest, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := c.cc.Invoke(ctx, "/google.devtools.cloudprofiler.v2.ProfilerService/CreateOfflineProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilerServiceClient) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := c.cc.Invoke(ctx, "/google.devtools.cloudprofiler.v2.ProfilerService/UpdateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfilerServiceServer is the server API for ProfilerService service.
type ProfilerServiceServer interface {
	// CreateProfile creates a new profile resource in the online mode.
	//
	// The server ensures that the new profiles are created at a constant rate per
	// deployment, so the creation request may hang for some time until the next
	// profile session is available.
	//
	// The request may fail with ABORTED error if the creation is not available
	// within ~1m, the response will indicate the duration of the backoff the
	// client should take before attempting creating a profile again. The backoff
	// duration is returned in google.rpc.RetryInfo extension on the response
	// status. To a gRPC client, the extension will be return as a
	// binary-serialized proto in the trailing metadata item named
	// "google.rpc.retryinfo-bin".
	CreateProfile(context.Context, *CreateProfileRequest) (*Profile, error)
	// CreateOfflineProfile creates a new profile resource in the offline mode.
	// The client provides the profile to create along with the profile bytes, the
	// server records it.
	CreateOfflineProfile(context.Context, *CreateOfflineProfileRequest) (*Profile, error)
	// UpdateProfile updates the profile bytes and labels on the profile resource
	// created in the online mode. Updating the bytes for profiles created in the
	// offline mode is currently not supported: the profile content must be
	// provided at the time of the profile creation.
	UpdateProfile(context.Context, *UpdateProfileRequest) (*Profile, error)
}

// UnimplementedProfilerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProfilerServiceServer struct {
}

func (*UnimplementedProfilerServiceServer) CreateProfile(ctx context.Context, req *CreateProfileRequest) (*Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProfile not implemented")
}
func (*UnimplementedProfilerServiceServer) CreateOfflineProfile(ctx context.Context, req *CreateOfflineProfileRequest) (*Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOfflineProfile not implemented")
}
func (*UnimplementedProfilerServiceServer) UpdateProfile(ctx context.Context, req *UpdateProfileRequest) (*Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}

func RegisterProfilerServiceServer(s *grpc.Server, srv ProfilerServiceServer) {
	s.RegisterService(&_ProfilerService_serviceDesc, srv)
}

func _ProfilerService_CreateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilerServiceServer).CreateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.cloudprofiler.v2.ProfilerService/CreateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilerServiceServer).CreateProfile(ctx, req.(*CreateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfilerService_CreateOfflineProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOfflineProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilerServiceServer).CreateOfflineProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.cloudprofiler.v2.ProfilerService/CreateOfflineProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilerServiceServer).CreateOfflineProfile(ctx, req.(*CreateOfflineProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfilerService_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilerServiceServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.cloudprofiler.v2.ProfilerService/UpdateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilerServiceServer).UpdateProfile(ctx, req.(*UpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProfilerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.devtools.cloudprofiler.v2.ProfilerService",
	HandlerType: (*ProfilerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProfile",
			Handler:    _ProfilerService_CreateProfile_Handler,
		},
		{
			MethodName: "CreateOfflineProfile",
			Handler:    _ProfilerService_CreateOfflineProfile_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _ProfilerService_UpdateProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/devtools/cloudprofiler/v2/profiler.proto",
}

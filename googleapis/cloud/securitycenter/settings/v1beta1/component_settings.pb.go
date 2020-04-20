// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/securitycenter/settings/v1beta1/component_settings.proto

package settings

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Valid states for a component
type ComponentEnablementState int32

const (
	// No state specified, equivalent of INHERIT
	ComponentEnablementState_COMPONENT_ENABLEMENT_STATE_UNSPECIFIED ComponentEnablementState = 0
	// Disable the component
	ComponentEnablementState_DISABLE ComponentEnablementState = 1
	// Enable the component
	ComponentEnablementState_ENABLE ComponentEnablementState = 2
	// Inherit the state from resources parent folder or organization.
	ComponentEnablementState_INHERIT ComponentEnablementState = 3
)

var ComponentEnablementState_name = map[int32]string{
	0: "COMPONENT_ENABLEMENT_STATE_UNSPECIFIED",
	1: "DISABLE",
	2: "ENABLE",
	3: "INHERIT",
}

var ComponentEnablementState_value = map[string]int32{
	"COMPONENT_ENABLEMENT_STATE_UNSPECIFIED": 0,
	"DISABLE":                                1,
	"ENABLE":                                 2,
	"INHERIT":                                3,
}

func (x ComponentEnablementState) String() string {
	return proto.EnumName(ComponentEnablementState_name, int32(x))
}

func (ComponentEnablementState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f91b4ec8b41218e5, []int{0}
}

// Component Settings for Security Command Center
type ComponentSettings struct {
	// The relative resource name of the component settings.
	// Formats:
	//  * `organizations/{organization}/components/{component}/settings`
	//  * `folders/{folder}/components/{component}/settings`
	//  * `projects/{project}/components/{component}/settings`
	//  * `projects/{project}/locations/{location}/clusters/{cluster}/components/{component}/settings`
	//  * `projects/{project}/regions/{region}/clusters/{cluster}/components/{component}/settings`
	//  * `projects/{project}/zones/{zone}/clusters/{cluster}/components/{component}/settings`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ENABLE to enable component, DISABLE to disable and INHERIT to inherit
	// setting from ancestors.
	State ComponentEnablementState `protobuf:"varint,2,opt,name=state,proto3,enum=google.cloud.securitycenter.settings.v1beta1.ComponentEnablementState" json:"state,omitempty"`
	// Output only. The service account to be used for security center component.
	// The component must have permission to "act as" the service account.
	ProjectServiceAccount string `protobuf:"bytes,3,opt,name=project_service_account,json=projectServiceAccount,proto3" json:"project_service_account,omitempty"`
	// Settings for detectors.  Not all detectors must have settings present at
	// each and every level in the hierarchy.  If it is not present the setting
	// will be inherited from its ancestors folders, organizations or the
	// defaults.
	DetectorSettings map[string]*ComponentSettings_DetectorSettings `protobuf:"bytes,4,rep,name=detector_settings,json=detectorSettings,proto3" json:"detector_settings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Output only. An fingerprint used for optimistic concurrency. If none is provided
	// on updates then the existing metadata will be blindly overwritten.
	Etag string `protobuf:"bytes,5,opt,name=etag,proto3" json:"etag,omitempty"`
	// Output only. The time these settings were last updated.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Component specific settings.  This must match the component value.
	//
	// Types that are valid to be assigned to SpecificSettings:
	//	*ComponentSettings_ContainerThreatDetectionSettings
	//	*ComponentSettings_EventThreatDetectionSettings
	//	*ComponentSettings_SecurityHealthAnalyticsSettings
	//	*ComponentSettings_WebSecurityScannerSettings
	SpecificSettings     isComponentSettings_SpecificSettings `protobuf_oneof:"specific_settings"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *ComponentSettings) Reset()         { *m = ComponentSettings{} }
func (m *ComponentSettings) String() string { return proto.CompactTextString(m) }
func (*ComponentSettings) ProtoMessage()    {}
func (*ComponentSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_f91b4ec8b41218e5, []int{0}
}

func (m *ComponentSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComponentSettings.Unmarshal(m, b)
}
func (m *ComponentSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComponentSettings.Marshal(b, m, deterministic)
}
func (m *ComponentSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComponentSettings.Merge(m, src)
}
func (m *ComponentSettings) XXX_Size() int {
	return xxx_messageInfo_ComponentSettings.Size(m)
}
func (m *ComponentSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_ComponentSettings.DiscardUnknown(m)
}

var xxx_messageInfo_ComponentSettings proto.InternalMessageInfo

func (m *ComponentSettings) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ComponentSettings) GetState() ComponentEnablementState {
	if m != nil {
		return m.State
	}
	return ComponentEnablementState_COMPONENT_ENABLEMENT_STATE_UNSPECIFIED
}

func (m *ComponentSettings) GetProjectServiceAccount() string {
	if m != nil {
		return m.ProjectServiceAccount
	}
	return ""
}

func (m *ComponentSettings) GetDetectorSettings() map[string]*ComponentSettings_DetectorSettings {
	if m != nil {
		return m.DetectorSettings
	}
	return nil
}

func (m *ComponentSettings) GetEtag() string {
	if m != nil {
		return m.Etag
	}
	return ""
}

func (m *ComponentSettings) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

type isComponentSettings_SpecificSettings interface {
	isComponentSettings_SpecificSettings()
}

type ComponentSettings_ContainerThreatDetectionSettings struct {
	ContainerThreatDetectionSettings *ContainerThreatDetectionSettings `protobuf:"bytes,41,opt,name=container_threat_detection_settings,json=containerThreatDetectionSettings,proto3,oneof"`
}

type ComponentSettings_EventThreatDetectionSettings struct {
	EventThreatDetectionSettings *EventThreatDetectionSettings `protobuf:"bytes,42,opt,name=event_threat_detection_settings,json=eventThreatDetectionSettings,proto3,oneof"`
}

type ComponentSettings_SecurityHealthAnalyticsSettings struct {
	SecurityHealthAnalyticsSettings *SecurityHealthAnalyticsSettings `protobuf:"bytes,44,opt,name=security_health_analytics_settings,json=securityHealthAnalyticsSettings,proto3,oneof"`
}

type ComponentSettings_WebSecurityScannerSettings struct {
	WebSecurityScannerSettings *WebSecurityScanner `protobuf:"bytes,40,opt,name=web_security_scanner_settings,json=webSecurityScannerSettings,proto3,oneof"`
}

func (*ComponentSettings_ContainerThreatDetectionSettings) isComponentSettings_SpecificSettings() {}

func (*ComponentSettings_EventThreatDetectionSettings) isComponentSettings_SpecificSettings() {}

func (*ComponentSettings_SecurityHealthAnalyticsSettings) isComponentSettings_SpecificSettings() {}

func (*ComponentSettings_WebSecurityScannerSettings) isComponentSettings_SpecificSettings() {}

func (m *ComponentSettings) GetSpecificSettings() isComponentSettings_SpecificSettings {
	if m != nil {
		return m.SpecificSettings
	}
	return nil
}

func (m *ComponentSettings) GetContainerThreatDetectionSettings() *ContainerThreatDetectionSettings {
	if x, ok := m.GetSpecificSettings().(*ComponentSettings_ContainerThreatDetectionSettings); ok {
		return x.ContainerThreatDetectionSettings
	}
	return nil
}

func (m *ComponentSettings) GetEventThreatDetectionSettings() *EventThreatDetectionSettings {
	if x, ok := m.GetSpecificSettings().(*ComponentSettings_EventThreatDetectionSettings); ok {
		return x.EventThreatDetectionSettings
	}
	return nil
}

func (m *ComponentSettings) GetSecurityHealthAnalyticsSettings() *SecurityHealthAnalyticsSettings {
	if x, ok := m.GetSpecificSettings().(*ComponentSettings_SecurityHealthAnalyticsSettings); ok {
		return x.SecurityHealthAnalyticsSettings
	}
	return nil
}

func (m *ComponentSettings) GetWebSecurityScannerSettings() *WebSecurityScanner {
	if x, ok := m.GetSpecificSettings().(*ComponentSettings_WebSecurityScannerSettings); ok {
		return x.WebSecurityScannerSettings
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ComponentSettings) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ComponentSettings_ContainerThreatDetectionSettings)(nil),
		(*ComponentSettings_EventThreatDetectionSettings)(nil),
		(*ComponentSettings_SecurityHealthAnalyticsSettings)(nil),
		(*ComponentSettings_WebSecurityScannerSettings)(nil),
	}
}

// Settings for each detector.
type ComponentSettings_DetectorSettings struct {
	// ENABLE to enable component, DISABLE to disable and INHERIT to inherit
	// setting from ancestors.
	State                ComponentEnablementState `protobuf:"varint,1,opt,name=state,proto3,enum=google.cloud.securitycenter.settings.v1beta1.ComponentEnablementState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ComponentSettings_DetectorSettings) Reset()         { *m = ComponentSettings_DetectorSettings{} }
func (m *ComponentSettings_DetectorSettings) String() string { return proto.CompactTextString(m) }
func (*ComponentSettings_DetectorSettings) ProtoMessage()    {}
func (*ComponentSettings_DetectorSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_f91b4ec8b41218e5, []int{0, 0}
}

func (m *ComponentSettings_DetectorSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComponentSettings_DetectorSettings.Unmarshal(m, b)
}
func (m *ComponentSettings_DetectorSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComponentSettings_DetectorSettings.Marshal(b, m, deterministic)
}
func (m *ComponentSettings_DetectorSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComponentSettings_DetectorSettings.Merge(m, src)
}
func (m *ComponentSettings_DetectorSettings) XXX_Size() int {
	return xxx_messageInfo_ComponentSettings_DetectorSettings.Size(m)
}
func (m *ComponentSettings_DetectorSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_ComponentSettings_DetectorSettings.DiscardUnknown(m)
}

var xxx_messageInfo_ComponentSettings_DetectorSettings proto.InternalMessageInfo

func (m *ComponentSettings_DetectorSettings) GetState() ComponentEnablementState {
	if m != nil {
		return m.State
	}
	return ComponentEnablementState_COMPONENT_ENABLEMENT_STATE_UNSPECIFIED
}

// User specified settings for Web Security Scanner
type WebSecurityScanner struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WebSecurityScanner) Reset()         { *m = WebSecurityScanner{} }
func (m *WebSecurityScanner) String() string { return proto.CompactTextString(m) }
func (*WebSecurityScanner) ProtoMessage()    {}
func (*WebSecurityScanner) Descriptor() ([]byte, []int) {
	return fileDescriptor_f91b4ec8b41218e5, []int{1}
}

func (m *WebSecurityScanner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebSecurityScanner.Unmarshal(m, b)
}
func (m *WebSecurityScanner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebSecurityScanner.Marshal(b, m, deterministic)
}
func (m *WebSecurityScanner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebSecurityScanner.Merge(m, src)
}
func (m *WebSecurityScanner) XXX_Size() int {
	return xxx_messageInfo_WebSecurityScanner.Size(m)
}
func (m *WebSecurityScanner) XXX_DiscardUnknown() {
	xxx_messageInfo_WebSecurityScanner.DiscardUnknown(m)
}

var xxx_messageInfo_WebSecurityScanner proto.InternalMessageInfo

// User specified settings for KTD
type ContainerThreatDetectionSettings struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContainerThreatDetectionSettings) Reset()         { *m = ContainerThreatDetectionSettings{} }
func (m *ContainerThreatDetectionSettings) String() string { return proto.CompactTextString(m) }
func (*ContainerThreatDetectionSettings) ProtoMessage()    {}
func (*ContainerThreatDetectionSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_f91b4ec8b41218e5, []int{2}
}

func (m *ContainerThreatDetectionSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerThreatDetectionSettings.Unmarshal(m, b)
}
func (m *ContainerThreatDetectionSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerThreatDetectionSettings.Marshal(b, m, deterministic)
}
func (m *ContainerThreatDetectionSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerThreatDetectionSettings.Merge(m, src)
}
func (m *ContainerThreatDetectionSettings) XXX_Size() int {
	return xxx_messageInfo_ContainerThreatDetectionSettings.Size(m)
}
func (m *ContainerThreatDetectionSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerThreatDetectionSettings.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerThreatDetectionSettings proto.InternalMessageInfo

// User specified settings for ETD
type EventThreatDetectionSettings struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventThreatDetectionSettings) Reset()         { *m = EventThreatDetectionSettings{} }
func (m *EventThreatDetectionSettings) String() string { return proto.CompactTextString(m) }
func (*EventThreatDetectionSettings) ProtoMessage()    {}
func (*EventThreatDetectionSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_f91b4ec8b41218e5, []int{3}
}

func (m *EventThreatDetectionSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventThreatDetectionSettings.Unmarshal(m, b)
}
func (m *EventThreatDetectionSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventThreatDetectionSettings.Marshal(b, m, deterministic)
}
func (m *EventThreatDetectionSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventThreatDetectionSettings.Merge(m, src)
}
func (m *EventThreatDetectionSettings) XXX_Size() int {
	return xxx_messageInfo_EventThreatDetectionSettings.Size(m)
}
func (m *EventThreatDetectionSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_EventThreatDetectionSettings.DiscardUnknown(m)
}

var xxx_messageInfo_EventThreatDetectionSettings proto.InternalMessageInfo

// User specified settings for Security Health Analytics
type SecurityHealthAnalyticsSettings struct {
	// Settings for "NON_ORG_IAM_MEMBER" scanner.
	NonOrgIamMemberSettings *SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings `protobuf:"bytes,1,opt,name=non_org_iam_member_settings,json=nonOrgIamMemberSettings,proto3" json:"non_org_iam_member_settings,omitempty"`
	// Settings for "ADMIN_SERVICE_ACCOUNT" scanner.
	AdminServiceAccountSettings *SecurityHealthAnalyticsSettings_AdminServiceAccountSettings `protobuf:"bytes,2,opt,name=admin_service_account_settings,json=adminServiceAccountSettings,proto3" json:"admin_service_account_settings,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}                                                     `json:"-"`
	XXX_unrecognized            []byte                                                       `json:"-"`
	XXX_sizecache               int32                                                        `json:"-"`
}

func (m *SecurityHealthAnalyticsSettings) Reset()         { *m = SecurityHealthAnalyticsSettings{} }
func (m *SecurityHealthAnalyticsSettings) String() string { return proto.CompactTextString(m) }
func (*SecurityHealthAnalyticsSettings) ProtoMessage()    {}
func (*SecurityHealthAnalyticsSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_f91b4ec8b41218e5, []int{4}
}

func (m *SecurityHealthAnalyticsSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecurityHealthAnalyticsSettings.Unmarshal(m, b)
}
func (m *SecurityHealthAnalyticsSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecurityHealthAnalyticsSettings.Marshal(b, m, deterministic)
}
func (m *SecurityHealthAnalyticsSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecurityHealthAnalyticsSettings.Merge(m, src)
}
func (m *SecurityHealthAnalyticsSettings) XXX_Size() int {
	return xxx_messageInfo_SecurityHealthAnalyticsSettings.Size(m)
}
func (m *SecurityHealthAnalyticsSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_SecurityHealthAnalyticsSettings.DiscardUnknown(m)
}

var xxx_messageInfo_SecurityHealthAnalyticsSettings proto.InternalMessageInfo

func (m *SecurityHealthAnalyticsSettings) GetNonOrgIamMemberSettings() *SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings {
	if m != nil {
		return m.NonOrgIamMemberSettings
	}
	return nil
}

func (m *SecurityHealthAnalyticsSettings) GetAdminServiceAccountSettings() *SecurityHealthAnalyticsSettings_AdminServiceAccountSettings {
	if m != nil {
		return m.AdminServiceAccountSettings
	}
	return nil
}

// Settings for "NON_ORG_IAM_MEMBER" scanner.
type SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings struct {
	// User emails ending in the provided identities are allowed to have IAM
	// permissions on a project or the organization. Otherwise a finding will
	// be created.
	// A valid identity can be:
	//   *  a domain that starts with "@", e.g. "@yourdomain.com".
	//   *  a fully specified email address that does not start with "@", e.g.
	//   "abc@gmail.com"
	// Regular expressions are not supported.
	// Service accounts are not examined by the scanner and will be omitted if
	// added to the list.
	// If not specified, only Gmail accounts will be considered as non-approved.
	ApprovedIdentities   []string `protobuf:"bytes,1,rep,name=approved_identities,json=approvedIdentities,proto3" json:"approved_identities,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings) Reset() {
	*m = SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings{}
}
func (m *SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings) String() string {
	return proto.CompactTextString(m)
}
func (*SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings) ProtoMessage() {}
func (*SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_f91b4ec8b41218e5, []int{4, 0}
}

func (m *SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings.Unmarshal(m, b)
}
func (m *SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings.Marshal(b, m, deterministic)
}
func (m *SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings.Merge(m, src)
}
func (m *SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings) XXX_Size() int {
	return xxx_messageInfo_SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings.Size(m)
}
func (m *SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings.DiscardUnknown(m)
}

var xxx_messageInfo_SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings proto.InternalMessageInfo

func (m *SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings) GetApprovedIdentities() []string {
	if m != nil {
		return m.ApprovedIdentities
	}
	return nil
}

// Settings for "ADMIN_SERVICE_ACCOUNT" scanner.
type SecurityHealthAnalyticsSettings_AdminServiceAccountSettings struct {
	// User-created service accounts ending in the provided identities are
	// allowed to have Admin, Owner or Editor roles granted to them. Otherwise
	// a finding will be created.
	// A valid identity can be:
	//   *  a partilly specified service account that starts with "@", e.g.
	//   "@myproject.iam.gserviceaccount.com". This approves all the service
	//   accounts suffixed with the specified identity.
	//   *  a fully specified service account that does not start with "@", e.g.
	//   "myadmin@myproject.iam.gserviceaccount.com".
	// Google-created service accounts are all approved.
	ApprovedIdentities   []string `protobuf:"bytes,1,rep,name=approved_identities,json=approvedIdentities,proto3" json:"approved_identities,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SecurityHealthAnalyticsSettings_AdminServiceAccountSettings) Reset() {
	*m = SecurityHealthAnalyticsSettings_AdminServiceAccountSettings{}
}
func (m *SecurityHealthAnalyticsSettings_AdminServiceAccountSettings) String() string {
	return proto.CompactTextString(m)
}
func (*SecurityHealthAnalyticsSettings_AdminServiceAccountSettings) ProtoMessage() {}
func (*SecurityHealthAnalyticsSettings_AdminServiceAccountSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_f91b4ec8b41218e5, []int{4, 1}
}

func (m *SecurityHealthAnalyticsSettings_AdminServiceAccountSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecurityHealthAnalyticsSettings_AdminServiceAccountSettings.Unmarshal(m, b)
}
func (m *SecurityHealthAnalyticsSettings_AdminServiceAccountSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecurityHealthAnalyticsSettings_AdminServiceAccountSettings.Marshal(b, m, deterministic)
}
func (m *SecurityHealthAnalyticsSettings_AdminServiceAccountSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecurityHealthAnalyticsSettings_AdminServiceAccountSettings.Merge(m, src)
}
func (m *SecurityHealthAnalyticsSettings_AdminServiceAccountSettings) XXX_Size() int {
	return xxx_messageInfo_SecurityHealthAnalyticsSettings_AdminServiceAccountSettings.Size(m)
}
func (m *SecurityHealthAnalyticsSettings_AdminServiceAccountSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_SecurityHealthAnalyticsSettings_AdminServiceAccountSettings.DiscardUnknown(m)
}

var xxx_messageInfo_SecurityHealthAnalyticsSettings_AdminServiceAccountSettings proto.InternalMessageInfo

func (m *SecurityHealthAnalyticsSettings_AdminServiceAccountSettings) GetApprovedIdentities() []string {
	if m != nil {
		return m.ApprovedIdentities
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.cloud.securitycenter.settings.v1beta1.ComponentEnablementState", ComponentEnablementState_name, ComponentEnablementState_value)
	proto.RegisterType((*ComponentSettings)(nil), "google.cloud.securitycenter.settings.v1beta1.ComponentSettings")
	proto.RegisterMapType((map[string]*ComponentSettings_DetectorSettings)(nil), "google.cloud.securitycenter.settings.v1beta1.ComponentSettings.DetectorSettingsEntry")
	proto.RegisterType((*ComponentSettings_DetectorSettings)(nil), "google.cloud.securitycenter.settings.v1beta1.ComponentSettings.DetectorSettings")
	proto.RegisterType((*WebSecurityScanner)(nil), "google.cloud.securitycenter.settings.v1beta1.WebSecurityScanner")
	proto.RegisterType((*ContainerThreatDetectionSettings)(nil), "google.cloud.securitycenter.settings.v1beta1.ContainerThreatDetectionSettings")
	proto.RegisterType((*EventThreatDetectionSettings)(nil), "google.cloud.securitycenter.settings.v1beta1.EventThreatDetectionSettings")
	proto.RegisterType((*SecurityHealthAnalyticsSettings)(nil), "google.cloud.securitycenter.settings.v1beta1.SecurityHealthAnalyticsSettings")
	proto.RegisterType((*SecurityHealthAnalyticsSettings_NonOrgIamMemberSettings)(nil), "google.cloud.securitycenter.settings.v1beta1.SecurityHealthAnalyticsSettings.NonOrgIamMemberSettings")
	proto.RegisterType((*SecurityHealthAnalyticsSettings_AdminServiceAccountSettings)(nil), "google.cloud.securitycenter.settings.v1beta1.SecurityHealthAnalyticsSettings.AdminServiceAccountSettings")
}

func init() {
	proto.RegisterFile("google/cloud/securitycenter/settings/v1beta1/component_settings.proto", fileDescriptor_f91b4ec8b41218e5)
}

var fileDescriptor_f91b4ec8b41218e5 = []byte{
	// 1031 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x41, 0x6f, 0xe3, 0x44,
	0x14, 0xae, 0x93, 0xb6, 0x68, 0xa7, 0xd2, 0x2a, 0x9d, 0x65, 0x69, 0x48, 0xcb, 0x36, 0x0a, 0x12,
	0x0a, 0x55, 0x65, 0xb7, 0xe5, 0x82, 0xb2, 0x1c, 0x36, 0x69, 0xbd, 0x34, 0x2b, 0x9a, 0x46, 0x4e,
	0xb6, 0x48, 0xab, 0x48, 0xd6, 0x64, 0x32, 0x71, 0x07, 0xec, 0x19, 0xcb, 0x9e, 0x64, 0xd5, 0x8d,
	0x7a, 0x41, 0xf0, 0x0b, 0x80, 0x1b, 0x12, 0x07, 0x6e, 0x1c, 0x38, 0xf2, 0x23, 0xf8, 0x0f, 0x48,
	0x70, 0xdd, 0x5f, 0xc0, 0x11, 0xd9, 0xe3, 0x71, 0xd2, 0x36, 0x69, 0x48, 0x55, 0x4e, 0x7e, 0x9e,
	0xf7, 0xbd, 0xef, 0x7d, 0xcf, 0x33, 0xef, 0x8d, 0x81, 0xe9, 0x70, 0xee, 0xb8, 0xc4, 0xc0, 0x2e,
	0x1f, 0xf4, 0x8c, 0x90, 0xe0, 0x41, 0x40, 0xc5, 0x05, 0x26, 0x4c, 0x90, 0xc0, 0x08, 0x89, 0x10,
	0x94, 0x39, 0xa1, 0x31, 0xdc, 0xef, 0x12, 0x81, 0xf6, 0x0d, 0xcc, 0x3d, 0x9f, 0x33, 0xc2, 0x84,
	0xad, 0x5c, 0xba, 0x1f, 0x70, 0xc1, 0xe1, 0xae, 0xa4, 0xd1, 0x63, 0x1a, 0xfd, 0x2a, 0x8d, 0x9e,
	0x62, 0x13, 0x9a, 0xc2, 0x76, 0x92, 0x14, 0xf9, 0xd4, 0xe8, 0x53, 0xe2, 0xf6, 0xec, 0x2e, 0x39,
	0x47, 0x43, 0xca, 0x03, 0x49, 0x57, 0x78, 0x7f, 0x02, 0x10, 0x90, 0x90, 0x0f, 0x02, 0x4c, 0x12,
	0x97, 0x8a, 0x8d, 0xdf, 0xba, 0x83, 0xbe, 0x21, 0xa8, 0x47, 0x42, 0x81, 0x3c, 0x3f, 0x01, 0x6c,
	0x4d, 0xc4, 0x22, 0xc6, 0xb8, 0x40, 0x82, 0x72, 0x96, 0x08, 0x2d, 0x7d, 0xfb, 0x10, 0xac, 0x1f,
	0xaa, 0x2a, 0x5a, 0x89, 0x30, 0x08, 0xc1, 0x32, 0x43, 0x1e, 0xc9, 0x6b, 0x45, 0xad, 0xfc, 0xc0,
	0x8a, 0x6d, 0xd8, 0x01, 0x2b, 0xa1, 0x40, 0x82, 0xe4, 0x33, 0x45, 0xad, 0xfc, 0xf0, 0xe0, 0xb9,
	0xbe, 0x48, 0x89, 0x7a, 0x9a, 0xc3, 0x64, 0xa8, 0xeb, 0x12, 0x2f, 0xca, 0x16, 0xb1, 0x59, 0x92,
	0x14, 0x3e, 0x05, 0x1b, 0x7e, 0xc0, 0xbf, 0x22, 0x38, 0xfa, 0x94, 0xc1, 0x90, 0x62, 0x62, 0x23,
	0x8c, 0xf9, 0x80, 0x89, 0x7c, 0x36, 0x12, 0x51, 0xcb, 0xfe, 0x55, 0xcd, 0x5a, 0x8f, 0x13, 0x4c,
	0x4b, 0x42, 0xaa, 0x12, 0x01, 0xbf, 0xd1, 0xc0, 0x7a, 0x8f, 0x08, 0x82, 0x05, 0x0f, 0xd2, 0x9d,
	0xc8, 0x2f, 0x17, 0xb3, 0xe5, 0xb5, 0x83, 0x97, 0x77, 0xd4, 0xa9, 0xbe, 0x85, 0x7e, 0x94, 0x10,
	0xab, 0x05, 0x93, 0x89, 0xe0, 0xc2, 0xca, 0xf5, 0xae, 0x2d, 0xc3, 0x0d, 0xb0, 0x4c, 0x04, 0x72,
	0xf2, 0x2b, 0x63, 0xb9, 0xf1, 0x02, 0x7c, 0x06, 0xd6, 0x06, 0x7e, 0x0f, 0x09, 0x62, 0x47, 0x5b,
	0x93, 0x5f, 0x2d, 0x6a, 0xe5, 0xb5, 0x83, 0x82, 0x92, 0xa5, 0xf6, 0x4d, 0x6f, 0xab, 0x7d, 0x93,
	0xb1, 0x40, 0xc6, 0x44, 0xab, 0xf0, 0x67, 0x0d, 0x7c, 0x88, 0x39, 0x13, 0x88, 0x32, 0x12, 0xd8,
	0xe2, 0x3c, 0x20, 0x48, 0xd8, 0x52, 0x00, 0xe5, 0x6c, 0x5c, 0xf1, 0xc7, 0x31, 0x75, 0x63, 0xd1,
	0x8a, 0x13, 0xe2, 0x76, 0xcc, 0x7b, 0xa4, 0x68, 0x55, 0x61, 0xc7, 0x4b, 0x56, 0x11, 0xcf, 0xc1,
	0xc0, 0xef, 0x35, 0xb0, 0x4d, 0x86, 0x51, 0x23, 0xcc, 0x56, 0xb7, 0x13, 0xab, 0x7b, 0xb1, 0x98,
	0x3a, 0x33, 0x22, 0x9d, 0xad, 0x6c, 0x8b, 0xdc, 0xe2, 0x87, 0x3f, 0x69, 0xa0, 0xa4, 0x12, 0xd8,
	0xe7, 0x04, 0xb9, 0xe2, 0xdc, 0x46, 0x0c, 0xb9, 0x17, 0x82, 0xe2, 0x70, 0x2c, 0x6c, 0x37, 0x16,
	0x76, 0xb2, 0x98, 0xb0, 0x56, 0xe2, 0x3f, 0x8e, 0x69, 0xab, 0x8a, 0x75, 0x42, 0xdb, 0x76, 0x78,
	0x3b, 0x04, 0x7e, 0xa7, 0x81, 0x0f, 0x5e, 0x93, 0xae, 0x9d, 0x4a, 0x0c, 0x31, 0x62, 0xd1, 0x0e,
	0xa7, 0xca, 0xca, 0xb1, 0xb2, 0x67, 0x8b, 0x29, 0xfb, 0x92, 0x74, 0x95, 0xb8, 0x96, 0x24, 0x3c,
	0x5e, 0xb2, 0x0a, 0xaf, 0x6f, 0xac, 0x2a, 0x1d, 0x05, 0x1f, 0xe4, 0xae, 0x1f, 0xf2, 0x71, 0xb7,
	0x6b, 0xff, 0x43, 0xb7, 0x17, 0x7e, 0xd4, 0xc0, 0xe3, 0xa9, 0x7d, 0x05, 0x73, 0x20, 0xfb, 0x35,
	0xb9, 0x48, 0x06, 0x4f, 0x64, 0xc2, 0x3e, 0x58, 0x19, 0x22, 0x77, 0x20, 0xe7, 0xce, 0xda, 0x41,
	0xf3, 0xbe, 0xfb, 0xd9, 0x92, 0xf4, 0x95, 0xcc, 0xa7, 0x5a, 0xe5, 0xef, 0xec, 0xdb, 0xea, 0x9f,
	0x59, 0x70, 0x6d, 0xee, 0x27, 0x19, 0x91, 0x4f, 0x43, 0x1d, 0x73, 0xcf, 0xb8, 0x39, 0x2f, 0x3f,
	0xe3, 0x81, 0x83, 0x18, 0x7d, 0x23, 0x47, 0xab, 0x31, 0x9a, 0x7c, 0xbd, 0x1c, 0xdf, 0x12, 0xa1,
	0x31, 0x4a, 0xed, 0xcb, 0xf4, 0x36, 0x81, 0x7b, 0x7d, 0xee, 0xf6, 0x48, 0x10, 0x1a, 0x23, 0x69,
	0xcc, 0x8f, 0x38, 0x48, 0xe6, 0x60, 0x68, 0x8c, 0x12, 0x6b, 0x7e, 0xcc, 0xab, 0x29, 0x31, 0x2e,
	0xc7, 0x4a, 0xb2, 0x32, 0x2f, 0x0d, 0xec, 0x0e, 0x42, 0x11, 0xcb, 0x49, 0xac, 0xf9, 0xdc, 0x67,
	0x53, 0xb8, 0x03, 0xe2, 0x48, 0x66, 0x69, 0xdc, 0x89, 0xd7, 0x9a, 0xc2, 0xfb, 0x86, 0x33, 0x12,
	0x1a, 0xa3, 0xe8, 0x71, 0x17, 0xce, 0xda, 0x23, 0xb0, 0x1e, 0xfa, 0x04, 0xd3, 0x3e, 0xc5, 0x69,
	0xa3, 0x95, 0xde, 0x05, 0xf0, 0x66, 0xdb, 0x94, 0x4a, 0xa0, 0x38, 0x6f, 0x3a, 0x96, 0x9e, 0x80,
	0xad, 0xdb, 0x66, 0x54, 0xe9, 0x87, 0x65, 0xb0, 0x3d, 0x67, 0x56, 0xc0, 0x5f, 0x34, 0xb0, 0xc9,
	0x38, 0xb3, 0x79, 0xe0, 0xd8, 0x14, 0x79, 0xb6, 0x47, 0xbc, 0xee, 0xe4, 0x18, 0xd0, 0xe2, 0x93,
	0x4f, 0xee, 0x75, 0x40, 0xe9, 0x0d, 0xce, 0x4e, 0x03, 0xa7, 0x8e, 0xbc, 0x93, 0x38, 0x5b, 0xda,
	0x0e, 0x1b, 0x6c, 0xba, 0x03, 0xfe, 0xa6, 0x81, 0x27, 0xa8, 0xe7, 0x51, 0x76, 0xfd, 0x86, 0x1e,
	0x0b, 0x95, 0x2d, 0x4a, 0xef, 0x57, 0x68, 0x35, 0xca, 0x79, 0xf5, 0xc6, 0x4f, 0xc5, 0x6e, 0xa2,
	0xd9, 0xce, 0xc2, 0x0b, 0xb0, 0x31, 0xa3, 0x48, 0x68, 0x80, 0x47, 0xc8, 0xf7, 0x03, 0x3e, 0x24,
	0x3d, 0x9b, 0xf6, 0x08, 0x13, 0x54, 0x50, 0x12, 0x7d, 0xe8, 0x6c, 0xf9, 0x81, 0x05, 0x95, 0xab,
	0x9e, 0x7a, 0x0a, 0x0d, 0xb0, 0x79, 0x8b, 0x8e, 0x85, 0xf9, 0x76, 0x5c, 0x90, 0x9f, 0x35, 0x24,
	0xe1, 0x0e, 0xf8, 0xe8, 0xf0, 0xf4, 0xa4, 0x79, 0xda, 0x30, 0x1b, 0x6d, 0xdb, 0x6c, 0x54, 0x6b,
	0x5f, 0x98, 0x27, 0x91, 0xd9, 0x6a, 0x57, 0xdb, 0xa6, 0xfd, 0xb2, 0xd1, 0x6a, 0x9a, 0x87, 0xf5,
	0xe7, 0x75, 0xf3, 0x28, 0xb7, 0x04, 0xd7, 0xc0, 0x3b, 0x47, 0xf5, 0x56, 0x04, 0xc9, 0x69, 0x10,
	0x80, 0x55, 0x09, 0xcf, 0x65, 0x22, 0x47, 0xbd, 0x71, 0x6c, 0x5a, 0xf5, 0x76, 0x2e, 0x5b, 0xfb,
	0x3d, 0x03, 0xf6, 0x30, 0xf7, 0x16, 0xda, 0x97, 0xda, 0x7b, 0x37, 0xe6, 0x5c, 0x33, 0xfa, 0x57,
	0x69, 0x6a, 0xaf, 0xda, 0x09, 0x8f, 0xc3, 0x5d, 0xc4, 0x1c, 0x9d, 0x07, 0x8e, 0xe1, 0x10, 0x16,
	0xff, 0xc9, 0x18, 0xe3, 0x59, 0xf9, 0xdf, 0xfe, 0xa1, 0x9f, 0xaa, 0x85, 0x7f, 0x34, 0xed, 0xd7,
	0xcc, 0xee, 0xe7, 0x92, 0xf9, 0x30, 0x56, 0xa8, 0x4e, 0xc6, 0xa1, 0x54, 0x98, 0x9e, 0x88, 0xb3,
	0xfd, 0x5a, 0x14, 0xfa, 0x87, 0x82, 0x77, 0x62, 0x78, 0xe7, 0x2a, 0xbc, 0xa3, 0xe0, 0x9d, 0x33,
	0x99, 0xe9, 0x6d, 0x66, 0x4f, 0xc2, 0x2b, 0x95, 0x18, 0x5f, 0xa9, 0x5c, 0x0d, 0x88, 0xde, 0x65,
	0x44, 0xa5, 0x92, 0x84, 0x74, 0x57, 0xe3, 0xa2, 0x3e, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0xab,
	0x24, 0x6c, 0x62, 0x17, 0x0c, 0x00, 0x00,
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/tasks/v2beta2/task.proto

package tasks

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// The view specifies a subset of [Task][google.cloud.tasks.v2beta2.Task] data.
//
// When a task is returned in a response, not all
// information is retrieved by default because some data, such as
// payloads, might be desirable to return only when needed because
// of its large size or because of the sensitivity of data that it
// contains.
type Task_View int32

const (
	// Unspecified. Defaults to BASIC.
	Task_VIEW_UNSPECIFIED Task_View = 0
	// The basic view omits fields which can be large or can contain
	// sensitive data.
	//
	// This view does not include the
	// ([payload in AppEngineHttpRequest][google.cloud.tasks.v2beta2.AppEngineHttpRequest] and
	// [payload in PullMessage][google.cloud.tasks.v2beta2.PullMessage.payload]). These payloads are
	// desirable to return only when needed, because they can be large
	// and because of the sensitivity of the data that you choose to
	// store in it.
	Task_BASIC Task_View = 1
	// All information is returned.
	//
	// Authorization for [FULL][google.cloud.tasks.v2beta2.Task.View.FULL] requires
	// `cloudtasks.tasks.fullView` [Google IAM](https://cloud.google.com/iam/)
	// permission on the [Queue][google.cloud.tasks.v2beta2.Queue] resource.
	Task_FULL Task_View = 2
)

var Task_View_name = map[int32]string{
	0: "VIEW_UNSPECIFIED",
	1: "BASIC",
	2: "FULL",
}

var Task_View_value = map[string]int32{
	"VIEW_UNSPECIFIED": 0,
	"BASIC":            1,
	"FULL":             2,
}

func (x Task_View) String() string {
	return proto.EnumName(Task_View_name, int32(x))
}

func (Task_View) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3fffa1a9946502fd, []int{0, 0}
}

// A unit of scheduled work.
type Task struct {
	// Optionally caller-specified in [CreateTask][google.cloud.tasks.v2beta2.CloudTasks.CreateTask].
	//
	// The task name.
	//
	// The task name must have the following format:
	// `projects/PROJECT_ID/locations/LOCATION_ID/queues/QUEUE_ID/tasks/TASK_ID`
	//
	// * `PROJECT_ID` can contain letters ([A-Za-z]), numbers ([0-9]),
	//    hyphens (-), colons (:), or periods (.).
	//    For more information, see
	//    [Identifying
	//    projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects#identifying_projects)
	// * `LOCATION_ID` is the canonical ID for the task's location.
	//    The list of available locations can be obtained by calling
	//    [ListLocations][google.cloud.location.Locations.ListLocations].
	//    For more information, see https://cloud.google.com/about/locations/.
	// * `QUEUE_ID` can contain letters ([A-Za-z]), numbers ([0-9]), or
	//   hyphens (-). The maximum length is 100 characters.
	// * `TASK_ID` can contain only letters ([A-Za-z]), numbers ([0-9]),
	//   hyphens (-), or underscores (_). The maximum length is 500 characters.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required.
	//
	// The task's payload is used by the task's target to process the task.
	// A payload is valid only if it is compatible with the queue's target.
	//
	// Types that are valid to be assigned to PayloadType:
	//	*Task_AppEngineHttpRequest
	//	*Task_PullMessage
	PayloadType isTask_PayloadType `protobuf_oneof:"payload_type"`
	// The time when the task is scheduled to be attempted.
	//
	// For App Engine queues, this is when the task will be attempted or retried.
	//
	// For pull queues, this is the time when the task is available to
	// be leased; if a task is currently leased, this is the time when
	// the current lease expires, that is, the time that the task was
	// leased plus the [lease_duration][google.cloud.tasks.v2beta2.LeaseTasksRequest.lease_duration].
	//
	// `schedule_time` will be truncated to the nearest microsecond.
	ScheduleTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=schedule_time,json=scheduleTime,proto3" json:"schedule_time,omitempty"`
	// Output only. The time that the task was created.
	//
	// `create_time` will be truncated to the nearest second.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The task status.
	Status *TaskStatus `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	// Output only. The view specifies which subset of the [Task][google.cloud.tasks.v2beta2.Task] has
	// been returned.
	View                 Task_View `protobuf:"varint,8,opt,name=view,proto3,enum=google.cloud.tasks.v2beta2.Task_View" json:"view,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fffa1a9946502fd, []int{0}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isTask_PayloadType interface {
	isTask_PayloadType()
}

type Task_AppEngineHttpRequest struct {
	AppEngineHttpRequest *AppEngineHttpRequest `protobuf:"bytes,3,opt,name=app_engine_http_request,json=appEngineHttpRequest,proto3,oneof"`
}

type Task_PullMessage struct {
	PullMessage *PullMessage `protobuf:"bytes,4,opt,name=pull_message,json=pullMessage,proto3,oneof"`
}

func (*Task_AppEngineHttpRequest) isTask_PayloadType() {}

func (*Task_PullMessage) isTask_PayloadType() {}

func (m *Task) GetPayloadType() isTask_PayloadType {
	if m != nil {
		return m.PayloadType
	}
	return nil
}

func (m *Task) GetAppEngineHttpRequest() *AppEngineHttpRequest {
	if x, ok := m.GetPayloadType().(*Task_AppEngineHttpRequest); ok {
		return x.AppEngineHttpRequest
	}
	return nil
}

func (m *Task) GetPullMessage() *PullMessage {
	if x, ok := m.GetPayloadType().(*Task_PullMessage); ok {
		return x.PullMessage
	}
	return nil
}

func (m *Task) GetScheduleTime() *timestamp.Timestamp {
	if m != nil {
		return m.ScheduleTime
	}
	return nil
}

func (m *Task) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Task) GetStatus() *TaskStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Task) GetView() Task_View {
	if m != nil {
		return m.View
	}
	return Task_VIEW_UNSPECIFIED
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Task) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Task_AppEngineHttpRequest)(nil),
		(*Task_PullMessage)(nil),
	}
}

// Status of the task.
type TaskStatus struct {
	// Output only. The number of attempts dispatched.
	//
	// This count includes attempts which have been dispatched but haven't
	// received a response.
	AttemptDispatchCount int32 `protobuf:"varint,1,opt,name=attempt_dispatch_count,json=attemptDispatchCount,proto3" json:"attempt_dispatch_count,omitempty"`
	// Output only. The number of attempts which have received a response.
	//
	// This field is not calculated for [pull tasks][google.cloud.tasks.v2beta2.PullMessage].
	AttemptResponseCount int32 `protobuf:"varint,2,opt,name=attempt_response_count,json=attemptResponseCount,proto3" json:"attempt_response_count,omitempty"`
	// Output only. The status of the task's first attempt.
	//
	// Only [dispatch_time][google.cloud.tasks.v2beta2.AttemptStatus.dispatch_time] will be set.
	// The other [AttemptStatus][google.cloud.tasks.v2beta2.AttemptStatus] information is not retained by Cloud Tasks.
	//
	// This field is not calculated for [pull tasks][google.cloud.tasks.v2beta2.PullMessage].
	FirstAttemptStatus *AttemptStatus `protobuf:"bytes,3,opt,name=first_attempt_status,json=firstAttemptStatus,proto3" json:"first_attempt_status,omitempty"`
	// Output only. The status of the task's last attempt.
	//
	// This field is not calculated for [pull tasks][google.cloud.tasks.v2beta2.PullMessage].
	LastAttemptStatus    *AttemptStatus `protobuf:"bytes,4,opt,name=last_attempt_status,json=lastAttemptStatus,proto3" json:"last_attempt_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TaskStatus) Reset()         { *m = TaskStatus{} }
func (m *TaskStatus) String() string { return proto.CompactTextString(m) }
func (*TaskStatus) ProtoMessage()    {}
func (*TaskStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fffa1a9946502fd, []int{1}
}

func (m *TaskStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskStatus.Unmarshal(m, b)
}
func (m *TaskStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskStatus.Marshal(b, m, deterministic)
}
func (m *TaskStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskStatus.Merge(m, src)
}
func (m *TaskStatus) XXX_Size() int {
	return xxx_messageInfo_TaskStatus.Size(m)
}
func (m *TaskStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskStatus.DiscardUnknown(m)
}

var xxx_messageInfo_TaskStatus proto.InternalMessageInfo

func (m *TaskStatus) GetAttemptDispatchCount() int32 {
	if m != nil {
		return m.AttemptDispatchCount
	}
	return 0
}

func (m *TaskStatus) GetAttemptResponseCount() int32 {
	if m != nil {
		return m.AttemptResponseCount
	}
	return 0
}

func (m *TaskStatus) GetFirstAttemptStatus() *AttemptStatus {
	if m != nil {
		return m.FirstAttemptStatus
	}
	return nil
}

func (m *TaskStatus) GetLastAttemptStatus() *AttemptStatus {
	if m != nil {
		return m.LastAttemptStatus
	}
	return nil
}

// The status of a task attempt.
type AttemptStatus struct {
	// Output only. The time that this attempt was scheduled.
	//
	// `schedule_time` will be truncated to the nearest microsecond.
	ScheduleTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=schedule_time,json=scheduleTime,proto3" json:"schedule_time,omitempty"`
	// Output only. The time that this attempt was dispatched.
	//
	// `dispatch_time` will be truncated to the nearest microsecond.
	DispatchTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=dispatch_time,json=dispatchTime,proto3" json:"dispatch_time,omitempty"`
	// Output only. The time that this attempt response was received.
	//
	// `response_time` will be truncated to the nearest microsecond.
	ResponseTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=response_time,json=responseTime,proto3" json:"response_time,omitempty"`
	// Output only. The response from the target for this attempt.
	//
	// If the task has not been attempted or the task is currently running
	// then the response status is unset.
	ResponseStatus       *status.Status `protobuf:"bytes,4,opt,name=response_status,json=responseStatus,proto3" json:"response_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AttemptStatus) Reset()         { *m = AttemptStatus{} }
func (m *AttemptStatus) String() string { return proto.CompactTextString(m) }
func (*AttemptStatus) ProtoMessage()    {}
func (*AttemptStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fffa1a9946502fd, []int{2}
}

func (m *AttemptStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttemptStatus.Unmarshal(m, b)
}
func (m *AttemptStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttemptStatus.Marshal(b, m, deterministic)
}
func (m *AttemptStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttemptStatus.Merge(m, src)
}
func (m *AttemptStatus) XXX_Size() int {
	return xxx_messageInfo_AttemptStatus.Size(m)
}
func (m *AttemptStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_AttemptStatus.DiscardUnknown(m)
}

var xxx_messageInfo_AttemptStatus proto.InternalMessageInfo

func (m *AttemptStatus) GetScheduleTime() *timestamp.Timestamp {
	if m != nil {
		return m.ScheduleTime
	}
	return nil
}

func (m *AttemptStatus) GetDispatchTime() *timestamp.Timestamp {
	if m != nil {
		return m.DispatchTime
	}
	return nil
}

func (m *AttemptStatus) GetResponseTime() *timestamp.Timestamp {
	if m != nil {
		return m.ResponseTime
	}
	return nil
}

func (m *AttemptStatus) GetResponseStatus() *status.Status {
	if m != nil {
		return m.ResponseStatus
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.cloud.tasks.v2beta2.Task_View", Task_View_name, Task_View_value)
	proto.RegisterType((*Task)(nil), "google.cloud.tasks.v2beta2.Task")
	proto.RegisterType((*TaskStatus)(nil), "google.cloud.tasks.v2beta2.TaskStatus")
	proto.RegisterType((*AttemptStatus)(nil), "google.cloud.tasks.v2beta2.AttemptStatus")
}

func init() {
	proto.RegisterFile("google/cloud/tasks/v2beta2/task.proto", fileDescriptor_3fffa1a9946502fd)
}

var fileDescriptor_3fffa1a9946502fd = []byte{
	// 672 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xdf, 0x4e, 0xdb, 0x3c,
	0x18, 0xc6, 0x49, 0x09, 0x7c, 0x60, 0xfe, 0x7c, 0xfd, 0xfc, 0x55, 0xa3, 0xab, 0x26, 0x86, 0x2a,
	0x31, 0xd8, 0x49, 0xbc, 0x75, 0x3b, 0xd9, 0x90, 0x86, 0x68, 0x29, 0xa2, 0x12, 0x9b, 0xaa, 0xf0,
	0x67, 0xda, 0x76, 0x10, 0x99, 0xf4, 0x25, 0xcd, 0x48, 0x62, 0x13, 0x3b, 0x20, 0x54, 0xf5, 0x1a,
	0x76, 0x2b, 0xbb, 0xae, 0xdd, 0xc1, 0xce, 0xa6, 0xd8, 0x0e, 0x50, 0xc1, 0x5a, 0x71, 0x94, 0xbc,
	0x7e, 0x9f, 0xe7, 0xe7, 0xf8, 0xf5, 0xd3, 0xa2, 0xf5, 0x80, 0xb1, 0x20, 0x02, 0xe2, 0x47, 0x2c,
	0xeb, 0x11, 0x49, 0xc5, 0xb9, 0x20, 0x97, 0x8d, 0x53, 0x90, 0xb4, 0xa1, 0x2a, 0x87, 0xa7, 0x4c,
	0x32, 0x5c, 0xd3, 0x32, 0x47, 0xc9, 0x1c, 0x25, 0x73, 0x8c, 0xac, 0xf6, 0xd4, 0x20, 0x28, 0x0f,
	0x49, 0x0a, 0x82, 0x65, 0xa9, 0x0f, 0xda, 0x56, 0xdb, 0x18, 0x4b, 0x4f, 0x03, 0x90, 0x46, 0xf8,
	0xdc, 0x08, 0x55, 0x75, 0x9a, 0x9d, 0x11, 0x19, 0xc6, 0x20, 0x24, 0x8d, 0xb9, 0x11, 0xac, 0x18,
	0x41, 0xca, 0x7d, 0x22, 0x24, 0x95, 0x99, 0x30, 0x8d, 0x67, 0x77, 0x76, 0xa7, 0x49, 0xc2, 0x24,
	0x95, 0x21, 0x4b, 0x4c, 0xb7, 0xfe, 0xdb, 0x46, 0xf6, 0x11, 0x15, 0xe7, 0x18, 0x23, 0x3b, 0xa1,
	0x31, 0x54, 0xad, 0x35, 0x6b, 0x73, 0xde, 0x55, 0xef, 0x38, 0x44, 0x2b, 0x94, 0x73, 0x0f, 0x92,
	0x20, 0x4c, 0xc0, 0xeb, 0x4b, 0xc9, 0xbd, 0x14, 0x2e, 0x32, 0x10, 0xb2, 0x3a, 0xbd, 0x66, 0x6d,
	0x2e, 0x34, 0x5e, 0x39, 0x7f, 0x3f, 0xb6, 0xb3, 0xc3, 0x79, 0x5b, 0x39, 0xf7, 0xa5, 0xe4, 0xae,
	0xf6, 0xed, 0x4f, 0xb9, 0x15, 0xfa, 0xc0, 0x3a, 0x3e, 0x40, 0x8b, 0x3c, 0x8b, 0x22, 0x2f, 0x06,
	0x21, 0x68, 0x00, 0x55, 0x5b, 0xf1, 0x37, 0xc6, 0xf1, 0xbb, 0x59, 0x14, 0x7d, 0xd4, 0xf2, 0xfd,
	0x29, 0x77, 0x81, 0xdf, 0x96, 0x78, 0x1b, 0x2d, 0x09, 0xbf, 0x0f, 0xbd, 0x2c, 0x02, 0x2f, 0x1f,
	0x54, 0x75, 0x46, 0xe1, 0x6a, 0x05, 0xae, 0x98, 0xa2, 0x73, 0x54, 0x4c, 0xd1, 0x5d, 0x2c, 0x0c,
	0xf9, 0x12, 0xde, 0x42, 0x0b, 0x7e, 0x0a, 0x54, 0x1a, 0xfb, 0xec, 0x44, 0x3b, 0xd2, 0x72, 0x65,
	0xfe, 0x80, 0x66, 0xf5, 0x0d, 0x54, 0xff, 0x51, 0xbe, 0x17, 0xe3, 0x4e, 0x91, 0x0f, 0xff, 0x50,
	0xa9, 0x5d, 0xe3, 0xc2, 0xef, 0x90, 0x7d, 0x19, 0xc2, 0x55, 0x75, 0x6e, 0xcd, 0xda, 0x5c, 0x6e,
	0xac, 0x4f, 0x72, 0x3b, 0x27, 0x21, 0x5c, 0xb9, 0xca, 0x52, 0x7f, 0x8d, 0xec, 0xbc, 0xc2, 0x15,
	0x54, 0x3e, 0xe9, 0xb4, 0x3f, 0x7b, 0xc7, 0x9f, 0x0e, 0xbb, 0xed, 0x56, 0x67, 0xaf, 0xd3, 0xde,
	0x2d, 0x4f, 0xe1, 0x79, 0x34, 0xd3, 0xdc, 0x39, 0xec, 0xb4, 0xca, 0x16, 0x9e, 0x43, 0xf6, 0xde,
	0xf1, 0xc1, 0x41, 0xb9, 0xf4, 0xbe, 0xff, 0x6b, 0x07, 0xd0, 0xaa, 0x82, 0x6b, 0xb6, 0xde, 0x8e,
	0xf2, 0x50, 0x38, 0x3e, 0x8b, 0x89, 0x4a, 0x47, 0x8b, 0xa7, 0xec, 0x3b, 0xf8, 0x52, 0x90, 0x81,
	0x79, 0x1b, 0x92, 0x88, 0xf9, 0x3a, 0x4a, 0x64, 0x50, 0xbc, 0x0e, 0xc9, 0x45, 0x06, 0x19, 0x08,
	0x32, 0x50, 0xcf, 0xa1, 0x49, 0xf4, 0x20, 0x7f, 0x0c, 0x9b, 0xcb, 0x68, 0x91, 0xd3, 0xeb, 0x88,
	0xd1, 0x9e, 0x27, 0xaf, 0x39, 0xd4, 0x7f, 0x96, 0x10, 0xba, 0x3d, 0x3e, 0x7e, 0x8b, 0x9e, 0x50,
	0x29, 0x21, 0xe6, 0xd2, 0xeb, 0x85, 0x82, 0x53, 0xe9, 0xf7, 0x3d, 0x9f, 0x65, 0x89, 0x54, 0x99,
	0x9c, 0x71, 0x2b, 0xa6, 0xbb, 0x6b, 0x9a, 0xad, 0xbc, 0x77, 0xd7, 0x95, 0x82, 0xe0, 0x2c, 0x11,
	0x60, 0x5c, 0xa5, 0x11, 0x97, 0x6b, 0x9a, 0xda, 0xf5, 0x0d, 0x55, 0xce, 0xc2, 0x54, 0x48, 0xaf,
	0xf0, 0x9a, 0x0b, 0xd3, 0xb1, 0x7e, 0x39, 0x36, 0xd6, 0xda, 0x61, 0xee, 0x0c, 0x2b, 0xcc, 0xc8,
	0x1a, 0xfe, 0x82, 0xfe, 0x8f, 0xe8, 0x7d, 0xb6, 0xfd, 0x58, 0xf6, 0x7f, 0x39, 0x65, 0x64, 0xa9,
	0xfe, 0xa3, 0x84, 0x96, 0x46, 0x37, 0xbb, 0x17, 0x75, 0xeb, 0x91, 0x51, 0xdf, 0x46, 0x4b, 0x37,
	0xe3, 0x56, 0x80, 0xd2, 0x64, 0x40, 0x61, 0x28, 0x00, 0x37, 0x93, 0x57, 0x80, 0xe9, 0xc9, 0x80,
	0xc2, 0x60, 0x7e, 0x6c, 0xff, 0xde, 0x00, 0x46, 0x66, 0x85, 0x0b, 0x44, 0xca, 0x7d, 0xc7, 0x0c,
	0x65, 0xb9, 0x90, 0xea, 0xba, 0x99, 0xa0, 0x55, 0x9f, 0xc5, 0x63, 0x86, 0xda, 0x9c, 0xcf, 0x33,
	0xd6, 0xcd, 0x3f, 0xa2, 0x6b, 0x7d, 0xdd, 0x36, 0xc2, 0x80, 0x45, 0x34, 0x09, 0x1c, 0x96, 0x06,
	0x24, 0x80, 0x44, 0x7d, 0x22, 0xb9, 0x0d, 0xfe, 0x43, 0xff, 0xc7, 0x5b, 0xaa, 0x3a, 0x9d, 0x55,
	0xda, 0x37, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x93, 0x3e, 0x7d, 0x18, 0x06, 0x00, 0x00,
}

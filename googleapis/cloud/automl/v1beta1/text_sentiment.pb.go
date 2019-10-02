// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/text_sentiment.proto

package automl

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

// Contains annotation details specific to text sentiment.
type TextSentimentAnnotation struct {
	// Output only. The sentiment with the semantic, as given to the
	// [AutoMl.ImportData][google.cloud.automl.v1beta1.AutoMl.ImportData] when populating the dataset from which the model used
	// for the prediction had been trained.
	// The sentiment values are between 0 and
	// Dataset.text_sentiment_dataset_metadata.sentiment_max (inclusive),
	// with higher value meaning more positive sentiment. They are completely
	// relative, i.e. 0 means least positive sentiment and sentiment_max means
	// the most positive from the sentiments present in the train data. Therefore
	//  e.g. if train data had only negative sentiment, then sentiment_max, would
	// be still negative (although least negative).
	// The sentiment shouldn't be confused with "score" or "magnitude"
	// from the previous Natural Language Sentiment Analysis API.
	Sentiment            int32    `protobuf:"varint,1,opt,name=sentiment,proto3" json:"sentiment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextSentimentAnnotation) Reset()         { *m = TextSentimentAnnotation{} }
func (m *TextSentimentAnnotation) String() string { return proto.CompactTextString(m) }
func (*TextSentimentAnnotation) ProtoMessage()    {}
func (*TextSentimentAnnotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c3765cde9ff2509, []int{0}
}

func (m *TextSentimentAnnotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextSentimentAnnotation.Unmarshal(m, b)
}
func (m *TextSentimentAnnotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextSentimentAnnotation.Marshal(b, m, deterministic)
}
func (m *TextSentimentAnnotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextSentimentAnnotation.Merge(m, src)
}
func (m *TextSentimentAnnotation) XXX_Size() int {
	return xxx_messageInfo_TextSentimentAnnotation.Size(m)
}
func (m *TextSentimentAnnotation) XXX_DiscardUnknown() {
	xxx_messageInfo_TextSentimentAnnotation.DiscardUnknown(m)
}

var xxx_messageInfo_TextSentimentAnnotation proto.InternalMessageInfo

func (m *TextSentimentAnnotation) GetSentiment() int32 {
	if m != nil {
		return m.Sentiment
	}
	return 0
}

// Model evaluation metrics for text sentiment problems.
type TextSentimentEvaluationMetrics struct {
	// Output only. Precision.
	Precision float32 `protobuf:"fixed32,1,opt,name=precision,proto3" json:"precision,omitempty"`
	// Output only. Recall.
	Recall float32 `protobuf:"fixed32,2,opt,name=recall,proto3" json:"recall,omitempty"`
	// Output only. The harmonic mean of recall and precision.
	F1Score float32 `protobuf:"fixed32,3,opt,name=f1_score,json=f1Score,proto3" json:"f1_score,omitempty"`
	// Output only. Mean absolute error. Only set for the overall model
	// evaluation, not for evaluation of a single annotation spec.
	MeanAbsoluteError float32 `protobuf:"fixed32,4,opt,name=mean_absolute_error,json=meanAbsoluteError,proto3" json:"mean_absolute_error,omitempty"`
	// Output only. Mean squared error. Only set for the overall model
	// evaluation, not for evaluation of a single annotation spec.
	MeanSquaredError float32 `protobuf:"fixed32,5,opt,name=mean_squared_error,json=meanSquaredError,proto3" json:"mean_squared_error,omitempty"`
	// Output only. Linear weighted kappa. Only set for the overall model
	// evaluation, not for evaluation of a single annotation spec.
	LinearKappa float32 `protobuf:"fixed32,6,opt,name=linear_kappa,json=linearKappa,proto3" json:"linear_kappa,omitempty"`
	// Output only. Quadratic weighted kappa. Only set for the overall model
	// evaluation, not for evaluation of a single annotation spec.
	QuadraticKappa float32 `protobuf:"fixed32,7,opt,name=quadratic_kappa,json=quadraticKappa,proto3" json:"quadratic_kappa,omitempty"`
	// Output only. Confusion matrix of the evaluation.
	// Only set for the overall model evaluation, not for evaluation of a single
	// annotation spec.
	ConfusionMatrix *ClassificationEvaluationMetrics_ConfusionMatrix `protobuf:"bytes,8,opt,name=confusion_matrix,json=confusionMatrix,proto3" json:"confusion_matrix,omitempty"`
	// Output only. The annotation spec ids used for this evaluation.
	// Deprecated .
	AnnotationSpecId     []string `protobuf:"bytes,9,rep,name=annotation_spec_id,json=annotationSpecId,proto3" json:"annotation_spec_id,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextSentimentEvaluationMetrics) Reset()         { *m = TextSentimentEvaluationMetrics{} }
func (m *TextSentimentEvaluationMetrics) String() string { return proto.CompactTextString(m) }
func (*TextSentimentEvaluationMetrics) ProtoMessage()    {}
func (*TextSentimentEvaluationMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c3765cde9ff2509, []int{1}
}

func (m *TextSentimentEvaluationMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextSentimentEvaluationMetrics.Unmarshal(m, b)
}
func (m *TextSentimentEvaluationMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextSentimentEvaluationMetrics.Marshal(b, m, deterministic)
}
func (m *TextSentimentEvaluationMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextSentimentEvaluationMetrics.Merge(m, src)
}
func (m *TextSentimentEvaluationMetrics) XXX_Size() int {
	return xxx_messageInfo_TextSentimentEvaluationMetrics.Size(m)
}
func (m *TextSentimentEvaluationMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_TextSentimentEvaluationMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_TextSentimentEvaluationMetrics proto.InternalMessageInfo

func (m *TextSentimentEvaluationMetrics) GetPrecision() float32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

func (m *TextSentimentEvaluationMetrics) GetRecall() float32 {
	if m != nil {
		return m.Recall
	}
	return 0
}

func (m *TextSentimentEvaluationMetrics) GetF1Score() float32 {
	if m != nil {
		return m.F1Score
	}
	return 0
}

func (m *TextSentimentEvaluationMetrics) GetMeanAbsoluteError() float32 {
	if m != nil {
		return m.MeanAbsoluteError
	}
	return 0
}

func (m *TextSentimentEvaluationMetrics) GetMeanSquaredError() float32 {
	if m != nil {
		return m.MeanSquaredError
	}
	return 0
}

func (m *TextSentimentEvaluationMetrics) GetLinearKappa() float32 {
	if m != nil {
		return m.LinearKappa
	}
	return 0
}

func (m *TextSentimentEvaluationMetrics) GetQuadraticKappa() float32 {
	if m != nil {
		return m.QuadraticKappa
	}
	return 0
}

func (m *TextSentimentEvaluationMetrics) GetConfusionMatrix() *ClassificationEvaluationMetrics_ConfusionMatrix {
	if m != nil {
		return m.ConfusionMatrix
	}
	return nil
}

// Deprecated: Do not use.
func (m *TextSentimentEvaluationMetrics) GetAnnotationSpecId() []string {
	if m != nil {
		return m.AnnotationSpecId
	}
	return nil
}

func init() {
	proto.RegisterType((*TextSentimentAnnotation)(nil), "google.cloud.automl.v1beta1.TextSentimentAnnotation")
	proto.RegisterType((*TextSentimentEvaluationMetrics)(nil), "google.cloud.automl.v1beta1.TextSentimentEvaluationMetrics")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/text_sentiment.proto", fileDescriptor_2c3765cde9ff2509)
}

var fileDescriptor_2c3765cde9ff2509 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0xd5, 0x94, 0x75, 0xab, 0x87, 0x58, 0x31, 0x12, 0x84, 0x6d, 0x1a, 0x65, 0x37, 0xf4,
	0x02, 0x25, 0x2b, 0x5c, 0x20, 0x85, 0xab, 0xb6, 0x9a, 0x10, 0x62, 0x95, 0x50, 0x8b, 0xb8, 0x40,
	0x95, 0xa2, 0x53, 0xe7, 0x34, 0xb2, 0x70, 0xec, 0xcc, 0x76, 0x46, 0xdf, 0x8c, 0x37, 0xe0, 0x82,
	0x47, 0xe1, 0x29, 0x50, 0xec, 0xac, 0x55, 0x05, 0xea, 0xa5, 0xcf, 0xf7, 0xfd, 0x7f, 0xdd, 0x1c,
	0x93, 0xab, 0x5c, 0xa9, 0x5c, 0x60, 0xcc, 0x84, 0xaa, 0xb2, 0x18, 0x2a, 0xab, 0x0a, 0x11, 0xdf,
	0x0d, 0x97, 0x68, 0x61, 0x18, 0x5b, 0x5c, 0xdb, 0xd4, 0xa0, 0xb4, 0xbc, 0x40, 0x69, 0xa3, 0x52,
	0x2b, 0xab, 0xe8, 0x99, 0x4f, 0x44, 0x2e, 0x11, 0xf9, 0x44, 0xd4, 0x24, 0x4e, 0xf7, 0xd6, 0x31,
	0x01, 0xc6, 0xf0, 0x15, 0x67, 0x60, 0xb9, 0x92, 0xbe, 0xee, 0xf4, 0xbc, 0x49, 0x40, 0xc9, 0x63,
	0x90, 0x52, 0x59, 0x07, 0x8d, 0xa7, 0x97, 0xef, 0xc8, 0xb3, 0x2f, 0xb8, 0xb6, 0xf3, 0xfb, 0x3b,
	0x8c, 0x36, 0x06, 0x3d, 0x27, 0xdd, 0xcd, 0xd5, 0xc2, 0x56, 0xbf, 0x35, 0x38, 0x98, 0x6d, 0x07,
	0x97, 0xbf, 0xda, 0xe4, 0x62, 0x27, 0x79, 0x7d, 0x07, 0xa2, 0x72, 0xc9, 0x29, 0x5a, 0xcd, 0x99,
	0xa9, 0x0b, 0x4a, 0x8d, 0x8c, 0x1b, 0xae, 0xa4, 0x2b, 0x08, 0x66, 0xdb, 0x01, 0x7d, 0x4a, 0x3a,
	0x1a, 0x19, 0x08, 0x11, 0x06, 0x0e, 0x35, 0x27, 0xfa, 0x9c, 0x1c, 0xad, 0x86, 0xa9, 0x61, 0x4a,
	0x63, 0xd8, 0x76, 0xe4, 0x70, 0x35, 0x9c, 0xd7, 0x47, 0x1a, 0x91, 0x27, 0x05, 0x82, 0x4c, 0x61,
	0x69, 0x94, 0xa8, 0x2c, 0xa6, 0xa8, 0xb5, 0xd2, 0xe1, 0x03, 0x67, 0x3d, 0xae, 0xd1, 0xa8, 0x21,
	0xd7, 0x35, 0xa0, 0xaf, 0x09, 0x75, 0xbe, 0xb9, 0xad, 0x40, 0x63, 0xd6, 0xe8, 0x07, 0x4e, 0xef,
	0xd5, 0x64, 0xee, 0x81, 0xb7, 0x5f, 0x92, 0x87, 0x82, 0x4b, 0x04, 0x9d, 0x7e, 0x87, 0xb2, 0x84,
	0xb0, 0xe3, 0xbc, 0x63, 0x3f, 0xfb, 0x54, 0x8f, 0xe8, 0x2b, 0x72, 0x72, 0x5b, 0x41, 0xa6, 0xc1,
	0x72, 0xd6, 0x58, 0x87, 0xce, 0x7a, 0xb4, 0x19, 0x7b, 0xf1, 0x07, 0xe9, 0x31, 0x25, 0x57, 0x55,
	0xfd, 0x4f, 0xd3, 0x02, 0xac, 0xe6, 0xeb, 0xf0, 0xa8, 0xdf, 0x1a, 0x1c, 0xbf, 0xb9, 0x89, 0xf6,
	0xac, 0x37, 0x9a, 0xec, 0x6c, 0xf0, 0x9f, 0x4f, 0x1a, 0x4d, 0xee, 0x4b, 0xa7, 0xae, 0x73, 0x76,
	0xc2, 0x76, 0x07, 0xf4, 0x8a, 0xd0, 0xed, 0x92, 0x53, 0x53, 0x22, 0x4b, 0x79, 0x16, 0x76, 0xfb,
	0xed, 0x41, 0x77, 0x1c, 0x84, 0xad, 0x59, 0x6f, 0x4b, 0xe7, 0x25, 0xb2, 0x8f, 0xd9, 0xf8, 0x67,
	0x8b, 0xbc, 0x60, 0xaa, 0xd8, 0x77, 0xad, 0x31, 0xdd, 0xd9, 0xf4, 0xe7, 0xfa, 0xe5, 0x7c, 0x1b,
	0x35, 0x81, 0x5c, 0x09, 0x90, 0x79, 0xa4, 0x74, 0x1e, 0xe7, 0x28, 0xdd, 0xab, 0x8a, 0x3d, 0x82,
	0x92, 0x9b, 0xff, 0x3e, 0xd4, 0xf7, 0xfe, 0xf8, 0x3b, 0x38, 0xfb, 0xe0, 0xc4, 0xc5, 0xa4, 0x96,
	0x16, 0xa3, 0xca, 0xaa, 0xa9, 0x58, 0x7c, 0xf5, 0xd2, 0x9f, 0xe0, 0xc2, 0xd3, 0x24, 0x71, 0x38,
	0x49, 0x1c, 0xbf, 0x49, 0x92, 0x46, 0x58, 0x76, 0xdc, 0x8f, 0xbd, 0xfd, 0x1b, 0x00, 0x00, 0xff,
	0xff, 0xae, 0xf0, 0x52, 0x23, 0x63, 0x03, 0x00, 0x00,
}

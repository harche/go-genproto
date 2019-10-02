// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/query_error.proto

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

// Enum describing possible query errors.
type QueryErrorEnum_QueryError int32

const (
	// Name unspecified.
	QueryErrorEnum_UNSPECIFIED QueryErrorEnum_QueryError = 0
	// The received error code is not known in this version.
	QueryErrorEnum_UNKNOWN QueryErrorEnum_QueryError = 1
	// Returned if all other query error reasons are not applicable.
	QueryErrorEnum_QUERY_ERROR QueryErrorEnum_QueryError = 50
	// A condition used in the query references an invalid enum constant.
	QueryErrorEnum_BAD_ENUM_CONSTANT QueryErrorEnum_QueryError = 18
	// Query contains an invalid escape sequence.
	QueryErrorEnum_BAD_ESCAPE_SEQUENCE QueryErrorEnum_QueryError = 7
	// Field name is invalid.
	QueryErrorEnum_BAD_FIELD_NAME QueryErrorEnum_QueryError = 12
	// Limit value is invalid (i.e. not a number)
	QueryErrorEnum_BAD_LIMIT_VALUE QueryErrorEnum_QueryError = 15
	// Encountered number can not be parsed.
	QueryErrorEnum_BAD_NUMBER QueryErrorEnum_QueryError = 5
	// Invalid operator encountered.
	QueryErrorEnum_BAD_OPERATOR QueryErrorEnum_QueryError = 3
	// Parameter unknown or not supported.
	QueryErrorEnum_BAD_PARAMETER_NAME QueryErrorEnum_QueryError = 61
	// Parameter have invalid value.
	QueryErrorEnum_BAD_PARAMETER_VALUE QueryErrorEnum_QueryError = 62
	// Invalid resource type was specified in the FROM clause.
	QueryErrorEnum_BAD_RESOURCE_TYPE_IN_FROM_CLAUSE QueryErrorEnum_QueryError = 45
	// Non-ASCII symbol encountered outside of strings.
	QueryErrorEnum_BAD_SYMBOL QueryErrorEnum_QueryError = 2
	// Value is invalid.
	QueryErrorEnum_BAD_VALUE QueryErrorEnum_QueryError = 4
	// Date filters fail to restrict date to a range smaller than 31 days.
	// Applicable if the query is segmented by date.
	QueryErrorEnum_DATE_RANGE_TOO_WIDE QueryErrorEnum_QueryError = 36
	// Filters on date/week/month/quarter have a start date after
	// end date.
	QueryErrorEnum_DATE_RANGE_TOO_NARROW QueryErrorEnum_QueryError = 60
	// Expected AND between values with BETWEEN operator.
	QueryErrorEnum_EXPECTED_AND QueryErrorEnum_QueryError = 30
	// Expecting ORDER BY to have BY.
	QueryErrorEnum_EXPECTED_BY QueryErrorEnum_QueryError = 14
	// There was no dimension field selected.
	QueryErrorEnum_EXPECTED_DIMENSION_FIELD_IN_SELECT_CLAUSE QueryErrorEnum_QueryError = 37
	// Missing filters on date related fields.
	QueryErrorEnum_EXPECTED_FILTERS_ON_DATE_RANGE QueryErrorEnum_QueryError = 55
	// Missing FROM clause.
	QueryErrorEnum_EXPECTED_FROM QueryErrorEnum_QueryError = 44
	// The operator used in the conditions requires the value to be a list.
	QueryErrorEnum_EXPECTED_LIST QueryErrorEnum_QueryError = 41
	// Fields used in WHERE or ORDER BY clauses are missing from the SELECT
	// clause.
	QueryErrorEnum_EXPECTED_REFERENCED_FIELD_IN_SELECT_CLAUSE QueryErrorEnum_QueryError = 16
	// SELECT is missing at the beginning of query.
	QueryErrorEnum_EXPECTED_SELECT QueryErrorEnum_QueryError = 13
	// A list was passed as a value to a condition whose operator expects a
	// single value.
	QueryErrorEnum_EXPECTED_SINGLE_VALUE QueryErrorEnum_QueryError = 42
	// Missing one or both values with BETWEEN operator.
	QueryErrorEnum_EXPECTED_VALUE_WITH_BETWEEN_OPERATOR QueryErrorEnum_QueryError = 29
	// Invalid date format. Expected 'YYYY-MM-DD'.
	QueryErrorEnum_INVALID_DATE_FORMAT QueryErrorEnum_QueryError = 38
	// Value passed was not a string when it should have been. I.e., it was a
	// number or unquoted literal.
	QueryErrorEnum_INVALID_STRING_VALUE QueryErrorEnum_QueryError = 57
	// A String value passed to the BETWEEN operator does not parse as a date.
	QueryErrorEnum_INVALID_VALUE_WITH_BETWEEN_OPERATOR QueryErrorEnum_QueryError = 26
	// The value passed to the DURING operator is not a Date range literal
	QueryErrorEnum_INVALID_VALUE_WITH_DURING_OPERATOR QueryErrorEnum_QueryError = 22
	// A non-string value was passed to the LIKE operator.
	QueryErrorEnum_INVALID_VALUE_WITH_LIKE_OPERATOR QueryErrorEnum_QueryError = 56
	// An operator was provided that is inapplicable to the field being
	// filtered.
	QueryErrorEnum_OPERATOR_FIELD_MISMATCH QueryErrorEnum_QueryError = 35
	// A Condition was found with an empty list.
	QueryErrorEnum_PROHIBITED_EMPTY_LIST_IN_CONDITION QueryErrorEnum_QueryError = 28
	// A condition used in the query references an unsupported enum constant.
	QueryErrorEnum_PROHIBITED_ENUM_CONSTANT QueryErrorEnum_QueryError = 54
	// Fields that are not allowed to be selected together were included in
	// the SELECT clause.
	QueryErrorEnum_PROHIBITED_FIELD_COMBINATION_IN_SELECT_CLAUSE QueryErrorEnum_QueryError = 31
	// A field that is not orderable was included in the ORDER BY clause.
	QueryErrorEnum_PROHIBITED_FIELD_IN_ORDER_BY_CLAUSE QueryErrorEnum_QueryError = 40
	// A field that is not selectable was included in the SELECT clause.
	QueryErrorEnum_PROHIBITED_FIELD_IN_SELECT_CLAUSE QueryErrorEnum_QueryError = 23
	// A field that is not filterable was included in the WHERE clause.
	QueryErrorEnum_PROHIBITED_FIELD_IN_WHERE_CLAUSE QueryErrorEnum_QueryError = 24
	// Resource type specified in the FROM clause is not supported by this
	// service.
	QueryErrorEnum_PROHIBITED_RESOURCE_TYPE_IN_FROM_CLAUSE QueryErrorEnum_QueryError = 43
	// A field that comes from an incompatible resource was included in the
	// SELECT clause.
	QueryErrorEnum_PROHIBITED_RESOURCE_TYPE_IN_SELECT_CLAUSE QueryErrorEnum_QueryError = 48
	// A field that comes from an incompatible resource was included in the
	// WHERE clause.
	QueryErrorEnum_PROHIBITED_RESOURCE_TYPE_IN_WHERE_CLAUSE QueryErrorEnum_QueryError = 58
	// A metric incompatible with the main resource or other selected
	// segmenting resources was included in the SELECT or WHERE clause.
	QueryErrorEnum_PROHIBITED_METRIC_IN_SELECT_OR_WHERE_CLAUSE QueryErrorEnum_QueryError = 49
	// A segment incompatible with the main resource or other selected
	// segmenting resources was included in the SELECT or WHERE clause.
	QueryErrorEnum_PROHIBITED_SEGMENT_IN_SELECT_OR_WHERE_CLAUSE QueryErrorEnum_QueryError = 51
	// A segment in the SELECT clause is incompatible with a metric in the
	// SELECT or WHERE clause.
	QueryErrorEnum_PROHIBITED_SEGMENT_WITH_METRIC_IN_SELECT_OR_WHERE_CLAUSE QueryErrorEnum_QueryError = 53
	// The value passed to the limit clause is too low.
	QueryErrorEnum_LIMIT_VALUE_TOO_LOW QueryErrorEnum_QueryError = 25
	// Query has a string containing a newline character.
	QueryErrorEnum_PROHIBITED_NEWLINE_IN_STRING QueryErrorEnum_QueryError = 8
	// List contains values of different types.
	QueryErrorEnum_PROHIBITED_VALUE_COMBINATION_IN_LIST QueryErrorEnum_QueryError = 10
	// The values passed to the BETWEEN operator are not of the same type.
	QueryErrorEnum_PROHIBITED_VALUE_COMBINATION_WITH_BETWEEN_OPERATOR QueryErrorEnum_QueryError = 21
	// Query contains unterminated string.
	QueryErrorEnum_STRING_NOT_TERMINATED QueryErrorEnum_QueryError = 6
	// Too many segments are specified in SELECT clause.
	QueryErrorEnum_TOO_MANY_SEGMENTS QueryErrorEnum_QueryError = 34
	// Query is incomplete and cannot be parsed.
	QueryErrorEnum_UNEXPECTED_END_OF_QUERY QueryErrorEnum_QueryError = 9
	// FROM clause cannot be specified in this query.
	QueryErrorEnum_UNEXPECTED_FROM_CLAUSE QueryErrorEnum_QueryError = 47
	// Query contains one or more unrecognized fields.
	QueryErrorEnum_UNRECOGNIZED_FIELD QueryErrorEnum_QueryError = 32
	// Query has an unexpected extra part.
	QueryErrorEnum_UNEXPECTED_INPUT QueryErrorEnum_QueryError = 11
	// Metrics cannot be requested for a manager account. To retrieve metrics,
	// issue separate requests against each client account under the manager
	// account.
	QueryErrorEnum_REQUESTED_METRICS_FOR_MANAGER QueryErrorEnum_QueryError = 59
)

var QueryErrorEnum_QueryError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	50: "QUERY_ERROR",
	18: "BAD_ENUM_CONSTANT",
	7:  "BAD_ESCAPE_SEQUENCE",
	12: "BAD_FIELD_NAME",
	15: "BAD_LIMIT_VALUE",
	5:  "BAD_NUMBER",
	3:  "BAD_OPERATOR",
	61: "BAD_PARAMETER_NAME",
	62: "BAD_PARAMETER_VALUE",
	45: "BAD_RESOURCE_TYPE_IN_FROM_CLAUSE",
	2:  "BAD_SYMBOL",
	4:  "BAD_VALUE",
	36: "DATE_RANGE_TOO_WIDE",
	60: "DATE_RANGE_TOO_NARROW",
	30: "EXPECTED_AND",
	14: "EXPECTED_BY",
	37: "EXPECTED_DIMENSION_FIELD_IN_SELECT_CLAUSE",
	55: "EXPECTED_FILTERS_ON_DATE_RANGE",
	44: "EXPECTED_FROM",
	41: "EXPECTED_LIST",
	16: "EXPECTED_REFERENCED_FIELD_IN_SELECT_CLAUSE",
	13: "EXPECTED_SELECT",
	42: "EXPECTED_SINGLE_VALUE",
	29: "EXPECTED_VALUE_WITH_BETWEEN_OPERATOR",
	38: "INVALID_DATE_FORMAT",
	57: "INVALID_STRING_VALUE",
	26: "INVALID_VALUE_WITH_BETWEEN_OPERATOR",
	22: "INVALID_VALUE_WITH_DURING_OPERATOR",
	56: "INVALID_VALUE_WITH_LIKE_OPERATOR",
	35: "OPERATOR_FIELD_MISMATCH",
	28: "PROHIBITED_EMPTY_LIST_IN_CONDITION",
	54: "PROHIBITED_ENUM_CONSTANT",
	31: "PROHIBITED_FIELD_COMBINATION_IN_SELECT_CLAUSE",
	40: "PROHIBITED_FIELD_IN_ORDER_BY_CLAUSE",
	23: "PROHIBITED_FIELD_IN_SELECT_CLAUSE",
	24: "PROHIBITED_FIELD_IN_WHERE_CLAUSE",
	43: "PROHIBITED_RESOURCE_TYPE_IN_FROM_CLAUSE",
	48: "PROHIBITED_RESOURCE_TYPE_IN_SELECT_CLAUSE",
	58: "PROHIBITED_RESOURCE_TYPE_IN_WHERE_CLAUSE",
	49: "PROHIBITED_METRIC_IN_SELECT_OR_WHERE_CLAUSE",
	51: "PROHIBITED_SEGMENT_IN_SELECT_OR_WHERE_CLAUSE",
	53: "PROHIBITED_SEGMENT_WITH_METRIC_IN_SELECT_OR_WHERE_CLAUSE",
	25: "LIMIT_VALUE_TOO_LOW",
	8:  "PROHIBITED_NEWLINE_IN_STRING",
	10: "PROHIBITED_VALUE_COMBINATION_IN_LIST",
	21: "PROHIBITED_VALUE_COMBINATION_WITH_BETWEEN_OPERATOR",
	6:  "STRING_NOT_TERMINATED",
	34: "TOO_MANY_SEGMENTS",
	9:  "UNEXPECTED_END_OF_QUERY",
	47: "UNEXPECTED_FROM_CLAUSE",
	32: "UNRECOGNIZED_FIELD",
	11: "UNEXPECTED_INPUT",
	59: "REQUESTED_METRICS_FOR_MANAGER",
}

var QueryErrorEnum_QueryError_value = map[string]int32{
	"UNSPECIFIED":                      0,
	"UNKNOWN":                          1,
	"QUERY_ERROR":                      50,
	"BAD_ENUM_CONSTANT":                18,
	"BAD_ESCAPE_SEQUENCE":              7,
	"BAD_FIELD_NAME":                   12,
	"BAD_LIMIT_VALUE":                  15,
	"BAD_NUMBER":                       5,
	"BAD_OPERATOR":                     3,
	"BAD_PARAMETER_NAME":               61,
	"BAD_PARAMETER_VALUE":              62,
	"BAD_RESOURCE_TYPE_IN_FROM_CLAUSE": 45,
	"BAD_SYMBOL":                       2,
	"BAD_VALUE":                        4,
	"DATE_RANGE_TOO_WIDE":              36,
	"DATE_RANGE_TOO_NARROW":            60,
	"EXPECTED_AND":                     30,
	"EXPECTED_BY":                      14,
	"EXPECTED_DIMENSION_FIELD_IN_SELECT_CLAUSE":                37,
	"EXPECTED_FILTERS_ON_DATE_RANGE":                           55,
	"EXPECTED_FROM":                                            44,
	"EXPECTED_LIST":                                            41,
	"EXPECTED_REFERENCED_FIELD_IN_SELECT_CLAUSE":               16,
	"EXPECTED_SELECT":                                          13,
	"EXPECTED_SINGLE_VALUE":                                    42,
	"EXPECTED_VALUE_WITH_BETWEEN_OPERATOR":                     29,
	"INVALID_DATE_FORMAT":                                      38,
	"INVALID_STRING_VALUE":                                     57,
	"INVALID_VALUE_WITH_BETWEEN_OPERATOR":                      26,
	"INVALID_VALUE_WITH_DURING_OPERATOR":                       22,
	"INVALID_VALUE_WITH_LIKE_OPERATOR":                         56,
	"OPERATOR_FIELD_MISMATCH":                                  35,
	"PROHIBITED_EMPTY_LIST_IN_CONDITION":                       28,
	"PROHIBITED_ENUM_CONSTANT":                                 54,
	"PROHIBITED_FIELD_COMBINATION_IN_SELECT_CLAUSE":            31,
	"PROHIBITED_FIELD_IN_ORDER_BY_CLAUSE":                      40,
	"PROHIBITED_FIELD_IN_SELECT_CLAUSE":                        23,
	"PROHIBITED_FIELD_IN_WHERE_CLAUSE":                         24,
	"PROHIBITED_RESOURCE_TYPE_IN_FROM_CLAUSE":                  43,
	"PROHIBITED_RESOURCE_TYPE_IN_SELECT_CLAUSE":                48,
	"PROHIBITED_RESOURCE_TYPE_IN_WHERE_CLAUSE":                 58,
	"PROHIBITED_METRIC_IN_SELECT_OR_WHERE_CLAUSE":              49,
	"PROHIBITED_SEGMENT_IN_SELECT_OR_WHERE_CLAUSE":             51,
	"PROHIBITED_SEGMENT_WITH_METRIC_IN_SELECT_OR_WHERE_CLAUSE": 53,
	"LIMIT_VALUE_TOO_LOW":                                      25,
	"PROHIBITED_NEWLINE_IN_STRING":                             8,
	"PROHIBITED_VALUE_COMBINATION_IN_LIST":                     10,
	"PROHIBITED_VALUE_COMBINATION_WITH_BETWEEN_OPERATOR":       21,
	"STRING_NOT_TERMINATED":                                    6,
	"TOO_MANY_SEGMENTS":                                        34,
	"UNEXPECTED_END_OF_QUERY":                                  9,
	"UNEXPECTED_FROM_CLAUSE":                                   47,
	"UNRECOGNIZED_FIELD":                                       32,
	"UNEXPECTED_INPUT":                                         11,
	"REQUESTED_METRICS_FOR_MANAGER":                            59,
}

func (x QueryErrorEnum_QueryError) String() string {
	return proto.EnumName(QueryErrorEnum_QueryError_name, int32(x))
}

func (QueryErrorEnum_QueryError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_08b7bf5560d34721, []int{0, 0}
}

// Container for enum describing possible query errors.
type QueryErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryErrorEnum) Reset()         { *m = QueryErrorEnum{} }
func (m *QueryErrorEnum) String() string { return proto.CompactTextString(m) }
func (*QueryErrorEnum) ProtoMessage()    {}
func (*QueryErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_08b7bf5560d34721, []int{0}
}

func (m *QueryErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryErrorEnum.Unmarshal(m, b)
}
func (m *QueryErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryErrorEnum.Marshal(b, m, deterministic)
}
func (m *QueryErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryErrorEnum.Merge(m, src)
}
func (m *QueryErrorEnum) XXX_Size() int {
	return xxx_messageInfo_QueryErrorEnum.Size(m)
}
func (m *QueryErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_QueryErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.QueryErrorEnum_QueryError", QueryErrorEnum_QueryError_name, QueryErrorEnum_QueryError_value)
	proto.RegisterType((*QueryErrorEnum)(nil), "google.ads.googleads.v2.errors.QueryErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/query_error.proto", fileDescriptor_08b7bf5560d34721)
}

var fileDescriptor_08b7bf5560d34721 = []byte{
	// 981 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0x6f, 0x73, 0xd3, 0xc6,
	0x13, 0xfe, 0x25, 0xfc, 0x0a, 0x65, 0x43, 0x92, 0xe3, 0x20, 0x7f, 0x48, 0x43, 0x1a, 0x4c, 0x80,
	0x40, 0x12, 0x3b, 0x98, 0x29, 0xa5, 0x86, 0x76, 0xe6, 0x24, 0xad, 0x9d, 0x1b, 0xa4, 0x3b, 0xe5,
	0x74, 0x8a, 0x6b, 0x26, 0x33, 0x37, 0x69, 0x93, 0xf1, 0x64, 0x06, 0xac, 0xd4, 0x0e, 0xcc, 0xf4,
	0xeb, 0xf4, 0x65, 0x3f, 0x45, 0x5f, 0xf7, 0x6d, 0xbf, 0x05, 0x9f, 0xa2, 0x73, 0x92, 0x25, 0xcb,
	0x60, 0xdc, 0x57, 0xd2, 0xed, 0x3e, 0xcf, 0x73, 0x7b, 0x7b, 0xbb, 0xb7, 0xb0, 0xdf, 0x4d, 0x92,
	0xee, 0xdb, 0xb3, 0xda, 0xc9, 0xe9, 0xa0, 0x96, 0xfd, 0xda, 0xbf, 0x0f, 0xf5, 0xda, 0x59, 0xbf,
	0x9f, 0xf4, 0x07, 0xb5, 0xdf, 0xde, 0x9f, 0xf5, 0x7f, 0x37, 0xe9, 0xa2, 0x7a, 0xd1, 0x4f, 0x2e,
	0x13, 0xba, 0x91, 0xc1, 0xaa, 0x27, 0xa7, 0x83, 0x6a, 0xc1, 0xa8, 0x7e, 0xa8, 0x57, 0x33, 0xc6,
	0xda, 0x7a, 0xae, 0x78, 0x71, 0x5e, 0x3b, 0xe9, 0xf5, 0x92, 0xcb, 0x93, 0xcb, 0xf3, 0xa4, 0x37,
	0xc8, 0xd8, 0x95, 0x7f, 0xe6, 0x61, 0xe1, 0xd0, 0x6a, 0xa2, 0x45, 0x63, 0xef, 0xfd, 0xbb, 0xca,
	0x5f, 0xf3, 0x00, 0x23, 0x13, 0x5d, 0x84, 0xb9, 0x58, 0x44, 0x21, 0xba, 0xbc, 0xc9, 0xd1, 0x23,
	0xff, 0xa3, 0x73, 0x70, 0x2d, 0x16, 0xaf, 0x85, 0x6c, 0x0b, 0x32, 0x63, 0xbd, 0x87, 0x31, 0xaa,
	0x8e, 0x41, 0xa5, 0xa4, 0x22, 0x75, 0xba, 0x04, 0x37, 0x1d, 0xe6, 0x19, 0x14, 0x71, 0x60, 0x5c,
	0x29, 0x22, 0xcd, 0x84, 0x26, 0x94, 0xae, 0xc0, 0xad, 0xd4, 0x1c, 0xb9, 0x2c, 0x44, 0x13, 0xe1,
	0x61, 0x8c, 0xc2, 0x45, 0x72, 0x8d, 0x52, 0x58, 0xb0, 0x8e, 0x26, 0x47, 0xdf, 0x33, 0x82, 0x05,
	0x48, 0x6e, 0xd0, 0x5b, 0xb0, 0x68, 0x6d, 0x3e, 0x0f, 0xb8, 0x36, 0x47, 0xcc, 0x8f, 0x91, 0x2c,
	0xd2, 0x05, 0x00, 0x6b, 0x14, 0x71, 0xe0, 0xa0, 0x22, 0x5f, 0x51, 0x02, 0x37, 0xec, 0x5a, 0x86,
	0xa8, 0x98, 0x96, 0x8a, 0x5c, 0xa1, 0xcb, 0x40, 0xad, 0x25, 0x64, 0x8a, 0x05, 0xa8, 0x51, 0x65,
	0x72, 0x3f, 0xe6, 0x7b, 0x8f, 0xec, 0x99, 0xe4, 0x4f, 0x74, 0x0b, 0x36, 0xad, 0x43, 0x61, 0x24,
	0x63, 0xe5, 0xa2, 0xd1, 0x9d, 0x10, 0x0d, 0x17, 0xa6, 0xa9, 0x64, 0x60, 0x5c, 0x9f, 0xc5, 0x11,
	0x92, 0xbd, 0x7c, 0xe3, 0xa8, 0x13, 0x38, 0xd2, 0x27, 0xb3, 0x74, 0x1e, 0xae, 0xdb, 0x75, 0x26,
	0xf2, 0x7f, 0xab, 0xee, 0x31, 0x8d, 0x46, 0x31, 0xd1, 0x42, 0xa3, 0xa5, 0x34, 0x6d, 0xee, 0x21,
	0xd9, 0xa2, 0x77, 0x60, 0xe9, 0x13, 0x87, 0x60, 0x4a, 0xc9, 0x36, 0x79, 0x65, 0x63, 0xc7, 0x9f,
	0x43, 0x74, 0x35, 0x7a, 0x86, 0x09, 0x8f, 0x6c, 0xd8, 0x3c, 0x16, 0x16, 0xa7, 0x43, 0x16, 0xe8,
	0x1e, 0x3c, 0x2e, 0x0c, 0x1e, 0x0f, 0x50, 0x44, 0x5c, 0x8a, 0x61, 0x9a, 0xb8, 0x30, 0x11, 0xfa,
	0xe8, 0xea, 0x3c, 0xc8, 0x07, 0xb4, 0x02, 0x1b, 0x05, 0xbc, 0xc9, 0x7d, 0x8d, 0x2a, 0x32, 0x52,
	0x98, 0x51, 0x00, 0xe4, 0x7b, 0x7a, 0x13, 0xe6, 0x47, 0x18, 0x25, 0x03, 0xb2, 0x3b, 0x66, 0xf2,
	0x79, 0xa4, 0xc9, 0x63, 0x5a, 0x85, 0x27, 0x85, 0x49, 0x61, 0x13, 0x95, 0xbd, 0x28, 0xef, 0x4b,
	0x3b, 0x13, 0x7b, 0x59, 0x05, 0x3e, 0xf3, 0x91, 0x79, 0x7b, 0xf6, 0x91, 0x91, 0x8b, 0x96, 0x8f,
	0xc3, 0x7c, 0x3d, 0xa1, 0xdb, 0xb0, 0x55, 0xb8, 0x52, 0x9b, 0x69, 0x73, 0x7d, 0x60, 0x1c, 0xd4,
	0x6d, 0x44, 0x31, 0xba, 0xcf, 0xbb, 0x36, 0xb3, 0x5c, 0x1c, 0x31, 0x9f, 0x7b, 0xd9, 0x39, 0x9a,
	0x52, 0x05, 0x4c, 0x93, 0x87, 0x74, 0x15, 0x6e, 0xe7, 0x8e, 0x48, 0x2b, 0x2e, 0x5a, 0x43, 0xf1,
	0x1f, 0xe8, 0x23, 0xb8, 0x9f, 0x7b, 0xa6, 0x69, 0xaf, 0xd1, 0x87, 0x50, 0x99, 0x00, 0xf4, 0xe2,
	0x54, 0xad, 0xc0, 0x2d, 0xdb, 0x12, 0x99, 0x80, 0xf3, 0xf9, 0x6b, 0x1c, 0xa1, 0x5e, 0xd0, 0x6f,
	0x60, 0x25, 0x5f, 0x0d, 0x13, 0x15, 0xf0, 0x28, 0x60, 0xda, 0x3d, 0x20, 0xf7, 0xed, 0x56, 0xa1,
	0x92, 0x07, 0xdc, 0xe1, 0xf6, 0xc8, 0x18, 0x84, 0xba, 0x93, 0xe6, 0xda, 0x26, 0xd3, 0x95, 0xc2,
	0xe3, 0x9a, 0x4b, 0x41, 0xd6, 0xe9, 0x3a, 0xac, 0x96, 0x71, 0x63, 0x0d, 0xf4, 0x9c, 0x3e, 0x85,
	0xbd, 0x92, 0x37, 0xdb, 0xc4, 0x95, 0x81, 0xc3, 0x05, 0xb3, 0xfc, 0xcf, 0x6f, 0xe6, 0x5b, 0x9b,
	0x8c, 0xcf, 0x28, 0x5c, 0x18, 0xa9, 0x3c, 0x54, 0xc6, 0xe9, 0xe4, 0xc0, 0x6d, 0xfa, 0x00, 0xee,
	0x4d, 0x02, 0x8e, 0xeb, 0xad, 0xd8, 0x5c, 0x4c, 0x82, 0xb5, 0x0f, 0x50, 0x61, 0x8e, 0x5a, 0xa5,
	0x3b, 0xf0, 0xa8, 0x84, 0x9a, 0xda, 0x5b, 0x3b, 0xb6, 0xca, 0xa7, 0x81, 0xc7, 0x23, 0xd8, 0xa7,
	0xbb, 0xb0, 0x3d, 0x0d, 0x3e, 0x16, 0x49, 0x83, 0xd6, 0x60, 0xa7, 0x84, 0x0e, 0x50, 0x2b, 0xee,
	0x96, 0x54, 0xa5, 0x1a, 0x27, 0x3c, 0xa5, 0xfb, 0xb0, 0x5b, 0x22, 0x44, 0xd8, 0x0a, 0x50, 0xe8,
	0x29, 0x8c, 0x67, 0xf4, 0x15, 0xbc, 0x98, 0xc0, 0x48, 0x6b, 0xe4, 0x3f, 0xf7, 0xfb, 0xce, 0x16,
	0x78, 0xe9, 0x8d, 0x4b, 0x9f, 0x08, 0x5f, 0xb6, 0xc9, 0x1d, 0xba, 0x09, 0xeb, 0x25, 0x59, 0x81,
	0x6d, 0x9f, 0x8b, 0x2c, 0x21, 0x69, 0xb9, 0x93, 0xaf, 0x6d, 0x17, 0x95, 0x10, 0x19, 0xff, 0x93,
	0x72, 0x48, 0xfb, 0x19, 0xe8, 0x73, 0xa8, 0x4f, 0x45, 0x4e, 0xee, 0x90, 0x25, 0xdb, 0xc2, 0xc3,
	0xe6, 0x12, 0x52, 0x1b, 0x8d, 0x2a, 0xb0, 0x0c, 0xf4, 0xc8, 0x55, 0xfb, 0xc6, 0xdb, 0x58, 0x03,
	0x26, 0x3a, 0xf9, 0x99, 0x23, 0x52, 0xb1, 0x5d, 0x10, 0x8b, 0xa2, 0xb7, 0x51, 0x78, 0x46, 0x36,
	0x4d, 0x3a, 0x1d, 0xc8, 0x75, 0xba, 0x06, 0xcb, 0x25, 0x67, 0xb9, 0x0a, 0x6a, 0xf6, 0xe1, 0x8e,
	0x85, 0x42, 0x57, 0xb6, 0x04, 0x7f, 0x93, 0x97, 0x16, 0xd9, 0xa4, 0xb7, 0x81, 0x94, 0x38, 0x5c,
	0x84, 0xb1, 0x26, 0x73, 0xf4, 0x1e, 0xdc, 0x55, 0x76, 0x7e, 0x44, 0xa3, 0x5b, 0x8d, 0xec, 0xdb,
	0x60, 0xe3, 0x61, 0x2d, 0x54, 0xe4, 0xa5, 0xf3, 0x71, 0x06, 0x2a, 0xbf, 0x26, 0xef, 0xaa, 0xd3,
	0x47, 0xa3, 0xb3, 0x38, 0x1a, 0x73, 0xa1, 0x9d, 0x86, 0xe1, 0xcc, 0x1b, 0x6f, 0x48, 0xe9, 0x26,
	0x6f, 0x4f, 0x7a, 0xdd, 0x6a, 0xd2, 0xef, 0xd6, 0xba, 0x67, 0xbd, 0x74, 0x56, 0xe6, 0xf3, 0xf8,
	0xe2, 0x7c, 0xf0, 0xa5, 0xf1, 0xfc, 0x32, 0xfb, 0xfc, 0x31, 0x7b, 0xa5, 0xc5, 0xd8, 0x9f, 0xb3,
	0x1b, 0xad, 0x4c, 0x8c, 0x9d, 0x0e, 0xaa, 0xd9, 0xaf, 0xfd, 0x3b, 0xaa, 0x57, 0xd3, 0x2d, 0x07,
	0x7f, 0xe7, 0x80, 0x63, 0x76, 0x3a, 0x38, 0x2e, 0x00, 0xc7, 0x47, 0xf5, 0xe3, 0x0c, 0xf0, 0x71,
	0xb6, 0x92, 0x59, 0x1b, 0x0d, 0x76, 0x3a, 0x68, 0x34, 0x0a, 0x48, 0xa3, 0x71, 0x54, 0x6f, 0x34,
	0x32, 0xd0, 0x2f, 0x57, 0xd3, 0xe8, 0x9e, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x85, 0x61,
	0xc1, 0x3b, 0x08, 0x00, 0x00,
}

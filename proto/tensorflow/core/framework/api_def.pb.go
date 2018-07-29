// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/api_def.proto

package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ApiDef_Visibility int32

const (
	// Normally this is "VISIBLE" unless you are inheriting a
	// different value from another ApiDef.
	ApiDef_DEFAULT_VISIBILITY ApiDef_Visibility = 0
	// Publicly visible in the API.
	ApiDef_VISIBLE ApiDef_Visibility = 1
	// Do not include this op in the generated API. If visibility is
	// set to 'SKIP', other fields are ignored for this op.
	ApiDef_SKIP ApiDef_Visibility = 2
	// Hide this op by putting it into an internal namespace (or whatever
	// is appropriate in the target language).
	ApiDef_HIDDEN ApiDef_Visibility = 3
)

var ApiDef_Visibility_name = map[int32]string{
	0: "DEFAULT_VISIBILITY",
	1: "VISIBLE",
	2: "SKIP",
	3: "HIDDEN",
}
var ApiDef_Visibility_value = map[string]int32{
	"DEFAULT_VISIBILITY": 0,
	"VISIBLE":            1,
	"SKIP":               2,
	"HIDDEN":             3,
}

func (x ApiDef_Visibility) String() string {
	return proto.EnumName(ApiDef_Visibility_name, int32(x))
}
func (ApiDef_Visibility) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_api_def_459d9a384fa34718, []int{0, 0}
}

// Used to specify and override the default API & behavior in the
// generated code for client languages, from what you would get from
// the OpDef alone. There will be a set of ApiDefs that are common
// to all client languages, and another set per client language.
// The per-client-language ApiDefs will inherit values from the
// common ApiDefs which it can either replace or modify.
//
// We separate the API definition from the OpDef so we can evolve the
// API while remaining backwards compatible when interpretting old
// graphs.  Overrides go in an "api_def.pbtxt" file with a text-format
// ApiDefs message.
//
// WARNING: Be *very* careful changing the API for any existing op --
// you can change the semantics of existing code.  These changes may
// need to wait until a major release of TensorFlow to avoid breaking
// our compatibility promises.
type ApiDef struct {
	// Name of the op (in the OpDef) to specify the API for.
	GraphOpName string             `protobuf:"bytes,1,opt,name=graph_op_name,json=graphOpName,proto3" json:"graph_op_name,omitempty"`
	Visibility  ApiDef_Visibility  `protobuf:"varint,2,opt,name=visibility,proto3,enum=tensorflow.ApiDef_Visibility" json:"visibility,omitempty"`
	Endpoint    []*ApiDef_Endpoint `protobuf:"bytes,3,rep,name=endpoint,proto3" json:"endpoint,omitempty"`
	InArg       []*ApiDef_Arg      `protobuf:"bytes,4,rep,name=in_arg,json=inArg,proto3" json:"in_arg,omitempty"`
	OutArg      []*ApiDef_Arg      `protobuf:"bytes,5,rep,name=out_arg,json=outArg,proto3" json:"out_arg,omitempty"`
	// List of original in_arg names to specify new argument order.
	// Length of arg_order should be either empty to keep current order
	// or match size of in_arg.
	ArgOrder []string       `protobuf:"bytes,11,rep,name=arg_order,json=argOrder,proto3" json:"arg_order,omitempty"`
	Attr     []*ApiDef_Attr `protobuf:"bytes,6,rep,name=attr,proto3" json:"attr,omitempty"`
	// One-line human-readable description of what the Op does.
	Summary string `protobuf:"bytes,7,opt,name=summary,proto3" json:"summary,omitempty"`
	// Additional, longer human-readable description of what the Op does.
	Description string `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	// Modify an existing/inherited description by adding text to the beginning
	// or end.
	DescriptionPrefix    string   `protobuf:"bytes,9,opt,name=description_prefix,json=descriptionPrefix,proto3" json:"description_prefix,omitempty"`
	DescriptionSuffix    string   `protobuf:"bytes,10,opt,name=description_suffix,json=descriptionSuffix,proto3" json:"description_suffix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiDef) Reset()         { *m = ApiDef{} }
func (m *ApiDef) String() string { return proto.CompactTextString(m) }
func (*ApiDef) ProtoMessage()    {}
func (*ApiDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_def_459d9a384fa34718, []int{0}
}
func (m *ApiDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiDef.Unmarshal(m, b)
}
func (m *ApiDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiDef.Marshal(b, m, deterministic)
}
func (dst *ApiDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiDef.Merge(dst, src)
}
func (m *ApiDef) XXX_Size() int {
	return xxx_messageInfo_ApiDef.Size(m)
}
func (m *ApiDef) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiDef.DiscardUnknown(m)
}

var xxx_messageInfo_ApiDef proto.InternalMessageInfo

func (m *ApiDef) GetGraphOpName() string {
	if m != nil {
		return m.GraphOpName
	}
	return ""
}

func (m *ApiDef) GetVisibility() ApiDef_Visibility {
	if m != nil {
		return m.Visibility
	}
	return ApiDef_DEFAULT_VISIBILITY
}

func (m *ApiDef) GetEndpoint() []*ApiDef_Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *ApiDef) GetInArg() []*ApiDef_Arg {
	if m != nil {
		return m.InArg
	}
	return nil
}

func (m *ApiDef) GetOutArg() []*ApiDef_Arg {
	if m != nil {
		return m.OutArg
	}
	return nil
}

func (m *ApiDef) GetArgOrder() []string {
	if m != nil {
		return m.ArgOrder
	}
	return nil
}

func (m *ApiDef) GetAttr() []*ApiDef_Attr {
	if m != nil {
		return m.Attr
	}
	return nil
}

func (m *ApiDef) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *ApiDef) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ApiDef) GetDescriptionPrefix() string {
	if m != nil {
		return m.DescriptionPrefix
	}
	return ""
}

func (m *ApiDef) GetDescriptionSuffix() string {
	if m != nil {
		return m.DescriptionSuffix
	}
	return ""
}

// If you specify any endpoint, this will replace all of the
// inherited endpoints.  The first endpoint should be the
// "canonical" endpoint, and should not be deprecated (unless all
// endpoints are deprecated).
type ApiDef_Endpoint struct {
	// Name should be either like "CamelCaseName" or
	// "Package.CamelCaseName". Client-language-specific ApiDefs may
	// use a snake_case convention instead of CamelCase.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// First GraphDef version at which the op is disallowed.
	DeprecationVersion   int32    `protobuf:"varint,2,opt,name=deprecation_version,json=deprecationVersion,proto3" json:"deprecation_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiDef_Endpoint) Reset()         { *m = ApiDef_Endpoint{} }
func (m *ApiDef_Endpoint) String() string { return proto.CompactTextString(m) }
func (*ApiDef_Endpoint) ProtoMessage()    {}
func (*ApiDef_Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_def_459d9a384fa34718, []int{0, 0}
}
func (m *ApiDef_Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiDef_Endpoint.Unmarshal(m, b)
}
func (m *ApiDef_Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiDef_Endpoint.Marshal(b, m, deterministic)
}
func (dst *ApiDef_Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiDef_Endpoint.Merge(dst, src)
}
func (m *ApiDef_Endpoint) XXX_Size() int {
	return xxx_messageInfo_ApiDef_Endpoint.Size(m)
}
func (m *ApiDef_Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiDef_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_ApiDef_Endpoint proto.InternalMessageInfo

func (m *ApiDef_Endpoint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ApiDef_Endpoint) GetDeprecationVersion() int32 {
	if m != nil {
		return m.DeprecationVersion
	}
	return 0
}

type ApiDef_Arg struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Change the name used to access this arg in the API from what
	// is used in the GraphDef.  Note that these names in `backticks`
	// will also be replaced in the summary & description fields.
	RenameTo string `protobuf:"bytes,2,opt,name=rename_to,json=renameTo,proto3" json:"rename_to,omitempty"`
	// Note: this will replace any inherited arg doc. There is no
	// current way of modifying arg descriptions (other than replacing
	// them entirely) as can be done with op descriptions.
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiDef_Arg) Reset()         { *m = ApiDef_Arg{} }
func (m *ApiDef_Arg) String() string { return proto.CompactTextString(m) }
func (*ApiDef_Arg) ProtoMessage()    {}
func (*ApiDef_Arg) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_def_459d9a384fa34718, []int{0, 1}
}
func (m *ApiDef_Arg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiDef_Arg.Unmarshal(m, b)
}
func (m *ApiDef_Arg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiDef_Arg.Marshal(b, m, deterministic)
}
func (dst *ApiDef_Arg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiDef_Arg.Merge(dst, src)
}
func (m *ApiDef_Arg) XXX_Size() int {
	return xxx_messageInfo_ApiDef_Arg.Size(m)
}
func (m *ApiDef_Arg) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiDef_Arg.DiscardUnknown(m)
}

var xxx_messageInfo_ApiDef_Arg proto.InternalMessageInfo

func (m *ApiDef_Arg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ApiDef_Arg) GetRenameTo() string {
	if m != nil {
		return m.RenameTo
	}
	return ""
}

func (m *ApiDef_Arg) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Description of the graph-construction-time configuration of this
// Op.  That is to say, this describes the attr fields that will
// be specified in the NodeDef.
type ApiDef_Attr struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Change the name used to access this attr in the API from what
	// is used in the GraphDef.  Note that these names in `backticks`
	// will also be replaced in the summary & description fields.
	RenameTo string `protobuf:"bytes,2,opt,name=rename_to,json=renameTo,proto3" json:"rename_to,omitempty"`
	// Specify a new default value to use for this attr.  This default
	// will be used when creating new graphs, as opposed to the
	// default in the OpDef, which will be used when interpreting old
	// GraphDefs.
	DefaultValue *AttrValue `protobuf:"bytes,3,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	// Note: this will replace any inherited attr doc, there is no current
	// way of modifying attr descriptions as can be done with op descriptions.
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApiDef_Attr) Reset()         { *m = ApiDef_Attr{} }
func (m *ApiDef_Attr) String() string { return proto.CompactTextString(m) }
func (*ApiDef_Attr) ProtoMessage()    {}
func (*ApiDef_Attr) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_def_459d9a384fa34718, []int{0, 2}
}
func (m *ApiDef_Attr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiDef_Attr.Unmarshal(m, b)
}
func (m *ApiDef_Attr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiDef_Attr.Marshal(b, m, deterministic)
}
func (dst *ApiDef_Attr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiDef_Attr.Merge(dst, src)
}
func (m *ApiDef_Attr) XXX_Size() int {
	return xxx_messageInfo_ApiDef_Attr.Size(m)
}
func (m *ApiDef_Attr) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiDef_Attr.DiscardUnknown(m)
}

var xxx_messageInfo_ApiDef_Attr proto.InternalMessageInfo

func (m *ApiDef_Attr) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ApiDef_Attr) GetRenameTo() string {
	if m != nil {
		return m.RenameTo
	}
	return ""
}

func (m *ApiDef_Attr) GetDefaultValue() *AttrValue {
	if m != nil {
		return m.DefaultValue
	}
	return nil
}

func (m *ApiDef_Attr) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type ApiDefs struct {
	Op                   []*ApiDef `protobuf:"bytes,1,rep,name=op,proto3" json:"op,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ApiDefs) Reset()         { *m = ApiDefs{} }
func (m *ApiDefs) String() string { return proto.CompactTextString(m) }
func (*ApiDefs) ProtoMessage()    {}
func (*ApiDefs) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_def_459d9a384fa34718, []int{1}
}
func (m *ApiDefs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiDefs.Unmarshal(m, b)
}
func (m *ApiDefs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiDefs.Marshal(b, m, deterministic)
}
func (dst *ApiDefs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiDefs.Merge(dst, src)
}
func (m *ApiDefs) XXX_Size() int {
	return xxx_messageInfo_ApiDefs.Size(m)
}
func (m *ApiDefs) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiDefs.DiscardUnknown(m)
}

var xxx_messageInfo_ApiDefs proto.InternalMessageInfo

func (m *ApiDefs) GetOp() []*ApiDef {
	if m != nil {
		return m.Op
	}
	return nil
}

func init() {
	proto.RegisterType((*ApiDef)(nil), "tensorflow.ApiDef")
	proto.RegisterType((*ApiDef_Endpoint)(nil), "tensorflow.ApiDef.Endpoint")
	proto.RegisterType((*ApiDef_Arg)(nil), "tensorflow.ApiDef.Arg")
	proto.RegisterType((*ApiDef_Attr)(nil), "tensorflow.ApiDef.Attr")
	proto.RegisterType((*ApiDefs)(nil), "tensorflow.ApiDefs")
	proto.RegisterEnum("tensorflow.ApiDef_Visibility", ApiDef_Visibility_name, ApiDef_Visibility_value)
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/api_def.proto", fileDescriptor_api_def_459d9a384fa34718)
}

var fileDescriptor_api_def_459d9a384fa34718 = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x8b, 0xd3, 0x40,
	0x14, 0xc6, 0x4d, 0x93, 0x4d, 0x93, 0x97, 0x5d, 0xa9, 0x23, 0xae, 0x43, 0x17, 0xa1, 0xf4, 0x62,
	0x51, 0xda, 0xc2, 0x7a, 0x10, 0x04, 0x0f, 0x5d, 0x5a, 0x35, 0xb8, 0x6c, 0x4b, 0x5a, 0x8b, 0x9e,
	0xc2, 0x6c, 0x33, 0x89, 0x83, 0x6d, 0x66, 0x98, 0x4c, 0xba, 0xee, 0x1f, 0xe2, 0x7f, 0xea, 0xc1,
	0xa3, 0x64, 0xb2, 0xdb, 0xc6, 0xb6, 0x2c, 0x78, 0xcb, 0x9b, 0xf7, 0xfb, 0xbe, 0x0c, 0xdf, 0xbc,
	0x07, 0x2f, 0x15, 0x4d, 0x33, 0x2e, 0xe3, 0x25, 0xbf, 0xe9, 0x2f, 0xb8, 0xa4, 0xfd, 0x58, 0x92,
	0x15, 0xbd, 0xe1, 0xf2, 0x47, 0x9f, 0x08, 0x16, 0x46, 0x34, 0xee, 0x09, 0xc9, 0x15, 0x47, 0xb0,
	0x05, 0x9b, 0xaf, 0x1e, 0x10, 0x29, 0x25, 0xc3, 0x35, 0x59, 0xe6, 0xb4, 0xd4, 0xb5, 0x7f, 0xdb,
	0x60, 0x0f, 0x04, 0x1b, 0xd2, 0x18, 0xb5, 0xe1, 0x24, 0x91, 0x44, 0x7c, 0x0f, 0xb9, 0x08, 0x53,
	0xb2, 0xa2, 0xd8, 0x68, 0x19, 0x1d, 0x37, 0xf0, 0xf4, 0xe1, 0x58, 0x5c, 0x91, 0x15, 0x45, 0xef,
	0x01, 0xd6, 0x2c, 0x63, 0xd7, 0x6c, 0xc9, 0xd4, 0x2d, 0xae, 0xb5, 0x8c, 0xce, 0xe3, 0xf3, 0x17,
	0xbd, 0xed, 0xff, 0x7a, 0xa5, 0x57, 0x6f, 0xbe, 0x81, 0x82, 0x8a, 0x00, 0xbd, 0x05, 0x87, 0xa6,
	0x91, 0xe0, 0x2c, 0x55, 0xd8, 0x6c, 0x99, 0x1d, 0xef, 0xfc, 0xec, 0x80, 0x78, 0x74, 0x87, 0x04,
	0x1b, 0x18, 0x75, 0xc1, 0x66, 0x69, 0x48, 0x64, 0x82, 0x2d, 0x2d, 0x3b, 0x3d, 0x20, 0x1b, 0xc8,
	0x24, 0x38, 0x62, 0xe9, 0x40, 0x26, 0xa8, 0x0f, 0x75, 0x9e, 0x2b, 0xcd, 0x1f, 0x3d, 0xc8, 0xdb,
	0x3c, 0x57, 0x85, 0xe0, 0x0c, 0x5c, 0x22, 0x93, 0x90, 0xcb, 0x88, 0x4a, 0xec, 0xb5, 0xcc, 0x8e,
	0x1b, 0x38, 0x44, 0x26, 0xe3, 0xa2, 0x46, 0xaf, 0xc1, 0x2a, 0x72, 0xc3, 0xb6, 0xb6, 0x7a, 0x7e,
	0xc8, 0x4a, 0x29, 0x19, 0x68, 0x08, 0x61, 0xa8, 0x67, 0xf9, 0x6a, 0x45, 0xe4, 0x2d, 0xae, 0xeb,
	0xfc, 0xee, 0x4b, 0xd4, 0x02, 0x2f, 0xa2, 0xd9, 0x42, 0x32, 0xa1, 0x18, 0x4f, 0xb1, 0x53, 0xa6,
	0x5b, 0x39, 0x42, 0x5d, 0x40, 0x95, 0x32, 0x14, 0x92, 0xc6, 0xec, 0x27, 0x76, 0x35, 0xf8, 0xa4,
	0xd2, 0x99, 0xe8, 0xc6, 0x2e, 0x9e, 0xe5, 0x71, 0x81, 0xc3, 0x1e, 0x3e, 0xd5, 0x8d, 0xe6, 0x18,
	0x9c, 0xfb, 0x64, 0x11, 0x02, 0xab, 0xf2, 0xc4, 0xfa, 0x1b, 0xf5, 0xe1, 0x69, 0x44, 0x85, 0xa4,
	0x0b, 0xa2, 0xed, 0xd6, 0x54, 0x66, 0xc5, 0x3d, 0x8b, 0x47, 0x3e, 0x0a, 0x50, 0xa5, 0x35, 0x2f,
	0x3b, 0xcd, 0xaf, 0x60, 0x16, 0xd9, 0x1d, 0xf2, 0x3a, 0x03, 0x57, 0xd2, 0xe2, 0x2b, 0x54, 0x5c,
	0x3b, 0xb8, 0x81, 0x53, 0x1e, 0xcc, 0xf8, 0x6e, 0x10, 0xe6, 0x5e, 0x10, 0xcd, 0x5f, 0x06, 0x58,
	0x45, 0xa6, 0xff, 0xef, 0xfd, 0x0e, 0x4e, 0x22, 0x1a, 0x93, 0x7c, 0xa9, 0xca, 0x31, 0xd7, 0xee,
	0xde, 0xf9, 0xb3, 0x7f, 0x1e, 0x4d, 0x29, 0x39, 0x2f, 0x9a, 0xc1, 0xf1, 0x1d, 0xab, 0xab, 0xdd,
	0x7b, 0x59, 0x7b, 0xf7, 0x6a, 0x7f, 0x04, 0xd8, 0x4e, 0x36, 0x3a, 0x05, 0x34, 0x1c, 0x7d, 0x18,
	0x7c, 0xb9, 0x9c, 0x85, 0x73, 0x7f, 0xea, 0x5f, 0xf8, 0x97, 0xfe, 0xec, 0x5b, 0xe3, 0x11, 0xf2,
	0xa0, 0xae, 0xeb, 0xcb, 0x51, 0xc3, 0x40, 0x0e, 0x58, 0xd3, 0xcf, 0xfe, 0xa4, 0x51, 0x43, 0x00,
	0xf6, 0x27, 0x7f, 0x38, 0x1c, 0x5d, 0x35, 0xcc, 0x76, 0x17, 0xea, 0xe5, 0xe8, 0x64, 0xa8, 0x0d,
	0x35, 0x2e, 0xb0, 0xa1, 0x67, 0x0b, 0xed, 0xcf, 0x56, 0x50, 0xe3, 0xe2, 0xa2, 0x0b, 0x98, 0xcb,
	0xa4, 0xda, 0xdc, 0xac, 0xf4, 0xc5, 0x71, 0xc9, 0x4d, 0x8a, 0x75, 0xce, 0x26, 0xc6, 0x1f, 0xc3,
	0xb8, 0xb6, 0xf5, 0x6e, 0xbf, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x11, 0x5d, 0x40, 0x58, 0x3e,
	0x04, 0x00, 0x00,
}

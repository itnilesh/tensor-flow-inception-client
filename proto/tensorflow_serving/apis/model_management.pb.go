// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow_serving/apis/model_management.proto

package tensorflow_serving

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import config "github.com/itnilesh/tensor-flow-inception-client/proto/tensorflow_serving/config"
import util "github.com/itnilesh/tensor-flow-inception-client/proto/tensorflow_serving/util"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ReloadConfigRequest struct {
	Config               *config.ModelServerConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ReloadConfigRequest) Reset()         { *m = ReloadConfigRequest{} }
func (m *ReloadConfigRequest) String() string { return proto.CompactTextString(m) }
func (*ReloadConfigRequest) ProtoMessage()    {}
func (*ReloadConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_management_e6371a463fdd84cf, []int{0}
}
func (m *ReloadConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReloadConfigRequest.Unmarshal(m, b)
}
func (m *ReloadConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReloadConfigRequest.Marshal(b, m, deterministic)
}
func (dst *ReloadConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReloadConfigRequest.Merge(dst, src)
}
func (m *ReloadConfigRequest) XXX_Size() int {
	return xxx_messageInfo_ReloadConfigRequest.Size(m)
}
func (m *ReloadConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReloadConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReloadConfigRequest proto.InternalMessageInfo

func (m *ReloadConfigRequest) GetConfig() *config.ModelServerConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type ReloadConfigResponse struct {
	Status               *util.StatusProto `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ReloadConfigResponse) Reset()         { *m = ReloadConfigResponse{} }
func (m *ReloadConfigResponse) String() string { return proto.CompactTextString(m) }
func (*ReloadConfigResponse) ProtoMessage()    {}
func (*ReloadConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_management_e6371a463fdd84cf, []int{1}
}
func (m *ReloadConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReloadConfigResponse.Unmarshal(m, b)
}
func (m *ReloadConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReloadConfigResponse.Marshal(b, m, deterministic)
}
func (dst *ReloadConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReloadConfigResponse.Merge(dst, src)
}
func (m *ReloadConfigResponse) XXX_Size() int {
	return xxx_messageInfo_ReloadConfigResponse.Size(m)
}
func (m *ReloadConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReloadConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReloadConfigResponse proto.InternalMessageInfo

func (m *ReloadConfigResponse) GetStatus() *util.StatusProto {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*ReloadConfigRequest)(nil), "tensorflow.serving.ReloadConfigRequest")
	proto.RegisterType((*ReloadConfigResponse)(nil), "tensorflow.serving.ReloadConfigResponse")
}

func init() {
	proto.RegisterFile("tensorflow_serving/apis/model_management.proto", fileDescriptor_model_management_e6371a463fdd84cf)
}

var fileDescriptor_model_management_e6371a463fdd84cf = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8e, 0x31, 0x6b, 0xc3, 0x30,
	0x14, 0x84, 0x11, 0x05, 0x0f, 0xea, 0xa6, 0x76, 0x28, 0x5e, 0x5a, 0x4c, 0x0b, 0x9d, 0x64, 0xa8,
	0x87, 0x4e, 0x5d, 0xda, 0xb9, 0xb4, 0xd8, 0xd9, 0x8d, 0x12, 0x3f, 0x1b, 0x81, 0xac, 0xe7, 0xe8,
	0xc9, 0xc9, 0x5f, 0xcf, 0x18, 0x2c, 0x09, 0x42, 0x88, 0xd7, 0xe3, 0xbb, 0xef, 0x8e, 0x4b, 0x0f,
	0x96, 0xd0, 0xf5, 0x06, 0x8f, 0x2d, 0x81, 0x3b, 0x68, 0x3b, 0x94, 0x6a, 0xd2, 0x54, 0x8e, 0xd8,
	0x81, 0x69, 0x47, 0x65, 0xd5, 0x00, 0x23, 0x58, 0x2f, 0x27, 0x87, 0x1e, 0x85, 0xb8, 0xf0, 0x32,
	0xf1, 0x79, 0xb5, 0xe2, 0xd8, 0xa1, 0xed, 0xf5, 0x90, 0x2c, 0x4b, 0x08, 0xae, 0x8d, 0x59, 0x14,
	0xe5, 0xaf, 0x2b, 0xa5, 0xd9, 0x6b, 0x53, 0x92, 0x57, 0x7e, 0xa6, 0x48, 0x15, 0x1b, 0xfe, 0x50,
	0x83, 0x41, 0xd5, 0xfd, 0x84, 0x6e, 0x0d, 0xfb, 0x19, 0xc8, 0x8b, 0x2f, 0x9e, 0x45, 0xd9, 0x13,
	0x7b, 0x61, 0xef, 0xf7, 0x1f, 0x6f, 0xf2, 0xf6, 0x96, 0xfc, 0x5d, 0xb6, 0x9b, 0x30, 0x9d, 0xda,
	0xa9, 0x54, 0xfc, 0xf1, 0xc7, 0x6b, 0x2b, 0x4d, 0x68, 0x09, 0xc4, 0x27, 0xcf, 0xe2, 0x7a, 0xd2,
	0x3e, 0xaf, 0x69, 0x9b, 0x40, 0xfc, 0x2f, 0xf7, 0xea, 0x84, 0x7f, 0xdf, 0x9d, 0x18, 0xdb, 0x66,
	0xe1, 0x72, 0x75, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x23, 0xf4, 0x92, 0xdb, 0x53, 0x01, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/cluster/v3/filter.proto

package envoy_config_cluster_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type Filter struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TypedConfig          *any.Any `protobuf:"bytes,2,opt,name=typed_config,json=typedConfig,proto3" json:"typed_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b51d51fa3c53520, []int{0}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Filter) GetTypedConfig() *any.Any {
	if m != nil {
		return m.TypedConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*Filter)(nil), "envoy.config.cluster.v3.Filter")
}

func init() {
	proto.RegisterFile("envoy/config/cluster/v3/filter.proto", fileDescriptor_9b51d51fa3c53520)
}

var fileDescriptor_9b51d51fa3c53520 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd0, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x06, 0x60, 0x52, 0x4a, 0xc5, 0xac, 0xa7, 0x45, 0x68, 0x6d, 0x51, 0xd6, 0xa2, 0xd0, 0xd3,
	0x04, 0xba, 0xa0, 0xe0, 0xcd, 0x15, 0x3c, 0x97, 0xbe, 0x80, 0xa4, 0xdd, 0xec, 0x12, 0x58, 0x67,
	0x42, 0x36, 0x1b, 0xcc, 0xd5, 0x93, 0xcf, 0xe0, 0xfb, 0xf8, 0x52, 0x9e, 0xc4, 0xa4, 0xf5, 0xe6,
	0x2d, 0x24, 0xdf, 0xf0, 0xff, 0x13, 0x7e, 0xa3, 0xd0, 0x53, 0x10, 0x7b, 0xc2, 0x46, 0xb7, 0x62,
	0xdf, 0x0d, 0xbd, 0x53, 0x56, 0xf8, 0x52, 0x34, 0xba, 0x73, 0xca, 0x82, 0xb1, 0xe4, 0x28, 0x9f,
	0x46, 0x05, 0x49, 0xc1, 0x41, 0x81, 0x2f, 0xe7, 0x17, 0x2d, 0x51, 0xdb, 0x29, 0x11, 0xd9, 0x6e,
	0x68, 0x84, 0xc4, 0x90, 0x66, 0xe6, 0xd7, 0x43, 0x6d, 0xa4, 0x90, 0x88, 0xe4, 0xa4, 0xd3, 0x84,
	0xbd, 0xf0, 0xca, 0xf6, 0x9a, 0x50, 0x63, 0x7b, 0x20, 0x53, 0x2f, 0x3b, 0x5d, 0x4b, 0xa7, 0xc4,
	0xf1, 0x90, 0x1e, 0x96, 0xef, 0x8c, 0x4f, 0x9e, 0x63, 0x81, 0x7c, 0xc1, 0xc7, 0x28, 0x5f, 0xd5,
	0x8c, 0x15, 0x6c, 0x75, 0x5a, 0x9d, 0x7c, 0x57, 0x63, 0x3b, 0x2a, 0xd8, 0x36, 0x5e, 0xe6, 0xf7,
	0xfc, 0xcc, 0x05, 0xa3, 0xea, 0x97, 0xd4, 0x6c, 0x36, 0x2a, 0xd8, 0x2a, 0x5b, 0x9f, 0x43, 0x6a,
	0x05, 0xc7, 0x56, 0xf0, 0x88, 0x61, 0x9b, 0x45, 0xf9, 0x14, 0xe1, 0xc3, 0xf2, 0xf3, 0xeb, 0xe3,
	0xea, 0x92, 0x2f, 0xd2, 0x5e, 0xd2, 0x68, 0xf0, 0xeb, 0xbf, 0xbd, 0x52, 0x72, 0x75, 0xc7, 0x6f,
	0x35, 0x41, 0x14, 0xc6, 0xd2, 0x5b, 0x80, 0x7f, 0x3e, 0xa1, 0xca, 0xd2, 0xc0, 0xe6, 0x37, 0x6d,
	0xc3, 0x76, 0x93, 0x18, 0x5b, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x22, 0x39, 0xfc, 0x42, 0x5b,
	0x01, 0x00, 0x00,
}

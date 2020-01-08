// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/listener/original_dst/v3alpha/original_dst.proto

package envoy_extensions_filters_listener_original_dst_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
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

type OriginalDst struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OriginalDst) Reset()         { *m = OriginalDst{} }
func (m *OriginalDst) String() string { return proto.CompactTextString(m) }
func (*OriginalDst) ProtoMessage()    {}
func (*OriginalDst) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed68029c7ea3e420, []int{0}
}

func (m *OriginalDst) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OriginalDst.Unmarshal(m, b)
}
func (m *OriginalDst) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OriginalDst.Marshal(b, m, deterministic)
}
func (m *OriginalDst) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OriginalDst.Merge(m, src)
}
func (m *OriginalDst) XXX_Size() int {
	return xxx_messageInfo_OriginalDst.Size(m)
}
func (m *OriginalDst) XXX_DiscardUnknown() {
	xxx_messageInfo_OriginalDst.DiscardUnknown(m)
}

var xxx_messageInfo_OriginalDst proto.InternalMessageInfo

func init() {
	proto.RegisterType((*OriginalDst)(nil), "envoy.extensions.filters.listener.original_dst.v3alpha.OriginalDst")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/listener/original_dst/v3alpha/original_dst.proto", fileDescriptor_ed68029c7ea3e420)
}

var fileDescriptor_ed68029c7ea3e420 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0x4c, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xad, 0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x4f, 0xcb, 0xcc,
	0x29, 0x49, 0x2d, 0x2a, 0xd6, 0xcf, 0xc9, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x2d, 0xd2, 0xcf, 0x2f,
	0xca, 0x4c, 0xcf, 0xcc, 0x4b, 0xcc, 0x89, 0x4f, 0x29, 0x2e, 0xd1, 0x2f, 0x33, 0x4e, 0xcc, 0x29,
	0xc8, 0x48, 0x44, 0x11, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x32, 0x03, 0x1b, 0xa5, 0x87,
	0x30, 0x4a, 0x0f, 0x6a, 0x94, 0x1e, 0xcc, 0x28, 0x3d, 0x14, 0x5d, 0x50, 0xa3, 0xa4, 0x14, 0x4b,
	0x53, 0x0a, 0x12, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0xc0, 0x4e, 0x28, 0x4b, 0x2d,
	0x02, 0x19, 0x90, 0x99, 0x97, 0x0e, 0x31, 0x5a, 0xc9, 0x8f, 0x8b, 0xdb, 0x1f, 0xaa, 0xd5, 0xa5,
	0xb8, 0xc4, 0xca, 0x7e, 0xd6, 0xd1, 0x0e, 0x39, 0x2b, 0x2e, 0x0b, 0x88, 0x85, 0xc9, 0xf9, 0x79,
	0x69, 0x99, 0xe9, 0x50, 0xcb, 0x70, 0xd9, 0x65, 0xa4, 0x87, 0x64, 0x80, 0x53, 0x14, 0x97, 0x4b,
	0x66, 0xbe, 0x1e, 0x58, 0x7b, 0x41, 0x51, 0x7e, 0x45, 0xa5, 0x1e, 0x79, 0x4e, 0x77, 0x12, 0x40,
	0x32, 0x34, 0x00, 0xe4, 0xd2, 0x00, 0xc6, 0x24, 0x36, 0xb0, 0x93, 0x8d, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xad, 0x64, 0xaa, 0x65, 0x5a, 0x01, 0x00, 0x00,
}
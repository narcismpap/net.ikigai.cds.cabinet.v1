// Code generated by protoc-gen-go. DO NOT EDIT.
// source: meta.proto

package cds_cabinet_v1

import (
	fmt "fmt"
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

type Meta struct {
	// Types that are valid to be assigned to Object:
	//	*Meta_Edge
	//	*Meta_Node
	Object               isMeta_Object `protobuf_oneof:"object"`
	Key                  uint32        `protobuf:"varint,10,opt,name=key,proto3" json:"key,omitempty"`
	Val                  []byte        `protobuf:"bytes,11,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Meta) Reset()         { *m = Meta{} }
func (m *Meta) String() string { return proto.CompactTextString(m) }
func (*Meta) ProtoMessage()    {}
func (*Meta) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b5ea8fe65782bcc, []int{0}
}

func (m *Meta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Meta.Unmarshal(m, b)
}
func (m *Meta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Meta.Marshal(b, m, deterministic)
}
func (m *Meta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Meta.Merge(m, src)
}
func (m *Meta) XXX_Size() int {
	return xxx_messageInfo_Meta.Size(m)
}
func (m *Meta) XXX_DiscardUnknown() {
	xxx_messageInfo_Meta.DiscardUnknown(m)
}

var xxx_messageInfo_Meta proto.InternalMessageInfo

type isMeta_Object interface {
	isMeta_Object()
}

type Meta_Edge struct {
	Edge *Edge `protobuf:"bytes,1,opt,name=edge,proto3,oneof"`
}

type Meta_Node struct {
	Node string `protobuf:"bytes,2,opt,name=node,proto3,oneof"`
}

func (*Meta_Edge) isMeta_Object() {}

func (*Meta_Node) isMeta_Object() {}

func (m *Meta) GetObject() isMeta_Object {
	if m != nil {
		return m.Object
	}
	return nil
}

func (m *Meta) GetEdge() *Edge {
	if x, ok := m.GetObject().(*Meta_Edge); ok {
		return x.Edge
	}
	return nil
}

func (m *Meta) GetNode() string {
	if x, ok := m.GetObject().(*Meta_Node); ok {
		return x.Node
	}
	return ""
}

func (m *Meta) GetKey() uint32 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *Meta) GetVal() []byte {
	if m != nil {
		return m.Val
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Meta) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Meta_Edge)(nil),
		(*Meta_Node)(nil),
	}
}

func init() {
	proto.RegisterType((*Meta)(nil), "cds.cabinet.v1.Meta")
}

func init() { proto.RegisterFile("meta.proto", fileDescriptor_3b5ea8fe65782bcc) }

var fileDescriptor_3b5ea8fe65782bcc = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0x4d, 0x2d, 0x49,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4b, 0x4e, 0x29, 0xd6, 0x4b, 0x4e, 0x4c, 0xca,
	0xcc, 0x4b, 0x2d, 0xd1, 0x2b, 0x33, 0x94, 0xe2, 0x4a, 0x4d, 0x49, 0x4f, 0x85, 0xc8, 0x29, 0x95,
	0x71, 0xb1, 0xf8, 0xa6, 0x96, 0x24, 0x0a, 0x69, 0x71, 0xb1, 0x80, 0x44, 0x25, 0x18, 0x15, 0x18,
	0x35, 0xb8, 0x8d, 0x44, 0xf4, 0x50, 0xb5, 0xe8, 0xb9, 0xa6, 0xa4, 0xa7, 0x7a, 0x30, 0x04, 0x81,
	0xd5, 0x08, 0x89, 0x70, 0xb1, 0xe4, 0xe5, 0xa7, 0xa4, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x82,
	0x44, 0x41, 0x3c, 0x21, 0x01, 0x2e, 0xe6, 0xec, 0xd4, 0x4a, 0x09, 0x2e, 0x05, 0x46, 0x0d, 0xde,
	0x20, 0x10, 0x13, 0x24, 0x52, 0x96, 0x98, 0x23, 0xc1, 0xad, 0xc0, 0xa8, 0xc1, 0x13, 0x04, 0x62,
	0x3a, 0x71, 0x70, 0xb1, 0xe5, 0x27, 0x65, 0xa5, 0x26, 0x97, 0x24, 0xb1, 0x81, 0xad, 0x37, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x5d, 0xea, 0xda, 0xa8, 0x00, 0x00, 0x00,
}

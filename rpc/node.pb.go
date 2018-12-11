// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node.proto

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

type Node struct {
	Type                 uint32            `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Id                   string            `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Version              uint32            `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	Created              uint32            `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              uint32            `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Prop                 map[uint32][]byte `protobuf:"bytes,6,rep,name=prop,proto3" json:"prop,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{0}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Node) GetCreated() uint32 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Node) GetUpdated() uint32 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *Node) GetProp() map[uint32][]byte {
	if m != nil {
		return m.Prop
	}
	return nil
}

func init() {
	proto.RegisterType((*Node)(nil), "cds.cabinet.v1.Node")
	proto.RegisterMapType((map[uint32][]byte)(nil), "cds.cabinet.v1.Node.PropEntry")
}

func init() { proto.RegisterFile("node.proto", fileDescriptor_0c843d59d2d938e7) }

var fileDescriptor_0c843d59d2d938e7 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xb1, 0x4a, 0xc6, 0x30,
	0x14, 0x46, 0x49, 0x9a, 0xff, 0x97, 0x5e, 0xb5, 0xc8, 0xc5, 0x21, 0x38, 0x48, 0x71, 0xea, 0x14,
	0xb0, 0x0e, 0x8a, 0xbb, 0xab, 0x48, 0xde, 0xa0, 0x6d, 0xee, 0x10, 0x94, 0x24, 0xa4, 0x69, 0xa1,
	0x0f, 0xec, 0x7b, 0x48, 0xd2, 0x2a, 0xb8, 0xdd, 0x93, 0xf3, 0x0d, 0x27, 0x00, 0xce, 0x1b, 0x52,
	0x21, 0xfa, 0xe4, 0xb1, 0x99, 0xcc, 0xac, 0xa6, 0x61, 0xb4, 0x8e, 0x92, 0x5a, 0x1f, 0x1f, 0xbe,
	0x19, 0x88, 0x77, 0x6f, 0x08, 0x11, 0x44, 0xda, 0x02, 0x49, 0xd6, 0xb2, 0xee, 0x5a, 0x97, 0x1b,
	0x1b, 0xe0, 0xd6, 0x48, 0xde, 0xb2, 0xae, 0xd6, 0xdc, 0x1a, 0x94, 0x70, 0xb1, 0x52, 0x9c, 0xad,
	0x77, 0xb2, 0x2a, 0xb3, 0x5f, 0xcc, 0x66, 0x8a, 0x34, 0x24, 0x32, 0x52, 0xec, 0xe6, 0xc0, 0x6c,
	0x96, 0x60, 0x8a, 0x39, 0xed, 0xe6, 0x40, 0xec, 0x41, 0x84, 0xe8, 0x83, 0x3c, 0xb7, 0x55, 0x77,
	0xd9, 0xdf, 0xab, 0xff, 0x65, 0x2a, 0x57, 0xa9, 0x8f, 0xe8, 0xc3, 0x9b, 0x4b, 0x71, 0xd3, 0x65,
	0x7b, 0xf7, 0x0c, 0xf5, 0xdf, 0x13, 0xde, 0x40, 0xf5, 0x49, 0xdb, 0x51, 0x9c, 0x4f, 0xbc, 0x85,
	0xd3, 0x3a, 0x7c, 0x2d, 0x54, 0x9a, 0xaf, 0xf4, 0x0e, 0xaf, 0xfc, 0x85, 0x8d, 0xe7, 0xf2, 0xfd,
	0xa7, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x39, 0xb0, 0x20, 0x0c, 0x01, 0x00, 0x00,
}
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
	Type                 uint32   `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Version              uint32   `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	Created              uint32   `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              uint32   `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Properties           []byte   `protobuf:"bytes,10,opt,name=properties,proto3" json:"properties,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
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

func (m *Node) GetProperties() []byte {
	if m != nil {
		return m.Properties
	}
	return nil
}

func init() {
	proto.RegisterType((*Node)(nil), "cds.cabinet.v1.Node")
}

func init() { proto.RegisterFile("node.proto", fileDescriptor_0c843d59d2d938e7) }

var fileDescriptor_0c843d59d2d938e7 = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xce, 0x31, 0x0e, 0xc2, 0x30,
	0x0c, 0x40, 0x51, 0xa5, 0x14, 0x10, 0x16, 0x74, 0xf0, 0xe4, 0x09, 0x55, 0x4c, 0x9d, 0x2a, 0x21,
	0xee, 0xc1, 0xd0, 0x1b, 0xb4, 0xb1, 0x87, 0x2c, 0x71, 0x94, 0x84, 0x4a, 0xdc, 0x84, 0xe3, 0xa2,
	0xa6, 0x45, 0x62, 0xf3, 0xf7, 0xf3, 0x60, 0x00, 0xaf, 0x2c, 0x7d, 0x88, 0x9a, 0x15, 0x1b, 0xcb,
	0xa9, 0xb7, 0xe3, 0xe4, 0xbc, 0xe4, 0x7e, 0xbe, 0xdf, 0x3e, 0x06, 0xea, 0xa7, 0xb2, 0x20, 0x42,
	0x9d, 0xdf, 0x41, 0xc8, 0xb4, 0xa6, 0xbb, 0x0c, 0x65, 0xc6, 0x06, 0x2a, 0xc7, 0x54, 0xb5, 0xa6,
	0x3b, 0x0d, 0x95, 0x63, 0x24, 0x38, 0xce, 0x12, 0x93, 0x53, 0x4f, 0xbb, 0x72, 0xf6, 0xcb, 0x45,
	0x6c, 0x94, 0x31, 0x0b, 0x53, 0xbd, 0xca, 0x96, 0x8b, 0xbc, 0x02, 0x17, 0xd9, 0xaf, 0xb2, 0x25,
	0x5e, 0x01, 0x42, 0xd4, 0x20, 0x31, 0x3b, 0x49, 0x04, 0xad, 0xe9, 0xce, 0xc3, 0xdf, 0x66, 0x3a,
	0x94, 0x8f, 0x1f, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa0, 0xb1, 0x48, 0xfa, 0xbf, 0x00, 0x00,
	0x00,
}

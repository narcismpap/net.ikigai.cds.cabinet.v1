// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service_counter.proto

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

type CounterValueResponse struct {
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CounterValueResponse) Reset()         { *m = CounterValueResponse{} }
func (m *CounterValueResponse) String() string { return proto.CompactTextString(m) }
func (*CounterValueResponse) ProtoMessage()    {}
func (*CounterValueResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd34d23255192a0d, []int{0}
}

func (m *CounterValueResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CounterValueResponse.Unmarshal(m, b)
}
func (m *CounterValueResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CounterValueResponse.Marshal(b, m, deterministic)
}
func (m *CounterValueResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CounterValueResponse.Merge(m, src)
}
func (m *CounterValueResponse) XXX_Size() int {
	return xxx_messageInfo_CounterValueResponse.Size(m)
}
func (m *CounterValueResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CounterValueResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CounterValueResponse proto.InternalMessageInfo

func (m *CounterValueResponse) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*CounterValueResponse)(nil), "cds.cabinet.v1.CounterValueResponse")
}

func init() { proto.RegisterFile("service_counter.proto", fileDescriptor_fd34d23255192a0d) }

var fileDescriptor_fd34d23255192a0d = []byte{
	// 114 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x8d, 0x4f, 0xce, 0x2f, 0xcd, 0x2b, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x4b, 0x4e, 0x29, 0xd6, 0x4b, 0x4e, 0x4c, 0xca, 0xcc, 0x4b, 0x2d, 0xd1, 0x2b,
	0x33, 0x94, 0xe2, 0x29, 0xce, 0x48, 0x2c, 0x4a, 0x4d, 0x81, 0xc8, 0x2a, 0xe9, 0x70, 0x89, 0x38,
	0x43, 0x94, 0x87, 0x25, 0xe6, 0x94, 0xa6, 0x06, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a,
	0x89, 0x70, 0xb1, 0x96, 0x81, 0x04, 0x24, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0x20, 0x9c, 0x24,
	0x36, 0xb0, 0x26, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x17, 0xec, 0x26, 0x0b, 0x6b, 0x00,
	0x00, 0x00,
}

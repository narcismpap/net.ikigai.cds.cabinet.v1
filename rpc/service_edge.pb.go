// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service_edge.proto

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

type EdgeGetRequest struct {
	Mode                 RetrieveMode `protobuf:"varint,1,opt,name=mode,proto3,enum=cds.cabinet.v1.RetrieveMode" json:"mode,omitempty"`
	Edge                 *Edge        `protobuf:"bytes,2,opt,name=edge,proto3" json:"edge,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *EdgeGetRequest) Reset()         { *m = EdgeGetRequest{} }
func (m *EdgeGetRequest) String() string { return proto.CompactTextString(m) }
func (*EdgeGetRequest) ProtoMessage()    {}
func (*EdgeGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ec6de44fafa439a, []int{0}
}

func (m *EdgeGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EdgeGetRequest.Unmarshal(m, b)
}
func (m *EdgeGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EdgeGetRequest.Marshal(b, m, deterministic)
}
func (m *EdgeGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EdgeGetRequest.Merge(m, src)
}
func (m *EdgeGetRequest) XXX_Size() int {
	return xxx_messageInfo_EdgeGetRequest.Size(m)
}
func (m *EdgeGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EdgeGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EdgeGetRequest proto.InternalMessageInfo

func (m *EdgeGetRequest) GetMode() RetrieveMode {
	if m != nil {
		return m.Mode
	}
	return RetrieveMode_ALL
}

func (m *EdgeGetRequest) GetEdge() *Edge {
	if m != nil {
		return m.Edge
	}
	return nil
}

type EdgeListRequest struct {
	Opt                  *ListOptions `protobuf:"bytes,1,opt,name=opt,proto3" json:"opt,omitempty"`
	Subject              string       `protobuf:"bytes,10,opt,name=subject,proto3" json:"subject,omitempty"`
	Predicate            uint32       `protobuf:"varint,11,opt,name=predicate,proto3" json:"predicate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *EdgeListRequest) Reset()         { *m = EdgeListRequest{} }
func (m *EdgeListRequest) String() string { return proto.CompactTextString(m) }
func (*EdgeListRequest) ProtoMessage()    {}
func (*EdgeListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ec6de44fafa439a, []int{1}
}

func (m *EdgeListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EdgeListRequest.Unmarshal(m, b)
}
func (m *EdgeListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EdgeListRequest.Marshal(b, m, deterministic)
}
func (m *EdgeListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EdgeListRequest.Merge(m, src)
}
func (m *EdgeListRequest) XXX_Size() int {
	return xxx_messageInfo_EdgeListRequest.Size(m)
}
func (m *EdgeListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EdgeListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EdgeListRequest proto.InternalMessageInfo

func (m *EdgeListRequest) GetOpt() *ListOptions {
	if m != nil {
		return m.Opt
	}
	return nil
}

func (m *EdgeListRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *EdgeListRequest) GetPredicate() uint32 {
	if m != nil {
		return m.Predicate
	}
	return 0
}

func init() {
	proto.RegisterType((*EdgeGetRequest)(nil), "cds.cabinet.v1.EdgeGetRequest")
	proto.RegisterType((*EdgeListRequest)(nil), "cds.cabinet.v1.EdgeListRequest")
}

func init() { proto.RegisterFile("service_edge.proto", fileDescriptor_5ec6de44fafa439a) }

var fileDescriptor_5ec6de44fafa439a = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xc1, 0x4a, 0xc4, 0x30,
	0x14, 0x45, 0x89, 0x0e, 0xca, 0xbc, 0x6a, 0x85, 0xe0, 0x22, 0x8c, 0xb3, 0x28, 0xb3, 0xca, 0xc6,
	0xa0, 0xf5, 0x1b, 0xc4, 0x8d, 0x22, 0xe4, 0x07, 0xa4, 0x4d, 0x2e, 0x63, 0x44, 0x27, 0x35, 0x79,
	0x53, 0xfc, 0x7c, 0x49, 0x45, 0x07, 0xbb, 0xbc, 0xc9, 0xb9, 0xe7, 0xf2, 0x48, 0x66, 0xa4, 0x31,
	0x38, 0xbc, 0xc0, 0x6f, 0x61, 0x86, 0x14, 0x39, 0xca, 0xda, 0xf9, 0x6c, 0x5c, 0xd7, 0x87, 0x1d,
	0xd8, 0x8c, 0xb7, 0xab, 0xb3, 0xfc, 0xda, 0x25, 0xf8, 0x9f, 0xdf, 0x15, 0x1d, 0xc8, 0xcd, 0x3b,
	0xd5, 0xf7, 0x7e, 0x8b, 0x07, 0xb0, 0xc5, 0xe7, 0x1e, 0x99, 0xe5, 0x0d, 0x2d, 0x3e, 0xa2, 0x87,
	0x12, 0x8d, 0xd0, 0x75, 0xbb, 0x36, 0xff, 0x55, 0xc6, 0x82, 0x53, 0xc0, 0x88, 0xa7, 0xe8, 0x61,
	0x27, 0x52, 0x6a, 0x5a, 0x14, 0xa3, 0x3a, 0x6a, 0x84, 0xae, 0xda, 0xcb, 0x79, 0xa3, 0xf8, 0xed,
	0x44, 0x6c, 0xbe, 0xe8, 0xa2, 0xa4, 0xc7, 0x90, 0xff, 0xe6, 0xae, 0xe9, 0x38, 0x0e, 0x3c, 0xad,
	0x55, 0xed, 0xd5, 0xbc, 0x5b, 0xc8, 0xe7, 0x81, 0x43, 0xdc, 0x65, 0x5b, 0x38, 0xa9, 0xe8, 0x34,
	0xef, 0xfb, 0x37, 0x38, 0x56, 0xd4, 0x08, 0xbd, 0xb4, 0xbf, 0x51, 0xae, 0x69, 0x39, 0x24, 0xf8,
	0xe0, 0x3a, 0x86, 0xaa, 0x1a, 0xa1, 0xcf, 0xed, 0xe1, 0xa1, 0x3f, 0x99, 0xce, 0xbd, 0xfb, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0xcc, 0xcc, 0x34, 0xa2, 0x2e, 0x01, 0x00, 0x00,
}
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service_index.proto

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

type IndexGetRequest struct {
	Index                *Index   `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IndexGetRequest) Reset()         { *m = IndexGetRequest{} }
func (m *IndexGetRequest) String() string { return proto.CompactTextString(m) }
func (*IndexGetRequest) ProtoMessage()    {}
func (*IndexGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e40c52b04100e5, []int{0}
}

func (m *IndexGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexGetRequest.Unmarshal(m, b)
}
func (m *IndexGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexGetRequest.Marshal(b, m, deterministic)
}
func (m *IndexGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexGetRequest.Merge(m, src)
}
func (m *IndexGetRequest) XXX_Size() int {
	return xxx_messageInfo_IndexGetRequest.Size(m)
}
func (m *IndexGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IndexGetRequest proto.InternalMessageInfo

func (m *IndexGetRequest) GetIndex() *Index {
	if m != nil {
		return m.Index
	}
	return nil
}

type IndexDropRequest struct {
	Index                uint32   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IndexDropRequest) Reset()         { *m = IndexDropRequest{} }
func (m *IndexDropRequest) String() string { return proto.CompactTextString(m) }
func (*IndexDropRequest) ProtoMessage()    {}
func (*IndexDropRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e40c52b04100e5, []int{1}
}

func (m *IndexDropRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexDropRequest.Unmarshal(m, b)
}
func (m *IndexDropRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexDropRequest.Marshal(b, m, deterministic)
}
func (m *IndexDropRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexDropRequest.Merge(m, src)
}
func (m *IndexDropRequest) XXX_Size() int {
	return xxx_messageInfo_IndexDropRequest.Size(m)
}
func (m *IndexDropRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexDropRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IndexDropRequest proto.InternalMessageInfo

func (m *IndexDropRequest) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

type IndexListRequest struct {
	Opt                  *ListOptions `protobuf:"bytes,1,opt,name=opt,proto3" json:"opt,omitempty"`
	Index                uint32       `protobuf:"varint,10,opt,name=index,proto3" json:"index,omitempty"`
	Value                string       `protobuf:"bytes,11,opt,name=value,proto3" json:"value,omitempty"`
	IncludeIndex         bool         `protobuf:"varint,20,opt,name=include_index,json=includeIndex,proto3" json:"include_index,omitempty"`
	IncludeValue         bool         `protobuf:"varint,21,opt,name=include_value,json=includeValue,proto3" json:"include_value,omitempty"`
	IncludeProp          bool         `protobuf:"varint,22,opt,name=include_prop,json=includeProp,proto3" json:"include_prop,omitempty"`
	IncludeNode          bool         `protobuf:"varint,23,opt,name=include_node,json=includeNode,proto3" json:"include_node,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *IndexListRequest) Reset()         { *m = IndexListRequest{} }
func (m *IndexListRequest) String() string { return proto.CompactTextString(m) }
func (*IndexListRequest) ProtoMessage()    {}
func (*IndexListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e40c52b04100e5, []int{2}
}

func (m *IndexListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexListRequest.Unmarshal(m, b)
}
func (m *IndexListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexListRequest.Marshal(b, m, deterministic)
}
func (m *IndexListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexListRequest.Merge(m, src)
}
func (m *IndexListRequest) XXX_Size() int {
	return xxx_messageInfo_IndexListRequest.Size(m)
}
func (m *IndexListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IndexListRequest proto.InternalMessageInfo

func (m *IndexListRequest) GetOpt() *ListOptions {
	if m != nil {
		return m.Opt
	}
	return nil
}

func (m *IndexListRequest) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *IndexListRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *IndexListRequest) GetIncludeIndex() bool {
	if m != nil {
		return m.IncludeIndex
	}
	return false
}

func (m *IndexListRequest) GetIncludeValue() bool {
	if m != nil {
		return m.IncludeValue
	}
	return false
}

func (m *IndexListRequest) GetIncludeProp() bool {
	if m != nil {
		return m.IncludeProp
	}
	return false
}

func (m *IndexListRequest) GetIncludeNode() bool {
	if m != nil {
		return m.IncludeNode
	}
	return false
}

type IndexChoiceRequest struct {
	Opt                  *ListOptions `protobuf:"bytes,1,opt,name=opt,proto3" json:"opt,omitempty"`
	Index                uint32       `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *IndexChoiceRequest) Reset()         { *m = IndexChoiceRequest{} }
func (m *IndexChoiceRequest) String() string { return proto.CompactTextString(m) }
func (*IndexChoiceRequest) ProtoMessage()    {}
func (*IndexChoiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_46e40c52b04100e5, []int{3}
}

func (m *IndexChoiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexChoiceRequest.Unmarshal(m, b)
}
func (m *IndexChoiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexChoiceRequest.Marshal(b, m, deterministic)
}
func (m *IndexChoiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexChoiceRequest.Merge(m, src)
}
func (m *IndexChoiceRequest) XXX_Size() int {
	return xxx_messageInfo_IndexChoiceRequest.Size(m)
}
func (m *IndexChoiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexChoiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IndexChoiceRequest proto.InternalMessageInfo

func (m *IndexChoiceRequest) GetOpt() *ListOptions {
	if m != nil {
		return m.Opt
	}
	return nil
}

func (m *IndexChoiceRequest) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func init() {
	proto.RegisterType((*IndexGetRequest)(nil), "cds.cabinet.v1.IndexGetRequest")
	proto.RegisterType((*IndexDropRequest)(nil), "cds.cabinet.v1.IndexDropRequest")
	proto.RegisterType((*IndexListRequest)(nil), "cds.cabinet.v1.IndexListRequest")
	proto.RegisterType((*IndexChoiceRequest)(nil), "cds.cabinet.v1.IndexChoiceRequest")
}

func init() { proto.RegisterFile("service_index.proto", fileDescriptor_46e40c52b04100e5) }

var fileDescriptor_46e40c52b04100e5 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0x41, 0x4b, 0xec, 0x30,
	0x14, 0x85, 0xe9, 0x7b, 0x28, 0x9a, 0xce, 0xa8, 0xd4, 0x19, 0x2d, 0xe3, 0xa6, 0xd6, 0x4d, 0x41,
	0x0c, 0xa8, 0x7b, 0x37, 0x0a, 0x22, 0x88, 0x4a, 0x16, 0x82, 0xab, 0xa1, 0x93, 0x5c, 0x98, 0xc0,
	0x90, 0x1b, 0x93, 0xb4, 0xf8, 0xe7, 0x05, 0x69, 0x9a, 0x6a, 0xea, 0xd6, 0xe5, 0xb9, 0xfd, 0xf8,
	0x4e, 0x39, 0x21, 0x87, 0x16, 0x4c, 0x2b, 0x39, 0x2c, 0xa5, 0x12, 0xf0, 0x41, 0xb5, 0x41, 0x87,
	0xd9, 0x1e, 0x17, 0x96, 0xf2, 0x7a, 0x25, 0x15, 0x38, 0xda, 0x5e, 0x2e, 0x26, 0x76, 0x5d, 0x1b,
	0x10, 0xfd, 0xd7, 0x45, 0x1a, 0xa1, 0xe5, 0x0d, 0xd9, 0x7f, 0xe8, 0xe2, 0x3d, 0x38, 0x06, 0xef,
	0x0d, 0x58, 0x97, 0x9d, 0x93, 0x2d, 0x4f, 0xe4, 0x49, 0x91, 0x54, 0xe9, 0xd5, 0x9c, 0x8e, 0x6d,
	0xd4, 0xf3, 0xac, 0x67, 0xca, 0x8a, 0x1c, 0xf8, 0x7c, 0x67, 0x50, 0x0f, 0x82, 0x59, 0x2c, 0x98,
	0x0e, 0xe4, 0x67, 0x12, 0xd0, 0x47, 0x69, 0xbf, 0xbb, 0x2e, 0xc8, 0x7f, 0xd4, 0x2e, 0x34, 0x9d,
	0xfc, 0x6e, 0xea, 0xc8, 0x67, 0xed, 0x24, 0x2a, 0xcb, 0x3a, 0xee, 0xc7, 0x4c, 0x22, 0x73, 0x77,
	0x6d, 0xeb, 0x4d, 0x03, 0x79, 0x5a, 0x24, 0xd5, 0x2e, 0xeb, 0x43, 0x76, 0x46, 0xa6, 0x52, 0xf1,
	0x4d, 0x23, 0xc2, 0x36, 0xf9, 0xac, 0x48, 0xaa, 0x1d, 0x36, 0x09, 0x47, 0xff, 0x2b, 0x31, 0xd4,
	0x2b, 0xe6, 0x23, 0xe8, 0xd5, 0x9b, 0x4e, 0xc9, 0x90, 0x97, 0xda, 0xa0, 0xce, 0x8f, 0x3c, 0x93,
	0x86, 0xdb, 0x8b, 0x41, 0x1d, 0x23, 0x0a, 0x05, 0xe4, 0xc7, 0x23, 0xe4, 0x09, 0x05, 0x94, 0x6f,
	0x24, 0xf3, 0x9d, 0xb7, 0x6b, 0x94, 0x1c, 0xfe, 0x3a, 0xc0, 0xbf, 0x68, 0x80, 0xd5, 0xb6, 0x7f,
	0xcb, 0xeb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8b, 0xfe, 0x3c, 0x22, 0x0d, 0x02, 0x00, 0x00,
}

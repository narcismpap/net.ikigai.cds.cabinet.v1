// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service_node.proto

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

type NodeCreateResponse struct {
	Status               MutationStatus `protobuf:"varint,1,opt,name=status,proto3,enum=cds.cabinet.v1.MutationStatus" json:"status,omitempty"`
	Id                   string         `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *NodeCreateResponse) Reset()         { *m = NodeCreateResponse{} }
func (m *NodeCreateResponse) String() string { return proto.CompactTextString(m) }
func (*NodeCreateResponse) ProtoMessage()    {}
func (*NodeCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_09240cbea398538b, []int{0}
}

func (m *NodeCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeCreateResponse.Unmarshal(m, b)
}
func (m *NodeCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeCreateResponse.Marshal(b, m, deterministic)
}
func (m *NodeCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeCreateResponse.Merge(m, src)
}
func (m *NodeCreateResponse) XXX_Size() int {
	return xxx_messageInfo_NodeCreateResponse.Size(m)
}
func (m *NodeCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NodeCreateResponse proto.InternalMessageInfo

func (m *NodeCreateResponse) GetStatus() MutationStatus {
	if m != nil {
		return m.Status
	}
	return MutationStatus_SUCCESS
}

func (m *NodeCreateResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type NodeGetRequest struct {
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	NodeType             uint32   `protobuf:"varint,3,opt,name=node_type,json=nodeType,proto3" json:"node_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeGetRequest) Reset()         { *m = NodeGetRequest{} }
func (m *NodeGetRequest) String() string { return proto.CompactTextString(m) }
func (*NodeGetRequest) ProtoMessage()    {}
func (*NodeGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_09240cbea398538b, []int{1}
}

func (m *NodeGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeGetRequest.Unmarshal(m, b)
}
func (m *NodeGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeGetRequest.Marshal(b, m, deterministic)
}
func (m *NodeGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeGetRequest.Merge(m, src)
}
func (m *NodeGetRequest) XXX_Size() int {
	return xxx_messageInfo_NodeGetRequest.Size(m)
}
func (m *NodeGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeGetRequest proto.InternalMessageInfo

func (m *NodeGetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NodeGetRequest) GetNodeType() uint32 {
	if m != nil {
		return m.NodeType
	}
	return 0
}

type NodeListRequest struct {
	Opt                  *ListOptions `protobuf:"bytes,1,opt,name=opt,proto3" json:"opt,omitempty"`
	NodeType             uint32       `protobuf:"varint,2,opt,name=node_type,json=nodeType,proto3" json:"node_type,omitempty"`
	IncludeType          bool         `protobuf:"varint,20,opt,name=include_type,json=includeType,proto3" json:"include_type,omitempty"`
	IncludeId            bool         `protobuf:"varint,21,opt,name=include_id,json=includeId,proto3" json:"include_id,omitempty"`
	IncludeProp          bool         `protobuf:"varint,22,opt,name=include_prop,json=includeProp,proto3" json:"include_prop,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *NodeListRequest) Reset()         { *m = NodeListRequest{} }
func (m *NodeListRequest) String() string { return proto.CompactTextString(m) }
func (*NodeListRequest) ProtoMessage()    {}
func (*NodeListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_09240cbea398538b, []int{2}
}

func (m *NodeListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeListRequest.Unmarshal(m, b)
}
func (m *NodeListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeListRequest.Marshal(b, m, deterministic)
}
func (m *NodeListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeListRequest.Merge(m, src)
}
func (m *NodeListRequest) XXX_Size() int {
	return xxx_messageInfo_NodeListRequest.Size(m)
}
func (m *NodeListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeListRequest proto.InternalMessageInfo

func (m *NodeListRequest) GetOpt() *ListOptions {
	if m != nil {
		return m.Opt
	}
	return nil
}

func (m *NodeListRequest) GetNodeType() uint32 {
	if m != nil {
		return m.NodeType
	}
	return 0
}

func (m *NodeListRequest) GetIncludeType() bool {
	if m != nil {
		return m.IncludeType
	}
	return false
}

func (m *NodeListRequest) GetIncludeId() bool {
	if m != nil {
		return m.IncludeId
	}
	return false
}

func (m *NodeListRequest) GetIncludeProp() bool {
	if m != nil {
		return m.IncludeProp
	}
	return false
}

func init() {
	proto.RegisterType((*NodeCreateResponse)(nil), "cds.cabinet.v1.NodeCreateResponse")
	proto.RegisterType((*NodeGetRequest)(nil), "cds.cabinet.v1.NodeGetRequest")
	proto.RegisterType((*NodeListRequest)(nil), "cds.cabinet.v1.NodeListRequest")
}

func init() { proto.RegisterFile("service_node.proto", fileDescriptor_09240cbea398538b) }

var fileDescriptor_09240cbea398538b = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0x47, 0xd9, 0x14, 0x4a, 0x33, 0xad, 0x11, 0x16, 0x95, 0x60, 0x51, 0x62, 0x4e, 0xb9, 0x18,
	0xb0, 0x82, 0x37, 0x4f, 0x1e, 0x44, 0xf0, 0x1f, 0xab, 0x47, 0xa1, 0xa4, 0xd9, 0x01, 0x17, 0x24,
	0xbb, 0xee, 0x4c, 0x0a, 0xfd, 0x7a, 0x7e, 0x32, 0xd9, 0x98, 0x4a, 0xd3, 0xeb, 0x6f, 0xde, 0x3e,
	0x1e, 0x0b, 0x92, 0xd0, 0xaf, 0x4d, 0x8d, 0xcb, 0xc6, 0x6a, 0x2c, 0x9d, 0xb7, 0x6c, 0x65, 0x52,
	0x6b, 0x2a, 0xeb, 0x6a, 0x65, 0x1a, 0xe4, 0x72, 0x7d, 0x75, 0x3a, 0xa3, 0xcf, 0xca, 0xa3, 0xfe,
	0xbb, 0xe6, 0x1f, 0x20, 0x9f, 0xad, 0xc6, 0x3b, 0x8f, 0x15, 0xa3, 0x42, 0x72, 0xb6, 0x21, 0x94,
	0x37, 0x30, 0x26, 0xae, 0xb8, 0xa5, 0x54, 0x64, 0xa2, 0x48, 0x16, 0xe7, 0xe5, 0x50, 0x52, 0x3e,
	0xb5, 0x5c, 0xb1, 0xb1, 0xcd, 0x5b, 0x47, 0xa9, 0x9e, 0x96, 0x09, 0x44, 0x46, 0xa7, 0x51, 0x26,
	0x8a, 0x58, 0x45, 0x46, 0xe7, 0xb7, 0x90, 0x04, 0xfb, 0x3d, 0xb2, 0xc2, 0xef, 0x16, 0x89, 0xf7,
	0x09, 0x39, 0x87, 0x38, 0xb4, 0x2e, 0x79, 0xe3, 0x30, 0x1d, 0x65, 0xa2, 0x38, 0x50, 0x93, 0x30,
	0xbc, 0x6f, 0x1c, 0xe6, 0x3f, 0x02, 0x0e, 0xc3, 0xfb, 0x47, 0x43, 0xff, 0x82, 0x4b, 0x18, 0x59,
	0xc7, 0x5d, 0xd7, 0x74, 0x31, 0xdf, 0xef, 0x0a, 0xe4, 0x8b, 0x0b, 0x65, 0xa4, 0x02, 0x37, 0xf4,
	0x47, 0x43, 0xbf, 0xbc, 0x80, 0x99, 0x69, 0xea, 0xaf, 0x76, 0x7b, 0x3f, 0xca, 0x44, 0x31, 0x51,
	0xd3, 0x7e, 0xeb, 0x90, 0x33, 0x80, 0x2d, 0x62, 0x74, 0x7a, 0xdc, 0x01, 0x71, 0xbf, 0x3c, 0xe8,
	0x5d, 0x83, 0xf3, 0xd6, 0xa5, 0x27, 0x03, 0xc3, 0xab, 0xb7, 0x6e, 0x35, 0xee, 0x3e, 0xfa, 0xfa,
	0x37, 0x00, 0x00, 0xff, 0xff, 0xad, 0xad, 0xc0, 0xc7, 0x9c, 0x01, 0x00, 0x00,
}

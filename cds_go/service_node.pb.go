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
	Mode                 RetrieveMode `protobuf:"varint,1,opt,name=mode,proto3,enum=cds.cabinet.v1.RetrieveMode" json:"mode,omitempty"`
	Id                   string       `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
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

func (m *NodeGetRequest) GetMode() RetrieveMode {
	if m != nil {
		return m.Mode
	}
	return RetrieveMode_ALL
}

func (m *NodeGetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type NodeListRequest struct {
	Opt                  *ListOptions `protobuf:"bytes,1,opt,name=opt,proto3" json:"opt,omitempty"`
	NodeType             string       `protobuf:"bytes,2,opt,name=node_type,json=nodeType,proto3" json:"node_type,omitempty"`
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

func (m *NodeListRequest) GetNodeType() string {
	if m != nil {
		return m.NodeType
	}
	return ""
}

func init() {
	proto.RegisterType((*NodeCreateResponse)(nil), "cds.cabinet.v1.NodeCreateResponse")
	proto.RegisterType((*NodeGetRequest)(nil), "cds.cabinet.v1.NodeGetRequest")
	proto.RegisterType((*NodeListRequest)(nil), "cds.cabinet.v1.NodeListRequest")
}

func init() { proto.RegisterFile("service_node.proto", fileDescriptor_09240cbea398538b) }

var fileDescriptor_09240cbea398538b = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0xd0, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x71, 0x5a, 0x65, 0xb8, 0x28, 0x15, 0x72, 0x1a, 0x4e, 0x64, 0xec, 0xb4, 0x8b, 0x41,
	0x27, 0xf8, 0x0f, 0x78, 0xf0, 0xe2, 0x14, 0xa2, 0x47, 0x65, 0x64, 0xcd, 0x0f, 0xcc, 0xc1, 0xbe,
	0x98, 0xf7, 0x5a, 0xd8, 0x7f, 0x2f, 0xa9, 0x45, 0xb0, 0xbb, 0x26, 0x5f, 0x3e, 0x3c, 0x7e, 0x4a,
	0x33, 0x52, 0x17, 0x6a, 0x6c, 0x1b, 0xf2, 0x30, 0x31, 0x91, 0x90, 0xae, 0x6a, 0xcf, 0xa6, 0x76,
	0xbb, 0xd0, 0x40, 0x4c, 0x77, 0x7b, 0x71, 0xc6, 0x9f, 0x2e, 0xc1, 0xff, 0xfe, 0x2e, 0xdf, 0x95,
	0x7e, 0x26, 0x8f, 0x87, 0x04, 0x27, 0xb0, 0xe0, 0x48, 0x0d, 0x43, 0xdf, 0xab, 0x09, 0x8b, 0x93,
	0x96, 0x67, 0xc5, 0xa2, 0x58, 0x55, 0xeb, 0x2b, 0xf3, 0x1f, 0x31, 0x9b, 0x56, 0x9c, 0x04, 0x6a,
	0x5e, 0xfb, 0xca, 0x0e, 0xb5, 0xae, 0x54, 0x19, 0xfc, 0xac, 0x5c, 0x14, 0xab, 0xa9, 0x2d, 0x83,
	0x5f, 0x5a, 0x55, 0x65, 0xfd, 0x11, 0x62, 0xf1, 0xdd, 0x82, 0x45, 0xdf, 0xa8, 0xe3, 0x2f, 0xf2,
	0x18, 0xdc, 0xcb, 0xb1, 0x6b, 0x21, 0x29, 0xa0, 0xc3, 0x86, 0x3c, 0x6c, 0x5f, 0x1e, 0x98, 0x1f,
	0xea, 0x3c, 0x9b, 0x4f, 0x81, 0xff, 0xd0, 0x6b, 0x75, 0x44, 0x51, 0x7a, 0xf3, 0x74, 0x3d, 0x1f,
	0x9b, 0xb9, 0x7c, 0x89, 0xf9, 0x5a, 0xb6, 0xb9, 0xd3, 0x73, 0x35, 0xcd, 0xfb, 0x6c, 0x65, 0x1f,
	0x31, 0xc0, 0x27, 0xf9, 0xe1, 0x6d, 0x1f, 0xb1, 0x9b, 0xf4, 0xbb, 0xdc, 0xfd, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xe7, 0x67, 0x8b, 0x3e, 0x4b, 0x01, 0x00, 0x00,
}
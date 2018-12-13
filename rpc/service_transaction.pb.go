// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service_transaction.proto

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

type TransactionAction struct {
	// Types that are valid to be assigned to Action:
	//	*TransactionAction_CounterIncrement
	//	*TransactionAction_CounterDelete
	//	*TransactionAction_CounterRegister
	//	*TransactionAction_EdgeUpdate
	//	*TransactionAction_EdgeDelete
	//	*TransactionAction_IndexUpdate
	//	*TransactionAction_IndexDelete
	//	*TransactionAction_MetaUpdate
	//	*TransactionAction_MetaDelete
	//	*TransactionAction_NodeCreate
	//	*TransactionAction_NodeUpdate
	//	*TransactionAction_NodeDelete
	//	*TransactionAction_ReadCheck
	Action               isTransactionAction_Action `protobuf_oneof:"action"`
	ActionId             uint32                     `protobuf:"varint,80,opt,name=action_id,json=actionId,proto3" json:"action_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *TransactionAction) Reset()         { *m = TransactionAction{} }
func (m *TransactionAction) String() string { return proto.CompactTextString(m) }
func (*TransactionAction) ProtoMessage()    {}
func (*TransactionAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_eae9ea52da333b22, []int{0}
}

func (m *TransactionAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionAction.Unmarshal(m, b)
}
func (m *TransactionAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionAction.Marshal(b, m, deterministic)
}
func (m *TransactionAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionAction.Merge(m, src)
}
func (m *TransactionAction) XXX_Size() int {
	return xxx_messageInfo_TransactionAction.Size(m)
}
func (m *TransactionAction) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionAction.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionAction proto.InternalMessageInfo

type isTransactionAction_Action interface {
	isTransactionAction_Action()
}

type TransactionAction_CounterIncrement struct {
	CounterIncrement *Counter `protobuf:"bytes,1,opt,name=counter_increment,json=counterIncrement,proto3,oneof"`
}

type TransactionAction_CounterDelete struct {
	CounterDelete *Counter `protobuf:"bytes,2,opt,name=counter_delete,json=counterDelete,proto3,oneof"`
}

type TransactionAction_CounterRegister struct {
	CounterRegister *Counter `protobuf:"bytes,3,opt,name=counter_register,json=counterRegister,proto3,oneof"`
}

type TransactionAction_EdgeUpdate struct {
	EdgeUpdate *Edge `protobuf:"bytes,11,opt,name=edge_update,json=edgeUpdate,proto3,oneof"`
}

type TransactionAction_EdgeDelete struct {
	EdgeDelete *Edge `protobuf:"bytes,12,opt,name=edge_delete,json=edgeDelete,proto3,oneof"`
}

type TransactionAction_IndexUpdate struct {
	IndexUpdate *Index `protobuf:"bytes,21,opt,name=index_update,json=indexUpdate,proto3,oneof"`
}

type TransactionAction_IndexDelete struct {
	IndexDelete *Index `protobuf:"bytes,22,opt,name=index_delete,json=indexDelete,proto3,oneof"`
}

type TransactionAction_MetaUpdate struct {
	MetaUpdate *Meta `protobuf:"bytes,31,opt,name=meta_update,json=metaUpdate,proto3,oneof"`
}

type TransactionAction_MetaDelete struct {
	MetaDelete *Meta `protobuf:"bytes,32,opt,name=meta_delete,json=metaDelete,proto3,oneof"`
}

type TransactionAction_NodeCreate struct {
	NodeCreate *Node `protobuf:"bytes,40,opt,name=node_create,json=nodeCreate,proto3,oneof"`
}

type TransactionAction_NodeUpdate struct {
	NodeUpdate *Node `protobuf:"bytes,41,opt,name=node_update,json=nodeUpdate,proto3,oneof"`
}

type TransactionAction_NodeDelete struct {
	NodeDelete *Node `protobuf:"bytes,42,opt,name=node_delete,json=nodeDelete,proto3,oneof"`
}

type TransactionAction_ReadCheck struct {
	ReadCheck *ReadCheckRequest `protobuf:"bytes,50,opt,name=read_check,json=readCheck,proto3,oneof"`
}

func (*TransactionAction_CounterIncrement) isTransactionAction_Action() {}

func (*TransactionAction_CounterDelete) isTransactionAction_Action() {}

func (*TransactionAction_CounterRegister) isTransactionAction_Action() {}

func (*TransactionAction_EdgeUpdate) isTransactionAction_Action() {}

func (*TransactionAction_EdgeDelete) isTransactionAction_Action() {}

func (*TransactionAction_IndexUpdate) isTransactionAction_Action() {}

func (*TransactionAction_IndexDelete) isTransactionAction_Action() {}

func (*TransactionAction_MetaUpdate) isTransactionAction_Action() {}

func (*TransactionAction_MetaDelete) isTransactionAction_Action() {}

func (*TransactionAction_NodeCreate) isTransactionAction_Action() {}

func (*TransactionAction_NodeUpdate) isTransactionAction_Action() {}

func (*TransactionAction_NodeDelete) isTransactionAction_Action() {}

func (*TransactionAction_ReadCheck) isTransactionAction_Action() {}

func (m *TransactionAction) GetAction() isTransactionAction_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *TransactionAction) GetCounterIncrement() *Counter {
	if x, ok := m.GetAction().(*TransactionAction_CounterIncrement); ok {
		return x.CounterIncrement
	}
	return nil
}

func (m *TransactionAction) GetCounterDelete() *Counter {
	if x, ok := m.GetAction().(*TransactionAction_CounterDelete); ok {
		return x.CounterDelete
	}
	return nil
}

func (m *TransactionAction) GetCounterRegister() *Counter {
	if x, ok := m.GetAction().(*TransactionAction_CounterRegister); ok {
		return x.CounterRegister
	}
	return nil
}

func (m *TransactionAction) GetEdgeUpdate() *Edge {
	if x, ok := m.GetAction().(*TransactionAction_EdgeUpdate); ok {
		return x.EdgeUpdate
	}
	return nil
}

func (m *TransactionAction) GetEdgeDelete() *Edge {
	if x, ok := m.GetAction().(*TransactionAction_EdgeDelete); ok {
		return x.EdgeDelete
	}
	return nil
}

func (m *TransactionAction) GetIndexUpdate() *Index {
	if x, ok := m.GetAction().(*TransactionAction_IndexUpdate); ok {
		return x.IndexUpdate
	}
	return nil
}

func (m *TransactionAction) GetIndexDelete() *Index {
	if x, ok := m.GetAction().(*TransactionAction_IndexDelete); ok {
		return x.IndexDelete
	}
	return nil
}

func (m *TransactionAction) GetMetaUpdate() *Meta {
	if x, ok := m.GetAction().(*TransactionAction_MetaUpdate); ok {
		return x.MetaUpdate
	}
	return nil
}

func (m *TransactionAction) GetMetaDelete() *Meta {
	if x, ok := m.GetAction().(*TransactionAction_MetaDelete); ok {
		return x.MetaDelete
	}
	return nil
}

func (m *TransactionAction) GetNodeCreate() *Node {
	if x, ok := m.GetAction().(*TransactionAction_NodeCreate); ok {
		return x.NodeCreate
	}
	return nil
}

func (m *TransactionAction) GetNodeUpdate() *Node {
	if x, ok := m.GetAction().(*TransactionAction_NodeUpdate); ok {
		return x.NodeUpdate
	}
	return nil
}

func (m *TransactionAction) GetNodeDelete() *Node {
	if x, ok := m.GetAction().(*TransactionAction_NodeDelete); ok {
		return x.NodeDelete
	}
	return nil
}

func (m *TransactionAction) GetReadCheck() *ReadCheckRequest {
	if x, ok := m.GetAction().(*TransactionAction_ReadCheck); ok {
		return x.ReadCheck
	}
	return nil
}

func (m *TransactionAction) GetActionId() uint32 {
	if m != nil {
		return m.ActionId
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TransactionAction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TransactionAction_CounterIncrement)(nil),
		(*TransactionAction_CounterDelete)(nil),
		(*TransactionAction_CounterRegister)(nil),
		(*TransactionAction_EdgeUpdate)(nil),
		(*TransactionAction_EdgeDelete)(nil),
		(*TransactionAction_IndexUpdate)(nil),
		(*TransactionAction_IndexDelete)(nil),
		(*TransactionAction_MetaUpdate)(nil),
		(*TransactionAction_MetaDelete)(nil),
		(*TransactionAction_NodeCreate)(nil),
		(*TransactionAction_NodeUpdate)(nil),
		(*TransactionAction_NodeDelete)(nil),
		(*TransactionAction_ReadCheck)(nil),
	}
}

type TransactionActionResponse struct {
	Status   MutationStatus `protobuf:"varint,1,opt,name=status,proto3,enum=cds.cabinet.v1.MutationStatus" json:"status,omitempty"`
	ActionId uint32         `protobuf:"varint,2,opt,name=action_id,json=actionId,proto3" json:"action_id,omitempty"`
	Error    string         `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	// Types that are valid to be assigned to Response:
	//	*TransactionActionResponse_NodeCreate
	//	*TransactionActionResponse_ReadCheck
	Response             isTransactionActionResponse_Response `protobuf_oneof:"response"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *TransactionActionResponse) Reset()         { *m = TransactionActionResponse{} }
func (m *TransactionActionResponse) String() string { return proto.CompactTextString(m) }
func (*TransactionActionResponse) ProtoMessage()    {}
func (*TransactionActionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eae9ea52da333b22, []int{1}
}

func (m *TransactionActionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionActionResponse.Unmarshal(m, b)
}
func (m *TransactionActionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionActionResponse.Marshal(b, m, deterministic)
}
func (m *TransactionActionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionActionResponse.Merge(m, src)
}
func (m *TransactionActionResponse) XXX_Size() int {
	return xxx_messageInfo_TransactionActionResponse.Size(m)
}
func (m *TransactionActionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionActionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionActionResponse proto.InternalMessageInfo

func (m *TransactionActionResponse) GetStatus() MutationStatus {
	if m != nil {
		return m.Status
	}
	return MutationStatus_SUCCESS
}

func (m *TransactionActionResponse) GetActionId() uint32 {
	if m != nil {
		return m.ActionId
	}
	return 0
}

func (m *TransactionActionResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type isTransactionActionResponse_Response interface {
	isTransactionActionResponse_Response()
}

type TransactionActionResponse_NodeCreate struct {
	NodeCreate *NodeCreateResponse `protobuf:"bytes,20,opt,name=node_create,json=nodeCreate,proto3,oneof"`
}

type TransactionActionResponse_ReadCheck struct {
	ReadCheck *ReadCheckResponse `protobuf:"bytes,30,opt,name=read_check,json=readCheck,proto3,oneof"`
}

func (*TransactionActionResponse_NodeCreate) isTransactionActionResponse_Response() {}

func (*TransactionActionResponse_ReadCheck) isTransactionActionResponse_Response() {}

func (m *TransactionActionResponse) GetResponse() isTransactionActionResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *TransactionActionResponse) GetNodeCreate() *NodeCreateResponse {
	if x, ok := m.GetResponse().(*TransactionActionResponse_NodeCreate); ok {
		return x.NodeCreate
	}
	return nil
}

func (m *TransactionActionResponse) GetReadCheck() *ReadCheckResponse {
	if x, ok := m.GetResponse().(*TransactionActionResponse_ReadCheck); ok {
		return x.ReadCheck
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TransactionActionResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TransactionActionResponse_NodeCreate)(nil),
		(*TransactionActionResponse_ReadCheck)(nil),
	}
}

type ReadCheckRequest struct {
	Source               string         `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Operator             CheckOperators `protobuf:"varint,2,opt,name=operator,proto3,enum=cds.cabinet.v1.CheckOperators" json:"operator,omitempty"`
	Target               *CheckTarget   `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReadCheckRequest) Reset()         { *m = ReadCheckRequest{} }
func (m *ReadCheckRequest) String() string { return proto.CompactTextString(m) }
func (*ReadCheckRequest) ProtoMessage()    {}
func (*ReadCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eae9ea52da333b22, []int{2}
}

func (m *ReadCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadCheckRequest.Unmarshal(m, b)
}
func (m *ReadCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadCheckRequest.Marshal(b, m, deterministic)
}
func (m *ReadCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadCheckRequest.Merge(m, src)
}
func (m *ReadCheckRequest) XXX_Size() int {
	return xxx_messageInfo_ReadCheckRequest.Size(m)
}
func (m *ReadCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadCheckRequest proto.InternalMessageInfo

func (m *ReadCheckRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *ReadCheckRequest) GetOperator() CheckOperators {
	if m != nil {
		return m.Operator
	}
	return CheckOperators_EXISTS
}

func (m *ReadCheckRequest) GetTarget() *CheckTarget {
	if m != nil {
		return m.Target
	}
	return nil
}

type ReadCheckResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadCheckResponse) Reset()         { *m = ReadCheckResponse{} }
func (m *ReadCheckResponse) String() string { return proto.CompactTextString(m) }
func (*ReadCheckResponse) ProtoMessage()    {}
func (*ReadCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eae9ea52da333b22, []int{3}
}

func (m *ReadCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadCheckResponse.Unmarshal(m, b)
}
func (m *ReadCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadCheckResponse.Marshal(b, m, deterministic)
}
func (m *ReadCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadCheckResponse.Merge(m, src)
}
func (m *ReadCheckResponse) XXX_Size() int {
	return xxx_messageInfo_ReadCheckResponse.Size(m)
}
func (m *ReadCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadCheckResponse proto.InternalMessageInfo

func (m *ReadCheckResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type CheckTarget struct {
	// Types that are valid to be assigned to Target:
	//	*CheckTarget_Val
	//	*CheckTarget_Iri
	Target               isCheckTarget_Target `protobuf_oneof:"target"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CheckTarget) Reset()         { *m = CheckTarget{} }
func (m *CheckTarget) String() string { return proto.CompactTextString(m) }
func (*CheckTarget) ProtoMessage()    {}
func (*CheckTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_eae9ea52da333b22, []int{4}
}

func (m *CheckTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckTarget.Unmarshal(m, b)
}
func (m *CheckTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckTarget.Marshal(b, m, deterministic)
}
func (m *CheckTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckTarget.Merge(m, src)
}
func (m *CheckTarget) XXX_Size() int {
	return xxx_messageInfo_CheckTarget.Size(m)
}
func (m *CheckTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckTarget.DiscardUnknown(m)
}

var xxx_messageInfo_CheckTarget proto.InternalMessageInfo

type isCheckTarget_Target interface {
	isCheckTarget_Target()
}

type CheckTarget_Val struct {
	Val string `protobuf:"bytes,1,opt,name=val,proto3,oneof"`
}

type CheckTarget_Iri struct {
	Iri string `protobuf:"bytes,2,opt,name=iri,proto3,oneof"`
}

func (*CheckTarget_Val) isCheckTarget_Target() {}

func (*CheckTarget_Iri) isCheckTarget_Target() {}

func (m *CheckTarget) GetTarget() isCheckTarget_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *CheckTarget) GetVal() string {
	if x, ok := m.GetTarget().(*CheckTarget_Val); ok {
		return x.Val
	}
	return ""
}

func (m *CheckTarget) GetIri() string {
	if x, ok := m.GetTarget().(*CheckTarget_Iri); ok {
		return x.Iri
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CheckTarget) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CheckTarget_Val)(nil),
		(*CheckTarget_Iri)(nil),
	}
}

func init() {
	proto.RegisterType((*TransactionAction)(nil), "cds.cabinet.v1.TransactionAction")
	proto.RegisterType((*TransactionActionResponse)(nil), "cds.cabinet.v1.TransactionActionResponse")
	proto.RegisterType((*ReadCheckRequest)(nil), "cds.cabinet.v1.ReadCheckRequest")
	proto.RegisterType((*ReadCheckResponse)(nil), "cds.cabinet.v1.ReadCheckResponse")
	proto.RegisterType((*CheckTarget)(nil), "cds.cabinet.v1.CheckTarget")
}

func init() { proto.RegisterFile("service_transaction.proto", fileDescriptor_eae9ea52da333b22) }

var fileDescriptor_eae9ea52da333b22 = []byte{
	// 624 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcb, 0x4e, 0xdb, 0x4c,
	0x14, 0x4e, 0x40, 0xbf, 0x95, 0x1c, 0x03, 0x3f, 0x4c, 0x81, 0x1a, 0x90, 0x68, 0xea, 0x15, 0x6d,
	0xa5, 0x48, 0x05, 0xa9, 0x48, 0x6c, 0x5a, 0x6e, 0x55, 0x58, 0xf4, 0xa2, 0x29, 0x5d, 0x5b, 0x83,
	0xe7, 0x28, 0x58, 0x05, 0x4f, 0x3a, 0x33, 0x46, 0x7d, 0x97, 0xaa, 0x6f, 0xd5, 0x07, 0xaa, 0xe6,
	0x96, 0xc4, 0x26, 0xaa, 0xbb, 0x89, 0x3c, 0x27, 0xdf, 0xe5, 0x9c, 0x6f, 0x2e, 0xb0, 0xa3, 0x50,
	0x3e, 0x14, 0x39, 0x66, 0x5a, 0xb2, 0x52, 0xb1, 0x5c, 0x17, 0xa2, 0x1c, 0x4e, 0xa4, 0xd0, 0x82,
	0xac, 0xe5, 0x5c, 0x0d, 0x73, 0x76, 0x53, 0x94, 0xa8, 0x87, 0x0f, 0xaf, 0x77, 0xb7, 0x02, 0x34,
	0x17, 0x55, 0xa9, 0x51, 0x3a, 0xd8, 0x2e, 0x09, 0x65, 0xe4, 0x63, 0xf4, 0xb5, 0x27, 0xa1, 0x56,
	0x94, 0x1c, 0x7f, 0x34, 0x81, 0xf7, 0xa8, 0x59, 0xb3, 0x56, 0x0a, 0x1e, 0xc8, 0xab, 0x75, 0x7d,
	0x98, 0xd3, 0x8d, 0xe7, 0xf5, 0x60, 0x4e, 0x07, 0xe6, 0xf8, 0x20, 0x26, 0x5a, 0xf9, 0xef, 0x15,
	0x75, 0xcb, 0x24, 0x72, 0xb7, 0x4a, 0x7f, 0x47, 0xb0, 0x71, 0x3d, 0x9b, 0xf3, 0xd4, 0xfe, 0x92,
	0xf7, 0xb0, 0xe1, 0x1d, 0xb3, 0xa2, 0xcc, 0x25, 0xde, 0x63, 0xa9, 0x93, 0xee, 0xa0, 0x7b, 0x10,
	0x1f, 0x3e, 0x1d, 0xd6, 0x33, 0x18, 0x9e, 0x3b, 0xe0, 0xa8, 0x43, 0xd7, 0x3d, 0xe7, 0x2a, 0x50,
	0xc8, 0x3b, 0x58, 0x0b, 0x3a, 0x1c, 0xef, 0x50, 0x63, 0xb2, 0xd4, 0x26, 0x12, 0x46, 0xbd, 0xb0,
	0x78, 0x72, 0x01, 0x41, 0x35, 0x93, 0x38, 0x2e, 0x94, 0x46, 0x99, 0x2c, 0xb7, 0x69, 0xfc, 0xef,
	0x29, 0xd4, 0x33, 0xc8, 0x31, 0xc4, 0x26, 0xb2, 0xac, 0x9a, 0x70, 0xa6, 0x31, 0x89, 0xad, 0xc0,
	0x66, 0x53, 0xe0, 0x92, 0x8f, 0x71, 0xd4, 0xa1, 0x36, 0xdd, 0xaf, 0x16, 0x39, 0x25, 0xfa, 0xee,
	0x57, 0xda, 0x89, 0xbe, 0xef, 0x13, 0x58, 0xb1, 0x1b, 0x13, 0x2c, 0xb7, 0x2c, 0x73, 0xab, 0xc9,
	0xbc, 0x32, 0x98, 0x51, 0x87, 0xba, 0x5d, 0xf4, 0xa6, 0x53, 0xae, 0x77, 0xdd, 0xfe, 0x17, 0xae,
	0xf7, 0x3d, 0x86, 0xd8, 0x9c, 0x81, 0x60, 0xfb, 0x6c, 0x71, 0xc3, 0x1f, 0x50, 0x33, 0xd3, 0xb0,
	0x81, 0xce, 0x26, 0xb5, 0x44, 0xef, 0x39, 0x68, 0x27, 0xce, 0x1c, 0xcd, 0x49, 0xcb, 0x72, 0x89,
	0xc6, 0xf1, 0x60, 0x31, 0xf1, 0xa3, 0xe0, 0x36, 0x22, 0x03, 0x3d, 0xb7, 0xc8, 0x29, 0xd1, 0xb7,
	0xfa, 0xa2, 0x9d, 0x38, 0x6b, 0xd5, 0x12, 0x7d, 0xab, 0x2f, 0xdb, 0x89, 0xbe, 0xd5, 0x53, 0x00,
	0x89, 0x8c, 0x67, 0xf9, 0x2d, 0xe6, 0xdf, 0x92, 0x43, 0xcb, 0x1b, 0x34, 0x79, 0x14, 0x19, 0x3f,
	0x37, 0x00, 0x8a, 0xdf, 0x2b, 0x54, 0x7a, 0xd4, 0xa1, 0x7d, 0x19, 0x6a, 0x64, 0x0f, 0xfa, 0xee,
	0xa6, 0x64, 0x05, 0x4f, 0x3e, 0x0f, 0xba, 0x07, 0xab, 0xb4, 0xe7, 0x0a, 0x57, 0xfc, 0xac, 0x07,
	0x91, 0xfb, 0x4e, 0x7f, 0x2d, 0xc1, 0xce, 0xa3, 0x6b, 0x45, 0x51, 0x4d, 0x44, 0xa9, 0x90, 0xbc,
	0x81, 0x48, 0x69, 0xa6, 0x2b, 0x65, 0xef, 0xd4, 0xda, 0xe1, 0xfe, 0xa3, 0x98, 0x2b, 0xcd, 0x0c,
	0xe3, 0x8b, 0x45, 0x51, 0x8f, 0xae, 0x9b, 0x2f, 0xd5, 0xcd, 0xc9, 0x26, 0xfc, 0x87, 0x52, 0x0a,
	0x77, 0x3d, 0xfa, 0xd4, 0x2d, 0xc8, 0x65, 0x7d, 0x77, 0x36, 0xed, 0xcc, 0xe9, 0xa2, 0xac, 0xdc,
	0xae, 0x84, 0x1e, 0x1b, 0x7b, 0x75, 0x56, 0x4b, 0x6e, 0xdf, 0xaa, 0x3c, 0xff, 0x4b, 0x72, 0x53,
	0x91, 0x59, 0x74, 0x67, 0x00, 0x3d, 0xe9, 0xff, 0x48, 0x7f, 0x76, 0x61, 0xbd, 0x19, 0x34, 0xd9,
	0x86, 0x48, 0x89, 0x4a, 0xe6, 0x68, 0x63, 0xe9, 0x53, 0xbf, 0x22, 0x27, 0xd0, 0x13, 0x13, 0x94,
	0x4c, 0x0b, 0x69, 0xa7, 0x5e, 0x10, 0x98, 0xd5, 0xf9, 0xe4, 0x41, 0x8a, 0x4e, 0xf1, 0xe4, 0x08,
	0x22, 0xcd, 0xe4, 0x18, 0xb5, 0x7f, 0x35, 0xf6, 0x16, 0x32, 0xaf, 0x2d, 0x84, 0x7a, 0x68, 0xfa,
	0x0a, 0x36, 0x1e, 0xcd, 0x62, 0xba, 0x93, 0xa8, 0xaa, 0x3b, 0xf7, 0x10, 0xf6, 0xa8, 0x5f, 0xa5,
	0x6f, 0x21, 0x9e, 0xd3, 0x20, 0x04, 0x96, 0x1f, 0xd8, 0x9d, 0x9b, 0x60, 0xd4, 0xa1, 0x66, 0x61,
	0x6a, 0x85, 0x2c, 0x6c, 0xef, 0xb6, 0x56, 0xc8, 0xc2, 0x9c, 0x15, 0xe7, 0x76, 0x13, 0xd9, 0x97,
	0xf8, 0xe8, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x63, 0xd1, 0xc8, 0x78, 0x06, 0x00, 0x00,
}

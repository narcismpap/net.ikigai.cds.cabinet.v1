// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cabinet.proto

package cds_cabinet_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

func init() { proto.RegisterFile("cabinet.proto", fileDescriptor_d677e6cffaacc14d) }

var fileDescriptor_d677e6cffaacc14d = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xdd, 0x6e, 0xd3, 0x40,
	0x10, 0x85, 0xe5, 0x1b, 0xea, 0x4e, 0x49, 0x55, 0x0d, 0x54, 0xc0, 0x5e, 0x50, 0x2a, 0x81, 0x04,
	0x37, 0x96, 0x81, 0x07, 0xe0, 0xc7, 0x85, 0x8a, 0x9f, 0xa2, 0x2a, 0x05, 0x6e, 0xab, 0xad, 0x77,
	0xd4, 0x5a, 0x84, 0x75, 0xf0, 0xae, 0x23, 0x5e, 0x99, 0xb7, 0x40, 0x6b, 0xef, 0xc6, 0x1b, 0x7b,
	0xed, 0x48, 0xbd, 0x89, 0xe3, 0xf9, 0xce, 0x9e, 0x39, 0x9e, 0x49, 0x0c, 0xb3, 0x9c, 0x5f, 0x15,
	0x92, 0x74, 0xb2, 0xac, 0x4a, 0x5d, 0xe2, 0x7e, 0x2e, 0x54, 0xe2, 0x4a, 0xab, 0x97, 0xec, 0x50,
	0x51, 0xb5, 0x2a, 0x72, 0xba, 0xcc, 0xcb, 0x5a, 0x6a, 0xaa, 0x5a, 0x19, 0x43, 0x57, 0x26, 0x71,
	0x4d, 0xb6, 0x76, 0xcf, 0xd5, 0x0a, 0x29, 0xe8, 0x6f, 0x5f, 0xf8, 0x9b, 0x34, 0xef, 0xd7, 0x64,
	0x29, 0xdc, 0xe1, 0x87, 0xae, 0xa6, 0xe8, 0x4f, 0x4d, 0x52, 0x17, 0x7c, 0x61, 0xc9, 0x23, 0x47,
	0x74, 0xc5, 0xa5, 0xe2, 0xb9, 0x2e, 0x4a, 0x69, 0xd1, 0x6c, 0x33, 0x14, 0x78, 0x61, 0xf6, 0xfc,
	0x10, 0xe0, 0x35, 0x07, 0xaf, 0xe9, 0xc1, 0xa0, 0xd9, 0x5d, 0x75, 0xc3, 0x2b, 0x12, 0xed, 0xdd,
	0xab, 0x7f, 0x31, 0x40, 0x76, 0x72, 0x91, 0xb5, 0xe3, 0xc0, 0x2f, 0x00, 0x59, 0xdb, 0xf0, 0x94,
	0x34, 0x3e, 0x48, 0x36, 0x47, 0x95, 0x58, 0xc6, 0x9e, 0x8e, 0x80, 0x9f, 0x7c, 0x51, 0xd3, 0x9c,
	0xd4, 0xb2, 0x94, 0x8a, 0xf0, 0x0d, 0xec, 0x7c, 0x10, 0xd7, 0x64, 0x9c, 0x1e, 0xf7, 0x0f, 0x58,
	0x30, 0x37, 0xe9, 0x94, 0x66, 0xf7, 0x43, 0x1c, 0x33, 0x88, 0xcd, 0xf5, 0x6b, 0xa1, 0x34, 0x1e,
	0x85, 0x14, 0x86, 0x4c, 0x5a, 0xa4, 0x11, 0xbe, 0x87, 0xf8, 0x93, 0x19, 0x94, 0x89, 0x31, 0x30,
	0x71, 0xc4, 0x99, 0x1c, 0x06, 0x05, 0xf8, 0x11, 0x76, 0x9b, 0x2f, 0x4d, 0x92, 0x27, 0x41, 0x8d,
	0x1f, 0x25, 0xec, 0x92, 0x46, 0xf8, 0x16, 0x76, 0xce, 0x48, 0x73, 0x13, 0x65, 0x10, 0xd7, 0x00,
	0x76, 0x14, 0xaa, 0x36, 0xf9, 0xec, 0x4c, 0x33, 0x88, 0x4d, 0x29, 0x3c, 0x12, 0x47, 0x46, 0x47,
	0x62, 0x04, 0x69, 0x64, 0x16, 0xf3, 0xad, 0x14, 0xe1, 0xc5, 0x58, 0x30, 0x6a, 0x61, 0xb8, 0x49,
	0x61, 0xae, 0xe1, 0x14, 0x8e, 0x4c, 0x5a, 0xa4, 0x11, 0x5e, 0xc2, 0xde, 0xf7, 0xee, 0xf7, 0x8e,
	0xc7, 0x7d, 0x99, 0x07, 0xdf, 0x35, 0x9f, 0xec, 0xc5, 0x56, 0x89, 0x9b, 0xd3, 0xf3, 0x28, 0x8d,
	0xf0, 0x1c, 0x76, 0xe7, 0xc4, 0x45, 0x76, 0x43, 0xf9, 0xaf, 0xe1, 0xd6, 0xd6, 0xc8, 0xe5, 0x3c,
	0x9e, 0x50, 0xd8, 0xe9, 0x7f, 0x86, 0x83, 0x8b, 0xf5, 0xff, 0x29, 0xab, 0x88, 0x6b, 0x42, 0xd6,
	0x3f, 0xd6, 0x29, 0xd8, 0x04, 0xc3, 0x73, 0xdf, 0xeb, 0x84, 0x16, 0xb4, 0xc5, 0x6b, 0xf0, 0x00,
	0x67, 0xb5, 0xe6, 0xfe, 0x33, 0xe3, 0x29, 0xcc, 0x3a, 0xbd, 0x59, 0xee, 0x6d, 0xa3, 0xfd, 0x80,
	0xfd, 0xee, 0xae, 0x59, 0xf2, 0xb3, 0x71, 0xb5, 0xbf, 0xea, 0x09, 0xd3, 0x34, 0xba, 0xba, 0xd3,
	0xbc, 0x72, 0x5e, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xf7, 0x35, 0xa3, 0x90, 0x05, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CDSCabinetClient is the client API for CDSCabinet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CDSCabinetClient interface {
	// @required: node_id|Edge, value
	// @throws: NodeInvalidID, EdgeInvalidSubject, EdgeInvalidTarget, EdgeInvalidPredicate
	CounterGet(ctx context.Context, in *Counter, opts ...grpc.CallOption) (*CounterValueResponse, error)
	// @required: edge.subject, edge.target, edge.predicate
	// @throws: EdgeInvalidSubject, EdgeInvalidTarget, EdgeInvalidPredicate, EdgeNotFound
	EdgeGet(ctx context.Context, in *EdgeGetRequest, opts ...grpc.CallOption) (*Edge, error)
	// @required: edge.subject
	// @optional: edge.predicate, edge.resume_from, edge.page_size, edge.mode
	// @throws: ListNoPagination, EdgeInvalidSubject, EdgeInvalidPredicate
	EdgeList(ctx context.Context, in *EdgeListRequest, opts ...grpc.CallOption) (CDSCabinet_EdgeListClient, error)
	// @required: index.type, index.field, index.node
	// @throws: IndexNotFound, IndexInvalidID, IndexInvalidQuery
	IndexGet(ctx context.Context, in *IndexGetRequest, opts ...grpc.CallOption) (*Index, error)
	// @required: index.type, index.field
	// @throws: IndexInvalidID
	IndexList(ctx context.Context, in *IndexListRequest, opts ...grpc.CallOption) (CDSCabinet_IndexListClient, error)
	// @required: field.edge|field.node_id, key
	// @throws: MetaNotFound, MetaInvalidObject, MetaInvalidKey
	MetaGet(ctx context.Context, in *Meta, opts ...grpc.CallOption) (*MetaGetResponse, error)
	// @required: field.edge|field.node_id
	// @throws: MetaInvalidObject, MetaInvalidKey
	MetaList(ctx context.Context, in *MetaListRequest, opts ...grpc.CallOption) (CDSCabinet_MetaListClient, error)
	// @required: id, type
	// @throws: NodeNotFound, NodeInvalidID, NodeInvalidType
	NodeGet(ctx context.Context, in *NodeGetRequest, opts ...grpc.CallOption) (*Node, error)
	// List node with pagination
	// @required: type, per_page
	NodeList(ctx context.Context, in *NodeListRequest, opts ...grpc.CallOption) (CDSCabinet_NodeListClient, error)
	// @required: <action.*>, retry_mode
	// @throws: TransactionInvalidAction, TransactionSyntaxError
	Transaction(ctx context.Context, opts ...grpc.CallOption) (CDSCabinet_TransactionClient, error)
	// @required: source, operator, target
	// @throws: @ReadCheckNaN
	ReadCheck(ctx context.Context, in *ReadCheckRequest, opts ...grpc.CallOption) (*ReadCheckResponse, error)
	// @required: type, node
	// @throws: NodeNotFound, SequentialInvalidType
	SequentialCreate(ctx context.Context, in *Sequential, opts ...grpc.CallOption) (*Sequential, error)
	// @required: type, seqid, node
	// @throws: NodeNotFound, SequentialInvalidType
	SequentialDelete(ctx context.Context, in *Sequential, opts ...grpc.CallOption) (*MutationResponse, error)
	// @required: type, seqid
	// @throws: NodeNotFound, SequentialInvalidType
	SequentialGet(ctx context.Context, in *Sequential, opts ...grpc.CallOption) (*Sequential, error)
	// @required: type, per_page
	// @throws: SequentialInvalidType
	SequentialList(ctx context.Context, in *SequentialListRequest, opts ...grpc.CallOption) (CDSCabinet_SequentialListClient, error)
}

type cDSCabinetClient struct {
	cc *grpc.ClientConn
}

func NewCDSCabinetClient(cc *grpc.ClientConn) CDSCabinetClient {
	return &cDSCabinetClient{cc}
}

func (c *cDSCabinetClient) CounterGet(ctx context.Context, in *Counter, opts ...grpc.CallOption) (*CounterValueResponse, error) {
	out := new(CounterValueResponse)
	err := c.cc.Invoke(ctx, "/cds.cabinet.v1.CDSCabinet/CounterGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cDSCabinetClient) EdgeGet(ctx context.Context, in *EdgeGetRequest, opts ...grpc.CallOption) (*Edge, error) {
	out := new(Edge)
	err := c.cc.Invoke(ctx, "/cds.cabinet.v1.CDSCabinet/EdgeGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cDSCabinetClient) EdgeList(ctx context.Context, in *EdgeListRequest, opts ...grpc.CallOption) (CDSCabinet_EdgeListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CDSCabinet_serviceDesc.Streams[0], "/cds.cabinet.v1.CDSCabinet/EdgeList", opts...)
	if err != nil {
		return nil, err
	}
	x := &cDSCabinetEdgeListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CDSCabinet_EdgeListClient interface {
	Recv() (*Edge, error)
	grpc.ClientStream
}

type cDSCabinetEdgeListClient struct {
	grpc.ClientStream
}

func (x *cDSCabinetEdgeListClient) Recv() (*Edge, error) {
	m := new(Edge)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cDSCabinetClient) IndexGet(ctx context.Context, in *IndexGetRequest, opts ...grpc.CallOption) (*Index, error) {
	out := new(Index)
	err := c.cc.Invoke(ctx, "/cds.cabinet.v1.CDSCabinet/IndexGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cDSCabinetClient) IndexList(ctx context.Context, in *IndexListRequest, opts ...grpc.CallOption) (CDSCabinet_IndexListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CDSCabinet_serviceDesc.Streams[1], "/cds.cabinet.v1.CDSCabinet/IndexList", opts...)
	if err != nil {
		return nil, err
	}
	x := &cDSCabinetIndexListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CDSCabinet_IndexListClient interface {
	Recv() (*Index, error)
	grpc.ClientStream
}

type cDSCabinetIndexListClient struct {
	grpc.ClientStream
}

func (x *cDSCabinetIndexListClient) Recv() (*Index, error) {
	m := new(Index)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cDSCabinetClient) MetaGet(ctx context.Context, in *Meta, opts ...grpc.CallOption) (*MetaGetResponse, error) {
	out := new(MetaGetResponse)
	err := c.cc.Invoke(ctx, "/cds.cabinet.v1.CDSCabinet/MetaGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cDSCabinetClient) MetaList(ctx context.Context, in *MetaListRequest, opts ...grpc.CallOption) (CDSCabinet_MetaListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CDSCabinet_serviceDesc.Streams[2], "/cds.cabinet.v1.CDSCabinet/MetaList", opts...)
	if err != nil {
		return nil, err
	}
	x := &cDSCabinetMetaListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CDSCabinet_MetaListClient interface {
	Recv() (*Meta, error)
	grpc.ClientStream
}

type cDSCabinetMetaListClient struct {
	grpc.ClientStream
}

func (x *cDSCabinetMetaListClient) Recv() (*Meta, error) {
	m := new(Meta)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cDSCabinetClient) NodeGet(ctx context.Context, in *NodeGetRequest, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/cds.cabinet.v1.CDSCabinet/NodeGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cDSCabinetClient) NodeList(ctx context.Context, in *NodeListRequest, opts ...grpc.CallOption) (CDSCabinet_NodeListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CDSCabinet_serviceDesc.Streams[3], "/cds.cabinet.v1.CDSCabinet/NodeList", opts...)
	if err != nil {
		return nil, err
	}
	x := &cDSCabinetNodeListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CDSCabinet_NodeListClient interface {
	Recv() (*Node, error)
	grpc.ClientStream
}

type cDSCabinetNodeListClient struct {
	grpc.ClientStream
}

func (x *cDSCabinetNodeListClient) Recv() (*Node, error) {
	m := new(Node)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cDSCabinetClient) Transaction(ctx context.Context, opts ...grpc.CallOption) (CDSCabinet_TransactionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CDSCabinet_serviceDesc.Streams[4], "/cds.cabinet.v1.CDSCabinet/Transaction", opts...)
	if err != nil {
		return nil, err
	}
	x := &cDSCabinetTransactionClient{stream}
	return x, nil
}

type CDSCabinet_TransactionClient interface {
	Send(*TransactionAction) error
	Recv() (*TransactionActionResponse, error)
	grpc.ClientStream
}

type cDSCabinetTransactionClient struct {
	grpc.ClientStream
}

func (x *cDSCabinetTransactionClient) Send(m *TransactionAction) error {
	return x.ClientStream.SendMsg(m)
}

func (x *cDSCabinetTransactionClient) Recv() (*TransactionActionResponse, error) {
	m := new(TransactionActionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cDSCabinetClient) ReadCheck(ctx context.Context, in *ReadCheckRequest, opts ...grpc.CallOption) (*ReadCheckResponse, error) {
	out := new(ReadCheckResponse)
	err := c.cc.Invoke(ctx, "/cds.cabinet.v1.CDSCabinet/ReadCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cDSCabinetClient) SequentialCreate(ctx context.Context, in *Sequential, opts ...grpc.CallOption) (*Sequential, error) {
	out := new(Sequential)
	err := c.cc.Invoke(ctx, "/cds.cabinet.v1.CDSCabinet/SequentialCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cDSCabinetClient) SequentialDelete(ctx context.Context, in *Sequential, opts ...grpc.CallOption) (*MutationResponse, error) {
	out := new(MutationResponse)
	err := c.cc.Invoke(ctx, "/cds.cabinet.v1.CDSCabinet/SequentialDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cDSCabinetClient) SequentialGet(ctx context.Context, in *Sequential, opts ...grpc.CallOption) (*Sequential, error) {
	out := new(Sequential)
	err := c.cc.Invoke(ctx, "/cds.cabinet.v1.CDSCabinet/SequentialGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cDSCabinetClient) SequentialList(ctx context.Context, in *SequentialListRequest, opts ...grpc.CallOption) (CDSCabinet_SequentialListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CDSCabinet_serviceDesc.Streams[5], "/cds.cabinet.v1.CDSCabinet/SequentialList", opts...)
	if err != nil {
		return nil, err
	}
	x := &cDSCabinetSequentialListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CDSCabinet_SequentialListClient interface {
	Recv() (*Sequential, error)
	grpc.ClientStream
}

type cDSCabinetSequentialListClient struct {
	grpc.ClientStream
}

func (x *cDSCabinetSequentialListClient) Recv() (*Sequential, error) {
	m := new(Sequential)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CDSCabinetServer is the server API for CDSCabinet service.
type CDSCabinetServer interface {
	// @required: node_id|Edge, value
	// @throws: NodeInvalidID, EdgeInvalidSubject, EdgeInvalidTarget, EdgeInvalidPredicate
	CounterGet(context.Context, *Counter) (*CounterValueResponse, error)
	// @required: edge.subject, edge.target, edge.predicate
	// @throws: EdgeInvalidSubject, EdgeInvalidTarget, EdgeInvalidPredicate, EdgeNotFound
	EdgeGet(context.Context, *EdgeGetRequest) (*Edge, error)
	// @required: edge.subject
	// @optional: edge.predicate, edge.resume_from, edge.page_size, edge.mode
	// @throws: ListNoPagination, EdgeInvalidSubject, EdgeInvalidPredicate
	EdgeList(*EdgeListRequest, CDSCabinet_EdgeListServer) error
	// @required: index.type, index.field, index.node
	// @throws: IndexNotFound, IndexInvalidID, IndexInvalidQuery
	IndexGet(context.Context, *IndexGetRequest) (*Index, error)
	// @required: index.type, index.field
	// @throws: IndexInvalidID
	IndexList(*IndexListRequest, CDSCabinet_IndexListServer) error
	// @required: field.edge|field.node_id, key
	// @throws: MetaNotFound, MetaInvalidObject, MetaInvalidKey
	MetaGet(context.Context, *Meta) (*MetaGetResponse, error)
	// @required: field.edge|field.node_id
	// @throws: MetaInvalidObject, MetaInvalidKey
	MetaList(*MetaListRequest, CDSCabinet_MetaListServer) error
	// @required: id, type
	// @throws: NodeNotFound, NodeInvalidID, NodeInvalidType
	NodeGet(context.Context, *NodeGetRequest) (*Node, error)
	// List node with pagination
	// @required: type, per_page
	NodeList(*NodeListRequest, CDSCabinet_NodeListServer) error
	// @required: <action.*>, retry_mode
	// @throws: TransactionInvalidAction, TransactionSyntaxError
	Transaction(CDSCabinet_TransactionServer) error
	// @required: source, operator, target
	// @throws: @ReadCheckNaN
	ReadCheck(context.Context, *ReadCheckRequest) (*ReadCheckResponse, error)
	// @required: type, node
	// @throws: NodeNotFound, SequentialInvalidType
	SequentialCreate(context.Context, *Sequential) (*Sequential, error)
	// @required: type, seqid, node
	// @throws: NodeNotFound, SequentialInvalidType
	SequentialDelete(context.Context, *Sequential) (*MutationResponse, error)
	// @required: type, seqid
	// @throws: NodeNotFound, SequentialInvalidType
	SequentialGet(context.Context, *Sequential) (*Sequential, error)
	// @required: type, per_page
	// @throws: SequentialInvalidType
	SequentialList(*SequentialListRequest, CDSCabinet_SequentialListServer) error
}

func RegisterCDSCabinetServer(s *grpc.Server, srv CDSCabinetServer) {
	s.RegisterService(&_CDSCabinet_serviceDesc, srv)
}

func _CDSCabinet_CounterGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Counter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CDSCabinetServer).CounterGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cds.cabinet.v1.CDSCabinet/CounterGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CDSCabinetServer).CounterGet(ctx, req.(*Counter))
	}
	return interceptor(ctx, in, info, handler)
}

func _CDSCabinet_EdgeGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EdgeGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CDSCabinetServer).EdgeGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cds.cabinet.v1.CDSCabinet/EdgeGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CDSCabinetServer).EdgeGet(ctx, req.(*EdgeGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CDSCabinet_EdgeList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EdgeListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CDSCabinetServer).EdgeList(m, &cDSCabinetEdgeListServer{stream})
}

type CDSCabinet_EdgeListServer interface {
	Send(*Edge) error
	grpc.ServerStream
}

type cDSCabinetEdgeListServer struct {
	grpc.ServerStream
}

func (x *cDSCabinetEdgeListServer) Send(m *Edge) error {
	return x.ServerStream.SendMsg(m)
}

func _CDSCabinet_IndexGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CDSCabinetServer).IndexGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cds.cabinet.v1.CDSCabinet/IndexGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CDSCabinetServer).IndexGet(ctx, req.(*IndexGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CDSCabinet_IndexList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(IndexListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CDSCabinetServer).IndexList(m, &cDSCabinetIndexListServer{stream})
}

type CDSCabinet_IndexListServer interface {
	Send(*Index) error
	grpc.ServerStream
}

type cDSCabinetIndexListServer struct {
	grpc.ServerStream
}

func (x *cDSCabinetIndexListServer) Send(m *Index) error {
	return x.ServerStream.SendMsg(m)
}

func _CDSCabinet_MetaGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Meta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CDSCabinetServer).MetaGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cds.cabinet.v1.CDSCabinet/MetaGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CDSCabinetServer).MetaGet(ctx, req.(*Meta))
	}
	return interceptor(ctx, in, info, handler)
}

func _CDSCabinet_MetaList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MetaListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CDSCabinetServer).MetaList(m, &cDSCabinetMetaListServer{stream})
}

type CDSCabinet_MetaListServer interface {
	Send(*Meta) error
	grpc.ServerStream
}

type cDSCabinetMetaListServer struct {
	grpc.ServerStream
}

func (x *cDSCabinetMetaListServer) Send(m *Meta) error {
	return x.ServerStream.SendMsg(m)
}

func _CDSCabinet_NodeGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CDSCabinetServer).NodeGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cds.cabinet.v1.CDSCabinet/NodeGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CDSCabinetServer).NodeGet(ctx, req.(*NodeGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CDSCabinet_NodeList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NodeListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CDSCabinetServer).NodeList(m, &cDSCabinetNodeListServer{stream})
}

type CDSCabinet_NodeListServer interface {
	Send(*Node) error
	grpc.ServerStream
}

type cDSCabinetNodeListServer struct {
	grpc.ServerStream
}

func (x *cDSCabinetNodeListServer) Send(m *Node) error {
	return x.ServerStream.SendMsg(m)
}

func _CDSCabinet_Transaction_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CDSCabinetServer).Transaction(&cDSCabinetTransactionServer{stream})
}

type CDSCabinet_TransactionServer interface {
	Send(*TransactionActionResponse) error
	Recv() (*TransactionAction, error)
	grpc.ServerStream
}

type cDSCabinetTransactionServer struct {
	grpc.ServerStream
}

func (x *cDSCabinetTransactionServer) Send(m *TransactionActionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *cDSCabinetTransactionServer) Recv() (*TransactionAction, error) {
	m := new(TransactionAction)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CDSCabinet_ReadCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CDSCabinetServer).ReadCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cds.cabinet.v1.CDSCabinet/ReadCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CDSCabinetServer).ReadCheck(ctx, req.(*ReadCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CDSCabinet_SequentialCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sequential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CDSCabinetServer).SequentialCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cds.cabinet.v1.CDSCabinet/SequentialCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CDSCabinetServer).SequentialCreate(ctx, req.(*Sequential))
	}
	return interceptor(ctx, in, info, handler)
}

func _CDSCabinet_SequentialDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sequential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CDSCabinetServer).SequentialDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cds.cabinet.v1.CDSCabinet/SequentialDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CDSCabinetServer).SequentialDelete(ctx, req.(*Sequential))
	}
	return interceptor(ctx, in, info, handler)
}

func _CDSCabinet_SequentialGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sequential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CDSCabinetServer).SequentialGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cds.cabinet.v1.CDSCabinet/SequentialGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CDSCabinetServer).SequentialGet(ctx, req.(*Sequential))
	}
	return interceptor(ctx, in, info, handler)
}

func _CDSCabinet_SequentialList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SequentialListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CDSCabinetServer).SequentialList(m, &cDSCabinetSequentialListServer{stream})
}

type CDSCabinet_SequentialListServer interface {
	Send(*Sequential) error
	grpc.ServerStream
}

type cDSCabinetSequentialListServer struct {
	grpc.ServerStream
}

func (x *cDSCabinetSequentialListServer) Send(m *Sequential) error {
	return x.ServerStream.SendMsg(m)
}

var _CDSCabinet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cds.cabinet.v1.CDSCabinet",
	HandlerType: (*CDSCabinetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CounterGet",
			Handler:    _CDSCabinet_CounterGet_Handler,
		},
		{
			MethodName: "EdgeGet",
			Handler:    _CDSCabinet_EdgeGet_Handler,
		},
		{
			MethodName: "IndexGet",
			Handler:    _CDSCabinet_IndexGet_Handler,
		},
		{
			MethodName: "MetaGet",
			Handler:    _CDSCabinet_MetaGet_Handler,
		},
		{
			MethodName: "NodeGet",
			Handler:    _CDSCabinet_NodeGet_Handler,
		},
		{
			MethodName: "ReadCheck",
			Handler:    _CDSCabinet_ReadCheck_Handler,
		},
		{
			MethodName: "SequentialCreate",
			Handler:    _CDSCabinet_SequentialCreate_Handler,
		},
		{
			MethodName: "SequentialDelete",
			Handler:    _CDSCabinet_SequentialDelete_Handler,
		},
		{
			MethodName: "SequentialGet",
			Handler:    _CDSCabinet_SequentialGet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EdgeList",
			Handler:       _CDSCabinet_EdgeList_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "IndexList",
			Handler:       _CDSCabinet_IndexList_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "MetaList",
			Handler:       _CDSCabinet_MetaList_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "NodeList",
			Handler:       _CDSCabinet_NodeList_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Transaction",
			Handler:       _CDSCabinet_Transaction_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SequentialList",
			Handler:       _CDSCabinet_SequentialList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cabinet.proto",
}

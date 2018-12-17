// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package server

import (
	"cds.ikigai.net/cabinet.v1/iri"
	"cds.ikigai.net/cabinet.v1/perms"
	pb "cds.ikigai.net/cabinet.v1/rpc"
	"fmt"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *CDSCabinetServer) MetaList(metaRq *pb.MetaListRequest, stream pb.CDSCabinet_MetaListServer) error{
	_, err := s.fdb.ReadTransact(func (rtr fdb.ReadTransaction) (interface{}, error) {
		if DebugServerRequests {
			s.logEvent(fmt.Sprintf("MetaList(%v)", metaRq))
		}

		metaPerms := &perms.Meta{
			AllowWildcardProperty: true,
		}

		metaIRI, err := iri.ResolveMetaIRI(metaRq.Meta, nil, metaPerms)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, err)
		}

		isEdgeMeta := false
		switch iOpt := metaRq.Meta.Object.(type) {
			case *pb.Meta_Edge:
				isEdgeMeta = true
			case *pb.Meta_Node:
				isEdgeMeta = false
			default:
				return nil, status.Errorf(codes.InvalidArgument, RPCErrorIRISpecific, fmt.Sprintf("%v is not valid Meta Object", iOpt))
		}

		ri := metaIRI.GetListRange(s.dbMeta, rtr, metaRq.Opt).Iterator()

		for ri.Advance() {
			kv := ri.MustGet()

			// [e, SUBJECT, PREDICATE, TARGET, PROP] = bin
			// [n, NODE_ID, PROP] = bin
			metaKeys, err := s.dbMeta.Unpack(kv.Key)

			if err != nil {
				return nil, status.Errorf(codes.DataLoss, RPCErrorDataCorrupted, "meta_key")
			}

			obj := &pb.Meta{}
			propKey := 2 // 2 on NODE, 4 on EDGE

			if !isEdgeMeta{
				if metaRq.IncludeNode {
					obj.Object = &pb.Meta_Node{Node: metaKeys[1].(string)}
				}

			}else{
				propKey = 4
				mEdge := &pb.Edge{}

				if metaRq.IncludeSubject {
					mEdge.Subject = metaKeys[1].(string)
				}

				if metaRq.IncludePredicate {
					propType, err := iri.KeyElementToInt(metaKeys[2].(string))
					if err != nil{
						return nil, status.Errorf(codes.DataLoss, RPCErrorDataCorrupted, "meta.edge.predicate")
					}

					mEdge.Predicate = uint32(propType)
				}

				if metaRq.IncludeTarget {
					mEdge.Target = metaKeys[3].(string)
				}

				obj.Object = &pb.Meta_Edge{Edge: mEdge}
			}

			if metaRq.IncludeProperty {
				propType, err := iri.KeyElementToInt(metaKeys[propKey].(string))
				if err != nil{
					return nil, status.Errorf(codes.DataLoss, RPCErrorDataCorrupted, "meta.key")
				}

				obj.Key = uint32(propType)
			}

			if metaRq.IncludeValue{
				obj.Val = kv.Value
			}

			if err := stream.Send(obj); err != nil {
				return nil, err
			}
		}

		return nil, nil
	})

	return err
}

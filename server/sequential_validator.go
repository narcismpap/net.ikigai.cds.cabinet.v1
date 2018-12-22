// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package server

import (
	pb "cds.ikigai.net/cabinet.v1/rpc"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validateSequentialRequest(seq *pb.Sequential, required []string, unexpected []string) error {
	for i := range required {
		if required[i] == "t" && len(seq.GetType()) == 0 {
			return status.Error(codes.InvalidArgument, fmt.Sprintf(RPCErrorFieldRequired, "seq.type"))
		} else if required[i] == "u" && len(seq.GetUuid()) == 0 {
			return status.Error(codes.InvalidArgument, fmt.Sprintf(RPCErrorFieldRequired, "seq.uuid"))
		} else if required[i] == "s" && seq.GetSeqid() == 0 {
			return status.Error(codes.InvalidArgument, fmt.Sprintf(RPCErrorFieldRequired, "seq.seqId"))
		}
	}

	for i := range unexpected {
		if unexpected[i] == "t" && len(seq.GetType()) > 0 {
			return status.Error(codes.InvalidArgument, fmt.Sprintf(RPCErrorFieldUnexpected, "seq.type"))
		} else if unexpected[i] == "u" && len(seq.GetUuid()) > 0 {
			return status.Error(codes.InvalidArgument, fmt.Sprintf(RPCErrorFieldUnexpected, "seq.uuid"))
		} else if unexpected[i] == "s" && seq.GetSeqid() != 0 {
			return status.Error(codes.InvalidArgument, fmt.Sprintf(RPCErrorFieldUnexpected, "seq.seqId"))
		}
	}

	return nil
}

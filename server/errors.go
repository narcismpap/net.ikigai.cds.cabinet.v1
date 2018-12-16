// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package server

const(
	DebugServerRequests = true
)

// RPC Errors
// Returned by gRPC, standardized on E(0xYYY)
const (
	//RPCError
	RPCErrorInvalidAction 	= "E(0x001) Unknown TransactionAction"
	RPCErrorRepeatAction    = "E(0x002) Repeat actionId in transaction"
	RPCErrorArgumentInvalid = "E(0x003) One or more arguments are invalid (generic)"
	RPCErrorInvalidIRI      = "E(0x004) Invalid IRI (generic)"
	RPCErrorNotFound        = "E(0x005) Requested record not found"
	RPCErrorFieldRequired   = "E(0x006) Field %s is required"
	RPCErrorFieldUnexpected = "E(0x007) Field %s is unexpected"
	RPCErrorFieldSpecific   = "E(0x008) Error %s on %s"
	RPCErrorDataCorrupted   = "E(0x009) Data is corrupted on %s"
	RPCErrorDuplicateRecord = "E(0x010) Requested record already exists"
	RPCErrorIRISpecific     = "E(0x011) IRI: %s"
)

func CheckFatalError(err error){
	if err != nil{
		panic(err)
	}
}

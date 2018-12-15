// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package main

import (
	"fmt"
	"log"
)

const (
	CDSErrFieldInvalid = iota
	CDSErrorFieldRequired
	CDSErrorFieldUnexpected

	CDSErrorNotFound
	CDSErrorPermission
	CDSErrorBadRecord

	CDSListNoPagination
)

const (
	//RPCError
	RPCErrorInvalidAction 	= "E(0x001) Unknown TransactionAction"
	RPCErrorRepeatAction	= "E(0x002) Repeat actionId in transaction"
	RPCErrorArgumentInvalid	= "E(0x003) One or more arguments are invalid"
	RPCErrorInvalidIRI 		= "E(0x004) Invalid IRI"
)

type CabinetError struct {
	err  	string
	code 	int16
	field 	string
}

func (e *CabinetError) Error() string {
	return fmt.Sprintf("Cabinet Error #%d on %s: %s", e.code, e.field, e.err)
}

func CheckFatalError(err error){
	if err != nil{
		log.Fatal(err)
		panic(err)
	}
}

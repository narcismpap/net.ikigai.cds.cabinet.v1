// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package main

import "fmt"

const (
	CDSErrFieldInvalid = iota
	CDSErrorFieldRequired
	CDSErrorFieldUnexpected

	CDSErrorNotFound
	CDSErrorPermission
	CDSErrorBadRecord

	CDSListNoPagination
)

type CabinetError struct {
	err  	string
	code 	int16
	field 	string
}

func (e *CabinetError) Error() string {
	return fmt.Sprintf("Cabinet Error #%d on %s: %s", e.code, e.field, e.err)
}


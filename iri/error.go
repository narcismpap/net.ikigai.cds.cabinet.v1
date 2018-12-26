// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package iri

import "fmt"

type ParsingError struct {
	msg   string
	field string
}

func (e *ParsingError) Error() string {
	return fmt.Sprintf("%s on %s", e.msg, e.field)
}

func (e *ParsingError) Message() string {
	return e.msg
}

func (e *ParsingError) Field() string {
	return e.field
}

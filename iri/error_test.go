// Package: net.ikigai.cds
// Module: cabinet.services.test
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package iri_test

import (
	"cds.ikigai.net/cabinet.v1/iri"
	"testing"
	"valencex.com/dev/testutil"
)

func TestIRIError(t *testing.T) {
	t.Parallel()
	r := testutil.NewTestRunner(t)

	er := iri.NewParsingError("null record", "john.doe")

	r.IsEqualString(er.Field(), "john.doe", "ParsingError.Field()")
	r.IsEqualString(er.Message(), "null record", "ParsingError.Message()")
	r.IsEqualString(er.Error(), "null record on john.doe", "ParsingError.Error()")
}

// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package iri_test

import (
	"cds.ikigai.net/cabinet.v1/iri"
	"fmt"
	"testing"
	"valencex.com/dev/testutil"
)

func TestValidateNodeId(t *testing.T) {
	t.Parallel()
	r := testutil.NewTestRunner(t)

	valid := []string{
		"1EshIpp5qyaXuSoDNLNzbu8QNFM", "1EshIrwqx0XbyFB5bMOoIRkpmyf", "1EshIp7Thwi9McdjBcwWQuw9pIo",
	}

	invalid := []string{
		"lorem", "1AF87096-FDC1-4709-88D0-2BAD5EC81AA9", "1EshIrwqx0XbyFB5bMOoIRkpm", "{cats}", "",
	}

	for v1 := range valid {
		_, err1 := iri.ValidateNodeId(valid[v1])
		r.AssertNilError(err1, fmt.Sprintf("ValidateNodeId(%s)", valid[v1]))
	}

	for v2 := range invalid {
		_, err2 := iri.ValidateNodeId(invalid[v2])
		r.AssertErrorPresent(err2, fmt.Sprintf("ValidateNodeId(%s)", invalid[v2]))
	}
}

func TestValidateUuid(t *testing.T) {
	t.Parallel()
	r := testutil.NewTestRunner(t)

	valid := []string{
		"1AF87096-FDC1-4709-88D0-2BAD5EC81AA9", "1AF87096FDC1470988D02BAD5EC81AA9",
	}

	invalid := []string{
		"1EshIpp5qyaXuSoDNLNzbu8QNFM", "", "1EshIrwqx0XbyFB5bMOoIRkpm", "{cats}", "",
		"1AF87096-FDC1-X709-88D0-2BAD5EC81AA9", // 4 replaced by X, illegal V4 type
	}

	for v1 := range valid {
		_, err1 := iri.ValidateUuid(valid[v1])
		r.AssertNilError(err1, fmt.Sprintf("ValidateUuid(%s)", valid[v1]))
	}

	for v2 := range invalid {
		_, err2 := iri.ValidateUuid(invalid[v2])
		r.AssertErrorPresent(err2, fmt.Sprintf("ValidateUuid(%s)", invalid[v2]))
	}
}

func TestValidateSequence(t *testing.T) {
	t.Parallel()
	r := testutil.NewTestRunner(t)

	for _, v1 := range []uint16{10, 12, 65000, 1000} {
		r.AssertTrue(iri.ValidateSequence(v1), fmt.Sprintf("ValidateSequence(%d)", v1))
	}

	for _, v2 := range []uint16{0} {
		r.AssertFalse(iri.ValidateSequence(v2), fmt.Sprintf("ValidateSequence(%d)", v2))
	}
}

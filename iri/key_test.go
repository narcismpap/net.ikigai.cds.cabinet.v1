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

func TestSequenceKeys(t *testing.T) {
	t.Parallel()
	r := testutil.NewTestRunner(t)

	valid := map[uint16][]byte{
		1: {0, 1},
		10: {0, 10},
		99: {0, 99},
		100: {0, 100},
		1000: {3, 232},
		5432: {21, 56},
		10000: {39, 16},
		25000: {97, 168},
		60000: {234, 96},
	}

	for num, knownBytes := range valid{
		bKey := iri.SequenceToSmallKey(num)
		r.AssertEqualBytes(bKey, knownBytes, fmt.Sprintf("SequenceToSmallKey(%d)", num))

		bRevKey, err := iri.SmallKeyToSequence(bKey)
		r.AssertNilError(err, fmt.Sprintf("SmallKeyToSequence(%v)", bKey))
		r.AssertEqualUInt16(bRevKey, num, "SequenceToSmallKey->SmallKeyToSequence")
	}
}


func TestCoreSequence16(t *testing.T) {
	t.Parallel()
	r := testutil.NewTestRunner(t)

	valid := map[string]uint16{
		"1": 1,
		"423": 423,
		"952": 952,
		"10413": 10413,
		"29426": 29426,
		"49134": 49134,
		"62013": 62013,
	}

	for vStr, vInt := range valid{
		ul, err := iri.ParseCoreSequence(vStr)

		r.AssertNilError(err, fmt.Sprintf("ParseCoreSequence(%v)", vStr))
		r.AssertEqualUInt16(ul, vInt, fmt.Sprintf("ParseCoreSequence(%v)", vStr))
	}
}

func TestCoreSequence32(t *testing.T) {
	t.Parallel()
	r := testutil.NewTestRunner(t)

	valid := map[string]uint32{
		"1": 1,
		"423": 423,
		"952": 952,
		"10413": 10413,
		"29426": 29426,
		"49134": 49134,
		"62013": 62013,
	}

	for vStr, vInt := range valid{
		ul, err := iri.ParseCoreSequence32(vStr)

		r.AssertNilError(err, fmt.Sprintf("ParseCoreSequence(%v)", vStr))
		r.AssertEqualUInt32(ul, vInt, fmt.Sprintf("ParseCoreSequence(%v)", vStr))
	}
}

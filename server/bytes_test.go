// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package server_test

import (
	"cds.ikigai.net/cabinet.v1/server"
	"fmt"
	"testing"
	"valencex.com/dev/testutil"
)

func TestInt64BytesConversion(t *testing.T) {
	t.Parallel()
	r := testutil.NewTestRunner(t)

	valid := map[int64][]byte{
		-1000000: {192, 189, 240, 255, 255, 255, 255, 255},
		-966424:  {232, 64, 241, 255, 255, 255, 255, 255},
		-65436:   {100, 0, 255, 255, 255, 255, 255, 255},
		-6545:    {111, 230, 255, 255, 255, 255, 255, 255},
		-654:     {114, 253, 255, 255, 255, 255, 255, 255},
		-6:       {250, 255, 255, 255, 255, 255, 255, 255},
		134:      {134, 0, 0, 0, 0, 0, 0, 0},
		100:      {100, 0, 0, 0, 0, 0, 0, 0},
		1000:     {232, 3, 0, 0, 0, 0, 0, 0},
		1765:     {229, 6, 0, 0, 0, 0, 0, 0},
		5432:     {56, 21, 0, 0, 0, 0, 0, 0},
		25765:    {165, 100, 0, 0, 0, 0, 0, 0},
		60654:    {238, 236, 0, 0, 0, 0, 0, 0},
		876543:   {255, 95, 13, 0, 0, 0, 0, 0},
		6543543:  {183, 216, 99, 0, 0, 0, 0, 0},
	}

	for num, knownBytes := range valid {
		bKey, err := server.Int64ToBytes(num)

		r.AssertNilError(err, fmt.Sprintf("Int64ToBytes(%v)", bKey))
		r.AssertEqualBytes(bKey, knownBytes, fmt.Sprintf("Int64ToBytes(%d)", num))

		bRevKey, err := server.BytesToInt(bKey)
		r.AssertNilError(err, fmt.Sprintf("BytesToInt(%v)", bKey))
		r.AssertEqualInt64(bRevKey, num, "Int64ToBytes->BytesToInt")
	}
}

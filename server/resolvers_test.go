// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package server_test

import (
	"cds.ikigai.net/cabinet.v1/server"
	"testing"
	"valencex.com/dev/testutil"
)

func TestNodeResolveId(t *testing.T) {
	t.Parallel()
	r := testutil.NewTestRunner(t)

	mp := map[string]string{"1": "1Et4vvPvqDfhGyGhpzaizwSEYU0", "2": "1Et4vwss75S9KYuZw2nXgVKiLQT"}
	em := make(map[string]string)

	mapTests := []struct {
		in  string
		out string
	}{
		{"test", "test"},
		{"1Et4vymFYgDHCKeQJpTAbJvqEtQ", "1Et4vymFYgDHCKeQJpTAbJvqEtQ"},
		{"tmp:1", "1Et4vvPvqDfhGyGhpzaizwSEYU0"},
		{"tmp:2", "1Et4vwss75S9KYuZw2nXgVKiLQT"},
	}

	for t := range mapTests {
		vl, err := server.NodeResolveId(mapTests[t].in, &mp)

		r.AssertNilError(err, "NodeResolveId")
		r.AssertEqualString(vl, mapTests[t].out, "NodeResolveId")
	}

	_, err := server.NodeResolveId("tmp:3", &mp)
	r.AssertErrorPresent(err, "NodeResolveId")

	_, err2 := server.NodeResolveId("tmp:5", &em)
	r.AssertErrorPresent(err2, "NodeResolveId")
}

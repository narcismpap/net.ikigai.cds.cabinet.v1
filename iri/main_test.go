// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package iri_test

import (
	"fmt"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"testing"
	"valencex.com/dev/testutil"
)

type ExpectedBytes struct {
	key   []byte
	start []byte
	end   []byte
}

func (s *ExpectedBytes) getKey(sub []byte) []byte{
	return append(sub, s.key...)
}

func (s *ExpectedBytes) getStart(sub []byte) []byte{
	return append(sub, s.start...)
}

func (s *ExpectedBytes) GetEnd(sub []byte) []byte{
	return append(sub, s.end...)
}

type IRITesterDb struct {
	Db        fdb.Database
	Container directory.DirectorySubspace

	DbNode     subspace.Subspace
	DbEdge     subspace.Subspace
	DbIndex    subspace.Subspace
	DbIndexCnt subspace.Subspace
	DbMeta     subspace.Subspace
	DbCount    subspace.Subspace
	DbSequence subspace.Subspace

	DbNodeBytes     []byte
	DbEdgeBytes     []byte
	DbIndexBytes    []byte
	DbIndexCntBytes []byte
	DbMetaBytes     []byte
	DbCountBytes    []byte
	DbSequenceBytes []byte
}

var testDb IRITesterDb

func init() {
	var err error

	fdb.MustAPIVersion(600)

	testDb.Db = fdb.MustOpenDefault()
	testDb.Container, err = directory.CreateOrOpen(testDb.Db, []string{"test"}, nil)

	if err != nil {
		panic(err)
	}

	testDb.DbNode = testDb.Container.Sub("n")
	testDb.DbEdge = testDb.Container.Sub("e")
	testDb.DbIndex = testDb.Container.Sub("i")
	testDb.DbIndexCnt = testDb.Container.Sub("k")
	testDb.DbMeta = testDb.Container.Sub("m")
	testDb.DbCount = testDb.Container.Sub("c")
	testDb.DbSequence = testDb.Container.Sub("s")

	testDb.DbNodeBytes = testDb.DbNode.Bytes()
	testDb.DbEdgeBytes = testDb.DbEdge.Bytes()
	testDb.DbIndexBytes = testDb.DbIndex.Bytes()
	testDb.DbIndexCntBytes = testDb.DbIndexCnt.Bytes()
	testDb.DbMetaBytes = testDb.DbMeta.Bytes()
	testDb.DbCountBytes = testDb.DbCount.Bytes()
	testDb.DbSequenceBytes = testDb.DbSequence.Bytes()
}

type IRITester struct {
	r *testutil.TestRunner
}

func NewIRITester(t *testing.T) *IRITester {
	ir := &IRITester{r: testutil.NewTestRunner(t)}
	t.Parallel()

	return ir
}

func (i *IRITester) path(a string, b string) {
	i.r.AssertEqualString(a, b, "IRITester.path()")
}

func (i *IRITester) key(a string, b string, name string) {
	i.r.AssertEqualString(a, b, fmt.Sprintf("IRITester.key(%s)", name))
}

func (i *IRITester) seqKey(a uint16, b uint16, name string) {
	i.r.AssertEqualUInt16(a, b, fmt.Sprintf("IRITester.seqKey(%s)", name))
}

func (i *IRITester) bytes(a []byte, expect []byte, name string) {
	i.r.AssertEqualBytes(a, expect, fmt.Sprintf("IRITester.bytes(%s)", name))
}

func (i *IRITester) seqKey32(a uint32, b uint32, name string) {
	i.r.AssertEqualUInt32(a, b, fmt.Sprintf("IRITester.seqKey32(%s)", name))
}

func (i *IRITester) error(err error, expected string) {
	i.r.AssertErrorExact(err, expected, "IRITester.error()")
}

func (i *IRITester) nil(err error) {
	i.r.AssertNilError(err, "IRITester.nil()")
}

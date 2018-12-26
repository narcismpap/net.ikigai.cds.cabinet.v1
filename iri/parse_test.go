// Package: net.ikigai.cds
// Module: cabinet.services.test
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package iri_test

import (
	"cds.ikigai.net/cabinet.v1/iri"
	"cds.ikigai.net/cabinet.v1/perms"
	"testing"
	"valencex.com/dev/testutil"
)

func TestParseSequenceId(t *testing.T) {
	t.Parallel()
	r := testutil.NewTestRunner(t)

	seqId, err := iri.Parse("s/test/66")
	r.IsNoError(err, "iri.Parse")

	seq := seqId.(*iri.Sequence)

	r.IsEqualString(seq.Type, "test", "seq.type")
	r.IsEqualUInt32(seq.SeqID, 66, "seq.SeqId")
	r.IsEqualString(seq.UUID, "", "seq.UUID")
	r.IsEqualString(seq.GetPath(), "s/test/66", "seq.GetPath()")

	p := &perms.Sequence{}
	r.IsNoError(seq.ValidateIRI(p), "seq.ValidateIRI")
}

func TestParseSequenceUuid(t *testing.T) {
	t.Parallel()
	r := testutil.NewTestRunner(t)

	seqId, err := iri.Parse("su/37A6ACAA-C704-45CD-9F89-8708B57AAD96")
	r.IsNoError(err, "iri.Parse")

	seq := seqId.(*iri.Sequence)

	r.IsEqualString(seq.Type, "", "seq.type")
	r.IsEqualUInt32(seq.SeqID, 0, "seq.SeqId")
	r.IsEqualString(seq.UUID, "37A6ACAA-C704-45CD-9F89-8708B57AAD96", "seq.UUID")
	r.IsEqualString(seq.GetPath(), "su/37A6ACAA-C704-45CD-9F89-8708B57AAD96", "seq.GetPath()")

	p := &perms.Sequence{}
	r.IsNoError(seq.ValidateIRI(p), "seq.ValidateIRI")
}

func TestParseEdge(t *testing.T) {
	t.Parallel()
	r := testutil.NewTestRunner(t)

	edgeRaw, err := iri.Parse("e/1Esvhs92tk27vcYQ0L12fORh6Jh/134/1EsvhtLglKL5jUKA0fdgz6A7gfy")
	r.IsNoError(err, "iri.Parse")

	edge := edgeRaw.(*iri.Edge)

	r.IsEqualString(edge.Subject, "1Esvhs92tk27vcYQ0L12fORh6Jh", "edge.Subject")
	r.IsEqualUInt16(edge.Predicate, 134, "edge.Predicate")
	r.IsEqualString(edge.Target, "1EsvhtLglKL5jUKA0fdgz6A7gfy", "edge.Target")
	r.IsEqualString(edge.GetPath(), "e/1Esvhs92tk27vcYQ0L12fORh6Jh/134/1EsvhtLglKL5jUKA0fdgz6A7gfy", "edge.GetPath()")

	p := &perms.Edge{}
	r.IsNoError(edge.ValidateIRI(p), "edge.ValidateIRI")
}

func TestParseIndex(t *testing.T) {
	t.Parallel()
	r := testutil.NewTestRunner(t)

	indexRaw, err := iri.Parse("i/95/cat/1EsJ4OwOAdywg8iM3dnH2ODHfjq")
	r.IsNoError(err, "iri.Parse")

	index := indexRaw.(*iri.NodeIndex)

	r.IsEqualString(index.Node, "1EsJ4OwOAdywg8iM3dnH2ODHfjq", "index.Node")
	r.IsEqualUInt16(index.IndexId, 95, "index.IndexId")
	r.IsEqualString(index.Value, "cat", "index.Value")
	r.IsEqualString(index.GetPath(), "i/95/cat/1EsJ4OwOAdywg8iM3dnH2ODHfjq", "index.GetPath()")

	p := &perms.Index{}
	r.IsNoError(index.ValidateIRI(p), "edge.ValidateIRI")
}

func TestParseNode(t *testing.T) {
	t.Parallel()
	r := testutil.NewTestRunner(t)

	nodeRaw, err := iri.Parse("n/7654/1Esvwjo5ltn4CArkd7HyzkWIK80")
	r.IsNoError(err, "iri.Parse")

	node := nodeRaw.(*iri.Node)

	r.IsEqualString(node.Id, "1Esvwjo5ltn4CArkd7HyzkWIK80", "node.Id")
	r.IsEqualUInt16(node.Type, 7654, "node.Type")
	r.IsEqualString(node.GetPath(), "n/7654/1Esvwjo5ltn4CArkd7HyzkWIK80", "node.GetPath()")

	p := &perms.Node{}
	r.IsNoError(node.ValidateIRI(p), "node.ValidateIRI")
}

func TestParseMetaNode(t *testing.T) {
	t.Parallel()
	r := testutil.NewTestRunner(t)

	metaRaw, err := iri.Parse("m/n/1EshIpkPqdT02Q6MPkhw4lWPMJk/6524")
	r.IsNoError(err, "iri.Parse")

	meta := metaRaw.(*iri.NodeMeta)

	r.IsEqualString(meta.Node, "1EshIpkPqdT02Q6MPkhw4lWPMJk", "meta.Node")
	r.IsEqualUInt16(meta.Property, 6524, "meta.Property")
	r.IsEqualString(meta.GetPath(), "m/n/1EshIpkPqdT02Q6MPkhw4lWPMJk/6524", "meta.GetPath()")

	p := &perms.Meta{}
	r.IsNoError(meta.ValidateIRI(p), "meta.ValidateIRI")
}

func TestParseMetaEdge(t *testing.T) {
	t.Parallel()
	r := testutil.NewTestRunner(t)

	metaRaw, err := iri.Parse("m/e/1EshIqcEmk4HxnwpZSdSFfMqmat/56/1EshIntZeJN1ubDCbUYS5zA1noN/5244")
	r.IsNoError(err, "iri.Parse")

	edgeMeta := metaRaw.(*iri.EdgeMeta)

	r.IsEqualString(edgeMeta.Subject, "1EshIqcEmk4HxnwpZSdSFfMqmat", "meta.edge.Subject")
	r.IsEqualUInt16(edgeMeta.Property, 5244, "meta.edge.Property")
	r.IsEqualUInt16(edgeMeta.Predicate, 56, "meta.Predicate")
	r.IsEqualString(edgeMeta.Target, "1EshIntZeJN1ubDCbUYS5zA1noN", "meta.edge.Target")
	r.IsEqualString(edgeMeta.GetPath(), "m/e/1EshIqcEmk4HxnwpZSdSFfMqmat/56/1EshIntZeJN1ubDCbUYS5zA1noN/5244", "meta.GetPath()")

	p := &perms.Meta{}
	r.IsNoError(edgeMeta.ValidateIRI(p), "meta.ValidateIRI")
}

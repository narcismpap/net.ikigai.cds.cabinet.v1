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
)

func TestIRIMetaEdgeCompose(t *testing.T) {
	x := NewIRITester(t)

	c1 := &iri.EdgeMeta{Subject: "1EshIqQAfyMqUGM4kj3wgo7v7ZQ", Predicate: 12, Target: "1EshIpkEEtXQa4KVidCkA2nQPVd", Property: 243}
	x.path(c1.GetPath(), "m/e/1EshIqQAfyMqUGM4kj3wgo7v7ZQ/12/1EshIpkEEtXQa4KVidCkA2nQPVd/243")

	x.key(c1.Subject, "1EshIqQAfyMqUGM4kj3wgo7v7ZQ", "subject")
	x.key(c1.Target, "1EshIpkEEtXQa4KVidCkA2nQPVd", "target")
	x.seqKey(c1.Predicate, 12, "predicate")
	x.seqKey(c1.Property, 243, "property")
}

func TestIRIMetaEdgeParse(t *testing.T) {
	x := NewIRITester(t)

	c2 := &iri.EdgeMeta{}
	if err := c2.Parse("m/e/1EshIqcEmk4HxnwpZSdSFfMqmat/56/1EshIntZeJN1ubDCbUYS5zA1noN/5244"); err != nil{
		x.t.Log(err)
	}

	x.path(c2.GetPath(), "m/e/1EshIqcEmk4HxnwpZSdSFfMqmat/56/1EshIntZeJN1ubDCbUYS5zA1noN/5244")

	x.key(c2.Subject, "1EshIqcEmk4HxnwpZSdSFfMqmat", "subject")
	x.key(c2.Target, "1EshIntZeJN1ubDCbUYS5zA1noN", "target")
	x.seqKey(c2.Predicate, 56, "predicate")
	x.seqKey(c2.Property, 5244, "property")
}

func TestIRIMetaEdgeBadSignature(t *testing.T) {
	x := NewIRITester(t)
	p := &perms.Meta{}

	// empty
	x.error((&iri.EdgeMeta{

	}).ValidateIRI(p), "null record on meta.property")

	// no prop
	x.error((&iri.EdgeMeta{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 100, Target: "1EsJ4O4FlJXmKdNxLk52Go4x0uE",
	}).ValidateIRI(p), "null record on meta.property")

	// wildcard prop (not allowed)
	x.error((&iri.EdgeMeta{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 100, Target: "1EsJ4O4FlJXmKdNxLk52Go4x0uE", Property: 0,
	}).ValidateIRI(p), "null record on meta.property")

	// no subject + bad
	x.error((&iri.EdgeMeta{
		Predicate: 100, Target: "1EsJ4O4FlJXmKdNxLk52Go4x0uE", Property: 143,
	}).ValidateIRI(p), "invalid Node ID on meta.edge.subject")

	x.error((&iri.EdgeMeta{
		Subject: "i am a test", Predicate: 100, Target: "1EsJ4O4FlJXmKdNxLk52Go4x0uE", Property: 255,
	}).ValidateIRI(p), "invalid Node ID on meta.edge.subject")

	// no predicate + bad
	x.error((&iri.EdgeMeta{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Target: "1EsJ4O4FlJXmKdNxLk52Go4x0uE", Property: 5243,
	}).ValidateIRI(p), "null record on meta.edge.predicate")

	x.error((&iri.EdgeMeta{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 0, Target: "1EsJ4O4FlJXmKdNxLk52Go4x0uE", Property: 524,
	}).ValidateIRI(p), "null record on meta.edge.predicate")

	// no target + bad
	x.error((&iri.EdgeMeta{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 100, Property: 24,
	}).ValidateIRI(p), "invalid Node ID on meta.edge.target")

	x.error((&iri.EdgeMeta{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 100, Property: 876, Target: "I like cats",
	}).ValidateIRI(p), "invalid Node ID on meta.edge.target")
}

func TestIRIMetaEdgeBadSignatureWildcard(t *testing.T) {
	x := NewIRITester(t)

	p1 := &perms.Meta{AllowWildcardProperty:false}
	p2 := &perms.Meta{AllowWildcardProperty:true}

	x.error((&iri.EdgeMeta{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 100, Target: "1EsJ4O4FlJXmKdNxLk52Go4x0uE", Property: 0,
	}).ValidateIRI(p1), "null record on meta.property")

	x.nil((&iri.EdgeMeta{
		Subject: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Predicate: 100, Target: "1EsJ4O4FlJXmKdNxLk52Go4x0uE", Property: 0,
	}).ValidateIRI(p2))
}

func TestIRIMetaNodeCompose(t *testing.T) {
	x := NewIRITester(t)

	c3 := &iri.NodeMeta{Node: "1EshIpkPqdT02Q6MPkhw4lWPMJk", Property: 524}
	x.path(c3.GetPath(), "m/n/1EshIpkPqdT02Q6MPkhw4lWPMJk/524")

	x.key(c3.Node, "1EshIpkPqdT02Q6MPkhw4lWPMJk", "node")
	x.seqKey(c3.Property, 524, "property")
}

func TestIRIMetaNodeParse(t *testing.T) {
	x := NewIRITester(t)

	c4 := &iri.NodeMeta{}
	if err := c4.Parse("m/n/1EshImtww6zGdxg1csBYGjgqo2T/524"); err != nil{
		x.t.Log(err)
	}

	x.path(c4.GetPath(), "m/n/1EshImtww6zGdxg1csBYGjgqo2T/524")

	x.key(c4.Node, "1EshImtww6zGdxg1csBYGjgqo2T", "node")
	x.seqKey(c4.Property, 524, "property")
}

func TestIRIMetaNodeBadSignature(t *testing.T) {
	x := NewIRITester(t)
	p := &perms.Meta{}

	// empty
	x.error((&iri.NodeMeta{

	}).ValidateIRI(p), "null record on meta.property")

	// no property
	x.error((&iri.NodeMeta{
		Node: "1EsJ4OwOAdywg8iM3dnH2ODHfjq",
	}).ValidateIRI(p), "null record on meta.property")

	// wildcard property (not allowed)
	x.error((&iri.NodeMeta{
		Node: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Property: 0,
	}).ValidateIRI(p), "null record on meta.property")

	// no valid ID
	x.error((&iri.NodeMeta{
		Node: "not good", Property: 100,
	}).ValidateIRI(p), "invalid Node ID on meta.node")

	// invalid size
	x.error((&iri.NodeMeta{
		Node: "1EsJ4OwOAdywg8iM3dnH2ODHfjqQ", Property: 100,
	}).ValidateIRI(p), "invalid Node ID on meta.node")

	// UUID
	x.error((&iri.NodeMeta{
		Node: "151E0C62-1ADD-4BA0-BFC5-840E9370592F", Property: 100,
	}).ValidateIRI(p), "invalid Node ID on meta.node")
}

func TestIRIMetaNodeBadSignatureWildcard(t *testing.T) {
	x := NewIRITester(t)

	p1 := &perms.Meta{AllowWildcardProperty:false}
	p2 := &perms.Meta{AllowWildcardProperty:true}

	x.error((&iri.NodeMeta{
		Node: "1EsJ4OwOAdywg8iM3dnH2ODHfjq", Property: 0,
	}).ValidateIRI(p1), "null record on meta.property")

	x.nil((&iri.NodeMeta{
		Node: "1EshIohhwTq3pya0tLVNCST07gN", Property: 0,
	}).ValidateIRI(p2))
}

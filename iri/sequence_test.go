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

var tSeqIdBytes ExpectedBytes
var tSeqUuidBytes ExpectedBytes

func init() {
	// IRI: s/q/66
	tSeqIdBytes.key = []byte{2, 113, 0, 2, 105, 0, 2, 48, 48, 48, 54, 54, 0}
	tSeqIdBytes.start = []byte{}
	tSeqIdBytes.end = []byte{}

	// IRI: su/19AEA360-E595-4B15-83A2-2CAB7FE31767
	tSeqUuidBytes.key = []byte{2, 115, 117, 0, 2, 49, 57, 65, 69, 65, 51, 54, 48, 45, 69, 53, 57, 53, 45, 52, 66, 49, 53, 45, 56, 51, 65, 50, 45, 50, 67, 65, 66, 55, 70, 69, 51, 49, 55, 54, 55, 0}
	tSeqUuidBytes.start = []byte{}
	tSeqUuidBytes.end = []byte{}
}

func TestIRISequenceIdCompose(t *testing.T) {
	x := NewIRITester(t)

	s1 := &iri.Sequence{Type: "q", SeqID: 66}
	x.path(s1.GetPath(), "s/q/66")

	x.key(s1.Type, "q", "type")
	x.key(s1.UUID, "", "uuid")
	x.seqKey32(s1.SeqID, 66, "seqId")

	x.bytes([]byte(s1.GetKey(testDb.DbSequence)), tSeqIdBytes.getKey(testDb.DbSequenceBytes), "GetKey()")
}

func TestIRISequenceIdParse(t *testing.T) {
	x := NewIRITester(t)

	s2 := &iri.Sequence{}
	if err := s2.Parse("s/q/66"); err != nil {
		t.Error(err)
	}

	x.path(s2.GetPath(), "s/q/66")

	x.key(s2.Type, "q", "type")
	x.key(s2.UUID, "", "uuid")
	x.seqKey32(s2.SeqID, 66, "seqId")

	x.bytes([]byte(s2.GetKey(testDb.DbSequence)), tSeqIdBytes.getKey(testDb.DbSequenceBytes), "GetKey()")
}

func TestIRISequenceUuidCompose(t *testing.T) {
	x := NewIRITester(t)

	s3 := &iri.Sequence{UUID: "19AEA360-E595-4B15-83A2-2CAB7FE31767"}
	x.path(s3.GetPath(), "su/19AEA360-E595-4B15-83A2-2CAB7FE31767")

	x.key(s3.Type, "", "type")
	x.key(s3.UUID, "19AEA360-E595-4B15-83A2-2CAB7FE31767", "uuid")
	x.seqKey32(s3.SeqID, 0, "seqId")

	x.bytes([]byte(s3.GetReverseKey(testDb.DbSequence)), tSeqUuidBytes.getKey(testDb.DbSequenceBytes), "GetReverseKey()")
}

func TestIRISequenceUuidParse(t *testing.T) {
	x := NewIRITester(t)

	s4 := &iri.Sequence{}
	if err := s4.Parse("su/19AEA360-E595-4B15-83A2-2CAB7FE31767"); err != nil {
		t.Error(err)
	}

	x.path(s4.GetPath(), "su/19AEA360-E595-4B15-83A2-2CAB7FE31767")

	x.key(s4.Type, "", "type")
	x.key(s4.UUID, "19AEA360-E595-4B15-83A2-2CAB7FE31767", "uuid")
	x.seqKey32(s4.SeqID, 0, "seqId")

	x.bytes([]byte(s4.GetReverseKey(testDb.DbSequence)), tSeqUuidBytes.getKey(testDb.DbSequenceBytes), "GetReverseKey()")
}

func TestIRISequenceBadSignature(t *testing.T) {
	x := NewIRITester(t)
	p := &perms.Sequence{}

	// empty
	x.error((&iri.Sequence{}).ValidateIRI(p), "id required on seq.uuid|seq.id")

	// bad type, missing & over-sized
	x.error((&iri.Sequence{
		SeqID: 100,
	}).ValidateIRI(p), "null record on seq.type")

	x.error((&iri.Sequence{
		SeqID: 100, Type: "",
	}).ValidateIRI(p), "null record on seq.type")

	x.error((&iri.Sequence{
		SeqID: 100, Type: "cats are quite fun",
	}).ValidateIRI(p), "len > 10 on seq.type")

	// seqid & uuid cannot co-exist
	x.error((&iri.Sequence{
		SeqID: 100, UUID: "5D5D892A-C4EA-40F9-9724-5B1CD5A68057",
	}).ValidateIRI(p), "mutually exclusive on seq.uuid,seq.id")

	// missing ID & UUID
	x.error((&iri.Sequence{
		Type: "test",
	}).ValidateIRI(p), "id required on seq.uuid|seq.id")

	// bad UUID
	x.error((&iri.Sequence{
		Type: "test", UUID: "cats are fun",
	}).ValidateIRI(p), "invalid UUID on seq.uuid")
}

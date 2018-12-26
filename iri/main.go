// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package iri

import (
	"encoding/binary"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"strconv"
)

type IRI interface {
	GetPath() string
	GetKey(db subspace.Subspace) fdb.Key
	GetClearRange(db subspace.Subspace) fdb.ExactRange
	GetListRange(db subspace.Subspace, rtr fdb.ReadTransaction, opt *ListOptions) fdb.RangeResult
}

type ListOptions struct{
	PageSize int
	Reverse bool
}

func SequenceToSmallKey(v uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, v)
	return b
}

func SmallKeyToSequence(k []byte) (uint16, error) {
	return binary.BigEndian.Uint16(k), nil
}

func SequenceToSortableKey(v uint16) string {
	return strconv.FormatUint(uint64(v), 36)
}

func SortableKeyToSequence(k string) (uint16, error) {
	v, e := strconv.ParseUint(k, 36, 32)

	if e != nil {
		return 0, e
	}

	return uint16(v), nil
}

func ParseCoreSequence(k string) (uint16, error) {
	v, e := strconv.ParseUint(k, 10, 32)

	if e != nil {
		return 0, e
	}

	return uint16(v), nil
}

func ParseCoreSequence32(k string) (uint32, error) {
	v, e := strconv.ParseUint(k, 10, 32)

	if e != nil {
		return 0, e
	}

	return uint32(v), nil
}
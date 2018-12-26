// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package iri

import (
	"encoding/binary"
	"strconv"
)

func SequenceToSmallKey(v uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, v)
	return b
}

func SmallKeyToSequence(k []byte) (uint16, error) {
	return binary.BigEndian.Uint16(k), nil
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

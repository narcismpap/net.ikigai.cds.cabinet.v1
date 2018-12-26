// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package iri

import (
	"github.com/gofrs/uuid"
	"github.com/segmentio/ksuid"
)

func ValidateNodeId(id string) (ksuid.KSUID, error) {
	return ksuid.Parse(id)
}

func ValidateUuid(sUuid string) (uuid.UUID, error) {
	return uuid.FromString(sUuid)
}

func ValidateSequence(seq uint16) bool {
	return seq > 0
}

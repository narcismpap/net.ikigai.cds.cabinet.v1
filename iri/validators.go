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

func validateNodeID(id string) (ksuid.KSUID, error){
	return ksuid.Parse(id)
}

func validateUUID(sUUID string) (uuid.UUID, error){
	return uuid.FromString(sUUID)
}

func validateSequence(seq uint16) bool{
	return seq > 0 && seq < 65535
}

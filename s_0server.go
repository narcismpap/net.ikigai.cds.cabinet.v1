// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package main

import (
	"fmt"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/directory"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/apple/foundationdb/bindings/go/src/fdb/tuple"
	"log"
	"strconv"
)

type CDSCabinetServer struct{
	version int32

	fDb fdb.Transactor
	dbContainer directory.DirectorySubspace

	dbNode subspace.Subspace
	dbEdge subspace.Subspace
	dbIndex subspace.Subspace
	dbMeta subspace.Subspace
	dbCnt subspace.Subspace
	dbSeq subspace.Subspace
}

func (s *CDSCabinetServer) logEvent(e string){
	fmt.Println(e)
}


func newCDSServer() *CDSCabinetServer {
	fdb.MustAPIVersion(600)
	db := fdb.MustOpenDefault()

	var activeContainer = "test"
	container, err := directory.CreateOrOpen(db, []string{activeContainer}, nil)

	if err != nil {
		log.Fatal(err)
	}

	s := &CDSCabinetServer{
		version: 1,
		fDb: db,
		dbContainer: container,

		dbNode: container.Sub("n"),
		dbEdge: container.Sub("e"),
		dbIndex: container.Sub("i"),
		dbMeta: container.Sub("m"),
		dbCnt: container.Sub("c"),
		dbSeq: container.Sub("s"),
	}

	// install db
	var coreSeq = []string{
		"n", "e", "i", "m", "c",
	}

	_, err = s.fDb.Transact(func (tr fdb.Transaction) (interface{}, error) {
		tr.ClearRange(s.dbContainer)

		for i := range coreSeq {
			tr.Set(s.dbSeq.Pack(tuple.Tuple{coreSeq[i], "l"}), []byte(strconv.FormatUint(uint64(1), 10)))
		}

		log.Printf("[I] Container [%s] is now initialized", activeContainer)
		return nil, nil
	})

	return s
}

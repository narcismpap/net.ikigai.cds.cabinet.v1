// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package server

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
	Version int32

	FdbConn fdb.Transactor
	DbContainer directory.DirectorySubspace

	DbNode subspace.Subspace
	DbEdge subspace.Subspace
	DbIndex subspace.Subspace
	DbMeta subspace.Subspace
	DbCnt subspace.Subspace
	DbSeq subspace.Subspace
}

func (s *CDSCabinetServer) logEvent(e string){
	fmt.Println(e)
}


func StartServer() *CDSCabinetServer {
	fdb.MustAPIVersion(600)
	db := fdb.MustOpenDefault()

	var activeContainer = "test"
	container, err := directory.CreateOrOpen(db, []string{activeContainer}, nil)

	if err != nil {
		log.Fatal(err)
	}

	server := &CDSCabinetServer{
		Version: 	 1,
		FdbConn: 	 db,
		DbContainer: container,

		DbNode: 	container.Sub("n"),
		DbEdge: 	container.Sub("e"),
		DbIndex: 	container.Sub("i"),
		DbMeta: 	container.Sub("m"),
		DbCnt: 		container.Sub("c"),
		DbSeq: 		container.Sub("server"),
	}

	// install db
	var coreSeq = []string{
		"n", "e", "i", "m", "c",
	}

	_, err = server.FdbConn.Transact(func (tr fdb.Transaction) (interface{}, error) {
		tr.ClearRange(server.DbContainer)

		for i := range coreSeq {
			tr.Set(server.DbSeq.Pack(tuple.Tuple{coreSeq[i], "l"}), []byte(strconv.FormatUint(uint64(1), 10)))
		}

		log.Printf("[I] Container [%server] is now initialized", activeContainer)
		return nil, nil
	})

	return server
}

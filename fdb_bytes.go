// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.


package main

import (
	"bytes"
	"encoding/binary"
)

func intToBytes(v int64) ([]byte, error){
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, v)

	if err != nil{
		return nil, err
	}

	return buf.Bytes(), nil
}

func BytesToInt(data []byte) (ret int64, err error){
	buf := bytes.NewBuffer(data)
	err = binary.Read(buf, binary.LittleEndian, &ret)
	return ret, err
}

func prepareProperties(src []byte) (prop []byte){
	if src == nil{
		prop = []byte("")
	}else{
		prop = src
	}

	return
}

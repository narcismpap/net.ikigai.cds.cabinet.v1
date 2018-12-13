// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.


package main

import (
	"bytes"
	"encoding/binary"
	"strconv"
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


func intToKeyElement(v uint16) string{
	return strconv.FormatUint(uint64(v), 36)
}

func KeyElementToInt(k string) (uint16, error){
	v, e := strconv.ParseUint(k, 36, 32)

	if e != nil{
		return 0, e
	}

	return uint16(v), nil
}

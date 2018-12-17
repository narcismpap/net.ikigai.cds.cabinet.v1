// Package: net.ikigai.cds
// Module: cabinet.services
//
// Author: Narcis M. PAP
// Copyright (c) 2018 Ikigai Cloud. All rights reserved.

package iri

import "errors"

func Parse(s string) (IRI, error){
	if s[:2] == "s/"{
		seq := &Sequence{}
		err := seq.Parse(s)
		return seq, err
	}else if s[:2] == "e/"{
		seq := &Edge{}
		err := seq.Parse(s)
		return seq, err
	}else if s[:2] == "i/"{
		seq := &NodeIndex{}
		err := seq.Parse(s)
		return seq, err
	}else if s[:2] == "m/"{
		if s[2:4] == "n/"{
			seq := &NodeMeta{}
			err := seq.Parse(s)
			return seq, err
		}else if s[2:4] == "e/"{
			seq := &EdgeMeta{}
			err := seq.Parse(s)
			return seq, err
		}
	}else if s[:2] == "n/"{
		seq := &Node{}
		err := seq.Parse(s)
		return seq, err
	}

	return nil, errors.New("unable to parse IRI")
}

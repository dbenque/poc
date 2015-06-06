package cliTypes

import (
	"encoding/json"
	"poc/server"
)

type Pod struct {
	Fieldstr string
	Fieldint int
	// here we omit any field that we don't want to see in CLI
	// This allow us to have a view on some field of the POD without the need of depends on K8s
}

type Silo struct {
	serverAPI.SiloBase
	Spec Pod
}

type RawSilo struct {
	serverAPI.SiloBase
	Spec json.RawMessage
}

func NewSiloFromBluePrint(bp string) (*RawSilo, *Silo, error) {
	var s Silo
	if err := json.Unmarshal([]byte(bp), &s); err != nil {
		return nil, nil, err
	}
	var r RawSilo
	if err := json.Unmarshal([]byte(bp), &r); err != nil {
		return nil, nil, err
	}

	return &r, &s, nil
}

package main

import (
	"errors"
	"strconv"
)

//go run main.go

type Result struct {
	Type   string
	Length int
	Value  string
}

func CrearResult(cosa string) (Result, error) {

	if len(cosa) > 5 {

		tipo := cosa[0:2]
		largo := cosa[2:4]
		valor := cosa[4:]

		largoParseado, _ := strconv.ParseInt(largo, 0, 8)

		return Result{tipo, int(largoParseado), valor}, nil

	} else {
		return Result{}, errors.New("cadena incorrecta")
	}
}

package main

import (
	"fmt"
	"strconv"
)

//go run main.go

func main() {
	cadena := "TX03ABC"
	resultado := CrearResult(cadena)
	fmt.Println(resultado)
}

type Result struct {
	Type   string
	Value  string
	Length int
}

func CrearResult(cosa string) Result {

	tipo := cosa[0:2]
	largo := cosa[2:4]
	valor := cosa[4:]

	largoParseado, _ := strconv.ParseInt(largo, 0, 8)

	return Result{tipo, valor, int(largoParseado)}
}

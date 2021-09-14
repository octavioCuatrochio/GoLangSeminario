package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrearResult(t *testing.T) {
	var cases = []struct {
		Input   string // input string in order to be parsed
		Success bool   // paser result
		Type    string // the input type
		Value   string // the input value
		Length  int    // value length
	}{
		{"TX02AB", true, "TX", "AB", 2},
		{"NN100987654321", true, "NN", "0987654321", 10},
		{"TX06ABCDE", false, "", "", 0},
		{"NN04000A", false, "", "", 0},
	}

	for _, testData := range cases {

		re := Result{}
		re, err := CrearResult(testData.Input)
		//re, err := ParseString([]byte(testData.Input))
		// if err == nil {
		// 	fmt.Println(re)
		// 	fmt.Println(testData)

		// } else {
		// 	panic(err)
		// }

		assert.Equal(re, err == nil, testData.Success)

		fmt.Println(re, err)
		// acá agregar chequeos propios del test por ejemplo:

	}
}

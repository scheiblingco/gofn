package main

import (
	"fmt"

	"github.com/scheiblingco/gofn/typetools"
)

// case "int", "int32", "int64", "uint", "uint32", "uint64", "float32", "float64", "uint8", "uint16", "uintptr":
func main() {
	intg := 123
	flot := 123.456
	strg := "123"

	fmt.Println(typetools.EnsureString(intg))
	fmt.Println(typetools.EnsureString(flot))
	fmt.Println(typetools.EnsureString(strg))

	fmt.Println("Hold")
}

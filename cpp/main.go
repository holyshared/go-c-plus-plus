package main

import (
	"fmt"

	"github.com/holyshared/go-c-plus-plus/cpp/example1"
)

// Dump - Dump
func Dump(p interface{}) {
	fmt.Printf("%v:%T\n", p, p)
}

func main() {
	Dump(example1.EchoInt(8))
	Dump(example1.EchoDouble(9))
}

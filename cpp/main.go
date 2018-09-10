package main

import (
	"fmt"

	"github.com/holyshared/go-c-plus-plus/cpp/example1"
	"github.com/holyshared/go-c-plus-plus/cpp/example2"
	"github.com/holyshared/go-c-plus-plus/cpp/example3"
)

// Dump - Dump
func Dump(p interface{}) {
	fmt.Printf("%v:%T\n", p, p)
}

func main() {
	Dump(example1.EchoInt(8))
	Dump(example1.EchoDouble(9))

	var p = example2.NewPoint(1, 5)
	p.Display()
	p.MoveRelative(1, 0)
	p.Display()

	fmt.Printf("%d\n", example3.Math_succ(1))
	fmt.Printf("%d\n", example3.Math_succ(example3.Math_succ(1)))
}

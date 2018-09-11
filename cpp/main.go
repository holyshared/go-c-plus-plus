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

func example1Test() {
	Dump(example1.EchoInt(8))
	Dump(example1.EchoDouble(9))
}

func example2Test() {
	var p = example2.NewPoint(1, 5)
	defer example2.DeletePoint(p)
	p.Display()
	p.MoveRelative(1, 0)
	p.Display()
}

func example3Test() {
	fmt.Printf("%d\n", example3.Math_succ(1))
	fmt.Printf("%d\n", example3.Math_succ(example3.Math_succ(1)))

	var p2 = example3.NewGraphicPoint(10, 5)
	defer example3.DeleteGraphicPoint(p2)
	p2.Display()
}

func main() {
	example1Test()
	example2Test()
	example3Test()
}

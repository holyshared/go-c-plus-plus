package main

import (
	"fmt"

	"github.com/holyshared/go-c-plus-plus/cpp/example1"
	"github.com/holyshared/go-c-plus-plus/cpp/example2"
	"github.com/holyshared/go-c-plus-plus/cpp/example3"
	"github.com/holyshared/go-c-plus-plus/cpp/example4"
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

func example4Test() {
	var item = example4.NewIntPair(1, 2)
	defer example4.DeleteIntPair(item)

	example4.Dump(item)

	var item2 = example4.NewStringPair("a", "b")
	defer example4.DeleteStringPair(item2)

	example4.Dump2(item2)

	var item3 = example4.NewStringPair("a", "b")
	defer example4.DeleteStringPair(item3)

	example4.Dump3(item3)

	var item4 = example4.Dump4()
	defer example4.DeleteStringPair(item4)
	fmt.Printf("f: %s\n", item4.GetFirst())
	fmt.Printf("s: %s\n", item4.GetSecond())
}

func main() {
	example1Test()
	example2Test()
	example3Test()
	example4Test()
}

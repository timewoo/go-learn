package main

import (
	"fmt"
)

type Abser interface {
	Abs() int32
}

type Demo struct {
	x int32
	y string
}

type Demo1 struct {
	a int32
}

// change Demo value copy
func (demo Demo) Abs() int32 {
	demo.x = demo.x + 1
	return demo.x
}

func (demo1 Demo1) Abs() int32 {
	return demo1.a
}

// change Demo value pointer
func (demo *Demo) Scale(i int) {
	demo.x = demo.x + int32(i)
}

// the func only input pointer
func ScaleFunc(demo *Demo) {
	demo.x = demo.x + 1
}

func main() {
	demo := Demo{4, "s"}
	fmt.Println(demo.Abs())
	// not change demo value
	fmt.Println(demo)
	demo.Scale(1)
	fmt.Println(demo.Abs())
	// change demo value
	fmt.Println(demo)
	ScaleFunc(&demo)
	// change demo value
	fmt.Println(demo)
}

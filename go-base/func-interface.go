package main

import (
	"fmt"
)

type Demo struct {
	x int32
	y string
}

// change Demo value copy
func (demo Demo) Abs() int32 {
	return demo.x
}

// change Demo value pointer
//
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
	fmt.Println(demo)
	demo.Scale(1)
	fmt.Println(demo.Abs())
	fmt.Println(demo)
	ScaleFunc(&demo)
}

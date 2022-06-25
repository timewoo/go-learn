package main

import (
	"fmt"
	"math/cmplx"
)

var (
	f float64 = 23.12
	b bool
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value is %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value is %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value is %v\n", z, z)
	fmt.Printf("Type: %T Value is %v\n", f, f)
	fmt.Printf("Type: %T Value is %v\n", b, b)
}

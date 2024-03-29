package main

import (
	"fmt"
	"math"
	"strings"
)

type Demo struct {
	A int
	B string
}

func main() {
	// pointers()
	// structData()
	// arrayData()
	// slicesData()
	// mapData()
	useFnAsParamter()
}

func pointers() {
	i, j := 10, 20
	// set pointers
	p := &i
	// get i value use pointers
	fmt.Println(*p)
	// set i value use pointers
	*p = 12
	fmt.Println(i)
	p = &j
	//calculation j value
	*p = *p / 10
	fmt.Println(j)
}

func structData() {
	// fmt.Println(Demo{1, "D"})
	data := Demo{1, "D"}
	// data.A = 3
	// fmt.Println(data.A)
	p := &data
	p.A = 10
	fmt.Println(data)
}

func arrayData() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{1, 3, 2, 45, 21, 2}
	fmt.Println(primes)
}

func slicesData() {
	primes := [7]int{1, 3, 2, 45, 21, 2, 5}
	// a[low,high],include low,exclude high
	var s []int = primes[1:5]
	fmt.Println(s)

	// slices like references to arrays
	s[0] = 5
	fmt.Println(primes)

	// create slices simple
	slices := []Demo{{1, "a"}, {2, "b"}, {3, "c"}}
	fmt.Println(slices)

	// default slices length
	s = primes[:2]
	fmt.Println(s)

	// get length and capacity
	s = s[:1]
	fmt.Printf("len=%d cap=%d \n", len(s), cap(s))

	// make create slices
	a := make([]int, 0, 5)
	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)

	// slices of slices
	board := [][]string{
		[]string{"a"},
		[]string{"b"},
		[]string{"c"},
	}
	board[0][0] = "A"
	board[1][0] = "B"
	board[2][0] = "C"
	for i := 0; i < len(board); i++ {
		fmt.Println(strings.Join(board[i], " "))
	}

	//slices append
	var increase []int
	increase = append(increase, 2, 5)
	fmt.Printf("len=%d cap=%d %v\n", len(increase), cap(increase), increase)

	// rang slices
	var loop = []int{1, 34, 542, 1, 43, 2, 4}
	for i, v := range loop {
		fmt.Printf("%d = %v\n", i, v)
	}
}

func mapData() {
	var m = make(map[string]Demo)
	m["test"] = Demo{1, "value"}
	fmt.Println(m["test"])

	var n = map[string]Demo{
		"test": Demo{2, "value"},
	}
	fmt.Println(n)

	// mutating map
	n["test"] = Demo{1, "1"}
	fmt.Println(n["test"])

	delete(n, "test")
	fmt.Println(n)

	// check the key has value,if not,the ok is false,and v is empty value
	v, ok := n["test"]
	fmt.Println("the value ", v, "Present?", ok)

	v1, ok := m["test"]
	fmt.Println("the value ", v1, "Present?", ok)

	// count every words
	var char = strings.Fields("sd sd aer efds sdf a")
	var count = make(map[string]int)
	for _, v := range char {
		v1, ok := count[v]
		if ok {
			count[v] = v1 + 1
		} else {
			count[v] = 1
		}
	}
	fmt.Println(count)
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func useFnAsParamter() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 3))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

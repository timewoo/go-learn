package main

import "fmt"

type Number interface {
	int64 | float64
}

func SumInts(m map[string]int64) int64 {
	var i int64
	for _, v := range m {
		i += v
	}
	return i
}

func SumFloats(m map[string]float64) float64 {
	var i float64
	for _, v := range m {
		i += v
	}
	return i
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var i V
	for _, v := range m {
		i += v
	}
	return i
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var i V
	for _, v := range m {
		i += v
	}
	return i
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}

package main

import (
	"fmt"
	"runtime"
)

func main() {
	// forbase()
	// forwhile()
	// ifloop()
	// switchloop()
	deferuse()
}

func forbase() {
	sum := 1
	// for i := 0; i < 10; i++ {
	// 	sum+=i
	// }
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func forwhile() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func ifloop() {
	sum := 1
	// if sum < 2 {
	// 	sum -= sum
	// }
	if v := 2; v > sum {
		fmt.Println(v)
		// return
	} else {
		fmt.Println(sum)
	}
}

func switchloop() {
	switch os := runtime.GOOS; os {
	case "drawin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func deferuse() {
	fmt.Println("start")
	for i := 0; i < 10; i++ {
		// delay to call the method when deferuse() return
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

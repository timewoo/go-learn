package main

import (
	"example/greetings"
	"fmt"
)

func main() {
	messge := greetings.Hello("timewoo")
	fmt.Println(messge)
}

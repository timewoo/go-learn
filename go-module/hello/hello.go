package main

import (
	"example/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings:")
	log.SetFlags(0)
	messge, err := greetings.Hello("timewoo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messge)
}

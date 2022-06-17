package greetings

import "fmt"

// return string for given name
func Hello(name string) string {
	message := fmt.Sprintf("Hi,%v .Welcome!", name)
	return message
}

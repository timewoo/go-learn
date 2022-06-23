package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// return string for given name
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	// loop request to call Hello method
	// return the messages
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}

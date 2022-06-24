package greetings

import (
	"regexp"
	"testing"
)

// test call Hello with name,check vaild return value
func TestHelloName(t *testing.T) {
	name := "test"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("test")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("test") = %q, %v, want match for %#q,nil`, msg, err, want)
	}
}

// test call Hello with empty name,check error
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v,want "",error`, msg, err)
	}
}

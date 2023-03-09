package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testcase := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}

	for _, s := range testcase {
		reverse, err := Reverse(s.in)
		if err != nil {
			return
		}
		if reverse != s.want {
			t.Errorf("Reverse: %q, want %q", reverse, s.want)
		}
	}
}

func FuzzReverse(f *testing.F) {
	testcase := []string{"Hello, world", " ", "!12345"}
	for _, s := range testcase {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, orig string) {
		reverse, err1 := Reverse(orig)
		if err1 != nil {
			return
		}
		doubleReverse, err2 := Reverse(reverse)
		if err2 != nil {
			return
		}
		//t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(reverse), utf8.RuneCountInString(doubleReverse))
		if orig != doubleReverse {
			t.Errorf("Before: %q, after: %q", orig, doubleReverse)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(reverse) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", reverse)
		}
	})
}

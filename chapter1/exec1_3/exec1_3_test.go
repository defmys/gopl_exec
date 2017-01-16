package main

import (
	"strings"
	"testing"
)

// go test -bench .

func BenchmarkInefficient(b *testing.B) {
	args := []string{"asdf", "123", "abc", "1a2b3c", "asdf", "123", "abc", "1a2b3c"}
	for i := 0; i < b.N; i++ {
		var s, sep string
		for i := 1; i < len(args); i++ {
			s += sep + args[i]
			sep = " "
		}
	}
}

func BenchmarkJoin(b *testing.B) {
	args := []string{"asdf", "123", "abc", "1a2b3c", "asdf", "123", "abc", "1a2b3c"}
	for i := 0; i < b.N; i++ {
		strings.Join(args[1:], " ")
	}
}

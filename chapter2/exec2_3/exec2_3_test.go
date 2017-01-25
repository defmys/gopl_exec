package main

import (
	"github.com/defmys/gopl_exec/chapter2/exec2_3/popcount"
	"testing"
)

const value = 8569258629352691

func TestPopCount(t *testing.T) {
	original := popcount.PopCount(value)
	loop := popcount.PopCountLoop(value)

	if original != loop {
		t.Errorf("%d != %d", original, loop)
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(value)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountLoop(value)
	}
}

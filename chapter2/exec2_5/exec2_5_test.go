package main

import (
	"github.com/defmys/gopl_exec/chapter2/exec2_5/popcount"
	"testing"
)

const value = 92875982696293579

func TestShiftArg(t *testing.T) {
	original := popcount.PopCount(value)
	shift := popcount.PopCountClearRight(value)
	if original != shift {
		t.Errorf("%d != %d", original, shift)
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(value)
	}
}

func BenchmarkPopCountClearRight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountClearRight(value)
	}
}

package main

import (
	"github.com/defmys/gopl_exec/chapter2/exec2_4/popcount"
	"testing"
)

const value = 92875982696293579

func TestShiftArg(t *testing.T) {
	original := popcount.PopCount(value)
	shift := popcount.PopCountShiftArg(value)
	if original != shift {
		t.Errorf("%d != %d", original, shift)
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(value)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountShiftArg(value)
	}
}

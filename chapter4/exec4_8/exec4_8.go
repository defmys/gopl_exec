package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	cateCounts := map[string]int{
		"digit":   0,
		"letter":  0,
		"space":   0,
		"control": 0,
	}
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsDigit(r) {
			cateCounts["digit"]++
		}
		if unicode.IsLetter(r) {
			cateCounts["letter"]++
		}
		if unicode.IsSpace(r) {
			cateCounts["space"]++
			fmt.Printf("space %q", r)
		}
		if unicode.IsControl(r) {
			cateCounts["control"]++
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	fmt.Printf("category\tcount\n")
	for cate, n := range cateCounts {
		fmt.Printf("%s\t\t%d\n", cate, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

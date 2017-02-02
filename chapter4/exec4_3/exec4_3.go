package main

import (
	"fmt"
)

const arraySize = 6

type array [arraySize]int

func main() {
	a := array{0, 1, 2, 3, 4, 5}
	a = reverse(&a)
	fmt.Println(a)
}

func reverse(sp *array) array {
	s := *sp
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

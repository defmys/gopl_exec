package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	rotate(s, 2)
	fmt.Println(s)
}

func rotate(s []int, idx int) {
	if idx > 0 && idx < len(s) {
		tmp := make([]int, idx)
		copy(tmp, s[:idx])
		copy(s[:len(s)-idx], s[idx:])
		copy(s[len(s)-idx:], tmp)
	}
}

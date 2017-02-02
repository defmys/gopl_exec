package main

import (
	"fmt"
	"strings"
)

func main() {
	s := strings.Split("hello world hello hello hello hello world world", " ")
	fmt.Println(inplace(s))
}

func inplace(s []string) []string {
	result := []string{}
	if len(s) > 0 {
		result = append(result, s[0])
	}
	for i, j := 0, 1; j < len(s); j++ {
		if s[i] != s[j] {
			result = append(result, s[j])
			i = j
		}
	}
	return result
}

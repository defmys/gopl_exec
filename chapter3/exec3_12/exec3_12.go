package main

import (
	"bytes"
	"fmt"
)

func main() {
	check("hello world", "world hello")
	check("hello asdf world", " asdf world hello")
	check("hello world asdf", "world hello")
	check("hello world", "hello")
	check("hello world", "hello wo")
}

func check(a, b string) {
	fmt.Printf("%20s %20s %10v\n", a, b, areAnagrams([]byte(a), []byte(b)))
}

func areAnagrams(a, b []byte) bool {
	aParts := bytes.Split(a, []byte(" "))
	for _, part := range aParts {
		if !bytes.Contains(b, part) {
			return false
		}
	}

	bParts := bytes.Split(b, []byte(" "))
	for _, part := range bParts {
		if !bytes.Contains(a, part) {
			return false
		}
	}

	return true
}

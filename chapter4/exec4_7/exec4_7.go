package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "test测试ab"
	fmt.Println(reverse([]byte(s)))
}

func reverse(s []byte) string {
	resultrunes := []rune{}

	for len(s) > 0 {
		r, size := utf8.DecodeRune(s)
		resultrunes = append(resultrunes, r)
		s = s[size:]
	}

	for i, j := 0, len(resultrunes)-1; i < j; i, j = i+1, j-1 {
		resultrunes[i], resultrunes[j] = resultrunes[j], resultrunes[i]
	}

	s = []byte{}
	buf := bytes.NewBuffer(s)
	for _, r := range resultrunes {
		buf.WriteRune(r)
	}

	return buf.String()
}

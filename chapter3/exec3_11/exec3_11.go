package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%s\n", comma("12345"))
	fmt.Printf("%s\n", comma("-12345678"))
	fmt.Printf("%s\n", comma("+123456"))
	fmt.Printf("%s\n", comma("-1234.5678"))
	fmt.Printf("%s\n", comma("+12345678.1"))
	fmt.Printf("%s\n", comma("-123456789.123"))
}

func comma(s string) string {
	var sign string
	if strings.HasPrefix(s, "-") {
		sign = "-"
	} else if strings.HasPrefix(s, "+") {
		sign = "+"
	}
	if sign != "" {
		s = s[1:]
	}

	dotIndex := strings.LastIndex(s, ".")

	n := len(s)
	if dotIndex > 0 {
		n = len(s[:dotIndex])
	}
	if n <= 3 {
		return s
	}
	return sign + comma(s[:n-3]) + "," + s[n-3:]
}

package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	var s bytes.Buffer
	for i := 1; i < 10; i++ {
		s.WriteString(strconv.Itoa(i))
		fmt.Println(comma(s))
	}
}

func comma(s bytes.Buffer) string {
	var result bytes.Buffer

	length := len(s.String())
	commaCount := length / 3
	if commaCount > 0 {
		skip := length % 3
		if skip > 0 {
			result.Write(s.Next(skip))
			result.WriteString(",")
		}
		commaCount--

		for commaCount > 0 {
			result.Write(s.Next(3))
			result.WriteString(",")
			commaCount--
		}
		result.Write(s.Next(3))
	} else {
		result = s
	}
	return result.String()
}

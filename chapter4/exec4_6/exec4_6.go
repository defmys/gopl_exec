package main

import (
	"bytes"
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := "tt　　t\t \tt"
	fmt.Println(s)
	fmt.Println(inplace([]byte(s)))
}

func inplace(s []byte) string {
	var runeList []rune
	for len(s) > 0 {
		r, size := utf8.DecodeRune(s)
		runeList = append(runeList, r)
		s = s[size:]
	}

	buf := bytes.NewBuffer([]byte{})
	for i, j := 0, 1; j < len(runeList); j++ {
		c1, c2 := runeList[i], runeList[j]
		if unicode.IsSpace(c1) && unicode.IsSpace(c2) {
			if i == j-1 {
				buf.WriteRune(' ')
			}
		} else {
			if i == j-1 {
				buf.WriteRune(c1)
			}
			i = j
		}
	}

	if len(runeList) > 0 {
		buf.WriteRune(runeList[len(runeList)-1])
	}

	return buf.String()
}

package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x byte) int {
	return int(pc[(x>>0)] + pc[x>>(1*8)])
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	var count int
	for i, c := range c1 {
		count += PopCount(c ^ c2[i])
	}
	fmt.Println(count)
}

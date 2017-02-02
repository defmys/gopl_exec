package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		data, _ := reader.ReadBytes(byte('\n'))
		data = data[:len(data)-1]
		if len(os.Args) > 1 {
			if os.Args[1] == "384" {
				fmt.Printf("%x\n", sha512.Sum384(data))
				continue
			} else if os.Args[1] == "512" {
				fmt.Printf("%x\n", sha512.Sum512(data))
				continue
			}
		}

		fmt.Printf("%x\n", sha256.Sum256(data))
	}
}

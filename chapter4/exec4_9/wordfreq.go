package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(bytes.NewReader(data))
	scanner.Split(bufio.ScanWords)

	count := make(map[string]int)
	total := 0
	for scanner.Scan() {
		count[scanner.Text()]++
		total++
	}

	fmt.Printf("%-15s|count\n", "word")
	fmt.Println("-------------------------")
	for word, n := range count {
		fmt.Printf("%-15s|%.2f%%\n", word, float64(n)/float64(total)*100)
	}
}

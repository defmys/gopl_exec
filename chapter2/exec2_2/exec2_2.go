package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/tempconv"
)

func parse(arg string) {
	t, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c))
}

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			parse(arg)
		}
	} else {
		for {
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				parse(scanner.Text())
			}
		}
	}
}

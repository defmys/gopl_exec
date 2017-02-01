package main

import (
	"fmt"
)

const (
	KB = 1000
	MB = KB * 1000
	GB = MB * 1000
	TB = GB * 1000
	PB = TB * 1000
	EB = PB * 1000
	ZB = EB * 1000
	YB = ZB * 1000
)

func main() {
	fmt.Printf("%e\n", float64(KB))
	fmt.Printf("%e\n", float64(MB))
	fmt.Printf("%e\n", float64(GB))
	fmt.Printf("%e\n", float64(TB))
	fmt.Printf("%e\n", float64(PB))
	fmt.Printf("%e\n", float64(EB))
	fmt.Printf("%e\n", float64(ZB))
	fmt.Printf("%e\n", float64(YB))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type LineInfo struct {
	count     int
	filenames []string
}

func main() {
	counts := make(map[string]LineInfo)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, filename := range files {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, lineinfo := range counts {
		if lineinfo.count > 1 {
			fmt.Printf("%d\t%s\t%s\n", lineinfo.count, line, strings.Join(lineinfo.filenames, "\t"))
		}
	}
}

func countLines(f *os.File, counts map[string]LineInfo) {
	input := bufio.NewScanner(f)
    filename := f.Name()
	for input.Scan() {
		fileinfo := counts[input.Text()]
		fileinfo.count++
		if !hasFile(fileinfo.filenames, filename) {
			fileinfo.filenames = append(fileinfo.filenames, filename)
		}
		counts[input.Text()] = fileinfo
	}
}

func hasFile(filenames []string, targetname string) bool {
	for _, filename := range filenames {
		if filename == targetname {
			return true
		}
	}

	return false
}

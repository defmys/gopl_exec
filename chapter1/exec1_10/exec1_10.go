package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// go run exec1_10 https://golang.org http://gopl.io https://godoc.org

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	f, fileErr := os.Create("./output")
	if fileErr != nil {
		fmt.Printf("while opening file: %v", fileErr)
		os.Exit(1)
	}

	var output string
	for range os.Args[1:] {
		output += fmt.Sprintf("%s\n", <-ch)
	}
	output += fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())
	_, writeErr := f.WriteString(output)
	if writeErr != nil {
		fmt.Printf("while writing data to file: %v\n", writeErr)
		f.Close()
		os.Exit(1)
	}

	f.Close()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%2.fs  %7d  %s", secs, nbytes, url)
}

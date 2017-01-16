package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	urlList := []string{
		"https://golang.org",
		"http://gopl.io",
		"https://baidu.com",
		"https://wikipedia.org",
		"http://www.qq.com",
		"https://vk.com",
		"http://sina.com.cn",
		"https://amazon.com",
		"http://google.com",
	}

	start := time.Now()
	ch := make(chan string)
	for _, url := range urlList {
		go fetch(url, ch) // start a goroutine
	}
	for range urlList {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

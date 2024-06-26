package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	f, err := os.Create("times.txt")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	for range os.Args[1:] {
		fmt.Fprintln(f, <-ch)

	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	//check if have old file
	if _, err := os.Stat(url[8:] + "_resp.txt"); err == nil {
		os.Remove(url[8:] + "_resp.txt")
	}

	// create file
	respF, err := os.Create(url[8:] + "_resp.txt")
	if err != nil {
		ch <- fmt.Sprintf("while creating %s file: %v", url, err)
		return
	}

	nbytes, err := io.Copy(respF, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start)
	ch <- fmt.Sprintf("%s %7d %s", secs, nbytes, url)
}

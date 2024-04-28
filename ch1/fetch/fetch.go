package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func fetch() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// orignal implement
		// b, err := io.ReadAll(resp.Body)
		// resp.Body.Close()
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "fetch:reading %s: %v", url, err)
		// 	os.Exit(1)
		// }

		// exercise 1.7
		written, err := io.Copy(os.Stdout, resp.Body)

		fmt.Printf("written: %i\n", written)
	}
}

func main() {
	fetch()
}

package main

import (
	"fmt"
)

func appendInt(x []int, y int) []int {
	fmt.Printf("origin slice: cap - %d, len - %d\n", cap(x), len(x))

	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	fmt.Printf("Final slice: cap - %d, len - %d\n", cap(z), len(z))

	z[len(x)] = y
	return z
}

func appendInts(x []int, y ...int) []int {
	fmt.Printf("origin slice: cap - %d, len - %d\n", cap(x), len(x))

	var z []int
	zlen := len(x) + len(y)

	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	fmt.Printf("Final slice: cap - %d, len - %d\n", cap(z), len(z))

	copy(z[len(x):], y)
	return z
}

func nonempty(x []string) []string {
	var i int = 0

	for _, s := range x {
		if s != "" {
			x[i] = s
			i++
		}
	}
	return x[:i]
}

func main() {
	var s = []string{"a", "", "", "asdf", "sdksdk"}
	fmt.Println(nonempty(s))
	fmt.Println(s)
}

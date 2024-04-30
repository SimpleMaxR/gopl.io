package main

import "fmt"

func fun() *int {
	var n int = 2
	fmt.Println(&n)
	return &n
}

func main() {
	var o *int
	o = fun()
	*o++
	fmt.Println(o)
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func echo1() {
	fmt.Println(strings.Join(os.Args, " "))
	// strings.Join快的原因是先计算了所有待拼接字符串的总长度，然后一次性分配内存.
	// 而循环拼接则有多次内存开销。
	// 所以 strings.Join 更快。
}

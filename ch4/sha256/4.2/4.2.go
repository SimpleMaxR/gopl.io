package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var encode = flag.String("e", "sha256", "define encoding method")

// 使用指定的编码方式计算字符串的哈希值
func main() {
	fmt.Print("Input a string:")
	var input string
	fmt.Scanln(&input)
	var inputBytes = []byte(input)

	flag.Parse()
	switch *encode {
	case "sha256":
		fmt.Printf("%x\n", sha256.Sum256(inputBytes))
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512(inputBytes))
	default:
		fmt.Printf("%x\n", sha256.Sum256(inputBytes))
	}
}

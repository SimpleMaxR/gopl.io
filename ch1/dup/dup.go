package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now().UTC()

	count := make(map[string]int, 500000)

	f, err := os.Open("test_file100.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	input := bufio.NewScanner(f)
	for input.Scan() {
		count[input.Text()]++
	}

	// data, err := os.ReadFile("test_file100.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, line := range strings.Split(string(data), "\n") {
	// 	count[line]++
	// }

	end := time.Now().UTC()
	duration := end.Sub(start)

	// 输出结果
	// for line, n := range count {
	// 	if n > 1 {
	// 		fmt.Printf("%d\t%s\n", n, line)
	// 	}
	// }

	fmt.Println(duration)
	// Program use 350ms for counting
	// Program with result printed will take 700ms
}

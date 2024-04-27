package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func generate_file() {
	f, err := os.Create("test_file100.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	words := []string{"Hello", "Golang", "Yonex", "Victor", "Typescript"}

	start := time.Now()
	for i := 0; i < 10000000; i++ {
		line := ""
		for j := 0; j < 10; j++ {
			line += words[rand.Intn(len(words))] + " "
		}
		line += "\n"
		f.WriteString(line)
	}
	end := time.Now()
	fmt.Println("Time used: ", end.Sub(start))

}

func main() {
	generate_file()
}

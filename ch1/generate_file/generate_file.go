package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func gen(filename string, lines int, wg *sync.WaitGroup) {
	defer wg.Done()

	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	words := []string{"Hello", "Golang", "Yonex", "Victor", "Typescript"}

	start := time.Now()
	for i := 0; i < lines; i++ {
		line := ""
		for j := 0; j < 10; j++ {
			line += words[rand.Intn(len(words))] + " "
		}
		line += "\n"
		f.WriteString(line)
	}
	end := time.Now()
	fmt.Printf("Time used for %s: %s\n", filename, end.Sub(start))
}
func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go gen("test_file1000.txt", 10000000, wg)
	go gen("test_file10.txt", 100000, wg)
	wg.Wait()
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func countWords(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}

	return count
}

func main() {
	files := make(chan string)
	words := make(chan int, 10)

	workers := 3

	fileSling := []string{"file1.txt", "file2.txt", "file3.txt"}
	var wg sync.WaitGroup

	for i := 1; i <= workers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for f := range files {
				words <- countWords(f)
			}
		}()
	}

	go func() {
		for _, f := range fileSling {
			files <- f
		}
		close(files)
	}()

	go func() {
		sum := 0
		for w := range words {

			sum += w
		}
		fmt.Println(sum)
	}()

	wg.Wait()
	close(words)
	time.Sleep(2 * time.Second)

}

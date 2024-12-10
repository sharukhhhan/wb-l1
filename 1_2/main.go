package main

import (
	"fmt"
	"sync"
)

func calculateSquare(num int, wg *sync.WaitGroup, results chan int) {
	defer wg.Done()
	results <- num * num
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	results := make(chan int, len(numbers))
	var wg sync.WaitGroup
	for _, num := range numbers {
		wg.Add(1)
		go calculateSquare(num, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for square := range results {
		fmt.Println(square)
	}
}

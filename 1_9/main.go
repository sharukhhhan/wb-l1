package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	input := make(chan int)
	output := make(chan int)

	go func() {
		for _, num := range numbers {
			input <- num
		}
		close(input)
	}()

	go func() {
		for num := range input {
			output <- num * 2
		}
		close(output)
	}()

	for result := range output {
		fmt.Println(result)
	}
}

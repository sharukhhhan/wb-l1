package main

import (
	"fmt"
	"time"
)

func main() {
	N := 5
	data := make(chan int)
	ticker := time.NewTicker(1 * time.Second)
	stop := time.After(time.Duration(N) * time.Second)

	go func() {
		counter := 1
		for {
			select {
			case <-stop:
				close(data)
				return
			case <-ticker.C:
				data <- counter
				counter++
			}
		}
	}()

	for value := range data {
		fmt.Println("Received:", value)
	}
	
}

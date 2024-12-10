package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, data <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopping\n", id)
			return
		case value := <-data:
			fmt.Printf("Worker %d did the value: %d\n", id, value)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	data := make(chan int)
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	var numWorkers int
	fmt.Print("Enter the number of workers: ")
	fmt.Scan(&numWorkers)

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, data, &wg)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(data)
				return
			default:
				data <- rand.Intn(100)
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

	fmt.Println("Shutting down...")
	cancel()
	wg.Wait()
	fmt.Println("Program finished.")
}

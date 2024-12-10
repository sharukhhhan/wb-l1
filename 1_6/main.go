package main

import (
	"context"
	"fmt"
	"time"
)

func withChannel(stop chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("Goroutine stopped")
			return
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func withContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Goroutine stopped")
			return
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

var stop bool

func withGlobalFlag() {
	for {
		if stop {
			fmt.Println("Goroutine stopped")
			return
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// через канал
	stopChan := make(chan bool)
	go withChannel(stopChan)
	time.Sleep(2 * time.Second)
	stopChan <- true
	time.Sleep(1 * time.Second)

	// через контекст
	ctx, cancel := context.WithCancel(context.Background())
	go withContext(ctx)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

	// через флаг
	go withGlobalFlag()
	time.Sleep(2 * time.Second)
	stop = true
	time.Sleep(1 * time.Second)

	fmt.Println("Main program finished.")
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	sharedMap := make(map[int]int)
	var wg sync.WaitGroup

	writeToMap := func(key, value int) {
		defer wg.Done()
		mu.Lock()
		sharedMap[key] = value
		mu.Unlock()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writeToMap(i, i*i)
	}

	wg.Wait()

	mu.Lock()
	for k, v := range sharedMap {
		fmt.Printf("Key: %d, Value: %d\n", k, v)
	}
	mu.Unlock()
}

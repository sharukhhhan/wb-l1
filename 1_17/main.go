package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, target int) int {
	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})

	if index < len(arr) && arr[index] == target {
		return index
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	target := 7

	result := binarySearch(arr, target)
	if result != -1 {
		fmt.Printf("Element found at index %d\n", result)
	} else {
		fmt.Println("Element not found")
	}
}

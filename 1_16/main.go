package main

import "fmt"

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1
	pivot := len(arr) / 2
	arr[pivot], arr[right] = arr[right], arr[pivot]
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]
	quicksort(arr[:left])
	quicksort(arr[left+1:])

	return arr
}

func main() {
	arr := []int{10, 3, 5, 1, 8, 7, 2, 4, 6, 9}
	sorted := quicksort(arr)
	fmt.Println("Sorted array:", sorted)
}

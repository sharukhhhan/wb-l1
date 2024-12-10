package main

import "fmt"

func intersection(set1, set2 []int) []int {
	setMap := make(map[int]struct{})
	result := []int{}

	for _, num := range set1 {
		setMap[num] = struct{}{}
	}

	for _, num := range set2 {
		if _, exists := setMap[num]; exists {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{3, 4, 5, 6, 7}

	intersect := intersection(s1, s2)
	fmt.Println("Intersection:", intersect)
}

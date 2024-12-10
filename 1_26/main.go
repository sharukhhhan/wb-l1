package main

import (
	"fmt"
	"strings"
)

func hasUniqueChars(input string) bool {
	seen := make(map[rune]struct{})
	input = strings.ToLower(input)

	for _, char := range input {
		if _, exists := seen[char]; exists {
			return false
		}
		seen[char] = struct{}{}
	}

	return true
}

func main() {
	fmt.Println("abcd:", hasUniqueChars("abcd"))
	fmt.Println("abCdefAaf:", hasUniqueChars("abCdefAaf"))
	fmt.Println("aabcd:", hasUniqueChars("aabcd"))
}

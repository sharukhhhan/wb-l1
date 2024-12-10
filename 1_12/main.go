package main

import "fmt"

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})

	for _, str := range strings {
		set[str] = struct{}{}
	}

	for key := range set {
		fmt.Println(key)
	}
}

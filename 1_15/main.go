package main

// Проблема заключается в том, что при срезе строки v[:100] результирующая строка justString будет ссылаться на базовую строку v.
// Если v является большой строкой, то вся эта строка останется в памяти,
// пока justString существует, даже если мы используем только первые 100 символов.

// Чтобы исправить, нужно скопировать строку с copy

import "fmt"

var justString string

func createHugeString(size int) string {
	return string(make([]byte, size))
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]byte(v[:100]))
}

func main() {
	someFunc()
	fmt.Println(len(justString))
}

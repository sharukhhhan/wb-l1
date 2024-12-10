package main

import (
	"fmt"
)

func setBit(value int64, i uint, bit int) int64 {
	if bit == 1 {
		return value | (1 << i)
	}
	return value & ^(1 << i)
}

func main() {
	var value int64 = 42
	var bitPosition uint = 3
	newBit := 1

	fmt.Printf("Initial value: %b\n", value)
	value = setBit(value, bitPosition, newBit)
	fmt.Printf("Modified value: %b\n", value)
}

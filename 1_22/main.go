package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int).SetInt64(1 << 21)
	b := new(big.Int).SetInt64(1 << 22)

	sum := new(big.Int).Add(a, b)
	fmt.Printf("Sum: %s\n", sum.String())

	diff := new(big.Int).Sub(a, b)
	fmt.Printf("Difference: %s\n", diff.String())

	product := new(big.Int).Mul(a, b)
	fmt.Printf("Product: %s\n", product.String())

	division := new(big.Int).Div(b, a)
	fmt.Printf("Quotient: %s\n", division.String())

	aFloat := new(big.Float).SetInt(a)
	bFloat := new(big.Float).SetInt(b)
	quotientFloat := new(big.Float).Quo(bFloat, aFloat)
	fmt.Printf("Divistion (float): %s\n", quotientFloat.Text('f', 15))
}

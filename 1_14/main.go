package main

import (
	"fmt"
	"reflect"
)

func detectType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Type: int")
	case string:
		fmt.Println("Type: string")
	case bool:
		fmt.Println("Type: bool")
	case chan int:
		fmt.Println("Type: channel of int")
	default:
		fmt.Println("Unknown type:", reflect.TypeOf(v))
	}
}

func main() {
	a := 42
	b := "hello"
	c := true
	d := make(chan int)

	detectType(a)
	detectType(b)
	detectType(c)
	detectType(d)
	detectType(3.14)
}

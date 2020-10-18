package main

import (
	"fmt"
)

type I interface{}

func do(i I) {
	switch v := i.(type) {
	case int:
		fmt.Printf("int %v, %T\n", v, v)
	case string:
		fmt.Printf("string %v, %T\n", v, v)
	case float32:
		fmt.Printf("float %v, %T\n", v, v)
	case float64:
		fmt.Printf("float %v, %T\n", v, v)
	case bool:
		fmt.Printf("Bool %v, %T\n", v, v)
	default:
		fmt.Printf("No idea!")
	}
}

func main() {
	do(21)
	do("Hello")
	do(true)
	do(21.1)
}

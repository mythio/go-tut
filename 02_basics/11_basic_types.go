package main

import (
	"fmt"
)

var (
	ToBe   bool       = false
	MaxInt uint32     = 1 << 64
	z      complex128 = 12i + 4
)

var (
	lol   int = 2312
	lol32 int = 213
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %v\n", lol32, lol32)
}

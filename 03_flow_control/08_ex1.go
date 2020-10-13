package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for z < x/z {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

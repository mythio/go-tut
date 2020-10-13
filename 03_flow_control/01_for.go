package main

import "fmt"

func main() {
	sum := 0
	i := 0
	for i < 10 {
		sum += i
		i++
	}
	fmt.Println(sum)
	fmt.Println(22.0 / 7)
}

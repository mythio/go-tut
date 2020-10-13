package main

import "fmt"

func add(x int, y int, s int) int {
	return x + y + s
}

func main() {
	fmt.Println(add(42, 13, 3))
}

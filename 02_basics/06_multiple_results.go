package main

import "fmt"

func swap(x, y string) (string, string, int) {
	return y, x, 34
}

func main() {
	a, b, c := swap("hello", "world")
	fmt.Println(a, b, c)
}

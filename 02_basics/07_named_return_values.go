package main

import "fmt"

func split(sum int) (x int) {
	x = sum * 4 / 9
	var y int = 2

	x = x + y

	return
}

func main() {
	fmt.Println(split(17))
}

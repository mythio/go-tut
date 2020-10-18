package main

import "fmt"

type I interface{}

func main() {
	var i I = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	var ii I = 45.2

	f, ok := ii.(float64)
	fmt.Println(f, ok)

	vari := ii.(string)
	fmt.Println(vari)
}

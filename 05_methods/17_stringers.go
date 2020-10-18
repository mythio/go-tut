package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{Name: "Arthur", Age: 42}
	z := Person{Name: "Bumble", Age: 123123}

	fmt.Println(a, z)
}

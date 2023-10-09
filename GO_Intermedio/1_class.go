package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func main() {
	e := Employee{}
	fmt.Printf("%v\n", e)
	e.id = 1
	e.name = "John"
	fmt.Printf("%v\n", e)
}

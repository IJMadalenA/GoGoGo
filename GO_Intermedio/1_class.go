package main

import "fmt"

type Employee struct {
	// This is a field declaration. It's a variable that's part of a struct.
	id   int
	name string
}

func (e *Employee) SetId(id int) {
	// This is a pointer receiver method. It can modify the value to which the receiver points.
	e.id = id
}

func (e *Employee) SetName(name string) {
	// This is a value receiver method. It can't modify the value to which the receiver points.
	e.name = name
}

func main() {
	e := Employee{}
	fmt.Printf("%v\n", e)
	e.id = 1
	e.name = "John"
	e.SetId(2)
	fmt.Printf("%v\n", e)
	e.SetId(3)
	e.SetName("Jane")
	fmt.Printf("%v\n", e)
}

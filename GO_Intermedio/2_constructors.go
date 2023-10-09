package main

import "fmt"

type Employee2 struct {
	// This is a field declaration. It's a variable that's part of a struct.
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee2 {
	return &Employee2{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	// first way to create a struct
	e := Employee2{}
	fmt.Printf("First way: %v\n", e)

	// second way to create a struct
	e2 := Employee2{
		1,
		"John",
		true,
	}
	fmt.Printf("Second way: %v\n", e2)

	// third way to create a struct
	// The reserved word new is used to create a pointer to a struct. The zero value of a pointer is nil.
	e3 := new(Employee2)
	fmt.Printf("Third way: %v\n", *e3)
	e3.id = 1
	e3.name = "John"

	// fourth way to create a struct
	e4 := NewEmployee(1, "John", true)
	fmt.Printf("Fourth way: %v\n", *e4)

}

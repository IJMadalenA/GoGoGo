package main

import "fmt"

type Person struct {
	name   string
	age    int
	gender string
}

type Employee3 struct {
	id int
}

type FullTimeEmployee struct {
	// This is an anonymous field. It's a field that has no name.
	Person
	Employee3
}

func GetMessage(p Person) string {
	fmt.Println("%s with age %d", p.name, p.age)
	if p.age > 40 {
		return "You're over the hill!"
	}
	return "You're so young!"
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "John"
	ftEmployee.age = 30
	ftEmployee.id = 1
	fmt.Printf("%v\n", ftEmployee)
	GetMessage(ftEmployee.Person)
}

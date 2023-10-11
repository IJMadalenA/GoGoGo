package main

import "fmt"

type Person struct {
	Name string
}

type Employee3 struct {
	ID int
}

type FullTimeEmployee struct {
	Person
	Employee3
	taxRate int
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full time employee"
}

type TemporaryEmployee struct {
	Person
	Employee3
	taxRate int
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temporary employee"
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	tEmployee := TemporaryEmployee{}
	getMessage(ftEmployee)
	getMessage(tEmployee)
}

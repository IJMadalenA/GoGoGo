package main

import "time"

func Sum(x, y int) int {
	return x + y
}

func GetMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

type Person struct {
	Name string
	Age  int
	DNI  string
}

type Employee struct {
	Person
	id       int
	Position string
}

type FullTimeEmployee struct {
	Employee
	Person
}

var GetPersonByDNI = func(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	// SELECT * FROM Person WHERE DNI = dni
	return Person{}, nil
}

var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	// SELECT * FROM Employee WHERE id = id
	return Employee{}, nil
}

func GetFullTimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var fullTimeEmployee FullTimeEmployee

	e, err := GetEmployeeById(id)
	if err != nil {
		return fullTimeEmployee, err
	}

	fullTimeEmployee.Employee = e
	p, err := GetPersonByDNI(dni)
	if err != nil {
		return fullTimeEmployee, err
	}
	fullTimeEmployee.Person = p

	return fullTimeEmployee, nil
}

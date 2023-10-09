package main

import "fmt"

type figure2D interface {
	area() float64
}

type square struct {
	side float64
}

type rectangle struct {
	height float64
	width  float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func calculate(f figure2D) {
	fmt.Println("area:", f.area())
}

func main() {
	mySquare := square{side: 2.0}
	myRectangle := rectangle{height: 2.0, width: 3.0}

	fmt.Println("mySquare:", mySquare)
	fmt.Println("myRectangle:", myRectangle)

	calculate(mySquare)
	calculate(myRectangle)

	// The empty interface can be used to hold values of any type.
	// The empty interface is used by the fmt package to print values of any type.
	var myEmptyInterface interface{}
	myEmptyInterface = 1
	fmt.Println("myEmptyInterface:", myEmptyInterface)
	myEmptyInterface = "Hello"
	fmt.Println("myEmptyInterface:", myEmptyInterface)
	myEmptyInterface = true
	fmt.Println("myEmptyInterface:", myEmptyInterface)
	myEmptyInterface = 1.0
	fmt.Println("myEmptyInterface:", myEmptyInterface)
	myEmptyInterface = []int{1, 2, 3}
	fmt.Println("myEmptyInterface:", myEmptyInterface)

	// list interface{} is a slice of empty interfaces, so it can hold values of any type.
	var myList []interface{}
	myList = append(myList, 1)
	myList = append(myList, "Hello")
	myList = append(myList, true)
	myList = append(myList, 1.0)
	myList = append(myList, []int{1, 2, 3})
	fmt.Println("myList:", myList)

	myInterface := []interface{}{}
	myInterface = append(myInterface, 1)
	myInterface = append(myInterface, "Hello")
	myInterface = append(myInterface, true)
	myInterface = append(myInterface, 1.0)
	myInterface = append(myInterface, []int{1, 2, 3})
	fmt.Println("myInterface:", myInterface)
}

package main

import "fmt"

type pc struct {
	ram  int
	disk int
	core int
}

func (myPc pc) ping() {
	fmt.Println(myPc.core, "cores")
}

func (myPc *pc) duplicateRAM() {
	myPc.ram = myPc.ram * 2
}

// A pointer is a variable that stores the memory address of another variable.
// The type of a pointer is a reference to the type of the variable whose address it stores, for example:
// var a int
// var b *int
// b = &a
// In the example above, b is a pointer to an int variable, because it stores the memory address of an int variable.

func main() {
	a := 1
	// The & operator generates a pointer to its operand, for example:
	// b := &a
	// In the example above, b is a pointer to an int variable, because it stores the memory address of an int variable.
	b := &a
	// The * operator denotes the pointer's underlying value, for example:
	// c := *b
	// In the example above, c is an int variable, because it stores the underlying value of a pointer to an int variable.
	c := *b
	// The * operator can also be used to change the value of the variable whose address is stored in a pointer, for example:
	// *b = 2
	// In the example above, the value of the variable whose address is stored in b is changed to 2.
	// *b = 2 is equivalent to a = 2, because b stores the memory address of a.
	// *b means the value of the variable whose address is stored in b, and *b = 2 means the value of the variable whose address is stored in b is changed to 2.
	*b = 2

	// a and b are different variables, but they share the same memory address.
	// If the value of a variable is changed, the value of the variable whose address is stored in a pointer is also changed.
	// a prints 2 because the memory address of a is stored in b, and the value of a is changed to 2.
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)

	// The zero value of a pointer is nil. nil is a special value that means the pointer is not pointing to any memory address.

	myPc := pc{ram: 16, disk: 200, core: 4}
	fmt.Println("myPc:", myPc)

	myPc.ping()

	myPc.duplicateRAM()
	fmt.Println("myPc after calling duplicateRAM:", myPc)
}

package main

import "fmt"

func main() {
	fmt.Println("Variables.")
	integerVariable()
	zeroValues()
}

// Declaración de una variable de tipo entero.
func integerVariable() {
	base := 12
	var height int = 14
	var area int = base * height / 2

	fmt.Println("base:", base)
	fmt.Println("height:", height)
	fmt.Println("area:", area)
}

// Zero values. Si no se inicializa una variable, se le asigna un valor por defecto.
func zeroValues() {
	// El valor por defecto de una variable de tipo entero es 0.
	var a int
	// El valor por defecto de una variable de tipo float es 0.
	var b float64
	// El valor por defecto de una variable de tipo string es "".
	var c string
	// El valor por defecto de una variable de tipo boolean es false.
	var d bool

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
}

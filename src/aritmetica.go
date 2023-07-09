package main

func main() {
	// Variables
	var a int = 12
	var b int = 14

	println(suma(a, b))
	println(resta(a, b))
	println(multiplicacion(a, b))
	println(division(a, b))
	println(modulo(a, b))
	println(incremento(a))
	println(decremento(a))
}

// Suma de dos números enteros.
func suma(a, b int) int {
	return a + b
}

// Resta de dos números enteros.
func resta(a, b int) int {
	return a - b
}

// Multiplicación de dos números enteros.
func multiplicacion(a, b int) int {
	return a * b
}

// División de dos números enteros.
func division(a, b int) int {
	return a / b
}

// Módulo de dos números enteros.
func modulo(a, b int) int {
	return a % b
}

// Incremento de un número entero.
func incremento(a int) int {
	a++
	return a
}

// Decremento de un número entero.
func decremento(a int) int {
	a--
	return a
}

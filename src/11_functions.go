package main

func normalFunction(message string) {
	println(message)
}

func tripleArgument(a int, b int, c string) {
	println(a+b, c)
}

// Return one value.
func returnValue(a int, b int) int {
	return a + b
}

// Return two or more values.
// Es posible retornar más de un valor en una función en Go y es posible asignar los valores retornados a variables.
func returnValues(a int, b int) (int, int) {
	return a, b
}

func main() {
	normalFunction("Hello World.")
	tripleArgument(12, 14, "Hello World.")

	// Retornar un valor en una función en Go.
	println(returnValue(12, 14))

	// Retornar más de un valor en una función en Go.
	println(returnValues(12, 14))
	value1, value2 := returnValues(12, 14)
	println(value1, value2)
	// Obtener un valor específico retornado.
	// Go permite ignorar valores retornados con el caracter `_`.
	// Obtener el primer valor retornado.
	value1, _ = returnValues(12, 14)
	println(value1)
	// Obtener el segundo valor retornado.
	_, value2 = returnValues(12, 14)
	println(value2)

}

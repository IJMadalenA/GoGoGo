package main

func main() {

	// AND operator - &&.
	// Indica que ambas expresiones deben ser verdaderas para que la expresión completa sea verdadera.
	// En caso contrario, la expresión completa es falsa.

	// true && true = true
	println(true && true)
	// true && false = false
	println(true && false)
	// false && true = false
	println(false && true)
	// false && false = false
	println(false && false)

	// OR operator - ||.
	// Indica que al menos una de las expresiones debe ser verdadera para que la expresión completa sea verdadera.
	// En caso contrario, la expresión completa es falsa.

	// true || true = true
	println(true || true)
	// true || false = true
	println(true || false)
	// false || true = true
	println(false || true)
	// false || false = false
	println(false || false)

	// NOT operator - !.
	// Indica que la expresión debe ser falsa para que la expresión completa sea verdadera.
	// En caso contrario, la expresión completa es falsa.

	// !true = false
	println(!true)
	// !false = true
	println(!false)

}

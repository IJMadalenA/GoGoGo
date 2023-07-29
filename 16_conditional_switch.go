package main

func main() {

	// switch statement.
	// switch <expression> {
	// case <expression>:
	// 	<code>
	// case <expression>:
	// 	<code>
	// default:
	// 	<code>
	// }

	// Switch en Go es como un if-else if-else, pero más limpio.
	// En Go, el switch no tiene un break implícito, por lo que no es necesario escribirlo, a menos que se desee.

	modulo := 5 % 2
	switch modulo {
	case 0:
		println("0")
	case 1:
		println("1")
	default:
		println("default")
	}

	// switch statement with short statement.
	// switch <short statement>; <expression> {
	// case <expression>:
	// 	<code>
	// case <expression>:
	// 	<code>
	// default:
	// 	<code>
	// }

	// En Go, el switch también puede tener una declaración corta, esto es, una declaración que se ejecuta antes de la evaluación de la expresión.
	// La declaración corta se ejecuta en cada caso, y el valor de la expresión se compara con el valor de cada caso.
	switch a := 1; a {
	case 1:
		println("1")
	case 2:
		println("2")
	default:
		println("default")
	}

	// switch statement without expression.
	// switch {
	// case <expression>:
	// 	<code>
	// case <expression>:
	// 	<code>
	// default:
	// 	<code>
	// }

	// En Go, el switch también puede no tener una expresión, en cuyo caso se evalúa como true.
	// En este caso, cada caso debe tener una expresión que se evalúa como true para que el caso se ejecute.
	value := 1
	switch {
	case value == 1:
		println("true")
	case value == 2:
		println("false")
	default:
		println("default")
	}
}

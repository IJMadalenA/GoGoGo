package main

import "fmt"

func main() {
	// Variables
	integerMessage := 12
	wordMessage := "Word message."

	// Println: string.
	// La función Println imprime en la consola el mensaje que se le pasa como parámetro y
	// añade un salto de línea al final del mensaje impreso en la consola (stdout).

	fmt.Println(wordMessage)
	fmt.Println(integerMessage)

	// Printf: string.
	// La función Printf imprime en la consola el mensaje que se le pasa como parámetro en el formato especificado.
	// %s: string.
	// %d: integer.
	// %f: float.
	// %t: boolean.
	// %v: any value.

	fmt.Printf("%s\n", wordMessage)
	fmt.Printf("%d\n", integerMessage)

	// Sprintf: string.
	// La función Sprintf devuelve un string con el mensaje que se le pasa como parámetro en el formato especificado.
	// Este no imprime en la consola el mensaje que se le pasa como parámetro sino que solo devuelve un string con el
	// mensaje, el cual se puede imprimir en la consola con la función Println.

	message := fmt.Sprintf("%s\n", wordMessage)
	fmt.Println(message)

	// Print: string.
	// La función Print imprime en la consola el mensaje que se le pasa como parámetro.
	// No añade un salto de línea al final del mensaje impreso en la consola (stdout).

	fmt.Print(wordMessage)
	fmt.Print(integerMessage)

	// Conocer el tipo de dato de una variable.
	// %T: type.

	fmt.Printf("%T\n", wordMessage)
	fmt.Printf("%T\n", integerMessage)
}

package main

import "fmt"

func main() {

	// En Go solo existe el ciclo for, no existe el ciclo while ni el ciclo do while, pero es posible simularlos con el ciclo for.
	// La estructura del ciclo for en Go es la siguiente: for <inicialización>; <condición>; <incremento> { <cuerpo> }.
	// La inicialización, la condición y el incremento son opcionales y pueden ser omitidos.
	// La condición debe ser una expresión booleana (true o false) y si no se especifica, se asume que es true.
	// La inicialización y el incremento pueden ser una o más expresiones separadas por comas, pero no es posible usar la asignación corta (:=), solo es posible usar el operador de asignación (=).
	// La inicialización se ejecuta una sola vez antes de que se ejecute el cuerpo del ciclo for, si se especifica. Si no se especifica, no se ejecuta nada antes de que se ejecute el cuerpo del ciclo for.
	// La condición se evalúa antes de que se ejecute el cuerpo del ciclo for, si se especifica. Si no se especifica, se asume que es true.
	// El incremento se ejecuta después de que se ejecute el cuerpo del ciclo for, si se especifica. Si no se especifica, no se ejecuta nada después de que se ejecute el cuerpo del ciclo for.
	// El cuerpo del ciclo for se ejecuta mientras la condición se cumpla, si se especifica. Si no se especifica, el cuerpo del ciclo for se ejecuta infinitamente.
	// Si la condición no se cumple, el ciclo for termina y el programa continúa con la siguiente instrucción después del ciclo for.
	// Si la condición no se cumple, el cuerpo del ciclo for no se ejecuta.

	// ===== FOR EN GO =====.
	// Para escribir un ciclo for en Go, se debe usar la palabra reservada for, seguida de la inicialización, la condición, el incremento y el cuerpo del ciclo for.
	// for <inicialización>; <condición>; <incremento> { <cuerpo> }
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// ===== WHILE EN GO =====.
	// Para simular un ciclo while en Go, se debe usar la palabra reservada for, seguida de la condición y el cuerpo del ciclo for.
	// for <condición> { <cuerpo> }
	j := 0
	for j <= 10 {
		fmt.Println(j)
		j++
	}

	// ===== DO WHILE EN GO =====.
	// El ciclo do-while es un ciclo while que se ejecuta al menos una vez, ya que la condición se evalúa después de ejecutar el cuerpo del ciclo.
	// Para simular un ciclo do-while en Go, se debe usar la palabra reservada for, seguida del cuerpo del ciclo for y la condición.
	// for { <cuerpo> <condición> }
	for {
		fmt.Println("Do while.")
		break
	}
	// Para asegurar que el ciclo for se ejecute al menos una vez, se debe usar la palabra reservada break para salir del ciclo for y evitar que se ejecute infinitamente.
	// Ademas la condición debe ser true para que el ciclo for se ejecute al menos una vez, por lo que se puede usar la palabra reservada true como condición.
	for next := true; next; {
		fmt.Println("Do while.")
		next = false
	}
}

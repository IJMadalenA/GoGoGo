package main

import "fmt"

func main() {
	fmt.Println("Hello World.")
	constant()
}

func constant() {
	// Declare a constant.
	const pi_1 float64 = 3.14
	const pi_2 = 3.14
	fmt.Println("pi_1:", pi_1)
	fmt.Println("pi_2:", pi_2)
}

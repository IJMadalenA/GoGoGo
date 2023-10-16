package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func sum_var(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func getValues(x int) (int, int, int) {
	return x, x * 2, x * 3
}

func getValues2(x int) (double int, triple int, quadruple int) {
	double = x * 2
	triple = x * 3
	quadruple = x * 4
	return
}

func main() {

	fmt.Println(sum(5, 5))
	fmt.Println(sum_var(5, 5, 5, 5, 5))
	printNames("Juan", "Pedro", "Pablo")
	fmt.Println(getValues(5))
	fmt.Println(getValues2(5))
}

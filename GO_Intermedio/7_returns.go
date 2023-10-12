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

func main() {

	fmt.Println(sum(5, 5))
	fmt.Println(sum_var(5, 5, 5, 5, 5))
}

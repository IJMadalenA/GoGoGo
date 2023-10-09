package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// Declaración de variables.
	var x int
	x = 8
	y := 7
	fmt.Println(x)
	fmt.Println(y)

	// Manejo de exceptions y errors.
	myValue, err := strconv.ParseInt("NaN", 0, 64)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)
	}

	// map objects - key:value.
	m := make(map[string]int)
	m["key"] = 6
	fmt.Println(m["key"])

	// Int slice.
	s := []int{1, 2, 3}
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}
	s = append(s, 16)
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}

	// GoRutines.
	c := make(chan int)
	go doSomething(c)
	<-c

	// Apuntadores.
	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)
	fmt.Println(*h)
}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done.")
	c <- 1
}

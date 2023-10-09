package main

import "fmt"

func say(s string, c chan<- string) {
	c <- s
}

func main() {
	c := make(chan string, 1)

	fmt.Println("Hello")

	go say("world", c)

	go say("Bye.", c)

	fmt.Println(<-c)
}

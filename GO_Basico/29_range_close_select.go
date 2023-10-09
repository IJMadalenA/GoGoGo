package main

import "fmt"

func message(s string, c chan<- string) {
	c <- s
}

func main() {
	// This channel can store 2 elements.
	c := make(chan string, 2)

	c <- "Hello"
	c <- "world."

	// len is the number of elements stored in the channel and cap is the capacity of the channel.
	fmt.Println(len(c), cap(c))

	// range and close are used to iterate over values received from a channel.
	// Only the sender should close a channel, never the receiver.
	close(c)

	// This action is not allowed because the channel is closed.
	// Even though the channel isn't closed, the number of elements stored in the channel is still 2.
	// c <- "Bye."

	// `range` iterates over each element as it's received from a channel, so it's not possible to use it on an open channel.
	// We need to close the channel to terminate the range loop, otherwise it will block.
	for message := range c {
		fmt.Println(message)
	}

	// `select` lets you wait on multiple channel operations.
	// `select` blocks until one of its cases can run, then it executes that case.
	// It chooses one at random if multiple are ready, so it's a good way to prevent goroutines from blocking each other when waiting for channels.
	string1 := make(chan string)
	string2 := make(chan string)
	go message("Hello", string1)
	go message("world.", string2)
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-string1:
			fmt.Println("msg1: ", msg1)
		case msg2 := <-string2:
			fmt.Println("msg2: ", msg2)
		}
	}
}

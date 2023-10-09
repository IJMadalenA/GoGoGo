package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println(text)
}

func main() {
	// A WaitGroup it's a group of goroutines that are waiting for each other to finish, so the main goroutine can wait for them to finish.
	// The main goroutine calls the Add method to set the number of goroutines to wait for.
	// Then each of the goroutines runs and calls the Done method when finishes.
	var wg sync.WaitGroup

	fmt.Println("Hello")
	// The main goroutine calls the Wait method to wait for the goroutines to finish.
	// The Wait method blocks until the counter is zero, so the main goroutine waits for the other goroutines to finish.
	// The method wg.Add(1) adds one to the counter.
	wg.Add(1)

	go say("World", &wg)

	// If you want to wait for multiple goroutines, you can call the Add method with the number of goroutines to wait for.
	// If you put a number greater than 1, the counter will be greater than 1, and the main goroutine will wait for more than one goroutine to finish.
	// If you put a number less than 1, the counter will be less than 1, and the main goroutine will not wait for any goroutine to finish.
	// The key to understanding how the WaitGroup works is to understand that the counter must be greater than zero when the main goroutine calls the Wait method.

	wg.Wait()
	// The method wg.Wait() blocks the execution until the counter is zero.
	// This method should be called from the main goroutine to wait for the other goroutines to finish.

	// The keyword go is used to create goroutines, and the keyword defer is used to call the Done method when the goroutine finishes.
	// This is an anonymous function, so it's called a function literal.
	go func(text string) {
		fmt.Println(text)
	}("Bye")

	// The main goroutine sleeps for 1 second to give time to the other goroutines to finish.
	time.Sleep(time.Second * 1)

	// Every time a goroutine is Done, the method wg.Done() is called, and the counter is decremented.
	// When the counter is zero, the main goroutine can continue with its execution.
}

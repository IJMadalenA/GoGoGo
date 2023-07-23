package main

import "fmt"

func main() {
	// defer. It will be executed the function after the function is finished.
	// It is used to close the file, close the connection, etc.
	// defer <expression>
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
}

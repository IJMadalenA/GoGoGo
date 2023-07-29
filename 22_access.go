package main

import (
	myPackage "../"
	"fmt"
)

func main() {
	var myDog myPackage.DogPublic
	myDog.Name = "Firulais"

	fmt.Println("myDog:", myDog)
}

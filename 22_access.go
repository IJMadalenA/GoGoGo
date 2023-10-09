package main

import (
	"fmt"
	myPackage "myPackage"
)

func main() {
	var myDog myPackage.DogPublic
	myDog.Name = "Firulais"

	fmt.Println("myDog:", myDog)
}

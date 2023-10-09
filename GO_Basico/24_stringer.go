package main

import "fmt"

type guitar struct {
	brand string
	model string
	price int
}

func (myGuitar guitar) String() string {
	return fmt.Sprintf("%v model %v - price: $%v", myGuitar.brand, myGuitar.model, myGuitar.price)
}

func main() {
	myGuitar := guitar{brand: "Gibson", model: "Les Paul", price: 3500}
	fmt.Println("myGuitar:", myGuitar)
}

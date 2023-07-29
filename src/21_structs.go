package main

import "fmt"

// A struct is a collection of fields.
type car struct {
	brand string
	model string
	year  int
}

func main() {
	// A struct is created using the following syntax:
	firstCar := car{"Ford", "Mustang", 1964}
	secondCar := car{brand: "Chevrolet", model: "Camaro", year: 2018}
	thirdCar := car{}

	thirdCar.brand = "Fiat"
	thirdCar.model = "Uno"
	thirdCar.year = 2010

	fourthCar := new(car)
	fourthCar.brand = "Volkswagen"
	fourthCar.model = "Gol"
	fourthCar.year = 2015

	fmt.Println("first_car:", firstCar)

	fifthCar := &car{brand: "Toyota", model: "Corolla", year: 2019}

	fmt.Println("first_car:", firstCar)
	fmt.Println("second_car:", secondCar)
	fmt.Println("third_car:", thirdCar)
	fmt.Println("fourth_car:", fourthCar)
	fmt.Println("fifth_car:", fifthCar)
}

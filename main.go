package main

import "fmt"

type Car struct {
	id		int64
	year 	int64
	brand string
}

func NewCar ( id int64, year int64, brand string) *Car {
	var c Car

	c.id = 001
	c.year = 2020
	c.brand = "BMW"

	return &c
}

func loops (y int) {
	for n := 0; n <= y; n++ {
		fmt.Println(n)
	}
}

// func main(){
// 	var newCar Car = Car{1, 2020, "BMW"}
// 	fmt.Println("My car id is ", newCar.id)
// 	fmt.Println("My car year is ", newCar.year)
// 	fmt.Println("My car brand is ", newCar.brand)
// 	fmt.Println("WELLL hello there")

// 	loops(7)

// }


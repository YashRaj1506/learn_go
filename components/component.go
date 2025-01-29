package main

import (
	"fmt"
)

type car struct {
	Make string
	Model string
	Height int
	Width int
	Frontwheel Wheel //Nester struct
	BackWheel Wheel
}

type Wheel struct {
	Radius int
	Material string
}

// Anonymous Structs
// myCar := car{}
// myCar.Frontwheel.Radius = 5

// myCar := struct {
// 	Make string
//	Model string
//	} {
//		Make : "tesla",
//		Model : "model 3"
//	}

//inheriting fields from one stryuct to another

type car struct {
	make string
	model string
}

type truck struct {
	// "car" is embedded, so tej defination of a
	// "truck" now also additionally contains all
	// of the fields of the car struct
    car
	bedsize int
}

// now instead of truck.car.model we just go for truck.model

lanesTruck = truck {
	bedsize: 10,
	car: car{
		make : "toyota",
		model: "camry",
	},
}




func top(x int ) int {
	return x + 1
}

func main() {
	// var sms int
	// var consti float64
	// var per bool
	// var username string

	// fmt.Printf(
	// 	"%v %f %v %q\n",
	// 	sms,
	// 	consti,
	// 	per,
	// 	username,
	// )

	// if length := getLength(email); length < 1 {
	// 	fmt.Println("Hye")
	// }

	a := top(5)
	fmt.Println(a)

	// var empty string or empty:= ""

}



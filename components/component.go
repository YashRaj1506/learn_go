package main

import (
	"fmt"
	// "math"
)


// // structs

// type car struct {
// 	Make string
// 	Model string
// 	Height int
// 	Width int
// 	Frontwheel Wheel //Nester struct
// 	BackWheel Wheel
// }

// type Wheel struct {
// 	Radius int
// 	Material string
// }

// // Anonymous Structs
// // myCar := car{}
// // myCar.Frontwheel.Radius = 5

// // myCar := struct {
// // 	Make string
// //	Model string
// //	} {
// //		Make : "tesla",
// //		Model : "model 3"
// //	}

// //inheriting fields from one stryuct to another

// type car struct {
// 	make string
// 	model string
// }

// type truck struct {
// 	// "car" is embedded, so tej defination of a
// 	// "truck" now also additionally contains all
// 	// of the fields of the car struct
//     car
// 	bedsize int
// }

// now instead of truck.car.model we just go for truck.model

// lanesTruck = truck {
// 	bedsize: 10,
// 	car: car{
// 		make : "toyota",
// 		model: "camry",
// 	},
// }

// interface

// type Shape interface {
// 	Area() float64
// }

// type Rectangle struct {
// 	width, height float64
// }

// func (r Rectangle) Area() float64 {
// 	return r.width * r.height
// }

// type Circle struct {
// 	radius float64
// }

// func (c Circle) Area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// func CalculateArea(s Shape) float64 {
// 	return s.Area()
// }

// func main() {
// 	rect := Rectangle{width: 5, height: 4}
// 	circle := Circle{radius: 2}

// 	fmt.Println(
// 		"Reactangle area:",
// 		CalculateArea(rect),
// 	)
// 	fmt.Println(
// 		"Circle area:",
// 		CalculateArea(circle),
// 	)

// 	mysteryBox := interface{}(10)
// 	describeValue(mysteryBox)

// 	retrieveInt, ok := mysteryBox.(string)
// 	if ok {
// 		fmt.Println("Retrieved int:", retrieveInt)
// 	} else {
// 		fmt.Println("Value is not an integer")
// 	}
// }


// //finding type of input

// func describeValue(t interface{}) {
// 	fmt.Printf("Type: %T, Value: %v\n",t ,t)
// }

// func describeValue(t interface{}) {
// 	fmt.Printf("Type: %T, Value: %v\n", t, t)
// }
// ✅ The function accepts any type as an argument (because interface{} can hold any type).
// ✅ Inside the function:

//slices

// var myIint [10]int

func main() {

	primes := [6]int{1,2,3,4,5,6}
	slices := primes[1:4]
	fmt.Println(slices)

	mySlice := make([]int, 5, 10) //minimum size is 5 and can be max to 10

}

func sum(nums ...int) int {
	logic
}

func main() {
	total := sum(1,2,3)
}
package main

import (
	"fmt"
)

type car struct {
	Make string
	Model string
	Height int
	Width int
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



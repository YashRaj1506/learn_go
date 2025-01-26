package main

import (
	"fmt"

	"learn_go/greetings"
)

func main() {
	message := greetings.Hello("Preeti")
	fmt.Println(message)
}
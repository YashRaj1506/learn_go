package main

import (
	"fmt"
	"log"

	"learn_go/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message,err := greetings.Hello("Yash")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
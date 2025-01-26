package main

import (
	"fmt"
	"log"

	"learn_go/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Yash","Samantha","Rachel"}

	message,err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
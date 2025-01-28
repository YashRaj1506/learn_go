package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//

	rand.Seed(time.Now().UnixNano())

	var input int
	fmt.Println("Rock(0), Paper(1), Scissors(2)")
	fmt.Scanln(&input)

	RandomNum := rand.Intn(3)

	if input == 0 && RandomNum == 0{
		fmt.Println("Tie")
	} else if input == 0 && RandomNum == 1 {
		fmt.Println("Computer wins!")
	} else if input == 0 && RandomNum == 2 {
		fmt.Println("You Win")
	} else if input == 1 && RandomNum == 1{
		fmt.Println("Tie")
	} else if input == 1 && RandomNum == 2 {
		fmt.Println("You wins!")
	} else if input == 1 && RandomNum == 0 {
		fmt.Println("Computer Win")
	} else if input == 2 && RandomNum == 2{
		fmt.Println("Tie")
	} else if input == 2 && RandomNum == 1 {
		fmt.Println("You wins!")
	} else if input == 2 && RandomNum == 0 {
		fmt.Println("Computer Win")
	}
}
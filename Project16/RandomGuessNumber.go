package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Random Guess Number Project 16 Week 1")

	// Generate a random integer between 0 and 99
	randomNumber := rand.Intn(100)
	var guess int
	fmt.Print("between 0 and 99: ")

	for {
		fmt.Print("Guess a Random number: ")
		fmt.Scan(&guess)
		if guess == randomNumber {
			fmt.Println("Congratrs you found the right number which is :", randomNumber)
			break
		} else if guess > randomNumber {
			fmt.Println("too high")
		} else {
			fmt.Println("too low")

		}
	}

}

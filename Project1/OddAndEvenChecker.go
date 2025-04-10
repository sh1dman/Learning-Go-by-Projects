package main

import "fmt"

func main() {
	fmt.Println("OddAndEvenChecker Project 1 Week 1")
	fmt.Println("Enter a Nummber: ")

	var num int
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

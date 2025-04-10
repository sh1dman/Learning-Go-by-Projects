package main

import "fmt"

func main() {
	fmt.Println("Temperature Convertion Project 2 Week 1")
	var input string
	fmt.Println("Enter F for Fahrenheit to Celsius or C for Celsius to Fahrenheit: ")
	fmt.Scanln(&input)
	if input == "F" {
		var Fahrenheit float64
		fmt.Println("Enter the Fahrenheit: ")
		fmt.Scanln(&Fahrenheit)
		C := (Fahrenheit - 32) * 5 / 9
		fmt.Println("The Celsius is: ", C)
	} else if input == "C" {
		var Celsius float64
		fmt.Println("Enter the Celsius: ")
		fmt.Scanln(&Celsius)
		F := Celsius*9/5 + 32
		fmt.Println("The Fahrenheit is: ", F)
	} else {
		fmt.Println("Invalid input")
	}

}

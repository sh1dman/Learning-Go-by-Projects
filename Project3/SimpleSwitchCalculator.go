package main

import "fmt"

func main() {
	fmt.Println("Simple Switch Calculator Project 3 Week 1")

	var num1, num2 float64
	var operation string

	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)
	fmt.Print("Enter the operation: ")
	fmt.Scan(&operation)

	switch operation {
	case "+":
		fmt.Println("The sum of the two numbers is: ", num1+num2)
	case "-":
		fmt.Println("The difference of the two numbers is: ", num1-num2)
	case "*":
		fmt.Println("The product of the two numbers is: ", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Division by zero is not allowed")
		} else {
			fmt.Println("The division of the two numbers is: ", num1/num2)
		}
	default:
		fmt.Println("Invalid operation")
	}
}

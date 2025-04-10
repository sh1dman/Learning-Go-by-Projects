package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Input error handling Project 22 week 1")

	// Get first input
	fmt.Println("Enter first input:")
firstInput:
	var input1 string
	fmt.Scan(&input1)
	num1, err1 := strconv.Atoi(input1)
	if err1 != nil {
		fmt.Println("Invalid input for the first number. Please enter an integer.")
		goto firstInput
	}

	// Get second input
	fmt.Println("Enter second input:")
secondInput:
	var input2 string
	fmt.Scan(&input2)
	num2, err2 := strconv.Atoi(input2)
	if err2 != nil {
		fmt.Println("Invalid input for the second number. Please enter an integer.")
		goto secondInput
	}

	// Check for division by zero
	if num2 == 0 {
		fmt.Println("Error: Cannot divide by zero.")
		return
	}

	// Perform division
	result := float64(num1) / float64(num2)
	fmt.Printf("The result of %d divided by %d is %.2f\n", num1, num2, result)
}

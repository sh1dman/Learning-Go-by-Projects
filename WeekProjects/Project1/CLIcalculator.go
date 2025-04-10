package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

func main() {

	log.Println("Week One Project One CLI Calculator")
	for {
		fmt.Print("Enter First Number: ")
		var input1 string
		fmt.Scan(&input1)

		num1, err := numCheck(input1)
		if err != nil {
			log.Println("Invalid input, please enter a number")
			continue
		}

		fmt.Print("Enter Second Number: ")
		var input2 string
		fmt.Scan(&input2)
		num2, err := numCheck(input2)
		if err != nil {
			log.Println("Invalid input, please enter a number")
			continue
		}

		fmt.Print("Choose operation (+, -, *, /): ")
		var opp string
		fmt.Scan(&opp)
		switch opp {
		case "+":
			log.Println(summation(num1, num2))
		case "-":
			log.Println(subtraction(num1, num2))
		case "*":
			log.Println(multiplication(num1, num2))
		case "/":
			log.Println(division(num1, num2))
		default:
			log.Println("Error: Invalid operation")
		}
		fmt.Print("Do you want to continue, Y for yes and N for no:")
		var YorN string
		fmt.Scan(&YorN)
		if YorN == "N" {
			break
		}
	}
	fmt.Println("End of the Project")

}

func numCheck(n string) (float64, error) { // float64 is a 64-bit floating-point number
	num, err := strconv.ParseFloat(n, 64)
	if err != nil {
		return 0, errors.New("Invalid imput please Enter a number")
	}
	return num, nil
}

func summation(a float64, b float64) float64 {
	return a + b
}
func subtraction(a float64, b float64) float64 {
	return a - b
}

func multiplication(a float64, b float64) float64 {
	return a * b
}
func division(a float64, b float64) float64 {
	n, check := safeCheck(b)
	if check {
		return a / n
	}
	return 0

}

func zeroCheck(n float64) float64 {
	if n == 0 {
		panic("Error u cant divide by zero")
	}
	return n
}

func safeCheck(a float64) (float64, bool) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recover from Panic: ", r)
		}
	}()
	num := zeroCheck(a)
	return num, true
}

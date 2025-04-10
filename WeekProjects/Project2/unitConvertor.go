package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	var input string
	var typeNum int

	for {
		fmt.Print("Tell us what you want to convert: \n 1.Celsius to Fahrenheit. \n 2.Fahrenheit to Celsius. \n 3.Kilometers to Miles. \n 4.Miles to Kilometers. \n 5.Exit. \n")
		fmt.Scan(&typeNum)

		if typeNum == 5 {
			fmt.Println("Exiting...")
			break
		} else if typeNum < 1 || typeNum > 5 { // Fixed logic operator
			log.Println("Invalid input. Please enter a number between 1 and 5")
			continue //! this continue if the user type invalid input then it will skip one iteration and make the loop start again over
		}

		fmt.Print("Enter number to convert: ")
		fmt.Scan(&input)

		num, err := IntCheck(input)

		if err == nil {
			switch typeNum {
			case 1:
				fmt.Printf("%.2f°C = %.2f°F\n", float64(num), CelsiusToFahrenheit(float64(num)))
			case 2:
				fmt.Printf("%.2f°F = %.2f°C\n", float64(num), FahrenheitToCelsius(float64(num)))
			case 3:
				fmt.Printf("%.2f km = %.2f miles\n", float64(num), KilometersToMiles(float64(num)))
			case 4:
				fmt.Printf("%.2f miles = %.2f km\n", float64(num), MilesToKilometers(float64(num)))
			}
		} else {
			log.Println("Your input is invalid, please enter a number")
			continue
		}

		for {
			fmt.Print("Do you want to continue? (Y/N): ")
			var choice string
			fmt.Scan(&choice)
			choice = strings.ToUpper(choice) //! for Handling both upper and lower case

			if choice == "Y" {
				break
			} else if choice == "N" {
				fmt.Println("Thanks for using our services!")
				return
			} else {
				fmt.Println("Please enter Y or N")
				continue
			}
		}
	}
}

func IntCheck(n string) (int, error) {
	val, err := strconv.Atoi(n)

	if err != nil {
		return 0, errors.New("Invalid input, please enter valid numbers.")
	}
	return val, nil
}

func CelsiusToFahrenheit(num float64) float64 {
	return (num * 9 / 5) + 32
}

func FahrenheitToCelsius(num float64) float64 {
	return (num - 32) * 5 / 9
}

func KilometersToMiles(num float64) float64 {
	return num * 0.621371
}

func MilesToKilometers(num float64) float64 {
	return num * 1.60934
}
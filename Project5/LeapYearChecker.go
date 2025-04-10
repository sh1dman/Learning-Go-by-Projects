package main

import "fmt"

func main() {
	fmt.Println("LeapYearChecker Project 5 Week 1")
	var num int

	fmt.Print("Enter the year: ")
	fmt.Scan(&num)

	if (num%4 == 0 && num%100 != 0) || num%400 == 0 {
		fmt.Println("The year is a leap year")
	} else {
		fmt.Println("The year is not a leap year")
	}
}

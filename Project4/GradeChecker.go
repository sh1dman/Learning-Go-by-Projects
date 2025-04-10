package main

import "fmt"

func main() {
	fmt.Println("GradeChecker Project 4 Week 1")
	var num int

	fmt.Print("Enter the number: ")
	fmt.Scan(&num)

	if num >= 0 && num <= 100 {
		switch {
		case num >= 90:
			fmt.Println("Grade: A")
		case num >= 80:
			fmt.Println("Grade: B")
		case num >= 70:
			fmt.Println("Grade: C")
		case num >= 60:
			fmt.Println("Grade: D")
		default:
			fmt.Println("Grade: F")
		}
	} else {
		fmt.Println("Invalid grade")
	}

}

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("Armstrong Project 14 Week 1")

	var num int
	fmt.Print("Enter a Number to check: ")
	fmt.Scan(&num)

	fmt.Println(" is it Arstrong ? ", Armstrong(num))

}

// func isArmstrong(n int) bool {
// 	check := Armstrong(n)
// 	if check == n {
// 		return true
// 	}
// 	return false
// }

func Armstrong(n int) bool {
	//! Convert the number to a string
	numStr := strconv.Itoa(n)

	//! now we have the digit number
	digit := len(numStr)
	sum := 0
	for _, char := range numStr {
		digitValue := int(char - '0')
		sum += int(math.Pow(float64(digitValue), float64(digit)))
	}
	return sum == n
}

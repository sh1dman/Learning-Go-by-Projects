package main

import (
	"Project18/sum"
	"fmt"
)

func main() {
	fmt.Println("custom pakcage Project 18 week 1")
	num1, num2 := 11, 10

	fmt.Println("the sum of numbers is: ", sum.Sum(num1, num2))
}

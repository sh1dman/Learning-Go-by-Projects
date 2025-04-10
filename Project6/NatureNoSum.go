package main

import "fmt"

func main() {
	fmt.Println("Sum of Nature Numbers Project 6 Week 1")

	fmt.Print("Enter a Start number: ")
	var start int
	fmt.Scan(&start)

	fmt.Print("Enter a End number: ")
	var end int
	fmt.Scan(&end)

	sum := 0
	if start < end {
		for i := start; i <= end; i++ {
			sum += i
		}
		fmt.Println("Sum of natural numbers between", start, "and", end, "is", sum)
	} else {
		for i := end; i <= start; i++ {
			sum += i
		}
		fmt.Println("Sum of natural numbers between", end, "and", start, "is", sum)
	}
}

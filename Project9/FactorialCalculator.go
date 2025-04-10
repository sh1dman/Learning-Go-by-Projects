package main

import "fmt"

func main() {
	fmt.Println("Factorial Calculator Project 9 Week 1")
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	fmt.Println("Factorial number of", n, "is", factorial(n))
}

// ! FACTORIAL using Recursion in method with negative number
func factorial(n int) int {
	if n >= 0 {
		if n == 0 || n == 1 {
			return 1
		}
		return n * factorial(n-1)
	} else {
		if n == -1 {
			return 1
		}
		return n * factorial(n+1)
	}
}


//! NORMAL FACTORIAL using Loop in method
// func factorial(n int) int {
// 	num := 1
// 	for i := 1; i <= n; i++ {
// 		num *= i
// 	}
// 	return num
// }

//! FACTORIAL using Recursion in method
// func factorial(n int) int {
// 	if n == 0 || n == 1 {
// 		return 1
// 	}
// 	return n * factorial(n-1)
// }

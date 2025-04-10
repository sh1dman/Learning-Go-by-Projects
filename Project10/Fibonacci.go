package main

import "fmt"

func main() {
	fmt.Println("Fibonacci Project 10 Week 1")

	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	fmt.Println("the fibpnanci for: ", num, " is ", fib(num))
}

//! FIBONACCI using Recursion in method

func fib(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Println("Buildin Package Test Project 19 week 1")

	// Read an environment variable
	path := os.Getenv("PATH")
	fmt.Println("Your PATH environment variable:", path)

	// Use the math package
	num := 16.0
	fmt.Printf("The square root of %.2f is %.2f\n", num, math.Sqrt(num))
}

package main

import (
	"Project17/mathutils"
	"fmt"
)

func main() {
	fmt.Println("Custom package Project 17 week 1")
	number := 10
	fmt.Printf("The square of %d is %d\n", number, mathutils.Square(number))
}

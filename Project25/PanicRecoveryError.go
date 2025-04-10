package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	log.Println("Panic recovery error Project 25 week 1")
	log.Print("Enter a number:")

	var num int
	fmt.Scan(&num)
	if safeCheck(num) {
		log.Println("Number is positive")
		log.Println("the Square of the number is:", math.Sqrt(float64(num)))
	} else {
		log.Println("Number is negative")
	}
}

func check(num int) {
	if num < 0 {
		panic("Number is negative")
	}
}

func safeCheck(num int) bool {
	//! recover is a built-in function that recovers from a panic and used with defer
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r)
		}
	}()
	//! check is the function that will cause the panic
	check(num)
	return true
}

package main

import "fmt"

func main() {
	fmt.Println("Max Min Number Project 13 Week 1")

	fmt.Print("Whats the length of the array: ")
	var length int
	fmt.Scan(&length)
	max, min := MaxMinChecker(length)

	fmt.Printf("the Max numer is %d and Min number is %d\n", max, min)
}

func MaxMinChecker(length int) (max int, min int) {

	//! ARRAY CREATION: we used {Slice} method to create an array with user entering length
	// in go we cant use a normal method such as {var arr [length]int}
	arr := make([]int, length)

	for i := 0; i < length; i++ {
		fmt.Print("Enter an array number: ")

		fmt.Scan(&arr[i])
	}
	max = arr[0]
	min = arr[0]

	for i := 0; i < length; i++ {
		if max < arr[i] {
			max = arr[i]
		}
		if min > arr[i] {
			min = arr[i]
		}
	}
	return
}

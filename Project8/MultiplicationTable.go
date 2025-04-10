package main

import "fmt"

func main() {
	fmt.Print("Enter the number to make a multiplication table: ")
	var num int
	fmt.Scan(&num)

	multiplicationTable(num)
}

func multiplicationTable(num int) {
	var end int
	fmt.Print("Enter the end number: ")
	fmt.Scan(&end)

	for i := 1; i <= end; i++ {
		fmt.Printf("%d * %d = %d\n", num, i, num*i)
	}
	return
}

package main

import "fmt"

func main() {

	var start int
	var end int

	fmt.Print("Print Start number: ")
	fmt.Scan(&start)
	fmt.Print("Print End number: ")
	fmt.Scan(&end)

	fmt.Print(prime(start, end))
}

func prime(start, end int) (i, p int) {

	for i := start; i <= end; i++ {
		p := i
		if p == 1 || p == 2 || p == 3 {
			fmt.Println("It is a prime number: ", p)
		} else if p%1 == 0 && p%i == 0 && p%2 != 0 && p%3 != 0 {
			fmt.Println("It is a prime number: ", p)
		}

	}
	return
}

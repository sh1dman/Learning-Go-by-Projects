package main

import "fmt"

func main() {
	var start int
	var end int
	fmt.Print("Print Start number: ")
	fmt.Scan(&start)
	fmt.Print("Print End number: ")
	fmt.Scan(&end)

	printPrimes(start, end)
}

func printPrimes(start, end int) {
	for num := start; num <= end; num++ {
		if isPrime(num) {
			fmt.Println("Its a prime number:", num)
		}
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

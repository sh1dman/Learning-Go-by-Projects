package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("File error handling Project 23 week 1")
	file, err := os.Open("file.txt")
	check(err)

	scanner := bufio.NewScanner(file) //! NewScanner is a function that returns a new Scanner for the file
	for scanner.Scan() { 
		fmt.Println(scanner.Text()) // Text is a function that //! returns the next line of the file
	}
}

func check(e error) {
	if e != nil {
		fmt.Println("ama Errora :", e)
	}
}

package main

import "fmt"

func main() {
	fmt.Println("Reverse String Project 11 Week 1")
	var str string
	fmt.Print("Enter aa String: ")
	fmt.Scan(&str)
	fmt.Println("Reverse of ", str, " Is: ", rev(str))
}

func rev(s string) string {
	text := ""
	for i := len(s) - 1; i >= 0; i-- {
		text += string(s[i])
	}
	return text
}

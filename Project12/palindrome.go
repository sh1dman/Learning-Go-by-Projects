package main

import (
	"fmt"
)

func main() {
	fmt.Println("Palindrome Project 12 Week 1")
	var str string
	fmt.Print("Enter a string: ")
	fmt.Scan(&str)
	fmt.Println("Is ", str, " a palindrome? ", isPalindrome(str))
}

//! better solution is to make one method
//! the difference we return bool and also the return is comparing both sides
//! that way we reduce repitation


func isPalindrome(s string) bool {
	text := ""
	for i := len(s) - 1; i >= 0; i-- {
		text += string(s[i])
	}
	return text == s
}

// func isPalindrome(s string) bool{
// 	text := rev(s)
// 	if text == s {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func rev(s string) string {
// 	text := ""
// 	for i := len(s) - 1; i >= 0; i-- {
// 		text += string(s[i])
// 	}
// 	return text
// }

package main

import (
	"fmt"
)

var balance int = 0

func main() {
	fmt.Println("Simple ATM Project 15 Week 1")

	const PIN = 123

	fmt.Print("Enter PIN number: ")
	var pinChecker int
	fmt.Scan(&pinChecker)
	if PIN == pinChecker {
		for {
			fmt.Println("Write D for Deposit, W for Withdraw, B for Balance checker: ")
			var letter string
			fmt.Scan(&letter)
			switch letter {
			case "D":
				var addMoney int
				fmt.Print("Enter amount to deposit: ")
				fmt.Scan(&addMoney)
				//! this return to balance is crutial becasue it will update the global balance
				balance = deposit(balance, addMoney)
			case "W":
				var subMoney int
				fmt.Print("Enter amount to withdraw: ")
				fmt.Scan(&subMoney)
				balance = withdraw(balance, subMoney)
			case "B":
				fmt.Println("Balance is: ", balanceChecker(balance))
			default:
				fmt.Println("Invalid option, please try again.")
			}
			fmt.Print("Do you want to continue? Y for yes and N for no: ")
			var yorN string
			fmt.Scan(&yorN)
			if yorN != "Y" {
				break
			}
		}
	} else {
		fmt.Println("Authentication failed")
	}
}

func deposit(balance, n int) int {
	if n > 0 {
		balance += n
		fmt.Println("Deposit succeeded, new balance is: ", balance)
	} else {
		fmt.Println("Your amount is below zero")
	}
	return balance
}

func withdraw(balance, n int) int {
	if n > 0 && n <= balance {
		balance -= n
		fmt.Println("Withdraw succeeded, new balance is: ", balance)
	} else if n > balance {
		fmt.Println("Insufficient funds")
	} else {
		fmt.Println("Your amount is below zero")
	}
	return balance
}

func balanceChecker(balance int) int {
	return balance
}

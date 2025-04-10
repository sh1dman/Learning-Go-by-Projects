package main

//! for these packages to be used we need make a MODULE for it
//? go mod init Project20
// this commadn above will make a go.mod file
// THEN we need to import the package in our code
//? go get github.com/fatih/color
// this commadn above will download the package then we can get
//? go.sum file,

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("External Package Test Project 20 week 1")
	color.Red("This is a red message!")
	color.Green("This is a green message!")
	color.Yellow("This is a yellow message!")
	color.Blue("This is a blue message!")
	color.Magenta("This is a magenta message!")
	color.Cyan("This is a cyan message!")
}

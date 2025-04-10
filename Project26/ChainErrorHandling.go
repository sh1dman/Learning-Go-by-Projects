package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	log.Println("Chain error handling Project 26 week 1")
	log.Print("Checking the files...")
	err := perfomanceTask()

	if err != nil {
		log.Printf("Fatal error encountered: %v\n", err)
		panic("Program terminated due to an error.")
	}
	fmt.Println("Program completed successfully.")
}

func perfomanceTask() error {
	filename := "file.tx"

	log.Println("Trying to Reading file...", filename)

	text, err := readFile(filename)
	if err != nil {
		log.Println("Error in perfomanceTask:", err)
		return fmt.Errorf("Reading file failed... %w", err)
	}

	log.Println("File read successfully:", text)
	return nil
}

func readFile(filename string) (string, error) {
	file, err := ioutil.ReadFile(filename) //! ioutil is a package that provides utility functions for io operations
	if err != nil {
		return "", errors.New("the file does not exist")
	}

	return string(file), nil
}

package main

import (
	"errors"
	"log"
)

func connectToDatabase() {
	log.Println("Start of the database...")

	log.Println("trying to connect to the database")

	err := startDatabase()
	if ifError(err) {
		return
	}

	defer closeDatabase()

}

func startDatabase() error {
	log.Println("Database started successfully...")

	//! return nil
	return errors.New("failed to connect to the database")
}

func closeDatabase() {
	log.Println("Database closed successfully...")
}

func ifError(err error) bool {
	if err != nil {
		log.Println("Error oCCured:", err)
		return true
	}
	log.Println("No error occured")
	return false
}

func main() {
	log.Println("Defer clean up error Project 24 week 1")

	log.Println("Start of the program")

	connectToDatabase()

	log.Println("End of the program")
}

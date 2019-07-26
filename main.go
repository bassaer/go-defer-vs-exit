package main

import (
	"errors"
	"log"
)

func cleanup() {
	log.Println("cleanup.")
}

func run() error {
	return errors.New("validation error.")
}

func main() {
	defer cleanup()
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

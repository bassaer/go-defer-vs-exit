package main

import (
	"errors"
	"log"
)

func cleanup() {
	log.Println("cleanup.")
}

func validate() error {
	return errors.New("validation error.")
}

func main() {
	defer cleanup()
	log.Println("start.")
	if err := validate(); err != nil {
		log.Fatal(err)
	}
	log.Println("end.")
}

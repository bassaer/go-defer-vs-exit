package main

import (
	"errors"
	"fmt"
	"log"
)

func cleanup() {
	fmt.Println("cleanup.")
}

func validate() error {
	return errors.New("validation error.")
}

func main() {
	defer cleanup()
	fmt.Println("start")
	if err := validate(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("end")
}

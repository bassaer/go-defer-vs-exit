package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("DONE.")
	fmt.Println("OK.")
	os.Exit(1)
}

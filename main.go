package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if any arguments were passed
	if len(os.Args) == 1 {
		fmt.Println("No arguments passed")
		fmt.Println("Usage: hydro <filename>.hyd")
		os.Exit(1)
	}

	arguments := os.Args[1:]

	data, err := os.ReadFile(arguments[0])
	check(err)
	fmt.Print(string(data))

	os.Exit(0)
}

// Generalized error handling
func check(e error) {
	if e != nil {
		panic(e)
	}
}

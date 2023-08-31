package main

import (
	"fmt"
	"os"
)

type TokenType int

const (
	illegal TokenType = iota
	eof
	white
	returns
	int_lit
	semi
)

type Token struct {
	Type  TokenType
	Value *string
}

func tokenize(data string) {
	for i := 0; i < len(data); i++ {
		var c rune = rune(data[i])
		fmt.Printf("%d: %c\n", i, c)
	}
}

func main() {
	// Check if any arguments were passed
	if len(os.Args) == 1 {
		fmt.Println("No arguments passed")
		fmt.Println("Usage: hydro <filename>.hyd")
		os.Exit(1)
	}

	arguments := os.Args[1:]

	data, err := os.ReadFile(arguments[0])
	if os.IsNotExist(err) {
		fmt.Printf("%s: File does not exist\n", arguments[0])
		os.Exit(1)
	} else {
		check(err)
	}

	fmt.Print(string(data))

	tokenize(string(data))

	os.Exit(0)
}

// Generalized error handling
func check(e error) {
	if e != nil {
		panic(e)
	}
}

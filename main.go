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
	if os.IsNotExist(err) {
		fmt.Printf("%s: File does not exist\n", arguments[0])
		os.Exit(1)
	} else {
		check(err)
	}

	tokenizer := Tokenizer{src: string(data), index: 0}
	parser := Parser{tokens: tokenizer.Tokenize()}
	tree, err := parser.ParseProg()
	check(err)
	generator := Generator{prog: &tree}
	{
		err := os.WriteFile("out.asm", []byte(generator.GenProg()), 0644)
		check(err)
	}

	os.Exit(0)
}

// Generalized error handling
func check(e error) {
	if e != nil {
		panic(e)
	}
}

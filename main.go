package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if any arguments were passed
	if len(os.Args) == 1 {
		fmt.Println("No arguments passed")
		fmt.Println("Usage: hydro <filename>.hy")
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
	generator.vars = make(map[string]Var)
	out, err := generator.GenProg()
	check(err)

	err = os.WriteFile("out.asm", []byte(out), 0644)
	check(err)
}

// Generalized error handling
func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

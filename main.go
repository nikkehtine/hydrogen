package main

import (
	"fmt"
	"os"
)

func tokensToAsm(tokens []Token) string {
	var output string = "global _start\n_start:\n"

	for i, token := range tokens {
		if token.Type == TokenType(exit) {
			if i+1 < len(tokens) && tokens[i+1].Type == TokenType(int_lit) {
				if i+2 < len(tokens) && tokens[i+2].Type == TokenType(semi) {
					// output += fmt.Sprintf("ret %s\n", tokens[i+1].Value)
					output += fmt.Sprintf("    mov rax, 60\n")
					output += fmt.Sprintf("    mov rdi, %s\n", tokens[i+1].Value)
					output += fmt.Sprintf("    syscall\n")
				}
			}
		}
	}

	return output
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

	tokenizer := Tokenizer{src: string(data), index: 0}

	var tokens []Token = tokenizer.Tokenize()
	{
		err := os.WriteFile("out.asm", []byte(tokensToAsm(tokens)), 0644)
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

package main

import (
	"bytes"
	"fmt"
	"os"
	"unicode"
)

type TokenType int

const (
	eof TokenType = iota
	returns
	int_lit
	semi
)

type Token struct {
	Type  TokenType
	Value string
}

func tokenize(data string) []Token {
	var tokens []Token

	var strBuf bytes.Buffer
	for i := 0; i < len(data); i++ {
		var c rune = rune(data[i])

		// fmt.Printf("%d: %c\n", i, c)

		if unicode.IsLetter(c) {
			strBuf.WriteRune(c)
			i++

			for unicode.IsLetter(rune(data[i])) {
				c = rune(data[i])
				strBuf.WriteRune(c)
				i++
			}
			i--

			if strBuf.String() == "return" {
				tokens = append(tokens, Token{Type: returns})
				strBuf.Reset()
			} else {
				fmt.Printf("%s: No such keyword\n", strBuf.String())
				os.Exit(1)
			}
		} else if unicode.IsDigit(c) {
			strBuf.WriteRune(c)
			i++
			for unicode.IsDigit(rune(data[i])) {
				strBuf.WriteRune(rune(data[i]))
				i++
			}
			i--
			tokens = append(tokens, Token{Type: int_lit, Value: strBuf.String()})
		} else if unicode.IsSpace(c) {
			continue
		} else if c == rune(';') {
			tokens = append(tokens, Token{Type: semi})
		} else {
			fmt.Println("Ya dun goofed")
			os.Exit(1)
		}
	}
	return tokens
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

	tokenize(string(data))

	os.Exit(0)
}

// Generalized error handling
func check(e error) {
	if e != nil {
		panic(e)
	}
}

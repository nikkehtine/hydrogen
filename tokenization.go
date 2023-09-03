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
	exit
	int_lit
	semi
)

type Token struct {
	Type  TokenType
	Value string
}

type Tokenizer struct {
	src   string
	index int
}

// Private methods

// Peek at the next character(s). Accepts only one argument and ignores the rest. By default peeks one character
func (Tokenizer *Tokenizer) peek(args ...int) rune {
	ahead := 1
	if args != nil {
		ahead = args[0]
	}
	// Check if we've reached the end
	if Tokenizer.index+ahead > len(Tokenizer.src) {
		return '\x00' // Null rune
	} else {
		return rune(Tokenizer.src[Tokenizer.index])
	}
}

// Consume and return a character
func (Tokenizer *Tokenizer) consume() rune {
	char := rune(Tokenizer.src[Tokenizer.index])
	Tokenizer.index++
	return char
}

// Public methods

// Turns source code into a slice of tokens
func (Tokenizer *Tokenizer) Tokenize() []Token {
	var tokens []Token
	var buf bytes.Buffer

	for Tokenizer.peek() != '\x00' {
		if unicode.IsLetter(Tokenizer.peek()) {
			buf.WriteRune(Tokenizer.consume())

			for Tokenizer.peek() != '\x00' && isAlNum(Tokenizer.peek()) {
				buf.WriteRune(Tokenizer.consume())
			}
			if buf.String() == "exit" {
				tokens = append(tokens, Token{Type: exit})
				buf.Reset()
				continue
			} else {
				fmt.Printf("%s: No such keyword\n", buf.String())
				os.Exit(1)
			}
		} else if unicode.IsDigit(Tokenizer.peek()) {
			buf.WriteRune(Tokenizer.consume())

			for Tokenizer.peek() != '\x00' && unicode.IsDigit(Tokenizer.peek()) {
				buf.WriteRune(Tokenizer.consume())
			}
			tokens = append(tokens, Token{Type: int_lit, Value: buf.String()})
			buf.Reset()
			continue
		} else if Tokenizer.peek() == ';' {
			_ = Tokenizer.consume()
			tokens = append(tokens, Token{Type: semi})
			continue
		} else if unicode.IsSpace(Tokenizer.peek()) {
			_ = Tokenizer.consume()
			continue
		} else {
			fmt.Printf("%c: Invalid character\n", Tokenizer.peek())
			os.Exit(1)
		}
	}
	Tokenizer.index = 0
	return tokens
}

// Checks if the given rune is alphanumeric
func isAlNum(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}

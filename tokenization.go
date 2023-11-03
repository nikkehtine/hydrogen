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
	assign
	open_paren
	close_paren
	ident
	let
)

type Token struct {
	ttype TokenType
	value string
}

type Tokenizer struct {
	src   string
	index int
}

// Private methods

// Peek at the next character(s). Accepts only one argument and ignores the rest. By default peeks one character
func (Tokenizer *Tokenizer) peek(args ...int) rune {
	offset := 0
	if args != nil {
		offset = args[0]
	}

	if Tokenizer.index+offset >= len(Tokenizer.src) {
		return '\x00' // Null rune
	} else {
		return rune(Tokenizer.src[Tokenizer.index+offset])
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
				tokens = append(tokens, Token{ttype: exit})
				buf.Reset()
				continue
			} else if buf.String() == "let" {
				tokens = append(tokens, Token{ttype: let})
				buf.Reset()
				continue
			} else {
				tokens = append(tokens, Token{ttype: ident, value: buf.String()})
				buf.Reset()
				continue
			}
		} else if unicode.IsDigit(Tokenizer.peek()) {
			buf.WriteRune(Tokenizer.consume())

			for Tokenizer.peek() != '\x00' && unicode.IsDigit(Tokenizer.peek()) {
				buf.WriteRune(Tokenizer.consume())
			}
			tokens = append(tokens, Token{ttype: int_lit, value: buf.String()})
			buf.Reset()
			continue
		} else if Tokenizer.peek() == ';' {
			Tokenizer.consume()
			tokens = append(tokens, Token{ttype: semi})
			continue
		} else if Tokenizer.peek() == '=' {
			Tokenizer.consume()
			tokens = append(tokens, Token{ttype: assign})
			continue
		} else if Tokenizer.peek() == '(' {
			Tokenizer.consume()
			tokens = append(tokens, Token{ttype: open_paren})
			continue
		} else if Tokenizer.peek() == ')' {
			Tokenizer.consume()
			tokens = append(tokens, Token{ttype: close_paren})
			continue
		} else if unicode.IsSpace(Tokenizer.peek()) {
			Tokenizer.consume()
			continue
		} else {
			fmt.Printf("Invalid character: %c\n", Tokenizer.peek())
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

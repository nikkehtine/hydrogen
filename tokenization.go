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
	char  rune
	index int
}

// Private methods

// Peek at the next character(s). By default peeks one character
func (Tokenizer Tokenizer) peek(ahead int) rune {
	if ahead == 0 {
		ahead = 1
	}
	if Tokenizer.index+ahead >= len(Tokenizer.src) {
		return '\x00'
	}
	return rune(Tokenizer.src[Tokenizer.index])
}

func (Tokenizer Tokenizer) consume() rune {
	return rune(Tokenizer.src[Tokenizer.index+1])
}

// Public methods

func (tokenizer Tokenizer) Tokenize() []Token {
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

			if strBuf.String() == "exit" {
				tokens = append(tokens, Token{Type: exit})
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

package main

import (
	"errors"
)

type Parser struct {
	tokens []Token
	index  int
}

type NodeExpr struct {
	int_lit Token
}

type NodeExit struct {
	expr NodeExpr
}

func (Parser *Parser) ParseExp() NodeExpr {
	if Parser.peek().Type == int_lit {
		return NodeExpr{int_lit: Parser.consume()}
	} else {
		return NodeExpr{}
	}
}

func (Parser *Parser) Parse() (NodeExit, error) {
	exitNode := NodeExit{}
	// Equivalent of 'while (peek().has_value())'
	for Parser.peek().Type != TokenType(eof) {
		if Parser.peek().Type == exit && Parser.peek(1).Type == TokenType(open_paren) {
			Parser.consume()
			Parser.consume()
			parseExp := Parser.ParseExp()
			if parseExp != (NodeExpr{}) {
				exitNode.expr = parseExp
			} else {
				return NodeExit{}, errors.New("invalid expression")
			}
			if Parser.peek() != (Token{}) && Parser.peek().Type == TokenType(close_paren) {
				Parser.consume()
			} else {
				return NodeExit{}, errors.New("expected `)`")
			}
			if Parser.peek() != (Token{}) && Parser.peek().Type == TokenType(semi) {
				Parser.consume()
			} else {
				return NodeExit{}, errors.New("expected `;`")
			}
		}
	}
	return exitNode, nil
}

// Peek at the next token(s). Accepts only one argument and ignores the rest. By default peeks one token
func (Parser *Parser) peek(args ...int) Token {
	offset := 0
	if args != nil {
		offset = args[0]
	}

	if Parser.index+offset >= len(Parser.tokens) {
		return Token{Type: eof}
	} else {
		return Parser.tokens[Parser.index+offset]
	}
}

// Consume and return a token
func (Parser *Parser) consume() Token {
	token := Parser.tokens[Parser.index]
	Parser.index++
	return token
}

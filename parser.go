package main

import "errors"

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
	for Parser.peek().Type != eof {
		if Parser.peek().Type == exit {
			Parser.consume()
			parseExp := Parser.ParseExp()
			if parseExp != (NodeExpr{}) {
				exitNode.expr = parseExp
			} else {
				return NodeExit{}, errors.New("invalid expression")
			}

			if Parser.peek() != (Token{}) && Parser.peek().Type == semi {
				Parser.consume()
			} else {
				return NodeExit{}, errors.New("invalid expression")
			}
		}
	}
	return exitNode, nil
}

// Peek at the next token(s). Accepts only one argument and ignores the rest. By default peeks one token
func (Parser *Parser) peek(args ...int) Token {
	ahead := 1
	if args != nil {
		ahead = args[0]
	}
	// Check if we've reached the end
	if Parser.index+ahead > len(Parser.tokens) {
		return Token{Type: eof}
	} else {
		return Parser.tokens[Parser.index]
	}
}

// Consume and return a token
func (Parser *Parser) consume() Token {
	token := Parser.tokens[Parser.index]
	Parser.index++
	return token
}

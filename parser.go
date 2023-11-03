package main

import (
	"errors"
)

type Parser struct {
	tokens []Token
	index  int
}

// Expression nodes

type NodeExprIntLit struct {
	int_lit Token
}

func (NodeExprIntLit) isNodeExprVariant() {}

type NodeExprIdent struct {
	ident Token
}

func (NodeExprIdent) isNodeExprVariant() {}

type NodeExpr struct {
	variant NodeExprVariant
}

type NodeExprVariant interface {
	isNodeExprVariant()
}

// Statement nodes

type NodeStmtExit struct {
	expr NodeExpr
}

func (NodeStmtExit) isNodeStmtVariant() {}

type NodeStmtLet struct {
	ident Token
	expr  NodeExpr
}

func (NodeStmtLet) isNodeStmtVariant() {}

type NodeStmt struct {
	variant NodeStmtVariant
}

type NodeStmtVariant interface {
	isNodeStmtVariant()
}

// Program node

type NodeProg struct {
	statements []NodeStmt
}

// Methods

func (Parser *Parser) ParseExpr() (NodeExpr, error) {
	if Parser.peek().ttype == TokenType(int_lit) {
		return NodeExpr{variant: NodeExprIntLit{int_lit: Parser.consume()}}, nil
	} else if Parser.peek().ttype == TokenType(ident) {
		return NodeExpr{variant: NodeExprIdent{ident: Parser.consume()}}, nil
	} else {
		return NodeExpr{}, errors.New("invalid expression")
	}
}

func (Parser *Parser) ParseStmt() (NodeStmt, error) {
	if Parser.peek().ttype == TokenType(exit) && Parser.peek(1).ttype == TokenType(open_paren) {
		Parser.consume()
		Parser.consume()
		var stmt_exit NodeStmtExit

		parse_expr, err := Parser.ParseExpr()
		if err == nil {
			stmt_exit.expr = parse_expr
		} else {
			return NodeStmt{}, err
		}

		if Parser.peek() != (Token{}) && Parser.peek().ttype == TokenType(close_paren) {
			Parser.consume()
		} else {
			return NodeStmt{}, errors.New("expected `)`")
		}

		if Parser.peek() != (Token{}) && Parser.peek().ttype == TokenType(semi) {
			Parser.consume()
		} else {
			return NodeStmt{}, errors.New("expected `;`")
		}

		return NodeStmt{variant: stmt_exit}, nil
	} else if Parser.peek().ttype == TokenType(let) &&
		Parser.peek(1).ttype == TokenType(ident) &&
		Parser.peek(2).ttype == TokenType(assign) {
		Parser.consume()
		stmt_let := NodeStmtLet{ident: Parser.consume()}
		Parser.consume()

		parse_expr, err := Parser.ParseExpr()
		if err == nil {
			stmt_let.expr = parse_expr
		} else {
			return NodeStmt{}, errors.New("invalid expression")
		}

		if Parser.peek() != (Token{}) && Parser.peek().ttype == TokenType(semi) {
			Parser.consume()
		} else {
			return NodeStmt{}, errors.New("expected `;`")
		}

		return NodeStmt{variant: stmt_let}, nil
	}

	return NodeStmt{}, errors.New("invalid statement")
}

func (Parser *Parser) ParseProg() (NodeProg, error) {
	var prog NodeProg

	for Parser.peek() != (Token{}) && Parser.peek().ttype != TokenType(eof) {
		parse_stmt, err := Parser.ParseStmt()
		if err == nil {
			prog.statements = append(prog.statements, parse_stmt)
		} else {
			return NodeProg{}, err
		}
	}

	return prog, nil
}

// Peek at the next token(s). Accepts only one argument and ignores the rest. By default peeks one token
func (Parser *Parser) peek(args ...int) Token {
	offset := 0
	if args != nil {
		offset = args[0]
	}

	if Parser.index+offset >= len(Parser.tokens) {
		return Token{ttype: eof}
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

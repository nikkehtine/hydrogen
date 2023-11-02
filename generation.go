package main

import (
	"fmt"
)

type Generator struct {
	prog   *NodeProg
	output string
}

func (Gen *Generator) GenExpr(expr *NodeExpr) {
	switch variant := expr.variant.(type) {
	case NodeExprIntLit:
		Gen.output += fmt.Sprintf(
			"    mov rax, %s\n", variant.int_lit.value,
		)
		Gen.output += "    push rax\n"
	case NodeExprIdent:
		// TODO
	}
}

func (Gen *Generator) GenStmt(stmt *NodeStmt) {
	switch variant := stmt.variant.(type) {
	case NodeStmtExit:
		Gen.GenExpr(&variant.expr)
		Gen.output += "    mov rax, 60\n"
		Gen.output += "    pop rdi\n"
		Gen.output += "    syscall\n"
	case NodeStmtLet:
	}
}

func (Gen *Generator) GenProg() string {
	Gen.output += "global _start\n_start:\n"

	for _, stmt := range Gen.prog.statements {
		Gen.GenStmt(&stmt)
	}

	Gen.output += "    mov rax, 60\n"
	Gen.output += "    mov rdi, 0\n"
	Gen.output += "    syscall\n"
	return Gen.output
}

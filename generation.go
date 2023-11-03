package main

import (
	"errors"
	"fmt"
)

type Generator struct {
	prog       *NodeProg
	output     string
	stack_size int
	vars       map[string]Var
}

type Var struct {
	stack_loc int
}

func (Gen *Generator) GenExpr(expr *NodeExpr) error {
	switch variant := expr.variant.(type) {
	case NodeExprIntLit:
		Gen.output += fmt.Sprintf("    mov rax, %s\n", variant.int_lit.value)
		Gen.push("rax")
	case NodeExprIdent:
		if variable, exists := Gen.vars[variant.ident.value]; !exists {
			return errors.New("undeclared identifier: " + variant.ident.value)
		} else {
			offset := fmt.Sprintf("QWORD [rsp + %d]\n", (Gen.stack_size-variable.stack_loc-1)*8)
			Gen.push(offset)
		}
	}
	return nil
}

func (Gen *Generator) GenStmt(stmt *NodeStmt) error {
	switch variant := stmt.variant.(type) {
	case NodeStmtExit:
		err := Gen.GenExpr(&variant.expr)
		if err != nil {
			return err
		}

		Gen.output += "    mov rax, 60\n"
		Gen.pop("rdi")
		Gen.output += "    syscall\n"
	case NodeStmtLet:
		if _, exists := Gen.vars[variant.ident.value]; exists {
			return errors.New("variable already exists: " + variant.ident.value)
		}

		Gen.vars[variant.ident.value] = Var{stack_loc: Gen.stack_size}

		err := Gen.GenExpr(&variant.expr)
		if err != nil {
			return err
		}
	}
	return nil
}

func (Gen *Generator) GenProg() (string, error) {
	Gen.output += "global _start\n_start:\n"

	for _, stmt := range Gen.prog.statements {
		if err := Gen.GenStmt(&stmt); err != nil {
			return "", err
		}
	}

	Gen.output += "    mov rax, 60\n"
	Gen.output += "    mov rdi, 0\n"
	Gen.output += "    syscall\n"
	return Gen.output, nil
}

func (Gen *Generator) push(reg string) {
	Gen.output += fmt.Sprintf("    push %s\n", reg)
	Gen.stack_size++
}

func (Gen *Generator) pop(reg string) {
	Gen.output += fmt.Sprintf("    pop %s\n", reg)
	Gen.stack_size--
}

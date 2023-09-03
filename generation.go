package main

import "fmt"

type Generator struct {
	root *NodeExit
}

func (Generator *Generator) Generate() string {
	var output string = "global _start\n_start:\n"
	output += fmt.Sprintf("    mov rax, 60\n")
	output += fmt.Sprintf("    mov rdi, %s\n", Generator.root.expr.int_lit.Value)
	output += fmt.Sprintf("    syscall\n")
	return output
}

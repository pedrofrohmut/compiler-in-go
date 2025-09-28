// monkey/compiler/compiler.go

package compiler

/*
	TODO: Understand why it is using objects and not the ast directly
*/

import (
	"monkey/ast"
	"monkey/code"
	"monkey/object"
)

type Compiler struct {
	instructions code.Instructions
	constants []object.Object
}

func NewCompiler() *Compiler {
	return &Compiler {
		instructions: code.Instructions {},
		constants: []object.Object {},
	}
}

func (this *Compiler) Compile(program ast.Node) error {
	return nil
}

type Bytecode struct {
	Instructions code.Instructions
	Constants []object.Object
}

func (this *Compiler) NewBytecode() *Bytecode {
	return &Bytecode {
		Instructions: this.instructions,
		Constants: this.constants,
	}
}

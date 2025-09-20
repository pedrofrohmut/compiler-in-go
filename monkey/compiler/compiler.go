// monkey/compiler/compiler.go

package compiler

type Compiler struct {
	instructions code.Instructions
	constants []object.Object
}

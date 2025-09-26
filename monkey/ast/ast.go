// monkey/parser/ast.go

/*
	AST - Abstract Syntax Tree

	A data struct in a form of a tree of nodes that represent a piece o source
	code
*/

package ast

import (
	"bytes"
	"fmt"
)

type Node interface {
	node()
	String() string
}

type Expression interface {
	Node
	expression()
}

type Statement interface {
	Node
	statement()
}

type Program struct {
	Statements []Statement
}

func (this *Program) node() {}

func (this *Program) String() string {
	var out bytes.Buffer

	for _, stm := range this.Statements {
		out.WriteString(stm.String() + ";\n")
	}

	return out.String()
}

type ExpressionStatement struct {
	Expression Expression
}

func (this *ExpressionStatement) node() {}

func (this *ExpressionStatement) statement() {}

func (this *ExpressionStatement) String() string {
	return this.Expression.String()
}

type IntegerLiteral struct {
	Value int64
}

func (this *IntegerLiteral) node() {}

func (this *IntegerLiteral) expression() {}

func (this *IntegerLiteral) String() string {
	return fmt.Sprintf("%d", this.Value)
}

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

// @Impl
func (this *Program) node() {}

// @Impl
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

// @Impl
func (this *ExpressionStatement) node() {}

// @Impl
func (this *ExpressionStatement) statement() {}

// @Impl
func (this *ExpressionStatement) String() string {
	return this.Expression.String()
}

type IntegerLiteral struct {
	Value int64
}

// @Impl
func (this *IntegerLiteral) node() {}

// @Impl
func (this *IntegerLiteral) expression() {}

// @Impl
func (this *IntegerLiteral) String() string {
	return fmt.Sprintf("%d", this.Value)
}

type InfixExpression struct {
	Left Expression
	Operator string
	Right Expression
}

// @Impl
func (this *InfixExpression) node() {}

// @Impl
func (this *InfixExpression) expression() {}

// @Impl
func (this *InfixExpression) String() string {
    return "(" + this.Left.String() + " " + this.Operator + " " + this.Right.String() + ")"
}

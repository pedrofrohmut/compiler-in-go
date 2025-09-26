// monkey/tests/parser_test.go

package tests

/*
	Parser ToDo:

	[ ] Precedence Expressions
	[ ] Infix Expressions
	[ ] Prefix Expressions

	[x] Integer Literals
	[ ] String Literals
	[ ] Char Literals
	[ ] Array Literals
	[ ] Hash Literals
	[ ] Function Literals

	[ ] Expression Statements
	[ ] Let Statements
	[ ] Return Statements
	[ ] Call Statements

	[ ] Identifiers Expressions
	[ ] Call Expressions
	[ ] If/Else Expressions
	[ ] Clojure Function
	[ ] Array Index
*/


import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/parser"
	"testing"
)

func printStatements(program *ast.Program) {
	for i, stm := range program.Statements {
		fmt.Printf("[%d]: %s\n", i, stm.String())
	}
}

func checkProgStmCount(t *testing.T, program *ast.Program, expected int) {
	if len(program.Statements) != expected {
		t.Fatalf("Expected program to have %d statements but got %d instead", expected, len(program.Statements))
	}
}

func Test_Parser_ExprStmIntegerLiteral(t *testing.T) {
	var input = "1;"
	var lexer = lexer.NewLexer(input)
	var parser = parser.NewParser(lexer)
	var program = parser.ParseProgram()

	// printStatements(program)
	checkProgStmCount(t, program, 1)

	var stm, okStm = program.Statements[0].(*ast.ExpressionStatement)
	if !okStm {
		t.Errorf("Expected program first statement to be *ast.ExpressionStatement but got %T instead",
			program.Statements[0])
		return
	}

	var integerLiteral, okIntegerLiteral = stm.Expression.(*ast.IntegerLiteral)
	if !okIntegerLiteral {
		t.Errorf("Expected statement expression to be *ast.IntegerLiteral but got %T instead", stm.Expression)
		return
	}

	var expectedValue = int64(1)
	if integerLiteral.Value != expectedValue {
		t.Errorf("Expected integer literal value to be %d but got %d instead", expectedValue, integerLiteral.Value)
	}
}

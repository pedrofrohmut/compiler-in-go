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

// func testIdentifier(t *testing.T, expression ast.Expression, expected string) {
//     var iden, ok = expression.(*ast.Identifier)
//     if !ok {
//         t.Errorf("Expression is not an Identifier. Got %T instead", expression)
//         return
//     }
//     if iden.Value != expected {
//         t.Errorf("Expected Identifier value to be %s but found %s instead", expected, iden.Value)
//     }
// }

func testIntegerLiteral(t *testing.T, expression ast.Expression, expectedValue int64) {
    var intLit, ok = expression.(*ast.IntegerLiteral)
    if !ok {
        t.Errorf("Expression is not an IntegerLiteral. Got %T instead", expression)
        return
    }
    if intLit.Value != expectedValue {
        t.Errorf("Expected IntegerLiteral value to be %d but got %d instead", expectedValue, intLit.Value)
    }
}

// func testBooleanLiteral(t *testing.T, expression ast.Expression, expected bool) {
//     var boo, ok = expression.(*ast.Boolean)
//     if !ok {
//         t.Errorf("Expression is not a Boolean. Got %T instead", expression)
//         return
//     }
//     if boo.Value != expected {
//         t.Errorf("Expected Boolean value to be '%t' but got '%t' instead", expected, boo.Value)
//     }
// }

func testLiteralExpression(t *testing.T, expression ast.Expression, expected any) {
    switch tmp := expected.(type) {
    case int:
        testIntegerLiteral(t, expression, int64(tmp))
    case int64:
        testIntegerLiteral(t, expression, tmp)
    // case string:
    //     testIdentifier(t, expression, tmp)
    // case bool:
    //     testBooleanLiteral(t, expression, tmp)
    default:
        t.Errorf("Tested expression type is not a covered on testLiteralExpression. Got %T", expression)
    }
}


func testInfixExpression(t *testing.T, expression ast.Expression, expectedLeft any,
        expectedOperator string, expectedRight any) {
    var inf, ok = expression.(*ast.InfixExpression)
    if !ok {
        t.Errorf("Expression is not an InfixExpression. Got %T instead", expression)
        return
    }
    // Test Left Value
    testLiteralExpression(t, inf.Left, expectedLeft)
    // Test Operator
    if inf.Operator != expectedOperator {
        t.Errorf("Expected Operator to be '%s' but got '%s' instead", expectedOperator, inf.Operator)
    }
    // Test Right Value
    testLiteralExpression(t, inf.Right, expectedRight)
}

func getExpressionStatement(t *testing.T, stm ast.Statement) (*ast.ExpressionStatement, bool) {
	var exprStm, okStm = stm.(*ast.ExpressionStatement)
	if !okStm {
		t.Errorf("Expected program first statement to be *ast.ExpressionStatement but got %T instead", stm)
		return nil, false
	}
	return exprStm, true
}

func Test_Parser_SingleIntegerLiteral(t *testing.T) {
	var input = "1;"
	var lexer = lexer.NewLexer(input)
	var parser = parser.NewParser(lexer)
	var program = parser.ParseProgram()

	// printStatements(program)
	checkProgStmCount(t, program, 1)

	var exprStm, ok = getExpressionStatement(t, program.Statements[0])
	if !ok { return }

	testLiteralExpression(t, exprStm.Expression, 1)
}

func Test_Parser_SingleInfixExpression(t *testing.T) {
	var input = "1 + 2;"
	var lexer = lexer.NewLexer(input)
	var parser = parser.NewParser(lexer)
	var program = parser.ParseProgram()

	// printStatements(program)
	checkProgStmCount(t, program, 1)

	var exprStm, ok = getExpressionStatement(t, program.Statements[0])
	if !ok { return }

	testInfixExpression(t, exprStm.Expression, 1, "+", 2)
}

// func Test_Parser_InfixExpression(t *testing.T) {
//     tests := []struct {
//         input string; left any; operator string; right any
//     } {
//         { "5 + 5",          5,     "+",  5     },
//         { "5 - 5",          5,     "-",  5     },
//         { "5 * 5",          5,     "*",  5     },
//         { "5 / 5",          5,     "/",  5     },
//         { "5 < 5",          5,     "<",  5     },
//         { "5 > 5",          5,     ">",  5     },
//         { "5 == 5",         5,     "==", 5     },
//         { "5 != 5",         5,     "!=", 5     },
//
//         // Booleans
//         { "true == true",   true,  "==", true  },
//         { "true != false",  true,  "!=", false },
//         { "false == false", false, "==", false },
//     }
//     var acc bytes.Buffer
//     for _, x := range tests { acc.WriteString(x.input + ";\n") }
//     input := acc.String()
//     lexer := lexer.NewLexer(input)
//     parser := NewParser(lexer)
//     program := parser.ParseProgram()
//
//     checkParserErrors(t, parser)
//     if len(program.Statements) != len(tests) {
//         t.Fatalf("Expected program to have %d statements but got %d instead", len(tests), len(program.Statements))
//     }
//     // program.PrintStatements()
//
//     for i, test := range tests {
//         stm, ok := program.Statements[i].(*ast.ExpressionStatement)
//         if !ok {
//             t.Fatalf("Program first statement is not an ExpressionStatement, got %s instead", program.Statements[0])
//         }
//
//         testInfixExpression(t, stm.Expression, test.left, test.operator, test.right)
//     }
// }

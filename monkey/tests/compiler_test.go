// monkey/tests/compiler_test.go

package tests

import (
	"fmt"
	"monkey/ast"
	"monkey/code"
	"monkey/compiler"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"testing"
)


type CompilerTestCase struct {
	input string
	expectedConstants []any
	expectedInstructions []code.Instructions
}

func parseProgram(input string) *ast.Program {
	var lexer = lexer.NewLexer(input)
	var parser = parser.NewParser(lexer)
	var program = parser.ParseProgram()

	for i, err := range parser.Errors() {
		fmt.Printf("[%d] %s\n", i, err)
	}

	return program
}

func concatInstructions(instructions []code.Instructions) code.Instructions {
	var acc = code.Instructions {}
	for _, instruction := range instructions {
		acc = append(acc, instruction...)
	}
	return acc
}

func testInstructions(instructions code.Instructions, expected []code.Instructions) error {
	var concatted = concatInstructions(expected)

	if len(instructions) != len(expected) {
		return fmt.Errorf("Wrong instructions length. Expected %q but got %q instead",
			concatted, instructions)
	}

	for i := range expected {
		if instructions[i] != concatted[i] {
			return fmt.Errorf("Wrong instruction at %d. Expected %d but got %d instead",
				i, instructions, concatted[i])
		}
	}

	return nil
}

func testIntegerObject(check object.Object, expected int64) error {
	var intObj, okIntObj = check.(*object.Integer)
	if !okIntObj {
		return fmt.Errorf("Expected object to be of type object.Integer but got %T instead", check)
	}

	if intObj.Value != expected {
		return fmt.Errorf("Expected object.Integer value to be %d but got %d instead", expected, intObj.Value)
	}

	return nil
}

func testConstants(constants []object.Object, expected []any) error {
	if len(constants) != len(expected) {
		return fmt.Errorf("Wrong constants length. Expected %d but got %d instead", len(expected), len(constants))
	}

	for i := range expected {
		switch constant := expected[i].(type) {
		case int:
			var err = testIntegerObject(constants[i], int64(constant))
			if err != nil {
				return fmt.Errorf("Constant %d - testIntegerObject failed: %s", i, err)
			}
		}
	}

	return nil
}

func Test_Compiler_IntegerArithmetic(t *testing.T) {
	var tests = []CompilerTestCase {
		{
			"1 + 2;",
		  	[]any { 1, 2 },
			[]code.Instructions { code.Make(code.OpConstant, 0), code.Make(code.OpConstant, 1) },
		},
	}

	for _, test := range tests {
		var program = parseProgram(test.input)

		var compiler = compiler.NewCompiler()
		var err = compiler.Compile(program)
		if err != nil {
			t.Errorf("Compiler error: %s\n", err)
			continue
		}

		var bytecode = compiler.NewBytecode()

		err = testInstructions(bytecode.Instructions, test.expectedInstructions)
		if err != nil {
			t.Errorf("Test instructions failed: %s\n", err)
		}

		err = testConstants(bytecode.Constants, test.expectedConstants)
		if err != nil {
			t.Errorf("Test constants failed: %s\n", err)
		}
	}
}

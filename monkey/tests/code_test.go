package tests

import (
	"monkey/code"
	"testing"
)

func Test_Code_Make(t *testing.T) {
	var tests = []struct {
		opcode code.Opcode
		operands []int
		expected []byte
	} {
		{ code.OpConstant, []int { 65534 }, []byte { byte(code.OpConstant), 255, 254 } },
	}

	for _, test := range tests {
		var instruction = code.Make(test.opcode, test.operands...)

		if len(instruction) != len(test.expected) {
			t.Errorf("Instruction has wrong length. Expected=%d but got=%d instead.",
				len(instruction), len(test.expected))
		}

		for i := range test.expected {
			if instruction[i] != test.expected[i] {
				t.Errorf("Wrong byte at position %d. Expected=%d but got=%d instead.",
					i, test.expected[i], instruction[i])
			}
		}
	}
}

func Test_Code_ReadOperands(t *testing.T) {
	var tests = []struct {
		opcode code.Opcode; operands []int; bytesRead int
	} {
		{ code.OpConstant, []int { 65535 }, 2 },
	}

	for _, test := range tests {
		var instruction = code.Make(test.opcode, test.operands...)

		var def, err = code.GetDefinition(byte(test.opcode))
		if err != nil {
			t.Errorf("Definition not found: %q\n", err)
			continue
		}

		var operandsRead, n = code.ReadOperands(def, instruction[1:])
		if n != test.bytesRead {
			t.Errorf("N Wrong. Wanted %d but got %d instead", test.bytesRead, n)
			continue
		}

		for i, expectedOperand := range test.operands {
			if operandsRead[i] != expectedOperand {
				t.Errorf("Operand wrong. Expected %q but got %q instead", expectedOperand, operandsRead)
			}
		}
	}
}

func Test_Code_InstructionsToString(t *testing.T) {
	var instructionsArr = []code.Instructions {
		code.Make(code.OpConstant, 1),
		code.Make(code.OpConstant, 2),
		code.Make(code.OpConstant, 65535),
	}

	// Do not fix the indentation of 'expected'. It will break the test
	var expected = `000 OpConstant 1
003 OpConstant 2
006 OpConstant 65535
`

	var concatted = code.Instructions{}
	for _, instruction := range instructionsArr {
		concatted = append(concatted, instruction...)
	}

	if concatted.String() != expected {
		t.Errorf("Instructions wrongly formatted. Wanted '%q' but got '%q' instead",
			expected, concatted.String())
	}
}

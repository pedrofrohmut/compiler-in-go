package tests

import (
	"monkey/code"
	"testing"
)

func TestMake(t *testing.T) {
	var tests = []struct {
		opcode code.Opcode
		operands []int
		expected []byte
	}{
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

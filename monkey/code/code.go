package code

import "fmt"

type Instruction []byte

type Opcode byte

const (
	OpConstant Opcode = iota
)

type Definition struct {
	Name string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition {
	OpConstant: { "Opcontant", []int {2} }, // uint16
}

func Lookup(op byte) (*Definition, error) {
	var definition, found = definitions[Opcode(op)]
	if !found {
		return nil, fmt.Errorf("No Opcode defined for %d", op)
	}
	return definition, nil
}

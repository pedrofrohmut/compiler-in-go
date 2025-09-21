package code

/*
   package intended to create the bytecode
   1. Create the instructions as byte arrays
 */

import (
	"encoding/binary"
	"fmt"
)

// 1 byte wise, the first byte in the instruction
type Opcode byte

const (
	OpConstant Opcode = iota // Holds an int that is a reference to the values in the constant pool
)

// A single instruction will not be explict defined to avoid too much casting and
// will be kept as a byte array for simplicity
type instructions []byte

// Struct mainly for the debugging purposes with a human readable name
type Definition struct {
	Name string					// Human readable
	OperandWidths []int			// Number of bytes each operand takes
}

var definitions = map[Opcode]*Definition {
	OpConstant: { "OpConstant", []int { 2 } },
}

func GetDefinition(opcode byte) (*Definition, error) {
	var definition, found = definitions[Opcode(opcode)]
	if !found {
		return nil, fmt.Errorf("No definition found for the opcode %d", opcode)
	}
	return definition, nil
}

// Makes an instruction that is just an array of bytes with the first byte reserved
// for the opcode
func Make(opcode Opcode, operands ...int) []byte {
	// TODO: Maybe not needed to check for definition not found in this method
	var definition = definitions[opcode]
	// var definition, found = definitions[opcode]
	// if !found { return []byte {} }

	// Find out the size of the instruction
	var instructionLength = 1 // Starts with 1 for the opcode
	for _, width := range definition.OperandWidths {
		instructionLength += width
	}

	// Allocate the calculated size for the instruction
	var instruction = make([]byte, instructionLength)
	instruction[0] = byte(opcode) // First byte to the opcode

	var offset = 1 // Offset the opcode
	for i, operand := range operands {
		var width = definition.OperandWidths[i]
		switch width {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(operand))
		}
		offset += width // Offset current operand
	}

	return instruction
}

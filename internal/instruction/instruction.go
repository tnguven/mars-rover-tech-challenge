package instruction

import (
	"fmt"
)

type Instruction int

const (
	Left Instruction = iota
	Right
	Move
)

var instructionMap = map[rune]Instruction{
	'L': Left,
	'R': Right,
	'M': Move,
}

var instructionToChar = map[Instruction]rune{
	Left:  'L',
	Right: 'R',
	Move:  'M',
}

func ParseInstruction(ch rune) (Instruction, error) {
	instr, ok := instructionMap[ch]
	if !ok {
		return 0, fmt.Errorf("invalid instruction: %c", ch)
	}
	return instr, nil
}

func (i Instruction) String() string {
	return string(instructionToChar[i])
}

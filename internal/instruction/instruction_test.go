package instruction_test

import (
   "testing"

   "github.com/stretchr/testify/assert"
   "github.com/tnguven/mars-rover-challenge/internal/instruction"
)

func TestParseInstruction_Valid(t *testing.T) {
   tests := []struct {
      input    rune
      expected instruction.Instruction
   }{
      {'L', instruction.Left},
      {'R', instruction.Right},
      {'M', instruction.Move},
   }

   for _, tt := range tests {
      result, err := instruction.ParseInstruction(tt.input)
      assert.NoError(t, err, "input %c should parse without error", tt.input)
      assert.Equal(t, tt.expected, result, "expected %v for input %c", tt.expected, tt.input)
   }
}

func TestParseInstruction_Invalid(t *testing.T) {
   invalidInputs := []rune{'X', 'A', ' '}

   for _, ch := range invalidInputs {
      _, err := instruction.ParseInstruction(ch)
      assert.Error(t, err, "input %c should return an error", ch)
   }
}

func TestInstructionString(t *testing.T) {
   tests := []struct {
      instruction instruction.Instruction
      expected    string
   }{
      {instruction.Left, "L"},
      {instruction.Right, "R"},
      {instruction.Move, "M"},
   }

   for _, tt := range tests {
      result := tt.instruction.String()
      assert.Equal(t, tt.expected, result, "expected string %q for instruction %v", tt.expected, tt.instruction)
   }
}

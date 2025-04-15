package input_test

import (
   "testing"

   "github.com/stretchr/testify/assert"
   "github.com/tnguven/mars-rover-challenge/internal/input"
)

func TestValidatePlateau_Valid(t *testing.T) {
   x, y, err := input.ValidatePlateau("5 5")
   assert.NoError(t, err)
   assert.Equal(t, 5, x)
   assert.Equal(t, 5, y)
}

func TestValidatePlateau_ExtraSpaces(t *testing.T) {
   x, y, err := input.ValidatePlateau("  5   5  ")
   assert.NoError(t, err)
   assert.Equal(t, 5, x)
   assert.Equal(t, 5, y)
}

func TestValidatePlateau_InvalidFormat(t *testing.T) {
   _, _, err := input.ValidatePlateau("5")
   assert.Error(t, err)
   _, _, err = input.ValidatePlateau("5 5 5")
   assert.Error(t, err)
   _, _, err = input.ValidatePlateau("")
   assert.Error(t, err)
}

func TestValidatePlateau_NegativeOrNonInteger(t *testing.T) {
   _, _, err := input.ValidatePlateau("-1 5")
   assert.Error(t, err)
   _, _, err = input.ValidatePlateau("5 -2")
   assert.Error(t, err)
   _, _, err = input.ValidatePlateau("a b")
   assert.Error(t, err)
}

func TestValidateStartPosition_Valid(t *testing.T) {
   out, err := input.ValidateStartPosition("1 2 N")
   assert.NoError(t, err)
   assert.Equal(t, "1 2 N", out)
}

func TestValidateStartPosition_ExtraSpaces(t *testing.T) {
   out, err := input.ValidateStartPosition(" 1 2 N ")
   assert.NoError(t, err)
   assert.Equal(t, "1 2 N", out)
}

func TestValidateStartPosition_InvalidFormat(t *testing.T) {
   _, err := input.ValidateStartPosition("1 2")
   assert.Error(t, err)
   _, err = input.ValidateStartPosition("1 2 N X")
   assert.Error(t, err)
   _, err = input.ValidateStartPosition("")
   assert.Error(t, err)
}

func TestValidateStartPosition_InvalidValues(t *testing.T) {
   _, err := input.ValidateStartPosition("a b N")
   assert.Error(t, err)
   _, err = input.ValidateStartPosition("1 2 X")
   assert.Error(t, err)
   _, err = input.ValidateStartPosition("-1 0 S")
   assert.Error(t, err)
}

func TestValidateInstructions_Valid(t *testing.T) {
   out, err := input.ValidateInstructions("LMLMLMLMM")
   assert.NoError(t, err)
   assert.Equal(t, "LMLMLMLMM", out)

   out, err = input.ValidateInstructions("lmr")
   assert.NoError(t, err)
   assert.Equal(t, "LMR", out)
}

func TestValidateInstructions_Whitespace(t *testing.T) {
   out, err := input.ValidateInstructions("  lmr  ")
   assert.NoError(t, err)
   assert.Equal(t, "LMR", out)
}

func TestValidateInstructions_Empty(t *testing.T) {
   _, err := input.ValidateInstructions("")
   assert.Error(t, err)
}

func TestValidateInstructions_InvalidChar(t *testing.T) {
   _, err := input.ValidateInstructions("LMX")
   assert.Error(t, err)
   _, err = input.ValidateInstructions("123")
   assert.Error(t, err)
   _, err = input.ValidateInstructions("!@#")
   assert.Error(t, err)
}

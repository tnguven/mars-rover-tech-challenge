package direction_test

import (
   "testing"

   "github.com/stretchr/testify/assert"
   "github.com/tnguven/mars-rover-challenge/internal/direction"
)

func TestParseDirection_Valid(t *testing.T) {
   tests := []struct {
      input    string
      expected direction.Direction
   }{
      {"N", direction.North},
      {"n", direction.North},
      {"E", direction.East},
      {"e", direction.East},
      {"S", direction.South},
      {"s", direction.South},
      {"W", direction.West},
      {"w", direction.West},
   }

   for _, tt := range tests {
      dir, err := direction.ParseDirection(tt.input)
      assert.NoError(t, err, "expected no error for input %s", tt.input)
      assert.Equal(t, tt.expected, dir, "expected %v for input %s", tt.expected, tt.input)
   }
}

func TestParseDirection_Invalid(t *testing.T) {
   invalidInputs := []string{"A", "X", "", "north"}

   for _, input := range invalidInputs {
      _, err := direction.ParseDirection(input)
      assert.Error(t, err, "expected error for input %q", input)
   }
}

func TestDirectionString(t *testing.T) {
   tests := []struct {
      direction direction.Direction
      expected  string
   }{
      {direction.North, "N"},
      {direction.East, "E"},
      {direction.South, "S"},
      {direction.West, "W"},
      {direction.Direction(10), "UNKNOWN"},
   }

   for _, tt := range tests {
      result := tt.direction.String()
      assert.Equal(t, tt.expected, result, "expected string representation of %v to be %q", tt.direction, tt.expected)
   }
}

func TestTurnLeft(t *testing.T) {
   tests := []struct {
      current  direction.Direction
      expected direction.Direction
   }{
      {direction.North, direction.West},
      {direction.West, direction.South},
      {direction.South, direction.East},
      {direction.East, direction.North},
   }

   for _, tt := range tests {
      result := tt.current.TurnLeft()
      assert.Equal(t, tt.expected, result, "TurnLeft on %v should result in %v", tt.current, tt.expected)
   }
}

func TestTurnRight(t *testing.T) {
   tests := []struct {
      current  direction.Direction
      expected direction.Direction
   }{
      {direction.North, direction.East},
      {direction.East, direction.South},
      {direction.South, direction.West},
      {direction.West, direction.North},
   }

   for _, tt := range tests {
      result := tt.current.TurnRight()
      assert.Equal(t, tt.expected, result, "TurnRight on %v should result in %v", tt.current, tt.expected)
   }
}

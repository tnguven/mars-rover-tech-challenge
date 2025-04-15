package rover_test

import (
   "testing"

   "github.com/stretchr/testify/assert"
   "github.com/tnguven/mars-rover-challenge/internal/direction"
   "github.com/tnguven/mars-rover-challenge/internal/instruction"
   "github.com/tnguven/mars-rover-challenge/internal/plateau"
   "github.com/tnguven/mars-rover-challenge/internal/rover"
)

func TestNewRover(t *testing.T) {
   p := plateau.New(5, 5)
   r := rover.New(1, 2, direction.North, p)

   assert.NotNil(t, r)
   assert.Equal(t, 1, r.X)
   assert.Equal(t, 2, r.Y)
   assert.Equal(t, direction.North, r.Dir)
   assert.Equal(t, p, r.Plateau)
}

func TestTurnLeft(t *testing.T) {
   p := plateau.New(5, 5)
   r := rover.New(1, 2, direction.North, p)

   r.TurnLeft()
   assert.Equal(t, direction.West, r.Dir, "turning left from North should yield West")

   r.TurnLeft()
   assert.Equal(t, direction.South, r.Dir, "turning left from West should yield South")
}

func TestTurnRight(t *testing.T) {
   p := plateau.New(5, 5)
   r := rover.New(1, 2, direction.North, p)

   r.TurnRight()
   assert.Equal(t, direction.East, r.Dir, "turning right from North should yield East")

   r.TurnRight()
   assert.Equal(t, direction.South, r.Dir, "turning right from East should yield South")
}

func TestMove(t *testing.T) {
   p := plateau.New(5, 5)
   r := rover.New(1, 1, direction.North, p)
   err := r.Move()
   assert.NoError(t, err, "move should not return an error when moving within bounds")

   assert.Equal(t, 1, r.X)
   assert.Equal(t, 2, r.Y)

   r.Dir = direction.East
   err = r.Move()
   assert.NoError(t, err)
   assert.Equal(t, 2, r.X)
   assert.Equal(t, 2, r.Y)
}

func TestMoveOutOfBounds(t *testing.T) {
   p := plateau.New(2, 2)
   r := rover.New(2, 2, direction.North, p)
   err := r.Move()
   assert.Error(t, err, "moving out of bounds should return an error")
   assert.Contains(t, err.Error(), "movement out of bounds", "error message should indicate out-of-bounds movement")
}

func TestExecuteInstruction(t *testing.T) {
   p := plateau.New(5, 5)
   r := rover.New(1, 1, direction.North, p)

   err := r.ExecuteInstruction(instruction.Left)
   assert.NoError(t, err)
   assert.Equal(t, direction.West, r.Dir, "after turning left from North, rover should face West")

   err = r.ExecuteInstruction(instruction.Right)
   assert.NoError(t, err)
   assert.Equal(t, direction.North, r.Dir, "after turning right from West, rover should face North again")

   err = r.ExecuteInstruction(instruction.Move)
   assert.NoError(t, err)
   assert.Equal(t, 1, r.X)
   assert.Equal(t, 2, r.Y)
}

func TestExecuteInstructionUnknown(t *testing.T) {
   p := plateau.New(5, 5)
   r := rover.New(1, 1, direction.North, p)

   unknownInstr := instruction.Instruction(99)
   err := r.ExecuteInstruction(unknownInstr)
   assert.Error(t, err, "executing an unknown instruction should return an error")
   assert.Contains(t, err.Error(), "unknown instruction", "error message should mention unknown instruction")
}

func TestRoverString(t *testing.T) {
   p := plateau.New(5, 5)
   r := rover.New(1, 2, direction.North, p)
   s := r.String()
   expected := "1 2 N"
   assert.Equal(t, expected, s, "rover's string representation should match expected output")
}

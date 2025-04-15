package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tnguven/mars-rover-challenge/internal/direction"
	"github.com/tnguven/mars-rover-challenge/internal/instruction"
	"github.com/tnguven/mars-rover-challenge/internal/parser"
)

func TestParseMissionInput(t *testing.T) {
	input := `
5 5
1 2 N
LMLMLMLMM
3 3 E
MMRMMRMRRM
`

	mission := parser.Parse(input)

	assert.Equal(t, 5, mission.PlateauMaxX, "expected PlateauMaxX to be 5")
	assert.Equal(t, 5, mission.PlateauMaxY, "expected PlateauMaxY to be 5")

	assert.Len(t, mission.Rovers, 2, "expected 2 rover inputs")

	rover1 := mission.Rovers[0]
	assert.Equal(t, 1, rover1.StartPosition.X, "first rover start X should be 1")
	assert.Equal(t, 2, rover1.StartPosition.Y, "first rover start Y should be 2")
	assert.Equal(t, direction.North, rover1.StartDirection, "first rover start direction should be North")
	expectedInsts1 := []instruction.Instruction{
		instruction.Left,
		instruction.Move,
		instruction.Left,
		instruction.Move,
		instruction.Left,
		instruction.Move,
		instruction.Left,
		instruction.Move,
		instruction.Move,
	}
	assert.Equal(t, expectedInsts1, rover1.Instructions, "first rover instructions do not match expected")

	rover2 := mission.Rovers[1]
	assert.Equal(t, 3, rover2.StartPosition.X, "second rover start X should be 3")
	assert.Equal(t, 3, rover2.StartPosition.Y, "second rover start Y should be 3")
	assert.Equal(t, direction.East, rover2.StartDirection, "second rover start direction should be East")
	expectedInsts2 := []instruction.Instruction{
		instruction.Move,
		instruction.Move,
		instruction.Right,
		instruction.Move,
		instruction.Move,
		instruction.Right,
		instruction.Move,
		instruction.Right,
		instruction.Right,
		instruction.Move,
	}
	assert.Equal(t, expectedInsts2, rover2.Instructions, "second rover instructions do not match expected")
}

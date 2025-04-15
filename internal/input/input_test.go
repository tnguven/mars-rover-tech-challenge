package input_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tnguven/mars-rover-challenge/internal/input"
)

func TestReadPlateau(t *testing.T) {
	inputStream := bytes.NewBufferString("invalid\n5 -5\n5 5\n")
	x, y := input.ReadPlateau(inputStream)
	assert.Equal(t, 5, x)
	assert.Equal(t, 5, y)
}

func TestReadStartPosition(t *testing.T) {
	inputStream := bytes.NewBufferString("wrong format\n1 x N\n1 2 N\n")
	result := input.ReadStartPosition(inputStream)
	assert.Equal(t, "1 2 N", result)
}

func TestReadInstructions(t *testing.T) {
	inputStream := bytes.NewBufferString("badinput\n123\nLMX\nLMLMLMLMM\n")
	result := input.ReadInstructions(inputStream)
	assert.Equal(t, "LMLMLMLMM", result)
}

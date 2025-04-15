package app_test

import (
   "testing"

   "github.com/stretchr/testify/assert"
   "github.com/tnguven/mars-rover-challenge/internal/app"
)

func TestRunMission_Success(t *testing.T) {
   input := `
5 5
1 2 N
LMLMLMLMM
3 3 E
MMRMMRMRRM
`

   rovers, err := app.RunMission(input)
   assert.NoError(t, err)
   assert.Len(t, rovers, 2)

   expected := []string{
      "1 3 N",
      "5 1 E",
   }

   for i, r := range rovers {
      assert.Equal(t, expected[i], r.String(), "rover %d final state mismatch", i+1)
   }
}

func TestRunMission_OutOfBoundsInstructionsAreIgnored(t *testing.T) {
   input := `
2 2
0 0 N
MMMM
`

   rovers, err := app.RunMission(input)
   assert.NoError(t, err)
   assert.Len(t, rovers, 1)
   assert.Equal(t, "0 2 N", rovers[0].String())
}

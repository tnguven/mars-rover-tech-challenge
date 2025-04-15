package parser

import (
   "strconv"
   "strings"

   "github.com/tnguven/mars-rover-challenge/internal/direction"
   "github.com/tnguven/mars-rover-challenge/internal/instruction"
)

type Position struct {
   X int
   Y int
}

type MissionInput struct {
   PlateauMaxX int
   PlateauMaxY int
   Rovers      []RoverInput
}

type RoverInput struct {
   StartPosition  Position
   StartDirection direction.Direction
   Instructions   []instruction.Instruction
}

// Parse input already validated
func Parse(input string) *MissionInput {
   lines := strings.Split(strings.TrimSpace(input), "\n")

   plateauParts := strings.Fields(lines[0])
   maxX, _ := strconv.Atoi(plateauParts[0])
   maxY, _ := strconv.Atoi(plateauParts[1])
   mission := &MissionInput{
      PlateauMaxX: maxX,
      PlateauMaxY: maxY,
   }

   for i := 1; i < len(lines); i += 2 {
      startParts := strings.Fields(strings.TrimSpace(lines[i]))
      x, _ := strconv.Atoi(startParts[0])
      y, _ := strconv.Atoi(startParts[1])
      dir, _ := direction.ParseDirection(strings.ToUpper(startParts[2]))

      instructionsLine := strings.TrimSpace(lines[i+1])
      var insts []instruction.Instruction
      for _, ch := range instructionsLine {
         inst, _ := instruction.ParseInstruction(ch)
         insts = append(insts, inst)
      }

      mission.Rovers = append(mission.Rovers, RoverInput{
         StartPosition:  Position{X: x, Y: y},
         StartDirection: dir,
         Instructions:   insts,
      })
   }

   return mission
}

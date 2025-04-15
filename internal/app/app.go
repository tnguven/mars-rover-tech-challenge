package app

import (
   "log"

   "github.com/tnguven/mars-rover-challenge/internal/parser"
   "github.com/tnguven/mars-rover-challenge/internal/plateau"
   "github.com/tnguven/mars-rover-challenge/internal/rover"
)

// RunMission input already validated
func RunMission(input string) ([]*rover.Rover, error) {
   mission := parser.Parse(input)

   p := plateau.New(mission.PlateauMaxX, mission.PlateauMaxY)
   var results []*rover.Rover

   for _, rInput := range mission.Rovers {
      rvr := rover.New(rInput.StartPosition.X, rInput.StartPosition.Y, rInput.StartDirection, p)

      for _, instr := range rInput.Instructions {
         if err := rvr.ExecuteInstruction(instr); err != nil {
            log.Printf("error executing instruction '%s': %v", instr, err)
            // continue to next instruction
            continue
         }
      }
      results = append(results, rvr)
   }

   return results, nil
}

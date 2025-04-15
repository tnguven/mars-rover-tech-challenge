package input

import (
   "fmt"
   "strconv"
   "strings"

   "github.com/tnguven/mars-rover-challenge/internal/direction"
   "github.com/tnguven/mars-rover-challenge/internal/instruction"
)

const (
   PlateauGridVarSize   = 2
   StartPositionVarSize = 3
)

func ValidatePlateau(input string) (int, int, error) {
   parts := strings.Fields(strings.TrimSpace(input))
   if len(parts) != PlateauGridVarSize {
      return 0, 0, fmt.Errorf("invalid format: must be 'X Y'")
   }

   x, err1 := strconv.Atoi(parts[0])
   y, err2 := strconv.Atoi(parts[1])
   if err1 != nil || err2 != nil || x < 0 || y < 0 {
      return 0, 0, fmt.Errorf("invalid coordinates: must be non-negative integers")
   }

   return x, y, nil
}

func ValidateStartPosition(input string) (string, error) {
   parts := strings.Fields(strings.TrimSpace(input))
   if len(parts) != StartPositionVarSize {
      return "", fmt.Errorf("invalid format: must be 'X Y D'")
   }

   x, errX := strconv.Atoi(parts[0])
   y, errY := strconv.Atoi(parts[1])
   dir := strings.ToUpper(parts[2])
   _, errDir := direction.ParseDirection(dir)

   if errX != nil || errY != nil || errDir != nil {
      return "", fmt.Errorf("invalid values: coordinates must be positive integers and direction must be N, E, S, or W")
   }

   if x < 0 || y < 0 {
      return "", fmt.Errorf("invalid coordinates: coordinates must be non-negative integers")
   }

   return fmt.Sprintf("%s %s %s", parts[0], parts[1], dir), nil
}

func ValidateInstructions(input string) (string, error) {
   line := strings.ToUpper(strings.TrimSpace(input))

   if len(line) == 0 {
      return "", fmt.Errorf("instructions cannot be empty")
   }

   for _, ch := range line {
      if _, err := instruction.ParseInstruction(ch); err != nil {
         return "", fmt.Errorf("invalid instruction: %c", ch)
      }
   }

   return line, nil
}

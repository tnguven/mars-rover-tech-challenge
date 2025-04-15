package direction

import (
   "fmt"
   "strings"
)

type Direction int

const (
   North Direction = iota
   East
   South
   West
)

const (
   directionCount = 4
   leftCount      = directionCount - 1
   rightCount     = 1
)

var directionMap = map[string]Direction{
   "N": North,
   "E": East,
   "S": South,
   "W": West,
}

var directionToString = map[Direction]string{
   North: "N",
   East:  "E",
   South: "S",
   West:  "W",
}

func ParseDirection(s string) (Direction, error) {
   dir, ok := directionMap[strings.ToUpper(s)]
   if !ok {
      return 0, fmt.Errorf("invalid direction: %s", s)
   }
   return dir, nil
}

func (d Direction) String() string {
   str, ok := directionToString[d]
   if !ok {
      return "UNKNOWN"
   }
   return str
}

func (d Direction) TurnLeft() Direction {
   return (d + leftCount) % directionCount
}

func (d Direction) TurnRight() Direction {
   return (d + rightCount) % directionCount
}

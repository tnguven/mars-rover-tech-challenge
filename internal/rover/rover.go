package rover

import (
   "fmt"

   "github.com/tnguven/mars-rover-challenge/internal/direction"
   "github.com/tnguven/mars-rover-challenge/internal/instruction"
   "github.com/tnguven/mars-rover-challenge/internal/plateau"
)

type Rover struct {
   X       int
   Y       int
   Dir     direction.Direction
   Plateau *plateau.Plateau
}

func New(x, y int, d direction.Direction, p *plateau.Plateau) *Rover {
   return &Rover{
      X:       x,
      Y:       y,
      Dir:     d,
      Plateau: p,
   }
}

func (r *Rover) TurnLeft() {
   r.Dir = r.Dir.TurnLeft()
}

func (r *Rover) TurnRight() {
   r.Dir = r.Dir.TurnRight()
}

func (r *Rover) Move() error {
   newX := r.X
   newY := r.Y

   switch r.Dir {
   case direction.North:
      newY = r.Y + 1
   case direction.East:
      newX = r.X + 1
   case direction.South:
      newY = r.Y - 1
   case direction.West:
      newX = r.X - 1
   }

   if !r.Plateau.IsinBounds(newX, newY) {
      return fmt.Errorf("movement out of bounds: trying to move to (%d, %d)", newX, newY)
   }

   r.X = newX
   r.Y = newY

   return nil
}

func (r *Rover) ExecuteInstruction(instr instruction.Instruction) error {
   switch instr {
   case instruction.Left:
      r.TurnLeft()
   case instruction.Right:
      r.TurnRight()
   case instruction.Move:
      return r.Move()
   default:
      return fmt.Errorf("unknown instruction")
   }
   return nil
}

func (r *Rover) String() string {
   return fmt.Sprintf("%d %d %s", r.X, r.Y, r.Dir.String())
}

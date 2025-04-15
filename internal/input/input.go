package input

import (
   "bufio"
   "fmt"
   "io"
   "log"
)

func ReadPlateau(reader io.Reader) (int, int) {
   scanner := bufio.NewScanner(reader)

   for {
      fmt.Print("Enter plateau size (e.g. '5 5'): ")
      if !scanner.Scan() {
         log.Println("failed to read input")
         continue
      }

      x, y, err := ValidatePlateau(scanner.Text())
      if err != nil {
         log.Printf("parse plateau errorr: %s", err)
         continue
      }

      return x, y
   }
}

func ReadStartPosition(reader io.Reader) string {
   scanner := bufio.NewScanner(reader)

   for {
      fmt.Print("Enter rover start position (e.g. '1 2 N'): ")
      if !scanner.Scan() {
         log.Println("Failed to read input")
         continue
      }

      pos, err := ValidateStartPosition(scanner.Text())
      if err != nil {
         log.Printf("invalid start position error: %s", err)
         continue
      }

      return pos
   }
}

func ReadInstructions(reader io.Reader) string {
   scanner := bufio.NewScanner(reader)

   for {
      fmt.Print("Enter movement instructions (e.g. 'LMLMLMLMM'): ")
      if !scanner.Scan() {
         log.Println("failed to read input")
         continue
      }

      instructions, err := ValidateInstructions(scanner.Text())
      if err != nil {
         log.Printf("parse instructions error: %s", err)
         continue
      }

      return instructions
   }
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tnguven/mars-rover-challenge/internal/app"
	"github.com/tnguven/mars-rover-challenge/internal/input"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var inputBuilder strings.Builder

	x, y := input.ReadPlateau(reader)
	inputBuilder.WriteString(fmt.Sprintf("%d %d\n", x, y))

	for {
		start := input.ReadStartPosition(reader)
		instr := input.ReadInstructions(reader)

		inputBuilder.WriteString(start + "\n")
		inputBuilder.WriteString(instr + "\n")

		fmt.Print("Add another rover? (y/n): ")
		another, _ := reader.ReadString('\n')
		if strings.ToLower(strings.TrimSpace(another)) != "y" {
			break
		}
	}

	results, err := app.RunMission(inputBuilder.String())
	if err != nil {
		log.Fatal(err)
	}

	for _, rover := range results {
		fmt.Println(rover.String())
	}
}

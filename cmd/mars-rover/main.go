package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter plateau size (e.g '5 5'): ")
	plateauSize, _ := reader.ReadString('\n')

	var inputBuilder strings.Builder
	inputBuilder.WriteString(strings.TrimSpace(plateauSize) + "\n")

	for {
		fmt.Print("Enter rover start position (e.g '1 2 N'): ")
		position, _ := reader.ReadString('\n')
		inputBuilder.WriteString(strings.TrimSpace(position) + "\n")

		fmt.Print("Enter rover instructions (e.g 'LMLMLMLMM'): ")
		instructions, _ := reader.ReadString('\n')
		inputBuilder.WriteString(strings.TrimSpace(instructions) + "\n")

		fmt.Print("Add another rover? (y/n): ")
		another, _ := reader.ReadString('\n')
		another = strings.ToLower(strings.TrimSpace(another))
		if another != "y" {
			break
		}
	}

	input := inputBuilder.String()

	fmt.Printf("Inputs: %s", input)
}

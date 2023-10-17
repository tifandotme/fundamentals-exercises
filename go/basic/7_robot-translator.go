package exercise

import (
	"fmt"
	"os"
	"strconv"
)

func RobotTranslator() {
	var commands string
	var count int

	fmt.Printf("Input robot commands (combination of R, L, or A): ")
	fmt.Scanln(&commands)

	for idx, command := range commands {
		count++

		isLastCommand := idx == len(commands)-1
		currCommand := string(command)
		var nextCommand string

		// Prevent out of range error. Only get next command if current command is not the last command.
		if !isLastCommand {
			nextCommand = commands[idx+1 : idx+2]
		}

		if isLastCommand || currCommand != nextCommand {
			var translation string

			switch currCommand {
			case "R":
				translation = formatTranslation("right", count)
			case "L":
				translation = formatTranslation("left", count)
			case "A":
				translation = formatTranslation("advance", count)
			default:
				// Assumption: when user input an invalid command, stop program.
				fmt.Printf("Invalid command. Stopping program.\n")
				os.Exit(0)
			}

			fmt.Println(translation)

			count = 0
		}
	}
}

func formatTranslation(direction string, count int) (translation string) {
	translation = fmt.Sprintf("Move %s %s time", direction, strconv.Itoa(count))

	if count > 1 {
		translation += "s"
	}

	return
}

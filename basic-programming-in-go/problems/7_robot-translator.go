package problems

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RobotTranslator() {
	var input string

	fmt.Printf("Input robot commands (combination of R, L, or A): ")
	fmt.Scanln(&input)

	commands := strings.Split(input, "")

	var count int

	for idx, val := range commands {
		count++

		if idx == len(commands)-1 || commands[idx+1] != val {
			var translation string

			switch val {
			case "R":
				translation = "Move right " + strconv.Itoa(count) + " time"
			case "L":
				translation = "Move left " + strconv.Itoa(count) + " time"
			case "A":
				translation = "Move advance " + strconv.Itoa(count) + " time"
			default:
				// Assumption: when user input an invalid command, stop program.
				fmt.Printf("Invalid command. Stopping program.\n")
				os.Exit(0)
			}

			if count > 1 {
				translation += "s"
			}

			fmt.Printf("%s\n", translation)

			count = 0
		}

	}
}

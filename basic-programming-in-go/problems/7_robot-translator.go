/*
Objectives
- Able to use the conditional statements
- Able to use the for statements
- Able to convert a string to a slice
- Able to concatenate a string

Description
A robot factory's test facility needs a program to translate robot commands.

The robots have three possible commands:

R for turn right
L for turn left
A for advance (forward)

For example, a valid commands would look like:

RRAAALA

However, you are a robot that cannot understand this string of characters just by looking at it. You need detailed instructions on how to follow the commands.

Your objective is to write a program to translate the input path to a set of detailed and readable instructions that even a robot like you could understand.

To do this, you must translate the previous example

RRAAALA

to a "line feed separated string" equivalent to:

Move right 2 times
Move advance 3 times
Move left 1 time
Move advance 1 time

Questions to mentors:
1. Is my error handling implementation correct?
*/

package problems

import (
	"fmt"
	"strings"
	"errors"
	"log"
)

func RobotTranslator() {
	var input string

	fmt.Printf("Input robot commands (combination of R, L, or A): ")
	fmt.Scanln(&input)

	commands := strings.Split(input, "")

	var amount int
	for idx, val := range commands {
		amount += 1

		if idx == len(commands)-1 || commands[idx+1] != val {
			switch val {
			case "R":
				fmt.Printf("Move right %d times\n", amount)
			case "L":
				fmt.Printf("Move left %d times\n", amount)
			case "A":
				fmt.Printf("Move advance %d times\n", amount)
			default:
				log.Fatal(errors.New("invalid command"))
			}

			amount = 0
		}
	}
}

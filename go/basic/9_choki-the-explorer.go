/*
Maze Explorer Exercise
Objectives
Able to create functions
Able to use functions (recursion)
Description
Choki is stranded in a maze when searching for treasure. You are tasked to help choki escaped from the 2D Maze. The grid contains the following symbols:

'S': Start point
'E': End point
'.': Open path
'#': Wall
'%': Power Up
'x': Path taken
Your program should guide Choki from the start point to the end point while avoiding walls. Choki can move up, down, left, or right. And he can get power ups which will automatically make Choki escaped. Choki will prioritize escaping first than getting a power ups.

Examples
Case 1
Input

S . . # . . .
# # . # # . #
. . . . . . #
. # # # # . #
. # . . . . E

Output

S x x # . . .
# # x # # . #
. . x x x x #
. # # # # x #
. # . . . x E

Path Found!

Case 2
Input

S . . # . . .
# # . # # . #
. . . . . . #
. # # # # . #
. # . . . . #

Output

S x x # x x x
# # x # # x #
x x x x x x #
x # # # # x #
x # x x x x #

No way to escape!

Case 3
Input

S . . # . . .
# # . # # . #
. . . . . . #
. # # # # . #
. # . % . . #

Output

S x x # . . .
# # x # # . #
. . x x x x #
. # # # # x #
. # . x x x #

Got powerups, Choki Escaped!
*/

package exercise

import (
	"fmt"
	"os"
	"strings"
)

func ChokiTheExplorer() {
	maze := [5][7]string{}

	fmt.Println("Input maze: ")
	for idx, row := range maze {
		var input string
		fmt.Printf("%d: ", idx+1)
		_, err := fmt.Scanln(&input)

		if len(input) != 7 || err != nil {
			fmt.Println("Invalid input. Exactly 7 char is needed")
			os.Exit(0)
		}

		foo := strings.Split(input, "")

		fmt.Println("foo length: ", len(foo))

		fmt.Println(copy(row[:], foo))
	}

	for _, row := range maze {
		for _, val := range row {
			fmt.Print(val)
		}
		fmt.Println()
	}
}

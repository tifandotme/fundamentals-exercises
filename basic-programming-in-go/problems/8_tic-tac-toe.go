/*
Objectives
Able to determine the appropriate data types
Able to use the conditional statements
Able to use the for statements
Able to use the slice data type
Able to create and use functions
Description
In a game of Tic-Tac-Toe, two players take turns marking an available cell in a 3-by-3 grid with their respective tokens (either X or O). When one player has placed three tokens in a horizontal, vertical, or diagonal row on the grid, the game is over and that player has won. A draw (no winner) occurs when all the cells on the grid have been filled with tokens and neither player has achieved a win. More details on the game here : https://en.wikipedia.org/wiki/Tic-tac-toe

Logo

Requirements
Problem statement: Given a string of ‘X’, ‘O’ and ‘-’ as input,

which represents a game board, print the state of the

game as output. The states must be one of:

X wins

O wins

Draw

Game in progress

Invalid grid

Examples
1. X wins - when X occupies either all cells in a vertical line, or all cells in a horizontal line, or all cells in a diagonal.
Sample Input :
XOXXOOXXO

Which represents the board:
| X | O | X |
| X | O | O |
| X | X | O |

Sample output:
X wins!

-----------------
2. O wins - when O occupies either all cells in a vertical line, or all cells in a horizontal line, or all cells in a diagonal.
Sample input:
XOOXOXOXO

Which represents the board:
X	O	O
X	O	X
O	X	O

Sample output:
O Wins!

-----------------
3. Draw - no player has won.
Sample input:
OXOXOXXOX

Which represents the board:

O	X	O
X	O	X
X	O	X

Sample output:
Its a draw!

-----------------
4. Game in progress - if no player has won and its not a draw
Sample input:
XOXX--O--

Which represents the board:

X	O	X
X
O

Sample output:
Game still in progress!

-----------------
5. Invalid Grid - a grid thats not possible to achieve in a real game
Sample Input:
XXXOOOXXO

Which represents the board:

X	X	X
O	O	O
X	X	O

Sample Output:
Invalid game board
*/

package problems

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var winCountX int
var winCountO int

func TicTacToe() {
	marks := promptBoard()

	analyzeXAxis(marks)

	analyzeYAxis(marks)

	analyzeDiagonals(marks)

	printResult()
}

func promptBoard() (marks string) {
	fmt.Print("Enter the board (X, O, or -): ")
	if _, err := fmt.Scanln(&marks); err != nil {
		fmt.Println("Invalid input")
		os.Exit(0)
	}

	isLengthCorrect := len(marks) != 9
	isMarksValid := regexp.MustCompile(`^[XO-]+$`).MatchString(marks)

	if !isMarksValid || isLengthCorrect {
		// Assumption: when user input an invalid mark, stop the program.
		fmt.Println("Invalid marks. Input a total of 9 marks of either X, O, or -")
		os.Exit(0)
	}

	if strings.Contains(marks, "-") {
		// State 4
		fmt.Println("Game still in progress!")
		os.Exit(0)
	}

	return
}

func analyzeXAxis(marks string) {
	for i := 0; i < 3; i++ {
		xAxis := marks[i*3 : (i*3)+3]

		checkLine(xAxis)
	}
}

func analyzeYAxis(marks string) {
	for i := 0; i < 3; i++ {
		yAxis := marks[i:i+1] + marks[i+3:i+4] + marks[i+6:i+7]

		checkLine(yAxis)
	}
}

func analyzeDiagonals(marks string) {
	diagonalAxis := []string{
		marks[:1] + marks[4:5] + marks[8:],
		marks[2:3] + marks[4:5] + marks[6:7],
	}

	for _, val := range diagonalAxis {
		checkLine(val)
	}
}

func printResult() {
	if winCountX > 1 || winCountO > 1 || winCountX >= 1 && winCountO >= 1 {
		// State 5
		fmt.Println("Invalid game board")
		return
	}

	if winCountX == 0 && winCountO == 0 {
		// State 3
		fmt.Println("It's a draw!")
		return
	}

	if winCountX == 1 {
		// State 1
		fmt.Println("X Wins!")
	} else {
		// State 2
		fmt.Println("O Wins!")
	}
}

func checkLine(line string) {
	switch line {
	case "XXX":
		winCountX++
	case "OOO":
		winCountO++
	}
}

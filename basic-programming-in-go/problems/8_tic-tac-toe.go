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
	"strings"
)

func TicTacToe() {
	var input string
	fmt.Printf("Enter the board (X, O, or -): ")
	fmt.Scanln(&input)

	if len(input) != 9 {
		fmt.Println("Input a total of 9 marks!")
		return
	}

	marks := strings.Split(input, "")

	var winCountX int
	var winCountO int

	for i := 0; i < 3; i++ {
		xAxis := marks[i*3 : (i*3)+3]

		switch strings.Join(xAxis, "") {
		case "XXX":
			winCountX++
		case "OOO":
			winCountO++
		}
	}

	for i := 0; i < 3; i++ {
		yAxis := []string{marks[i], marks[i+3], marks[i+6]}

		switch strings.Join(yAxis, "") {
		case "XXX":
			winCountX++
		case "OOO":
			winCountO++
		}
	}

	diagonal := [][]string{
		{marks[0], marks[4], marks[8]},
		{marks[2], marks[4], marks[6]},
	}

	for _, v := range diagonal {
		switch strings.Join(v, "") {
		case "XXX":
			winCountX++
		case "OOO":
			winCountO++
		}
	}

	fmt.Println(winCountX)
	fmt.Println(winCountO)

	if winCountX >= 1 && winCountO >= 1 {
		fmt.Println("Invalid game board")
		return
	}

	if winCountX == 0 && winCountO == 0 {
		fmt.Println("It's a draw!")
		return
	}

	if winCountX == 1 {
		fmt.Println("X Wins!")
	} else {
		fmt.Println("O Wins!")
	}
	// var board [3][3]string

	// for i := 0; i < 3; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		mark := marks[i+(j*3)]

	// 		if mark == "-" {
	// 			// Scenario 1
	// 			fmt.Println("Game still in progress!")
	// 			return
	// 		}

	// 		board[i][j] = mark
	// 	}
	// }
}

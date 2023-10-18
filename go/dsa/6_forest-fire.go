/*
Forest Fire
Description

You are given a task to determine on how long it takes for a fire to reach one specific place in a two-dimensional array.

A fire is marked as X in a two-dimensional array. A fire can travel through trees which marked as -. Each day, the fire will travel through the adjacent trees. However, the fire cannot travel through a rock, whcih marked as R. Your goal is to determine the least number of days for the fire to reach the O.

If there is no possible path for the fire to reach the O, then return -1 in your function. Otherwise, return the least number of days to reach the O.
*/

package exercise

import "fmt"

func ForestFire() {
	// input assumes a 2 dimensional slices of rune with an object of either X, -, or O

	// Example 1
	// forest := [][]rune{
	// 	{'-', '-', '-'},
	// 	{'-', 'X', '-'},
	// 	{'-', '-', 'O'},
	// }

	// Example 2
	// forest := [][]rune{
	// 	{'-', '-', '-'},
	// 	{'-', 'X', 'R'},
	// 	{'-', 'R', 'O'},
	// }

	// Example 3
	forest := [][]rune{
		{'-', '-', '-'},
		{'X', 'R', 'O'},
		{'R', '-', '-'},
		{'-', '-', '-'},
	}

	// forest := [][]rune{
	// 	{'-', '-', '-'},
	// 	{'-', 'X', '-'},
	// 	{'-', 'R', 'R'},
	// 	{'-', '-', 'O'},
	// }

	days := forestFireSimulation(forest, 0, false)

	fmt.Println("Days it took:", days)
}

func forestFireSimulation(forest [][]rune, days int, isOver bool) int {
	if isOver {
		return -1
	}

	printForest(forest, days)

	toBeBurned := make([][]bool, len(forest))
	for idx := range toBeBurned {
		toBeBurned[idx] = make([]bool, len(forest[idx]))
	}

	var isFound bool

out:
	for yAxis, rows := range forest {
		for xAxis, object := range rows {

			if object == 'X' {

				if yAxis != 0 {
					if forest[yAxis-1][xAxis] != 'R' {
						if forest[yAxis-1][xAxis] == 'O' {
							isFound = true
							break out
						}
						toBeBurned[yAxis-1][xAxis] = true
					}
				}

				if yAxis != len(forest)-1 {
					if forest[yAxis+1][xAxis] != 'R' {
						if forest[yAxis+1][xAxis] == 'O' {
							isFound = true
							break out
						}
						toBeBurned[yAxis+1][xAxis] = true
					}
				}

				if xAxis != 0 {
					if forest[yAxis][xAxis-1] != 'R' {
						if forest[yAxis][xAxis-1] == 'O' {
							isFound = true
							break out
						}
						toBeBurned[yAxis][xAxis-1] = true
					}
				}

				if xAxis != len(rows)-1 {
					if forest[yAxis][xAxis+1] != 'R' {
						if forest[yAxis][xAxis+1] == 'O' {
							isFound = true
							break out
						}
						toBeBurned[yAxis][xAxis+1] = true
					}
				}
			}
		}
	}

	isOver = true
	for yAxis, rows := range forest {
		for xAxis, object := range rows {
			if toBeBurned[yAxis][xAxis] && object != 'X' {
				forest[yAxis][xAxis] = 'X'
				isOver = false
			}
		}
	}

	if isFound {
		printForest(forest, days+1)
		return days + 1
	}

	return forestFireSimulation(forest, days+1, isOver)
}

func printForest(forest [][]rune, day int) {
	fmt.Printf("Day %d\n", day)
	for _, rows := range forest {
		for idx, object := range rows {
			fmt.Printf("%s", string(object))

			if idx != len(rows)-1 {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}

	fmt.Println()
}

/*
Printer Task (LIFO)
Description
Given printer with max task capacity is 5, if all the capacity is filled then there is a new task so the new task will be the first task to be print.

Example
Input
current: ["newspaper", "cartoon", "fabel", "book", "tabloid"]
next: ["catalog"]

Output
["newspaper", "cartoon", "fabel", "book", "catalog"]
*/

package exercise

import "fmt"

func lifo(input []string, next []string) []string {
	ceiling := 0

	if len(next) <= len(input) {
		ceiling = len(input) - len(next)
	}

	input = input[:ceiling]

	input = append(input, next[:5]...)

	return input
}

func PrinterTask() {
	input := []string{"newspaper", "cartoon", "fabel", "book", "tabloid"}

	next := []string{"catalog", "one", "two", "three", "four", "five"}

	output := lifo(input, next)

	fmt.Println(output)
}

/*
All You Can Eat Ticketing (FIFO)

Description
Given a restaurant with 10 tables inside and everyone has limited time to eat so the schema will be First in First Out (FIFO), if all of tables is filled and there is a new person come to eat so the first one will be changed by the new one. You can ignore the limit time, just consider the FIFO schema.

Example
Input
current: ["Dani", "Rian", "Toni", "Loni", "Poki", "Vali", "Juli", "Kris", "Veny", "Luis"]
next: ["Edward", "Same"]

Output
["Toni", "Loni", "Poki", "Vali", "Juli", "Kris", "Veny", "Luis", "Edward", "Same",]
*/

package exercise

import "fmt"

// func fifo(input *[]string, next []string) {
// 	for _, v := range next {
// 		*input = append(*input, v)
// 		*input = (*input)[1:]
// 	}
// }

func fifo(input []string, next []string) []string {
	for idx, val := range next {
		input = append(input, val)

		if idx != len(input)-1 {
			input = input[1:]
		}
	}

	return input
}

func AllYouCanEat() {
	input := []string{"Dani", "Rian", "Toni", "Loni", "Poki", "Vali", "Juli", "Kris", "Veny", "Luis"}

	next := []string{"Edward", "Same"}

	output := fifo(input, next)

	fmt.Println(output)
}

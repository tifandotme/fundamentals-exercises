/*
Objectives
- Able to use basic operators
- Able to use conditional statements
- Able to use slices

Description
Imagine you are a software developer who has been tasked with integrating a third-party API into your company's system. This API is crucial for your project, but it has a unique way of communicating errors. Instead of conventional error messages, it sends integer values, each representing a specific error.

You've been given the following mappings of error codes to their meanings:

Code	Description
0			No Error
1			Incorrect input
2			The server encounters internal error
4			The server is overloaded by too much traffic
8			You are not authorized to proceed with the input

To complicate matters, the API might combine error codes to represent multiple problems occurring at the same time.

The output should be an array (or a slice in Go) of error messages.

Examples
Input: 0
Output: ["No Error"]

Input: 8
Output: ["You are not authorized to proceed with the input"]

Input: 3
Output: ["Incorrect input", "The server encounters internal error"]

Questions to mentor:
1. Why do we need to use slice? not map
*/

package problems

import "fmt"

type errorCodes struct {
	code        int
	description string
}

var errors = []errorCodes{
	{0, "No Error"},
	{1, "Incorrect input"},
	{2, "The server encounters internal error"},
	{4, "The server is overloaded by too much traffic"},
	{8, "You are not authorized to proceed with the input"},
}

func ApiErrorCode() {

	// Example with map:
	// errors := map[string]struct {
	// 	{0, "No Error"},
	// 	{1, "Incorrect input"},
	// 	{2, "The server encounters internal error"},
	// 	{4, "The server is overloaded by too much traffic"},
	// 	{8, "You are not authorized to proceed with the input"},
	// }

	var input uint8
	var result []string

	fmt.Print("Input: ")
	fmt.Scanln(&input)

	if input == 0 {
		result = append(result, accessErrorCodes(0))
	}
	if input >= 8 {
		result = append(result, accessErrorCodes(8))
		input -= 8
	}
	if input >= 4 {
		result = append(result, accessErrorCodes(4))
		input -= 4
	}
	if input >= 2 {
		result = append(result, accessErrorCodes(2))
		input -= 2
	}
	if input >= 1 {
		result = append(result, accessErrorCodes(1))
		input -= 1
	}

	fmt.Printf("%+q\n", result)
}

func accessErrorCodes(code int) (description string) {
	for _, v := range errors {
		if v.code == code {
			description = v.description
		}
	}
	return
}

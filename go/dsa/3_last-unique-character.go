/*
Last Unique Character

Description
You are tasked to find the last unique character from a string. The input will be lowercase character (a-z) only.

Examples

Example 1
jakarta

If we count the characters:

j -> 1
a -> 3
k -> 1
r -> 1
t -> 1

As we can see, there are several unique characters such as j, k, r, and t. Since we want the last unique character, our function needs to return t.

Example 2
tartar

If we count the characters:

t -> 2
a -> 2
r -> 2
As we can see, there is no unique characters. Therefore, our function returns null (ASCII 0).
*/

package exercise

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func LastUniqueCharacter() {
	input := promptInput()

	charCount := map[string]int{}

	for _, char := range strings.ToLower(input) {
		charCount[string(char)]++
	}

	uniqueChar := findLastUniqueChar(input, charCount)

	fmt.Printf("Last unique character for '%s' is %x (%s)\n", input, uniqueChar, string(uniqueChar))
}

func promptInput() (input string) {
	fmt.Print("Input: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Invalid input")
		os.Exit(0)
	}

	isValid := regexp.MustCompile(`^[a-z]+$`).MatchString(input)

	if !isValid {
		fmt.Println("Invalid input")
		os.Exit(0)
	}

	return
}

func findLastUniqueChar(input string, charCount map[string]int) byte {
	for i := len(input) - 1; i >= 0; i-- {
		char := byte(input[i])

		if charCount[string(char)] == 1 {
			return char
		}
	}

	return 0
}

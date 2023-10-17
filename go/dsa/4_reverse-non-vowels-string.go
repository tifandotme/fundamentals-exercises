/*
Reverse the Non-Vowels of a String
Description

String is somewhat a form of array of bytes. In Go, we can represent a string as an array of byte. This is sufficient for ASCII text.

You are tasked to reverse of the non-vowels of a string.

Let's say that vowels are a, i, u, e, o. The input which will be given to you will be only characters from a-z or A-Z. Please preserve the uppercase or lowercase of the character.

Examples

Example 1

Input: Ambulance

Output: Acnulabme

Example 2

Input: Five

Output: viFe

Example 3

Input: Example

Output: Elapmxe

Example 4

Input: Build

Output: duilB
*/

package exercise

import (
	"fmt"
	"os"
	"regexp"

	"slices"
)

func ReverseNonVowelsOfString() {
	input := promptInput2()

	nonVowels := []rune{}

	for _, char := range input {
		if !isVowel(char) {
			nonVowels = append(nonVowels, char)
		}
	}

	slices.Reverse(nonVowels)

	result := []rune{}

	for _, char := range input {
		if !isVowel(char) {
			result = append(result, nonVowels[0])
			nonVowels = nonVowels[1:]
		} else {
			result = append(result, char)
		}
	}

	fmt.Println("Output:", string(result))
}

func isVowel(char rune) bool {
	return regexp.MustCompile(`a|i|u|e|o|A|I|U|E|O`).MatchString(string(char))
}

func promptInput2() (input string) {
	fmt.Print("Input: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Invalid input")
		os.Exit(0)
	}

	isValid := regexp.MustCompile(`^[A-Za-z]+$`).MatchString(input)

	if !isValid {
		fmt.Println("Invalid input")
		os.Exit(0)
	}

	return
}

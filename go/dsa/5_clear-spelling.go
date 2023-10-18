/*
Clear Spelling
Description

One of our friends encountered a problem when he was asked for his address on the telephone. Sometimes, the other person thinks he said n when actually he said m. One day, he found out that there is something called NATO phonetic alphabet:

The (International) Radiotelephony Spelling Alphabet, commonly known as the NATO phonetic alphabet, is the most widely used set of clear code words for communicating the letters of the Roman alphabet...

...A spelling alphabet is used to spell parts of a message containing letters and numbers to avoid confusion, because many letters sound similar, for instance "n" and "m" or "f" and "s"; the potential for confusion increases if static or other interference is present.

Source: Nato Phonetic Alphabet

Now, he asks you to transform alphabet characters in a string (a-z A-Z) to into the NATO phonetic alphabet. The input will be only alphabetical characters and space.

Here is the table for the letter code words with pronunciation:

NATO phonetic alphabet
If there is a space between words, show Space.

Examples

Example 1 Input: Budi Output: Bravo Uniform Delta India

Example 2 Input: ABC Output: Alfa Bravo Charlie

Example 3 Input: Ini Budi Output: India November India Space Bravo Uniform Delta India
*/

package exercise

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var natoDict = map[rune]string{
	'A': "Alfa",
	'B': "Bravo",
	'C': "Charlie",
	'D': "Delta",
	'E': "Echo",
	'F': "Foxtrot",
	'G': "Golf",
	'H': "Hotel",
	'I': "India",
	'J': "Juliett",
	'K': "Kilo",
	'L': "Lima",
	'M': "Mike",
	'N': "November",
	'O': "Oscar",
	'P': "Papa",
	'Q': "Quebec",
	'R': "Romeo",
	'S': "Sierra",
	'T': "Tango",
	'U': "Uniform",
	'V': "Victor",
	'W': "Whiskey",
	'X': "X-ray",
	'Y': "Yankee",
	'Z': "Zulu",
}

func ClearSpelling() {
	input := promptInput3()

	for idx, char := range strings.ToUpper(input) {
		if char == ' ' {
			fmt.Printf("Space")
		} else {
			fmt.Printf("%s", natoDict[char])
		}

		if idx != len(input)-1 {
			fmt.Printf(" ")
		} else {
			fmt.Printf("\n")
		}
	}
}

func promptInput3() (input string) {
	fmt.Print("Input: ")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Invalid input")
		os.Exit(0)
	}

	input = strings.TrimSpace(input)

	if !regexp.MustCompile(`^[A-Za-z ]+$`).MatchString(input) {
		fmt.Println("Invalid input")
		os.Exit(0)
	}

	return
}

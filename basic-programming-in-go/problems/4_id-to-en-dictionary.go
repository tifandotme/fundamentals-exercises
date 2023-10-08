/*
Objectives
- Able to use the maps data type
- Able to use for-range statements
- Able to handle an error
- Able to use sort from standard library

Description
To upgrade our Digitalent speaking skill, we would like to make an application that can translate an Indonesia word to English.

Questions to mentor:
1. Shouldn't the excercise title be Indonesia to English Dictionary CLI?
*/

package problems

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"sort"
	"strings"
)

// Dictionary with ID as key and EN as value
var dictionary = map[string]string{
	"makan": "eat",
	"tidur": "sleep",
}

func IndonesiaToEnglishDictionary() {
forLoop:
	for {
		fmt.Printf(`ID to EN Dictionary
Menu:
1. Translate
2. Add word
3. Remove word
4. Print dictionary
0. Exit`)

		var input int8

		fmt.Printf("\nInput (number): ")
		fmt.Scanln(&input)

		switch input {
		case 1: // Translate
			var word string

			fmt.Printf("Word to translate (ID): ")
			fmt.Scanln(&word)

			val, ok := dictionary[word]

			clearScreen()
			if ok {
				fmt.Printf("ID: %s\n", titleCase(word))
				fmt.Printf("EN: %s\n\n", titleCase(val))
			} else {
				fmt.Printf("Sorry, \"%s\" is not found in dictionary\n\n", word)
			}

		case 2: // Add word
			var word string

			fmt.Printf("Word to be added: ")
			fmt.Scanln(&word)

			// Trim leading and trailing whitespace and convert to lowercase
			word = strings.TrimSpace(strings.ToLower(word))

			// Index of substring "#", if not found index will be -1
			hashIndex := strings.Index(word, "#")

			// When "#" is not present, break out of switch
			if hashIndex == -1 {
				clearScreen()
				fmt.Printf("Word must be separated with # char\n\n")
				break
			}

			id := word[:hashIndex]
			en := word[hashIndex+1:]

			dictionary[id] = en

			clearScreen()
			fmt.Printf("New word succesfully added\n\n")

		case 3: // Remove word
			var word string

			fmt.Printf("Word to be removed (ID): ")
			fmt.Scanln(&word)

			clearScreen()
			if _, ok := dictionary[word]; ok {
				delete(dictionary, word)
				fmt.Printf("\"%s\" has been removed\n\n", word)
			} else {
				fmt.Printf("Sorry, \"%s\" is not found in dictionary\n\n", word)
			}

		case 4: // Print dictionary
			// Sort dictionary by key
			keys := make([]string, 0, len(dictionary))
			for k := range dictionary {
				keys = append(keys, k)
			}
			sort.Strings(keys)

			clearScreen()
			fmt.Println(`ID to EN Dictionary:`)
			for idx, k := range keys {
				fmt.Printf("%d. %s: %s\n", idx+1, titleCase(k), titleCase(dictionary[k]))
			}
			fmt.Println("")

		case 0: // Exit
			break forLoop

		default:
			clearScreen()
			fmt.Printf("Invalid input\n\n")
		}
	}
}

// Convert string to title case
func titleCase(str string) string {
	caser := cases.Title(language.AmericanEnglish)

	return caser.String(str)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

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
	"os"
	"sort"
	"strings"
)

func IndonesiaToEnglishDictionary() {
	// Dictionary with ID as key and EN as value
	var dictionary = map[string]string{
		"makan": "eat",
		"tidur": "sleep",
	}

	for {
		fmt.Printf(`ID to EN Dictionary
Menu:
1. Translate
2. Add word
3. Remove word
4. Print dictionary
0. Exit`)

		var input uint8

		fmt.Printf("\nInput (number): ")
		fmt.Scanln(&input)

		switch input {
		case 1: // Translate
			var word string

			fmt.Printf("Word to translate (ID): ")
			fmt.Scanln(&word)

			clearScreen()

			if val, ok := dictionary[word]; ok {
				fmt.Printf("ID: %s\n", convertToTitleCase(word))
				fmt.Printf("EN: %s\n\n", convertToTitleCase(val))
			} else {
				fmt.Printf("Sorry, \"%s\" is not found in dictionary\n\n", word)
			}

		case 2: // Add word
			var word string

			fmt.Printf("Word to be added: ")
			fmt.Scanln(&word)

			clearScreen()

			// Trim leading and trailing whitespace and convert to lowercase
			word = strings.TrimSpace(strings.ToLower(word))

			// Index of substring "#", if not found index will be -1
			hashIndex := strings.Index(word, "#")

			// When "#" is not present, break out of switch
			if hashIndex == -1 {
				fmt.Printf("Word must be separated with # char\n\n")
				break
			}

			id := word[:hashIndex]
			en := word[hashIndex+1:]

			if _, ok := dictionary[id]; ok {
				fmt.Printf("Cannot add existing word \"%s\"\n\n", id)
				break
			}

			dictionary[id] = en
			fmt.Printf("New word succesfully added\n\n")

		case 3: // Remove word
			var word string

			fmt.Printf("Word to be removed (ID): ")
			fmt.Scanln(&word)

			clearScreen()

			if _, ok := dictionary[word]; !ok {
				fmt.Printf("Sorry, \"%s\" is not found in dictionary\n\n", word)
				break
			}

			delete(dictionary, word)
			fmt.Printf("\"%s\" has been removed\n\n", word)

		case 4: // Print dictionary
			keys := sortDictionary(dictionary)

			clearScreen()

			fmt.Println(`ID to EN Dictionary:`)
			for idx, val := range keys {
				fmt.Printf("%d. %s: %s\n", idx+1, convertToTitleCase(val), convertToTitleCase(dictionary[val]))
			}
			fmt.Printf("\n")

		case 0: // Exit
			os.Exit(0)

		default:
			fmt.Printf("Invalid input\n\n")
		}
	}
}

func convertToTitleCase(str string) string {
	caser := cases.Title(language.AmericanEnglish)

	return caser.String(str)
}

func sortDictionary(dictionary map[string]string) (keys []string) {
	keys = make([]string, 0, len(dictionary))

	for k := range dictionary {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

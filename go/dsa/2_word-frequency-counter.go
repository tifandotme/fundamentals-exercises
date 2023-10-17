/*
Word Frequency Counter
Description Given a text document, and you want to count the frequency of each word in the text, should be incase sesitive and ignoring symbol.

Example
Input
The quick brown fox jumps over the lazy dog. The dog barks loudly.

Output
the: 3
quick: 1
brown: 1
fox: 1
jumps: 1
over: 1
lazy: 1
dog: 2
barks: 1
loudly: 1
*/

package exercise

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func WordFrequencyCounter() {
	input := "The quick brown fox jumps over the lazy dog. The dog barks loudly."

	sanitized := regexp.MustCompile(`[^a-zA-Z0-9' ]+`).ReplaceAllString(input, "")

	lowercased := strings.ToLower(sanitized)

	output := map[string]int{}

	for _, val := range strings.Split(lowercased, " ") {
		output[val]++
	}

	printOutput(output)
}

func printOutput(output map[string]int) {
	keys := make([]string, 0, len(output))
	for k := range output {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("%s: %d\n", key, output[key])
	}
}

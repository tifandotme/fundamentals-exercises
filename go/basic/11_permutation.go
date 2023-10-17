package exercise

import "fmt"

// TODO: understand this
func generatePermutations(input string) []string {
	if len(input) == 1 {
		return []string{input}
	}

	permutations := []string{}

	for i, c := range input {
		remaining := input[:i] + input[i+1:]
		subPermutations := generatePermutations(remaining) // bc

		for _, p := range subPermutations {
			permutations = append(permutations, string(c)+p)
		}
	}

	return permutations
}

func Permutation() {
	input := "abc"
	permutations := generatePermutations(input)
	for _, permutation := range permutations {
		fmt.Println(permutation)
	}
}

package exercise

import "fmt"

func hasDuplicatesBruteForce(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				return true
			}
		}
	}
	return false
}

func hasDuplicatesAlt(arr []int) bool {
	log := map[int]bool{}

	for _, v := range arr {
		_, ok := log[v]

		if ok {
			return true
		}

		log[v] = true
	}

	return false
}

func AtleastOofN() {
	arr := []int{1, 3, 5, 7, 10, 9}
	fmt.Println("Has Duplicates (Brute Force):", hasDuplicatesBruteForce(arr))
	fmt.Println("Has Duplicates (Alt):", hasDuplicatesAlt(arr))
}

package exercise

import "fmt"

func findMaxElement(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	maxElement := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > maxElement {
			maxElement = arr[i]
		}
	}
	return maxElement
}

func FixTheEdgeCase() {
	// var arr []int

	arr := []int{5}

	fmt.Println("Maximum Element:", findMaxElement(arr))
}

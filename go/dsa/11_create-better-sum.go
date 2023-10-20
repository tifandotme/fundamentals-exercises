package exercise

import "fmt"

func recursiveSum(arr []int, n int) int {
	if n == 0 {
		return 0
	}
	return arr[n-1] + recursiveSum(arr, n-1)
}

func iterativeSum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func CreateBetterSum() {
	arr := []int{1, 3, 5, 7, 9}
	fmt.Println("Sum of elements (Recursive Sum):", recursiveSum(arr, len(arr)))
	fmt.Println("Sum of elements (Iterative Sum):", iterativeSum(arr))
}

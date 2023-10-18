package exercise

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(char string) {
	*s = append(*s, char)
}

func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}
}

func (s *Stack) Print() {
	fmt.Println(strings.Join(*s, ""))
}

func PrefixNotation() {
	// Example 1
	// input := "+ 1 2"

	// Example 2
	// input := "+ 1 * 2 + 3 4"
	input := "/ 1 * 2 + 3 4"

	// input := "+ + 1 * 2 3 4"

	// input := "+ 3 4"

	fmt.Println(input)

	input = strings.ReplaceAll(input, " ", "")

	bucket := Stack{}

	for i := len(input) - 1; i >= 0; i-- {
		if isOperator(string(input[i])) {
			leftOperand := toInteger(bucket.Pop())
			rightOperand := toInteger(bucket.Pop())
			result := 0

			switch input[i] {
			case '+':
				result = leftOperand + rightOperand
			case '-':
				result = leftOperand - rightOperand
			case '*':
				result = leftOperand * rightOperand
			case '/':
				result = rightOperand / leftOperand
			}

			bucket.Push(toString(result))
		} else {
			bucket.Push(string(input[i]))
		}
	}

	bucket.Print()
}

func isOperator(char string) bool {
	return regexp.MustCompile(`^[-+*\/]+`).MatchString(char)
}

func toInteger(char string) int {
	num, _ := strconv.Atoi(char)

	return num
}

func toString(num int) string {
	return strconv.Itoa(num)
}

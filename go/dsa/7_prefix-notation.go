/*
Evaluating Polish Notation
In this exercise, you are tasked with creating a program that can evaluate a mathematical expression expressed using the Polish Notation.

description

Usually, we evaluate mathematical expressions using an infix notation, such as: 1 + 2

The operators, in this case, the plus (+) operator, are placed between the operands. So, when we evaluate the expression, it results in: one plus two equals three.

The challenge arises when we need to evaluate expressions with multiple operators that have different orders of operations.

For example: 1 + 2 * (3 + 4)

The convention is to perform operations in this order:

Parentheses
Exponentiation
Multiplication and Division
Addition and Subtraction
However, there was a Polish mathematician (hence the name) who had an idea to use a notation that eliminates the need for parentheses and provides a clear order of operations: the Polish Notation.

Polish notation offers another way to evaluate mathematical expressions. It represents a classic challenge when learning data structures. This notation is also known as prefix notation.

In prefix notation, the expression 1 + 2 is written as + 1 2.

A more complex expression like 1 + 2 * (3 + 4) would be written as: + 1 * 2 + 3 4

Here is the link for more explanation regarding this notation.

The input will be a string in Polish Notation, which contains:

mathematical operator symbols such as: +, -, *, /
positive and negative numbers between -2147483648 and 2147483647 such as: 1, -2, 3, 5000
the operators and operands will always be separated by a space ( )
Your output should be the result of the Polish Notation evaluation.

Examples

Example 1

Input: + 1 2

Output: 3

Example 2

Input: + 1 * 2 + 3 4

Output: 15
*/

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

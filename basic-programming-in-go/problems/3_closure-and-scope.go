/*
Objectives
Understand how variable scope and closure works

Description
We want to create a function, which returns an array of functions, which return their index in the array.

Examples:

package main

import "fmt"

func main() {
    callbacks := CreateCallbacks(5)

    fmt.Println(callbacks[0]()) // must return 0
    fmt.Println(callbacks[1]()) // must return 1
    fmt.Println(callbacks[2]()) // must return 2
    fmt.Println(callbacks[3]()) // must return 3
    fmt.Println(callbacks[4]()) // must return 4
}
*/

package problems

import "fmt"

func createCallback(num int) []func() int {
	arr := make([]func() int, num)

	for i := 0; i < num; i++ {
		index := i

		arr[i] = func() int {
			return index
		}
	}

	return arr
}

func ClosureAndScope() {
	callbacks := createCallback(5)

	fmt.Println(callbacks[0]()) // must return 0
	fmt.Println(callbacks[1]()) // must return 1
	fmt.Println(callbacks[2]()) // must return 2
	fmt.Println(callbacks[3]()) // must return 3
	fmt.Println(callbacks[4]()) // must return 4
}

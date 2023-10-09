package main

import "digiwiki/problems"

func main() {
	problems.TicTacToe()
}

/*
In Go, function names that start with an uppercase letter will be exported,
while function names that start with a lowercase letter will not be exported.

Exported functions can be accessed from other packages,
while unexported functions can only be accessed from the same package.
*/

package main

import (
	"fmt"
	"strings"
)
/*board is a slice of 3 rows and each row is a slice of 3 rows and each row is a 
slice of 3 strings */


func main() {
	// Create a tic-tac-toe game ``
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X" //row 0 col 0
	board[2][2] = "O" //row 2 col 2
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

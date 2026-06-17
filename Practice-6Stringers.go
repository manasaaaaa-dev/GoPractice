// Problem 2 — Medium

// Create a Point struct with fields X and Y (both int). Implement Stringer so that printing a Point gives:
// (3,7)

package main

import "fmt"

type Point struct {
	X int 
	Y int
}

func (p Point) String () string {
	return fmt.Sprintf("(%v,%v)",p.X,p.Y)
}

func main() {
	a:=Point{3,7}
	fmt.Println(a)
}
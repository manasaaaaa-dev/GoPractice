package main

import "fmt"
// by using the make keyword we can create dynamic arrays

func main() {
	a := make([]int, 5)
	//makes a zero sized array 
	printSlice("a", a)

	b := make([]int, 0, 5)
	//To specify capacity we need to add another field 
	printSlice("b", b)

	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
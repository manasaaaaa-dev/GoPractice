package main

import "fmt"
//the zero value of the slice is nil
//a nil slice has a length = capacity = 0  
func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

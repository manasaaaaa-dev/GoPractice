package main

import "fmt"

func main() {
	a, b := 100, 200
	p := &a
	fmt.Println(*p) 
	fmt.Println(*p) 
	*p = 999
	fmt.Println(b)  
}
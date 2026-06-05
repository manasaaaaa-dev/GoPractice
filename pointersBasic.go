package main

import "fmt"

func main() {
	name := "Go"
	p := &name
	fmt.Println(*p)  
	*p = "Golang"    
	fmt.Println(name) 
}
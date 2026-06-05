package main

import "fmt"

func main() {
	names := [4]string{
		"manasa",
		"veda",
		"jyoshitha",
		"shruthi",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "check"
	fmt.Println(a, b)
	fmt.Println(names)
	//they will change the original array also 
}

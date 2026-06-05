package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]//normal way 
	fmt.Println(s)

	s = s[:2]// indices 0,1 are printed  
		fmt.Println(s)

	s = s[1:]// defaults -> omits the rest of the array 
	fmt.Println(s)
}

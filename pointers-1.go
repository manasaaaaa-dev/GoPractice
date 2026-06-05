package main
import "fmt"

func main() {
	a,b := 100,200
	p:=&a
	fmt.Println(*p)//pointing to a 
	p = &b
	fmt.Println(*p)//pointing to b 
}
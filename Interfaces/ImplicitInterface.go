package main

/*A type implements an interface by implementing its methods. 
There is no explicit declaration of intent, no "implements" keyword.

Implicit interfaces decouple the definition of an interface from its implementation, 
which could then appear in any package without prearrangement.*/


import "fmt"

type interface {
	S string
}

type T struct {
	S string
}
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}

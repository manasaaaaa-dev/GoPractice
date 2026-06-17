package main
import "fmt"
//It's an interface with zero methods. Since every type in Go automatically satisfies it, 
// it can hold anything , be it a number, a string, a struct, whatever.
func main() {
	var i interface{}
	describe(i)
	i = 42
	describe(i)
	i="hello"
	describe(i)

}

func describe (i interface{}) {
	fmt.Println("(%v,%T)\n", i ,i)
}
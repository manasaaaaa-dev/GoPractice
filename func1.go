//// 2. Create a function add(x int, y int) that returns the sum of two numbers.
package main
import "fmt"

func add (x int , y int ) int {
	return x+y
}

func main(){
	fmt.Println(add(5,9))
}

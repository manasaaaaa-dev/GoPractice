// Level 2 — Medium

//4. Write a function that returns two values — the sum and product of two integers.
package main
import "fmt"

func sumProd (x int , y int) (int, int) {
	sum := x+y
	product := x*y
	return sum , product
}
func main() {
	fmt.Println(sumProd(4,8))
}








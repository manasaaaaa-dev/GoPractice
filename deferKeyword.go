package main
import "fmt"
func main() {
	defer fmt.Println("turn off the lights")//executes at the last even though it is written at the begining
	fmt.Println("turn off the geyser")
	fmt.Println("fold clothes")
	fmt.Println("take keys")
}
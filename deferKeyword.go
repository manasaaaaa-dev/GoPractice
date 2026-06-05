package main
import "fmt"
func main() {
	defer fmt.Println("turn off the lights")
	fmt.Println("turn off the geyser")
	fmt.Println("fold clothes")
	fmt.Println("take keys")
}
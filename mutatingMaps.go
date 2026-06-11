package main 
import "fmt"
func main() {
	var mcq = make(map[string]string )

	mcq["op-1"] = "A"
	fmt.Println("The value of first option is :", mcq["op-1"])

	mcq["op-2"] = "B"
	fmt.Println("The value of second option is ",mcq["op-2"])
}
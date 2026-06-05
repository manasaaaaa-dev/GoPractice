package main 
import "fmt"

func main() {
	rollNums := [5] int {2,3,8,5,34}

	var s []int = rollNums[1:4]//4 is excluded 
	fmt.Println(s)

	//[3 5 7]
}